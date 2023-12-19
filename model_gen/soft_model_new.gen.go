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

func newSoftModelNew(db *gorm.DB, opts ...gen.DOOption) softModelNew {
	_softModelNew := softModelNew{}

	_softModelNew.softModelNewDo.UseDB(db, opts...)
	_softModelNew.softModelNewDo.UseModel(&model.SoftModelNew{})

	tableName := _softModelNew.softModelNewDo.TableName()
	_softModelNew.ALL = field.NewAsterisk(tableName)
	_softModelNew.SoftModelName = field.NewString(tableName, "soft_model_name")
	_softModelNew.SoftType = field.NewString(tableName, "soft_type")
	_softModelNew.SoftDataJSON = field.NewString(tableName, "soft_data_json")
	_softModelNew.UUIDID = field.NewString(tableName, "uuid_id")
	_softModelNew.IsDelete = field.NewString(tableName, "is_delete")
	_softModelNew.OpUser = field.NewString(tableName, "op_user")

	_softModelNew.fillFieldMap()

	return _softModelNew
}

type softModelNew struct {
	softModelNewDo

	ALL           field.Asterisk
	SoftModelName field.String
	SoftType      field.String
	SoftDataJSON  field.String
	UUIDID        field.String
	IsDelete      field.String
	OpUser        field.String

	fieldMap map[string]field.Expr
}

func (s softModelNew) Table(newTableName string) *softModelNew {
	s.softModelNewDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s softModelNew) As(alias string) *softModelNew {
	s.softModelNewDo.DO = *(s.softModelNewDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *softModelNew) updateTableName(table string) *softModelNew {
	s.ALL = field.NewAsterisk(table)
	s.SoftModelName = field.NewString(table, "soft_model_name")
	s.SoftType = field.NewString(table, "soft_type")
	s.SoftDataJSON = field.NewString(table, "soft_data_json")
	s.UUIDID = field.NewString(table, "uuid_id")
	s.IsDelete = field.NewString(table, "is_delete")
	s.OpUser = field.NewString(table, "op_user")

	s.fillFieldMap()

	return s
}

func (s *softModelNew) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *softModelNew) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["soft_model_name"] = s.SoftModelName
	s.fieldMap["soft_type"] = s.SoftType
	s.fieldMap["soft_data_json"] = s.SoftDataJSON
	s.fieldMap["uuid_id"] = s.UUIDID
	s.fieldMap["is_delete"] = s.IsDelete
	s.fieldMap["op_user"] = s.OpUser
}

func (s softModelNew) clone(db *gorm.DB) softModelNew {
	s.softModelNewDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s softModelNew) replaceDB(db *gorm.DB) softModelNew {
	s.softModelNewDo.ReplaceDB(db)
	return s
}

type softModelNewDo struct{ gen.DO }

type ISoftModelNewDo interface {
	gen.SubQuery
	Debug() ISoftModelNewDo
	WithContext(ctx context.Context) ISoftModelNewDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISoftModelNewDo
	WriteDB() ISoftModelNewDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISoftModelNewDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISoftModelNewDo
	Not(conds ...gen.Condition) ISoftModelNewDo
	Or(conds ...gen.Condition) ISoftModelNewDo
	Select(conds ...field.Expr) ISoftModelNewDo
	Where(conds ...gen.Condition) ISoftModelNewDo
	Order(conds ...field.Expr) ISoftModelNewDo
	Distinct(cols ...field.Expr) ISoftModelNewDo
	Omit(cols ...field.Expr) ISoftModelNewDo
	Join(table schema.Tabler, on ...field.Expr) ISoftModelNewDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISoftModelNewDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISoftModelNewDo
	Group(cols ...field.Expr) ISoftModelNewDo
	Having(conds ...gen.Condition) ISoftModelNewDo
	Limit(limit int) ISoftModelNewDo
	Offset(offset int) ISoftModelNewDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISoftModelNewDo
	Unscoped() ISoftModelNewDo
	Create(values ...*model.SoftModelNew) error
	CreateInBatches(values []*model.SoftModelNew, batchSize int) error
	Save(values ...*model.SoftModelNew) error
	First() (*model.SoftModelNew, error)
	Take() (*model.SoftModelNew, error)
	Last() (*model.SoftModelNew, error)
	Find() ([]*model.SoftModelNew, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SoftModelNew, err error)
	FindInBatches(result *[]*model.SoftModelNew, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SoftModelNew) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISoftModelNewDo
	Assign(attrs ...field.AssignExpr) ISoftModelNewDo
	Joins(fields ...field.RelationField) ISoftModelNewDo
	Preload(fields ...field.RelationField) ISoftModelNewDo
	FirstOrInit() (*model.SoftModelNew, error)
	FirstOrCreate() (*model.SoftModelNew, error)
	FindByPage(offset int, limit int) (result []*model.SoftModelNew, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISoftModelNewDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s softModelNewDo) Debug() ISoftModelNewDo {
	return s.withDO(s.DO.Debug())
}

func (s softModelNewDo) WithContext(ctx context.Context) ISoftModelNewDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s softModelNewDo) ReadDB() ISoftModelNewDo {
	return s.Clauses(dbresolver.Read)
}

func (s softModelNewDo) WriteDB() ISoftModelNewDo {
	return s.Clauses(dbresolver.Write)
}

func (s softModelNewDo) Session(config *gorm.Session) ISoftModelNewDo {
	return s.withDO(s.DO.Session(config))
}

func (s softModelNewDo) Clauses(conds ...clause.Expression) ISoftModelNewDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s softModelNewDo) Returning(value interface{}, columns ...string) ISoftModelNewDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s softModelNewDo) Not(conds ...gen.Condition) ISoftModelNewDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s softModelNewDo) Or(conds ...gen.Condition) ISoftModelNewDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s softModelNewDo) Select(conds ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s softModelNewDo) Where(conds ...gen.Condition) ISoftModelNewDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s softModelNewDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISoftModelNewDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s softModelNewDo) Order(conds ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s softModelNewDo) Distinct(cols ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s softModelNewDo) Omit(cols ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s softModelNewDo) Join(table schema.Tabler, on ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s softModelNewDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s softModelNewDo) RightJoin(table schema.Tabler, on ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s softModelNewDo) Group(cols ...field.Expr) ISoftModelNewDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s softModelNewDo) Having(conds ...gen.Condition) ISoftModelNewDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s softModelNewDo) Limit(limit int) ISoftModelNewDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s softModelNewDo) Offset(offset int) ISoftModelNewDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s softModelNewDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISoftModelNewDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s softModelNewDo) Unscoped() ISoftModelNewDo {
	return s.withDO(s.DO.Unscoped())
}

func (s softModelNewDo) Create(values ...*model.SoftModelNew) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s softModelNewDo) CreateInBatches(values []*model.SoftModelNew, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s softModelNewDo) Save(values ...*model.SoftModelNew) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s softModelNewDo) First() (*model.SoftModelNew, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SoftModelNew), nil
	}
}

func (s softModelNewDo) Take() (*model.SoftModelNew, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SoftModelNew), nil
	}
}

func (s softModelNewDo) Last() (*model.SoftModelNew, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SoftModelNew), nil
	}
}

func (s softModelNewDo) Find() ([]*model.SoftModelNew, error) {
	result, err := s.DO.Find()
	return result.([]*model.SoftModelNew), err
}

func (s softModelNewDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SoftModelNew, err error) {
	buf := make([]*model.SoftModelNew, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s softModelNewDo) FindInBatches(result *[]*model.SoftModelNew, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s softModelNewDo) Attrs(attrs ...field.AssignExpr) ISoftModelNewDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s softModelNewDo) Assign(attrs ...field.AssignExpr) ISoftModelNewDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s softModelNewDo) Joins(fields ...field.RelationField) ISoftModelNewDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s softModelNewDo) Preload(fields ...field.RelationField) ISoftModelNewDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s softModelNewDo) FirstOrInit() (*model.SoftModelNew, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SoftModelNew), nil
	}
}

func (s softModelNewDo) FirstOrCreate() (*model.SoftModelNew, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SoftModelNew), nil
	}
}

func (s softModelNewDo) FindByPage(offset int, limit int) (result []*model.SoftModelNew, count int64, err error) {
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

func (s softModelNewDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s softModelNewDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s softModelNewDo) Delete(models ...*model.SoftModelNew) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *softModelNewDo) withDO(do gen.Dao) *softModelNewDo {
	s.DO = *do.(*gen.DO)
	return s
}
