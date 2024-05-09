package subscribe

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"on_one_heels_go/global"
)

type Subscribe struct {
	client *ethclient.Client
}

func NewSubscribe() (*Subscribe, error) {
	client, err := ethclient.Dial(global.UniSwapSetting.Rpc)
	if err != nil {
		return nil, err
	}
	return &Subscribe{client: client}, err
}

func (s *Subscribe) SubscribeIncreaseLiquidityEvent() {
	contractAddress := common.HexToAddress(global.UniSwapSetting.PositionManagerAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

//func (s *Subscribe) Subscribe
