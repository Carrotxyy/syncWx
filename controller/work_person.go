package controller

import (
	"fmt"
	"github.com/Carrotxyy/syncWx/common/api"
	"github.com/Carrotxyy/syncWx/models"
)

/**
	微信同步- 业主数据同步
 */

/**
	上传可接收访客访问的业主信息

	返回参数 (错误)

 */
func (w *Work)PersonUpload()error{

	var persons []*models.Personinfo
	// 查找修改过数据的业主,并且这些数据没有被同步到微信系统中
	where := models.Personinfo{
		Per_WXMark: "0",
	}
	// 获取数据
	err,count := w.Repository.Get(where,&persons,"Per_ID,Per_Name,Per_ContactTel,Per_Status,Per_AllowVisit")
	if err != nil {
		fmt.Println("获取用户数据错误:",err)
		return err
	}
	fmt.Println("数据数量:",count)
	if count == 0{
		fmt.Println("暂无新数据")
		return nil
	}
	// 构建请求 路径  携带 加密钥匙
	url := w.SplicUrl("/interactive/saveperson")
	// 数据类型
	contentType := "application/json"
	var res struct{
		Status bool `json:"status"`
	}
	// 发送请求
	err = api.HttpPost(url,persons,&res,contentType)
	if err != nil || !res.Status {
		fmt.Println("发送同步请求错误!")
	}else {
		// 上传同步成功后，需要将所有同步过数据的标志位设置为1，表示数据已经同步
		err = w.recovery(persons)
	}

	return err
}
