package syncWx

import "syncWx/controller"

type Init struct {
	Work controller.Work
}

// 初始化
func Create()*Init {
	work := controller.CreateWork()
	return &Init{Work: work}
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