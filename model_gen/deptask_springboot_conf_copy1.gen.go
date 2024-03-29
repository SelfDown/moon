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

func newDeptaskSpringbootConfCopy1(db *gorm.DB, opts ...gen.DOOption) deptaskSpringbootConfCopy1 {
	_deptaskSpringbootConfCopy1 := deptaskSpringbootConfCopy1{}

	_deptaskSpringbootConfCopy1.deptaskSpringbootConfCopy1Do.UseDB(db, opts...)
	_deptaskSpringbootConfCopy1.deptaskSpringbootConfCopy1Do.UseModel(&model.DeptaskSpringbootConfCopy1{})

	tableName := _deptaskSpringbootConfCopy1.deptaskSpringbootConfCopy1Do.TableName()
	_deptaskSpringbootConfCopy1.ALL = field.NewAsterisk(tableName)
	_deptaskSpringbootConfCopy1.SpringbootConfID = field.NewString(tableName, "springboot_conf_id")
	_deptaskSpringbootConfCopy1.DepTaskID = field.NewString(tableName, "dep_task_id")
	_deptaskSpringbootConfCopy1.EnvConf = field.NewString(tableName, "env_conf")
	_deptaskSpringbootConfCopy1.Isvalid = field.NewString(tableName, "isvalid")
	_deptaskSpringbootConfCopy1.OpUser = field.NewString(tableName, "op_user")
	_deptaskSpringbootConfCopy1.CreateTime = field.NewTime(tableName, "create_time")
	_deptaskSpringbootConfCopy1.ModifyTime = field.NewTime(tableName, "modify_time")
	_deptaskSpringbootConfCopy1.Comments = field.NewString(tableName, "comments")

	_deptaskSpringbootConfCopy1.fillFieldMap()

	return _deptaskSpringbootConfCopy1
}

type deptaskSpringbootConfCopy1 struct {
	deptaskSpringbootConfCopy1Do

	ALL              field.Asterisk
	SpringbootConfID field.String
	DepTaskID        field.String
	EnvConf          field.String
	Isvalid          field.String // 0 无效   1有效
	OpUser           field.String
	CreateTime       field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime       field.Time   // 记录修改时间（数据库自动写入）
	Comments         field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (d deptaskSpringbootConfCopy1) Table(newTableName string) *deptaskSpringbootConfCopy1 {
	d.deptaskSpringbootConfCopy1Do.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deptaskSpringbootConfCopy1) As(alias string) *deptaskSpringbootConfCopy1 {
	d.deptaskSpringbootConfCopy1Do.DO = *(d.deptaskSpringbootConfCopy1Do.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deptaskSpringbootConfCopy1) updateTableName(table string) *deptaskSpringbootConfCopy1 {
	d.ALL = field.NewAsterisk(table)
	d.SpringbootConfID = field.NewString(table, "springboot_conf_id")
	d.DepTaskID = field.NewString(table, "dep_task_id")
	d.EnvConf = field.NewString(table, "env_conf")
	d.Isvalid = field.NewString(table, "isvalid")
	d.OpUser = field.NewString(table, "op_user")
	d.CreateTime = field.NewTime(table, "create_time")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")

	d.fillFieldMap()

	return d
}

func (d *deptaskSpringbootConfCopy1) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deptaskSpringbootConfCopy1) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["springboot_conf_id"] = d.SpringbootConfID
	d.fieldMap["dep_task_id"] = d.DepTaskID
	d.fieldMap["env_conf"] = d.EnvConf
	d.fieldMap["isvalid"] = d.Isvalid
	d.fieldMap["op_user"] = d.OpUser
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
}

func (d deptaskSpringbootConfCopy1) clone(db *gorm.DB) deptaskSpringbootConfCopy1 {
	d.deptaskSpringbootConfCopy1Do.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deptaskSpringbootConfCopy1) replaceDB(db *gorm.DB) deptaskSpringbootConfCopy1 {
	d.deptaskSpringbootConfCopy1Do.ReplaceDB(db)
	return d
}

type deptaskSpringbootConfCopy1Do struct{ gen.DO }

type IDeptaskSpringbootConfCopy1Do interface {
	gen.SubQuery
	Debug() IDeptaskSpringbootConfCopy1Do
	WithContext(ctx context.Context) IDeptaskSpringbootConfCopy1Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeptaskSpringbootConfCopy1Do
	WriteDB() IDeptaskSpringbootConfCopy1Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeptaskSpringbootConfCopy1Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeptaskSpringbootConfCopy1Do
	Not(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do
	Or(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do
	Select(conds ...field.Expr) IDeptaskSpringbootConfCopy1Do
	Where(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do
	Order(conds ...field.Expr) IDeptaskSpringbootConfCopy1Do
	Distinct(cols ...field.Expr) IDeptaskSpringbootConfCopy1Do
	Omit(cols ...field.Expr) IDeptaskSpringbootConfCopy1Do
	Join(table schema.Tabler, on ...field.Expr) IDeptaskSpringbootConfCopy1Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeptaskSpringbootConfCopy1Do
	RightJoin(table schema.Tabler, on ...field.Expr) IDeptaskSpringbootConfCopy1Do
	Group(cols ...field.Expr) IDeptaskSpringbootConfCopy1Do
	Having(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do
	Limit(limit int) IDeptaskSpringbootConfCopy1Do
	Offset(offset int) IDeptaskSpringbootConfCopy1Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeptaskSpringbootConfCopy1Do
	Unscoped() IDeptaskSpringbootConfCopy1Do
	Create(values ...*model.DeptaskSpringbootConfCopy1) error
	CreateInBatches(values []*model.DeptaskSpringbootConfCopy1, batchSize int) error
	Save(values ...*model.DeptaskSpringbootConfCopy1) error
	First() (*model.DeptaskSpringbootConfCopy1, error)
	Take() (*model.DeptaskSpringbootConfCopy1, error)
	Last() (*model.DeptaskSpringbootConfCopy1, error)
	Find() ([]*model.DeptaskSpringbootConfCopy1, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeptaskSpringbootConfCopy1, err error)
	FindInBatches(result *[]*model.DeptaskSpringbootConfCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeptaskSpringbootConfCopy1) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeptaskSpringbootConfCopy1Do
	Assign(attrs ...field.AssignExpr) IDeptaskSpringbootConfCopy1Do
	Joins(fields ...field.RelationField) IDeptaskSpringbootConfCopy1Do
	Preload(fields ...field.RelationField) IDeptaskSpringbootConfCopy1Do
	FirstOrInit() (*model.DeptaskSpringbootConfCopy1, error)
	FirstOrCreate() (*model.DeptaskSpringbootConfCopy1, error)
	FindByPage(offset int, limit int) (result []*model.DeptaskSpringbootConfCopy1, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeptaskSpringbootConfCopy1Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deptaskSpringbootConfCopy1Do) Debug() IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Debug())
}

func (d deptaskSpringbootConfCopy1Do) WithContext(ctx context.Context) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deptaskSpringbootConfCopy1Do) ReadDB() IDeptaskSpringbootConfCopy1Do {
	return d.Clauses(dbresolver.Read)
}

func (d deptaskSpringbootConfCopy1Do) WriteDB() IDeptaskSpringbootConfCopy1Do {
	return d.Clauses(dbresolver.Write)
}

func (d deptaskSpringbootConfCopy1Do) Session(config *gorm.Session) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Session(config))
}

func (d deptaskSpringbootConfCopy1Do) Clauses(conds ...clause.Expression) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Returning(value interface{}, columns ...string) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deptaskSpringbootConfCopy1Do) Not(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Not(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Or(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Or(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Select(conds ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Select(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Where(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Where(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDeptaskSpringbootConfCopy1Do {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d deptaskSpringbootConfCopy1Do) Order(conds ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Order(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Distinct(cols ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deptaskSpringbootConfCopy1Do) Omit(cols ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deptaskSpringbootConfCopy1Do) Join(table schema.Tabler, on ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deptaskSpringbootConfCopy1Do) LeftJoin(table schema.Tabler, on ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deptaskSpringbootConfCopy1Do) RightJoin(table schema.Tabler, on ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deptaskSpringbootConfCopy1Do) Group(cols ...field.Expr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Group(cols...))
}

func (d deptaskSpringbootConfCopy1Do) Having(conds ...gen.Condition) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Having(conds...))
}

func (d deptaskSpringbootConfCopy1Do) Limit(limit int) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Limit(limit))
}

func (d deptaskSpringbootConfCopy1Do) Offset(offset int) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Offset(offset))
}

func (d deptaskSpringbootConfCopy1Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deptaskSpringbootConfCopy1Do) Unscoped() IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Unscoped())
}

func (d deptaskSpringbootConfCopy1Do) Create(values ...*model.DeptaskSpringbootConfCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deptaskSpringbootConfCopy1Do) CreateInBatches(values []*model.DeptaskSpringbootConfCopy1, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deptaskSpringbootConfCopy1Do) Save(values ...*model.DeptaskSpringbootConfCopy1) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deptaskSpringbootConfCopy1Do) First() (*model.DeptaskSpringbootConfCopy1, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeptaskSpringbootConfCopy1), nil
	}
}

func (d deptaskSpringbootConfCopy1Do) Take() (*model.DeptaskSpringbootConfCopy1, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeptaskSpringbootConfCopy1), nil
	}
}

func (d deptaskSpringbootConfCopy1Do) Last() (*model.DeptaskSpringbootConfCopy1, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeptaskSpringbootConfCopy1), nil
	}
}

func (d deptaskSpringbootConfCopy1Do) Find() ([]*model.DeptaskSpringbootConfCopy1, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeptaskSpringbootConfCopy1), err
}

func (d deptaskSpringbootConfCopy1Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeptaskSpringbootConfCopy1, err error) {
	buf := make([]*model.DeptaskSpringbootConfCopy1, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deptaskSpringbootConfCopy1Do) FindInBatches(result *[]*model.DeptaskSpringbootConfCopy1, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deptaskSpringbootConfCopy1Do) Attrs(attrs ...field.AssignExpr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deptaskSpringbootConfCopy1Do) Assign(attrs ...field.AssignExpr) IDeptaskSpringbootConfCopy1Do {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deptaskSpringbootConfCopy1Do) Joins(fields ...field.RelationField) IDeptaskSpringbootConfCopy1Do {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deptaskSpringbootConfCopy1Do) Preload(fields ...field.RelationField) IDeptaskSpringbootConfCopy1Do {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deptaskSpringbootConfCopy1Do) FirstOrInit() (*model.DeptaskSpringbootConfCopy1, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeptaskSpringbootConfCopy1), nil
	}
}

func (d deptaskSpringbootConfCopy1Do) FirstOrCreate() (*model.DeptaskSpringbootConfCopy1, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeptaskSpringbootConfCopy1), nil
	}
}

func (d deptaskSpringbootConfCopy1Do) FindByPage(offset int, limit int) (result []*model.DeptaskSpringbootConfCopy1, count int64, err error) {
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

func (d deptaskSpringbootConfCopy1Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deptaskSpringbootConfCopy1Do) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deptaskSpringbootConfCopy1Do) Delete(models ...*model.DeptaskSpringbootConfCopy1) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deptaskSpringbootConfCopy1Do) withDO(do gen.Dao) *deptaskSpringbootConfCopy1Do {
	d.DO = *do.(*gen.DO)
	return d
}
