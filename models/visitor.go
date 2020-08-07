package models

type Visitor struct {
	Vis_Name       string `gorm:"cloumn:Vis_Name" json:"name"`           // 业主姓名
	Vis_Number     string `gorm:"column:Vis_Number" json:"number"`       // 业主电话
	Vis_UName      string `gorm:"column:Vis_UName" json:"uname"`         // 访客姓名
	Vis_UNumber    string `gorm:"column:Vis_UNumber" json:"unumber"`     // 访客电话
	Vis_IDType     string `gorm:"column:Vis_IDType" json:"idtype"`       // 证件类型
	Vis_IDNum      string `gorm:"column:Vis_IDNum" json:"idnum"`         // 证件号码
	Vis_StartTime  string `gorm:"column:Vis_StartTime" json:"starttime"` // 访问开始时间
	Vis_EndTime    string `gorm:"column:Vis_EndTime" json:"endtime"`     // 访问结束时间
	Vis_Message    string `gorm:"column:Vis_Message" json:"message"`     // 访问留言
	Vis_IsAcc      bool   `gorm:"column:Vis_IsAcc" json:"isAcc"`         // 业主是否对当前的预约进行了操作
	Vis_State      bool   `gorm:"column:Vis_State" json:"state"`         // 操作结果
	Vis_FileName   string `gorm:"column:Vis_FileName" json:"filename"`   // 访客图片名(根据图片名-加上远程微信访客系统路径，动态获取图片，本地不保存图片数据)
	Vis_PerID      int    `gorm:"column:Vis_PerID" json:"Per_ID"`        // 业主ID
	Vis_SenseVisID string `gorm:"column:Vis_SenseVisID"`                 // 商汤访客ID
	Vis_PeakVisID  string `gorm:"column:Vis_PeakVisID"`                  // 披克访客ID
	Vis_MegVisID   string `gorm:"column:Vis_MegVisID"`                   // 旷视访客ID
	Vis_SenseMark  string `gorm:"column:Vis_SenseMark"`                  // 同步商汤系统标志位
	Vis_PeakMark   string `gorm:"column:Vis_PeakMark"`                   // 同步披克系统标志位
	Vis_MegMark    string `gorm:"column:Vis_MegMark"`                    // 同步旷视系统标志位
}
