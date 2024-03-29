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

func newFrErrrecord(db *gorm.DB, opts ...gen.DOOption) frErrrecord {
	_frErrrecord := frErrrecord{}

	_frErrrecord.frErrrecordDo.UseDB(db, opts...)
	_frErrrecord.frErrrecordDo.UseModel(&model.FrErrrecord{})

	tableName := _frErrrecord.frErrrecordDo.TableName()
	_frErrrecord.ALL = field.NewAsterisk(tableName)
	_frErrrecord.ID = field.NewInt32(tableName, "id")
	_frErrrecord.Tname = field.NewString(tableName, "tname")
	_frErrrecord.Sinfo = field.NewString(tableName, "sinfo")
	_frErrrecord.Type = field.NewInt32(tableName, "type")
	_frErrrecord.Msg = field.NewString(tableName, "msg")
	_frErrrecord.Logtime = field.NewTime(tableName, "logtime")
	_frErrrecord.Project = field.NewString(tableName, "project")

	_frErrrecord.fillFieldMap()

	return _frErrrecord
}

type frErrrecord struct {
	frErrrecordDo

	ALL     field.Asterisk
	ID      field.Int32
	Tname   field.String
	Sinfo   field.String
	Type    field.Int32
	Msg     field.String
	Logtime field.Time
	Project field.String

	fieldMap map[string]field.Expr
}

func (f frErrrecord) Table(newTableName string) *frErrrecord {
	f.frErrrecordDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f frErrrecord) As(alias string) *frErrrecord {
	f.frErrrecordDo.DO = *(f.frErrrecordDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *frErrrecord) updateTableName(table string) *frErrrecord {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.Tname = field.NewString(table, "tname")
	f.Sinfo = field.NewString(table, "sinfo")
	f.Type = field.NewInt32(table, "type")
	f.Msg = field.NewString(table, "msg")
	f.Logtime = field.NewTime(table, "logtime")
	f.Project = field.NewString(table, "project")

	f.fillFieldMap()

	return f
}

func (f *frErrrecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *frErrrecord) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 7)
	f.fieldMap["id"] = f.ID
	f.fieldMap["tname"] = f.Tname
	f.fieldMap["sinfo"] = f.Sinfo
	f.fieldMap["type"] = f.Type
	f.fieldMap["msg"] = f.Msg
	f.fieldMap["logtime"] = f.Logtime
	f.fieldMap["project"] = f.Project
}

func (f frErrrecord) clone(db *gorm.DB) frErrrecord {
	f.frErrrecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f frErrrecord) replaceDB(db *gorm.DB) frErrrecord {
	f.frErrrecordDo.ReplaceDB(db)
	return f
}

type frErrrecordDo struct{ gen.DO }

type IFrErrrecordDo interface {
	gen.SubQuery
	Debug() IFrErrrecordDo
	WithContext(ctx context.Context) IFrErrrecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFrErrrecordDo
	WriteDB() IFrErrrecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFrErrrecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFrErrrecordDo
	Not(conds ...gen.Condition) IFrErrrecordDo
	Or(conds ...gen.Condition) IFrErrrecordDo
	Select(conds ...field.Expr) IFrErrrecordDo
	Where(conds ...gen.Condition) IFrErrrecordDo
	Order(conds ...field.Expr) IFrErrrecordDo
	Distinct(cols ...field.Expr) IFrErrrecordDo
	Omit(cols ...field.Expr) IFrErrrecordDo
	Join(table schema.Tabler, on ...field.Expr) IFrErrrecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFrErrrecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFrErrrecordDo
	Group(cols ...field.Expr) IFrErrrecordDo
	Having(conds ...gen.Condition) IFrErrrecordDo
	Limit(limit int) IFrErrrecordDo
	Offset(offset int) IFrErrrecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFrErrrecordDo
	Unscoped() IFrErrrecordDo
	Create(values ...*model.FrErrrecord) error
	CreateInBatches(values []*model.FrErrrecord, batchSize int) error
	Save(values ...*model.FrErrrecord) error
	First() (*model.FrErrrecord, error)
	Take() (*model.FrErrrecord, error)
	Last() (*model.FrErrrecord, error)
	Find() ([]*model.FrErrrecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FrErrrecord, err error)
	FindInBatches(result *[]*model.FrErrrecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FrErrrecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFrErrrecordDo
	Assign(attrs ...field.AssignExpr) IFrErrrecordDo
	Joins(fields ...field.RelationField) IFrErrrecordDo
	Preload(fields ...field.RelationField) IFrErrrecordDo
	FirstOrInit() (*model.FrErrrecord, error)
	FirstOrCreate() (*model.FrErrrecord, error)
	FindByPage(offset int, limit int) (result []*model.FrErrrecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFrErrrecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f frErrrecordDo) Debug() IFrErrrecordDo {
	return f.withDO(f.DO.Debug())
}

func (f frErrrecordDo) WithContext(ctx context.Context) IFrErrrecordDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f frErrrecordDo) ReadDB() IFrErrrecordDo {
	return f.Clauses(dbresolver.Read)
}

func (f frErrrecordDo) WriteDB() IFrErrrecordDo {
	return f.Clauses(dbresolver.Write)
}

func (f frErrrecordDo) Session(config *gorm.Session) IFrErrrecordDo {
	return f.withDO(f.DO.Session(config))
}

func (f frErrrecordDo) Clauses(conds ...clause.Expression) IFrErrrecordDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f frErrrecordDo) Returning(value interface{}, columns ...string) IFrErrrecordDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f frErrrecordDo) Not(conds ...gen.Condition) IFrErrrecordDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f frErrrecordDo) Or(conds ...gen.Condition) IFrErrrecordDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f frErrrecordDo) Select(conds ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f frErrrecordDo) Where(conds ...gen.Condition) IFrErrrecordDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f frErrrecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFrErrrecordDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f frErrrecordDo) Order(conds ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f frErrrecordDo) Distinct(cols ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f frErrrecordDo) Omit(cols ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f frErrrecordDo) Join(table schema.Tabler, on ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f frErrrecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f frErrrecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f frErrrecordDo) Group(cols ...field.Expr) IFrErrrecordDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f frErrrecordDo) Having(conds ...gen.Condition) IFrErrrecordDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f frErrrecordDo) Limit(limit int) IFrErrrecordDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f frErrrecordDo) Offset(offset int) IFrErrrecordDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f frErrrecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFrErrrecordDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f frErrrecordDo) Unscoped() IFrErrrecordDo {
	return f.withDO(f.DO.Unscoped())
}

func (f frErrrecordDo) Create(values ...*model.FrErrrecord) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f frErrrecordDo) CreateInBatches(values []*model.FrErrrecord, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f frErrrecordDo) Save(values ...*model.FrErrrecord) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f frErrrecordDo) First() (*model.FrErrrecord, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrErrrecord), nil
	}
}

func (f frErrrecordDo) Take() (*model.FrErrrecord, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrErrrecord), nil
	}
}

func (f frErrrecordDo) Last() (*model.FrErrrecord, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrErrrecord), nil
	}
}

func (f frErrrecordDo) Find() ([]*model.FrErrrecord, error) {
	result, err := f.DO.Find()
	return result.([]*model.FrErrrecord), err
}

func (f frErrrecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FrErrrecord, err error) {
	buf := make([]*model.FrErrrecord, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f frErrrecordDo) FindInBatches(result *[]*model.FrErrrecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f frErrrecordDo) Attrs(attrs ...field.AssignExpr) IFrErrrecordDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f frErrrecordDo) Assign(attrs ...field.AssignExpr) IFrErrrecordDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f frErrrecordDo) Joins(fields ...field.RelationField) IFrErrrecordDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f frErrrecordDo) Preload(fields ...field.RelationField) IFrErrrecordDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f frErrrecordDo) FirstOrInit() (*model.FrErrrecord, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrErrrecord), nil
	}
}

func (f frErrrecordDo) FirstOrCreate() (*model.FrErrrecord, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrErrrecord), nil
	}
}

func (f frErrrecordDo) FindByPage(offset int, limit int) (result []*model.FrErrrecord, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f frErrrecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f frErrrecordDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f frErrrecordDo) Delete(models ...*model.FrErrrecord) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *frErrrecordDo) withDO(do gen.Dao) *frErrrecordDo {
	f.DO = *do.(*gen.DO)
	return f
}
