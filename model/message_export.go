// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

type MessageExport struct {
	ChannelId          *string
	ChannelDisplayName *string
	ChannelType        *string

	UserId    *string
	UserEmail *string
	Username  *string

	PostId       *string
	PostCreateAt *int64
	PostMessage  *string
	PostType     *string
	PostFileIds  StringArray
}
