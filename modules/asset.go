package auraland

import (
	"encoding/json"
)

type Asset struct {
	Id               string `json:"id,omitempty"`
	Url              string `json:"url,omitempty"`
	Name             string `json:"name,omitempty"`
	ThumbnailUrl     string `json:"thumbnail_url,omitempty"`
	Type             string `json:"type,omitempty"`
	Category         string `json:"category,omitempty"`
	Metadata         string `json:"metadata,omitempty"`
	CreatedAt        int64  `json:"created_at,omitempty"`
	UpdatedAt        int64  `json:"updated_at,omitempty"`
	Version          int32  `json:"version,omitempty"`
	ThumbnailVersion int32  `json:"thumbnail_version,omitempty"`
}

/*
	convert asset to []byte
*/
func (asset *Asset) ToByte() []byte {
	raw, _ := json.Marshal(asset)
	return raw
}

/*
	convert asset to string
*/
func (asset *Asset) ToString() string {
	return string(asset.ToByte())
}

/*
	assign id to asset
*/
func (asset *Asset) AssignId(id string) {
	asset.Id = id
}