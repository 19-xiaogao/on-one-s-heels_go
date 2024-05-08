package dao

import "on_one_heels_go/internal/model"

func (d *Dao) CountPosition() (int64, error) {
	position := model.Position{}
	return position.AllCount(d.engine)
}

// CreatePosition  create position
func (d *Dao) CreatePosition(TokenId uint64,
	Liquidity uint32,
	Amount0 uint64,
	Amount1 uint64,
	Nonce uint32,
	Operator string,
	TickLower uint32,
	TickUpper uint32,
	FeeGrowthInside0LastX128 uint64,
	FeeGrowthInside1LastX128 uint64,
	TokensOwed0 uint32,
	TokensOwed1 uint32,
) error {
	position := model.Position{
		TokenId:                  TokenId,
		Liquidity:                Liquidity,
		Amount0:                  Amount0,
		Amount1:                  Amount1,
		Nonce:                    Nonce,
		Operator:                 Operator,
		TickLower:                TickLower,
		TickUpper:                TickUpper,
		FeeGrowthInside0LastX128: FeeGrowthInside0LastX128,
		FeeGrowthInside1LastX128: FeeGrowthInside1LastX128,
		TokensOwed0:              TokensOwed0,
		TokensOwed1:              TokensOwed1,
	}
	return position.Create(d.engine)
}
