// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model_gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"moon/model"
)

func newSysBoard(db *gorm.DB, opts ...gen.DOOption) sysBoard {
	_sysBoard := sysBoard{}

	_sysBoard.sysBoardDo.UseDB(db, opts...)
	_sysBoard.sysBoardDo.UseModel(&model.SysBoard{})

	tableName := _sysBoard.sysBoardDo.TableName()
	_sysBoard.ALL = field.NewAsterisk(tableName)
	_sysBoard.SysBoardID = field.NewInt32(tableName, "sys_board_id")
	_sysBoard.SysBoardMessage = field.NewString(tableName, "sys_board_message")
	_sysBoard.MessageStatus = field.NewString(tableName, "message_status")
	_sysBoard.OpUser = field.NewString(tableName, "op_user")
	_sysBoard.DetailURL = field.NewString(tableName, "detail_url")
	_sysBoard.OrderIdx = field.NewInt32(tableName, "order_idx")
	_sysBoard.CreateTime = field.NewTime(tableName, "create_time")
	_sysBoard.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysBoard.Comments = field.NewString(tableName, "comments")

	_sysBoard.fillFieldMap()

	return _sysBoard
}

type sysBoard struct {
	sysBoardDo

	ALL             field.Asterisk
	SysBoardID      field.Int32
	SysBoardMessage field.String // 公告内容，支持html
	/*
		0 ，失效
		1 或空，有效
	*/
	MessageStatus field.String
	OpUser        field.String // 操作人
	DetailURL     field.String // 查看页面详情
	OrderIdx      field.Int32  // 公告排序
	CreateTime    field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime    field.Time   // 记录修改时间（数据库自动写入）
	Comments      field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysBoard) Table(newTableName string) *sysBoard {
	s.sysBoardDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysBoard) As(alias string) *sysBoard {
	s.sysBoardDo.DO = *(s.sysBoardDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysBoard) updateTableName(table string) *sysBoard {
	s.ALL = field.NewAsterisk(table)
	s.SysBoardID = field.NewInt32(table, "sys_board_id")
	s.SysBoardMessage = field.NewString(table, "sys_board_message")
	s.MessageStatus = field.NewString(table, "message_status")
	s.OpUser = field.NewString(table, "op_user")
	s.DetailURL = field.NewString(table, "detail_url")
	s.OrderIdx = field.NewInt32(table, "order_idx")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysBoard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysBoard) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["sys_board_id"] = s.SysBoardID
	s.fieldMap["sys_board_message"] = s.SysBoardMessage
	s.fieldMap["message_status"] = s.MessageStatus
	s.fieldMap["op_user"] = s.OpUser
	s.fieldMap["detail_url"] = s.DetailURL
	s.fieldMap["order_idx"] = s.OrderIdx
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysBoard) clone(db *gorm.DB) sysBoard {
	s.sysBoardDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysBoard) replaceDB(db *gorm.DB) sysBoard {
	s.sysBoardDo.ReplaceDB(db)
	return s
}

type sysBoardDo struct{ gen.DO }

type ISysBoardDo interface {
	gen.SubQuery
	Debug() ISysBoardDo
	WithContext(ctx context.Context) ISysBoardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysBoardDo
	WriteDB() ISysBoardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysBoardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysBoardDo
	Not(conds ...gen.Condition) ISysBoardDo
	Or(conds ...gen.Condition) ISysBoardDo
	Select(conds ...field.Expr) ISysBoardDo
	Where(conds ...gen.Condition) ISysBoardDo
	Order(conds ...field.Expr) ISysBoardDo
	Distinct(cols ...field.Expr) ISysBoardDo
	Omit(cols ...field.Expr) ISysBoardDo
	Join(table schema.Tabler, on ...field.Expr) ISysBoardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysBoardDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysBoardDo
	Group(cols ...field.Expr) ISysBoardDo
	Having(conds ...gen.Condition) ISysBoardDo
	Limit(limit int) ISysBoardDo
	Offset(offset int) ISysBoardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysBoardDo
	Unscoped() ISysBoardDo
	Create(values ...*model.SysBoard) error
	CreateInBatches(values []*model.SysBoard, batchSize int) error
	Save(values ...*model.SysBoard) error
	First() (*model.SysBoard, error)
	Take() (*model.SysBoard, error)
	Last() (*model.SysBoard, error)
	Find() ([]*model.SysBoard, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysBoard, err error)
	FindInBatches(result *[]*model.SysBoard, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysBoard) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysBoardDo
	Assign(attrs ...field.AssignExpr) ISysBoardDo
	Joins(fields ...field.RelationField) ISysBoardDo
	Preload(fields ...field.RelationField) ISysBoardDo
	FirstOrInit() (*model.SysBoard, error)
	FirstOrCreate() (*model.SysBoard, error)
	FindByPage(offset int, limit int) (result []*model.SysBoard, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysBoardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysBoardDo) Debug() ISysBoardDo {
	return s.withDO(s.DO.Debug())
}

func (s sysBoardDo) WithContext(ctx context.Context) ISysBoardDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysBoardDo) ReadDB() ISysBoardDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysBoardDo) WriteDB() ISysBoardDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysBoardDo) Session(config *gorm.Session) ISysBoardDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysBoardDo) Clauses(conds ...clause.Expression) ISysBoardDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysBoardDo) Returning(value interface{}, columns ...string) ISysBoardDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysBoardDo) Not(conds ...gen.Condition) ISysBoardDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysBoardDo) Or(conds ...gen.Condition) ISysBoardDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysBoardDo) Select(conds ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysBoardDo) Where(conds ...gen.Condition) ISysBoardDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysBoardDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysBoardDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysBoardDo) Order(conds ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysBoardDo) Distinct(cols ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysBoardDo) Omit(cols ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysBoardDo) Join(table schema.Tabler, on ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysBoardDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysBoardDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysBoardDo) Group(cols ...field.Expr) ISysBoardDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysBoardDo) Having(conds ...gen.Condition) ISysBoardDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysBoardDo) Limit(limit int) ISysBoardDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysBoardDo) Offset(offset int) ISysBoardDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysBoardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysBoardDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysBoardDo) Unscoped() ISysBoardDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysBoardDo) Create(values ...*model.SysBoard) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysBoardDo) CreateInBatches(values []*model.SysBoard, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysBoardDo) Save(values ...*model.SysBoard) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysBoardDo) First() (*model.SysBoard, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysBoard), nil
	}
}

func (s sysBoardDo) Take() (*model.SysBoard, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysBoard), nil
	}
}

func (s sysBoardDo) Last() (*model.SysBoard, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysBoard), nil
	}
}

func (s sysBoardDo) Find() ([]*model.SysBoard, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysBoard), err
}

func (s sysBoardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysBoard, err error) {
	buf := make([]*model.SysBoard, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysBoardDo) FindInBatches(result *[]*model.SysBoard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysBoardDo) Attrs(attrs ...field.AssignExpr) ISysBoardDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysBoardDo) Assign(attrs ...field.AssignExpr) ISysBoardDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysBoardDo) Joins(fields ...field.RelationField) ISysBoardDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysBoardDo) Preload(fields ...field.RelationField) ISysBoardDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysBoardDo) FirstOrInit() (*model.SysBoard, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysBoard), nil
	}
}

func (s sysBoardDo) FirstOrCreate() (*model.SysBoard, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysBoard), nil
	}
}

func (s sysBoardDo) FindByPage(offset int, limit int) (result []*model.SysBoard, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysBoardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysBoardDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysBoardDo) Delete(models ...*model.SysBoard) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysBoardDo) withDO(do gen.Dao) *sysBoardDo {
	s.DO = *do.(*gen.DO)
	return s
}