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

func newNoticeServerEnv(db *gorm.DB, opts ...gen.DOOption) noticeServerEnv {
	_noticeServerEnv := noticeServerEnv{}

	_noticeServerEnv.noticeServerEnvDo.UseDB(db, opts...)
	_noticeServerEnv.noticeServerEnvDo.UseModel(&model.NoticeServerEnv{})

	tableName := _noticeServerEnv.noticeServerEnvDo.TableName()
	_noticeServerEnv.ALL = field.NewAsterisk(tableName)
	_noticeServerEnv.NoticeEnvID = field.NewString(tableName, "notice_env_id")
	_noticeServerEnv.ServerEnvID = field.NewString(tableName, "server_env_id")
	_noticeServerEnv.NoticeID = field.NewString(tableName, "notice_id")

	_noticeServerEnv.fillFieldMap()

	return _noticeServerEnv
}

type noticeServerEnv struct {
	noticeServerEnvDo

	ALL         field.Asterisk
	NoticeEnvID field.String
	ServerEnvID field.String
	NoticeID    field.String

	fieldMap map[string]field.Expr
}

func (n noticeServerEnv) Table(newTableName string) *noticeServerEnv {
	n.noticeServerEnvDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n noticeServerEnv) As(alias string) *noticeServerEnv {
	n.noticeServerEnvDo.DO = *(n.noticeServerEnvDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *noticeServerEnv) updateTableName(table string) *noticeServerEnv {
	n.ALL = field.NewAsterisk(table)
	n.NoticeEnvID = field.NewString(table, "notice_env_id")
	n.ServerEnvID = field.NewString(table, "server_env_id")
	n.NoticeID = field.NewString(table, "notice_id")

	n.fillFieldMap()

	return n
}

func (n *noticeServerEnv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *noticeServerEnv) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 3)
	n.fieldMap["notice_env_id"] = n.NoticeEnvID
	n.fieldMap["server_env_id"] = n.ServerEnvID
	n.fieldMap["notice_id"] = n.NoticeID
}

func (n noticeServerEnv) clone(db *gorm.DB) noticeServerEnv {
	n.noticeServerEnvDo.ReplaceConnPool(db.Statement.ConnPool)
	return n
}

func (n noticeServerEnv) replaceDB(db *gorm.DB) noticeServerEnv {
	n.noticeServerEnvDo.ReplaceDB(db)
	return n
}

type noticeServerEnvDo struct{ gen.DO }

type INoticeServerEnvDo interface {
	gen.SubQuery
	Debug() INoticeServerEnvDo
	WithContext(ctx context.Context) INoticeServerEnvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() INoticeServerEnvDo
	WriteDB() INoticeServerEnvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) INoticeServerEnvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INoticeServerEnvDo
	Not(conds ...gen.Condition) INoticeServerEnvDo
	Or(conds ...gen.Condition) INoticeServerEnvDo
	Select(conds ...field.Expr) INoticeServerEnvDo
	Where(conds ...gen.Condition) INoticeServerEnvDo
	Order(conds ...field.Expr) INoticeServerEnvDo
	Distinct(cols ...field.Expr) INoticeServerEnvDo
	Omit(cols ...field.Expr) INoticeServerEnvDo
	Join(table schema.Tabler, on ...field.Expr) INoticeServerEnvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INoticeServerEnvDo
	RightJoin(table schema.Tabler, on ...field.Expr) INoticeServerEnvDo
	Group(cols ...field.Expr) INoticeServerEnvDo
	Having(conds ...gen.Condition) INoticeServerEnvDo
	Limit(limit int) INoticeServerEnvDo
	Offset(offset int) INoticeServerEnvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INoticeServerEnvDo
	Unscoped() INoticeServerEnvDo
	Create(values ...*model.NoticeServerEnv) error
	CreateInBatches(values []*model.NoticeServerEnv, batchSize int) error
	Save(values ...*model.NoticeServerEnv) error
	First() (*model.NoticeServerEnv, error)
	Take() (*model.NoticeServerEnv, error)
	Last() (*model.NoticeServerEnv, error)
	Find() ([]*model.NoticeServerEnv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NoticeServerEnv, err error)
	FindInBatches(result *[]*model.NoticeServerEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NoticeServerEnv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INoticeServerEnvDo
	Assign(attrs ...field.AssignExpr) INoticeServerEnvDo
	Joins(fields ...field.RelationField) INoticeServerEnvDo
	Preload(fields ...field.RelationField) INoticeServerEnvDo
	FirstOrInit() (*model.NoticeServerEnv, error)
	FirstOrCreate() (*model.NoticeServerEnv, error)
	FindByPage(offset int, limit int) (result []*model.NoticeServerEnv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INoticeServerEnvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n noticeServerEnvDo) Debug() INoticeServerEnvDo {
	return n.withDO(n.DO.Debug())
}

func (n noticeServerEnvDo) WithContext(ctx context.Context) INoticeServerEnvDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n noticeServerEnvDo) ReadDB() INoticeServerEnvDo {
	return n.Clauses(dbresolver.Read)
}

func (n noticeServerEnvDo) WriteDB() INoticeServerEnvDo {
	return n.Clauses(dbresolver.Write)
}

func (n noticeServerEnvDo) Session(config *gorm.Session) INoticeServerEnvDo {
	return n.withDO(n.DO.Session(config))
}

func (n noticeServerEnvDo) Clauses(conds ...clause.Expression) INoticeServerEnvDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n noticeServerEnvDo) Returning(value interface{}, columns ...string) INoticeServerEnvDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n noticeServerEnvDo) Not(conds ...gen.Condition) INoticeServerEnvDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n noticeServerEnvDo) Or(conds ...gen.Condition) INoticeServerEnvDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n noticeServerEnvDo) Select(conds ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n noticeServerEnvDo) Where(conds ...gen.Condition) INoticeServerEnvDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n noticeServerEnvDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INoticeServerEnvDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n noticeServerEnvDo) Order(conds ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n noticeServerEnvDo) Distinct(cols ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n noticeServerEnvDo) Omit(cols ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n noticeServerEnvDo) Join(table schema.Tabler, on ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n noticeServerEnvDo) LeftJoin(table schema.Tabler, on ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n noticeServerEnvDo) RightJoin(table schema.Tabler, on ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n noticeServerEnvDo) Group(cols ...field.Expr) INoticeServerEnvDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n noticeServerEnvDo) Having(conds ...gen.Condition) INoticeServerEnvDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n noticeServerEnvDo) Limit(limit int) INoticeServerEnvDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n noticeServerEnvDo) Offset(offset int) INoticeServerEnvDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n noticeServerEnvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INoticeServerEnvDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n noticeServerEnvDo) Unscoped() INoticeServerEnvDo {
	return n.withDO(n.DO.Unscoped())
}

func (n noticeServerEnvDo) Create(values ...*model.NoticeServerEnv) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n noticeServerEnvDo) CreateInBatches(values []*model.NoticeServerEnv, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n noticeServerEnvDo) Save(values ...*model.NoticeServerEnv) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n noticeServerEnvDo) First() (*model.NoticeServerEnv, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoticeServerEnv), nil
	}
}

func (n noticeServerEnvDo) Take() (*model.NoticeServerEnv, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoticeServerEnv), nil
	}
}

func (n noticeServerEnvDo) Last() (*model.NoticeServerEnv, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoticeServerEnv), nil
	}
}

func (n noticeServerEnvDo) Find() ([]*model.NoticeServerEnv, error) {
	result, err := n.DO.Find()
	return result.([]*model.NoticeServerEnv), err
}

func (n noticeServerEnvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NoticeServerEnv, err error) {
	buf := make([]*model.NoticeServerEnv, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n noticeServerEnvDo) FindInBatches(result *[]*model.NoticeServerEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n noticeServerEnvDo) Attrs(attrs ...field.AssignExpr) INoticeServerEnvDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n noticeServerEnvDo) Assign(attrs ...field.AssignExpr) INoticeServerEnvDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n noticeServerEnvDo) Joins(fields ...field.RelationField) INoticeServerEnvDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n noticeServerEnvDo) Preload(fields ...field.RelationField) INoticeServerEnvDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n noticeServerEnvDo) FirstOrInit() (*model.NoticeServerEnv, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoticeServerEnv), nil
	}
}

func (n noticeServerEnvDo) FirstOrCreate() (*model.NoticeServerEnv, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NoticeServerEnv), nil
	}
}

func (n noticeServerEnvDo) FindByPage(offset int, limit int) (result []*model.NoticeServerEnv, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n noticeServerEnvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n noticeServerEnvDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n noticeServerEnvDo) Delete(models ...*model.NoticeServerEnv) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *noticeServerEnvDo) withDO(do gen.Dao) *noticeServerEnvDo {
	n.DO = *do.(*gen.DO)
	return n
}
