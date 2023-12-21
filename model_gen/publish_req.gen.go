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

func newPublishReq(db *gorm.DB, opts ...gen.DOOption) publishReq {
	_publishReq := publishReq{}

	_publishReq.publishReqDo.UseDB(db, opts...)
	_publishReq.publishReqDo.UseModel(&model.PublishReq{})

	tableName := _publishReq.publishReqDo.TableName()
	_publishReq.ALL = field.NewAsterisk(tableName)
	_publishReq.ReqID = field.NewString(tableName, "req_id")
	_publishReq.ReqSummary = field.NewString(tableName, "req_summary")
	_publishReq.DevUser = field.NewString(tableName, "dev_user")
	_publishReq.TestUser = field.NewString(tableName, "test_user")
	_publishReq.PublishVersion = field.NewString(tableName, "publish_version")
	_publishReq.ReqStatu = field.NewString(tableName, "req_statu")
	_publishReq.ProductDomain = field.NewString(tableName, "product_domain")
	_publishReq.SysProjectID = field.NewString(tableName, "sys_project_id")
	_publishReq.CreateTime = field.NewTime(tableName, "create_time")
	_publishReq.ModifyTime = field.NewTime(tableName, "modify_time")
	_publishReq.Comments = field.NewString(tableName, "comments")
	_publishReq.VersionType = field.NewString(tableName, "version_type")
	_publishReq.IsDelete = field.NewString(tableName, "is_delete")
	_publishReq.CreateUser = field.NewString(tableName, "create_user")
	_publishReq.ModifyUser = field.NewString(tableName, "modify_user")
	_publishReq.Num = field.NewInt32(tableName, "num")
	_publishReq.LastDeployTime = field.NewTime(tableName, "last_deploy_time")
	_publishReq.LastRollbackTime = field.NewTime(tableName, "last_rollback_time")
	_publishReq.PrepareDeployTime = field.NewTime(tableName, "prepare_deploy_time")
	_publishReq.DbscriptEventID = field.NewString(tableName, "dbscript_event_id")
	_publishReq.DeployTime = field.NewTime(tableName, "deploy_time")
	_publishReq.PlanPublishTime = field.NewTime(tableName, "plan_publish_time")
	_publishReq.InnerTest = field.NewString(tableName, "inner_test")
	_publishReq.DevExpectFinishTime = field.NewString(tableName, "dev_expect_finish_time")
	_publishReq.TestExpectFinishTime = field.NewString(tableName, "test_expect_finish_time")
	_publishReq.TestServerEnvID = field.NewString(tableName, "test_server_env_id")
	_publishReq.PublishTestEnv = field.NewString(tableName, "publish_test_env")
	_publishReq.ProductFrom = field.NewString(tableName, "product_from")
	_publishReq.IsProduct = field.NewString(tableName, "is_product")
	_publishReq.DutyTeam = field.NewString(tableName, "duty_team")
	_publishReq.RequireFinishTime = field.NewString(tableName, "require_finish_time")

	_publishReq.fillFieldMap()

	return _publishReq
}

type publishReq struct {
	publishReqDo

	ALL                  field.Asterisk
	ReqID                field.String
	ReqSummary           field.String // 升级单概述
	DevUser              field.String
	TestUser             field.String
	PublishVersion       field.String
	ReqStatu             field.String
	ProductDomain        field.String // 产品域
	SysProjectID         field.String
	CreateTime           field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime           field.Time   // 记录修改时间（数据库自动写入）
	Comments             field.String // 备注说明
	VersionType          field.String
	IsDelete             field.String // 是否删除
	CreateUser           field.String
	ModifyUser           field.String
	Num                  field.Int32
	LastDeployTime       field.Time
	LastRollbackTime     field.Time
	PrepareDeployTime    field.Time
	DbscriptEventID      field.String
	DeployTime           field.Time // 升级单在点击部署时记录升级单部署时间到数据库，虽然可能不是确切的时间（因为有可能有回退或者有部署失败）但是基本上可以作为参考实际部署时间；
	PlanPublishTime      field.Time
	InnerTest            field.String // 内部测试
	DevExpectFinishTime  field.String
	TestExpectFinishTime field.String
	TestServerEnvID      field.String
	PublishTestEnv       field.String
	ProductFrom          field.String
	IsProduct            field.String
	DutyTeam             field.String // 责任团队
	RequireFinishTime    field.String

	fieldMap map[string]field.Expr
}

func (p publishReq) Table(newTableName string) *publishReq {
	p.publishReqDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publishReq) As(alias string) *publishReq {
	p.publishReqDo.DO = *(p.publishReqDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publishReq) updateTableName(table string) *publishReq {
	p.ALL = field.NewAsterisk(table)
	p.ReqID = field.NewString(table, "req_id")
	p.ReqSummary = field.NewString(table, "req_summary")
	p.DevUser = field.NewString(table, "dev_user")
	p.TestUser = field.NewString(table, "test_user")
	p.PublishVersion = field.NewString(table, "publish_version")
	p.ReqStatu = field.NewString(table, "req_statu")
	p.ProductDomain = field.NewString(table, "product_domain")
	p.SysProjectID = field.NewString(table, "sys_project_id")
	p.CreateTime = field.NewTime(table, "create_time")
	p.ModifyTime = field.NewTime(table, "modify_time")
	p.Comments = field.NewString(table, "comments")
	p.VersionType = field.NewString(table, "version_type")
	p.IsDelete = field.NewString(table, "is_delete")
	p.CreateUser = field.NewString(table, "create_user")
	p.ModifyUser = field.NewString(table, "modify_user")
	p.Num = field.NewInt32(table, "num")
	p.LastDeployTime = field.NewTime(table, "last_deploy_time")
	p.LastRollbackTime = field.NewTime(table, "last_rollback_time")
	p.PrepareDeployTime = field.NewTime(table, "prepare_deploy_time")
	p.DbscriptEventID = field.NewString(table, "dbscript_event_id")
	p.DeployTime = field.NewTime(table, "deploy_time")
	p.PlanPublishTime = field.NewTime(table, "plan_publish_time")
	p.InnerTest = field.NewString(table, "inner_test")
	p.DevExpectFinishTime = field.NewString(table, "dev_expect_finish_time")
	p.TestExpectFinishTime = field.NewString(table, "test_expect_finish_time")
	p.TestServerEnvID = field.NewString(table, "test_server_env_id")
	p.PublishTestEnv = field.NewString(table, "publish_test_env")
	p.ProductFrom = field.NewString(table, "product_from")
	p.IsProduct = field.NewString(table, "is_product")
	p.DutyTeam = field.NewString(table, "duty_team")
	p.RequireFinishTime = field.NewString(table, "require_finish_time")

	p.fillFieldMap()

	return p
}

func (p *publishReq) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publishReq) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 31)
	p.fieldMap["req_id"] = p.ReqID
	p.fieldMap["req_summary"] = p.ReqSummary
	p.fieldMap["dev_user"] = p.DevUser
	p.fieldMap["test_user"] = p.TestUser
	p.fieldMap["publish_version"] = p.PublishVersion
	p.fieldMap["req_statu"] = p.ReqStatu
	p.fieldMap["product_domain"] = p.ProductDomain
	p.fieldMap["sys_project_id"] = p.SysProjectID
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["modify_time"] = p.ModifyTime
	p.fieldMap["comments"] = p.Comments
	p.fieldMap["version_type"] = p.VersionType
	p.fieldMap["is_delete"] = p.IsDelete
	p.fieldMap["create_user"] = p.CreateUser
	p.fieldMap["modify_user"] = p.ModifyUser
	p.fieldMap["num"] = p.Num
	p.fieldMap["last_deploy_time"] = p.LastDeployTime
	p.fieldMap["last_rollback_time"] = p.LastRollbackTime
	p.fieldMap["prepare_deploy_time"] = p.PrepareDeployTime
	p.fieldMap["dbscript_event_id"] = p.DbscriptEventID
	p.fieldMap["deploy_time"] = p.DeployTime
	p.fieldMap["plan_publish_time"] = p.PlanPublishTime
	p.fieldMap["inner_test"] = p.InnerTest
	p.fieldMap["dev_expect_finish_time"] = p.DevExpectFinishTime
	p.fieldMap["test_expect_finish_time"] = p.TestExpectFinishTime
	p.fieldMap["test_server_env_id"] = p.TestServerEnvID
	p.fieldMap["publish_test_env"] = p.PublishTestEnv
	p.fieldMap["product_from"] = p.ProductFrom
	p.fieldMap["is_product"] = p.IsProduct
	p.fieldMap["duty_team"] = p.DutyTeam
	p.fieldMap["require_finish_time"] = p.RequireFinishTime
}

func (p publishReq) clone(db *gorm.DB) publishReq {
	p.publishReqDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publishReq) replaceDB(db *gorm.DB) publishReq {
	p.publishReqDo.ReplaceDB(db)
	return p
}

type publishReqDo struct{ gen.DO }

type IPublishReqDo interface {
	gen.SubQuery
	Debug() IPublishReqDo
	WithContext(ctx context.Context) IPublishReqDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublishReqDo
	WriteDB() IPublishReqDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublishReqDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublishReqDo
	Not(conds ...gen.Condition) IPublishReqDo
	Or(conds ...gen.Condition) IPublishReqDo
	Select(conds ...field.Expr) IPublishReqDo
	Where(conds ...gen.Condition) IPublishReqDo
	Order(conds ...field.Expr) IPublishReqDo
	Distinct(cols ...field.Expr) IPublishReqDo
	Omit(cols ...field.Expr) IPublishReqDo
	Join(table schema.Tabler, on ...field.Expr) IPublishReqDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqDo
	Group(cols ...field.Expr) IPublishReqDo
	Having(conds ...gen.Condition) IPublishReqDo
	Limit(limit int) IPublishReqDo
	Offset(offset int) IPublishReqDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqDo
	Unscoped() IPublishReqDo
	Create(values ...*model.PublishReq) error
	CreateInBatches(values []*model.PublishReq, batchSize int) error
	Save(values ...*model.PublishReq) error
	First() (*model.PublishReq, error)
	Take() (*model.PublishReq, error)
	Last() (*model.PublishReq, error)
	Find() ([]*model.PublishReq, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReq, err error)
	FindInBatches(result *[]*model.PublishReq, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PublishReq) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublishReqDo
	Assign(attrs ...field.AssignExpr) IPublishReqDo
	Joins(fields ...field.RelationField) IPublishReqDo
	Preload(fields ...field.RelationField) IPublishReqDo
	FirstOrInit() (*model.PublishReq, error)
	FirstOrCreate() (*model.PublishReq, error)
	FindByPage(offset int, limit int) (result []*model.PublishReq, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublishReqDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p publishReqDo) Debug() IPublishReqDo {
	return p.withDO(p.DO.Debug())
}

func (p publishReqDo) WithContext(ctx context.Context) IPublishReqDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publishReqDo) ReadDB() IPublishReqDo {
	return p.Clauses(dbresolver.Read)
}

func (p publishReqDo) WriteDB() IPublishReqDo {
	return p.Clauses(dbresolver.Write)
}

func (p publishReqDo) Session(config *gorm.Session) IPublishReqDo {
	return p.withDO(p.DO.Session(config))
}

func (p publishReqDo) Clauses(conds ...clause.Expression) IPublishReqDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publishReqDo) Returning(value interface{}, columns ...string) IPublishReqDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publishReqDo) Not(conds ...gen.Condition) IPublishReqDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publishReqDo) Or(conds ...gen.Condition) IPublishReqDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publishReqDo) Select(conds ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publishReqDo) Where(conds ...gen.Condition) IPublishReqDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publishReqDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPublishReqDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publishReqDo) Order(conds ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publishReqDo) Distinct(cols ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publishReqDo) Omit(cols ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publishReqDo) Join(table schema.Tabler, on ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publishReqDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publishReqDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publishReqDo) Group(cols ...field.Expr) IPublishReqDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publishReqDo) Having(conds ...gen.Condition) IPublishReqDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publishReqDo) Limit(limit int) IPublishReqDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publishReqDo) Offset(offset int) IPublishReqDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publishReqDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublishReqDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publishReqDo) Unscoped() IPublishReqDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publishReqDo) Create(values ...*model.PublishReq) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publishReqDo) CreateInBatches(values []*model.PublishReq, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publishReqDo) Save(values ...*model.PublishReq) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publishReqDo) First() (*model.PublishReq, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReq), nil
	}
}

func (p publishReqDo) Take() (*model.PublishReq, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReq), nil
	}
}

func (p publishReqDo) Last() (*model.PublishReq, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReq), nil
	}
}

func (p publishReqDo) Find() ([]*model.PublishReq, error) {
	result, err := p.DO.Find()
	return result.([]*model.PublishReq), err
}

func (p publishReqDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublishReq, err error) {
	buf := make([]*model.PublishReq, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publishReqDo) FindInBatches(result *[]*model.PublishReq, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publishReqDo) Attrs(attrs ...field.AssignExpr) IPublishReqDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publishReqDo) Assign(attrs ...field.AssignExpr) IPublishReqDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publishReqDo) Joins(fields ...field.RelationField) IPublishReqDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publishReqDo) Preload(fields ...field.RelationField) IPublishReqDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publishReqDo) FirstOrInit() (*model.PublishReq, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReq), nil
	}
}

func (p publishReqDo) FirstOrCreate() (*model.PublishReq, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublishReq), nil
	}
}

func (p publishReqDo) FindByPage(offset int, limit int) (result []*model.PublishReq, count int64, err error) {
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

func (p publishReqDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publishReqDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publishReqDo) Delete(models ...*model.PublishReq) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publishReqDo) withDO(do gen.Dao) *publishReqDo {
	p.DO = *do.(*gen.DO)
	return p
}
