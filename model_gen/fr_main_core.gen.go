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

func newFrMainCore(db *gorm.DB, opts ...gen.DOOption) frMainCore {
	_frMainCore := frMainCore{}

	_frMainCore.frMainCoreDo.UseDB(db, opts...)
	_frMainCore.frMainCoreDo.UseModel(&model.FrMainCore{})

	tableName := _frMainCore.frMainCoreDo.TableName()
	_frMainCore.ALL = field.NewAsterisk(tableName)
	_frMainCore.ID = field.NewInt32(tableName, "id")
	_frMainCore.CptMaj = field.NewString(tableName, "cpt_maj")
	_frMainCore.CptType = field.NewString(tableName, "cpt_type")
	_frMainCore.InOutFlag = field.NewString(tableName, "in_out_flag")
	_frMainCore.CptName = field.NewString(tableName, "cpt_name")
	_frMainCore.CptNameEng = field.NewString(tableName, "cpt_name_eng")
	_frMainCore.CptExp = field.NewString(tableName, "cpt_exp")
	_frMainCore.CptForm = field.NewString(tableName, "cpt_form")
	_frMainCore.CptFlag = field.NewString(tableName, "cpt_flag")
	_frMainCore.CptReq = field.NewString(tableName, "cpt_req")
	_frMainCore.ReqPerson = field.NewString(tableName, "req_person")
	_frMainCore.ShowFlag = field.NewString(tableName, "show_flag")
	_frMainCore.CptKey = field.NewString(tableName, "cpt_key")
	_frMainCore.MidTable = field.NewString(tableName, "mid_table")
	_frMainCore.Descn = field.NewString(tableName, "descn")
	_frMainCore.Menu = field.NewString(tableName, "menu")
	_frMainCore.CreateTime = field.NewTime(tableName, "create_time")
	_frMainCore.UpdateTime = field.NewTime(tableName, "update_time")
	_frMainCore.ReportPath = field.NewString(tableName, "report_path")
	_frMainCore.Project = field.NewString(tableName, "project")

	_frMainCore.fillFieldMap()

	return _frMainCore
}

type frMainCore struct {
	frMainCoreDo

	ALL        field.Asterisk
	ID         field.Int32
	CptMaj     field.String
	CptType    field.String
	InOutFlag  field.String
	CptName    field.String
	CptNameEng field.String
	CptExp     field.String
	CptForm    field.String
	CptFlag    field.String
	CptReq     field.String
	ReqPerson  field.String
	ShowFlag   field.String
	CptKey     field.String
	MidTable   field.String
	Descn      field.String
	Menu       field.String
	CreateTime field.Time
	UpdateTime field.Time
	ReportPath field.String
	Project    field.String

	fieldMap map[string]field.Expr
}

func (f frMainCore) Table(newTableName string) *frMainCore {
	f.frMainCoreDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f frMainCore) As(alias string) *frMainCore {
	f.frMainCoreDo.DO = *(f.frMainCoreDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *frMainCore) updateTableName(table string) *frMainCore {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.CptMaj = field.NewString(table, "cpt_maj")
	f.CptType = field.NewString(table, "cpt_type")
	f.InOutFlag = field.NewString(table, "in_out_flag")
	f.CptName = field.NewString(table, "cpt_name")
	f.CptNameEng = field.NewString(table, "cpt_name_eng")
	f.CptExp = field.NewString(table, "cpt_exp")
	f.CptForm = field.NewString(table, "cpt_form")
	f.CptFlag = field.NewString(table, "cpt_flag")
	f.CptReq = field.NewString(table, "cpt_req")
	f.ReqPerson = field.NewString(table, "req_person")
	f.ShowFlag = field.NewString(table, "show_flag")
	f.CptKey = field.NewString(table, "cpt_key")
	f.MidTable = field.NewString(table, "mid_table")
	f.Descn = field.NewString(table, "descn")
	f.Menu = field.NewString(table, "menu")
	f.CreateTime = field.NewTime(table, "create_time")
	f.UpdateTime = field.NewTime(table, "update_time")
	f.ReportPath = field.NewString(table, "report_path")
	f.Project = field.NewString(table, "project")

	f.fillFieldMap()

	return f
}

func (f *frMainCore) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *frMainCore) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 20)
	f.fieldMap["id"] = f.ID
	f.fieldMap["cpt_maj"] = f.CptMaj
	f.fieldMap["cpt_type"] = f.CptType
	f.fieldMap["in_out_flag"] = f.InOutFlag
	f.fieldMap["cpt_name"] = f.CptName
	f.fieldMap["cpt_name_eng"] = f.CptNameEng
	f.fieldMap["cpt_exp"] = f.CptExp
	f.fieldMap["cpt_form"] = f.CptForm
	f.fieldMap["cpt_flag"] = f.CptFlag
	f.fieldMap["cpt_req"] = f.CptReq
	f.fieldMap["req_person"] = f.ReqPerson
	f.fieldMap["show_flag"] = f.ShowFlag
	f.fieldMap["cpt_key"] = f.CptKey
	f.fieldMap["mid_table"] = f.MidTable
	f.fieldMap["descn"] = f.Descn
	f.fieldMap["menu"] = f.Menu
	f.fieldMap["create_time"] = f.CreateTime
	f.fieldMap["update_time"] = f.UpdateTime
	f.fieldMap["report_path"] = f.ReportPath
	f.fieldMap["project"] = f.Project
}

func (f frMainCore) clone(db *gorm.DB) frMainCore {
	f.frMainCoreDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f frMainCore) replaceDB(db *gorm.DB) frMainCore {
	f.frMainCoreDo.ReplaceDB(db)
	return f
}

type frMainCoreDo struct{ gen.DO }

type IFrMainCoreDo interface {
	gen.SubQuery
	Debug() IFrMainCoreDo
	WithContext(ctx context.Context) IFrMainCoreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFrMainCoreDo
	WriteDB() IFrMainCoreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFrMainCoreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFrMainCoreDo
	Not(conds ...gen.Condition) IFrMainCoreDo
	Or(conds ...gen.Condition) IFrMainCoreDo
	Select(conds ...field.Expr) IFrMainCoreDo
	Where(conds ...gen.Condition) IFrMainCoreDo
	Order(conds ...field.Expr) IFrMainCoreDo
	Distinct(cols ...field.Expr) IFrMainCoreDo
	Omit(cols ...field.Expr) IFrMainCoreDo
	Join(table schema.Tabler, on ...field.Expr) IFrMainCoreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFrMainCoreDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFrMainCoreDo
	Group(cols ...field.Expr) IFrMainCoreDo
	Having(conds ...gen.Condition) IFrMainCoreDo
	Limit(limit int) IFrMainCoreDo
	Offset(offset int) IFrMainCoreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFrMainCoreDo
	Unscoped() IFrMainCoreDo
	Create(values ...*model.FrMainCore) error
	CreateInBatches(values []*model.FrMainCore, batchSize int) error
	Save(values ...*model.FrMainCore) error
	First() (*model.FrMainCore, error)
	Take() (*model.FrMainCore, error)
	Last() (*model.FrMainCore, error)
	Find() ([]*model.FrMainCore, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FrMainCore, err error)
	FindInBatches(result *[]*model.FrMainCore, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FrMainCore) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFrMainCoreDo
	Assign(attrs ...field.AssignExpr) IFrMainCoreDo
	Joins(fields ...field.RelationField) IFrMainCoreDo
	Preload(fields ...field.RelationField) IFrMainCoreDo
	FirstOrInit() (*model.FrMainCore, error)
	FirstOrCreate() (*model.FrMainCore, error)
	FindByPage(offset int, limit int) (result []*model.FrMainCore, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFrMainCoreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f frMainCoreDo) Debug() IFrMainCoreDo {
	return f.withDO(f.DO.Debug())
}

func (f frMainCoreDo) WithContext(ctx context.Context) IFrMainCoreDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f frMainCoreDo) ReadDB() IFrMainCoreDo {
	return f.Clauses(dbresolver.Read)
}

func (f frMainCoreDo) WriteDB() IFrMainCoreDo {
	return f.Clauses(dbresolver.Write)
}

func (f frMainCoreDo) Session(config *gorm.Session) IFrMainCoreDo {
	return f.withDO(f.DO.Session(config))
}

func (f frMainCoreDo) Clauses(conds ...clause.Expression) IFrMainCoreDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f frMainCoreDo) Returning(value interface{}, columns ...string) IFrMainCoreDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f frMainCoreDo) Not(conds ...gen.Condition) IFrMainCoreDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f frMainCoreDo) Or(conds ...gen.Condition) IFrMainCoreDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f frMainCoreDo) Select(conds ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f frMainCoreDo) Where(conds ...gen.Condition) IFrMainCoreDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f frMainCoreDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFrMainCoreDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f frMainCoreDo) Order(conds ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f frMainCoreDo) Distinct(cols ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f frMainCoreDo) Omit(cols ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f frMainCoreDo) Join(table schema.Tabler, on ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f frMainCoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f frMainCoreDo) RightJoin(table schema.Tabler, on ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f frMainCoreDo) Group(cols ...field.Expr) IFrMainCoreDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f frMainCoreDo) Having(conds ...gen.Condition) IFrMainCoreDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f frMainCoreDo) Limit(limit int) IFrMainCoreDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f frMainCoreDo) Offset(offset int) IFrMainCoreDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f frMainCoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFrMainCoreDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f frMainCoreDo) Unscoped() IFrMainCoreDo {
	return f.withDO(f.DO.Unscoped())
}

func (f frMainCoreDo) Create(values ...*model.FrMainCore) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f frMainCoreDo) CreateInBatches(values []*model.FrMainCore, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f frMainCoreDo) Save(values ...*model.FrMainCore) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f frMainCoreDo) First() (*model.FrMainCore, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrMainCore), nil
	}
}

func (f frMainCoreDo) Take() (*model.FrMainCore, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrMainCore), nil
	}
}

func (f frMainCoreDo) Last() (*model.FrMainCore, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrMainCore), nil
	}
}

func (f frMainCoreDo) Find() ([]*model.FrMainCore, error) {
	result, err := f.DO.Find()
	return result.([]*model.FrMainCore), err
}

func (f frMainCoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FrMainCore, err error) {
	buf := make([]*model.FrMainCore, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f frMainCoreDo) FindInBatches(result *[]*model.FrMainCore, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f frMainCoreDo) Attrs(attrs ...field.AssignExpr) IFrMainCoreDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f frMainCoreDo) Assign(attrs ...field.AssignExpr) IFrMainCoreDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f frMainCoreDo) Joins(fields ...field.RelationField) IFrMainCoreDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f frMainCoreDo) Preload(fields ...field.RelationField) IFrMainCoreDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f frMainCoreDo) FirstOrInit() (*model.FrMainCore, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrMainCore), nil
	}
}

func (f frMainCoreDo) FirstOrCreate() (*model.FrMainCore, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FrMainCore), nil
	}
}

func (f frMainCoreDo) FindByPage(offset int, limit int) (result []*model.FrMainCore, count int64, err error) {
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

func (f frMainCoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f frMainCoreDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f frMainCoreDo) Delete(models ...*model.FrMainCore) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *frMainCoreDo) withDO(do gen.Dao) *frMainCoreDo {
	f.DO = *do.(*gen.DO)
	return f
}
