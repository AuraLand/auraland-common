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

/*
	Description: Update a space according to the input info
	Param: info: space info for the space
	return:
*/
func (asset *Asset) Update(info map[string]interface{}) {
	if value, ok := info["url"]; ok {
		asset.Url = value.(string)
	}
	if value, ok := info["name"]; ok {
		asset.Name = value.(string)
	}
	if value, ok := info["thumbnail_url"]; ok {
		asset.ThumbnailUrl = value.(string)
	}
	if value, ok := info["type"]; ok {
		asset.Type = value.(string)
	}
	if value, ok := info["category"]; ok {
		asset.Category = value.(string)
	}
	if value, ok := info["metadata"]; ok {
		asset.Metadata = value.(string)
	}
	if value, ok := info["created_at"]; ok {
		asset.CreatedAt = value.(int64)
	}
	if value, ok := info["updated_at"]; ok {
		asset.UpdatedAt = value.(int64)
	}
	if value, ok := info["version"]; ok {
		asset.Version = value.(int32)
	}
	if value, ok := info["thumbnail_version"]; ok {
		asset.ThumbnailVersion = value.(int32)
	}
}