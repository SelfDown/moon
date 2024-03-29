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

func newDeployTasklist88(db *gorm.DB, opts ...gen.DOOption) deployTasklist88 {
	_deployTasklist88 := deployTasklist88{}

	_deployTasklist88.deployTasklist88Do.UseDB(db, opts...)
	_deployTasklist88.deployTasklist88Do.UseModel(&model.DeployTasklist88{})

	tableName := _deployTasklist88.deployTasklist88Do.TableName()
	_deployTasklist88.ALL = field.NewAsterisk(tableName)
	_deployTasklist88.DeployFlag = field.NewString(tableName, "deploy_flag")
	_deployTasklist88.DepTaskID = field.NewString(tableName, "dep_task_id")
	_deployTasklist88.ServerID = field.NewString(tableName, "server_id")
	_deployTasklist88.WarArtifactid = field.NewString(tableName, "war_artifactid")
	_deployTasklist88.WarGroupid = field.NewString(tableName, "war_groupid")
	_deployTasklist88.Notes = field.NewString(tableName, "notes")
	_deployTasklist88.HospitalCode = field.NewString(tableName, "hospital_code")
	_deployTasklist88.Appstatu = field.NewString(tableName, "appstatu")
	_deployTasklist88.Lastversion = field.NewString(tableName, "lastversion")
	_deployTasklist88.Lastglobalgroupid = field.NewString(tableName, "lastglobalgroupid")
	_deployTasklist88.Deploydir = field.NewString(tableName, "deploydir")
	_deployTasklist88.Appurl = field.NewString(tableName, "appurl")
	_deployTasklist88.Deploystatu = field.NewString(tableName, "deploystatu")
	_deployTasklist88.Locked = field.NewString(tableName, "locked")
	_deployTasklist88.EnvCode = field.NewString(tableName, "env_code")
	_deployTasklist88.GlobalGroupID = field.NewString(tableName, "global_group_id")
	_deployTasklist88.InstallSoftID = field.NewString(tableName, "install_soft_id")
	_deployTasklist88.Orderindex = field.NewInt32(tableName, "orderindex")
	_deployTasklist88.Lastdeployver = field.NewString(tableName, "lastdeployver")
	_deployTasklist88.PublishReqVer = field.NewString(tableName, "publish_req_ver")
	_deployTasklist88.WarRemotefilesize = field.NewString(tableName, "war_remotefilesize")
	_deployTasklist88.WarDownloadPercent = field.NewString(tableName, "war_download_percent")
	_deployTasklist88.WarDownloadComplete = field.NewString(tableName, "war_download_complete")
	_deployTasklist88.OpUser = field.NewString(tableName, "op_user")
	_deployTasklist88.WarDownloadSpeed = field.NewString(tableName, "war_download_speed")
	_deployTasklist88.BeforehandID = field.NewString(tableName, "beforehand_id")
	_deployTasklist88.ThreadName = field.NewString(tableName, "thread_name")
	_deployTasklist88.ArtifactPath = field.NewString(tableName, "artifact_path")
	_deployTasklist88.ArtifactMd5 = field.NewString(tableName, "artifact_md5")
	_deployTasklist88.CheckWarStatusCode = field.NewString(tableName, "check_war_status_code")
	_deployTasklist88.PublishReqVer = field.NewString(tableName, "publishReqVer")
	_deployTasklist88.Deploy = field.NewInt32(tableName, "deploy_")
	_deployTasklist88.FrameworkVersion = field.NewString(tableName, "framework_version")
	_deployTasklist88.CreateTime = field.NewTime(tableName, "create_time")
	_deployTasklist88.ModifyTime = field.NewTime(tableName, "modify_time")
	_deployTasklist88.Comments = field.NewString(tableName, "comments")
	_deployTasklist88.StabilityAssuranceEventID = field.NewString(tableName, "stability_assurance_event_id")
	_deployTasklist88.JbossCliPid = field.NewString(tableName, "jboss_cli_pid")
	_deployTasklist88.JbossCliPidStartTime = field.NewTime(tableName, "jboss_cli_pid_start_time")

	_deployTasklist88.fillFieldMap()

	return _deployTasklist88
}

type deployTasklist88 struct {
	deployTasklist88Do

	ALL field.Asterisk
	/*
		部署是否完成
		0 部署未完成
		1 部署完成。
	*/
	DeployFlag          field.String
	DepTaskID           field.String
	ServerID            field.String // 服务器ID
	WarArtifactid       field.String
	WarGroupid          field.String
	Notes               field.String // 备注
	HospitalCode        field.String // 院区
	Appstatu            field.String // jboss应用HTTP状态
	Lastversion         field.String // 最后一次部署的版本号
	Lastglobalgroupid   field.String // 最近一次使用的global文件ID
	Deploydir           field.String
	Appurl              field.String // 程序网址,用于检查状态
	Deploystatu         field.String // 应用部署状态
	Locked              field.String // 部署锁定状态 1 表示被锁 0 未被锁，可以部署
	EnvCode             field.String
	GlobalGroupID       field.String
	InstallSoftID       field.String
	Orderindex          field.Int32
	Lastdeployver       field.String // 上一次部署成功的版本号
	PublishReqVer       field.String // 升级单离线下载后，将下载时候的WAR版本更新到此，用于下次离线部署的新旧版本对比
	WarRemotefilesize   field.String // war包大小
	WarDownloadPercent  field.String // 下载进度
	WarDownloadComplete field.String // 0 待下载，下载完成。
	OpUser              field.String // userid
	WarDownloadSpeed    field.String // 下载速度
	BeforehandID        field.String // 预下载UUID,用于预部署读取目录路径。
	ThreadName          field.String // 任务现场号
	ArtifactPath        field.String // 最近一次下载的war包存放路径
	ArtifactMd5         field.String // 最近一次下载的war包文件md5
	/*
		check.war web状态
		值：存HTTP_code
	*/
	CheckWarStatusCode        field.String
	PublishReqVer             field.String // 同字段，publish_req_ver  ，用于新旧版本兼容
	Deploy                    field.Int32
	FrameworkVersion          field.String // java应用依赖的框架版本
	CreateTime                field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime                field.Time   // 记录修改时间（数据库自动写入）
	Comments                  field.String
	StabilityAssuranceEventID field.String // 稳定性保障中心操作事件ID
	JbossCliPid               field.String // jboss_cli进程ID
	JbossCliPidStartTime      field.Time   // jboss cli 进程启动时间

	fieldMap map[string]field.Expr
}

func (d deployTasklist88) Table(newTableName string) *deployTasklist88 {
	d.deployTasklist88Do.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deployTasklist88) As(alias string) *deployTasklist88 {
	d.deployTasklist88Do.DO = *(d.deployTasklist88Do.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deployTasklist88) updateTableName(table string) *deployTasklist88 {
	d.ALL = field.NewAsterisk(table)
	d.DeployFlag = field.NewString(table, "deploy_flag")
	d.DepTaskID = field.NewString(table, "dep_task_id")
	d.ServerID = field.NewString(table, "server_id")
	d.WarArtifactid = field.NewString(table, "war_artifactid")
	d.WarGroupid = field.NewString(table, "war_groupid")
	d.Notes = field.NewString(table, "notes")
	d.HospitalCode = field.NewString(table, "hospital_code")
	d.Appstatu = field.NewString(table, "appstatu")
	d.Lastversion = field.NewString(table, "lastversion")
	d.Lastglobalgroupid = field.NewString(table, "lastglobalgroupid")
	d.Deploydir = field.NewString(table, "deploydir")
	d.Appurl = field.NewString(table, "appurl")
	d.Deploystatu = field.NewString(table, "deploystatu")
	d.Locked = field.NewString(table, "locked")
	d.EnvCode = field.NewString(table, "env_code")
	d.GlobalGroupID = field.NewString(table, "global_group_id")
	d.InstallSoftID = field.NewString(table, "install_soft_id")
	d.Orderindex = field.NewInt32(table, "orderindex")
	d.Lastdeployver = field.NewString(table, "lastdeployver")
	d.PublishReqVer = field.NewString(table, "publish_req_ver")
	d.WarRemotefilesize = field.NewString(table, "war_remotefilesize")
	d.WarDownloadPercent = field.NewString(table, "war_download_percent")
	d.WarDownloadComplete = field.NewString(table, "war_download_complete")
	d.OpUser = field.NewString(table, "op_user")
	d.WarDownloadSpeed = field.NewString(table, "war_download_speed")
	d.BeforehandID = field.NewString(table, "beforehand_id")
	d.ThreadName = field.NewString(table, "thread_name")
	d.ArtifactPath = field.NewString(table, "artifact_path")
	d.ArtifactMd5 = field.NewString(table, "artifact_md5")
	d.CheckWarStatusCode = field.NewString(table, "check_war_status_code")
	d.PublishReqVer = field.NewString(table, "publishReqVer")
	d.Deploy = field.NewInt32(table, "deploy_")
	d.FrameworkVersion = field.NewString(table, "framework_version")
	d.CreateTime = field.NewTime(table, "create_time")
	d.ModifyTime = field.NewTime(table, "modify_time")
	d.Comments = field.NewString(table, "comments")
	d.StabilityAssuranceEventID = field.NewString(table, "stability_assurance_event_id")
	d.JbossCliPid = field.NewString(table, "jboss_cli_pid")
	d.JbossCliPidStartTime = field.NewTime(table, "jboss_cli_pid_start_time")

	d.fillFieldMap()

	return d
}

func (d *deployTasklist88) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deployTasklist88) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 39)
	d.fieldMap["deploy_flag"] = d.DeployFlag
	d.fieldMap["dep_task_id"] = d.DepTaskID
	d.fieldMap["server_id"] = d.ServerID
	d.fieldMap["war_artifactid"] = d.WarArtifactid
	d.fieldMap["war_groupid"] = d.WarGroupid
	d.fieldMap["notes"] = d.Notes
	d.fieldMap["hospital_code"] = d.HospitalCode
	d.fieldMap["appstatu"] = d.Appstatu
	d.fieldMap["lastversion"] = d.Lastversion
	d.fieldMap["lastglobalgroupid"] = d.Lastglobalgroupid
	d.fieldMap["deploydir"] = d.Deploydir
	d.fieldMap["appurl"] = d.Appurl
	d.fieldMap["deploystatu"] = d.Deploystatu
	d.fieldMap["locked"] = d.Locked
	d.fieldMap["env_code"] = d.EnvCode
	d.fieldMap["global_group_id"] = d.GlobalGroupID
	d.fieldMap["install_soft_id"] = d.InstallSoftID
	d.fieldMap["orderindex"] = d.Orderindex
	d.fieldMap["lastdeployver"] = d.Lastdeployver
	d.fieldMap["publish_req_ver"] = d.PublishReqVer
	d.fieldMap["war_remotefilesize"] = d.WarRemotefilesize
	d.fieldMap["war_download_percent"] = d.WarDownloadPercent
	d.fieldMap["war_download_complete"] = d.WarDownloadComplete
	d.fieldMap["op_user"] = d.OpUser
	d.fieldMap["war_download_speed"] = d.WarDownloadSpeed
	d.fieldMap["beforehand_id"] = d.BeforehandID
	d.fieldMap["thread_name"] = d.ThreadName
	d.fieldMap["artifact_path"] = d.ArtifactPath
	d.fieldMap["artifact_md5"] = d.ArtifactMd5
	d.fieldMap["check_war_status_code"] = d.CheckWarStatusCode
	d.fieldMap["publishReqVer"] = d.PublishReqVer
	d.fieldMap["deploy_"] = d.Deploy
	d.fieldMap["framework_version"] = d.FrameworkVersion
	d.fieldMap["create_time"] = d.CreateTime
	d.fieldMap["modify_time"] = d.ModifyTime
	d.fieldMap["comments"] = d.Comments
	d.fieldMap["stability_assurance_event_id"] = d.StabilityAssuranceEventID
	d.fieldMap["jboss_cli_pid"] = d.JbossCliPid
	d.fieldMap["jboss_cli_pid_start_time"] = d.JbossCliPidStartTime
}

func (d deployTasklist88) clone(db *gorm.DB) deployTasklist88 {
	d.deployTasklist88Do.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deployTasklist88) replaceDB(db *gorm.DB) deployTasklist88 {
	d.deployTasklist88Do.ReplaceDB(db)
	return d
}

type deployTasklist88Do struct{ gen.DO }

type IDeployTasklist88Do interface {
	gen.SubQuery
	Debug() IDeployTasklist88Do
	WithContext(ctx context.Context) IDeployTasklist88Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeployTasklist88Do
	WriteDB() IDeployTasklist88Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeployTasklist88Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeployTasklist88Do
	Not(conds ...gen.Condition) IDeployTasklist88Do
	Or(conds ...gen.Condition) IDeployTasklist88Do
	Select(conds ...field.Expr) IDeployTasklist88Do
	Where(conds ...gen.Condition) IDeployTasklist88Do
	Order(conds ...field.Expr) IDeployTasklist88Do
	Distinct(cols ...field.Expr) IDeployTasklist88Do
	Omit(cols ...field.Expr) IDeployTasklist88Do
	Join(table schema.Tabler, on ...field.Expr) IDeployTasklist88Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeployTasklist88Do
	RightJoin(table schema.Tabler, on ...field.Expr) IDeployTasklist88Do
	Group(cols ...field.Expr) IDeployTasklist88Do
	Having(conds ...gen.Condition) IDeployTasklist88Do
	Limit(limit int) IDeployTasklist88Do
	Offset(offset int) IDeployTasklist88Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployTasklist88Do
	Unscoped() IDeployTasklist88Do
	Create(values ...*model.DeployTasklist88) error
	CreateInBatches(values []*model.DeployTasklist88, batchSize int) error
	Save(values ...*model.DeployTasklist88) error
	First() (*model.DeployTasklist88, error)
	Take() (*model.DeployTasklist88, error)
	Last() (*model.DeployTasklist88, error)
	Find() ([]*model.DeployTasklist88, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployTasklist88, err error)
	FindInBatches(result *[]*model.DeployTasklist88, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeployTasklist88) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeployTasklist88Do
	Assign(attrs ...field.AssignExpr) IDeployTasklist88Do
	Joins(fields ...field.RelationField) IDeployTasklist88Do
	Preload(fields ...field.RelationField) IDeployTasklist88Do
	FirstOrInit() (*model.DeployTasklist88, error)
	FirstOrCreate() (*model.DeployTasklist88, error)
	FindByPage(offset int, limit int) (result []*model.DeployTasklist88, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeployTasklist88Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deployTasklist88Do) Debug() IDeployTasklist88Do {
	return d.withDO(d.DO.Debug())
}

func (d deployTasklist88Do) WithContext(ctx context.Context) IDeployTasklist88Do {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deployTasklist88Do) ReadDB() IDeployTasklist88Do {
	return d.Clauses(dbresolver.Read)
}

func (d deployTasklist88Do) WriteDB() IDeployTasklist88Do {
	return d.Clauses(dbresolver.Write)
}

func (d deployTasklist88Do) Session(config *gorm.Session) IDeployTasklist88Do {
	return d.withDO(d.DO.Session(config))
}

func (d deployTasklist88Do) Clauses(conds ...clause.Expression) IDeployTasklist88Do {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deployTasklist88Do) Returning(value interface{}, columns ...string) IDeployTasklist88Do {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deployTasklist88Do) Not(conds ...gen.Condition) IDeployTasklist88Do {
	return d.withDO(d.DO.Not(conds...))
}

func (d deployTasklist88Do) Or(conds ...gen.Condition) IDeployTasklist88Do {
	return d.withDO(d.DO.Or(conds...))
}

func (d deployTasklist88Do) Select(conds ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.Select(conds...))
}

func (d deployTasklist88Do) Where(conds ...gen.Condition) IDeployTasklist88Do {
	return d.withDO(d.DO.Where(conds...))
}

func (d deployTasklist88Do) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDeployTasklist88Do {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d deployTasklist88Do) Order(conds ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.Order(conds...))
}

func (d deployTasklist88Do) Distinct(cols ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deployTasklist88Do) Omit(cols ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deployTasklist88Do) Join(table schema.Tabler, on ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deployTasklist88Do) LeftJoin(table schema.Tabler, on ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deployTasklist88Do) RightJoin(table schema.Tabler, on ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deployTasklist88Do) Group(cols ...field.Expr) IDeployTasklist88Do {
	return d.withDO(d.DO.Group(cols...))
}

func (d deployTasklist88Do) Having(conds ...gen.Condition) IDeployTasklist88Do {
	return d.withDO(d.DO.Having(conds...))
}

func (d deployTasklist88Do) Limit(limit int) IDeployTasklist88Do {
	return d.withDO(d.DO.Limit(limit))
}

func (d deployTasklist88Do) Offset(offset int) IDeployTasklist88Do {
	return d.withDO(d.DO.Offset(offset))
}

func (d deployTasklist88Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployTasklist88Do {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deployTasklist88Do) Unscoped() IDeployTasklist88Do {
	return d.withDO(d.DO.Unscoped())
}

func (d deployTasklist88Do) Create(values ...*model.DeployTasklist88) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deployTasklist88Do) CreateInBatches(values []*model.DeployTasklist88, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deployTasklist88Do) Save(values ...*model.DeployTasklist88) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deployTasklist88Do) First() (*model.DeployTasklist88, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklist88), nil
	}
}

func (d deployTasklist88Do) Take() (*model.DeployTasklist88, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklist88), nil
	}
}

func (d deployTasklist88Do) Last() (*model.DeployTasklist88, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklist88), nil
	}
}

func (d deployTasklist88Do) Find() ([]*model.DeployTasklist88, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeployTasklist88), err
}

func (d deployTasklist88Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployTasklist88, err error) {
	buf := make([]*model.DeployTasklist88, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deployTasklist88Do) FindInBatches(result *[]*model.DeployTasklist88, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deployTasklist88Do) Attrs(attrs ...field.AssignExpr) IDeployTasklist88Do {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deployTasklist88Do) Assign(attrs ...field.AssignExpr) IDeployTasklist88Do {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deployTasklist88Do) Joins(fields ...field.RelationField) IDeployTasklist88Do {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deployTasklist88Do) Preload(fields ...field.RelationField) IDeployTasklist88Do {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deployTasklist88Do) FirstOrInit() (*model.DeployTasklist88, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklist88), nil
	}
}

func (d deployTasklist88Do) FirstOrCreate() (*model.DeployTasklist88, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklist88), nil
	}
}

func (d deployTasklist88Do) FindByPage(offset int, limit int) (result []*model.DeployTasklist88, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deployTasklist88Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deployTasklist88Do) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deployTasklist88Do) Delete(models ...*model.DeployTasklist88) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deployTasklist88Do) withDO(do gen.Dao) *deployTasklist88Do {
	d.DO = *do.(*gen.DO)
	return d
}
