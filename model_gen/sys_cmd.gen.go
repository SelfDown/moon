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

func newSysCmd(db *gorm.DB, opts ...gen.DOOption) sysCmd {
	_sysCmd := sysCmd{}

	_sysCmd.sysCmdDo.UseDB(db, opts...)
	_sysCmd.sysCmdDo.UseModel(&model.SysCmd{})

	tableName := _sysCmd.sysCmdDo.TableName()
	_sysCmd.ALL = field.NewAsterisk(tableName)
	_sysCmd.SysCmdID = field.NewString(tableName, "sys_cmd_id")
	_sysCmd.Cmd = field.NewString(tableName, "cmd")
	_sysCmd.Note = field.NewString(tableName, "note")
	_sysCmd.CmdText = field.NewString(tableName, "cmd_text")
	_sysCmd.Root = field.NewString(tableName, "root")
	_sysCmd.CreateTime = field.NewTime(tableName, "create_time")
	_sysCmd.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysCmd.Comments = field.NewString(tableName, "comments")

	_sysCmd.fillFieldMap()

	return _sysCmd
}

type sysCmd struct {
	sysCmdDo

	ALL        field.Asterisk
	SysCmdID   field.String
	Cmd        field.String // 具体脚本
	Note       field.String // 备注
	CmdText    field.String // 脚本中文说明
	Root       field.String // 是否需要root权限
	CreateTime field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime field.Time   // 记录修改时间（数据库自动写入）
	Comments   field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (s sysCmd) Table(newTableName string) *sysCmd {
	s.sysCmdDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysCmd) As(alias string) *sysCmd {
	s.sysCmdDo.DO = *(s.sysCmdDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysCmd) updateTableName(table string) *sysCmd {
	s.ALL = field.NewAsterisk(table)
	s.SysCmdID = field.NewString(table, "sys_cmd_id")
	s.Cmd = field.NewString(table, "cmd")
	s.Note = field.NewString(table, "note")
	s.CmdText = field.NewString(table, "cmd_text")
	s.Root = field.NewString(table, "root")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")

	s.fillFieldMap()

	return s
}

func (s *sysCmd) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysCmd) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["sys_cmd_id"] = s.SysCmdID
	s.fieldMap["cmd"] = s.Cmd
	s.fieldMap["note"] = s.Note
	s.fieldMap["cmd_text"] = s.CmdText
	s.fieldMap["root"] = s.Root
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
}

func (s sysCmd) clone(db *gorm.DB) sysCmd {
	s.sysCmdDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysCmd) replaceDB(db *gorm.DB) sysCmd {
	s.sysCmdDo.ReplaceDB(db)
	return s
}

type sysCmdDo struct{ gen.DO }

type ISysCmdDo interface {
	gen.SubQuery
	Debug() ISysCmdDo
	WithContext(ctx context.Context) ISysCmdDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysCmdDo
	WriteDB() ISysCmdDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysCmdDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysCmdDo
	Not(conds ...gen.Condition) ISysCmdDo
	Or(conds ...gen.Condition) ISysCmdDo
	Select(conds ...field.Expr) ISysCmdDo
	Where(conds ...gen.Condition) ISysCmdDo
	Order(conds ...field.Expr) ISysCmdDo
	Distinct(cols ...field.Expr) ISysCmdDo
	Omit(cols ...field.Expr) ISysCmdDo
	Join(table schema.Tabler, on ...field.Expr) ISysCmdDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysCmdDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysCmdDo
	Group(cols ...field.Expr) ISysCmdDo
	Having(conds ...gen.Condition) ISysCmdDo
	Limit(limit int) ISysCmdDo
	Offset(offset int) ISysCmdDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysCmdDo
	Unscoped() ISysCmdDo
	Create(values ...*model.SysCmd) error
	CreateInBatches(values []*model.SysCmd, batchSize int) error
	Save(values ...*model.SysCmd) error
	First() (*model.SysCmd, error)
	Take() (*model.SysCmd, error)
	Last() (*model.SysCmd, error)
	Find() ([]*model.SysCmd, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysCmd, err error)
	FindInBatches(result *[]*model.SysCmd, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysCmd) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysCmdDo
	Assign(attrs ...field.AssignExpr) ISysCmdDo
	Joins(fields ...field.RelationField) ISysCmdDo
	Preload(fields ...field.RelationField) ISysCmdDo
	FirstOrInit() (*model.SysCmd, error)
	FirstOrCreate() (*model.SysCmd, error)
	FindByPage(offset int, limit int) (result []*model.SysCmd, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysCmdDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysCmdDo) Debug() ISysCmdDo {
	return s.withDO(s.DO.Debug())
}

func (s sysCmdDo) WithContext(ctx context.Context) ISysCmdDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysCmdDo) ReadDB() ISysCmdDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysCmdDo) WriteDB() ISysCmdDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysCmdDo) Session(config *gorm.Session) ISysCmdDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysCmdDo) Clauses(conds ...clause.Expression) ISysCmdDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysCmdDo) Returning(value interface{}, columns ...string) ISysCmdDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysCmdDo) Not(conds ...gen.Condition) ISysCmdDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysCmdDo) Or(conds ...gen.Condition) ISysCmdDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysCmdDo) Select(conds ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysCmdDo) Where(conds ...gen.Condition) ISysCmdDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysCmdDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysCmdDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysCmdDo) Order(conds ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysCmdDo) Distinct(cols ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysCmdDo) Omit(cols ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysCmdDo) Join(table schema.Tabler, on ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysCmdDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysCmdDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysCmdDo) Group(cols ...field.Expr) ISysCmdDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysCmdDo) Having(conds ...gen.Condition) ISysCmdDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysCmdDo) Limit(limit int) ISysCmdDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysCmdDo) Offset(offset int) ISysCmdDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysCmdDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysCmdDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysCmdDo) Unscoped() ISysCmdDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysCmdDo) Create(values ...*model.SysCmd) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysCmdDo) CreateInBatches(values []*model.SysCmd, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysCmdDo) Save(values ...*model.SysCmd) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysCmdDo) First() (*model.SysCmd, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCmd), nil
	}
}

func (s sysCmdDo) Take() (*model.SysCmd, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCmd), nil
	}
}

func (s sysCmdDo) Last() (*model.SysCmd, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCmd), nil
	}
}

func (s sysCmdDo) Find() ([]*model.SysCmd, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysCmd), err
}

func (s sysCmdDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysCmd, err error) {
	buf := make([]*model.SysCmd, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysCmdDo) FindInBatches(result *[]*model.SysCmd, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysCmdDo) Attrs(attrs ...field.AssignExpr) ISysCmdDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysCmdDo) Assign(attrs ...field.AssignExpr) ISysCmdDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysCmdDo) Joins(fields ...field.RelationField) ISysCmdDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysCmdDo) Preload(fields ...field.RelationField) ISysCmdDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysCmdDo) FirstOrInit() (*model.SysCmd, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCmd), nil
	}
}

func (s sysCmdDo) FirstOrCreate() (*model.SysCmd, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysCmd), nil
	}
}

func (s sysCmdDo) FindByPage(offset int, limit int) (result []*model.SysCmd, count int64, err error) {
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

func (s sysCmdDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysCmdDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysCmdDo) Delete(models ...*model.SysCmd) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysCmdDo) withDO(do gen.Dao) *sysCmdDo {
	s.DO = *do.(*gen.DO)
	return s
}