package auraland

import (
	"encoding/json"
	"fmt"
)

type Space struct {
	Id               string   `json:"id"`
	CreatorId        string   `json:"creator_id,omitempty"`
	DisplayName      string   `json:"display_name"`
	Description      string   `json:"description,omitempty"`
	ThumbnailUrl     string   `json:"thumbnail_url,omitempty"`
	Lang             string   `json:"lang,omitempty"` // Language tag corresponding to the BCP 47 spec. This is important and used for Space search.
	Theme            string   `json:"theme,omitempty"`
	UtcOffsetMs      int64    `json:"utc_offset_ms,omitempty"` // Offset time in millisecond from UTC.
	Metadata         string   `json:"metadata,omitempty"`      // Group metadata information.
	Count            int64    `json:"count,omitempty"`         // Current number of users in this space.
	Private          bool     `json:"private,omitempty"`
	CreatedAt        int64    `json:"created_at,omitempty"`
	UpdatedAt        int64    `json:"updated_at,omitempty"`
	Url              string   `json:"url,omitempty"`
	Version          string   `json:"version,omitempty"`
	ThumbnailVersion int32    `json:"thumbnail_version,omitempty"`
	CreatorFullName  string   `json:"creator_full_name,omitempty"`
	UserIDs          []string `json:"user_ids,omitempty"`
	EntityIDs        []string `json:"entity_ids,omitempty"`
	AvatarIDs        []string `json:"avatar_ids,omitempty"`
	AssetIDs         []string `json:"asset_ids,omitempty"`
}

func (space *Space) AssignId(id string) {
	space.Id = id
}

/*
	Description: Update a space according to the input info
	Param: info: space info for the space
	return:
*/
func (space *Space) Update(info map[string]interface{}) {
	if value, ok := info["display_name"]; ok {
		space.DisplayName = value.(string)
	}
	if value, ok := info["description"]; ok {
		space.Description = value.(string)
	}
	if value, ok := info["thumbnail_url"]; ok {
		space.ThumbnailUrl = value.(string)
	}
	if value, ok := info["lang"]; ok {
		space.Lang = value.(string)
	}
	if value, ok := info["theme"]; ok {
		space.Theme = value.(string)
	}
	if value, ok := info["utc_offset_ms"]; ok {
		space.UtcOffsetMs = value.(int64)
	}
	if value, ok := info["metadata"]; ok {
		space.Metadata = value.(string)
	}
	if value, ok := info["count"]; ok {
		space.Count = value.(int64)
	}
	if value, ok := info["private"]; ok {
		space.Private = value.(bool)
	}
	if value, ok := info["updated_at"]; ok {
		space.UpdatedAt = value.(int64)
	}
	if value, ok := info["url"]; ok {
		space.Url = value.(string)
	}
	if value, ok := info["version"]; ok {
		space.Version = value.(string)
	}
	if value, ok := info["thumbnail_version"]; ok {
		space.ThumbnailVersion = value.(int32)
	}
	if value, ok := info["creator_fullname"]; ok {
		space.CreatorFullName = value.(string)
	}

}

/*
	member joined the space
*/
func (space *Space) JoinSpace(userId string) error {
	for _, id := range space.UserIDs {
		if id == userId {
			return fmt.Errorf("user: %+v is already in the space", userId)
		}
	}
	space.UserIDs = append(space.UserIDs, userId)
	return nil
}

/*
	member leave the space
*/
func (space *Space) LeaveSpace(ids []string) {
	for i := 0; i < len(space.UserIDs); i++ {
		item := space.UserIDs[i]
		for _, rem := range ids {
			if item == rem {
				space.UserIDs = append(space.UserIDs[:i], space.UserIDs[i+1:]...)
				i--
				break
			}
		}
	}
}

/*
	convert space to []byte
*/
func (space *Space) ToByte() []byte {
	raw, _ := json.Marshal(space)
	return raw
}

/*
	convert space to string
*/
func (space *Space) ToString() string {
	return string(space.ToByte())
}

/*
	add entity ids to the space
*/
func (space *Space) AddEntityIDs(ids []string) {
	checkMap := make(map[string]bool)
	entityList := append(space.EntityIDs, ids...)
	space.EntityIDs = nil
	for _, item := range entityList {
		if _, value := checkMap[item]; !value {
			checkMap[item] = true
			space.EntityIDs = append(space.EntityIDs, item)
		}
	}
}

/*
	remove entity ids from the space
*/
func (space *Space) RemoveEntityIDs(ids []string) {
	for i := 0; i < len(space.EntityIDs); i++ {
		item := space.EntityIDs[i]
		for _, rem := range ids {
			if item == rem {
				space.EntityIDs = append(space.EntityIDs[:i], space.EntityIDs[i+1:]...)
				i--
				break
			}
		}
	}
}

/*
	add asset ids to the space
*/
func (space *Space) AddAssetIDs(ids []string) {
	checkMap := make(map[string]bool)
	assetList := append(space.AssetIDs, ids...)
	space.AssetIDs = nil
	for _, item := range assetList {
		if _, value := checkMap[item]; !value {
			checkMap[item] = true
			space.AssetIDs = append(space.AssetIDs, item)
		}
	}
}

/*
	remove asset ids from the space
*/
func (space *Space) RemoveAssetIDs(ids []string) {
	for i := 0; i < len(space.AssetIDs); i++ {
		item := space.AssetIDs[i]
		for _, rem := range ids {
			if item == rem {
				space.AssetIDs = append(space.AssetIDs[:i], space.AssetIDs[i+1:]...)
				i--
				break
			}
		}
	}
}

/*
	add avatar ids to the space
*/
func (space *Space) AddAvatarIDs(ids []string) {
	checkMap := make(map[string]bool)
	avatarList := append(space.AvatarIDs, ids...)
	space.AvatarIDs = nil
	for _, item := range avatarList {
		if _, value := checkMap[item]; !value {
			checkMap[item] = true
			space.AvatarIDs = append(space.AvatarIDs, item)
		}
	}
}

/*
	remove avatar ids from the space
*/
func (space *Space) RemoveAvatarID(ids []string) {
	for i := 0; i < len(space.AvatarIDs); i++ {
		item := space.AvatarIDs[i]
		for _, rem := range ids {
			if item == rem {
				space.AvatarIDs = append(space.AvatarIDs[:i], space.AvatarIDs[i+1:]...)
				i--
				break
			}
		}
	}
}
