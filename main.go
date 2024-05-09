package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"on_one_heels_go/global"
	"on_one_heels_go/internal/model"
	"on_one_heels_go/pkg/settings"
	"on_one_heels_go/uniswap/subscribe"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	//router := routers.NewRouter()
	//
	//err := router.Run(":" + global.ServerSetting.HttpPort)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	newSubscribe, err := subscribe.NewSubscribe()
	if err != nil {
		return
	}
	newSubscribe.SubscribeIncreaseLiquidityEvent()
}

func setupSetting() error {
	setting, err := settings.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)

	if err != nil {
		return err
	}

	err = setting.ReadSection("UniSwap", &global.UniSwapSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	//init all table
	global.DBEngine.AutoMigrate(&model.Position{})
	return nil
}
