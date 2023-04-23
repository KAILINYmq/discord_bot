package syncer

type MintData struct {
	Nft struct {
		MintAddress    string `json:"mint_address"`
		TokenId        string `json:"token_id"`
		CreatorAddress string `json:"creator_address"`
	}
}

type MintInfo struct {
	Status string   `json:"status"`
	Data   MintData `json:"data"`
}

type LevelData struct {
	MintAddress  string `json:"mint_address"`
	OwnerAddress string `json:"owner_address"`
	Level        int    `json:"level"`
}

type LevelInfo struct {
	Status string    `json:"status"`
	Data   LevelData `json:"data"`
}

type CardMetaInfos struct {
	Status string         `json:"status"`
	Tokens []CardMetadata `json:"tokens"`
}

type CardMetadata struct {
	TokenId     string `json:"token_id"`
	UserAddress string `json:"user_address"`
}
