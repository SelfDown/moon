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

func newWebshellWatch(db *gorm.DB, opts ...gen.DOOption) webshellWatch {
	_webshellWatch := webshellWatch{}

	_webshellWatch.webshellWatchDo.UseDB(db, opts...)
	_webshellWatch.webshellWatchDo.UseModel(&model.WebshellWatch{})

	tableName := _webshellWatch.webshellWatchDo.TableName()
	_webshellWatch.ALL = field.NewAsterisk(tableName)
	_webshellWatch.WatchID = field.NewString(tableName, "watch_id")
	_webshellWatch.Userid = field.NewString(tableName, "userid")
	_webshellWatch.ServerID = field.NewString(tableName, "server_id")

	_webshellWatch.fillFieldMap()

	return _webshellWatch
}

type webshellWatch struct {
	webshellWatchDo

	ALL      field.Asterisk
	WatchID  field.String
	Userid   field.String
	ServerID field.String

	fieldMap map[string]field.Expr
}

func (w webshellWatch) Table(newTableName string) *webshellWatch {
	w.webshellWatchDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w webshellWatch) As(alias string) *webshellWatch {
	w.webshellWatchDo.DO = *(w.webshellWatchDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *webshellWatch) updateTableName(table string) *webshellWatch {
	w.ALL = field.NewAsterisk(table)
	w.WatchID = field.NewString(table, "watch_id")
	w.Userid = field.NewString(table, "userid")
	w.ServerID = field.NewString(table, "server_id")

	w.fillFieldMap()

	return w
}

func (w *webshellWatch) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *webshellWatch) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 3)
	w.fieldMap["watch_id"] = w.WatchID
	w.fieldMap["userid"] = w.Userid
	w.fieldMap["server_id"] = w.ServerID
}

func (w webshellWatch) clone(db *gorm.DB) webshellWatch {
	w.webshellWatchDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w webshellWatch) replaceDB(db *gorm.DB) webshellWatch {
	w.webshellWatchDo.ReplaceDB(db)
	return w
}

type webshellWatchDo struct{ gen.DO }

type IWebshellWatchDo interface {
	gen.SubQuery
	Debug() IWebshellWatchDo
	WithContext(ctx context.Context) IWebshellWatchDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWebshellWatchDo
	WriteDB() IWebshellWatchDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWebshellWatchDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWebshellWatchDo
	Not(conds ...gen.Condition) IWebshellWatchDo
	Or(conds ...gen.Condition) IWebshellWatchDo
	Select(conds ...field.Expr) IWebshellWatchDo
	Where(conds ...gen.Condition) IWebshellWatchDo
	Order(conds ...field.Expr) IWebshellWatchDo
	Distinct(cols ...field.Expr) IWebshellWatchDo
	Omit(cols ...field.Expr) IWebshellWatchDo
	Join(table schema.Tabler, on ...field.Expr) IWebshellWatchDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWebshellWatchDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWebshellWatchDo
	Group(cols ...field.Expr) IWebshellWatchDo
	Having(conds ...gen.Condition) IWebshellWatchDo
	Limit(limit int) IWebshellWatchDo
	Offset(offset int) IWebshellWatchDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWebshellWatchDo
	Unscoped() IWebshellWatchDo
	Create(values ...*model.WebshellWatch) error
	CreateInBatches(values []*model.WebshellWatch, batchSize int) error
	Save(values ...*model.WebshellWatch) error
	First() (*model.WebshellWatch, error)
	Take() (*model.WebshellWatch, error)
	Last() (*model.WebshellWatch, error)
	Find() ([]*model.WebshellWatch, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WebshellWatch, err error)
	FindInBatches(result *[]*model.WebshellWatch, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WebshellWatch) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWebshellWatchDo
	Assign(attrs ...field.AssignExpr) IWebshellWatchDo
	Joins(fields ...field.RelationField) IWebshellWatchDo
	Preload(fields ...field.RelationField) IWebshellWatchDo
	FirstOrInit() (*model.WebshellWatch, error)
	FirstOrCreate() (*model.WebshellWatch, error)
	FindByPage(offset int, limit int) (result []*model.WebshellWatch, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWebshellWatchDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w webshellWatchDo) Debug() IWebshellWatchDo {
	return w.withDO(w.DO.Debug())
}

func (w webshellWatchDo) WithContext(ctx context.Context) IWebshellWatchDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w webshellWatchDo) ReadDB() IWebshellWatchDo {
	return w.Clauses(dbresolver.Read)
}

func (w webshellWatchDo) WriteDB() IWebshellWatchDo {
	return w.Clauses(dbresolver.Write)
}

func (w webshellWatchDo) Session(config *gorm.Session) IWebshellWatchDo {
	return w.withDO(w.DO.Session(config))
}

func (w webshellWatchDo) Clauses(conds ...clause.Expression) IWebshellWatchDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w webshellWatchDo) Returning(value interface{}, columns ...string) IWebshellWatchDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w webshellWatchDo) Not(conds ...gen.Condition) IWebshellWatchDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w webshellWatchDo) Or(conds ...gen.Condition) IWebshellWatchDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w webshellWatchDo) Select(conds ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w webshellWatchDo) Where(conds ...gen.Condition) IWebshellWatchDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w webshellWatchDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IWebshellWatchDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w webshellWatchDo) Order(conds ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w webshellWatchDo) Distinct(cols ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w webshellWatchDo) Omit(cols ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w webshellWatchDo) Join(table schema.Tabler, on ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w webshellWatchDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w webshellWatchDo) RightJoin(table schema.Tabler, on ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w webshellWatchDo) Group(cols ...field.Expr) IWebshellWatchDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w webshellWatchDo) Having(conds ...gen.Condition) IWebshellWatchDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w webshellWatchDo) Limit(limit int) IWebshellWatchDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w webshellWatchDo) Offset(offset int) IWebshellWatchDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w webshellWatchDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWebshellWatchDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w webshellWatchDo) Unscoped() IWebshellWatchDo {
	return w.withDO(w.DO.Unscoped())
}

func (w webshellWatchDo) Create(values ...*model.WebshellWatch) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w webshellWatchDo) CreateInBatches(values []*model.WebshellWatch, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w webshellWatchDo) Save(values ...*model.WebshellWatch) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w webshellWatchDo) First() (*model.WebshellWatch, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WebshellWatch), nil
	}
}

func (w webshellWatchDo) Take() (*model.WebshellWatch, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WebshellWatch), nil
	}
}

func (w webshellWatchDo) Last() (*model.WebshellWatch, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WebshellWatch), nil
	}
}

func (w webshellWatchDo) Find() ([]*model.WebshellWatch, error) {
	result, err := w.DO.Find()
	return result.([]*model.WebshellWatch), err
}

func (w webshellWatchDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WebshellWatch, err error) {
	buf := make([]*model.WebshellWatch, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w webshellWatchDo) FindInBatches(result *[]*model.WebshellWatch, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w webshellWatchDo) Attrs(attrs ...field.AssignExpr) IWebshellWatchDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w webshellWatchDo) Assign(attrs ...field.AssignExpr) IWebshellWatchDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w webshellWatchDo) Joins(fields ...field.RelationField) IWebshellWatchDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w webshellWatchDo) Preload(fields ...field.RelationField) IWebshellWatchDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w webshellWatchDo) FirstOrInit() (*model.WebshellWatch, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WebshellWatch), nil
	}
}

func (w webshellWatchDo) FirstOrCreate() (*model.WebshellWatch, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WebshellWatch), nil
	}
}

func (w webshellWatchDo) FindByPage(offset int, limit int) (result []*model.WebshellWatch, count int64, err error) {
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

func (w webshellWatchDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w webshellWatchDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w webshellWatchDo) Delete(models ...*model.WebshellWatch) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *webshellWatchDo) withDO(do gen.Dao) *webshellWatchDo {
	w.DO = *do.(*gen.DO)
	return w
}
