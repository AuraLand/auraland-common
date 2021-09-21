package auraland

import (
	"encoding/json"
)

type Entity struct {
	Id              string   `json:"id,omitempty"`
	SpaceId         string   `json:"space_id,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	AuthorityUserId string   `json:"authority_user_id,omitempty"`
	DisplayName     string   `json:"display_name,omitempty"`
	AssetUrl        string   `json:"asset_url,omitempty"`
	AssetType       string   `json:"asset_type,omitempty"`
	Version         int32    `json:"version,omitempty"`
	Position        *V3      `json:"position,omitempty"`
	Rotation        *V3      `json:"rotation,omitempty"`
	Scale           *V3      `json:"scale,omitempty"`
	Metadata        string   `json:"metadata,omitempty"`
	CreatedAt       int64    `json:"created_at,omitempty"`
	UpdatedAt       int64    `json:"updated_at,omitempty"`
	Online          bool     `json:"online,omitempty"`
	GroupId         string   `json:"group_id,omitempty"`
	SegmentId       int64    `json:"segment_id,omitempty"`
	IsSegment       bool     `json:"is_segment,omitempty"`
	AssetID         string   `json:"asset_id,omitempty"`
	UpdatedFields   []string `json:"updated_fields,omitempty"`
}

type V3 struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
	Z float32 `json:"z,omitempty"`
}

/*
	add asset id to the entity
*/
func (entity *Entity) AddAssetID(id string) {
	entity.AssetID = id
}

/*
	remove asset id from the entity
*/
func (entity *Entity) RemoveAssetID() {
	entity.AssetID = ""
}

func (entity *Entity) AssignId(id string) {
	entity.Id = id
}

func (entity *Entity) ToString() string {
	return string(entity.ToByte())
}

func (entity *Entity) ToByte() []byte {
	raw, _ := json.Marshal(entity)
	return raw
}
