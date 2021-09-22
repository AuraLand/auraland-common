package auralandapi

const (
	/*
		Asset's RPC API
	*/
	RPCAssetList    = "assetList"
	RPCAssetGetByID = "assetGetById"
	RPCAssetCreate  = "assetCreate"
	RPCAssetUpdate  = "assetUpdate"
	RPCAssetDrop    = "assetDrop"
	/*
		Avatar's RPC API
	*/
	RPCAvatarList   = "avatarList"
	RPCAvatarChoose = "avatarChoose"

	/*
		Entity's RPC API
	*/
	RPCEntityLinkAsset         = "entityAssetLink"
	RPCEntityUnlinkAsset       = "entityAssetUnlink"
	RPCEntityCreate            = "entityCreate"
	RPCEntityUpdate            = "entityUpdate"
	RPCEntityDrop              = "entityDrop"
	RPCEntityGetById           = "entityGet"
	RPCEntityList              = "entityList"
	RPCEntityListByUserID      = "entityListByUserID"
	RPCEntityListByUserSpaceId = "entityListByUserSpaceId"

	/*
		Space's RPC API
	*/
	RPCSpaceCreate            = "spaceCreate"
	RPCSpaceUpdate            = "spaceUpdate"
	RPCSpaceDrop              = "spaceDrop"
	RPCSpaceGetById           = "spaceGetById"
	RPCSpaceList              = "spaceList"
	RPCSpaceListByCreatorID   = "spaceListByCreatorID"
	RPCSpaceLinkEntity        = "spaceEntityLink"
	RPCSpaceUnlinkEntity      = "spaceEntityUnlink"
	RPCSpaceLinkAsset         = "spaceAssetLink"
	RPCSpaceUnlinkAsset       = "spaceAssetUnlink"
	RPCSpaceLinkAvatar        = "spaceAvatarLink"
	RPCSpaceUnlinkAvatar      = "spaceAvatarUnlink"
	RPCSpaceSendStreamMessage = "spaceSendStreamMessage"
	RPCSpaceUserList          = "spaceUserList"
	RPCSpaceUserJoin          = "spaceUserJoin"
	RPCSpaceUserLeave         = "spaceUserLeave"
)
