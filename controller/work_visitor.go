package controller

import (
	"fmt"
	"syncWx/common/api"
	"syncWx/models"
)



// 下载 所有访客数据
/**
	下载访客数据

	返回参数 (错误)
 */
func (w *Work)VisitorLoad()error{

	// 构建url
	url := w.SplicUrl("/interactive/visitor")
	obj := struct {
		Data   []*models.Visitor `json:"data"`
		Status bool              `json:"status"`
	}{}
	// 获取数据
	err  := api.HttpGet(url,&obj)
	if err != nil {
		fmt.Println("拉取访客数据错误!")
		return err
	}
	if len(obj.Data) <= 0 {
		fmt.Println("暂无数据!")
		return nil
	}
	err = w.Repository.BatchSave(obj.Data)
	// 测试代码
	fmt.Println("访客数据:",obj.Data)
	if err != nil {
		fmt.Println("保存访客数据错误:",err)
		return err
	}

	return nil
}


