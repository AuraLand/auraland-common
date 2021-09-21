package auraland

import (
	"encoding/json"
)


type Avatar struct {
	Id             string `json:"id,omitempty"`
	Url            string `json:"url,omitempty"`
	Name           string `json:"name,omitempty"`
	ThumbnailUrl   string `json:"thumbnail_url,omitempty"`
	ThumbnailHdUrl string `json:"thumbnail_hd_url,omitempty"`
	Metadata       string `json:"metadata,omitempty"`
	Version        int32  `json:"version,omitempty"`
}

/*
	Description: RPC function for loading avatar info into storage engine
	Param: avatarID: the avatar ID to acquire
	return: error and the stored avatar
	TODO: add helper method to pre-load avatar
*/
//func RpcLoadAvatar(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, avatarID string) (string, error) {
//	return "", nil
//}

/*
	convert avatar to []byte
*/
func (avatar *Avatar) ToByte() []byte {
	raw, _ := json.Marshal(avatar)
	return raw
}

/*
	convert avatar to string
*/
func (avatar *Avatar) ToString() string {
	return string(avatar.ToByte())
}

func (avatar *Avatar) AssignId(id string) {
	avatar.Id = id
}
