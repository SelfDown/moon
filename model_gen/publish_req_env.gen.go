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

func newPublishReqEnv(db *gorm.DB, opts ...gen.DOOption) publishReqEnv {
	_publishReqEnv := publishReqEnv{}

	_publishReqEnv.publishReqEnvDo.UseDB(db, opts...)
	_publishReqEnv.publishReqEnvDo.UseModel(&model.PublishReqEnv{})

	tableName := _publishReqEnv.publishReqEnvDo.TableName()
	_publishReqEnv.ALL = field.NewAsterisk(tableName)
	_publishReqEnv.ReqEnvID = field.NewString(tableName, "req_env_id")
	_publishReqEnv.ReqID = field.NewString(tableName, "req_id")
	_publishReqEnv.EnvCode = field.NewString(tableName, "env_code")
	_publishReqEnv.ServerEnvID = field.NewString(tableName, "server_env_id")
	_publishReqEnv.CreateTime = field.NewTime(tableName, "create_time")
	_publishReqEnv.ModifyTime = field.NewTime(tableName, "modify_time")
	_publishReqEnv.Comments = field.NewString(tableName, "comments")

	_publishReqEnv.fillFieldMap()

	return _publishReqEnv
}

type publishReqEnv struct {
	publishReqEnvDo

	ALL         field.Asterisk
	ReqEnvID    field.String
	ReqID       field.String // 确认单申请ID
	EnvCode     field.String
	ServerEnvID field.String
	CreateTime  field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime  field.Time   // 记录修改时间（数据库自动写入）
	Comments    field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (p publishReqEnv) Table(newTableName string) *publishReqEnv {
	p.publishReqEnvDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publishReqEnv) As(alias string) *publishReqEnv {
	p.publishReqEnvDo.DO = *(p.publishReqEnvDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publishReqEnv) updateTableName(table string) *publishReqEnv {
	p.ALL = field.NewAsterisk(table)
	p.ReqEnvID = field.NewString(table, "req_env_id")
	p.ReqID = field.NewString(table, "req_id")
	p.EnvCode = field.NewString(table, "env_code")
	p.ServerEnvID = field.NewString(table, "server_env_id")
	p.CreateTime = field.NewTime(table, "create_time")
	p.ModifyTime = field.NewTime(table, "modify_time")
	p.Comments = field.NewString(table, "comments")

	p.fillFieldMap()

	return p
}

func (p *publishReqEnv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publishReqEnv) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["req_env_id"] = p.ReqEnvID
	p.fieldMap["req_id"] = p.ReqID
	p.fieldMap["env_code"] = p.EnvCode
	p.fieldMap["server_env_id"] = p.ServerEnvID
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["modify_time"] = p.ModifyTime
	p.fieldMap["comments"] = p.Comments
}

func (p publishReqEnv) clone(db *gorm.DB) publishReqEnv {
	p.publishReqEnvDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publishReqEnv) replaceDB(db *gorm.DB) publishReqEnv {
	p.publishReqEnvDo.ReplaceDB(db)
	return p
}

type publishReqEnvDo struct{ gen.DO }

type IPublishReqEnvDo interface {
	gen.SubQuery
	Debug() IPublishReqEnvDo
	WithContext(ctx context.Context) IPublishReqEnvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublishReqEnvDo
	WriteDB() IPublishReqEnvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublishReqEnvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublishReqEnvDo
	Not(conds ...gen.Condition) IPublishReqEnvDo
	Or(conds ...gen.Condition) IPublishReqEnvDo
	Select(conds ...field.Expr) IPublishReqEnvDo
	Where(conds ...gen.Condition) IPublishReqEnvDo
	Order(conds ...field.Expr) IPublishReqEnvDo
	Distinct(cols ...field.Expr) IPublishReqEnvDo
	Omit(cols ...field.Expr) IPublishReqEnvDo
	Join(table schema.Tabler, on ...field.Expr) IPublishReqEnvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqEnvDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqEnvDo
	Group(cols ...field.Expr) IPublishReqEnvDo
	Having(conds ...gen.Condition) IPublishReqEnvDo
	Limit(limit int) IPublishReqEnvDo
	Offset(offset int) IPublishReqEnvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqEnvDo
	Unscoped() IPublishReqEnvDo
	Create(values ...*model.PublishReqEnv) error
	CreateInBatches(values []*model.PublishReqEnv, batchSize int) error
	Save(values ...*model.PublishReqEnv) error
	First() (*model.PublishReqEnv, error)
	Take() (*model.PublishReqEnv, error)
	Last() (*model.PublishReqEnv, error)
	Find() ([]*model.PublishReqEnv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReqEnv, err error)
	FindInBatches(result *[]*model.PublishReqEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PublishReqEnv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublishReqEnvDo
	Assign(attrs ...field.AssignExpr) IPublishReqEnvDo
	Joins(fields ...field.RelationField) IPublishReqEnvDo
	Preload(fields ...field.RelationField) IPublishReqEnvDo
	FirstOrInit() (*model.PublishReqEnv, error)
	FirstOrCreate() (*model.PublishReqEnv, error)
	FindByPage(offset int, limit int) (result []*model.PublishReqEnv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublishReqEnvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p publishReqEnvDo) Debug() IPublishReqEnvDo {
	return p.withDO(p.DO.Debug())
}

func (p publishReqEnvDo) WithContext(ctx context.Context) IPublishReqEnvDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publishReqEnvDo) ReadDB() IPublishReqEnvDo {
	return p.Clauses(dbresolver.Read)
}

func (p publishReqEnvDo) WriteDB() IPublishReqEnvDo {
	return p.Clauses(dbresolver.Write)
}

func (p publishReqEnvDo) Session(config *gorm.Session) IPublishReqEnvDo {
	return p.withDO(p.DO.Session(config))
}

func (p publishReqEnvDo) Clauses(conds ...clause.Expression) IPublishReqEnvDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publishReqEnvDo) Returning(value interface{}, columns ...string) IPublishReqEnvDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publishReqEnvDo) Not(conds ...gen.Condition) IPublishReqEnvDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publishReqEnvDo) Or(conds ...gen.Condition) IPublishReqEnvDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publishReqEnvDo) Select(conds ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publishReqEnvDo) Where(conds ...gen.Condition) IPublishReqEnvDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publishReqEnvDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPublishReqEnvDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publishReqEnvDo) Order(conds ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publishReqEnvDo) Distinct(cols ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publishReqEnvDo) Omit(cols ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publishReqEnvDo) Join(table schema.Tabler, on ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publishReqEnvDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publishReqEnvDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publishReqEnvDo) Group(cols ...field.Expr) IPublishReqEnvDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publishReqEnvDo) Having(conds ...gen.Condition) IPublishReqEnvDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publishReqEnvDo) Limit(limit int) IPublishReqEnvDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publishReqEnvDo) Offset(offset int) IPublishReqEnvDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publishReqEnvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqEnvDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publishReqEnvDo) Unscoped() IPublishReqEnvDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publishReqEnvDo) Create(values ...*model.PublishReqEnv) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publishReqEnvDo) CreateInBatches(values []*model.PublishReqEnv, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publishReqEnvDo) Save(values ...*model.PublishReqEnv) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publishReqEnvDo) First() (*model.PublishReqEnv, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqEnv), nil
	}
}

func (p publishReqEnvDo) Take() (*model.PublishReqEnv, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqEnv), nil
	}
}

func (p publishReqEnvDo) Last() (*model.PublishReqEnv, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqEnv), nil
	}
}

func (p publishReqEnvDo) Find() ([]*model.PublishReqEnv, error) {
	result, err := p.DO.Find()
	return result.([]*model.PublishReqEnv), err
}

func (p publishReqEnvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReqEnv, err error) {
	buf := make([]*model.PublishReqEnv, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publishReqEnvDo) FindInBatches(result *[]*model.PublishReqEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publishReqEnvDo) Attrs(attrs ...field.AssignExpr) IPublishReqEnvDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publishReqEnvDo) Assign(attrs ...field.AssignExpr) IPublishReqEnvDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publishReqEnvDo) Joins(fields ...field.RelationField) IPublishReqEnvDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publishReqEnvDo) Preload(fields ...field.RelationField) IPublishReqEnvDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publishReqEnvDo) FirstOrInit() (*model.PublishReqEnv, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqEnv), nil
	}
}

func (p publishReqEnvDo) FirstOrCreate() (*model.PublishReqEnv, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqEnv), nil
	}
}

func (p publishReqEnvDo) FindByPage(offset int, limit int) (result []*model.PublishReqEnv, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p publishReqEnvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publishReqEnvDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publishReqEnvDo) Delete(models ...*model.PublishReqEnv) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publishReqEnvDo) withDO(do gen.Dao) *publishReqEnvDo {
	p.DO = *do.(*gen.DO)
	return p
}
