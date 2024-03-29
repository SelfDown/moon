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

func newVersionPlan(db *gorm.DB, opts ...gen.DOOption) versionPlan {
	_versionPlan := versionPlan{}

	_versionPlan.versionPlanDo.UseDB(db, opts...)
	_versionPlan.versionPlanDo.UseModel(&model.VersionPlan{})

	tableName := _versionPlan.versionPlanDo.TableName()
	_versionPlan.ALL = field.NewAsterisk(tableName)
	_versionPlan.VersionPlanID = field.NewString(tableName, "version_plan_id")
	_versionPlan.ReqSummary = field.NewString(tableName, "req_summary")
	_versionPlan.DevUser = field.NewString(tableName, "dev_user")
	_versionPlan.TestUser = field.NewString(tableName, "test_user")
	_versionPlan.DutyTeam = field.NewString(tableName, "duty_team")
	_versionPlan.PublishVersion = field.NewString(tableName, "publish_version")
	_versionPlan.ReqStatu = field.NewString(tableName, "req_statu")
	_versionPlan.ProductDomain = field.NewString(tableName, "product_domain")
	_versionPlan.SysProjectID = field.NewString(tableName, "sys_project_id")
	_versionPlan.CreateTime = field.NewTime(tableName, "create_time")
	_versionPlan.ModifyTime = field.NewTime(tableName, "modify_time")
	_versionPlan.Comments = field.NewString(tableName, "comments")
	_versionPlan.VersionType = field.NewString(tableName, "version_type")
	_versionPlan.IsDelete = field.NewString(tableName, "is_delete")
	_versionPlan.CreateUser = field.NewString(tableName, "create_user")
	_versionPlan.ModifyUser = field.NewString(tableName, "modify_user")

	_versionPlan.fillFieldMap()

	return _versionPlan
}

type versionPlan struct {
	versionPlanDo

	ALL            field.Asterisk
	VersionPlanID  field.String
	ReqSummary     field.String // 概述
	DevUser        field.String
	TestUser       field.String
	DutyTeam       field.String
	PublishVersion field.String
	ReqStatu       field.String
	ProductDomain  field.String // 产品域
	SysProjectID   field.String
	CreateTime     field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime     field.Time   // 记录修改时间（数据库自动写入）
	Comments       field.String // 备注说明
	VersionType    field.String
	IsDelete       field.String // 是否删除
	CreateUser     field.String
	ModifyUser     field.String

	fieldMap map[string]field.Expr
}

func (v versionPlan) Table(newTableName string) *versionPlan {
	v.versionPlanDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v versionPlan) As(alias string) *versionPlan {
	v.versionPlanDo.DO = *(v.versionPlanDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *versionPlan) updateTableName(table string) *versionPlan {
	v.ALL = field.NewAsterisk(table)
	v.VersionPlanID = field.NewString(table, "version_plan_id")
	v.ReqSummary = field.NewString(table, "req_summary")
	v.DevUser = field.NewString(table, "dev_user")
	v.TestUser = field.NewString(table, "test_user")
	v.DutyTeam = field.NewString(table, "duty_team")
	v.PublishVersion = field.NewString(table, "publish_version")
	v.ReqStatu = field.NewString(table, "req_statu")
	v.ProductDomain = field.NewString(table, "product_domain")
	v.SysProjectID = field.NewString(table, "sys_project_id")
	v.CreateTime = field.NewTime(table, "create_time")
	v.ModifyTime = field.NewTime(table, "modify_time")
	v.Comments = field.NewString(table, "comments")
	v.VersionType = field.NewString(table, "version_type")
	v.IsDelete = field.NewString(table, "is_delete")
	v.CreateUser = field.NewString(table, "create_user")
	v.ModifyUser = field.NewString(table, "modify_user")

	v.fillFieldMap()

	return v
}

func (v *versionPlan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *versionPlan) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 16)
	v.fieldMap["version_plan_id"] = v.VersionPlanID
	v.fieldMap["req_summary"] = v.ReqSummary
	v.fieldMap["dev_user"] = v.DevUser
	v.fieldMap["test_user"] = v.TestUser
	v.fieldMap["duty_team"] = v.DutyTeam
	v.fieldMap["publish_version"] = v.PublishVersion
	v.fieldMap["req_statu"] = v.ReqStatu
	v.fieldMap["product_domain"] = v.ProductDomain
	v.fieldMap["sys_project_id"] = v.SysProjectID
	v.fieldMap["create_time"] = v.CreateTime
	v.fieldMap["modify_time"] = v.ModifyTime
	v.fieldMap["comments"] = v.Comments
	v.fieldMap["version_type"] = v.VersionType
	v.fieldMap["is_delete"] = v.IsDelete
	v.fieldMap["create_user"] = v.CreateUser
	v.fieldMap["modify_user"] = v.ModifyUser
}

func (v versionPlan) clone(db *gorm.DB) versionPlan {
	v.versionPlanDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v versionPlan) replaceDB(db *gorm.DB) versionPlan {
	v.versionPlanDo.ReplaceDB(db)
	return v
}

type versionPlanDo struct{ gen.DO }

type IVersionPlanDo interface {
	gen.SubQuery
	Debug() IVersionPlanDo
	WithContext(ctx context.Context) IVersionPlanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVersionPlanDo
	WriteDB() IVersionPlanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVersionPlanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVersionPlanDo
	Not(conds ...gen.Condition) IVersionPlanDo
	Or(conds ...gen.Condition) IVersionPlanDo
	Select(conds ...field.Expr) IVersionPlanDo
	Where(conds ...gen.Condition) IVersionPlanDo
	Order(conds ...field.Expr) IVersionPlanDo
	Distinct(cols ...field.Expr) IVersionPlanDo
	Omit(cols ...field.Expr) IVersionPlanDo
	Join(table schema.Tabler, on ...field.Expr) IVersionPlanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVersionPlanDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVersionPlanDo
	Group(cols ...field.Expr) IVersionPlanDo
	Having(conds ...gen.Condition) IVersionPlanDo
	Limit(limit int) IVersionPlanDo
	Offset(offset int) IVersionPlanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVersionPlanDo
	Unscoped() IVersionPlanDo
	Create(values ...*model.VersionPlan) error
	CreateInBatches(values []*model.VersionPlan, batchSize int) error
	Save(values ...*model.VersionPlan) error
	First() (*model.VersionPlan, error)
	Take() (*model.VersionPlan, error)
	Last() (*model.VersionPlan, error)
	Find() ([]*model.VersionPlan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VersionPlan, err error)
	FindInBatches(result *[]*model.VersionPlan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VersionPlan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVersionPlanDo
	Assign(attrs ...field.AssignExpr) IVersionPlanDo
	Joins(fields ...field.RelationField) IVersionPlanDo
	Preload(fields ...field.RelationField) IVersionPlanDo
	FirstOrInit() (*model.VersionPlan, error)
	FirstOrCreate() (*model.VersionPlan, error)
	FindByPage(offset int, limit int) (result []*model.VersionPlan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVersionPlanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v versionPlanDo) Debug() IVersionPlanDo {
	return v.withDO(v.DO.Debug())
}

func (v versionPlanDo) WithContext(ctx context.Context) IVersionPlanDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v versionPlanDo) ReadDB() IVersionPlanDo {
	return v.Clauses(dbresolver.Read)
}

func (v versionPlanDo) WriteDB() IVersionPlanDo {
	return v.Clauses(dbresolver.Write)
}

func (v versionPlanDo) Session(config *gorm.Session) IVersionPlanDo {
	return v.withDO(v.DO.Session(config))
}

func (v versionPlanDo) Clauses(conds ...clause.Expression) IVersionPlanDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v versionPlanDo) Returning(value interface{}, columns ...string) IVersionPlanDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v versionPlanDo) Not(conds ...gen.Condition) IVersionPlanDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v versionPlanDo) Or(conds ...gen.Condition) IVersionPlanDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v versionPlanDo) Select(conds ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v versionPlanDo) Where(conds ...gen.Condition) IVersionPlanDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v versionPlanDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IVersionPlanDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v versionPlanDo) Order(conds ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v versionPlanDo) Distinct(cols ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v versionPlanDo) Omit(cols ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v versionPlanDo) Join(table schema.Tabler, on ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v versionPlanDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v versionPlanDo) RightJoin(table schema.Tabler, on ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v versionPlanDo) Group(cols ...field.Expr) IVersionPlanDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v versionPlanDo) Having(conds ...gen.Condition) IVersionPlanDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v versionPlanDo) Limit(limit int) IVersionPlanDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v versionPlanDo) Offset(offset int) IVersionPlanDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v versionPlanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVersionPlanDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v versionPlanDo) Unscoped() IVersionPlanDo {
	return v.withDO(v.DO.Unscoped())
}

func (v versionPlanDo) Create(values ...*model.VersionPlan) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v versionPlanDo) CreateInBatches(values []*model.VersionPlan, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v versionPlanDo) Save(values ...*model.VersionPlan) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v versionPlanDo) First() (*model.VersionPlan, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VersionPlan), nil
	}
}

func (v versionPlanDo) Take() (*model.VersionPlan, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VersionPlan), nil
	}
}

func (v versionPlanDo) Last() (*model.VersionPlan, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VersionPlan), nil
	}
}

func (v versionPlanDo) Find() ([]*model.VersionPlan, error) {
	result, err := v.DO.Find()
	return result.([]*model.VersionPlan), err
}

func (v versionPlanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VersionPlan, err error) {
	buf := make([]*model.VersionPlan, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v versionPlanDo) FindInBatches(result *[]*model.VersionPlan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v versionPlanDo) Attrs(attrs ...field.AssignExpr) IVersionPlanDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v versionPlanDo) Assign(attrs ...field.AssignExpr) IVersionPlanDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v versionPlanDo) Joins(fields ...field.RelationField) IVersionPlanDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v versionPlanDo) Preload(fields ...field.RelationField) IVersionPlanDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v versionPlanDo) FirstOrInit() (*model.VersionPlan, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VersionPlan), nil
	}
}

func (v versionPlanDo) FirstOrCreate() (*model.VersionPlan, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VersionPlan), nil
	}
}

func (v versionPlanDo) FindByPage(offset int, limit int) (result []*model.VersionPlan, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v versionPlanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v versionPlanDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v versionPlanDo) Delete(models ...*model.VersionPlan) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *versionPlanDo) withDO(do gen.Dao) *versionPlanDo {
	v.DO = *do.(*gen.DO)
	return v
}
