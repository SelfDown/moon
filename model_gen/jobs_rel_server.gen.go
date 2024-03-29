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

func newJobsRelServer(db *gorm.DB, opts ...gen.DOOption) jobsRelServer {
	_jobsRelServer := jobsRelServer{}

	_jobsRelServer.jobsRelServerDo.UseDB(db, opts...)
	_jobsRelServer.jobsRelServerDo.UseModel(&model.JobsRelServer{})

	tableName := _jobsRelServer.jobsRelServerDo.TableName()
	_jobsRelServer.ALL = field.NewAsterisk(tableName)
	_jobsRelServer.JobsRelServerID = field.NewString(tableName, "jobs_rel_server_id")
	_jobsRelServer.ServerID = field.NewString(tableName, "server_id")
	_jobsRelServer.JobID = field.NewString(tableName, "job_id")
	_jobsRelServer.CreateTime = field.NewTime(tableName, "create_time")
	_jobsRelServer.ModifyTime = field.NewTime(tableName, "modify_time")
	_jobsRelServer.Comments = field.NewString(tableName, "comments")

	_jobsRelServer.fillFieldMap()

	return _jobsRelServer
}

type jobsRelServer struct {
	jobsRelServerDo

	ALL             field.Asterisk
	JobsRelServerID field.String // uuid
	ServerID        field.String
	JobID           field.String // 任务编号
	CreateTime      field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime      field.Time   // 记录修改时间（数据库自动写入）
	Comments        field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (j jobsRelServer) Table(newTableName string) *jobsRelServer {
	j.jobsRelServerDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jobsRelServer) As(alias string) *jobsRelServer {
	j.jobsRelServerDo.DO = *(j.jobsRelServerDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jobsRelServer) updateTableName(table string) *jobsRelServer {
	j.ALL = field.NewAsterisk(table)
	j.JobsRelServerID = field.NewString(table, "jobs_rel_server_id")
	j.ServerID = field.NewString(table, "server_id")
	j.JobID = field.NewString(table, "job_id")
	j.CreateTime = field.NewTime(table, "create_time")
	j.ModifyTime = field.NewTime(table, "modify_time")
	j.Comments = field.NewString(table, "comments")

	j.fillFieldMap()

	return j
}

func (j *jobsRelServer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jobsRelServer) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 6)
	j.fieldMap["jobs_rel_server_id"] = j.JobsRelServerID
	j.fieldMap["server_id"] = j.ServerID
	j.fieldMap["job_id"] = j.JobID
	j.fieldMap["create_time"] = j.CreateTime
	j.fieldMap["modify_time"] = j.ModifyTime
	j.fieldMap["comments"] = j.Comments
}

func (j jobsRelServer) clone(db *gorm.DB) jobsRelServer {
	j.jobsRelServerDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jobsRelServer) replaceDB(db *gorm.DB) jobsRelServer {
	j.jobsRelServerDo.ReplaceDB(db)
	return j
}

type jobsRelServerDo struct{ gen.DO }

type IJobsRelServerDo interface {
	gen.SubQuery
	Debug() IJobsRelServerDo
	WithContext(ctx context.Context) IJobsRelServerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJobsRelServerDo
	WriteDB() IJobsRelServerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJobsRelServerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJobsRelServerDo
	Not(conds ...gen.Condition) IJobsRelServerDo
	Or(conds ...gen.Condition) IJobsRelServerDo
	Select(conds ...field.Expr) IJobsRelServerDo
	Where(conds ...gen.Condition) IJobsRelServerDo
	Order(conds ...field.Expr) IJobsRelServerDo
	Distinct(cols ...field.Expr) IJobsRelServerDo
	Omit(cols ...field.Expr) IJobsRelServerDo
	Join(table schema.Tabler, on ...field.Expr) IJobsRelServerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJobsRelServerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJobsRelServerDo
	Group(cols ...field.Expr) IJobsRelServerDo
	Having(conds ...gen.Condition) IJobsRelServerDo
	Limit(limit int) IJobsRelServerDo
	Offset(offset int) IJobsRelServerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJobsRelServerDo
	Unscoped() IJobsRelServerDo
	Create(values ...*model.JobsRelServer) error
	CreateInBatches(values []*model.JobsRelServer, batchSize int) error
	Save(values ...*model.JobsRelServer) error
	First() (*model.JobsRelServer, error)
	Take() (*model.JobsRelServer, error)
	Last() (*model.JobsRelServer, error)
	Find() ([]*model.JobsRelServer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JobsRelServer, err error)
	FindInBatches(result *[]*model.JobsRelServer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.JobsRelServer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJobsRelServerDo
	Assign(attrs ...field.AssignExpr) IJobsRelServerDo
	Joins(fields ...field.RelationField) IJobsRelServerDo
	Preload(fields ...field.RelationField) IJobsRelServerDo
	FirstOrInit() (*model.JobsRelServer, error)
	FirstOrCreate() (*model.JobsRelServer, error)
	FindByPage(offset int, limit int) (result []*model.JobsRelServer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJobsRelServerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jobsRelServerDo) Debug() IJobsRelServerDo {
	return j.withDO(j.DO.Debug())
}

func (j jobsRelServerDo) WithContext(ctx context.Context) IJobsRelServerDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobsRelServerDo) ReadDB() IJobsRelServerDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobsRelServerDo) WriteDB() IJobsRelServerDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobsRelServerDo) Session(config *gorm.Session) IJobsRelServerDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobsRelServerDo) Clauses(conds ...clause.Expression) IJobsRelServerDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobsRelServerDo) Returning(value interface{}, columns ...string) IJobsRelServerDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobsRelServerDo) Not(conds ...gen.Condition) IJobsRelServerDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobsRelServerDo) Or(conds ...gen.Condition) IJobsRelServerDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobsRelServerDo) Select(conds ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobsRelServerDo) Where(conds ...gen.Condition) IJobsRelServerDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobsRelServerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IJobsRelServerDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j jobsRelServerDo) Order(conds ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobsRelServerDo) Distinct(cols ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobsRelServerDo) Omit(cols ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobsRelServerDo) Join(table schema.Tabler, on ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobsRelServerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobsRelServerDo) RightJoin(table schema.Tabler, on ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobsRelServerDo) Group(cols ...field.Expr) IJobsRelServerDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobsRelServerDo) Having(conds ...gen.Condition) IJobsRelServerDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobsRelServerDo) Limit(limit int) IJobsRelServerDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobsRelServerDo) Offset(offset int) IJobsRelServerDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobsRelServerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJobsRelServerDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobsRelServerDo) Unscoped() IJobsRelServerDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobsRelServerDo) Create(values ...*model.JobsRelServer) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobsRelServerDo) CreateInBatches(values []*model.JobsRelServer, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobsRelServerDo) Save(values ...*model.JobsRelServer) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobsRelServerDo) First() (*model.JobsRelServer, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsRelServer), nil
	}
}

func (j jobsRelServerDo) Take() (*model.JobsRelServer, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsRelServer), nil
	}
}

func (j jobsRelServerDo) Last() (*model.JobsRelServer, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsRelServer), nil
	}
}

func (j jobsRelServerDo) Find() ([]*model.JobsRelServer, error) {
	result, err := j.DO.Find()
	return result.([]*model.JobsRelServer), err
}

func (j jobsRelServerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JobsRelServer, err error) {
	buf := make([]*model.JobsRelServer, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobsRelServerDo) FindInBatches(result *[]*model.JobsRelServer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobsRelServerDo) Attrs(attrs ...field.AssignExpr) IJobsRelServerDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobsRelServerDo) Assign(attrs ...field.AssignExpr) IJobsRelServerDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobsRelServerDo) Joins(fields ...field.RelationField) IJobsRelServerDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobsRelServerDo) Preload(fields ...field.RelationField) IJobsRelServerDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobsRelServerDo) FirstOrInit() (*model.JobsRelServer, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsRelServer), nil
	}
}

func (j jobsRelServerDo) FirstOrCreate() (*model.JobsRelServer, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsRelServer), nil
	}
}

func (j jobsRelServerDo) FindByPage(offset int, limit int) (result []*model.JobsRelServer, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j jobsRelServerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobsRelServerDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobsRelServerDo) Delete(models ...*model.JobsRelServer) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobsRelServerDo) withDO(do gen.Dao) *jobsRelServerDo {
	j.DO = *do.(*gen.DO)
	return j
}
