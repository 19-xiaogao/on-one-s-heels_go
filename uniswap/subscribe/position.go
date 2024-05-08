package subscribe

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"on_one_heels_go/global"
)

type Subscribe struct {
	client *ethclient.Client
}

func newSubscribe() (*Subscribe, error) {
	client, err := ethclient.Dial(global.UniSwapSetting.Rpc)
	if err != nil {
		return nil, err
	}
	return &Subscribe{client: client}, err
}

func (s *Subscribe) SubscribeIncreaseLiquidityEvent() {
}
