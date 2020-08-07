package repository

import (
	"github.com/Carrotxyy/syncWx/common/db"
	"github.com/Carrotxyy/syncWx/models"
	"strings"
)

type BaseRepository struct {
	DB *db.DB `inject:""`
}

// 初始化
//func NewRepository()*BaseRepository{
//	db := db.DB{}
//	err := db.Connect()
//	if err != nil {
//		log.Panic("初始化数据库错误:",err)
//		os.Exit(1)
//	}
//	return &BaseRepository{
//		&db,
//	}
//}

/**
	根据条件，获取人员表数据

	@where 查询条件
	@out   装载结果
	@sel   查询的字段

	返回参数 (错误，数据数量)
 */
func (b *BaseRepository) Get(where , out interface{} ,sel string)(error,int){
	var count int
	// 获取数据库对象,Not 反向获取
	db := b.DB.Conn.Not(where)
	if sel != "" {
		// 检索的字段
		db = db.Select(sel)
	}
	err := db.Find(out).Count(&count).Error
	return err,count
}

/**
	新增数据

	@value 新数据

	返回参数 (错误)
 */
func (b *BaseRepository) Save(value interface{})error{
	return b.DB.Conn.Create(value).Error
}

/**
	批量保存访客数据

	@value 访客数据

	返回参数 (错误)
 */
func (b *BaseRepository) BatchSave(value []*models.Visitor)error{
	sql := "insert into `go_visitor` (`Vis_Name`,`Vis_Number`,`Vis_UName`,`Vis_UNumber`,`Vis_IDType`,`Vis_IDNum`,`Vis_StartTime`,`Vis_EndTime`,`Vis_Message`,`Vis_IsAcc`,`Vis_State`,`Vis_FileName`,`Vis_PerID`,`Vis_SenseMark`,`Vis_PeakMark`,`Vis_MegMark`) values "
	// 实际参数
	rels := []interface{}{}
	// sql语句参数
	rowSql := "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	var insert []string
	// 构建批量添加的sql
	for _, e := range value {
		insert = append(insert,rowSql)
		rels = append(rels,e.Vis_Name, e.Vis_Number, e.Vis_UName ,e.Vis_UNumber,e.Vis_IDType,e.Vis_IDNum,e.Vis_StartTime,e.Vis_EndTime,e.Vis_Message,e.Vis_IsAcc,e.Vis_State,e.Vis_FileName,e.Vis_PerID,"1","1","1")

	}
	// 拼接sql
	sql = sql + strings.Join(insert,",")
	// 执行sql
	return b.DB.Conn.Exec(sql,rels...).Error
}


/**
	业主表 恢复标志位

	@ids 需要恢复标志位数据的ID

	返回参数 (错误)
 */
func (b *BaseRepository) BatchUpdate(ids []int)error{
	return b.DB.Conn.Table("go_personinfo").Where("Per_ID IN (?)",ids).Update("Per_WXMark","0").Error
}
