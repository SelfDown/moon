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

func newBeanReportTeam(db *gorm.DB, opts ...gen.DOOption) beanReportTeam {
	_beanReportTeam := beanReportTeam{}

	_beanReportTeam.beanReportTeamDo.UseDB(db, opts...)
	_beanReportTeam.beanReportTeamDo.UseModel(&model.BeanReportTeam{})

	tableName := _beanReportTeam.beanReportTeamDo.TableName()
	_beanReportTeam.ALL = field.NewAsterisk(tableName)
	_beanReportTeam.BeanReportTeamID = field.NewString(tableName, "bean_report_team_id")
	_beanReportTeam.Name = field.NewString(tableName, "name")
	_beanReportTeam.JiraKey = field.NewString(tableName, "jira_key")
	_beanReportTeam.OrderIndex = field.NewInt32(tableName, "order_index")
	_beanReportTeam.ReportShow = field.NewString(tableName, "report_show")

	_beanReportTeam.fillFieldMap()

	return _beanReportTeam
}

type beanReportTeam struct {
	beanReportTeamDo

	ALL              field.Asterisk
	BeanReportTeamID field.String
	Name             field.String // 团队名称
	JiraKey          field.String
	OrderIndex       field.Int32
	ReportShow       field.String

	fieldMap map[string]field.Expr
}

func (b beanReportTeam) Table(newTableName string) *beanReportTeam {
	b.beanReportTeamDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b beanReportTeam) As(alias string) *beanReportTeam {
	b.beanReportTeamDo.DO = *(b.beanReportTeamDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *beanReportTeam) updateTableName(table string) *beanReportTeam {
	b.ALL = field.NewAsterisk(table)
	b.BeanReportTeamID = field.NewString(table, "bean_report_team_id")
	b.Name = field.NewString(table, "name")
	b.JiraKey = field.NewString(table, "jira_key")
	b.OrderIndex = field.NewInt32(table, "order_index")
	b.ReportShow = field.NewString(table, "report_show")

	b.fillFieldMap()

	return b
}

func (b *beanReportTeam) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *beanReportTeam) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 5)
	b.fieldMap["bean_report_team_id"] = b.BeanReportTeamID
	b.fieldMap["name"] = b.Name
	b.fieldMap["jira_key"] = b.JiraKey
	b.fieldMap["order_index"] = b.OrderIndex
	b.fieldMap["report_show"] = b.ReportShow
}

func (b beanReportTeam) clone(db *gorm.DB) beanReportTeam {
	b.beanReportTeamDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b beanReportTeam) replaceDB(db *gorm.DB) beanReportTeam {
	b.beanReportTeamDo.ReplaceDB(db)
	return b
}

type beanReportTeamDo struct{ gen.DO }

type IBeanReportTeamDo interface {
	gen.SubQuery
	Debug() IBeanReportTeamDo
	WithContext(ctx context.Context) IBeanReportTeamDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBeanReportTeamDo
	WriteDB() IBeanReportTeamDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBeanReportTeamDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBeanReportTeamDo
	Not(conds ...gen.Condition) IBeanReportTeamDo
	Or(conds ...gen.Condition) IBeanReportTeamDo
	Select(conds ...field.Expr) IBeanReportTeamDo
	Where(conds ...gen.Condition) IBeanReportTeamDo
	Order(conds ...field.Expr) IBeanReportTeamDo
	Distinct(cols ...field.Expr) IBeanReportTeamDo
	Omit(cols ...field.Expr) IBeanReportTeamDo
	Join(table schema.Tabler, on ...field.Expr) IBeanReportTeamDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBeanReportTeamDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBeanReportTeamDo
	Group(cols ...field.Expr) IBeanReportTeamDo
	Having(conds ...gen.Condition) IBeanReportTeamDo
	Limit(limit int) IBeanReportTeamDo
	Offset(offset int) IBeanReportTeamDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBeanReportTeamDo
	Unscoped() IBeanReportTeamDo
	Create(values ...*model.BeanReportTeam) error
	CreateInBatches(values []*model.BeanReportTeam, batchSize int) error
	Save(values ...*model.BeanReportTeam) error
	First() (*model.BeanReportTeam, error)
	Take() (*model.BeanReportTeam, error)
	Last() (*model.BeanReportTeam, error)
	Find() ([]*model.BeanReportTeam, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BeanReportTeam, err error)
	FindInBatches(result *[]*model.BeanReportTeam, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BeanReportTeam) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBeanReportTeamDo
	Assign(attrs ...field.AssignExpr) IBeanReportTeamDo
	Joins(fields ...field.RelationField) IBeanReportTeamDo
	Preload(fields ...field.RelationField) IBeanReportTeamDo
	FirstOrInit() (*model.BeanReportTeam, error)
	FirstOrCreate() (*model.BeanReportTeam, error)
	FindByPage(offset int, limit int) (result []*model.BeanReportTeam, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBeanReportTeamDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b beanReportTeamDo) Debug() IBeanReportTeamDo {
	return b.withDO(b.DO.Debug())
}

func (b beanReportTeamDo) WithContext(ctx context.Context) IBeanReportTeamDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b beanReportTeamDo) ReadDB() IBeanReportTeamDo {
	return b.Clauses(dbresolver.Read)
}

func (b beanReportTeamDo) WriteDB() IBeanReportTeamDo {
	return b.Clauses(dbresolver.Write)
}

func (b beanReportTeamDo) Session(config *gorm.Session) IBeanReportTeamDo {
	return b.withDO(b.DO.Session(config))
}

func (b beanReportTeamDo) Clauses(conds ...clause.Expression) IBeanReportTeamDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b beanReportTeamDo) Returning(value interface{}, columns ...string) IBeanReportTeamDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b beanReportTeamDo) Not(conds ...gen.Condition) IBeanReportTeamDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b beanReportTeamDo) Or(conds ...gen.Condition) IBeanReportTeamDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b beanReportTeamDo) Select(conds ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b beanReportTeamDo) Where(conds ...gen.Condition) IBeanReportTeamDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b beanReportTeamDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBeanReportTeamDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b beanReportTeamDo) Order(conds ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b beanReportTeamDo) Distinct(cols ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b beanReportTeamDo) Omit(cols ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b beanReportTeamDo) Join(table schema.Tabler, on ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b beanReportTeamDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b beanReportTeamDo) RightJoin(table schema.Tabler, on ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b beanReportTeamDo) Group(cols ...field.Expr) IBeanReportTeamDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b beanReportTeamDo) Having(conds ...gen.Condition) IBeanReportTeamDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b beanReportTeamDo) Limit(limit int) IBeanReportTeamDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b beanReportTeamDo) Offset(offset int) IBeanReportTeamDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b beanReportTeamDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBeanReportTeamDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b beanReportTeamDo) Unscoped() IBeanReportTeamDo {
	return b.withDO(b.DO.Unscoped())
}

func (b beanReportTeamDo) Create(values ...*model.BeanReportTeam) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b beanReportTeamDo) CreateInBatches(values []*model.BeanReportTeam, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b beanReportTeamDo) Save(values ...*model.BeanReportTeam) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b beanReportTeamDo) First() (*model.BeanReportTeam, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportTeam), nil
	}
}

func (b beanReportTeamDo) Take() (*model.BeanReportTeam, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportTeam), nil
	}
}

func (b beanReportTeamDo) Last() (*model.BeanReportTeam, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportTeam), nil
	}
}

func (b beanReportTeamDo) Find() ([]*model.BeanReportTeam, error) {
	result, err := b.DO.Find()
	return result.([]*model.BeanReportTeam), err
}

func (b beanReportTeamDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BeanReportTeam, err error) {
	buf := make([]*model.BeanReportTeam, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b beanReportTeamDo) FindInBatches(result *[]*model.BeanReportTeam, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b beanReportTeamDo) Attrs(attrs ...field.AssignExpr) IBeanReportTeamDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b beanReportTeamDo) Assign(attrs ...field.AssignExpr) IBeanReportTeamDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b beanReportTeamDo) Joins(fields ...field.RelationField) IBeanReportTeamDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b beanReportTeamDo) Preload(fields ...field.RelationField) IBeanReportTeamDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b beanReportTeamDo) FirstOrInit() (*model.BeanReportTeam, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportTeam), nil
	}
}

func (b beanReportTeamDo) FirstOrCreate() (*model.BeanReportTeam, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BeanReportTeam), nil
	}
}

func (b beanReportTeamDo) FindByPage(offset int, limit int) (result []*model.BeanReportTeam, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b beanReportTeamDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b beanReportTeamDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b beanReportTeamDo) Delete(models ...*model.BeanReportTeam) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *beanReportTeamDo) withDO(do gen.Dao) *beanReportTeamDo {
	b.DO = *do.(*gen.DO)
	return b
}
