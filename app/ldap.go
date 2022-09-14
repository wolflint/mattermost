// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/shared/i18n"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

// SyncLdap starts an LDAP sync job.
// If includeRemovedMembers is true, then members who left or were removed from a team/channel will
// be re-added; otherwise, they will not be re-added.
func (a *App) SyncLdap(c request.CTX, includeRemovedMembers bool) {
	a.Srv().Go(func() {
		if license := a.Srv().License(c); license != nil && *license.Features.LDAP && *a.Config().LdapSettings.EnableSync {
			if ldapI := a.Ldap(); ldapI != nil {
				ldapI.StartSynchronizeJob(false, includeRemovedMembers)
			} else {
				c.Logger().Error("Not executing ldap sync because ldap is not available")
			}
		}
	})
}

func (a *App) TestLdap(c request.CTX) *model.AppError {
	license := a.Srv().License(c)
	if ldapI := a.Ldap(); ldapI != nil && license != nil && *license.Features.LDAP && (*a.Config().LdapSettings.Enable || *a.Config().LdapSettings.EnableSync) {
		if err := ldapI.RunTest(); err != nil {
			err.StatusCode = 500
			return err
		}
	} else {
		err := model.NewAppError("TestLdap", "ent.ldap.disabled.app_error", nil, "", http.StatusNotImplemented)
		return err
	}

	return nil
}

// GetLdapGroup retrieves a single LDAP group by the given LDAP group id.
func (a *App) GetLdapGroup(ldapGroupID string) (*model.Group, *model.AppError) {
	var group *model.Group

	if a.Ldap() != nil {
		var err *model.AppError
		group, err = a.Ldap().GetGroup(ldapGroupID)
		if err != nil {
			return nil, err
		}
	} else {
		ae := model.NewAppError("GetLdapGroup", "ent.ldap.app_error", map[string]any{"ldap_group_id": ldapGroupID}, "", http.StatusNotImplemented)
		return nil, ae
	}

	return group, nil
}

// GetAllLdapGroupsPage retrieves all LDAP groups under the configured base DN using the default or configured group
// filter.
func (a *App) GetAllLdapGroupsPage(page int, perPage int, opts model.LdapGroupSearchOpts) ([]*model.Group, int, *model.AppError) {
	var groups []*model.Group
	var total int

	if a.Ldap() != nil {
		var err *model.AppError
		groups, total, err = a.Ldap().GetAllGroupsPage(page, perPage, opts)
		if err != nil {
			return nil, 0, err
		}
	} else {
		ae := model.NewAppError("GetAllLdapGroupsPage", "ent.ldap.app_error", nil, "", http.StatusNotImplemented)
		return nil, 0, ae
	}

	return groups, total, nil
}

func (a *App) SwitchEmailToLdap(c request.CTX, email, password, code, ldapLoginId, ldapPassword string) (string, *model.AppError) {
	if a.Srv().License(c) != nil && !*a.Config().ServiceSettings.ExperimentalEnableAuthenticationTransfer {
		return "", model.NewAppError("emailToLdap", "api.user.email_to_ldap.not_available.app_error", nil, "", http.StatusForbidden)
	}

	user, err := a.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := a.CheckPasswordAndAllCriteria(user, password, code); err != nil {
		return "", err
	}

	if err := a.RevokeAllSessions(user.Id); err != nil {
		return "", err
	}

	ldapInterface := a.Ldap()
	if ldapInterface == nil {
		return "", model.NewAppError("SwitchEmailToLdap", "api.user.email_to_ldap.not_available.app_error", nil, "", http.StatusNotImplemented)
	}

	if err := ldapInterface.SwitchToLdap(user.Id, ldapLoginId, ldapPassword); err != nil {
		return "", err
	}

	a.Srv().Go(func() {
		if err := a.Srv().EmailService.SendSignInChangeEmail(user.Email, "AD/LDAP", user.Locale, a.GetSiteURL()); err != nil {
			c.Logger().Error("Could not send sign in method changed e-mail", mlog.Err(err))
		}
	})

	return "/login?extra=signin_change", nil
}

func (a *App) SwitchLdapToEmail(c request.CTX, ldapPassword, code, email, newPassword string) (string, *model.AppError) {
	if a.Srv().License(c) != nil && !*a.Config().ServiceSettings.ExperimentalEnableAuthenticationTransfer {
		return "", model.NewAppError("ldapToEmail", "api.user.ldap_to_email.not_available.app_error", nil, "", http.StatusForbidden)
	}

	user, err := a.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if user.AuthService != model.UserAuthServiceLdap {
		return "", model.NewAppError("SwitchLdapToEmail", "api.user.ldap_to_email.not_ldap_account.app_error", nil, "", http.StatusBadRequest)
	}

	ldapInterface := a.Ldap()
	if ldapInterface == nil || user.AuthData == nil {
		return "", model.NewAppError("SwitchLdapToEmail", "api.user.ldap_to_email.not_available.app_error", nil, "", http.StatusNotImplemented)
	}

	if err := ldapInterface.CheckPasswordAuthData(*user.AuthData, ldapPassword); err != nil {
		return "", err
	}

	if err := a.CheckUserMfa(user, code); err != nil {
		return "", err
	}

	if err := a.UpdatePassword(user, newPassword); err != nil {
		return "", err
	}

	if err := a.RevokeAllSessions(user.Id); err != nil {
		return "", err
	}

	T := i18n.GetUserTranslations(user.Locale)

	a.Srv().Go(func() {
		if err := a.Srv().EmailService.SendSignInChangeEmail(user.Email, T("api.templates.signin_change_email.body.method_email"), user.Locale, a.GetSiteURL()); err != nil {
			c.Logger().Error("Could not send sign in method changed e-mail", mlog.Err(err))
		}
	})

	return "/login?extra=signin_change", nil
}

func (a *App) MigrateIdLDAP(toAttribute string) *model.AppError {
	if ldapI := a.Ldap(); ldapI != nil {
		if err := ldapI.MigrateIDAttribute(toAttribute); err != nil {
			switch err := err.(type) {
			case *model.AppError:
				return err
			default:
				return model.NewAppError("IdMigrateLDAP", "ent.ldap_id_migrate.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
			}
		}
		return nil
	}
	return model.NewAppError("IdMigrateLDAP", "ent.ldap.disabled.app_error", nil, "", http.StatusNotImplemented)
}

func (a *App) writeLdapFile(filename string, fileData *multipart.FileHeader) *model.AppError {
	file, err := fileData.Open()
	if err != nil {
		return model.NewAppError("AddLdapCertificate", "api.admin.add_certificate.open.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return model.NewAppError("AddLdapCertificate", "api.admin.add_certificate.saving.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}

	err = a.Srv().platform.SetConfigFile(filename, data)
	if err != nil {
		return model.NewAppError("AddLdapCertificate", "api.admin.add_certificate.saving.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
	}

	return nil
}

func (a *App) AddLdapPublicCertificate(fileData *multipart.FileHeader) *model.AppError {
	if err := a.writeLdapFile(model.LdapPublicCertificateName, fileData); err != nil {
		return err
	}

	cfg := a.Config().Clone()
	*cfg.LdapSettings.PublicCertificateFile = model.LdapPublicCertificateName

	if err := cfg.IsValid(); err != nil {
		return err
	}

	a.UpdateConfig(func(dest *model.Config) { *dest = *cfg })

	return nil
}

func (a *App) AddLdapPrivateCertificate(fileData *multipart.FileHeader) *model.AppError {
	if err := a.writeLdapFile(model.LdapPrivateKeyName, fileData); err != nil {
		return err
	}

	cfg := a.Config().Clone()
	*cfg.LdapSettings.PrivateKeyFile = model.LdapPrivateKeyName

	if err := cfg.IsValid(); err != nil {
		return err
	}

	a.UpdateConfig(func(dest *model.Config) { *dest = *cfg })

	return nil
}

func (a *App) removeLdapFile(filename string) *model.AppError {
	if err := a.Srv().platform.RemoveConfigFile(filename); err != nil {
		return model.NewAppError("RemoveLdapFile", "api.admin.remove_certificate.delete.app_error", map[string]any{"Filename": filename}, "", http.StatusInternalServerError).Wrap(err)
	}
	return nil
}

func (a *App) RemoveLdapPublicCertificate() *model.AppError {
	if err := a.removeLdapFile(*a.Config().LdapSettings.PublicCertificateFile); err != nil {
		return err
	}

	cfg := a.Config().Clone()
	*cfg.LdapSettings.PublicCertificateFile = ""

	if err := cfg.IsValid(); err != nil {
		return err
	}

	a.UpdateConfig(func(dest *model.Config) { *dest = *cfg })

	return nil
}

func (a *App) RemoveLdapPrivateCertificate() *model.AppError {
	if err := a.removeLdapFile(*a.Config().LdapSettings.PrivateKeyFile); err != nil {
		return err
	}

	cfg := a.Config().Clone()
	*cfg.LdapSettings.PrivateKeyFile = ""

	if err := cfg.IsValid(); err != nil {
		return err
	}

	a.UpdateConfig(func(dest *model.Config) { *dest = *cfg })

	return nil
}
