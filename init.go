package main

import (
	"fmt"
	db "github.com/Carrotxyy/syncWx/common/db"
	"github.com/Carrotxyy/syncWx/common/setting"
	"github.com/Carrotxyy/syncWx/controller"
	"github.com/facebookgo/inject"
	"log"
)

type Init struct {
	Work *controller.Work `inject:""`
}

// 初始化
func Create(config setting.Config)(*Init,error) {
	conn := db.DB{}
	// 连接数据库
	err := conn.Connect(config)
	if err != nil {
		fmt.Println("连接数据库错误:",err)
		return nil,err
	}

	var init Init

	// 依赖注入
	var injector inject.Graph
	if err := injector.Provide(
		&inject.Object{Value: &init},
		&inject.Object{Value: &conn},
		&inject.Object{Value: &config},
	); err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("injector fatal: ", err)
	}

	return &init,nil
}

// 获取配置对象
func GetConfig()setting.Config{
	return setting.Config{}
}

/**
创建任务，用于定时任务时调用
*/
func (init *Init)RunWork(){
	// 运行 上传业主
	init.Work.PersonUpload()
	// 运行 下载访客
	init.Work.VisitorLoad()
}

/**
	以路由方式加载
 */
//func (inti *Init)HttpWork(router *gin.Engine){
//	// 加载controller
//}



