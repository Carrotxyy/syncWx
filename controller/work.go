package controller

import (
	"fmt"
	"github.com/Carrotxyy/syncWx/common/api"
	"github.com/Carrotxyy/syncWx/common/setting"
	"github.com/Carrotxyy/syncWx/models"
	"github.com/Carrotxyy/syncWx/repository"
)

/**
	微信访客同步任务
*/


type Work struct {
	Repository *repository.BaseRepository `inject:""`
	Config *setting.Config `inject:""`
}

//// 创建业务
//func CreateWork()Work{
//	r := repository.NewRepository()
//	config := setting.LoadingConf()
//	return Work{Repository: r,Config: config}
//}

/**
	恢复标志位

	@persons 需要恢复标志位的数据

	返回参数 (错误)
 */
func (w *Work)recovery(persons []*models.Personinfo)error{
	var ids []int
	// 获取所有数据的id
	for _, v := range persons {
		// 获取Per_ID
		ids = append(ids,v.Per_ID)
	}
	return w.Repository.BatchUpdate(ids)
}

// 创建访问路由
// @path : 具体的路由 例: http://xyz.szlimaiyun.cn/cloudSync/test   path = /cloudSync/test
/**
	添加路由参数

	@path 访问路由
 */
func (w *Work)SplicUrl(path string)string{

	// 获取 key
	key := api.GetKey(w.Config.WxAddr +"/interactive/key").(string)
	// 加密 key
	enkey := api.Encryption(key)

	// 构建url
	url := fmt.Sprintf("%s%s?EnKey=%s",w.Config.WxAddr,path,enkey)
	return url
}