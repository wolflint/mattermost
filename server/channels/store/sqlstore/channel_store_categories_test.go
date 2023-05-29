// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/server/v8/channels/store/storetest"
)

func TestChannelStoreCategories(t *testing.T) {
	StoreTestWithSqlStore(t, storetest.TestChannelStoreCategories)
	StoreTestWithSqlStore(t, storetest.TestChannelStoreCategoriesWithAppsCategory)
}