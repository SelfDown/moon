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

func newPublishReqDeployProductEnv(db *gorm.DB, opts ...gen.DOOption) publishReqDeployProductEnv {
	_publishReqDeployProductEnv := publishReqDeployProductEnv{}

	_publishReqDeployProductEnv.publishReqDeployProductEnvDo.UseDB(db, opts...)
	_publishReqDeployProductEnv.publishReqDeployProductEnvDo.UseModel(&model.PublishReqDeployProductEnv{})

	tableName := _publishReqDeployProductEnv.publishReqDeployProductEnvDo.TableName()
	_publishReqDeployProductEnv.ALL = field.NewAsterisk(tableName)
	_publishReqDeployProductEnv.PublishReqDeployProductEnvID = field.NewString(tableName, "publish_req_deploy_product_env_id")
	_publishReqDeployProductEnv.ReqID = field.NewString(tableName, "req_id")
	_publishReqDeployProductEnv.OpUser = field.NewString(tableName, "op_user")
	_publishReqDeployProductEnv.ServerEnvID = field.NewString(tableName, "server_env_id")
	_publishReqDeployProductEnv.CreateTime = field.NewTime(tableName, "create_time")
	_publishReqDeployProductEnv.ModifyTime = field.NewTime(tableName, "modify_time")
	_publishReqDeployProductEnv.Comments = field.NewString(tableName, "comments")

	_publishReqDeployProductEnv.fillFieldMap()

	return _publishReqDeployProductEnv
}

type publishReqDeployProductEnv struct {
	publishReqDeployProductEnvDo

	ALL                          field.Asterisk
	PublishReqDeployProductEnvID field.String
	ReqID                        field.String // 确认单申请ID
	OpUser                       field.String
	ServerEnvID                  field.String
	CreateTime                   field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime                   field.Time   // 记录修改时间（数据库自动写入）
	Comments                     field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (p publishReqDeployProductEnv) Table(newTableName string) *publishReqDeployProductEnv {
	p.publishReqDeployProductEnvDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publishReqDeployProductEnv) As(alias string) *publishReqDeployProductEnv {
	p.publishReqDeployProductEnvDo.DO = *(p.publishReqDeployProductEnvDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publishReqDeployProductEnv) updateTableName(table string) *publishReqDeployProductEnv {
	p.ALL = field.NewAsterisk(table)
	p.PublishReqDeployProductEnvID = field.NewString(table, "publish_req_deploy_product_env_id")
	p.ReqID = field.NewString(table, "req_id")
	p.OpUser = field.NewString(table, "op_user")
	p.ServerEnvID = field.NewString(table, "server_env_id")
	p.CreateTime = field.NewTime(table, "create_time")
	p.ModifyTime = field.NewTime(table, "modify_time")
	p.Comments = field.NewString(table, "comments")

	p.fillFieldMap()

	return p
}

func (p *publishReqDeployProductEnv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publishReqDeployProductEnv) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["publish_req_deploy_product_env_id"] = p.PublishReqDeployProductEnvID
	p.fieldMap["req_id"] = p.ReqID
	p.fieldMap["op_user"] = p.OpUser
	p.fieldMap["server_env_id"] = p.ServerEnvID
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["modify_time"] = p.ModifyTime
	p.fieldMap["comments"] = p.Comments
}

func (p publishReqDeployProductEnv) clone(db *gorm.DB) publishReqDeployProductEnv {
	p.publishReqDeployProductEnvDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publishReqDeployProductEnv) replaceDB(db *gorm.DB) publishReqDeployProductEnv {
	p.publishReqDeployProductEnvDo.ReplaceDB(db)
	return p
}

type publishReqDeployProductEnvDo struct{ gen.DO }

type IPublishReqDeployProductEnvDo interface {
	gen.SubQuery
	Debug() IPublishReqDeployProductEnvDo
	WithContext(ctx context.Context) IPublishReqDeployProductEnvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublishReqDeployProductEnvDo
	WriteDB() IPublishReqDeployProductEnvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublishReqDeployProductEnvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublishReqDeployProductEnvDo
	Not(conds ...gen.Condition) IPublishReqDeployProductEnvDo
	Or(conds ...gen.Condition) IPublishReqDeployProductEnvDo
	Select(conds ...field.Expr) IPublishReqDeployProductEnvDo
	Where(conds ...gen.Condition) IPublishReqDeployProductEnvDo
	Order(conds ...field.Expr) IPublishReqDeployProductEnvDo
	Distinct(cols ...field.Expr) IPublishReqDeployProductEnvDo
	Omit(cols ...field.Expr) IPublishReqDeployProductEnvDo
	Join(table schema.Tabler, on ...field.Expr) IPublishReqDeployProductEnvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqDeployProductEnvDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqDeployProductEnvDo
	Group(cols ...field.Expr) IPublishReqDeployProductEnvDo
	Having(conds ...gen.Condition) IPublishReqDeployProductEnvDo
	Limit(limit int) IPublishReqDeployProductEnvDo
	Offset(offset int) IPublishReqDeployProductEnvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqDeployProductEnvDo
	Unscoped() IPublishReqDeployProductEnvDo
	Create(values ...*model.PublishReqDeployProductEnv) error
	CreateInBatches(values []*model.PublishReqDeployProductEnv, batchSize int) error
	Save(values ...*model.PublishReqDeployProductEnv) error
	First() (*model.PublishReqDeployProductEnv, error)
	Take() (*model.PublishReqDeployProductEnv, error)
	Last() (*model.PublishReqDeployProductEnv, error)
	Find() ([]*model.PublishReqDeployProductEnv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReqDeployProductEnv, err error)
	FindInBatches(result *[]*model.PublishReqDeployProductEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PublishReqDeployProductEnv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublishReqDeployProductEnvDo
	Assign(attrs ...field.AssignExpr) IPublishReqDeployProductEnvDo
	Joins(fields ...field.RelationField) IPublishReqDeployProductEnvDo
	Preload(fields ...field.RelationField) IPublishReqDeployProductEnvDo
	FirstOrInit() (*model.PublishReqDeployProductEnv, error)
	FirstOrCreate() (*model.PublishReqDeployProductEnv, error)
	FindByPage(offset int, limit int) (result []*model.PublishReqDeployProductEnv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublishReqDeployProductEnvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p publishReqDeployProductEnvDo) Debug() IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Debug())
}

func (p publishReqDeployProductEnvDo) WithContext(ctx context.Context) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publishReqDeployProductEnvDo) ReadDB() IPublishReqDeployProductEnvDo {
	return p.Clauses(dbresolver.Read)
}

func (p publishReqDeployProductEnvDo) WriteDB() IPublishReqDeployProductEnvDo {
	return p.Clauses(dbresolver.Write)
}

func (p publishReqDeployProductEnvDo) Session(config *gorm.Session) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Session(config))
}

func (p publishReqDeployProductEnvDo) Clauses(conds ...clause.Expression) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publishReqDeployProductEnvDo) Returning(value interface{}, columns ...string) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publishReqDeployProductEnvDo) Not(conds ...gen.Condition) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publishReqDeployProductEnvDo) Or(conds ...gen.Condition) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publishReqDeployProductEnvDo) Select(conds ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publishReqDeployProductEnvDo) Where(conds ...gen.Condition) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publishReqDeployProductEnvDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPublishReqDeployProductEnvDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publishReqDeployProductEnvDo) Order(conds ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publishReqDeployProductEnvDo) Distinct(cols ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publishReqDeployProductEnvDo) Omit(cols ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publishReqDeployProductEnvDo) Join(table schema.Tabler, on ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publishReqDeployProductEnvDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publishReqDeployProductEnvDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publishReqDeployProductEnvDo) Group(cols ...field.Expr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publishReqDeployProductEnvDo) Having(conds ...gen.Condition) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publishReqDeployProductEnvDo) Limit(limit int) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publishReqDeployProductEnvDo) Offset(offset int) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publishReqDeployProductEnvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publishReqDeployProductEnvDo) Unscoped() IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publishReqDeployProductEnvDo) Create(values ...*model.PublishReqDeployProductEnv) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publishReqDeployProductEnvDo) CreateInBatches(values []*model.PublishReqDeployProductEnv, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publishReqDeployProductEnvDo) Save(values ...*model.PublishReqDeployProductEnv) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publishReqDeployProductEnvDo) First() (*model.PublishReqDeployProductEnv, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqDeployProductEnv), nil
	}
}

func (p publishReqDeployProductEnvDo) Take() (*model.PublishReqDeployProductEnv, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqDeployProductEnv), nil
	}
}

func (p publishReqDeployProductEnvDo) Last() (*model.PublishReqDeployProductEnv, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqDeployProductEnv), nil
	}
}

func (p publishReqDeployProductEnvDo) Find() ([]*model.PublishReqDeployProductEnv, error) {
	result, err := p.DO.Find()
	return result.([]*model.PublishReqDeployProductEnv), err
}

func (p publishReqDeployProductEnvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReqDeployProductEnv, err error) {
	buf := make([]*model.PublishReqDeployProductEnv, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publishReqDeployProductEnvDo) FindInBatches(result *[]*model.PublishReqDeployProductEnv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publishReqDeployProductEnvDo) Attrs(attrs ...field.AssignExpr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publishReqDeployProductEnvDo) Assign(attrs ...field.AssignExpr) IPublishReqDeployProductEnvDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publishReqDeployProductEnvDo) Joins(fields ...field.RelationField) IPublishReqDeployProductEnvDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publishReqDeployProductEnvDo) Preload(fields ...field.RelationField) IPublishReqDeployProductEnvDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publishReqDeployProductEnvDo) FirstOrInit() (*model.PublishReqDeployProductEnv, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqDeployProductEnv), nil
	}
}

func (p publishReqDeployProductEnvDo) FirstOrCreate() (*model.PublishReqDeployProductEnv, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReqDeployProductEnv), nil
	}
}

func (p publishReqDeployProductEnvDo) FindByPage(offset int, limit int) (result []*model.PublishReqDeployProductEnv, count int64, err error) {
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

func (p publishReqDeployProductEnvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publishReqDeployProductEnvDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publishReqDeployProductEnvDo) Delete(models ...*model.PublishReqDeployProductEnv) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publishReqDeployProductEnvDo) withDO(do gen.Dao) *publishReqDeployProductEnvDo {
	p.DO = *do.(*gen.DO)
	return p
}