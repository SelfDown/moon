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

func newReqInfoModifyDetail(db *gorm.DB, opts ...gen.DOOption) reqInfoModifyDetail {
	_reqInfoModifyDetail := reqInfoModifyDetail{}

	_reqInfoModifyDetail.reqInfoModifyDetailDo.UseDB(db, opts...)
	_reqInfoModifyDetail.reqInfoModifyDetailDo.UseModel(&model.ReqInfoModifyDetail{})

	tableName := _reqInfoModifyDetail.reqInfoModifyDetailDo.TableName()
	_reqInfoModifyDetail.ALL = field.NewAsterisk(tableName)
	_reqInfoModifyDetail.TraceDetailID = field.NewString(tableName, "trace_detail_id")
	_reqInfoModifyDetail.TraceID = field.NewString(tableName, "trace_id")
	_reqInfoModifyDetail.ModifyKey = field.NewString(tableName, "modify_key")
	_reqInfoModifyDetail.ModifyKeyDec = field.NewString(tableName, "modify_key_dec")
	_reqInfoModifyDetail.ModifyBeforeValue = field.NewString(tableName, "modify_before_value")
	_reqInfoModifyDetail.ModifyAfterValue = field.NewString(tableName, "modify_after_value")
	_reqInfoModifyDetail.Note = field.NewString(tableName, "note")
	_reqInfoModifyDetail.CreateTime = field.NewTime(tableName, "create_time")
	_reqInfoModifyDetail.ModifyTime = field.NewTime(tableName, "modify_time")
	_reqInfoModifyDetail.Comments = field.NewString(tableName, "comments")
	_reqInfoModifyDetail.Operation = field.NewString(tableName, "operation")

	_reqInfoModifyDetail.fillFieldMap()

	return _reqInfoModifyDetail
}

type reqInfoModifyDetail struct {
	reqInfoModifyDetailDo

	ALL               field.Asterisk
	TraceDetailID     field.String
	TraceID           field.String
	ModifyKey         field.String
	ModifyKeyDec      field.String
	ModifyBeforeValue field.String
	ModifyAfterValue  field.String
	Note              field.String
	CreateTime        field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime        field.Time   // 记录修改时间（数据库自动写入）
	Comments          field.String // 备注说明
	Operation         field.String

	fieldMap map[string]field.Expr
}

func (r reqInfoModifyDetail) Table(newTableName string) *reqInfoModifyDetail {
	r.reqInfoModifyDetailDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reqInfoModifyDetail) As(alias string) *reqInfoModifyDetail {
	r.reqInfoModifyDetailDo.DO = *(r.reqInfoModifyDetailDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reqInfoModifyDetail) updateTableName(table string) *reqInfoModifyDetail {
	r.ALL = field.NewAsterisk(table)
	r.TraceDetailID = field.NewString(table, "trace_detail_id")
	r.TraceID = field.NewString(table, "trace_id")
	r.ModifyKey = field.NewString(table, "modify_key")
	r.ModifyKeyDec = field.NewString(table, "modify_key_dec")
	r.ModifyBeforeValue = field.NewString(table, "modify_before_value")
	r.ModifyAfterValue = field.NewString(table, "modify_after_value")
	r.Note = field.NewString(table, "note")
	r.CreateTime = field.NewTime(table, "create_time")
	r.ModifyTime = field.NewTime(table, "modify_time")
	r.Comments = field.NewString(table, "comments")
	r.Operation = field.NewString(table, "operation")

	r.fillFieldMap()

	return r
}

func (r *reqInfoModifyDetail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reqInfoModifyDetail) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 11)
	r.fieldMap["trace_detail_id"] = r.TraceDetailID
	r.fieldMap["trace_id"] = r.TraceID
	r.fieldMap["modify_key"] = r.ModifyKey
	r.fieldMap["modify_key_dec"] = r.ModifyKeyDec
	r.fieldMap["modify_before_value"] = r.ModifyBeforeValue
	r.fieldMap["modify_after_value"] = r.ModifyAfterValue
	r.fieldMap["note"] = r.Note
	r.fieldMap["create_time"] = r.CreateTime
	r.fieldMap["modify_time"] = r.ModifyTime
	r.fieldMap["comments"] = r.Comments
	r.fieldMap["operation"] = r.Operation
}

func (r reqInfoModifyDetail) clone(db *gorm.DB) reqInfoModifyDetail {
	r.reqInfoModifyDetailDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reqInfoModifyDetail) replaceDB(db *gorm.DB) reqInfoModifyDetail {
	r.reqInfoModifyDetailDo.ReplaceDB(db)
	return r
}

type reqInfoModifyDetailDo struct{ gen.DO }

type IReqInfoModifyDetailDo interface {
	gen.SubQuery
	Debug() IReqInfoModifyDetailDo
	WithContext(ctx context.Context) IReqInfoModifyDetailDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReqInfoModifyDetailDo
	WriteDB() IReqInfoModifyDetailDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReqInfoModifyDetailDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReqInfoModifyDetailDo
	Not(conds ...gen.Condition) IReqInfoModifyDetailDo
	Or(conds ...gen.Condition) IReqInfoModifyDetailDo
	Select(conds ...field.Expr) IReqInfoModifyDetailDo
	Where(conds ...gen.Condition) IReqInfoModifyDetailDo
	Order(conds ...field.Expr) IReqInfoModifyDetailDo
	Distinct(cols ...field.Expr) IReqInfoModifyDetailDo
	Omit(cols ...field.Expr) IReqInfoModifyDetailDo
	Join(table schema.Tabler, on ...field.Expr) IReqInfoModifyDetailDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReqInfoModifyDetailDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReqInfoModifyDetailDo
	Group(cols ...field.Expr) IReqInfoModifyDetailDo
	Having(conds ...gen.Condition) IReqInfoModifyDetailDo
	Limit(limit int) IReqInfoModifyDetailDo
	Offset(offset int) IReqInfoModifyDetailDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReqInfoModifyDetailDo
	Unscoped() IReqInfoModifyDetailDo
	Create(values ...*model.ReqInfoModifyDetail) error
	CreateInBatches(values []*model.ReqInfoModifyDetail, batchSize int) error
	Save(values ...*model.ReqInfoModifyDetail) error
	First() (*model.ReqInfoModifyDetail, error)
	Take() (*model.ReqInfoModifyDetail, error)
	Last() (*model.ReqInfoModifyDetail, error)
	Find() ([]*model.ReqInfoModifyDetail, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReqInfoModifyDetail, err error)
	FindInBatches(result *[]*model.ReqInfoModifyDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ReqInfoModifyDetail) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReqInfoModifyDetailDo
	Assign(attrs ...field.AssignExpr) IReqInfoModifyDetailDo
	Joins(fields ...field.RelationField) IReqInfoModifyDetailDo
	Preload(fields ...field.RelationField) IReqInfoModifyDetailDo
	FirstOrInit() (*model.ReqInfoModifyDetail, error)
	FirstOrCreate() (*model.ReqInfoModifyDetail, error)
	FindByPage(offset int, limit int) (result []*model.ReqInfoModifyDetail, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReqInfoModifyDetailDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reqInfoModifyDetailDo) Debug() IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Debug())
}

func (r reqInfoModifyDetailDo) WithContext(ctx context.Context) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reqInfoModifyDetailDo) ReadDB() IReqInfoModifyDetailDo {
	return r.Clauses(dbresolver.Read)
}

func (r reqInfoModifyDetailDo) WriteDB() IReqInfoModifyDetailDo {
	return r.Clauses(dbresolver.Write)
}

func (r reqInfoModifyDetailDo) Session(config *gorm.Session) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Session(config))
}

func (r reqInfoModifyDetailDo) Clauses(conds ...clause.Expression) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reqInfoModifyDetailDo) Returning(value interface{}, columns ...string) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reqInfoModifyDetailDo) Not(conds ...gen.Condition) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reqInfoModifyDetailDo) Or(conds ...gen.Condition) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reqInfoModifyDetailDo) Select(conds ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reqInfoModifyDetailDo) Where(conds ...gen.Condition) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reqInfoModifyDetailDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IReqInfoModifyDetailDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r reqInfoModifyDetailDo) Order(conds ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reqInfoModifyDetailDo) Distinct(cols ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reqInfoModifyDetailDo) Omit(cols ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reqInfoModifyDetailDo) Join(table schema.Tabler, on ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reqInfoModifyDetailDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reqInfoModifyDetailDo) RightJoin(table schema.Tabler, on ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reqInfoModifyDetailDo) Group(cols ...field.Expr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reqInfoModifyDetailDo) Having(conds ...gen.Condition) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reqInfoModifyDetailDo) Limit(limit int) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reqInfoModifyDetailDo) Offset(offset int) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reqInfoModifyDetailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reqInfoModifyDetailDo) Unscoped() IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reqInfoModifyDetailDo) Create(values ...*model.ReqInfoModifyDetail) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reqInfoModifyDetailDo) CreateInBatches(values []*model.ReqInfoModifyDetail, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reqInfoModifyDetailDo) Save(values ...*model.ReqInfoModifyDetail) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reqInfoModifyDetailDo) First() (*model.ReqInfoModifyDetail, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReqInfoModifyDetail), nil
	}
}

func (r reqInfoModifyDetailDo) Take() (*model.ReqInfoModifyDetail, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReqInfoModifyDetail), nil
	}
}

func (r reqInfoModifyDetailDo) Last() (*model.ReqInfoModifyDetail, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReqInfoModifyDetail), nil
	}
}

func (r reqInfoModifyDetailDo) Find() ([]*model.ReqInfoModifyDetail, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReqInfoModifyDetail), err
}

func (r reqInfoModifyDetailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReqInfoModifyDetail, err error) {
	buf := make([]*model.ReqInfoModifyDetail, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reqInfoModifyDetailDo) FindInBatches(result *[]*model.ReqInfoModifyDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reqInfoModifyDetailDo) Attrs(attrs ...field.AssignExpr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reqInfoModifyDetailDo) Assign(attrs ...field.AssignExpr) IReqInfoModifyDetailDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reqInfoModifyDetailDo) Joins(fields ...field.RelationField) IReqInfoModifyDetailDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reqInfoModifyDetailDo) Preload(fields ...field.RelationField) IReqInfoModifyDetailDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reqInfoModifyDetailDo) FirstOrInit() (*model.ReqInfoModifyDetail, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReqInfoModifyDetail), nil
	}
}

func (r reqInfoModifyDetailDo) FirstOrCreate() (*model.ReqInfoModifyDetail, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReqInfoModifyDetail), nil
	}
}

func (r reqInfoModifyDetailDo) FindByPage(offset int, limit int) (result []*model.ReqInfoModifyDetail, count int64, err error) {
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

func (r reqInfoModifyDetailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reqInfoModifyDetailDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reqInfoModifyDetailDo) Delete(models ...*model.ReqInfoModifyDetail) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reqInfoModifyDetailDo) withDO(do gen.Dao) *reqInfoModifyDetailDo {
	r.DO = *do.(*gen.DO)
	return r
}
