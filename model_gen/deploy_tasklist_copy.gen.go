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

func newDeployTasklistCopy(db *gorm.DB, opts ...gen.DOOption) deployTasklistCopy {
	_deployTasklistCopy := deployTasklistCopy{}

	_deployTasklistCopy.deployTasklistCopyDo.UseDB(db, opts...)
	_deployTasklistCopy.deployTasklistCopyDo.UseModel(&model.DeployTasklistCopy{})

	tableName := _deployTasklistCopy.deployTasklistCopyDo.TableName()
	_deployTasklistCopy.ALL = field.NewAsterisk(tableName)
	_deployTasklistCopy.DeployFlag = field.NewString(tableName, "deploy_flag")
	_deployTasklistCopy.DepTaskID = field.NewString(tableName, "dep_task_id")
	_deployTasklistCopy.ServerID = field.NewString(tableName, "server_id")
	_deployTasklistCopy.WarArtifactid = field.NewString(tableName, "war_artifactid")
	_deployTasklistCopy.WarGroupid = field.NewString(tableName, "war_groupid")
	_deployTasklistCopy.Notes = field.NewString(tableName, "notes")
	_deployTasklistCopy.HospitalCode = field.NewString(tableName, "hospital_code")
	_deployTasklistCopy.Appstatu = field.NewString(tableName, "appstatu")
	_deployTasklistCopy.Lastversion = field.NewString(tableName, "lastversion")
	_deployTasklistCopy.Lastglobalgroupid = field.NewString(tableName, "lastglobalgroupid")
	_deployTasklistCopy.Deploydir = field.NewString(tableName, "deploydir")
	_deployTasklistCopy.Appurl = field.NewString(tableName, "appurl")
	_deployTasklistCopy.Deploystatu = field.NewString(tableName, "deploystatu")
	_deployTasklistCopy.Locked = field.NewString(tableName, "locked")
	_deployTasklistCopy.EnvCode = field.NewString(tableName, "env_code")
	_deployTasklistCopy.GlobalGroupID = field.NewString(tableName, "global_group_id")
	_deployTasklistCopy.InstallSoftID = field.NewString(tableName, "install_soft_id")
	_deployTasklistCopy.Orderindex = field.NewInt32(tableName, "orderindex")
	_deployTasklistCopy.Lastdeployver = field.NewString(tableName, "lastdeployver")
	_deployTasklistCopy.PublishReqVer = field.NewString(tableName, "publish_req_ver")
	_deployTasklistCopy.WarRemotefilesize = field.NewString(tableName, "war_remotefilesize")
	_deployTasklistCopy.WarDownloadPercent = field.NewString(tableName, "war_download_percent")
	_deployTasklistCopy.WarDownloadComplete = field.NewString(tableName, "war_download_complete")
	_deployTasklistCopy.OpUser = field.NewString(tableName, "op_user")
	_deployTasklistCopy.WarDownloadSpeed = field.NewString(tableName, "war_download_speed")
	_deployTasklistCopy.BeforehandID = field.NewString(tableName, "beforehand_id")
	_deployTasklistCopy.ThreadName = field.NewString(tableName, "thread_name")
	_deployTasklistCopy.ArtifactPath = field.NewString(tableName, "artifact_path")
	_deployTasklistCopy.ArtifactMd5 = field.NewString(tableName, "artifact_md5")
	_deployTasklistCopy.CheckWarStatusCode = field.NewString(tableName, "check_war_status_code")
	_deployTasklistCopy.PublishReqVer = field.NewString(tableName, "publishReqVer")
	_deployTasklistCopy.Deploy = field.NewInt32(tableName, "deploy_")
	_deployTasklistCopy.FrameworkVersion = field.NewString(tableName, "framework_version")
	_deployTasklistCopy.CreateTime = field.NewTime(tableName, "create_time")
	_deployTasklistCopy.ModifyTime = field.NewTime(tableName, "modify_time")
	_deployTasklistCopy.Comments = field.NewString(tableName, "comments")
	_deployTasklistCopy.StabilityAssuranceEventID = field.NewString(tableName, "stability_assurance_event_id")
	_deployTasklistCopy.JbossCliPid = field.NewString(tableName, "jboss_cli_pid")
	_deployTasklistCopy.JbossCliPidStartTime = field.NewTime(tableName, "jboss_cli_pid_start_time")
	_deployTasklistCopy.Checkwarstatu = field.NewString(tableName, "checkwarstatu")

	_deployTasklistCopy.fillFieldMap()

	return _deployTasklistCopy
}

type deployTasklistCopy struct {
	deployTasklistCopyDo

	ALL                       field.Asterisk
	DeployFlag                field.String // 部署是否完成\n0 部署未完成\n1 部署完成。
	DepTaskID                 field.String
	ServerID                  field.String // 服务器ID
	WarArtifactid             field.String
	WarGroupid                field.String
	Notes                     field.String // 备注
	HospitalCode              field.String // 院区
	Appstatu                  field.String // jboss应用HTTP状态
	Lastversion               field.String // 最后一次部署的版本号
	Lastglobalgroupid         field.String // 最近一次使用的global文件ID
	Deploydir                 field.String
	Appurl                    field.String // 程序网址,用于检查状态
	Deploystatu               field.String // 应用部署状态
	Locked                    field.String // 部署锁定状态 1 表示被锁 0 未被锁，可以部署
	EnvCode                   field.String
	GlobalGroupID             field.String
	InstallSoftID             field.String
	Orderindex                field.Int32
	Lastdeployver             field.String // 上一次部署成功的版本号
	PublishReqVer             field.String // 升级单离线下载后，将下载时候的WAR版本更新到此，用于下次离线部署的新旧版本对比
	WarRemotefilesize         field.String // war包大小
	WarDownloadPercent        field.String // 下载进度
	WarDownloadComplete       field.String // 0 待下载，下载完成。
	OpUser                    field.String // userid
	WarDownloadSpeed          field.String // 下载速度
	BeforehandID              field.String // 预下载UUID,用于预部署读取目录路径。
	ThreadName                field.String // 任务现场号
	ArtifactPath              field.String // 最近一次下载的war包存放路径
	ArtifactMd5               field.String // 最近一次下载的war包文件md5
	CheckWarStatusCode        field.String // check.war web状态\n值：存HTTP_code
	PublishReqVer             field.String // 同字段，publish_req_ver  ，用于新旧版本兼容
	Deploy                    field.Int32
	FrameworkVersion          field.String // java应用依赖的框架版本
	CreateTime                field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime                field.Time   // 记录修改时间（数据库自动写入）
	Comments                  field.String
	StabilityAssuranceEventID field.String // 稳定性保障中心操作事件ID
	JbossCliPid               field.String // jboss_cli进程ID
	JbossCliPidStartTime      field.Time   // jboss cli 进程启动时间
	Checkwarstatu             field.String

	fieldMap map[string]field.Expr
}

func (d deployTasklistCopy) Table(newTableName string) *deployTasklistCopy {
	d.deployTasklistCopyDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deployTasklistCopy) As(alias string) *deployTasklistCopy {
	d.deployTasklistCopyDo.DO = *(d.deployTasklistCopyDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deployTasklistCopy) updateTableName(table string) *deployTasklistCopy {
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
	d.Checkwarstatu = field.NewString(table, "checkwarstatu")

	d.fillFieldMap()

	return d
}

func (d *deployTasklistCopy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deployTasklistCopy) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 40)
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
	d.fieldMap["checkwarstatu"] = d.Checkwarstatu
}

func (d deployTasklistCopy) clone(db *gorm.DB) deployTasklistCopy {
	d.deployTasklistCopyDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deployTasklistCopy) replaceDB(db *gorm.DB) deployTasklistCopy {
	d.deployTasklistCopyDo.ReplaceDB(db)
	return d
}

type deployTasklistCopyDo struct{ gen.DO }

type IDeployTasklistCopyDo interface {
	gen.SubQuery
	Debug() IDeployTasklistCopyDo
	WithContext(ctx context.Context) IDeployTasklistCopyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeployTasklistCopyDo
	WriteDB() IDeployTasklistCopyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeployTasklistCopyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeployTasklistCopyDo
	Not(conds ...gen.Condition) IDeployTasklistCopyDo
	Or(conds ...gen.Condition) IDeployTasklistCopyDo
	Select(conds ...field.Expr) IDeployTasklistCopyDo
	Where(conds ...gen.Condition) IDeployTasklistCopyDo
	Order(conds ...field.Expr) IDeployTasklistCopyDo
	Distinct(cols ...field.Expr) IDeployTasklistCopyDo
	Omit(cols ...field.Expr) IDeployTasklistCopyDo
	Join(table schema.Tabler, on ...field.Expr) IDeployTasklistCopyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeployTasklistCopyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeployTasklistCopyDo
	Group(cols ...field.Expr) IDeployTasklistCopyDo
	Having(conds ...gen.Condition) IDeployTasklistCopyDo
	Limit(limit int) IDeployTasklistCopyDo
	Offset(offset int) IDeployTasklistCopyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployTasklistCopyDo
	Unscoped() IDeployTasklistCopyDo
	Create(values ...*model.DeployTasklistCopy) error
	CreateInBatches(values []*model.DeployTasklistCopy, batchSize int) error
	Save(values ...*model.DeployTasklistCopy) error
	First() (*model.DeployTasklistCopy, error)
	Take() (*model.DeployTasklistCopy, error)
	Last() (*model.DeployTasklistCopy, error)
	Find() ([]*model.DeployTasklistCopy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployTasklistCopy, err error)
	FindInBatches(result *[]*model.DeployTasklistCopy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DeployTasklistCopy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeployTasklistCopyDo
	Assign(attrs ...field.AssignExpr) IDeployTasklistCopyDo
	Joins(fields ...field.RelationField) IDeployTasklistCopyDo
	Preload(fields ...field.RelationField) IDeployTasklistCopyDo
	FirstOrInit() (*model.DeployTasklistCopy, error)
	FirstOrCreate() (*model.DeployTasklistCopy, error)
	FindByPage(offset int, limit int) (result []*model.DeployTasklistCopy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeployTasklistCopyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deployTasklistCopyDo) Debug() IDeployTasklistCopyDo {
	return d.withDO(d.DO.Debug())
}

func (d deployTasklistCopyDo) WithContext(ctx context.Context) IDeployTasklistCopyDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deployTasklistCopyDo) ReadDB() IDeployTasklistCopyDo {
	return d.Clauses(dbresolver.Read)
}

func (d deployTasklistCopyDo) WriteDB() IDeployTasklistCopyDo {
	return d.Clauses(dbresolver.Write)
}

func (d deployTasklistCopyDo) Session(config *gorm.Session) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Session(config))
}

func (d deployTasklistCopyDo) Clauses(conds ...clause.Expression) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deployTasklistCopyDo) Returning(value interface{}, columns ...string) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deployTasklistCopyDo) Not(conds ...gen.Condition) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deployTasklistCopyDo) Or(conds ...gen.Condition) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deployTasklistCopyDo) Select(conds ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deployTasklistCopyDo) Where(conds ...gen.Condition) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deployTasklistCopyDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDeployTasklistCopyDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d deployTasklistCopyDo) Order(conds ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deployTasklistCopyDo) Distinct(cols ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deployTasklistCopyDo) Omit(cols ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deployTasklistCopyDo) Join(table schema.Tabler, on ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deployTasklistCopyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deployTasklistCopyDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deployTasklistCopyDo) Group(cols ...field.Expr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deployTasklistCopyDo) Having(conds ...gen.Condition) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deployTasklistCopyDo) Limit(limit int) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deployTasklistCopyDo) Offset(offset int) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deployTasklistCopyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deployTasklistCopyDo) Unscoped() IDeployTasklistCopyDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deployTasklistCopyDo) Create(values ...*model.DeployTasklistCopy) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deployTasklistCopyDo) CreateInBatches(values []*model.DeployTasklistCopy, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deployTasklistCopyDo) Save(values ...*model.DeployTasklistCopy) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deployTasklistCopyDo) First() (*model.DeployTasklistCopy, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklistCopy), nil
	}
}

func (d deployTasklistCopyDo) Take() (*model.DeployTasklistCopy, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklistCopy), nil
	}
}

func (d deployTasklistCopyDo) Last() (*model.DeployTasklistCopy, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklistCopy), nil
	}
}

func (d deployTasklistCopyDo) Find() ([]*model.DeployTasklistCopy, error) {
	result, err := d.DO.Find()
	return result.([]*model.DeployTasklistCopy), err
}

func (d deployTasklistCopyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DeployTasklistCopy, err error) {
	buf := make([]*model.DeployTasklistCopy, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deployTasklistCopyDo) FindInBatches(result *[]*model.DeployTasklistCopy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deployTasklistCopyDo) Attrs(attrs ...field.AssignExpr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deployTasklistCopyDo) Assign(attrs ...field.AssignExpr) IDeployTasklistCopyDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deployTasklistCopyDo) Joins(fields ...field.RelationField) IDeployTasklistCopyDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deployTasklistCopyDo) Preload(fields ...field.RelationField) IDeployTasklistCopyDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deployTasklistCopyDo) FirstOrInit() (*model.DeployTasklistCopy, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklistCopy), nil
	}
}

func (d deployTasklistCopyDo) FirstOrCreate() (*model.DeployTasklistCopy, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DeployTasklistCopy), nil
	}
}

func (d deployTasklistCopyDo) FindByPage(offset int, limit int) (result []*model.DeployTasklistCopy, count int64, err error) {
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

func (d deployTasklistCopyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deployTasklistCopyDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deployTasklistCopyDo) Delete(models ...*model.DeployTasklistCopy) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deployTasklistCopyDo) withDO(do gen.Dao) *deployTasklistCopyDo {
	d.DO = *do.(*gen.DO)
	return d
}
