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

func newJbosslogcategoryOperator(db *gorm.DB, opts ...gen.DOOption) jbosslogcategoryOperator {
	_jbosslogcategoryOperator := jbosslogcategoryOperator{}

	_jbosslogcategoryOperator.jbosslogcategoryOperatorDo.UseDB(db, opts...)
	_jbosslogcategoryOperator.jbosslogcategoryOperatorDo.UseModel(&model.JbosslogcategoryOperator{})

	tableName := _jbosslogcategoryOperator.jbosslogcategoryOperatorDo.TableName()
	_jbosslogcategoryOperator.ALL = field.NewAsterisk(tableName)
	_jbosslogcategoryOperator.EventID = field.NewString(tableName, "event_id")
	_jbosslogcategoryOperator.InstallSoftID = field.NewString(tableName, "install_soft_id")
	_jbosslogcategoryOperator.ServerIP = field.NewString(tableName, "server_ip")
	_jbosslogcategoryOperator.ServerPort = field.NewString(tableName, "server_port")
	_jbosslogcategoryOperator.ExpiryTime = field.NewTime(tableName, "expiry_time")
	_jbosslogcategoryOperator.LastExeRs = field.NewString(tableName, "last_exe_rs")
	_jbosslogcategoryOperator.CategoryName = field.NewString(tableName, "category_name")
	_jbosslogcategoryOperator.CategoryLevel = field.NewString(tableName, "category_level")
	_jbosslogcategoryOperator.Status = field.NewString(tableName, "status")
	_jbosslogcategoryOperator.OpUser = field.NewString(tableName, "op_user")
	_jbosslogcategoryOperator.RemoveTime = field.NewTime(tableName, "remove_time")
	_jbosslogcategoryOperator.AddTime = field.NewTime(tableName, "add_time")
	_jbosslogcategoryOperator.SoftType = field.NewString(tableName, "soft_type")
	_jbosslogcategoryOperator.Comments = field.NewString(tableName, "comments")
	_jbosslogcategoryOperator.Project = field.NewString(tableName, "project")

	_jbosslogcategoryOperator.fillFieldMap()

	return _jbosslogcategoryOperator
}

type jbosslogcategoryOperator struct {
	jbosslogcategoryOperatorDo

	ALL           field.Asterisk
	EventID       field.String
	InstallSoftID field.String
	ServerIP      field.String
	ServerPort    field.String
	ExpiryTime    field.Time   // 过期时间
	LastExeRs     field.String // 最近执行结果日志
	CategoryName  field.String
	CategoryLevel field.String
	Status        field.String // 1、新入  2、已撤销
	OpUser        field.String
	RemoveTime    field.Time
	AddTime       field.Time // 记录创建时间（数据库自动写入）
	SoftType      field.String
	Comments      field.String // 备注说明
	Project       field.String

	fieldMap map[string]field.Expr
}

func (j jbosslogcategoryOperator) Table(newTableName string) *jbosslogcategoryOperator {
	j.jbosslogcategoryOperatorDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jbosslogcategoryOperator) As(alias string) *jbosslogcategoryOperator {
	j.jbosslogcategoryOperatorDo.DO = *(j.jbosslogcategoryOperatorDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jbosslogcategoryOperator) updateTableName(table string) *jbosslogcategoryOperator {
	j.ALL = field.NewAsterisk(table)
	j.EventID = field.NewString(table, "event_id")
	j.InstallSoftID = field.NewString(table, "install_soft_id")
	j.ServerIP = field.NewString(table, "server_ip")
	j.ServerPort = field.NewString(table, "server_port")
	j.ExpiryTime = field.NewTime(table, "expiry_time")
	j.LastExeRs = field.NewString(table, "last_exe_rs")
	j.CategoryName = field.NewString(table, "category_name")
	j.CategoryLevel = field.NewString(table, "category_level")
	j.Status = field.NewString(table, "status")
	j.OpUser = field.NewString(table, "op_user")
	j.RemoveTime = field.NewTime(table, "remove_time")
	j.AddTime = field.NewTime(table, "add_time")
	j.SoftType = field.NewString(table, "soft_type")
	j.Comments = field.NewString(table, "comments")
	j.Project = field.NewString(table, "project")

	j.fillFieldMap()

	return j
}

func (j *jbosslogcategoryOperator) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jbosslogcategoryOperator) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 15)
	j.fieldMap["event_id"] = j.EventID
	j.fieldMap["install_soft_id"] = j.InstallSoftID
	j.fieldMap["server_ip"] = j.ServerIP
	j.fieldMap["server_port"] = j.ServerPort
	j.fieldMap["expiry_time"] = j.ExpiryTime
	j.fieldMap["last_exe_rs"] = j.LastExeRs
	j.fieldMap["category_name"] = j.CategoryName
	j.fieldMap["category_level"] = j.CategoryLevel
	j.fieldMap["status"] = j.Status
	j.fieldMap["op_user"] = j.OpUser
	j.fieldMap["remove_time"] = j.RemoveTime
	j.fieldMap["add_time"] = j.AddTime
	j.fieldMap["soft_type"] = j.SoftType
	j.fieldMap["comments"] = j.Comments
	j.fieldMap["project"] = j.Project
}

func (j jbosslogcategoryOperator) clone(db *gorm.DB) jbosslogcategoryOperator {
	j.jbosslogcategoryOperatorDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jbosslogcategoryOperator) replaceDB(db *gorm.DB) jbosslogcategoryOperator {
	j.jbosslogcategoryOperatorDo.ReplaceDB(db)
	return j
}

type jbosslogcategoryOperatorDo struct{ gen.DO }

type IJbosslogcategoryOperatorDo interface {
	gen.SubQuery
	Debug() IJbosslogcategoryOperatorDo
	WithContext(ctx context.Context) IJbosslogcategoryOperatorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJbosslogcategoryOperatorDo
	WriteDB() IJbosslogcategoryOperatorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJbosslogcategoryOperatorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJbosslogcategoryOperatorDo
	Not(conds ...gen.Condition) IJbosslogcategoryOperatorDo
	Or(conds ...gen.Condition) IJbosslogcategoryOperatorDo
	Select(conds ...field.Expr) IJbosslogcategoryOperatorDo
	Where(conds ...gen.Condition) IJbosslogcategoryOperatorDo
	Order(conds ...field.Expr) IJbosslogcategoryOperatorDo
	Distinct(cols ...field.Expr) IJbosslogcategoryOperatorDo
	Omit(cols ...field.Expr) IJbosslogcategoryOperatorDo
	Join(table schema.Tabler, on ...field.Expr) IJbosslogcategoryOperatorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJbosslogcategoryOperatorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJbosslogcategoryOperatorDo
	Group(cols ...field.Expr) IJbosslogcategoryOperatorDo
	Having(conds ...gen.Condition) IJbosslogcategoryOperatorDo
	Limit(limit int) IJbosslogcategoryOperatorDo
	Offset(offset int) IJbosslogcategoryOperatorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJbosslogcategoryOperatorDo
	Unscoped() IJbosslogcategoryOperatorDo
	Create(values ...*model.JbosslogcategoryOperator) error
	CreateInBatches(values []*model.JbosslogcategoryOperator, batchSize int) error
	Save(values ...*model.JbosslogcategoryOperator) error
	First() (*model.JbosslogcategoryOperator, error)
	Take() (*model.JbosslogcategoryOperator, error)
	Last() (*model.JbosslogcategoryOperator, error)
	Find() ([]*model.JbosslogcategoryOperator, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JbosslogcategoryOperator, err error)
	FindInBatches(result *[]*model.JbosslogcategoryOperator, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.JbosslogcategoryOperator) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJbosslogcategoryOperatorDo
	Assign(attrs ...field.AssignExpr) IJbosslogcategoryOperatorDo
	Joins(fields ...field.RelationField) IJbosslogcategoryOperatorDo
	Preload(fields ...field.RelationField) IJbosslogcategoryOperatorDo
	FirstOrInit() (*model.JbosslogcategoryOperator, error)
	FirstOrCreate() (*model.JbosslogcategoryOperator, error)
	FindByPage(offset int, limit int) (result []*model.JbosslogcategoryOperator, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJbosslogcategoryOperatorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jbosslogcategoryOperatorDo) Debug() IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Debug())
}

func (j jbosslogcategoryOperatorDo) WithContext(ctx context.Context) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jbosslogcategoryOperatorDo) ReadDB() IJbosslogcategoryOperatorDo {
	return j.Clauses(dbresolver.Read)
}

func (j jbosslogcategoryOperatorDo) WriteDB() IJbosslogcategoryOperatorDo {
	return j.Clauses(dbresolver.Write)
}

func (j jbosslogcategoryOperatorDo) Session(config *gorm.Session) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Session(config))
}

func (j jbosslogcategoryOperatorDo) Clauses(conds ...clause.Expression) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jbosslogcategoryOperatorDo) Returning(value interface{}, columns ...string) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jbosslogcategoryOperatorDo) Not(conds ...gen.Condition) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jbosslogcategoryOperatorDo) Or(conds ...gen.Condition) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jbosslogcategoryOperatorDo) Select(conds ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jbosslogcategoryOperatorDo) Where(conds ...gen.Condition) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jbosslogcategoryOperatorDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IJbosslogcategoryOperatorDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j jbosslogcategoryOperatorDo) Order(conds ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jbosslogcategoryOperatorDo) Distinct(cols ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jbosslogcategoryOperatorDo) Omit(cols ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jbosslogcategoryOperatorDo) Join(table schema.Tabler, on ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jbosslogcategoryOperatorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jbosslogcategoryOperatorDo) RightJoin(table schema.Tabler, on ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jbosslogcategoryOperatorDo) Group(cols ...field.Expr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jbosslogcategoryOperatorDo) Having(conds ...gen.Condition) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jbosslogcategoryOperatorDo) Limit(limit int) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jbosslogcategoryOperatorDo) Offset(offset int) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jbosslogcategoryOperatorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jbosslogcategoryOperatorDo) Unscoped() IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jbosslogcategoryOperatorDo) Create(values ...*model.JbosslogcategoryOperator) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jbosslogcategoryOperatorDo) CreateInBatches(values []*model.JbosslogcategoryOperator, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jbosslogcategoryOperatorDo) Save(values ...*model.JbosslogcategoryOperator) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jbosslogcategoryOperatorDo) First() (*model.JbosslogcategoryOperator, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JbosslogcategoryOperator), nil
	}
}

func (j jbosslogcategoryOperatorDo) Take() (*model.JbosslogcategoryOperator, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JbosslogcategoryOperator), nil
	}
}

func (j jbosslogcategoryOperatorDo) Last() (*model.JbosslogcategoryOperator, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JbosslogcategoryOperator), nil
	}
}

func (j jbosslogcategoryOperatorDo) Find() ([]*model.JbosslogcategoryOperator, error) {
	result, err := j.DO.Find()
	return result.([]*model.JbosslogcategoryOperator), err
}

func (j jbosslogcategoryOperatorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JbosslogcategoryOperator, err error) {
	buf := make([]*model.JbosslogcategoryOperator, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jbosslogcategoryOperatorDo) FindInBatches(result *[]*model.JbosslogcategoryOperator, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jbosslogcategoryOperatorDo) Attrs(attrs ...field.AssignExpr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jbosslogcategoryOperatorDo) Assign(attrs ...field.AssignExpr) IJbosslogcategoryOperatorDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jbosslogcategoryOperatorDo) Joins(fields ...field.RelationField) IJbosslogcategoryOperatorDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jbosslogcategoryOperatorDo) Preload(fields ...field.RelationField) IJbosslogcategoryOperatorDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jbosslogcategoryOperatorDo) FirstOrInit() (*model.JbosslogcategoryOperator, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JbosslogcategoryOperator), nil
	}
}

func (j jbosslogcategoryOperatorDo) FirstOrCreate() (*model.JbosslogcategoryOperator, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JbosslogcategoryOperator), nil
	}
}

func (j jbosslogcategoryOperatorDo) FindByPage(offset int, limit int) (result []*model.JbosslogcategoryOperator, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jbosslogcategoryOperatorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jbosslogcategoryOperatorDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jbosslogcategoryOperatorDo) Delete(models ...*model.JbosslogcategoryOperator) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jbosslogcategoryOperatorDo) withDO(do gen.Dao) *jbosslogcategoryOperatorDo {
	j.DO = *do.(*gen.DO)
	return j
}
