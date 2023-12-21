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

func newHisIssueStatu(db *gorm.DB, opts ...gen.DOOption) hisIssueStatu {
	_hisIssueStatu := hisIssueStatu{}

	_hisIssueStatu.hisIssueStatuDo.UseDB(db, opts...)
	_hisIssueStatu.hisIssueStatuDo.UseModel(&model.HisIssueStatu{})

	tableName := _hisIssueStatu.hisIssueStatuDo.TableName()
	_hisIssueStatu.ALL = field.NewAsterisk(tableName)
	_hisIssueStatu.HisIssueStatuID = field.NewString(tableName, "his_issue_statu_id")
	_hisIssueStatu.StatuText = field.NewString(tableName, "statu_text")
	_hisIssueStatu.Notes = field.NewString(tableName, "notes")
	_hisIssueStatu.Weight = field.NewInt32(tableName, "weight")
	_hisIssueStatu.OrderIndex = field.NewInt32(tableName, "order_index")
	_hisIssueStatu.NotifyURL = field.NewString(tableName, "notify_url")
	_hisIssueStatu.StatuTextCustomer = field.NewString(tableName, "statu_text_customer")

	_hisIssueStatu.fillFieldMap()

	return _hisIssueStatu
}

type hisIssueStatu struct {
	hisIssueStatuDo

	ALL               field.Asterisk
	HisIssueStatuID   field.String
	StatuText         field.String
	Notes             field.String
	Weight            field.Int32
	OrderIndex        field.Int32
	NotifyURL         field.String
	StatuTextCustomer field.String

	fieldMap map[string]field.Expr
}

func (h hisIssueStatu) Table(newTableName string) *hisIssueStatu {
	h.hisIssueStatuDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hisIssueStatu) As(alias string) *hisIssueStatu {
	h.hisIssueStatuDo.DO = *(h.hisIssueStatuDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hisIssueStatu) updateTableName(table string) *hisIssueStatu {
	h.ALL = field.NewAsterisk(table)
	h.HisIssueStatuID = field.NewString(table, "his_issue_statu_id")
	h.StatuText = field.NewString(table, "statu_text")
	h.Notes = field.NewString(table, "notes")
	h.Weight = field.NewInt32(table, "weight")
	h.OrderIndex = field.NewInt32(table, "order_index")
	h.NotifyURL = field.NewString(table, "notify_url")
	h.StatuTextCustomer = field.NewString(table, "statu_text_customer")

	h.fillFieldMap()

	return h
}

func (h *hisIssueStatu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hisIssueStatu) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 7)
	h.fieldMap["his_issue_statu_id"] = h.HisIssueStatuID
	h.fieldMap["statu_text"] = h.StatuText
	h.fieldMap["notes"] = h.Notes
	h.fieldMap["weight"] = h.Weight
	h.fieldMap["order_index"] = h.OrderIndex
	h.fieldMap["notify_url"] = h.NotifyURL
	h.fieldMap["statu_text_customer"] = h.StatuTextCustomer
}

func (h hisIssueStatu) clone(db *gorm.DB) hisIssueStatu {
	h.hisIssueStatuDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hisIssueStatu) replaceDB(db *gorm.DB) hisIssueStatu {
	h.hisIssueStatuDo.ReplaceDB(db)
	return h
}

type hisIssueStatuDo struct{ gen.DO }

type IHisIssueStatuDo interface {
	gen.SubQuery
	Debug() IHisIssueStatuDo
	WithContext(ctx context.Context) IHisIssueStatuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHisIssueStatuDo
	WriteDB() IHisIssueStatuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHisIssueStatuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHisIssueStatuDo
	Not(conds ...gen.Condition) IHisIssueStatuDo
	Or(conds ...gen.Condition) IHisIssueStatuDo
	Select(conds ...field.Expr) IHisIssueStatuDo
	Where(conds ...gen.Condition) IHisIssueStatuDo
	Order(conds ...field.Expr) IHisIssueStatuDo
	Distinct(cols ...field.Expr) IHisIssueStatuDo
	Omit(cols ...field.Expr) IHisIssueStatuDo
	Join(table schema.Tabler, on ...field.Expr) IHisIssueStatuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHisIssueStatuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHisIssueStatuDo
	Group(cols ...field.Expr) IHisIssueStatuDo
	Having(conds ...gen.Condition) IHisIssueStatuDo
	Limit(limit int) IHisIssueStatuDo
	Offset(offset int) IHisIssueStatuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHisIssueStatuDo
	Unscoped() IHisIssueStatuDo
	Create(values ...*model.HisIssueStatu) error
	CreateInBatches(values []*model.HisIssueStatu, batchSize int) error
	Save(values ...*model.HisIssueStatu) error
	First() (*model.HisIssueStatu, error)
	Take() (*model.HisIssueStatu, error)
	Last() (*model.HisIssueStatu, error)
	Find() ([]*model.HisIssueStatu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HisIssueStatu, err error)
	FindInBatches(result *[]*model.HisIssueStatu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.HisIssueStatu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHisIssueStatuDo
	Assign(attrs ...field.AssignExpr) IHisIssueStatuDo
	Joins(fields ...field.RelationField) IHisIssueStatuDo
	Preload(fields ...field.RelationField) IHisIssueStatuDo
	FirstOrInit() (*model.HisIssueStatu, error)
	FirstOrCreate() (*model.HisIssueStatu, error)
	FindByPage(offset int, limit int) (result []*model.HisIssueStatu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHisIssueStatuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hisIssueStatuDo) Debug() IHisIssueStatuDo {
	return h.withDO(h.DO.Debug())
}

func (h hisIssueStatuDo) WithContext(ctx context.Context) IHisIssueStatuDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hisIssueStatuDo) ReadDB() IHisIssueStatuDo {
	return h.Clauses(dbresolver.Read)
}

func (h hisIssueStatuDo) WriteDB() IHisIssueStatuDo {
	return h.Clauses(dbresolver.Write)
}

func (h hisIssueStatuDo) Session(config *gorm.Session) IHisIssueStatuDo {
	return h.withDO(h.DO.Session(config))
}

func (h hisIssueStatuDo) Clauses(conds ...clause.Expression) IHisIssueStatuDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hisIssueStatuDo) Returning(value interface{}, columns ...string) IHisIssueStatuDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hisIssueStatuDo) Not(conds ...gen.Condition) IHisIssueStatuDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hisIssueStatuDo) Or(conds ...gen.Condition) IHisIssueStatuDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hisIssueStatuDo) Select(conds ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hisIssueStatuDo) Where(conds ...gen.Condition) IHisIssueStatuDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hisIssueStatuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IHisIssueStatuDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h hisIssueStatuDo) Order(conds ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hisIssueStatuDo) Distinct(cols ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hisIssueStatuDo) Omit(cols ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hisIssueStatuDo) Join(table schema.Tabler, on ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hisIssueStatuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hisIssueStatuDo) RightJoin(table schema.Tabler, on ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hisIssueStatuDo) Group(cols ...field.Expr) IHisIssueStatuDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hisIssueStatuDo) Having(conds ...gen.Condition) IHisIssueStatuDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hisIssueStatuDo) Limit(limit int) IHisIssueStatuDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hisIssueStatuDo) Offset(offset int) IHisIssueStatuDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hisIssueStatuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHisIssueStatuDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hisIssueStatuDo) Unscoped() IHisIssueStatuDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hisIssueStatuDo) Create(values ...*model.HisIssueStatu) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hisIssueStatuDo) CreateInBatches(values []*model.HisIssueStatu, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hisIssueStatuDo) Save(values ...*model.HisIssueStatu) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hisIssueStatuDo) First() (*model.HisIssueStatu, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.HisIssueStatu), nil
	}
}

func (h hisIssueStatuDo) Take() (*model.HisIssueStatu, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.HisIssueStatu), nil
	}
}

func (h hisIssueStatuDo) Last() (*model.HisIssueStatu, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.HisIssueStatu), nil
	}
}

func (h hisIssueStatuDo) Find() ([]*model.HisIssueStatu, error) {
	result, err := h.DO.Find()
	return result.([]*model.HisIssueStatu), err
}

func (h hisIssueStatuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.HisIssueStatu, err error) {
	buf := make([]*model.HisIssueStatu, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hisIssueStatuDo) FindInBatches(result *[]*model.HisIssueStatu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hisIssueStatuDo) Attrs(attrs ...field.AssignExpr) IHisIssueStatuDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hisIssueStatuDo) Assign(attrs ...field.AssignExpr) IHisIssueStatuDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hisIssueStatuDo) Joins(fields ...field.RelationField) IHisIssueStatuDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hisIssueStatuDo) Preload(fields ...field.RelationField) IHisIssueStatuDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hisIssueStatuDo) FirstOrInit() (*model.HisIssueStatu, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.HisIssueStatu), nil
	}
}

func (h hisIssueStatuDo) FirstOrCreate() (*model.HisIssueStatu, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.HisIssueStatu), nil
	}
}

func (h hisIssueStatuDo) FindByPage(offset int, limit int) (result []*model.HisIssueStatu, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hisIssueStatuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hisIssueStatuDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hisIssueStatuDo) Delete(models ...*model.HisIssueStatu) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hisIssueStatuDo) withDO(do gen.Dao) *hisIssueStatuDo {
	h.DO = *do.(*gen.DO)
	return h
}