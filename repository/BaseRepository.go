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
	sql := "insert into `go_visitor` (`vis_name`,`vis_number`,`vis_uname`,`vis_unumber`,`vis_idtype`,`vis_idnum`,`vis_starttime`,`vis_endtime`,`vis_message`,`vis_isacc`,`vis_state`,`vis_filename`,`vis_perid`) values "
	// 实际参数
	rels := []interface{}{}
	// sql语句参数
	rowSql := "(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	var insert []string
	// 构建批量添加的sql
	for _, e := range value {
		insert = append(insert,rowSql)
		rels = append(rels,e.Name, e.Number, e.UName ,e.UNumber,e.IdType,e.IdNum,e.StartTime,e.EndTime,e.Message,e.IsAcc,e.State,e.FileName,e.Per_ID)

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