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

func newSysFaq(db *gorm.DB, opts ...gen.DOOption) sysFaq {
	_sysFaq := sysFaq{}

	_sysFaq.sysFaqDo.UseDB(db, opts...)
	_sysFaq.sysFaqDo.UseModel(&model.SysFaq{})

	tableName := _sysFaq.sysFaqDo.TableName()
	_sysFaq.ALL = field.NewAsterisk(tableName)
	_sysFaq.SysFaqID = field.NewString(tableName, "sys_faq_id")
	_sysFaq.Question = field.NewString(tableName, "question")
	_sysFaq.Answer = field.NewString(tableName, "answer")
	_sysFaq.CreateTime = field.NewTime(tableName, "create_time")
	_sysFaq.ModifyTime = field.NewTime(tableName, "modify_time")
	_sysFaq.Comments = field.NewString(tableName, "comments")
	_sysFaq.AgreeCount = field.NewInt32(tableName, "agree_count")
	_sysFaq.OpposeCount = field.NewInt32(tableName, "oppose_count")

	_sysFaq.fillFieldMap()

	return _sysFaq
}

type sysFaq struct {
	sysFaqDo

	ALL         field.Asterisk
	SysFaqID    field.String
	Question    field.String // 问题
	Answer      field.String // 问题答案
	CreateTime  field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime  field.Time
	Comments    field.String
	AgreeCount  field.Int32 // 问题被赞同的数量
	OpposeCount field.Int32 // 问题被不赞同的数量

	fieldMap map[string]field.Expr
}

func (s sysFaq) Table(newTableName string) *sysFaq {
	s.sysFaqDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysFaq) As(alias string) *sysFaq {
	s.sysFaqDo.DO = *(s.sysFaqDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysFaq) updateTableName(table string) *sysFaq {
	s.ALL = field.NewAsterisk(table)
	s.SysFaqID = field.NewString(table, "sys_faq_id")
	s.Question = field.NewString(table, "question")
	s.Answer = field.NewString(table, "answer")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")
	s.AgreeCount = field.NewInt32(table, "agree_count")
	s.OpposeCount = field.NewInt32(table, "oppose_count")

	s.fillFieldMap()

	return s
}

func (s *sysFaq) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysFaq) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["sys_faq_id"] = s.SysFaqID
	s.fieldMap["question"] = s.Question
	s.fieldMap["answer"] = s.Answer
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["agree_count"] = s.AgreeCount
	s.fieldMap["oppose_count"] = s.OpposeCount
}

func (s sysFaq) clone(db *gorm.DB) sysFaq {
	s.sysFaqDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysFaq) replaceDB(db *gorm.DB) sysFaq {
	s.sysFaqDo.ReplaceDB(db)
	return s
}

type sysFaqDo struct{ gen.DO }

type ISysFaqDo interface {
	gen.SubQuery
	Debug() ISysFaqDo
	WithContext(ctx context.Context) ISysFaqDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysFaqDo
	WriteDB() ISysFaqDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysFaqDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysFaqDo
	Not(conds ...gen.Condition) ISysFaqDo
	Or(conds ...gen.Condition) ISysFaqDo
	Select(conds ...field.Expr) ISysFaqDo
	Where(conds ...gen.Condition) ISysFaqDo
	Order(conds ...field.Expr) ISysFaqDo
	Distinct(cols ...field.Expr) ISysFaqDo
	Omit(cols ...field.Expr) ISysFaqDo
	Join(table schema.Tabler, on ...field.Expr) ISysFaqDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysFaqDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysFaqDo
	Group(cols ...field.Expr) ISysFaqDo
	Having(conds ...gen.Condition) ISysFaqDo
	Limit(limit int) ISysFaqDo
	Offset(offset int) ISysFaqDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysFaqDo
	Unscoped() ISysFaqDo
	Create(values ...*model.SysFaq) error
	CreateInBatches(values []*model.SysFaq, batchSize int) error
	Save(values ...*model.SysFaq) error
	First() (*model.SysFaq, error)
	Take() (*model.SysFaq, error)
	Last() (*model.SysFaq, error)
	Find() ([]*model.SysFaq, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysFaq, err error)
	FindInBatches(result *[]*model.SysFaq, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysFaq) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysFaqDo
	Assign(attrs ...field.AssignExpr) ISysFaqDo
	Joins(fields ...field.RelationField) ISysFaqDo
	Preload(fields ...field.RelationField) ISysFaqDo
	FirstOrInit() (*model.SysFaq, error)
	FirstOrCreate() (*model.SysFaq, error)
	FindByPage(offset int, limit int) (result []*model.SysFaq, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysFaqDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysFaqDo) Debug() ISysFaqDo {
	return s.withDO(s.DO.Debug())
}

func (s sysFaqDo) WithContext(ctx context.Context) ISysFaqDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysFaqDo) ReadDB() ISysFaqDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysFaqDo) WriteDB() ISysFaqDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysFaqDo) Session(config *gorm.Session) ISysFaqDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysFaqDo) Clauses(conds ...clause.Expression) ISysFaqDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysFaqDo) Returning(value interface{}, columns ...string) ISysFaqDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysFaqDo) Not(conds ...gen.Condition) ISysFaqDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysFaqDo) Or(conds ...gen.Condition) ISysFaqDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysFaqDo) Select(conds ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysFaqDo) Where(conds ...gen.Condition) ISysFaqDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysFaqDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysFaqDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysFaqDo) Order(conds ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysFaqDo) Distinct(cols ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysFaqDo) Omit(cols ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysFaqDo) Join(table schema.Tabler, on ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysFaqDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysFaqDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysFaqDo) Group(cols ...field.Expr) ISysFaqDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysFaqDo) Having(conds ...gen.Condition) ISysFaqDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysFaqDo) Limit(limit int) ISysFaqDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysFaqDo) Offset(offset int) ISysFaqDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysFaqDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysFaqDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysFaqDo) Unscoped() ISysFaqDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysFaqDo) Create(values ...*model.SysFaq) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysFaqDo) CreateInBatches(values []*model.SysFaq, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysFaqDo) Save(values ...*model.SysFaq) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysFaqDo) First() (*model.SysFaq, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFaq), nil
	}
}

func (s sysFaqDo) Take() (*model.SysFaq, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFaq), nil
	}
}

func (s sysFaqDo) Last() (*model.SysFaq, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFaq), nil
	}
}

func (s sysFaqDo) Find() ([]*model.SysFaq, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysFaq), err
}

func (s sysFaqDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysFaq, err error) {
	buf := make([]*model.SysFaq, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysFaqDo) FindInBatches(result *[]*model.SysFaq, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysFaqDo) Attrs(attrs ...field.AssignExpr) ISysFaqDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysFaqDo) Assign(attrs ...field.AssignExpr) ISysFaqDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysFaqDo) Joins(fields ...field.RelationField) ISysFaqDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysFaqDo) Preload(fields ...field.RelationField) ISysFaqDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysFaqDo) FirstOrInit() (*model.SysFaq, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFaq), nil
	}
}

func (s sysFaqDo) FirstOrCreate() (*model.SysFaq, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysFaq), nil
	}
}

func (s sysFaqDo) FindByPage(offset int, limit int) (result []*model.SysFaq, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysFaqDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysFaqDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysFaqDo) Delete(models ...*model.SysFaq) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysFaqDo) withDO(do gen.Dao) *sysFaqDo {
	s.DO = *do.(*gen.DO)
	return s
}