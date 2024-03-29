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

func newOracleAwr(db *gorm.DB, opts ...gen.DOOption) oracleAwr {
	_oracleAwr := oracleAwr{}

	_oracleAwr.oracleAwrDo.UseDB(db, opts...)
	_oracleAwr.oracleAwrDo.UseModel(&model.OracleAwr{})

	tableName := _oracleAwr.oracleAwrDo.TableName()
	_oracleAwr.ALL = field.NewAsterisk(tableName)
	_oracleAwr.OracleAwrID = field.NewString(tableName, "oracle_awr_id")
	_oracleAwr.OralceInstanceID = field.NewString(tableName, "oralce_instance_id")
	_oracleAwr.CreateTime = field.NewTime(tableName, "create_time")
	_oracleAwr.ModifyTime = field.NewTime(tableName, "modify_time")
	_oracleAwr.UserID = field.NewString(tableName, "user_id")
	_oracleAwr.Comments = field.NewString(tableName, "comments")
	_oracleAwr.ReportHTMLContent = field.NewString(tableName, "report_html_content")
	_oracleAwr.SanpStartTime = field.NewString(tableName, "sanp_start_time")
	_oracleAwr.SanpEndTime = field.NewString(tableName, "sanp_end_time")

	_oracleAwr.fillFieldMap()

	return _oracleAwr
}

type oracleAwr struct {
	oracleAwrDo

	ALL               field.Asterisk
	OracleAwrID       field.String
	OralceInstanceID  field.String // 关联oracle_instance
	CreateTime        field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime        field.Time
	UserID            field.String
	Comments          field.String
	ReportHTMLContent field.String // AWR报告HTML内容
	SanpStartTime     field.String
	SanpEndTime       field.String

	fieldMap map[string]field.Expr
}

func (o oracleAwr) Table(newTableName string) *oracleAwr {
	o.oracleAwrDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o oracleAwr) As(alias string) *oracleAwr {
	o.oracleAwrDo.DO = *(o.oracleAwrDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *oracleAwr) updateTableName(table string) *oracleAwr {
	o.ALL = field.NewAsterisk(table)
	o.OracleAwrID = field.NewString(table, "oracle_awr_id")
	o.OralceInstanceID = field.NewString(table, "oralce_instance_id")
	o.CreateTime = field.NewTime(table, "create_time")
	o.ModifyTime = field.NewTime(table, "modify_time")
	o.UserID = field.NewString(table, "user_id")
	o.Comments = field.NewString(table, "comments")
	o.ReportHTMLContent = field.NewString(table, "report_html_content")
	o.SanpStartTime = field.NewString(table, "sanp_start_time")
	o.SanpEndTime = field.NewString(table, "sanp_end_time")

	o.fillFieldMap()

	return o
}

func (o *oracleAwr) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *oracleAwr) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 9)
	o.fieldMap["oracle_awr_id"] = o.OracleAwrID
	o.fieldMap["oralce_instance_id"] = o.OralceInstanceID
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["modify_time"] = o.ModifyTime
	o.fieldMap["user_id"] = o.UserID
	o.fieldMap["comments"] = o.Comments
	o.fieldMap["report_html_content"] = o.ReportHTMLContent
	o.fieldMap["sanp_start_time"] = o.SanpStartTime
	o.fieldMap["sanp_end_time"] = o.SanpEndTime
}

func (o oracleAwr) clone(db *gorm.DB) oracleAwr {
	o.oracleAwrDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o oracleAwr) replaceDB(db *gorm.DB) oracleAwr {
	o.oracleAwrDo.ReplaceDB(db)
	return o
}

type oracleAwrDo struct{ gen.DO }

type IOracleAwrDo interface {
	gen.SubQuery
	Debug() IOracleAwrDo
	WithContext(ctx context.Context) IOracleAwrDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOracleAwrDo
	WriteDB() IOracleAwrDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOracleAwrDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOracleAwrDo
	Not(conds ...gen.Condition) IOracleAwrDo
	Or(conds ...gen.Condition) IOracleAwrDo
	Select(conds ...field.Expr) IOracleAwrDo
	Where(conds ...gen.Condition) IOracleAwrDo
	Order(conds ...field.Expr) IOracleAwrDo
	Distinct(cols ...field.Expr) IOracleAwrDo
	Omit(cols ...field.Expr) IOracleAwrDo
	Join(table schema.Tabler, on ...field.Expr) IOracleAwrDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOracleAwrDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOracleAwrDo
	Group(cols ...field.Expr) IOracleAwrDo
	Having(conds ...gen.Condition) IOracleAwrDo
	Limit(limit int) IOracleAwrDo
	Offset(offset int) IOracleAwrDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOracleAwrDo
	Unscoped() IOracleAwrDo
	Create(values ...*model.OracleAwr) error
	CreateInBatches(values []*model.OracleAwr, batchSize int) error
	Save(values ...*model.OracleAwr) error
	First() (*model.OracleAwr, error)
	Take() (*model.OracleAwr, error)
	Last() (*model.OracleAwr, error)
	Find() ([]*model.OracleAwr, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OracleAwr, err error)
	FindInBatches(result *[]*model.OracleAwr, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OracleAwr) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOracleAwrDo
	Assign(attrs ...field.AssignExpr) IOracleAwrDo
	Joins(fields ...field.RelationField) IOracleAwrDo
	Preload(fields ...field.RelationField) IOracleAwrDo
	FirstOrInit() (*model.OracleAwr, error)
	FirstOrCreate() (*model.OracleAwr, error)
	FindByPage(offset int, limit int) (result []*model.OracleAwr, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOracleAwrDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o oracleAwrDo) Debug() IOracleAwrDo {
	return o.withDO(o.DO.Debug())
}

func (o oracleAwrDo) WithContext(ctx context.Context) IOracleAwrDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o oracleAwrDo) ReadDB() IOracleAwrDo {
	return o.Clauses(dbresolver.Read)
}

func (o oracleAwrDo) WriteDB() IOracleAwrDo {
	return o.Clauses(dbresolver.Write)
}

func (o oracleAwrDo) Session(config *gorm.Session) IOracleAwrDo {
	return o.withDO(o.DO.Session(config))
}

func (o oracleAwrDo) Clauses(conds ...clause.Expression) IOracleAwrDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o oracleAwrDo) Returning(value interface{}, columns ...string) IOracleAwrDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o oracleAwrDo) Not(conds ...gen.Condition) IOracleAwrDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o oracleAwrDo) Or(conds ...gen.Condition) IOracleAwrDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o oracleAwrDo) Select(conds ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o oracleAwrDo) Where(conds ...gen.Condition) IOracleAwrDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o oracleAwrDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOracleAwrDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o oracleAwrDo) Order(conds ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o oracleAwrDo) Distinct(cols ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o oracleAwrDo) Omit(cols ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o oracleAwrDo) Join(table schema.Tabler, on ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o oracleAwrDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o oracleAwrDo) RightJoin(table schema.Tabler, on ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o oracleAwrDo) Group(cols ...field.Expr) IOracleAwrDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o oracleAwrDo) Having(conds ...gen.Condition) IOracleAwrDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o oracleAwrDo) Limit(limit int) IOracleAwrDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o oracleAwrDo) Offset(offset int) IOracleAwrDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o oracleAwrDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOracleAwrDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o oracleAwrDo) Unscoped() IOracleAwrDo {
	return o.withDO(o.DO.Unscoped())
}

func (o oracleAwrDo) Create(values ...*model.OracleAwr) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o oracleAwrDo) CreateInBatches(values []*model.OracleAwr, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o oracleAwrDo) Save(values ...*model.OracleAwr) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o oracleAwrDo) First() (*model.OracleAwr, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OracleAwr), nil
	}
}

func (o oracleAwrDo) Take() (*model.OracleAwr, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OracleAwr), nil
	}
}

func (o oracleAwrDo) Last() (*model.OracleAwr, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OracleAwr), nil
	}
}

func (o oracleAwrDo) Find() ([]*model.OracleAwr, error) {
	result, err := o.DO.Find()
	return result.([]*model.OracleAwr), err
}

func (o oracleAwrDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OracleAwr, err error) {
	buf := make([]*model.OracleAwr, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o oracleAwrDo) FindInBatches(result *[]*model.OracleAwr, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o oracleAwrDo) Attrs(attrs ...field.AssignExpr) IOracleAwrDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o oracleAwrDo) Assign(attrs ...field.AssignExpr) IOracleAwrDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o oracleAwrDo) Joins(fields ...field.RelationField) IOracleAwrDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o oracleAwrDo) Preload(fields ...field.RelationField) IOracleAwrDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o oracleAwrDo) FirstOrInit() (*model.OracleAwr, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OracleAwr), nil
	}
}

func (o oracleAwrDo) FirstOrCreate() (*model.OracleAwr, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OracleAwr), nil
	}
}

func (o oracleAwrDo) FindByPage(offset int, limit int) (result []*model.OracleAwr, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o oracleAwrDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o oracleAwrDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o oracleAwrDo) Delete(models ...*model.OracleAwr) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *oracleAwrDo) withDO(do gen.Dao) *oracleAwrDo {
	o.DO = *do.(*gen.DO)
	return o
}
