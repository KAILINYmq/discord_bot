package model

import (
	"gorm.io/gorm"
	"time"
)

type NftEvent struct {
	gorm.Model
	NftAddress    string    `json:"nft_address" gorm:"index"`
	Contract      string    `json:"contract" gorm:"index:idx_contract_token_id"`
	TokenId       int64     `json:"token_id" gorm:"index:idx_contract_token_id"`
	Event         string    `json:"event"`
	Price         string    `json:"price"`
	FromAddress   string    `json:"from_address" gorm:"index"`
	ToAddress     string    `json:"to_address" gorm:"index"`
	Date          time.Time `json:"date"`
	SolanaAddress string    `json:"solana_address"`
}

func (this *NftEvent) CreateOne(db *gorm.DB) error {
	return db.Model(&NftEvent{}).Where("contract = ? and token_id = ? and event = ? and solana_address = ?",
		this.Contract, this.TokenId, this.Event, this.SolanaAddress).
		FirstOrCreate(&this).Error
}

func (this *NftEvent) FindEventByTokenId(db *gorm.DB) ([]NftEvent, error) {
	var result []NftEvent
	err := db.Model(&NftEvent{}).Where("contract = ? and token_id = ?", this.Contract, this.TokenId).
		Order("date DESC").
		Order("id").
		Find(&result).Error
	return result, err
}

func (this *NftEvent) FindEventByAddress(db *gorm.DB, ownerAddress string) ([]NftEvent, error) {
	var result []NftEvent
	err := db.Model(&NftEvent{}).Where("from_address = ? or to_address = ?", ownerAddress, ownerAddress).
		Order("date DESC").
		Order("id").
		Find(&result).Error
	return result, err
}
