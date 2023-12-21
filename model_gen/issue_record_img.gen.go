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

func newIssueRecordImg(db *gorm.DB, opts ...gen.DOOption) issueRecordImg {
	_issueRecordImg := issueRecordImg{}

	_issueRecordImg.issueRecordImgDo.UseDB(db, opts...)
	_issueRecordImg.issueRecordImgDo.UseModel(&model.IssueRecordImg{})

	tableName := _issueRecordImg.issueRecordImgDo.TableName()
	_issueRecordImg.ALL = field.NewAsterisk(tableName)
	_issueRecordImg.IssueRecordImgID = field.NewString(tableName, "issue_record_img_id")
	_issueRecordImg.Data = field.NewString(tableName, "data")
	_issueRecordImg.AddTime = field.NewString(tableName, "add_time")

	_issueRecordImg.fillFieldMap()

	return _issueRecordImg
}

type issueRecordImg struct {
	issueRecordImgDo

	ALL              field.Asterisk
	IssueRecordImgID field.String
	Data             field.String // 文件内容
	AddTime          field.String

	fieldMap map[string]field.Expr
}

func (i issueRecordImg) Table(newTableName string) *issueRecordImg {
	i.issueRecordImgDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i issueRecordImg) As(alias string) *issueRecordImg {
	i.issueRecordImgDo.DO = *(i.issueRecordImgDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *issueRecordImg) updateTableName(table string) *issueRecordImg {
	i.ALL = field.NewAsterisk(table)
	i.IssueRecordImgID = field.NewString(table, "issue_record_img_id")
	i.Data = field.NewString(table, "data")
	i.AddTime = field.NewString(table, "add_time")

	i.fillFieldMap()

	return i
}

func (i *issueRecordImg) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *issueRecordImg) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 3)
	i.fieldMap["issue_record_img_id"] = i.IssueRecordImgID
	i.fieldMap["data"] = i.Data
	i.fieldMap["add_time"] = i.AddTime
}

func (i issueRecordImg) clone(db *gorm.DB) issueRecordImg {
	i.issueRecordImgDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i issueRecordImg) replaceDB(db *gorm.DB) issueRecordImg {
	i.issueRecordImgDo.ReplaceDB(db)
	return i
}

type issueRecordImgDo struct{ gen.DO }

type IIssueRecordImgDo interface {
	gen.SubQuery
	Debug() IIssueRecordImgDo
	WithContext(ctx context.Context) IIssueRecordImgDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIssueRecordImgDo
	WriteDB() IIssueRecordImgDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIssueRecordImgDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIssueRecordImgDo
	Not(conds ...gen.Condition) IIssueRecordImgDo
	Or(conds ...gen.Condition) IIssueRecordImgDo
	Select(conds ...field.Expr) IIssueRecordImgDo
	Where(conds ...gen.Condition) IIssueRecordImgDo
	Order(conds ...field.Expr) IIssueRecordImgDo
	Distinct(cols ...field.Expr) IIssueRecordImgDo
	Omit(cols ...field.Expr) IIssueRecordImgDo
	Join(table schema.Tabler, on ...field.Expr) IIssueRecordImgDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIssueRecordImgDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIssueRecordImgDo
	Group(cols ...field.Expr) IIssueRecordImgDo
	Having(conds ...gen.Condition) IIssueRecordImgDo
	Limit(limit int) IIssueRecordImgDo
	Offset(offset int) IIssueRecordImgDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIssueRecordImgDo
	Unscoped() IIssueRecordImgDo
	Create(values ...*model.IssueRecordImg) error
	CreateInBatches(values []*model.IssueRecordImg, batchSize int) error
	Save(values ...*model.IssueRecordImg) error
	First() (*model.IssueRecordImg, error)
	Take() (*model.IssueRecordImg, error)
	Last() (*model.IssueRecordImg, error)
	Find() ([]*model.IssueRecordImg, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IssueRecordImg, err error)
	FindInBatches(result *[]*model.IssueRecordImg, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.IssueRecordImg) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIssueRecordImgDo
	Assign(attrs ...field.AssignExpr) IIssueRecordImgDo
	Joins(fields ...field.RelationField) IIssueRecordImgDo
	Preload(fields ...field.RelationField) IIssueRecordImgDo
	FirstOrInit() (*model.IssueRecordImg, error)
	FirstOrCreate() (*model.IssueRecordImg, error)
	FindByPage(offset int, limit int) (result []*model.IssueRecordImg, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIssueRecordImgDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i issueRecordImgDo) Debug() IIssueRecordImgDo {
	return i.withDO(i.DO.Debug())
}

func (i issueRecordImgDo) WithContext(ctx context.Context) IIssueRecordImgDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i issueRecordImgDo) ReadDB() IIssueRecordImgDo {
	return i.Clauses(dbresolver.Read)
}

func (i issueRecordImgDo) WriteDB() IIssueRecordImgDo {
	return i.Clauses(dbresolver.Write)
}

func (i issueRecordImgDo) Session(config *gorm.Session) IIssueRecordImgDo {
	return i.withDO(i.DO.Session(config))
}

func (i issueRecordImgDo) Clauses(conds ...clause.Expression) IIssueRecordImgDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i issueRecordImgDo) Returning(value interface{}, columns ...string) IIssueRecordImgDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i issueRecordImgDo) Not(conds ...gen.Condition) IIssueRecordImgDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i issueRecordImgDo) Or(conds ...gen.Condition) IIssueRecordImgDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i issueRecordImgDo) Select(conds ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i issueRecordImgDo) Where(conds ...gen.Condition) IIssueRecordImgDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i issueRecordImgDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IIssueRecordImgDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i issueRecordImgDo) Order(conds ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i issueRecordImgDo) Distinct(cols ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i issueRecordImgDo) Omit(cols ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i issueRecordImgDo) Join(table schema.Tabler, on ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i issueRecordImgDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i issueRecordImgDo) RightJoin(table schema.Tabler, on ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i issueRecordImgDo) Group(cols ...field.Expr) IIssueRecordImgDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i issueRecordImgDo) Having(conds ...gen.Condition) IIssueRecordImgDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i issueRecordImgDo) Limit(limit int) IIssueRecordImgDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i issueRecordImgDo) Offset(offset int) IIssueRecordImgDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i issueRecordImgDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIssueRecordImgDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i issueRecordImgDo) Unscoped() IIssueRecordImgDo {
	return i.withDO(i.DO.Unscoped())
}

func (i issueRecordImgDo) Create(values ...*model.IssueRecordImg) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i issueRecordImgDo) CreateInBatches(values []*model.IssueRecordImg, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i issueRecordImgDo) Save(values ...*model.IssueRecordImg) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i issueRecordImgDo) First() (*model.IssueRecordImg, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueRecordImg), nil
	}
}

func (i issueRecordImgDo) Take() (*model.IssueRecordImg, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueRecordImg), nil
	}
}

func (i issueRecordImgDo) Last() (*model.IssueRecordImg, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueRecordImg), nil
	}
}

func (i issueRecordImgDo) Find() ([]*model.IssueRecordImg, error) {
	result, err := i.DO.Find()
	return result.([]*model.IssueRecordImg), err
}

func (i issueRecordImgDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.IssueRecordImg, err error) {
	buf := make([]*model.IssueRecordImg, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i issueRecordImgDo) FindInBatches(result *[]*model.IssueRecordImg, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i issueRecordImgDo) Attrs(attrs ...field.AssignExpr) IIssueRecordImgDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i issueRecordImgDo) Assign(attrs ...field.AssignExpr) IIssueRecordImgDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i issueRecordImgDo) Joins(fields ...field.RelationField) IIssueRecordImgDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i issueRecordImgDo) Preload(fields ...field.RelationField) IIssueRecordImgDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i issueRecordImgDo) FirstOrInit() (*model.IssueRecordImg, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueRecordImg), nil
	}
}

func (i issueRecordImgDo) FirstOrCreate() (*model.IssueRecordImg, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.IssueRecordImg), nil
	}
}

func (i issueRecordImgDo) FindByPage(offset int, limit int) (result []*model.IssueRecordImg, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i issueRecordImgDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i issueRecordImgDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i issueRecordImgDo) Delete(models ...*model.IssueRecordImg) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *issueRecordImgDo) withDO(do gen.Dao) *issueRecordImgDo {
	i.DO = *do.(*gen.DO)
	return i
}