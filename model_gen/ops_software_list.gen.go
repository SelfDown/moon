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

func newOpsSoftwareList(db *gorm.DB, opts ...gen.DOOption) opsSoftwareList {
	_opsSoftwareList := opsSoftwareList{}

	_opsSoftwareList.opsSoftwareListDo.UseDB(db, opts...)
	_opsSoftwareList.opsSoftwareListDo.UseModel(&model.OpsSoftwareList{})

	tableName := _opsSoftwareList.opsSoftwareListDo.TableName()
	_opsSoftwareList.ALL = field.NewAsterisk(tableName)
	_opsSoftwareList.OpsSoftwareListID = field.NewString(tableName, "ops_software_list_id")
	_opsSoftwareList.OpsSoftwareNexusRepository = field.NewString(tableName, "ops_software_nexus_repository")
	_opsSoftwareList.OpsSoftwareArtifactID = field.NewString(tableName, "ops_software_artifact_id")
	_opsSoftwareList.OpsSoftwareGroupID = field.NewString(tableName, "ops_software_group_id")
	_opsSoftwareList.CreateTime = field.NewTime(tableName, "create_time")
	_opsSoftwareList.ModifyTime = field.NewTime(tableName, "modify_time")
	_opsSoftwareList.Comments = field.NewString(tableName, "comments")
	_opsSoftwareList.SoftName = field.NewString(tableName, "soft_name")
	_opsSoftwareList.SoftGitPath = field.NewString(tableName, "soft_git_path")

	_opsSoftwareList.fillFieldMap()

	return _opsSoftwareList
}

type opsSoftwareList struct {
	opsSoftwareListDo

	ALL                        field.Asterisk
	OpsSoftwareListID          field.String // 主键
	OpsSoftwareNexusRepository field.String
	OpsSoftwareArtifactID      field.String // nexus 坐标
	OpsSoftwareGroupID         field.String // nexus 坐标
	CreateTime                 field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime                 field.Time   // 记录修改时间（数据库自动写入）
	Comments                   field.String // 备注说明
	SoftName                   field.String
	SoftGitPath                field.String

	fieldMap map[string]field.Expr
}

func (o opsSoftwareList) Table(newTableName string) *opsSoftwareList {
	o.opsSoftwareListDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o opsSoftwareList) As(alias string) *opsSoftwareList {
	o.opsSoftwareListDo.DO = *(o.opsSoftwareListDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *opsSoftwareList) updateTableName(table string) *opsSoftwareList {
	o.ALL = field.NewAsterisk(table)
	o.OpsSoftwareListID = field.NewString(table, "ops_software_list_id")
	o.OpsSoftwareNexusRepository = field.NewString(table, "ops_software_nexus_repository")
	o.OpsSoftwareArtifactID = field.NewString(table, "ops_software_artifact_id")
	o.OpsSoftwareGroupID = field.NewString(table, "ops_software_group_id")
	o.CreateTime = field.NewTime(table, "create_time")
	o.ModifyTime = field.NewTime(table, "modify_time")
	o.Comments = field.NewString(table, "comments")
	o.SoftName = field.NewString(table, "soft_name")
	o.SoftGitPath = field.NewString(table, "soft_git_path")

	o.fillFieldMap()

	return o
}

func (o *opsSoftwareList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *opsSoftwareList) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 9)
	o.fieldMap["ops_software_list_id"] = o.OpsSoftwareListID
	o.fieldMap["ops_software_nexus_repository"] = o.OpsSoftwareNexusRepository
	o.fieldMap["ops_software_artifact_id"] = o.OpsSoftwareArtifactID
	o.fieldMap["ops_software_group_id"] = o.OpsSoftwareGroupID
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["modify_time"] = o.ModifyTime
	o.fieldMap["comments"] = o.Comments
	o.fieldMap["soft_name"] = o.SoftName
	o.fieldMap["soft_git_path"] = o.SoftGitPath
}

func (o opsSoftwareList) clone(db *gorm.DB) opsSoftwareList {
	o.opsSoftwareListDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o opsSoftwareList) replaceDB(db *gorm.DB) opsSoftwareList {
	o.opsSoftwareListDo.ReplaceDB(db)
	return o
}

type opsSoftwareListDo struct{ gen.DO }

type IOpsSoftwareListDo interface {
	gen.SubQuery
	Debug() IOpsSoftwareListDo
	WithContext(ctx context.Context) IOpsSoftwareListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOpsSoftwareListDo
	WriteDB() IOpsSoftwareListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOpsSoftwareListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOpsSoftwareListDo
	Not(conds ...gen.Condition) IOpsSoftwareListDo
	Or(conds ...gen.Condition) IOpsSoftwareListDo
	Select(conds ...field.Expr) IOpsSoftwareListDo
	Where(conds ...gen.Condition) IOpsSoftwareListDo
	Order(conds ...field.Expr) IOpsSoftwareListDo
	Distinct(cols ...field.Expr) IOpsSoftwareListDo
	Omit(cols ...field.Expr) IOpsSoftwareListDo
	Join(table schema.Tabler, on ...field.Expr) IOpsSoftwareListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOpsSoftwareListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOpsSoftwareListDo
	Group(cols ...field.Expr) IOpsSoftwareListDo
	Having(conds ...gen.Condition) IOpsSoftwareListDo
	Limit(limit int) IOpsSoftwareListDo
	Offset(offset int) IOpsSoftwareListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOpsSoftwareListDo
	Unscoped() IOpsSoftwareListDo
	Create(values ...*model.OpsSoftwareList) error
	CreateInBatches(values []*model.OpsSoftwareList, batchSize int) error
	Save(values ...*model.OpsSoftwareList) error
	First() (*model.OpsSoftwareList, error)
	Take() (*model.OpsSoftwareList, error)
	Last() (*model.OpsSoftwareList, error)
	Find() ([]*model.OpsSoftwareList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OpsSoftwareList, err error)
	FindInBatches(result *[]*model.OpsSoftwareList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OpsSoftwareList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOpsSoftwareListDo
	Assign(attrs ...field.AssignExpr) IOpsSoftwareListDo
	Joins(fields ...field.RelationField) IOpsSoftwareListDo
	Preload(fields ...field.RelationField) IOpsSoftwareListDo
	FirstOrInit() (*model.OpsSoftwareList, error)
	FirstOrCreate() (*model.OpsSoftwareList, error)
	FindByPage(offset int, limit int) (result []*model.OpsSoftwareList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOpsSoftwareListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o opsSoftwareListDo) Debug() IOpsSoftwareListDo {
	return o.withDO(o.DO.Debug())
}

func (o opsSoftwareListDo) WithContext(ctx context.Context) IOpsSoftwareListDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o opsSoftwareListDo) ReadDB() IOpsSoftwareListDo {
	return o.Clauses(dbresolver.Read)
}

func (o opsSoftwareListDo) WriteDB() IOpsSoftwareListDo {
	return o.Clauses(dbresolver.Write)
}

func (o opsSoftwareListDo) Session(config *gorm.Session) IOpsSoftwareListDo {
	return o.withDO(o.DO.Session(config))
}

func (o opsSoftwareListDo) Clauses(conds ...clause.Expression) IOpsSoftwareListDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o opsSoftwareListDo) Returning(value interface{}, columns ...string) IOpsSoftwareListDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o opsSoftwareListDo) Not(conds ...gen.Condition) IOpsSoftwareListDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o opsSoftwareListDo) Or(conds ...gen.Condition) IOpsSoftwareListDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o opsSoftwareListDo) Select(conds ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o opsSoftwareListDo) Where(conds ...gen.Condition) IOpsSoftwareListDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o opsSoftwareListDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IOpsSoftwareListDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o opsSoftwareListDo) Order(conds ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o opsSoftwareListDo) Distinct(cols ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o opsSoftwareListDo) Omit(cols ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o opsSoftwareListDo) Join(table schema.Tabler, on ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o opsSoftwareListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o opsSoftwareListDo) RightJoin(table schema.Tabler, on ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o opsSoftwareListDo) Group(cols ...field.Expr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o opsSoftwareListDo) Having(conds ...gen.Condition) IOpsSoftwareListDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o opsSoftwareListDo) Limit(limit int) IOpsSoftwareListDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o opsSoftwareListDo) Offset(offset int) IOpsSoftwareListDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o opsSoftwareListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOpsSoftwareListDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o opsSoftwareListDo) Unscoped() IOpsSoftwareListDo {
	return o.withDO(o.DO.Unscoped())
}

func (o opsSoftwareListDo) Create(values ...*model.OpsSoftwareList) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o opsSoftwareListDo) CreateInBatches(values []*model.OpsSoftwareList, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o opsSoftwareListDo) Save(values ...*model.OpsSoftwareList) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o opsSoftwareListDo) First() (*model.OpsSoftwareList, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpsSoftwareList), nil
	}
}

func (o opsSoftwareListDo) Take() (*model.OpsSoftwareList, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpsSoftwareList), nil
	}
}

func (o opsSoftwareListDo) Last() (*model.OpsSoftwareList, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpsSoftwareList), nil
	}
}

func (o opsSoftwareListDo) Find() ([]*model.OpsSoftwareList, error) {
	result, err := o.DO.Find()
	return result.([]*model.OpsSoftwareList), err
}

func (o opsSoftwareListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OpsSoftwareList, err error) {
	buf := make([]*model.OpsSoftwareList, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o opsSoftwareListDo) FindInBatches(result *[]*model.OpsSoftwareList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o opsSoftwareListDo) Attrs(attrs ...field.AssignExpr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o opsSoftwareListDo) Assign(attrs ...field.AssignExpr) IOpsSoftwareListDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o opsSoftwareListDo) Joins(fields ...field.RelationField) IOpsSoftwareListDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o opsSoftwareListDo) Preload(fields ...field.RelationField) IOpsSoftwareListDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o opsSoftwareListDo) FirstOrInit() (*model.OpsSoftwareList, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpsSoftwareList), nil
	}
}

func (o opsSoftwareListDo) FirstOrCreate() (*model.OpsSoftwareList, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OpsSoftwareList), nil
	}
}

func (o opsSoftwareListDo) FindByPage(offset int, limit int) (result []*model.OpsSoftwareList, count int64, err error) {
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

func (o opsSoftwareListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o opsSoftwareListDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o opsSoftwareListDo) Delete(models ...*model.OpsSoftwareList) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *opsSoftwareListDo) withDO(do gen.Dao) *opsSoftwareListDo {
	o.DO = *do.(*gen.DO)
	return o
}
