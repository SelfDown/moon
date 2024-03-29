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

func newEnvVariable(db *gorm.DB, opts ...gen.DOOption) envVariable {
	_envVariable := envVariable{}

	_envVariable.envVariableDo.UseDB(db, opts...)
	_envVariable.envVariableDo.UseModel(&model.EnvVariable{})

	tableName := _envVariable.envVariableDo.TableName()
	_envVariable.ALL = field.NewAsterisk(tableName)
	_envVariable.ParamID = field.NewString(tableName, "param_id")
	_envVariable.ParamKey = field.NewString(tableName, "param_key")
	_envVariable.ParamValue = field.NewString(tableName, "param_value")
	_envVariable.ServerEnvID = field.NewString(tableName, "server_env_id")
	_envVariable.ParamNotes = field.NewString(tableName, "param_notes")
	_envVariable.CreateTime = field.NewTime(tableName, "create_time")
	_envVariable.ModifyTime = field.NewTime(tableName, "modify_time")
	_envVariable.Comments = field.NewString(tableName, "comments")

	_envVariable.fillFieldMap()

	return _envVariable
}

type envVariable struct {
	envVariableDo

	ALL         field.Asterisk
	ParamID     field.String
	ParamKey    field.String
	ParamValue  field.String
	ServerEnvID field.String
	ParamNotes  field.String
	CreateTime  field.Time
	ModifyTime  field.Time
	Comments    field.String

	fieldMap map[string]field.Expr
}

func (e envVariable) Table(newTableName string) *envVariable {
	e.envVariableDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e envVariable) As(alias string) *envVariable {
	e.envVariableDo.DO = *(e.envVariableDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *envVariable) updateTableName(table string) *envVariable {
	e.ALL = field.NewAsterisk(table)
	e.ParamID = field.NewString(table, "param_id")
	e.ParamKey = field.NewString(table, "param_key")
	e.ParamValue = field.NewString(table, "param_value")
	e.ServerEnvID = field.NewString(table, "server_env_id")
	e.ParamNotes = field.NewString(table, "param_notes")
	e.CreateTime = field.NewTime(table, "create_time")
	e.ModifyTime = field.NewTime(table, "modify_time")
	e.Comments = field.NewString(table, "comments")

	e.fillFieldMap()

	return e
}

func (e *envVariable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *envVariable) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 8)
	e.fieldMap["param_id"] = e.ParamID
	e.fieldMap["param_key"] = e.ParamKey
	e.fieldMap["param_value"] = e.ParamValue
	e.fieldMap["server_env_id"] = e.ServerEnvID
	e.fieldMap["param_notes"] = e.ParamNotes
	e.fieldMap["create_time"] = e.CreateTime
	e.fieldMap["modify_time"] = e.ModifyTime
	e.fieldMap["comments"] = e.Comments
}

func (e envVariable) clone(db *gorm.DB) envVariable {
	e.envVariableDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e envVariable) replaceDB(db *gorm.DB) envVariable {
	e.envVariableDo.ReplaceDB(db)
	return e
}

type envVariableDo struct{ gen.DO }

type IEnvVariableDo interface {
	gen.SubQuery
	Debug() IEnvVariableDo
	WithContext(ctx context.Context) IEnvVariableDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEnvVariableDo
	WriteDB() IEnvVariableDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEnvVariableDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEnvVariableDo
	Not(conds ...gen.Condition) IEnvVariableDo
	Or(conds ...gen.Condition) IEnvVariableDo
	Select(conds ...field.Expr) IEnvVariableDo
	Where(conds ...gen.Condition) IEnvVariableDo
	Order(conds ...field.Expr) IEnvVariableDo
	Distinct(cols ...field.Expr) IEnvVariableDo
	Omit(cols ...field.Expr) IEnvVariableDo
	Join(table schema.Tabler, on ...field.Expr) IEnvVariableDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEnvVariableDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEnvVariableDo
	Group(cols ...field.Expr) IEnvVariableDo
	Having(conds ...gen.Condition) IEnvVariableDo
	Limit(limit int) IEnvVariableDo
	Offset(offset int) IEnvVariableDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEnvVariableDo
	Unscoped() IEnvVariableDo
	Create(values ...*model.EnvVariable) error
	CreateInBatches(values []*model.EnvVariable, batchSize int) error
	Save(values ...*model.EnvVariable) error
	First() (*model.EnvVariable, error)
	Take() (*model.EnvVariable, error)
	Last() (*model.EnvVariable, error)
	Find() ([]*model.EnvVariable, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EnvVariable, err error)
	FindInBatches(result *[]*model.EnvVariable, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.EnvVariable) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEnvVariableDo
	Assign(attrs ...field.AssignExpr) IEnvVariableDo
	Joins(fields ...field.RelationField) IEnvVariableDo
	Preload(fields ...field.RelationField) IEnvVariableDo
	FirstOrInit() (*model.EnvVariable, error)
	FirstOrCreate() (*model.EnvVariable, error)
	FindByPage(offset int, limit int) (result []*model.EnvVariable, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEnvVariableDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e envVariableDo) Debug() IEnvVariableDo {
	return e.withDO(e.DO.Debug())
}

func (e envVariableDo) WithContext(ctx context.Context) IEnvVariableDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e envVariableDo) ReadDB() IEnvVariableDo {
	return e.Clauses(dbresolver.Read)
}

func (e envVariableDo) WriteDB() IEnvVariableDo {
	return e.Clauses(dbresolver.Write)
}

func (e envVariableDo) Session(config *gorm.Session) IEnvVariableDo {
	return e.withDO(e.DO.Session(config))
}

func (e envVariableDo) Clauses(conds ...clause.Expression) IEnvVariableDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e envVariableDo) Returning(value interface{}, columns ...string) IEnvVariableDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e envVariableDo) Not(conds ...gen.Condition) IEnvVariableDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e envVariableDo) Or(conds ...gen.Condition) IEnvVariableDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e envVariableDo) Select(conds ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e envVariableDo) Where(conds ...gen.Condition) IEnvVariableDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e envVariableDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IEnvVariableDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e envVariableDo) Order(conds ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e envVariableDo) Distinct(cols ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e envVariableDo) Omit(cols ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e envVariableDo) Join(table schema.Tabler, on ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e envVariableDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e envVariableDo) RightJoin(table schema.Tabler, on ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e envVariableDo) Group(cols ...field.Expr) IEnvVariableDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e envVariableDo) Having(conds ...gen.Condition) IEnvVariableDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e envVariableDo) Limit(limit int) IEnvVariableDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e envVariableDo) Offset(offset int) IEnvVariableDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e envVariableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEnvVariableDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e envVariableDo) Unscoped() IEnvVariableDo {
	return e.withDO(e.DO.Unscoped())
}

func (e envVariableDo) Create(values ...*model.EnvVariable) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e envVariableDo) CreateInBatches(values []*model.EnvVariable, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e envVariableDo) Save(values ...*model.EnvVariable) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e envVariableDo) First() (*model.EnvVariable, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnvVariable), nil
	}
}

func (e envVariableDo) Take() (*model.EnvVariable, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnvVariable), nil
	}
}

func (e envVariableDo) Last() (*model.EnvVariable, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnvVariable), nil
	}
}

func (e envVariableDo) Find() ([]*model.EnvVariable, error) {
	result, err := e.DO.Find()
	return result.([]*model.EnvVariable), err
}

func (e envVariableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.EnvVariable, err error) {
	buf := make([]*model.EnvVariable, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e envVariableDo) FindInBatches(result *[]*model.EnvVariable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e envVariableDo) Attrs(attrs ...field.AssignExpr) IEnvVariableDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e envVariableDo) Assign(attrs ...field.AssignExpr) IEnvVariableDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e envVariableDo) Joins(fields ...field.RelationField) IEnvVariableDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e envVariableDo) Preload(fields ...field.RelationField) IEnvVariableDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e envVariableDo) FirstOrInit() (*model.EnvVariable, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnvVariable), nil
	}
}

func (e envVariableDo) FirstOrCreate() (*model.EnvVariable, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.EnvVariable), nil
	}
}

func (e envVariableDo) FindByPage(offset int, limit int) (result []*model.EnvVariable, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e envVariableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e envVariableDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e envVariableDo) Delete(models ...*model.EnvVariable) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *envVariableDo) withDO(do gen.Dao) *envVariableDo {
	e.DO = *do.(*gen.DO)
	return e
}
