## 微信访客——管理系统  本地同步模块

#### 简介

用于本地管理系统与云端访客系统的数据同步模块。

上传业主信息，下载访客信息

#### 下载

```go
import github.com/Carrotxyy/syncWx
```

#### 使用

```go 
	// 获取配置对象
	config := syncWx.GetConfig()
	// 配置参数
	config.WxAddr = "http://xyz.szlimaiyun.cn"
	config.DbType = "mysql"
	config.DbUser = "root"
	config.DbPassword = "123456"
	config.DbIP = "127.0.0.1"
	config.DbName = "gin-vue"
	config.TablePrefix = "go_"
	// 创建任务
	wxWork,err := Create(config)
	if err != nil {
		fmt.Println("创建微信同步任务错误！")
		return
	}
	// 以定时任务方式运行
	wxWork.RunWork()

	

```

