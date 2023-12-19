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

func newJobsExecLog(db *gorm.DB, opts ...gen.DOOption) jobsExecLog {
	_jobsExecLog := jobsExecLog{}

	_jobsExecLog.jobsExecLogDo.UseDB(db, opts...)
	_jobsExecLog.jobsExecLogDo.UseModel(&model.JobsExecLog{})

	tableName := _jobsExecLog.jobsExecLogDo.TableName()
	_jobsExecLog.ALL = field.NewAsterisk(tableName)
	_jobsExecLog.JobsExecLogID = field.NewString(tableName, "jobs_exec_log_id")
	_jobsExecLog.JobID = field.NewString(tableName, "job_id")
	_jobsExecLog.ExecTime = field.NewTime(tableName, "exec_time")
	_jobsExecLog.JobsExecLog = field.NewString(tableName, "jobs_exec_log")
	_jobsExecLog.ExceStatus = field.NewString(tableName, "exce_status")
	_jobsExecLog.CreateTime = field.NewTime(tableName, "create_time")
	_jobsExecLog.ModifyTime = field.NewTime(tableName, "modify_time")
	_jobsExecLog.Comments = field.NewString(tableName, "comments")

	_jobsExecLog.fillFieldMap()

	return _jobsExecLog
}

type jobsExecLog struct {
	jobsExecLogDo

	ALL           field.Asterisk
	JobsExecLogID field.String // UUID
	JobID         field.String // 任务ID
	ExecTime      field.Time
	JobsExecLog   field.String // 执行日志
	ExceStatus    field.String // 0 执行成功，1表示执行失败
	CreateTime    field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime    field.Time   // 记录修改时间（数据库自动写入）
	Comments      field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (j jobsExecLog) Table(newTableName string) *jobsExecLog {
	j.jobsExecLogDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j jobsExecLog) As(alias string) *jobsExecLog {
	j.jobsExecLogDo.DO = *(j.jobsExecLogDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *jobsExecLog) updateTableName(table string) *jobsExecLog {
	j.ALL = field.NewAsterisk(table)
	j.JobsExecLogID = field.NewString(table, "jobs_exec_log_id")
	j.JobID = field.NewString(table, "job_id")
	j.ExecTime = field.NewTime(table, "exec_time")
	j.JobsExecLog = field.NewString(table, "jobs_exec_log")
	j.ExceStatus = field.NewString(table, "exce_status")
	j.CreateTime = field.NewTime(table, "create_time")
	j.ModifyTime = field.NewTime(table, "modify_time")
	j.Comments = field.NewString(table, "comments")

	j.fillFieldMap()

	return j
}

func (j *jobsExecLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *jobsExecLog) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 8)
	j.fieldMap["jobs_exec_log_id"] = j.JobsExecLogID
	j.fieldMap["job_id"] = j.JobID
	j.fieldMap["exec_time"] = j.ExecTime
	j.fieldMap["jobs_exec_log"] = j.JobsExecLog
	j.fieldMap["exce_status"] = j.ExceStatus
	j.fieldMap["create_time"] = j.CreateTime
	j.fieldMap["modify_time"] = j.ModifyTime
	j.fieldMap["comments"] = j.Comments
}

func (j jobsExecLog) clone(db *gorm.DB) jobsExecLog {
	j.jobsExecLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j jobsExecLog) replaceDB(db *gorm.DB) jobsExecLog {
	j.jobsExecLogDo.ReplaceDB(db)
	return j
}

type jobsExecLogDo struct{ gen.DO }

type IJobsExecLogDo interface {
	gen.SubQuery
	Debug() IJobsExecLogDo
	WithContext(ctx context.Context) IJobsExecLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IJobsExecLogDo
	WriteDB() IJobsExecLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IJobsExecLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IJobsExecLogDo
	Not(conds ...gen.Condition) IJobsExecLogDo
	Or(conds ...gen.Condition) IJobsExecLogDo
	Select(conds ...field.Expr) IJobsExecLogDo
	Where(conds ...gen.Condition) IJobsExecLogDo
	Order(conds ...field.Expr) IJobsExecLogDo
	Distinct(cols ...field.Expr) IJobsExecLogDo
	Omit(cols ...field.Expr) IJobsExecLogDo
	Join(table schema.Tabler, on ...field.Expr) IJobsExecLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IJobsExecLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IJobsExecLogDo
	Group(cols ...field.Expr) IJobsExecLogDo
	Having(conds ...gen.Condition) IJobsExecLogDo
	Limit(limit int) IJobsExecLogDo
	Offset(offset int) IJobsExecLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IJobsExecLogDo
	Unscoped() IJobsExecLogDo
	Create(values ...*model.JobsExecLog) error
	CreateInBatches(values []*model.JobsExecLog, batchSize int) error
	Save(values ...*model.JobsExecLog) error
	First() (*model.JobsExecLog, error)
	Take() (*model.JobsExecLog, error)
	Last() (*model.JobsExecLog, error)
	Find() ([]*model.JobsExecLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JobsExecLog, err error)
	FindInBatches(result *[]*model.JobsExecLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.JobsExecLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IJobsExecLogDo
	Assign(attrs ...field.AssignExpr) IJobsExecLogDo
	Joins(fields ...field.RelationField) IJobsExecLogDo
	Preload(fields ...field.RelationField) IJobsExecLogDo
	FirstOrInit() (*model.JobsExecLog, error)
	FirstOrCreate() (*model.JobsExecLog, error)
	FindByPage(offset int, limit int) (result []*model.JobsExecLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IJobsExecLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (j jobsExecLogDo) Debug() IJobsExecLogDo {
	return j.withDO(j.DO.Debug())
}

func (j jobsExecLogDo) WithContext(ctx context.Context) IJobsExecLogDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j jobsExecLogDo) ReadDB() IJobsExecLogDo {
	return j.Clauses(dbresolver.Read)
}

func (j jobsExecLogDo) WriteDB() IJobsExecLogDo {
	return j.Clauses(dbresolver.Write)
}

func (j jobsExecLogDo) Session(config *gorm.Session) IJobsExecLogDo {
	return j.withDO(j.DO.Session(config))
}

func (j jobsExecLogDo) Clauses(conds ...clause.Expression) IJobsExecLogDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j jobsExecLogDo) Returning(value interface{}, columns ...string) IJobsExecLogDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j jobsExecLogDo) Not(conds ...gen.Condition) IJobsExecLogDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j jobsExecLogDo) Or(conds ...gen.Condition) IJobsExecLogDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j jobsExecLogDo) Select(conds ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j jobsExecLogDo) Where(conds ...gen.Condition) IJobsExecLogDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j jobsExecLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IJobsExecLogDo {
	return j.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (j jobsExecLogDo) Order(conds ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j jobsExecLogDo) Distinct(cols ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j jobsExecLogDo) Omit(cols ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j jobsExecLogDo) Join(table schema.Tabler, on ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j jobsExecLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j jobsExecLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j jobsExecLogDo) Group(cols ...field.Expr) IJobsExecLogDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j jobsExecLogDo) Having(conds ...gen.Condition) IJobsExecLogDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j jobsExecLogDo) Limit(limit int) IJobsExecLogDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j jobsExecLogDo) Offset(offset int) IJobsExecLogDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j jobsExecLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IJobsExecLogDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j jobsExecLogDo) Unscoped() IJobsExecLogDo {
	return j.withDO(j.DO.Unscoped())
}

func (j jobsExecLogDo) Create(values ...*model.JobsExecLog) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j jobsExecLogDo) CreateInBatches(values []*model.JobsExecLog, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j jobsExecLogDo) Save(values ...*model.JobsExecLog) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j jobsExecLogDo) First() (*model.JobsExecLog, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsExecLog), nil
	}
}

func (j jobsExecLogDo) Take() (*model.JobsExecLog, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsExecLog), nil
	}
}

func (j jobsExecLogDo) Last() (*model.JobsExecLog, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsExecLog), nil
	}
}

func (j jobsExecLogDo) Find() ([]*model.JobsExecLog, error) {
	result, err := j.DO.Find()
	return result.([]*model.JobsExecLog), err
}

func (j jobsExecLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.JobsExecLog, err error) {
	buf := make([]*model.JobsExecLog, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j jobsExecLogDo) FindInBatches(result *[]*model.JobsExecLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j jobsExecLogDo) Attrs(attrs ...field.AssignExpr) IJobsExecLogDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j jobsExecLogDo) Assign(attrs ...field.AssignExpr) IJobsExecLogDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j jobsExecLogDo) Joins(fields ...field.RelationField) IJobsExecLogDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j jobsExecLogDo) Preload(fields ...field.RelationField) IJobsExecLogDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j jobsExecLogDo) FirstOrInit() (*model.JobsExecLog, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsExecLog), nil
	}
}

func (j jobsExecLogDo) FirstOrCreate() (*model.JobsExecLog, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.JobsExecLog), nil
	}
}

func (j jobsExecLogDo) FindByPage(offset int, limit int) (result []*model.JobsExecLog, count int64, err error) {
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

func (j jobsExecLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j jobsExecLogDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j jobsExecLogDo) Delete(models ...*model.JobsExecLog) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *jobsExecLogDo) withDO(do gen.Dao) *jobsExecLogDo {
	j.DO = *do.(*gen.DO)
	return j
}
