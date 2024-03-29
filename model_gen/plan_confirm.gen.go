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

func newPlanConfirm(db *gorm.DB, opts ...gen.DOOption) planConfirm {
	_planConfirm := planConfirm{}

	_planConfirm.planConfirmDo.UseDB(db, opts...)
	_planConfirm.planConfirmDo.UseModel(&model.PlanConfirm{})

	tableName := _planConfirm.planConfirmDo.TableName()
	_planConfirm.ALL = field.NewAsterisk(tableName)
	_planConfirm.ReqConfirmID = field.NewString(tableName, "req_confirm_id")
	_planConfirm.ServerEnvID = field.NewString(tableName, "server_env_id")
	_planConfirm.ProjectType = field.NewString(tableName, "project_type")
	_planConfirm.CreateTime = field.NewString(tableName, "create_time")
	_planConfirm.OpUser = field.NewString(tableName, "op_user")
	_planConfirm.ReqID = field.NewString(tableName, "req_id")

	_planConfirm.fillFieldMap()

	return _planConfirm
}

type planConfirm struct {
	planConfirmDo

	ALL          field.Asterisk
	ReqConfirmID field.String
	ServerEnvID  field.String
	ProjectType  field.String
	CreateTime   field.String
	OpUser       field.String
	ReqID        field.String

	fieldMap map[string]field.Expr
}

func (p planConfirm) Table(newTableName string) *planConfirm {
	p.planConfirmDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p planConfirm) As(alias string) *planConfirm {
	p.planConfirmDo.DO = *(p.planConfirmDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *planConfirm) updateTableName(table string) *planConfirm {
	p.ALL = field.NewAsterisk(table)
	p.ReqConfirmID = field.NewString(table, "req_confirm_id")
	p.ServerEnvID = field.NewString(table, "server_env_id")
	p.ProjectType = field.NewString(table, "project_type")
	p.CreateTime = field.NewString(table, "create_time")
	p.OpUser = field.NewString(table, "op_user")
	p.ReqID = field.NewString(table, "req_id")

	p.fillFieldMap()

	return p
}

func (p *planConfirm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *planConfirm) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["req_confirm_id"] = p.ReqConfirmID
	p.fieldMap["server_env_id"] = p.ServerEnvID
	p.fieldMap["project_type"] = p.ProjectType
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["op_user"] = p.OpUser
	p.fieldMap["req_id"] = p.ReqID
}

func (p planConfirm) clone(db *gorm.DB) planConfirm {
	p.planConfirmDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p planConfirm) replaceDB(db *gorm.DB) planConfirm {
	p.planConfirmDo.ReplaceDB(db)
	return p
}

type planConfirmDo struct{ gen.DO }

type IPlanConfirmDo interface {
	gen.SubQuery
	Debug() IPlanConfirmDo
	WithContext(ctx context.Context) IPlanConfirmDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPlanConfirmDo
	WriteDB() IPlanConfirmDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPlanConfirmDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPlanConfirmDo
	Not(conds ...gen.Condition) IPlanConfirmDo
	Or(conds ...gen.Condition) IPlanConfirmDo
	Select(conds ...field.Expr) IPlanConfirmDo
	Where(conds ...gen.Condition) IPlanConfirmDo
	Order(conds ...field.Expr) IPlanConfirmDo
	Distinct(cols ...field.Expr) IPlanConfirmDo
	Omit(cols ...field.Expr) IPlanConfirmDo
	Join(table schema.Tabler, on ...field.Expr) IPlanConfirmDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPlanConfirmDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPlanConfirmDo
	Group(cols ...field.Expr) IPlanConfirmDo
	Having(conds ...gen.Condition) IPlanConfirmDo
	Limit(limit int) IPlanConfirmDo
	Offset(offset int) IPlanConfirmDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPlanConfirmDo
	Unscoped() IPlanConfirmDo
	Create(values ...*model.PlanConfirm) error
	CreateInBatches(values []*model.PlanConfirm, batchSize int) error
	Save(values ...*model.PlanConfirm) error
	First() (*model.PlanConfirm, error)
	Take() (*model.PlanConfirm, error)
	Last() (*model.PlanConfirm, error)
	Find() ([]*model.PlanConfirm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PlanConfirm, err error)
	FindInBatches(result *[]*model.PlanConfirm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PlanConfirm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPlanConfirmDo
	Assign(attrs ...field.AssignExpr) IPlanConfirmDo
	Joins(fields ...field.RelationField) IPlanConfirmDo
	Preload(fields ...field.RelationField) IPlanConfirmDo
	FirstOrInit() (*model.PlanConfirm, error)
	FirstOrCreate() (*model.PlanConfirm, error)
	FindByPage(offset int, limit int) (result []*model.PlanConfirm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPlanConfirmDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p planConfirmDo) Debug() IPlanConfirmDo {
	return p.withDO(p.DO.Debug())
}

func (p planConfirmDo) WithContext(ctx context.Context) IPlanConfirmDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p planConfirmDo) ReadDB() IPlanConfirmDo {
	return p.Clauses(dbresolver.Read)
}

func (p planConfirmDo) WriteDB() IPlanConfirmDo {
	return p.Clauses(dbresolver.Write)
}

func (p planConfirmDo) Session(config *gorm.Session) IPlanConfirmDo {
	return p.withDO(p.DO.Session(config))
}

func (p planConfirmDo) Clauses(conds ...clause.Expression) IPlanConfirmDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p planConfirmDo) Returning(value interface{}, columns ...string) IPlanConfirmDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p planConfirmDo) Not(conds ...gen.Condition) IPlanConfirmDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p planConfirmDo) Or(conds ...gen.Condition) IPlanConfirmDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p planConfirmDo) Select(conds ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p planConfirmDo) Where(conds ...gen.Condition) IPlanConfirmDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p planConfirmDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPlanConfirmDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p planConfirmDo) Order(conds ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p planConfirmDo) Distinct(cols ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p planConfirmDo) Omit(cols ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p planConfirmDo) Join(table schema.Tabler, on ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p planConfirmDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p planConfirmDo) RightJoin(table schema.Tabler, on ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p planConfirmDo) Group(cols ...field.Expr) IPlanConfirmDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p planConfirmDo) Having(conds ...gen.Condition) IPlanConfirmDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p planConfirmDo) Limit(limit int) IPlanConfirmDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p planConfirmDo) Offset(offset int) IPlanConfirmDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p planConfirmDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPlanConfirmDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p planConfirmDo) Unscoped() IPlanConfirmDo {
	return p.withDO(p.DO.Unscoped())
}

func (p planConfirmDo) Create(values ...*model.PlanConfirm) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p planConfirmDo) CreateInBatches(values []*model.PlanConfirm, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p planConfirmDo) Save(values ...*model.PlanConfirm) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p planConfirmDo) First() (*model.PlanConfirm, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlanConfirm), nil
	}
}

func (p planConfirmDo) Take() (*model.PlanConfirm, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlanConfirm), nil
	}
}

func (p planConfirmDo) Last() (*model.PlanConfirm, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlanConfirm), nil
	}
}

func (p planConfirmDo) Find() ([]*model.PlanConfirm, error) {
	result, err := p.DO.Find()
	return result.([]*model.PlanConfirm), err
}

func (p planConfirmDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PlanConfirm, err error) {
	buf := make([]*model.PlanConfirm, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p planConfirmDo) FindInBatches(result *[]*model.PlanConfirm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p planConfirmDo) Attrs(attrs ...field.AssignExpr) IPlanConfirmDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p planConfirmDo) Assign(attrs ...field.AssignExpr) IPlanConfirmDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p planConfirmDo) Joins(fields ...field.RelationField) IPlanConfirmDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p planConfirmDo) Preload(fields ...field.RelationField) IPlanConfirmDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p planConfirmDo) FirstOrInit() (*model.PlanConfirm, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlanConfirm), nil
	}
}

func (p planConfirmDo) FirstOrCreate() (*model.PlanConfirm, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PlanConfirm), nil
	}
}

func (p planConfirmDo) FindByPage(offset int, limit int) (result []*model.PlanConfirm, count int64, err error) {
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

func (p planConfirmDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p planConfirmDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p planConfirmDo) Delete(models ...*model.PlanConfirm) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *planConfirmDo) withDO(do gen.Dao) *planConfirmDo {
	p.DO = *do.(*gen.DO)
	return p
}
