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

func newAutoDeploy(db *gorm.DB, opts ...gen.DOOption) autoDeploy {
	_autoDeploy := autoDeploy{}

	_autoDeploy.autoDeployDo.UseDB(db, opts...)
	_autoDeploy.autoDeployDo.UseModel(&model.AutoDeploy{})

	tableName := _autoDeploy.autoDeployDo.TableName()
	_autoDeploy.ALL = field.NewAsterisk(tableName)
	_autoDeploy.AutoBuildID = field.NewString(tableName, "auto_build_id")
	_autoDeploy.BuildGroupID = field.NewString(tableName, "build_group_id")
	_autoDeploy.ProjectCode = field.NewString(tableName, "project_code")
	_autoDeploy.ServerEnvCode = field.NewString(tableName, "server_env_code")
	_autoDeploy.WarArtifactid = field.NewString(tableName, "war_artifactid")
	_autoDeploy.WarGroupid = field.NewString(tableName, "war_groupid")
	_autoDeploy.Version = field.NewString(tableName, "version")
	_autoDeploy.CreateTime = field.NewString(tableName, "create_time")
	_autoDeploy.CreateUser = field.NewString(tableName, "create_user")
	_autoDeploy.Msg = field.NewString(tableName, "msg")
	_autoDeploy.DeployStatus = field.NewString(tableName, "deploy_status")

	_autoDeploy.fillFieldMap()

	return _autoDeploy
}

type autoDeploy struct {
	autoDeployDo

	ALL           field.Asterisk
	AutoBuildID   field.String
	BuildGroupID  field.String
	ProjectCode   field.String // 服务器ID
	ServerEnvCode field.String
	WarArtifactid field.String // 服务器IP
	WarGroupid    field.String
	Version       field.String
	CreateTime    field.String
	CreateUser    field.String
	Msg           field.String
	DeployStatus  field.String

	fieldMap map[string]field.Expr
}

func (a autoDeploy) Table(newTableName string) *autoDeploy {
	a.autoDeployDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a autoDeploy) As(alias string) *autoDeploy {
	a.autoDeployDo.DO = *(a.autoDeployDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *autoDeploy) updateTableName(table string) *autoDeploy {
	a.ALL = field.NewAsterisk(table)
	a.AutoBuildID = field.NewString(table, "auto_build_id")
	a.BuildGroupID = field.NewString(table, "build_group_id")
	a.ProjectCode = field.NewString(table, "project_code")
	a.ServerEnvCode = field.NewString(table, "server_env_code")
	a.WarArtifactid = field.NewString(table, "war_artifactid")
	a.WarGroupid = field.NewString(table, "war_groupid")
	a.Version = field.NewString(table, "version")
	a.CreateTime = field.NewString(table, "create_time")
	a.CreateUser = field.NewString(table, "create_user")
	a.Msg = field.NewString(table, "msg")
	a.DeployStatus = field.NewString(table, "deploy_status")

	a.fillFieldMap()

	return a
}

func (a *autoDeploy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *autoDeploy) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["auto_build_id"] = a.AutoBuildID
	a.fieldMap["build_group_id"] = a.BuildGroupID
	a.fieldMap["project_code"] = a.ProjectCode
	a.fieldMap["server_env_code"] = a.ServerEnvCode
	a.fieldMap["war_artifactid"] = a.WarArtifactid
	a.fieldMap["war_groupid"] = a.WarGroupid
	a.fieldMap["version"] = a.Version
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["create_user"] = a.CreateUser
	a.fieldMap["msg"] = a.Msg
	a.fieldMap["deploy_status"] = a.DeployStatus
}

func (a autoDeploy) clone(db *gorm.DB) autoDeploy {
	a.autoDeployDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a autoDeploy) replaceDB(db *gorm.DB) autoDeploy {
	a.autoDeployDo.ReplaceDB(db)
	return a
}

type autoDeployDo struct{ gen.DO }

type IAutoDeployDo interface {
	gen.SubQuery
	Debug() IAutoDeployDo
	WithContext(ctx context.Context) IAutoDeployDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAutoDeployDo
	WriteDB() IAutoDeployDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAutoDeployDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAutoDeployDo
	Not(conds ...gen.Condition) IAutoDeployDo
	Or(conds ...gen.Condition) IAutoDeployDo
	Select(conds ...field.Expr) IAutoDeployDo
	Where(conds ...gen.Condition) IAutoDeployDo
	Order(conds ...field.Expr) IAutoDeployDo
	Distinct(cols ...field.Expr) IAutoDeployDo
	Omit(cols ...field.Expr) IAutoDeployDo
	Join(table schema.Tabler, on ...field.Expr) IAutoDeployDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAutoDeployDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAutoDeployDo
	Group(cols ...field.Expr) IAutoDeployDo
	Having(conds ...gen.Condition) IAutoDeployDo
	Limit(limit int) IAutoDeployDo
	Offset(offset int) IAutoDeployDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAutoDeployDo
	Unscoped() IAutoDeployDo
	Create(values ...*model.AutoDeploy) error
	CreateInBatches(values []*model.AutoDeploy, batchSize int) error
	Save(values ...*model.AutoDeploy) error
	First() (*model.AutoDeploy, error)
	Take() (*model.AutoDeploy, error)
	Last() (*model.AutoDeploy, error)
	Find() ([]*model.AutoDeploy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AutoDeploy, err error)
	FindInBatches(result *[]*model.AutoDeploy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AutoDeploy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAutoDeployDo
	Assign(attrs ...field.AssignExpr) IAutoDeployDo
	Joins(fields ...field.RelationField) IAutoDeployDo
	Preload(fields ...field.RelationField) IAutoDeployDo
	FirstOrInit() (*model.AutoDeploy, error)
	FirstOrCreate() (*model.AutoDeploy, error)
	FindByPage(offset int, limit int) (result []*model.AutoDeploy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAutoDeployDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a autoDeployDo) Debug() IAutoDeployDo {
	return a.withDO(a.DO.Debug())
}

func (a autoDeployDo) WithContext(ctx context.Context) IAutoDeployDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a autoDeployDo) ReadDB() IAutoDeployDo {
	return a.Clauses(dbresolver.Read)
}

func (a autoDeployDo) WriteDB() IAutoDeployDo {
	return a.Clauses(dbresolver.Write)
}

func (a autoDeployDo) Session(config *gorm.Session) IAutoDeployDo {
	return a.withDO(a.DO.Session(config))
}

func (a autoDeployDo) Clauses(conds ...clause.Expression) IAutoDeployDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a autoDeployDo) Returning(value interface{}, columns ...string) IAutoDeployDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a autoDeployDo) Not(conds ...gen.Condition) IAutoDeployDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a autoDeployDo) Or(conds ...gen.Condition) IAutoDeployDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a autoDeployDo) Select(conds ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a autoDeployDo) Where(conds ...gen.Condition) IAutoDeployDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a autoDeployDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAutoDeployDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a autoDeployDo) Order(conds ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a autoDeployDo) Distinct(cols ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a autoDeployDo) Omit(cols ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a autoDeployDo) Join(table schema.Tabler, on ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a autoDeployDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a autoDeployDo) RightJoin(table schema.Tabler, on ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a autoDeployDo) Group(cols ...field.Expr) IAutoDeployDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a autoDeployDo) Having(conds ...gen.Condition) IAutoDeployDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a autoDeployDo) Limit(limit int) IAutoDeployDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a autoDeployDo) Offset(offset int) IAutoDeployDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a autoDeployDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAutoDeployDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a autoDeployDo) Unscoped() IAutoDeployDo {
	return a.withDO(a.DO.Unscoped())
}

func (a autoDeployDo) Create(values ...*model.AutoDeploy) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a autoDeployDo) CreateInBatches(values []*model.AutoDeploy, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a autoDeployDo) Save(values ...*model.AutoDeploy) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a autoDeployDo) First() (*model.AutoDeploy, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoDeploy), nil
	}
}

func (a autoDeployDo) Take() (*model.AutoDeploy, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoDeploy), nil
	}
}

func (a autoDeployDo) Last() (*model.AutoDeploy, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoDeploy), nil
	}
}

func (a autoDeployDo) Find() ([]*model.AutoDeploy, error) {
	result, err := a.DO.Find()
	return result.([]*model.AutoDeploy), err
}

func (a autoDeployDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AutoDeploy, err error) {
	buf := make([]*model.AutoDeploy, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a autoDeployDo) FindInBatches(result *[]*model.AutoDeploy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a autoDeployDo) Attrs(attrs ...field.AssignExpr) IAutoDeployDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a autoDeployDo) Assign(attrs ...field.AssignExpr) IAutoDeployDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a autoDeployDo) Joins(fields ...field.RelationField) IAutoDeployDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a autoDeployDo) Preload(fields ...field.RelationField) IAutoDeployDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a autoDeployDo) FirstOrInit() (*model.AutoDeploy, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoDeploy), nil
	}
}

func (a autoDeployDo) FirstOrCreate() (*model.AutoDeploy, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AutoDeploy), nil
	}
}

func (a autoDeployDo) FindByPage(offset int, limit int) (result []*model.AutoDeploy, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a autoDeployDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a autoDeployDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a autoDeployDo) Delete(models ...*model.AutoDeploy) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *autoDeployDo) withDO(do gen.Dao) *autoDeployDo {
	a.DO = *do.(*gen.DO)
	return a
}