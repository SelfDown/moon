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

func newWeekUserTeam(db *gorm.DB, opts ...gen.DOOption) weekUserTeam {
	_weekUserTeam := weekUserTeam{}

	_weekUserTeam.weekUserTeamDo.UseDB(db, opts...)
	_weekUserTeam.weekUserTeamDo.UseModel(&model.WeekUserTeam{})

	tableName := _weekUserTeam.weekUserTeamDo.TableName()
	_weekUserTeam.ALL = field.NewAsterisk(tableName)
	_weekUserTeam.WeekUserTeamID = field.NewString(tableName, "week_user_team_id")
	_weekUserTeam.BeanReportID = field.NewString(tableName, "bean_report_id")
	_weekUserTeam.Username = field.NewString(tableName, "username")
	_weekUserTeam.Nick = field.NewString(tableName, "nick")
	_weekUserTeam.Team = field.NewString(tableName, "team")

	_weekUserTeam.fillFieldMap()

	return _weekUserTeam
}

type weekUserTeam struct {
	weekUserTeamDo

	ALL            field.Asterisk
	WeekUserTeamID field.String
	BeanReportID   field.String
	Username       field.String
	Nick           field.String
	Team           field.String

	fieldMap map[string]field.Expr
}

func (w weekUserTeam) Table(newTableName string) *weekUserTeam {
	w.weekUserTeamDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w weekUserTeam) As(alias string) *weekUserTeam {
	w.weekUserTeamDo.DO = *(w.weekUserTeamDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *weekUserTeam) updateTableName(table string) *weekUserTeam {
	w.ALL = field.NewAsterisk(table)
	w.WeekUserTeamID = field.NewString(table, "week_user_team_id")
	w.BeanReportID = field.NewString(table, "bean_report_id")
	w.Username = field.NewString(table, "username")
	w.Nick = field.NewString(table, "nick")
	w.Team = field.NewString(table, "team")

	w.fillFieldMap()

	return w
}

func (w *weekUserTeam) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *weekUserTeam) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["week_user_team_id"] = w.WeekUserTeamID
	w.fieldMap["bean_report_id"] = w.BeanReportID
	w.fieldMap["username"] = w.Username
	w.fieldMap["nick"] = w.Nick
	w.fieldMap["team"] = w.Team
}

func (w weekUserTeam) clone(db *gorm.DB) weekUserTeam {
	w.weekUserTeamDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w weekUserTeam) replaceDB(db *gorm.DB) weekUserTeam {
	w.weekUserTeamDo.ReplaceDB(db)
	return w
}

type weekUserTeamDo struct{ gen.DO }

type IWeekUserTeamDo interface {
	gen.SubQuery
	Debug() IWeekUserTeamDo
	WithContext(ctx context.Context) IWeekUserTeamDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWeekUserTeamDo
	WriteDB() IWeekUserTeamDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWeekUserTeamDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWeekUserTeamDo
	Not(conds ...gen.Condition) IWeekUserTeamDo
	Or(conds ...gen.Condition) IWeekUserTeamDo
	Select(conds ...field.Expr) IWeekUserTeamDo
	Where(conds ...gen.Condition) IWeekUserTeamDo
	Order(conds ...field.Expr) IWeekUserTeamDo
	Distinct(cols ...field.Expr) IWeekUserTeamDo
	Omit(cols ...field.Expr) IWeekUserTeamDo
	Join(table schema.Tabler, on ...field.Expr) IWeekUserTeamDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWeekUserTeamDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWeekUserTeamDo
	Group(cols ...field.Expr) IWeekUserTeamDo
	Having(conds ...gen.Condition) IWeekUserTeamDo
	Limit(limit int) IWeekUserTeamDo
	Offset(offset int) IWeekUserTeamDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWeekUserTeamDo
	Unscoped() IWeekUserTeamDo
	Create(values ...*model.WeekUserTeam) error
	CreateInBatches(values []*model.WeekUserTeam, batchSize int) error
	Save(values ...*model.WeekUserTeam) error
	First() (*model.WeekUserTeam, error)
	Take() (*model.WeekUserTeam, error)
	Last() (*model.WeekUserTeam, error)
	Find() ([]*model.WeekUserTeam, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WeekUserTeam, err error)
	FindInBatches(result *[]*model.WeekUserTeam, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WeekUserTeam) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWeekUserTeamDo
	Assign(attrs ...field.AssignExpr) IWeekUserTeamDo
	Joins(fields ...field.RelationField) IWeekUserTeamDo
	Preload(fields ...field.RelationField) IWeekUserTeamDo
	FirstOrInit() (*model.WeekUserTeam, error)
	FirstOrCreate() (*model.WeekUserTeam, error)
	FindByPage(offset int, limit int) (result []*model.WeekUserTeam, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWeekUserTeamDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w weekUserTeamDo) Debug() IWeekUserTeamDo {
	return w.withDO(w.DO.Debug())
}

func (w weekUserTeamDo) WithContext(ctx context.Context) IWeekUserTeamDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w weekUserTeamDo) ReadDB() IWeekUserTeamDo {
	return w.Clauses(dbresolver.Read)
}

func (w weekUserTeamDo) WriteDB() IWeekUserTeamDo {
	return w.Clauses(dbresolver.Write)
}

func (w weekUserTeamDo) Session(config *gorm.Session) IWeekUserTeamDo {
	return w.withDO(w.DO.Session(config))
}

func (w weekUserTeamDo) Clauses(conds ...clause.Expression) IWeekUserTeamDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w weekUserTeamDo) Returning(value interface{}, columns ...string) IWeekUserTeamDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w weekUserTeamDo) Not(conds ...gen.Condition) IWeekUserTeamDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w weekUserTeamDo) Or(conds ...gen.Condition) IWeekUserTeamDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w weekUserTeamDo) Select(conds ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w weekUserTeamDo) Where(conds ...gen.Condition) IWeekUserTeamDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w weekUserTeamDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IWeekUserTeamDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w weekUserTeamDo) Order(conds ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w weekUserTeamDo) Distinct(cols ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w weekUserTeamDo) Omit(cols ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w weekUserTeamDo) Join(table schema.Tabler, on ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w weekUserTeamDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w weekUserTeamDo) RightJoin(table schema.Tabler, on ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w weekUserTeamDo) Group(cols ...field.Expr) IWeekUserTeamDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w weekUserTeamDo) Having(conds ...gen.Condition) IWeekUserTeamDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w weekUserTeamDo) Limit(limit int) IWeekUserTeamDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w weekUserTeamDo) Offset(offset int) IWeekUserTeamDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w weekUserTeamDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWeekUserTeamDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w weekUserTeamDo) Unscoped() IWeekUserTeamDo {
	return w.withDO(w.DO.Unscoped())
}

func (w weekUserTeamDo) Create(values ...*model.WeekUserTeam) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w weekUserTeamDo) CreateInBatches(values []*model.WeekUserTeam, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w weekUserTeamDo) Save(values ...*model.WeekUserTeam) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w weekUserTeamDo) First() (*model.WeekUserTeam, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeekUserTeam), nil
	}
}

func (w weekUserTeamDo) Take() (*model.WeekUserTeam, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeekUserTeam), nil
	}
}

func (w weekUserTeamDo) Last() (*model.WeekUserTeam, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeekUserTeam), nil
	}
}

func (w weekUserTeamDo) Find() ([]*model.WeekUserTeam, error) {
	result, err := w.DO.Find()
	return result.([]*model.WeekUserTeam), err
}

func (w weekUserTeamDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WeekUserTeam, err error) {
	buf := make([]*model.WeekUserTeam, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w weekUserTeamDo) FindInBatches(result *[]*model.WeekUserTeam, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w weekUserTeamDo) Attrs(attrs ...field.AssignExpr) IWeekUserTeamDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w weekUserTeamDo) Assign(attrs ...field.AssignExpr) IWeekUserTeamDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w weekUserTeamDo) Joins(fields ...field.RelationField) IWeekUserTeamDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w weekUserTeamDo) Preload(fields ...field.RelationField) IWeekUserTeamDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w weekUserTeamDo) FirstOrInit() (*model.WeekUserTeam, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeekUserTeam), nil
	}
}

func (w weekUserTeamDo) FirstOrCreate() (*model.WeekUserTeam, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeekUserTeam), nil
	}
}

func (w weekUserTeamDo) FindByPage(offset int, limit int) (result []*model.WeekUserTeam, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w weekUserTeamDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w weekUserTeamDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w weekUserTeamDo) Delete(models ...*model.WeekUserTeam) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *weekUserTeamDo) withDO(do gen.Dao) *weekUserTeamDo {
	w.DO = *do.(*gen.DO)
	return w
}
