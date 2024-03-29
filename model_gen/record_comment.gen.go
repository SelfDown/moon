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

func newRecordComment(db *gorm.DB, opts ...gen.DOOption) recordComment {
	_recordComment := recordComment{}

	_recordComment.recordCommentDo.UseDB(db, opts...)
	_recordComment.recordCommentDo.UseModel(&model.RecordComment{})

	tableName := _recordComment.recordCommentDo.TableName()
	_recordComment.ALL = field.NewAsterisk(tableName)
	_recordComment.RecordCommentID = field.NewString(tableName, "record_comment_id")
	_recordComment.HisIssueRecordID = field.NewString(tableName, "his_issue_record_id")
	_recordComment.HasFix = field.NewString(tableName, "has_fix")
	_recordComment.WorkCode = field.NewString(tableName, "work_code")
	_recordComment.Presenter = field.NewString(tableName, "presenter")
	_recordComment.Star = field.NewString(tableName, "star")
	_recordComment.Comment = field.NewString(tableName, "comment")
	_recordComment.CreateTime = field.NewString(tableName, "create_time")

	_recordComment.fillFieldMap()

	return _recordComment
}

type recordComment struct {
	recordCommentDo

	ALL              field.Asterisk
	RecordCommentID  field.String
	HisIssueRecordID field.String
	HasFix           field.String
	WorkCode         field.String
	Presenter        field.String
	Star             field.String
	Comment          field.String
	CreateTime       field.String

	fieldMap map[string]field.Expr
}

func (r recordComment) Table(newTableName string) *recordComment {
	r.recordCommentDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r recordComment) As(alias string) *recordComment {
	r.recordCommentDo.DO = *(r.recordCommentDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *recordComment) updateTableName(table string) *recordComment {
	r.ALL = field.NewAsterisk(table)
	r.RecordCommentID = field.NewString(table, "record_comment_id")
	r.HisIssueRecordID = field.NewString(table, "his_issue_record_id")
	r.HasFix = field.NewString(table, "has_fix")
	r.WorkCode = field.NewString(table, "work_code")
	r.Presenter = field.NewString(table, "presenter")
	r.Star = field.NewString(table, "star")
	r.Comment = field.NewString(table, "comment")
	r.CreateTime = field.NewString(table, "create_time")

	r.fillFieldMap()

	return r
}

func (r *recordComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *recordComment) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 8)
	r.fieldMap["record_comment_id"] = r.RecordCommentID
	r.fieldMap["his_issue_record_id"] = r.HisIssueRecordID
	r.fieldMap["has_fix"] = r.HasFix
	r.fieldMap["work_code"] = r.WorkCode
	r.fieldMap["presenter"] = r.Presenter
	r.fieldMap["star"] = r.Star
	r.fieldMap["comment"] = r.Comment
	r.fieldMap["create_time"] = r.CreateTime
}

func (r recordComment) clone(db *gorm.DB) recordComment {
	r.recordCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r recordComment) replaceDB(db *gorm.DB) recordComment {
	r.recordCommentDo.ReplaceDB(db)
	return r
}

type recordCommentDo struct{ gen.DO }

type IRecordCommentDo interface {
	gen.SubQuery
	Debug() IRecordCommentDo
	WithContext(ctx context.Context) IRecordCommentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRecordCommentDo
	WriteDB() IRecordCommentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRecordCommentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRecordCommentDo
	Not(conds ...gen.Condition) IRecordCommentDo
	Or(conds ...gen.Condition) IRecordCommentDo
	Select(conds ...field.Expr) IRecordCommentDo
	Where(conds ...gen.Condition) IRecordCommentDo
	Order(conds ...field.Expr) IRecordCommentDo
	Distinct(cols ...field.Expr) IRecordCommentDo
	Omit(cols ...field.Expr) IRecordCommentDo
	Join(table schema.Tabler, on ...field.Expr) IRecordCommentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRecordCommentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRecordCommentDo
	Group(cols ...field.Expr) IRecordCommentDo
	Having(conds ...gen.Condition) IRecordCommentDo
	Limit(limit int) IRecordCommentDo
	Offset(offset int) IRecordCommentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRecordCommentDo
	Unscoped() IRecordCommentDo
	Create(values ...*model.RecordComment) error
	CreateInBatches(values []*model.RecordComment, batchSize int) error
	Save(values ...*model.RecordComment) error
	First() (*model.RecordComment, error)
	Take() (*model.RecordComment, error)
	Last() (*model.RecordComment, error)
	Find() ([]*model.RecordComment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RecordComment, err error)
	FindInBatches(result *[]*model.RecordComment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RecordComment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRecordCommentDo
	Assign(attrs ...field.AssignExpr) IRecordCommentDo
	Joins(fields ...field.RelationField) IRecordCommentDo
	Preload(fields ...field.RelationField) IRecordCommentDo
	FirstOrInit() (*model.RecordComment, error)
	FirstOrCreate() (*model.RecordComment, error)
	FindByPage(offset int, limit int) (result []*model.RecordComment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRecordCommentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r recordCommentDo) Debug() IRecordCommentDo {
	return r.withDO(r.DO.Debug())
}

func (r recordCommentDo) WithContext(ctx context.Context) IRecordCommentDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r recordCommentDo) ReadDB() IRecordCommentDo {
	return r.Clauses(dbresolver.Read)
}

func (r recordCommentDo) WriteDB() IRecordCommentDo {
	return r.Clauses(dbresolver.Write)
}

func (r recordCommentDo) Session(config *gorm.Session) IRecordCommentDo {
	return r.withDO(r.DO.Session(config))
}

func (r recordCommentDo) Clauses(conds ...clause.Expression) IRecordCommentDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r recordCommentDo) Returning(value interface{}, columns ...string) IRecordCommentDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r recordCommentDo) Not(conds ...gen.Condition) IRecordCommentDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r recordCommentDo) Or(conds ...gen.Condition) IRecordCommentDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r recordCommentDo) Select(conds ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r recordCommentDo) Where(conds ...gen.Condition) IRecordCommentDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r recordCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRecordCommentDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r recordCommentDo) Order(conds ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r recordCommentDo) Distinct(cols ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r recordCommentDo) Omit(cols ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r recordCommentDo) Join(table schema.Tabler, on ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r recordCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r recordCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r recordCommentDo) Group(cols ...field.Expr) IRecordCommentDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r recordCommentDo) Having(conds ...gen.Condition) IRecordCommentDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r recordCommentDo) Limit(limit int) IRecordCommentDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r recordCommentDo) Offset(offset int) IRecordCommentDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r recordCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRecordCommentDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r recordCommentDo) Unscoped() IRecordCommentDo {
	return r.withDO(r.DO.Unscoped())
}

func (r recordCommentDo) Create(values ...*model.RecordComment) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r recordCommentDo) CreateInBatches(values []*model.RecordComment, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r recordCommentDo) Save(values ...*model.RecordComment) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r recordCommentDo) First() (*model.RecordComment, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordComment), nil
	}
}

func (r recordCommentDo) Take() (*model.RecordComment, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordComment), nil
	}
}

func (r recordCommentDo) Last() (*model.RecordComment, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordComment), nil
	}
}

func (r recordCommentDo) Find() ([]*model.RecordComment, error) {
	result, err := r.DO.Find()
	return result.([]*model.RecordComment), err
}

func (r recordCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RecordComment, err error) {
	buf := make([]*model.RecordComment, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r recordCommentDo) FindInBatches(result *[]*model.RecordComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r recordCommentDo) Attrs(attrs ...field.AssignExpr) IRecordCommentDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r recordCommentDo) Assign(attrs ...field.AssignExpr) IRecordCommentDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r recordCommentDo) Joins(fields ...field.RelationField) IRecordCommentDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r recordCommentDo) Preload(fields ...field.RelationField) IRecordCommentDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r recordCommentDo) FirstOrInit() (*model.RecordComment, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordComment), nil
	}
}

func (r recordCommentDo) FirstOrCreate() (*model.RecordComment, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RecordComment), nil
	}
}

func (r recordCommentDo) FindByPage(offset int, limit int) (result []*model.RecordComment, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r recordCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r recordCommentDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r recordCommentDo) Delete(models ...*model.RecordComment) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *recordCommentDo) withDO(do gen.Dao) *recordCommentDo {
	r.DO = *do.(*gen.DO)
	return r
}
