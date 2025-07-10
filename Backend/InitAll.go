package Backend

import (
	"Siu_shop/Backend/model"
	"Siu_shop/Backend/router"
)

func StartBackend() {
	//启动数据库和http
	model.InitMysql()
	defer func() {
		model.Close()
	}()
	router.InitGin()
}
