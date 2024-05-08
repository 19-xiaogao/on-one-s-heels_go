package model

import "gorm.io/gorm"

type Position struct {
	TokenId                  uint64
	Liquidity                uint32
	Amount0                  uint64
	Amount1                  uint64
	Nonce                    uint32
	Operator                 string //  the address that is approved for spending this token
	PoolId                   string // pool address
	TickLower                uint32
	TickUpper                uint32
	FeeGrowthInside0LastX128 uint64
	FeeGrowthInside1LastX128 uint64
	TokensOwed0              uint32
	TokensOwed1              uint32
}

func (a Position) TableName() string {
	return "uniswap_position"
}

func (a Position) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Position) AllCount(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&a).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil

}
