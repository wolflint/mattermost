// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make pluginapi"
// DO NOT EDIT

package plugin

import (
	"errors"
	"io"
	"reflect"

	"github.com/mattermost/mattermost-server/v6/model"
)

type OnConfigurationChangeIFace interface {
	OnConfigurationChange() error
}

type ExecuteCommandIFace interface {
	ExecuteCommand(c *Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError)
}

type UserHasBeenCreatedIFace interface {
	UserHasBeenCreated(c *Context, user *model.User)
}

type UserWillLogInIFace interface {
	UserWillLogIn(c *Context, user *model.User) string
}

type UserHasLoggedInIFace interface {
	UserHasLoggedIn(c *Context, user *model.User)
}

type MessageWillBePostedIFace interface {
	MessageWillBePosted(c *Context, post *model.Post) (*model.Post, string)
}

type MessageWillBeUpdatedIFace interface {
	MessageWillBeUpdated(c *Context, newPost, oldPost *model.Post) (*model.Post, string)
}

type MessageHasBeenPostedIFace interface {
	MessageHasBeenPosted(c *Context, post *model.Post)
}

type MessageHasBeenUpdatedIFace interface {
	MessageHasBeenUpdated(c *Context, newPost, oldPost *model.Post)
}

type MessageWillBeConsumedIFace interface {
	MessageWillBeConsumed(post *model.Post) (*model.Post)
}

type ChannelHasBeenCreatedIFace interface {
	ChannelHasBeenCreated(c *Context, channel *model.Channel)
}

type UserHasJoinedChannelIFace interface {
	UserHasJoinedChannel(c *Context, channelMember *model.ChannelMember, actor *model.User)
}

type UserHasLeftChannelIFace interface {
	UserHasLeftChannel(c *Context, channelMember *model.ChannelMember, actor *model.User)
}

type UserHasJoinedTeamIFace interface {
	UserHasJoinedTeam(c *Context, teamMember *model.TeamMember, actor *model.User)
}

type UserHasLeftTeamIFace interface {
	UserHasLeftTeam(c *Context, teamMember *model.TeamMember, actor *model.User)
}

type FileWillBeUploadedIFace interface {
	FileWillBeUploaded(c *Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string)
}

type ReactionHasBeenAddedIFace interface {
	ReactionHasBeenAdded(c *Context, reaction *model.Reaction)
}

type ReactionHasBeenRemovedIFace interface {
	ReactionHasBeenRemoved(c *Context, reaction *model.Reaction)
}

type OnPluginClusterEventIFace interface {
	OnPluginClusterEvent(c *Context, ev model.PluginClusterEvent)
}

type OnWebSocketConnectIFace interface {
	OnWebSocketConnect(webConnID, userID string)
}

type OnWebSocketDisconnectIFace interface {
	OnWebSocketDisconnect(webConnID, userID string)
}

type WebSocketMessageHasBeenPostedIFace interface {
	WebSocketMessageHasBeenPosted(webConnID, userID string, req *model.WebSocketRequest)
}

type RunDataRetentionIFace interface {
	RunDataRetention(nowTime, batchSize int64) (int64, error)
}

type OnInstallIFace interface {
	OnInstall(c *Context, event model.OnInstallEvent) error
}

type OnSendDailyTelemetryIFace interface {
	OnSendDailyTelemetry()
}

type OnCloudLimitsUpdatedIFace interface {
	OnCloudLimitsUpdated(limits *model.ProductLimits)
}

type UserHasPermissionToCollectionIFace interface {
	UserHasPermissionToCollection(c *Context, userID string, collectionType, collectionId string, permission *model.Permission) (bool, error)
}

type GetAllCollectionIDsForUserIFace interface {
	GetAllCollectionIDsForUser(c *Context, userID, collectionType string) ([]string, error)
}

type GetAllUserIdsForCollectionIFace interface {
	GetAllUserIdsForCollection(c *Context, collectionType, collectionID string) ([]string, error)
}

type GetTopicRedirectIFace interface {
	GetTopicRedirect(c *Context, topicType, topicID string) (string, error)
}

type GetCollectionMetadataByIdsIFace interface {
	GetCollectionMetadataByIds(c *Context, collectionType string, collectionIds []string) (map[string]*model.CollectionMetadata, error)
}

type GetTopicMetadataByIdsIFace interface {
	GetTopicMetadataByIds(c *Context, topicType string, topicIds []string) (map[string]*model.TopicMetadata, error)
}

type HooksAdapter struct {
	implemented  map[int]struct{}
	productHooks any
}

func NewAdapter(productHooks any) (*HooksAdapter, error) {
	a := &HooksAdapter{
		implemented:  make(map[int]struct{}),
		productHooks: productHooks,
	}
	var tt reflect.Type
	ft := reflect.TypeOf(productHooks)

	// Assessing the type of the productHooks if it individually implements OnConfigurationChange interface.
	tt = reflect.TypeOf((*OnConfigurationChangeIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnConfigurationChangeID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnConfigurationChange"); ok {
		return nil, errors.New("hook has OnConfigurationChange method but does not implement plugin.OnConfigurationChange interface")
	}

	// Assessing the type of the productHooks if it individually implements ExecuteCommand interface.
	tt = reflect.TypeOf((*ExecuteCommandIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[ExecuteCommandID] = struct{}{}
	} else if _, ok := ft.MethodByName("ExecuteCommand"); ok {
		return nil, errors.New("hook has ExecuteCommand method but does not implement plugin.ExecuteCommand interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasBeenCreated interface.
	tt = reflect.TypeOf((*UserHasBeenCreatedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasBeenCreatedID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasBeenCreated"); ok {
		return nil, errors.New("hook has UserHasBeenCreated method but does not implement plugin.UserHasBeenCreated interface")
	}

	// Assessing the type of the productHooks if it individually implements UserWillLogIn interface.
	tt = reflect.TypeOf((*UserWillLogInIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserWillLogInID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserWillLogIn"); ok {
		return nil, errors.New("hook has UserWillLogIn method but does not implement plugin.UserWillLogIn interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasLoggedIn interface.
	tt = reflect.TypeOf((*UserHasLoggedInIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasLoggedInID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasLoggedIn"); ok {
		return nil, errors.New("hook has UserHasLoggedIn method but does not implement plugin.UserHasLoggedIn interface")
	}

	// Assessing the type of the productHooks if it individually implements MessageWillBePosted interface.
	tt = reflect.TypeOf((*MessageWillBePostedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[MessageWillBePostedID] = struct{}{}
	} else if _, ok := ft.MethodByName("MessageWillBePosted"); ok {
		return nil, errors.New("hook has MessageWillBePosted method but does not implement plugin.MessageWillBePosted interface")
	}

	// Assessing the type of the productHooks if it individually implements MessageWillBeUpdated interface.
	tt = reflect.TypeOf((*MessageWillBeUpdatedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[MessageWillBeUpdatedID] = struct{}{}
	} else if _, ok := ft.MethodByName("MessageWillBeUpdated"); ok {
		return nil, errors.New("hook has MessageWillBeUpdated method but does not implement plugin.MessageWillBeUpdated interface")
	}

	// Assessing the type of the productHooks if it individually implements MessageHasBeenPosted interface.
	tt = reflect.TypeOf((*MessageHasBeenPostedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[MessageHasBeenPostedID] = struct{}{}
	} else if _, ok := ft.MethodByName("MessageHasBeenPosted"); ok {
		return nil, errors.New("hook has MessageHasBeenPosted method but does not implement plugin.MessageHasBeenPosted interface")
	}

	// Assessing the type of the productHooks if it individually implements MessageHasBeenUpdated interface.
	tt = reflect.TypeOf((*MessageHasBeenUpdatedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[MessageHasBeenUpdatedID] = struct{}{}
	} else if _, ok := ft.MethodByName("MessageHasBeenUpdated"); ok {
		return nil, errors.New("hook has MessageHasBeenUpdated method but does not implement plugin.MessageHasBeenUpdated interface")
	}

	// Assessing the type of the productHooks if it individually implements MessageWillBeConsumed interface.
	tt = reflect.TypeOf((*MessageWillBeConsumedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[MessageWillBeConsumedID] = struct{}{}
	} else if _, ok := ft.MethodByName("MessageWillBeConsumed"); ok {
		return nil, errors.New("hook has MessageWillBeConsumed method but does not implement plugin.MessageWillBeConsumed interface")
	}

	// Assessing the type of the productHooks if it individually implements ChannelHasBeenCreated interface.
	tt = reflect.TypeOf((*ChannelHasBeenCreatedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[ChannelHasBeenCreatedID] = struct{}{}
	} else if _, ok := ft.MethodByName("ChannelHasBeenCreated"); ok {
		return nil, errors.New("hook has ChannelHasBeenCreated method but does not implement plugin.ChannelHasBeenCreated interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasJoinedChannel interface.
	tt = reflect.TypeOf((*UserHasJoinedChannelIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasJoinedChannelID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasJoinedChannel"); ok {
		return nil, errors.New("hook has UserHasJoinedChannel method but does not implement plugin.UserHasJoinedChannel interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasLeftChannel interface.
	tt = reflect.TypeOf((*UserHasLeftChannelIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasLeftChannelID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasLeftChannel"); ok {
		return nil, errors.New("hook has UserHasLeftChannel method but does not implement plugin.UserHasLeftChannel interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasJoinedTeam interface.
	tt = reflect.TypeOf((*UserHasJoinedTeamIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasJoinedTeamID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasJoinedTeam"); ok {
		return nil, errors.New("hook has UserHasJoinedTeam method but does not implement plugin.UserHasJoinedTeam interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasLeftTeam interface.
	tt = reflect.TypeOf((*UserHasLeftTeamIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasLeftTeamID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasLeftTeam"); ok {
		return nil, errors.New("hook has UserHasLeftTeam method but does not implement plugin.UserHasLeftTeam interface")
	}

	// Assessing the type of the productHooks if it individually implements FileWillBeUploaded interface.
	tt = reflect.TypeOf((*FileWillBeUploadedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[FileWillBeUploadedID] = struct{}{}
	} else if _, ok := ft.MethodByName("FileWillBeUploaded"); ok {
		return nil, errors.New("hook has FileWillBeUploaded method but does not implement plugin.FileWillBeUploaded interface")
	}

	// Assessing the type of the productHooks if it individually implements ReactionHasBeenAdded interface.
	tt = reflect.TypeOf((*ReactionHasBeenAddedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[ReactionHasBeenAddedID] = struct{}{}
	} else if _, ok := ft.MethodByName("ReactionHasBeenAdded"); ok {
		return nil, errors.New("hook has ReactionHasBeenAdded method but does not implement plugin.ReactionHasBeenAdded interface")
	}

	// Assessing the type of the productHooks if it individually implements ReactionHasBeenRemoved interface.
	tt = reflect.TypeOf((*ReactionHasBeenRemovedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[ReactionHasBeenRemovedID] = struct{}{}
	} else if _, ok := ft.MethodByName("ReactionHasBeenRemoved"); ok {
		return nil, errors.New("hook has ReactionHasBeenRemoved method but does not implement plugin.ReactionHasBeenRemoved interface")
	}

	// Assessing the type of the productHooks if it individually implements OnPluginClusterEvent interface.
	tt = reflect.TypeOf((*OnPluginClusterEventIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnPluginClusterEventID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnPluginClusterEvent"); ok {
		return nil, errors.New("hook has OnPluginClusterEvent method but does not implement plugin.OnPluginClusterEvent interface")
	}

	// Assessing the type of the productHooks if it individually implements OnWebSocketConnect interface.
	tt = reflect.TypeOf((*OnWebSocketConnectIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnWebSocketConnectID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnWebSocketConnect"); ok {
		return nil, errors.New("hook has OnWebSocketConnect method but does not implement plugin.OnWebSocketConnect interface")
	}

	// Assessing the type of the productHooks if it individually implements OnWebSocketDisconnect interface.
	tt = reflect.TypeOf((*OnWebSocketDisconnectIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnWebSocketDisconnectID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnWebSocketDisconnect"); ok {
		return nil, errors.New("hook has OnWebSocketDisconnect method but does not implement plugin.OnWebSocketDisconnect interface")
	}

	// Assessing the type of the productHooks if it individually implements WebSocketMessageHasBeenPosted interface.
	tt = reflect.TypeOf((*WebSocketMessageHasBeenPostedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[WebSocketMessageHasBeenPostedID] = struct{}{}
	} else if _, ok := ft.MethodByName("WebSocketMessageHasBeenPosted"); ok {
		return nil, errors.New("hook has WebSocketMessageHasBeenPosted method but does not implement plugin.WebSocketMessageHasBeenPosted interface")
	}

	// Assessing the type of the productHooks if it individually implements RunDataRetention interface.
	tt = reflect.TypeOf((*RunDataRetentionIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[RunDataRetentionID] = struct{}{}
	} else if _, ok := ft.MethodByName("RunDataRetention"); ok {
		return nil, errors.New("hook has RunDataRetention method but does not implement plugin.RunDataRetention interface")
	}

	// Assessing the type of the productHooks if it individually implements OnInstall interface.
	tt = reflect.TypeOf((*OnInstallIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnInstallID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnInstall"); ok {
		return nil, errors.New("hook has OnInstall method but does not implement plugin.OnInstall interface")
	}

	// Assessing the type of the productHooks if it individually implements OnSendDailyTelemetry interface.
	tt = reflect.TypeOf((*OnSendDailyTelemetryIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnSendDailyTelemetryID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnSendDailyTelemetry"); ok {
		return nil, errors.New("hook has OnSendDailyTelemetry method but does not implement plugin.OnSendDailyTelemetry interface")
	}

	// Assessing the type of the productHooks if it individually implements OnCloudLimitsUpdated interface.
	tt = reflect.TypeOf((*OnCloudLimitsUpdatedIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[OnCloudLimitsUpdatedID] = struct{}{}
	} else if _, ok := ft.MethodByName("OnCloudLimitsUpdated"); ok {
		return nil, errors.New("hook has OnCloudLimitsUpdated method but does not implement plugin.OnCloudLimitsUpdated interface")
	}

	// Assessing the type of the productHooks if it individually implements UserHasPermissionToCollection interface.
	tt = reflect.TypeOf((*UserHasPermissionToCollectionIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[UserHasPermissionToCollectionID] = struct{}{}
	} else if _, ok := ft.MethodByName("UserHasPermissionToCollection"); ok {
		return nil, errors.New("hook has UserHasPermissionToCollection method but does not implement plugin.UserHasPermissionToCollection interface")
	}

	// Assessing the type of the productHooks if it individually implements GetAllCollectionIDsForUser interface.
	tt = reflect.TypeOf((*GetAllCollectionIDsForUserIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[GetAllCollectionIDsForUserID] = struct{}{}
	} else if _, ok := ft.MethodByName("GetAllCollectionIDsForUser"); ok {
		return nil, errors.New("hook has GetAllCollectionIDsForUser method but does not implement plugin.GetAllCollectionIDsForUser interface")
	}

	// Assessing the type of the productHooks if it individually implements GetAllUserIdsForCollection interface.
	tt = reflect.TypeOf((*GetAllUserIdsForCollectionIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[GetAllUserIdsForCollectionID] = struct{}{}
	} else if _, ok := ft.MethodByName("GetAllUserIdsForCollection"); ok {
		return nil, errors.New("hook has GetAllUserIdsForCollection method but does not implement plugin.GetAllUserIdsForCollection interface")
	}

	// Assessing the type of the productHooks if it individually implements GetTopicRedirect interface.
	tt = reflect.TypeOf((*GetTopicRedirectIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[GetTopicRedirectID] = struct{}{}
	} else if _, ok := ft.MethodByName("GetTopicRedirect"); ok {
		return nil, errors.New("hook has GetTopicRedirect method but does not implement plugin.GetTopicRedirect interface")
	}

	// Assessing the type of the productHooks if it individually implements GetCollectionMetadataByIds interface.
	tt = reflect.TypeOf((*GetCollectionMetadataByIdsIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[GetCollectionMetadataByIdsID] = struct{}{}
	} else if _, ok := ft.MethodByName("GetCollectionMetadataByIds"); ok {
		return nil, errors.New("hook has GetCollectionMetadataByIds method but does not implement plugin.GetCollectionMetadataByIds interface")
	}

	// Assessing the type of the productHooks if it individually implements GetTopicMetadataByIds interface.
	tt = reflect.TypeOf((*GetTopicMetadataByIdsIFace)(nil)).Elem()

	if ft.Implements(tt) {
		a.implemented[GetTopicMetadataByIdsID] = struct{}{}
	} else if _, ok := ft.MethodByName("GetTopicMetadataByIds"); ok {
		return nil, errors.New("hook has GetTopicMetadataByIds method but does not implement plugin.GetTopicMetadataByIds interface")
	}

	return a, nil
}

func (a *HooksAdapter) OnConfigurationChange() error {
	if _, ok := a.implemented[OnConfigurationChangeID]; !ok {
		panic("product hooks must implement OnConfigurationChange")
	}

	return a.productHooks.(OnConfigurationChangeIFace).OnConfigurationChange()

}

func (a *HooksAdapter) ExecuteCommand(c *Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	if _, ok := a.implemented[ExecuteCommandID]; !ok {
		panic("product hooks must implement ExecuteCommand")
	}

	return a.productHooks.(ExecuteCommandIFace).ExecuteCommand(c, args)

}

func (a *HooksAdapter) UserHasBeenCreated(c *Context, user *model.User) {
	if _, ok := a.implemented[UserHasBeenCreatedID]; !ok {
		panic("product hooks must implement UserHasBeenCreated")
	}

	a.productHooks.(UserHasBeenCreatedIFace).UserHasBeenCreated(c, user)

}

func (a *HooksAdapter) UserWillLogIn(c *Context, user *model.User) string {
	if _, ok := a.implemented[UserWillLogInID]; !ok {
		panic("product hooks must implement UserWillLogIn")
	}

	return a.productHooks.(UserWillLogInIFace).UserWillLogIn(c, user)

}

func (a *HooksAdapter) UserHasLoggedIn(c *Context, user *model.User) {
	if _, ok := a.implemented[UserHasLoggedInID]; !ok {
		panic("product hooks must implement UserHasLoggedIn")
	}

	a.productHooks.(UserHasLoggedInIFace).UserHasLoggedIn(c, user)

}

func (a *HooksAdapter) MessageWillBePosted(c *Context, post *model.Post) (*model.Post, string) {
	if _, ok := a.implemented[MessageWillBePostedID]; !ok {
		panic("product hooks must implement MessageWillBePosted")
	}

	return a.productHooks.(MessageWillBePostedIFace).MessageWillBePosted(c, post)

}

func (a *HooksAdapter) MessageWillBeUpdated(c *Context, newPost, oldPost *model.Post) (*model.Post, string) {
	if _, ok := a.implemented[MessageWillBeUpdatedID]; !ok {
		panic("product hooks must implement MessageWillBeUpdated")
	}

	return a.productHooks.(MessageWillBeUpdatedIFace).MessageWillBeUpdated(c, newPost, oldPost)

}

func (a *HooksAdapter) MessageHasBeenPosted(c *Context, post *model.Post) {
	if _, ok := a.implemented[MessageHasBeenPostedID]; !ok {
		panic("product hooks must implement MessageHasBeenPosted")
	}

	a.productHooks.(MessageHasBeenPostedIFace).MessageHasBeenPosted(c, post)

}

func (a *HooksAdapter) MessageHasBeenUpdated(c *Context, newPost, oldPost *model.Post) {
	if _, ok := a.implemented[MessageHasBeenUpdatedID]; !ok {
		panic("product hooks must implement MessageHasBeenUpdated")
	}

	a.productHooks.(MessageHasBeenUpdatedIFace).MessageHasBeenUpdated(c, newPost, oldPost)

}

func (a *HooksAdapter) MessageWillBeConsumed(post *model.Post) (*model.Post) {
	if _, ok := a.implemented[MessageWillBeConsumedID]; !ok {
		panic("product hooks must implement MessageWillBeConsumed")
	}
	p := a.productHooks.(MessageWillBeConsumedIFace).MessageWillBeConsumed(post)
	return p
}

func (a *HooksAdapter) ChannelHasBeenCreated(c *Context, channel *model.Channel) {
	if _, ok := a.implemented[ChannelHasBeenCreatedID]; !ok {
		panic("product hooks must implement ChannelHasBeenCreated")
	}

	a.productHooks.(ChannelHasBeenCreatedIFace).ChannelHasBeenCreated(c, channel)

}

func (a *HooksAdapter) UserHasJoinedChannel(c *Context, channelMember *model.ChannelMember, actor *model.User) {
	if _, ok := a.implemented[UserHasJoinedChannelID]; !ok {
		panic("product hooks must implement UserHasJoinedChannel")
	}

	a.productHooks.(UserHasJoinedChannelIFace).UserHasJoinedChannel(c, channelMember, actor)

}

func (a *HooksAdapter) UserHasLeftChannel(c *Context, channelMember *model.ChannelMember, actor *model.User) {
	if _, ok := a.implemented[UserHasLeftChannelID]; !ok {
		panic("product hooks must implement UserHasLeftChannel")
	}

	a.productHooks.(UserHasLeftChannelIFace).UserHasLeftChannel(c, channelMember, actor)

}

func (a *HooksAdapter) UserHasJoinedTeam(c *Context, teamMember *model.TeamMember, actor *model.User) {
	if _, ok := a.implemented[UserHasJoinedTeamID]; !ok {
		panic("product hooks must implement UserHasJoinedTeam")
	}

	a.productHooks.(UserHasJoinedTeamIFace).UserHasJoinedTeam(c, teamMember, actor)

}

func (a *HooksAdapter) UserHasLeftTeam(c *Context, teamMember *model.TeamMember, actor *model.User) {
	if _, ok := a.implemented[UserHasLeftTeamID]; !ok {
		panic("product hooks must implement UserHasLeftTeam")
	}

	a.productHooks.(UserHasLeftTeamIFace).UserHasLeftTeam(c, teamMember, actor)

}

func (a *HooksAdapter) FileWillBeUploaded(c *Context, info *model.FileInfo, file io.Reader, output io.Writer) (*model.FileInfo, string) {
	if _, ok := a.implemented[FileWillBeUploadedID]; !ok {
		panic("product hooks must implement FileWillBeUploaded")
	}

	return a.productHooks.(FileWillBeUploadedIFace).FileWillBeUploaded(c, info, file, output)

}

func (a *HooksAdapter) ReactionHasBeenAdded(c *Context, reaction *model.Reaction) {
	if _, ok := a.implemented[ReactionHasBeenAddedID]; !ok {
		panic("product hooks must implement ReactionHasBeenAdded")
	}

	a.productHooks.(ReactionHasBeenAddedIFace).ReactionHasBeenAdded(c, reaction)

}

func (a *HooksAdapter) ReactionHasBeenRemoved(c *Context, reaction *model.Reaction) {
	if _, ok := a.implemented[ReactionHasBeenRemovedID]; !ok {
		panic("product hooks must implement ReactionHasBeenRemoved")
	}

	a.productHooks.(ReactionHasBeenRemovedIFace).ReactionHasBeenRemoved(c, reaction)

}

func (a *HooksAdapter) OnPluginClusterEvent(c *Context, ev model.PluginClusterEvent) {
	if _, ok := a.implemented[OnPluginClusterEventID]; !ok {
		panic("product hooks must implement OnPluginClusterEvent")
	}

	a.productHooks.(OnPluginClusterEventIFace).OnPluginClusterEvent(c, ev)

}

func (a *HooksAdapter) OnWebSocketConnect(webConnID, userID string) {
	if _, ok := a.implemented[OnWebSocketConnectID]; !ok {
		panic("product hooks must implement OnWebSocketConnect")
	}

	a.productHooks.(OnWebSocketConnectIFace).OnWebSocketConnect(webConnID, userID)

}

func (a *HooksAdapter) OnWebSocketDisconnect(webConnID, userID string) {
	if _, ok := a.implemented[OnWebSocketDisconnectID]; !ok {
		panic("product hooks must implement OnWebSocketDisconnect")
	}

	a.productHooks.(OnWebSocketDisconnectIFace).OnWebSocketDisconnect(webConnID, userID)

}

func (a *HooksAdapter) WebSocketMessageHasBeenPosted(webConnID, userID string, req *model.WebSocketRequest) {
	if _, ok := a.implemented[WebSocketMessageHasBeenPostedID]; !ok {
		panic("product hooks must implement WebSocketMessageHasBeenPosted")
	}

	a.productHooks.(WebSocketMessageHasBeenPostedIFace).WebSocketMessageHasBeenPosted(webConnID, userID, req)

}

func (a *HooksAdapter) RunDataRetention(nowTime, batchSize int64) (int64, error) {
	if _, ok := a.implemented[RunDataRetentionID]; !ok {
		panic("product hooks must implement RunDataRetention")
	}

	return a.productHooks.(RunDataRetentionIFace).RunDataRetention(nowTime, batchSize)

}

func (a *HooksAdapter) OnInstall(c *Context, event model.OnInstallEvent) error {
	if _, ok := a.implemented[OnInstallID]; !ok {
		panic("product hooks must implement OnInstall")
	}

	return a.productHooks.(OnInstallIFace).OnInstall(c, event)

}

func (a *HooksAdapter) OnSendDailyTelemetry() {
	if _, ok := a.implemented[OnSendDailyTelemetryID]; !ok {
		panic("product hooks must implement OnSendDailyTelemetry")
	}

	a.productHooks.(OnSendDailyTelemetryIFace).OnSendDailyTelemetry()

}

func (a *HooksAdapter) OnCloudLimitsUpdated(limits *model.ProductLimits) {
	if _, ok := a.implemented[OnCloudLimitsUpdatedID]; !ok {
		panic("product hooks must implement OnCloudLimitsUpdated")
	}

	a.productHooks.(OnCloudLimitsUpdatedIFace).OnCloudLimitsUpdated(limits)

}

func (a *HooksAdapter) UserHasPermissionToCollection(c *Context, userID string, collectionType, collectionId string, permission *model.Permission) (bool, error) {
	if _, ok := a.implemented[UserHasPermissionToCollectionID]; !ok {
		panic("product hooks must implement UserHasPermissionToCollection")
	}

	return a.productHooks.(UserHasPermissionToCollectionIFace).UserHasPermissionToCollection(c, userID, collectionType, collectionId, permission)

}

func (a *HooksAdapter) GetAllCollectionIDsForUser(c *Context, userID, collectionType string) ([]string, error) {
	if _, ok := a.implemented[GetAllCollectionIDsForUserID]; !ok {
		panic("product hooks must implement GetAllCollectionIDsForUser")
	}

	return a.productHooks.(GetAllCollectionIDsForUserIFace).GetAllCollectionIDsForUser(c, userID, collectionType)

}

func (a *HooksAdapter) GetAllUserIdsForCollection(c *Context, collectionType, collectionID string) ([]string, error) {
	if _, ok := a.implemented[GetAllUserIdsForCollectionID]; !ok {
		panic("product hooks must implement GetAllUserIdsForCollection")
	}

	return a.productHooks.(GetAllUserIdsForCollectionIFace).GetAllUserIdsForCollection(c, collectionType, collectionID)

}

func (a *HooksAdapter) GetTopicRedirect(c *Context, topicType, topicID string) (string, error) {
	if _, ok := a.implemented[GetTopicRedirectID]; !ok {
		panic("product hooks must implement GetTopicRedirect")
	}

	return a.productHooks.(GetTopicRedirectIFace).GetTopicRedirect(c, topicType, topicID)

}

func (a *HooksAdapter) GetCollectionMetadataByIds(c *Context, collectionType string, collectionIds []string) (map[string]*model.CollectionMetadata, error) {
	if _, ok := a.implemented[GetCollectionMetadataByIdsID]; !ok {
		panic("product hooks must implement GetCollectionMetadataByIds")
	}

	return a.productHooks.(GetCollectionMetadataByIdsIFace).GetCollectionMetadataByIds(c, collectionType, collectionIds)

}

func (a *HooksAdapter) GetTopicMetadataByIds(c *Context, topicType string, topicIds []string) (map[string]*model.TopicMetadata, error) {
	if _, ok := a.implemented[GetTopicMetadataByIdsID]; !ok {
		panic("product hooks must implement GetTopicMetadataByIds")
	}

	return a.productHooks.(GetTopicMetadataByIdsIFace).GetTopicMetadataByIds(c, topicType, topicIds)

}
