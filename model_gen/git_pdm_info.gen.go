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

func newGitPdmInfo(db *gorm.DB, opts ...gen.DOOption) gitPdmInfo {
	_gitPdmInfo := gitPdmInfo{}

	_gitPdmInfo.gitPdmInfoDo.UseDB(db, opts...)
	_gitPdmInfo.gitPdmInfoDo.UseModel(&model.GitPdmInfo{})

	tableName := _gitPdmInfo.gitPdmInfoDo.TableName()
	_gitPdmInfo.ALL = field.NewAsterisk(tableName)
	_gitPdmInfo.GitPdmInfoID = field.NewString(tableName, "git_pdm_info_id")
	_gitPdmInfo.ProjectID = field.NewString(tableName, "project_id")
	_gitPdmInfo.ProjectName = field.NewString(tableName, "project_name")
	_gitPdmInfo.PdmPath = field.NewString(tableName, "pdm_path")
	_gitPdmInfo.Branch = field.NewString(tableName, "branch")
	_gitPdmInfo.Comment = field.NewString(tableName, "comment")
	_gitPdmInfo.CreateTime = field.NewString(tableName, "create_time")

	_gitPdmInfo.fillFieldMap()

	return _gitPdmInfo
}

type gitPdmInfo struct {
	gitPdmInfoDo

	ALL          field.Asterisk
	GitPdmInfoID field.String
	ProjectID    field.String
	ProjectName  field.String
	PdmPath      field.String
	Branch       field.String
	Comment      field.String
	CreateTime   field.String

	fieldMap map[string]field.Expr
}

func (g gitPdmInfo) Table(newTableName string) *gitPdmInfo {
	g.gitPdmInfoDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gitPdmInfo) As(alias string) *gitPdmInfo {
	g.gitPdmInfoDo.DO = *(g.gitPdmInfoDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gitPdmInfo) updateTableName(table string) *gitPdmInfo {
	g.ALL = field.NewAsterisk(table)
	g.GitPdmInfoID = field.NewString(table, "git_pdm_info_id")
	g.ProjectID = field.NewString(table, "project_id")
	g.ProjectName = field.NewString(table, "project_name")
	g.PdmPath = field.NewString(table, "pdm_path")
	g.Branch = field.NewString(table, "branch")
	g.Comment = field.NewString(table, "comment")
	g.CreateTime = field.NewString(table, "create_time")

	g.fillFieldMap()

	return g
}

func (g *gitPdmInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gitPdmInfo) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["git_pdm_info_id"] = g.GitPdmInfoID
	g.fieldMap["project_id"] = g.ProjectID
	g.fieldMap["project_name"] = g.ProjectName
	g.fieldMap["pdm_path"] = g.PdmPath
	g.fieldMap["branch"] = g.Branch
	g.fieldMap["comment"] = g.Comment
	g.fieldMap["create_time"] = g.CreateTime
}

func (g gitPdmInfo) clone(db *gorm.DB) gitPdmInfo {
	g.gitPdmInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gitPdmInfo) replaceDB(db *gorm.DB) gitPdmInfo {
	g.gitPdmInfoDo.ReplaceDB(db)
	return g
}

type gitPdmInfoDo struct{ gen.DO }

type IGitPdmInfoDo interface {
	gen.SubQuery
	Debug() IGitPdmInfoDo
	WithContext(ctx context.Context) IGitPdmInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGitPdmInfoDo
	WriteDB() IGitPdmInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGitPdmInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGitPdmInfoDo
	Not(conds ...gen.Condition) IGitPdmInfoDo
	Or(conds ...gen.Condition) IGitPdmInfoDo
	Select(conds ...field.Expr) IGitPdmInfoDo
	Where(conds ...gen.Condition) IGitPdmInfoDo
	Order(conds ...field.Expr) IGitPdmInfoDo
	Distinct(cols ...field.Expr) IGitPdmInfoDo
	Omit(cols ...field.Expr) IGitPdmInfoDo
	Join(table schema.Tabler, on ...field.Expr) IGitPdmInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGitPdmInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGitPdmInfoDo
	Group(cols ...field.Expr) IGitPdmInfoDo
	Having(conds ...gen.Condition) IGitPdmInfoDo
	Limit(limit int) IGitPdmInfoDo
	Offset(offset int) IGitPdmInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGitPdmInfoDo
	Unscoped() IGitPdmInfoDo
	Create(values ...*model.GitPdmInfo) error
	CreateInBatches(values []*model.GitPdmInfo, batchSize int) error
	Save(values ...*model.GitPdmInfo) error
	First() (*model.GitPdmInfo, error)
	Take() (*model.GitPdmInfo, error)
	Last() (*model.GitPdmInfo, error)
	Find() ([]*model.GitPdmInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GitPdmInfo, err error)
	FindInBatches(result *[]*model.GitPdmInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GitPdmInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGitPdmInfoDo
	Assign(attrs ...field.AssignExpr) IGitPdmInfoDo
	Joins(fields ...field.RelationField) IGitPdmInfoDo
	Preload(fields ...field.RelationField) IGitPdmInfoDo
	FirstOrInit() (*model.GitPdmInfo, error)
	FirstOrCreate() (*model.GitPdmInfo, error)
	FindByPage(offset int, limit int) (result []*model.GitPdmInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGitPdmInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gitPdmInfoDo) Debug() IGitPdmInfoDo {
	return g.withDO(g.DO.Debug())
}

func (g gitPdmInfoDo) WithContext(ctx context.Context) IGitPdmInfoDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gitPdmInfoDo) ReadDB() IGitPdmInfoDo {
	return g.Clauses(dbresolver.Read)
}

func (g gitPdmInfoDo) WriteDB() IGitPdmInfoDo {
	return g.Clauses(dbresolver.Write)
}

func (g gitPdmInfoDo) Session(config *gorm.Session) IGitPdmInfoDo {
	return g.withDO(g.DO.Session(config))
}

func (g gitPdmInfoDo) Clauses(conds ...clause.Expression) IGitPdmInfoDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gitPdmInfoDo) Returning(value interface{}, columns ...string) IGitPdmInfoDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gitPdmInfoDo) Not(conds ...gen.Condition) IGitPdmInfoDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gitPdmInfoDo) Or(conds ...gen.Condition) IGitPdmInfoDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gitPdmInfoDo) Select(conds ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gitPdmInfoDo) Where(conds ...gen.Condition) IGitPdmInfoDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gitPdmInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IGitPdmInfoDo {
	return g.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (g gitPdmInfoDo) Order(conds ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gitPdmInfoDo) Distinct(cols ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gitPdmInfoDo) Omit(cols ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gitPdmInfoDo) Join(table schema.Tabler, on ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gitPdmInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gitPdmInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gitPdmInfoDo) Group(cols ...field.Expr) IGitPdmInfoDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gitPdmInfoDo) Having(conds ...gen.Condition) IGitPdmInfoDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gitPdmInfoDo) Limit(limit int) IGitPdmInfoDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gitPdmInfoDo) Offset(offset int) IGitPdmInfoDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gitPdmInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGitPdmInfoDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gitPdmInfoDo) Unscoped() IGitPdmInfoDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gitPdmInfoDo) Create(values ...*model.GitPdmInfo) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gitPdmInfoDo) CreateInBatches(values []*model.GitPdmInfo, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gitPdmInfoDo) Save(values ...*model.GitPdmInfo) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gitPdmInfoDo) First() (*model.GitPdmInfo, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GitPdmInfo), nil
	}
}

func (g gitPdmInfoDo) Take() (*model.GitPdmInfo, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GitPdmInfo), nil
	}
}

func (g gitPdmInfoDo) Last() (*model.GitPdmInfo, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GitPdmInfo), nil
	}
}

func (g gitPdmInfoDo) Find() ([]*model.GitPdmInfo, error) {
	result, err := g.DO.Find()
	return result.([]*model.GitPdmInfo), err
}

func (g gitPdmInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GitPdmInfo, err error) {
	buf := make([]*model.GitPdmInfo, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gitPdmInfoDo) FindInBatches(result *[]*model.GitPdmInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gitPdmInfoDo) Attrs(attrs ...field.AssignExpr) IGitPdmInfoDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gitPdmInfoDo) Assign(attrs ...field.AssignExpr) IGitPdmInfoDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gitPdmInfoDo) Joins(fields ...field.RelationField) IGitPdmInfoDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gitPdmInfoDo) Preload(fields ...field.RelationField) IGitPdmInfoDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gitPdmInfoDo) FirstOrInit() (*model.GitPdmInfo, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GitPdmInfo), nil
	}
}

func (g gitPdmInfoDo) FirstOrCreate() (*model.GitPdmInfo, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GitPdmInfo), nil
	}
}

func (g gitPdmInfoDo) FindByPage(offset int, limit int) (result []*model.GitPdmInfo, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gitPdmInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gitPdmInfoDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gitPdmInfoDo) Delete(models ...*model.GitPdmInfo) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gitPdmInfoDo) withDO(do gen.Dao) *gitPdmInfoDo {
	g.DO = *do.(*gen.DO)
	return g
}
