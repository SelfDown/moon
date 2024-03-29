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

func newDeployConfDiffDetail(db *gorm.DB, opts ...gen.DOOption) deployConfDiffDetail {
	_deployConfDiffDetail := deployConfDiffDetail{}

	_deployConfDiffDetail.deployConfDiffDetailDo.UseDB(db, opts...)
	_deployConfDiffDetail.deployConfDiffDetailDo.UseModel(&model.DeployConfDiffDetail{})

	tableName := _deployConfDiffDetail.deployConfDiffDetailDo.TableName()
	_deployConfDiffDetail.ALL = field.NewAsterisk(tableName)
	_deployConfDiffDetail.DiffID = field.NewString(tableName, "diff_id")
	_deployConfDiffDetail.RepWarID = field.NewString(tableName, "rep_war_id")
	_deployConfDiffDetail.ParamKey = field.NewString(tableName, "param_key")
	_deployConfDiffDetail.ParamOrgValue = field.NewString(tableName, "param_org_value")
	_deployConfDiffDetail.ParamRepValue = field.NewString(tableName, "param_rep_value")
	_deployConfDiffDetail.NotReplace = field.NewInt32(tableName, "not_replace")
	_deployConfDiffDetail.CreateTime = field.NewTime(tableName, "create_time")
	_deployConfDiffDetail.ModifyTime = field.NewTime(tableName, "modify_time")
	_deployConfDiffDetail.Comments = field.NewString(tableName, "comments")

	_deployConfDiffDetail.fillFieldMap()

	return _deployConfDiffDetail
}

type deployConfDiffDetail struct {
	deployConfDiffDetailDo

	ALL           field.Asterisk
	DiffID        field.String
	RepWarID      field.String // global参数对比临时表
	ParamKey      field.String
	ParamOrgValue field.String // 替换原始值
	ParamRepValue field.String // 替换后的值
	NotReplace    field.Int32  // 是否被替换 1、被替换 0、未被替换 2 新增KEY
	CreateTime    field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime    field.Time   // 记录修改时间（数据库自动写入）
	Comments      field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (d deployConfDiffDetail) Table(newTableName string) *deployConfDiffDetail {
	d.deployConfDiffDetailDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deployConfDiffDetail) As(alias string) *deployConfDiffDetail {
	d.deployConfDiffDetailDo.DO = *(d.deployConfDiffDetailDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deployConfDiffDetail) updateTableName(table string) *deployConfDiffDetail {
	d.ALL = field.NewAsterisk(table)
	d.DiffID = field.NewString(table, "diff_id")
	d.RepWarID = field.NewString(table, "rep_war_id")
	d.ParamKey = field.NewString(table, "param_key")
	d.ParamOrgValue = field.NewString(table, "param_org_value")
	d.ParamRepValue = field.NewString(table, "param_rep_value")
	d.NotReplace = field.NewInt32(table, "not_replace")
	d.CreateTime = field.NewTime(table, "create_time")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")

	d.fillFieldMap()

	return d
}

func (d *deployConfDiffDetail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deployConfDiffDetail) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 9)
	d.fieldMap["diff_id"] = d.DiffID
	d.fieldMap["rep_war_id"] = d.RepWarID
	d.fieldMap["param_key"] = d.ParamKey
	d.fieldMap["param_org_value"] = d.ParamOrgValue
	d.fieldMap["param_rep_value"] = d.ParamRepValue
	d.fieldMap["not_replace"] = d.NotReplace
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
}

func (d deployConfDiffDetail) clone(db *gorm.DB) deployConfDiffDetail {
	d.deployConfDiffDetailDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deployConfDiffDetail) replaceDB(db *gorm.DB) deployConfDiffDetail {
	d.deployConfDiffDetailDo.ReplaceDB(db)
	return d
}

type deployConfDiffDetailDo struct{ gen.DO }

type IDeployConfDiffDetailDo interface {
	gen.SubQuery
	Debug() IDeployConfDiffDetailDo
	WithContext(ctx context.Context) IDeployConfDiffDetailDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeployConfDiffDetailDo
	WriteDB() IDeployConfDiffDetailDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeployConfDiffDetailDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeployConfDiffDetailDo
	Not(conds ...gen.Condition) IDeployConfDiffDetailDo
	Or(conds ...gen.Condition) IDeployConfDiffDetailDo
	Select(conds ...field.Expr) IDeployConfDiffDetailDo
	Where(conds ...gen.Condition) IDeployConfDiffDetailDo
	Order(conds ...field.Expr) IDeployConfDiffDetailDo
	Distinct(cols ...field.Expr) IDeployConfDiffDetailDo
	Omit(cols ...field.Expr) IDeployConfDiffDetailDo
	Join(table schema.Tabler, on ...field.Expr) IDeployConfDiffDetailDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeployConfDiffDetailDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeployConfDiffDetailDo
	Group(cols ...field.Expr) IDeployConfDiffDetailDo
	Having(conds ...gen.Condition) IDeployConfDiffDetailDo
	Limit(limit int) IDeployConfDiffDetailDo
	Offset(offset int) IDeployConfDiffDetailDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployConfDiffDetailDo
	Unscoped() IDeployConfDiffDetailDo
	Create(values ...*model.DeployConfDiffDetail) error
	CreateInBatches(values []*model.DeployConfDiffDetail, batchSize int) error
	Save(values ...*model.DeployConfDiffDetail) error
	First() (*model.DeployConfDiffDetail, error)
	Take() (*model.DeployConfDiffDetail, error)
	Last() (*model.DeployConfDiffDetail, error)
	Find() ([]*model.DeployConfDiffDetail, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployConfDiffDetail, err error)
	FindInBatches(result *[]*model.DeployConfDiffDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeployConfDiffDetail) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeployConfDiffDetailDo
	Assign(attrs ...field.AssignExpr) IDeployConfDiffDetailDo
	Joins(fields ...field.RelationField) IDeployConfDiffDetailDo
	Preload(fields ...field.RelationField) IDeployConfDiffDetailDo
	FirstOrInit() (*model.DeployConfDiffDetail, error)
	FirstOrCreate() (*model.DeployConfDiffDetail, error)
	FindByPage(offset int, limit int) (result []*model.DeployConfDiffDetail, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeployConfDiffDetailDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deployConfDiffDetailDo) Debug() IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Debug())
}

func (d deployConfDiffDetailDo) WithContext(ctx context.Context) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deployConfDiffDetailDo) ReadDB() IDeployConfDiffDetailDo {
	return d.Clauses(dbresolver.Read)
}

func (d deployConfDiffDetailDo) WriteDB() IDeployConfDiffDetailDo {
	return d.Clauses(dbresolver.Write)
}

func (d deployConfDiffDetailDo) Session(config *gorm.Session) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Session(config))
}

func (d deployConfDiffDetailDo) Clauses(conds ...clause.Expression) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deployConfDiffDetailDo) Returning(value interface{}, columns ...string) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deployConfDiffDetailDo) Not(conds ...gen.Condition) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deployConfDiffDetailDo) Or(conds ...gen.Condition) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deployConfDiffDetailDo) Select(conds ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deployConfDiffDetailDo) Where(conds ...gen.Condition) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deployConfDiffDetailDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDeployConfDiffDetailDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d deployConfDiffDetailDo) Order(conds ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deployConfDiffDetailDo) Distinct(cols ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deployConfDiffDetailDo) Omit(cols ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deployConfDiffDetailDo) Join(table schema.Tabler, on ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deployConfDiffDetailDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deployConfDiffDetailDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deployConfDiffDetailDo) Group(cols ...field.Expr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deployConfDiffDetailDo) Having(conds ...gen.Condition) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deployConfDiffDetailDo) Limit(limit int) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deployConfDiffDetailDo) Offset(offset int) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deployConfDiffDetailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deployConfDiffDetailDo) Unscoped() IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deployConfDiffDetailDo) Create(values ...*model.DeployConfDiffDetail) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deployConfDiffDetailDo) CreateInBatches(values []*model.DeployConfDiffDetail, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deployConfDiffDetailDo) Save(values ...*model.DeployConfDiffDetail) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deployConfDiffDetailDo) First() (*model.DeployConfDiffDetail, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployConfDiffDetail), nil
	}
}

func (d deployConfDiffDetailDo) Take() (*model.DeployConfDiffDetail, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployConfDiffDetail), nil
	}
}

func (d deployConfDiffDetailDo) Last() (*model.DeployConfDiffDetail, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployConfDiffDetail), nil
	}
}

func (d deployConfDiffDetailDo) Find() ([]*model.DeployConfDiffDetail, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeployConfDiffDetail), err
}

func (d deployConfDiffDetailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployConfDiffDetail, err error) {
	buf := make([]*model.DeployConfDiffDetail, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deployConfDiffDetailDo) FindInBatches(result *[]*model.DeployConfDiffDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deployConfDiffDetailDo) Attrs(attrs ...field.AssignExpr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deployConfDiffDetailDo) Assign(attrs ...field.AssignExpr) IDeployConfDiffDetailDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deployConfDiffDetailDo) Joins(fields ...field.RelationField) IDeployConfDiffDetailDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deployConfDiffDetailDo) Preload(fields ...field.RelationField) IDeployConfDiffDetailDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deployConfDiffDetailDo) FirstOrInit() (*model.DeployConfDiffDetail, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployConfDiffDetail), nil
	}
}

func (d deployConfDiffDetailDo) FirstOrCreate() (*model.DeployConfDiffDetail, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployConfDiffDetail), nil
	}
}

func (d deployConfDiffDetailDo) FindByPage(offset int, limit int) (result []*model.DeployConfDiffDetail, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deployConfDiffDetailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deployConfDiffDetailDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deployConfDiffDetailDo) Delete(models ...*model.DeployConfDiffDetail) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deployConfDiffDetailDo) withDO(do gen.Dao) *deployConfDiffDetailDo {
	d.DO = *do.(*gen.DO)
	return d
}
