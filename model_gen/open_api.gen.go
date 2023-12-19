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

func newOpenAPI(db *gorm.DB, opts ...gen.DOOption) openAPI {
	_openAPI := openAPI{}

	_openAPI.openAPIDo.UseDB(db, opts...)
	_openAPI.openAPIDo.UseModel(&model.OpenAPI{})

	tableName := _openAPI.openAPIDo.TableName()
	_openAPI.ALL = field.NewAsterisk(tableName)
	_openAPI.OpenAPIID = field.NewString(tableName, "open_api_id")
	_openAPI.OpenAPIURI = field.NewString(tableName, "open_api_uri")
	_openAPI.OpenAPIName = field.NewString(tableName, "open_api_name")
	_openAPI.CreateTime = field.NewTime(tableName, "create_time")
	_openAPI.ModifyTime = field.NewTime(tableName, "modify_time")
	_openAPI.Comments = field.NewString(tableName, "comments")

	_openAPI.fillFieldMap()

	return _openAPI
}

type openAPI struct {
	openAPIDo

	ALL       field.Asterisk
	OpenAPIID field.String // UUID
	/*
		API资源地址
		比如：/open/jboss?api_key=xxx&action=restart&ip={zabbix host ip}
	*/
	OpenAPIURI  field.String
	OpenAPIName field.String // API名称
	CreateTime  field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime  field.Time   // 记录修改时间（数据库自动写入）
	Comments    field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (o openAPI) Table(newTableName string) *openAPI {
	o.openAPIDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o openAPI) As(alias string) *openAPI {
	o.openAPIDo.DO = *(o.openAPIDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *openAPI) updateTableName(table string) *openAPI {
	o.ALL = field.NewAsterisk(table)
	o.OpenAPIID = field.NewString(table, "open_api_id")
	o.OpenAPIURI = field.NewString(table, "open_api_uri")
	o.OpenAPIName = field.NewString(table, "open_api_name")
	o.CreateTime = field.NewTime(table, "create_time")
	o.ModifyTime = field.NewTime(table, "modify_time")
	o.Comments = field.NewString(table, "comments")

	o.fillFieldMap()

	return o
}

func (o *openAPI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *openAPI) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 6)
	o.fieldMap["open_api_id"] = o.OpenAPIID
	o.fieldMap["open_api_uri"] = o.OpenAPIURI
	o.fieldMap["open_api_name"] = o.OpenAPIName
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["modify_time"] = o.ModifyTime
	o.fieldMap["comments"] = o.Comments
}

func (o openAPI) clone(db *gorm.DB) openAPI {
	o.openAPIDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o openAPI) replaceDB(db *gorm.DB) openAPI {
	o.openAPIDo.ReplaceDB(db)
	return o
}

type openAPIDo struct{ gen.DO }

type IOpenAPIDo interface {
	gen.SubQuery
	Debug() IOpenAPIDo
	WithContext(ctx context.Context) IOpenAPIDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOpenAPIDo
	WriteDB() IOpenAPIDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOpenAPIDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOpenAPIDo
	Not(conds ...gen.Condition) IOpenAPIDo
	Or(conds ...gen.Condition) IOpenAPIDo
	Select(conds ...field.Expr) IOpenAPIDo
	Where(conds ...gen.Condition) IOpenAPIDo
	Order(conds ...field.Expr) IOpenAPIDo
	Distinct(cols ...field.Expr) IOpenAPIDo
	Omit(cols ...field.Expr) IOpenAPIDo
	Join(table schema.Tabler, on ...field.Expr) IOpenAPIDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOpenAPIDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOpenAPIDo
	Group(cols ...field.Expr) IOpenAPIDo
	Having(conds ...gen.Condition) IOpenAPIDo
	Limit(limit int) IOpenAPIDo
	Offset(offset int) IOpenAPIDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOpenAPIDo
	Unscoped() IOpenAPIDo
	Create(values ...*model.OpenAPI) error
	CreateInBatches(values []*model.OpenAPI, batchSize int) error
	Save(values ...*model.OpenAPI) error
	First() (*model.OpenAPI, error)
	Take() (*model.OpenAPI, error)
	Last() (*model.OpenAPI, error)
	Find() ([]*model.OpenAPI, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OpenAPI, err error)
	FindInBatches(result *[]*model.OpenAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OpenAPI) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOpenAPIDo
	Assign(attrs ...field.AssignExpr) IOpenAPIDo
	Joins(fields ...field.RelationField) IOpenAPIDo
	Preload(fields ...field.RelationField) IOpenAPIDo
	FirstOrInit() (*model.OpenAPI, error)
	FirstOrCreate() (*model.OpenAPI, error)
	FindByPage(offset int, limit int) (result []*model.OpenAPI, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOpenAPIDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o openAPIDo) Debug() IOpenAPIDo {
	return o.withDO(o.DO.Debug())
}

func (o openAPIDo) WithContext(ctx context.Context) IOpenAPIDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o openAPIDo) ReadDB() IOpenAPIDo {
	return o.Clauses(dbresolver.Read)
}

func (o openAPIDo) WriteDB() IOpenAPIDo {
	return o.Clauses(dbresolver.Write)
}

func (o openAPIDo) Session(config *gorm.Session) IOpenAPIDo {
	return o.withDO(o.DO.Session(config))
}

func (o openAPIDo) Clauses(conds ...clause.Expression) IOpenAPIDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o openAPIDo) Returning(value interface{}, columns ...string) IOpenAPIDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o openAPIDo) Not(conds ...gen.Condition) IOpenAPIDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o openAPIDo) Or(conds ...gen.Condition) IOpenAPIDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o openAPIDo) Select(conds ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o openAPIDo) Where(conds ...gen.Condition) IOpenAPIDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o openAPIDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOpenAPIDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o openAPIDo) Order(conds ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o openAPIDo) Distinct(cols ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o openAPIDo) Omit(cols ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o openAPIDo) Join(table schema.Tabler, on ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o openAPIDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o openAPIDo) RightJoin(table schema.Tabler, on ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o openAPIDo) Group(cols ...field.Expr) IOpenAPIDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o openAPIDo) Having(conds ...gen.Condition) IOpenAPIDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o openAPIDo) Limit(limit int) IOpenAPIDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o openAPIDo) Offset(offset int) IOpenAPIDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o openAPIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOpenAPIDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o openAPIDo) Unscoped() IOpenAPIDo {
	return o.withDO(o.DO.Unscoped())
}

func (o openAPIDo) Create(values ...*model.OpenAPI) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o openAPIDo) CreateInBatches(values []*model.OpenAPI, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o openAPIDo) Save(values ...*model.OpenAPI) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o openAPIDo) First() (*model.OpenAPI, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenAPI), nil
	}
}

func (o openAPIDo) Take() (*model.OpenAPI, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenAPI), nil
	}
}

func (o openAPIDo) Last() (*model.OpenAPI, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenAPI), nil
	}
}

func (o openAPIDo) Find() ([]*model.OpenAPI, error) {
	result, err := o.DO.Find()
	return result.([]*model.OpenAPI), err
}

func (o openAPIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OpenAPI, err error) {
	buf := make([]*model.OpenAPI, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o openAPIDo) FindInBatches(result *[]*model.OpenAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o openAPIDo) Attrs(attrs ...field.AssignExpr) IOpenAPIDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o openAPIDo) Assign(attrs ...field.AssignExpr) IOpenAPIDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o openAPIDo) Joins(fields ...field.RelationField) IOpenAPIDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o openAPIDo) Preload(fields ...field.RelationField) IOpenAPIDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o openAPIDo) FirstOrInit() (*model.OpenAPI, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenAPI), nil
	}
}

func (o openAPIDo) FirstOrCreate() (*model.OpenAPI, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpenAPI), nil
	}
}

func (o openAPIDo) FindByPage(offset int, limit int) (result []*model.OpenAPI, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o openAPIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o openAPIDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o openAPIDo) Delete(models ...*model.OpenAPI) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *openAPIDo) withDO(do gen.Dao) *openAPIDo {
	o.DO = *do.(*gen.DO)
	return o
}
