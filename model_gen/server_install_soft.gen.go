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

func newServerInstallSoft(db *gorm.DB, opts ...gen.DOOption) serverInstallSoft {
	_serverInstallSoft := serverInstallSoft{}

	_serverInstallSoft.serverInstallSoftDo.UseDB(db, opts...)
	_serverInstallSoft.serverInstallSoftDo.UseModel(&model.ServerInstallSoft{})

	tableName := _serverInstallSoft.serverInstallSoftDo.TableName()
	_serverInstallSoft.ALL = field.NewAsterisk(tableName)
	_serverInstallSoft.ServicePort = field.NewString(tableName, "service_port")
	_serverInstallSoft.InstallSoftID = field.NewString(tableName, "install_soft_id")
	_serverInstallSoft.ServerID = field.NewString(tableName, "server_id")
	_serverInstallSoft.SoftName = field.NewString(tableName, "soft_name")
	_serverInstallSoft.SoftType = field.NewString(tableName, "soft_type")
	_serverInstallSoft.SoftHome = field.NewString(tableName, "soft_home")
	_serverInstallSoft.Portoffset = field.NewString(tableName, "portoffset")
	_serverInstallSoft.CreateTime = field.NewTime(tableName, "create_time")
	_serverInstallSoft.ModifyTime = field.NewTime(tableName, "modify_time")
	_serverInstallSoft.Comments = field.NewString(tableName, "comments")
	_serverInstallSoft.DockerContainerName = field.NewString(tableName, "docker_container_name")
	_serverInstallSoft.StartSoftwareCommand = field.NewString(tableName, "start_software_command")
	_serverInstallSoft.StopSoftwareCommand = field.NewString(tableName, "stop_software_command")
	_serverInstallSoft.RestartSoftwareCommand = field.NewString(tableName, "restart_software_command")
	_serverInstallSoft.ServerSoftLogsID = field.NewString(tableName, "server_soft_logs_id")
	_serverInstallSoft.ServerOsUserID = field.NewInt32(tableName, "server_os_user_id")
	_serverInstallSoft.SoftID = field.NewInt32(tableName, "soft_id")
	_serverInstallSoft.SoftwareLiveCheckCommand = field.NewString(tableName, "software_live_check_command")
	_serverInstallSoft.ServerSoftLogPath = field.NewString(tableName, "server_soft_log_path")
	_serverInstallSoft.IsMainSoft = field.NewString(tableName, "is_main_soft")
	_serverInstallSoft.IsDocker = field.NewString(tableName, "is_docker")
	_serverInstallSoft.PidShell = field.NewString(tableName, "pid_shell")
	_serverInstallSoft.DockerServerOsUserID = field.NewInt32(tableName, "docker_server_os_user_id")
	_serverInstallSoft.DbSchema = field.NewString(tableName, "db_schema")
	_serverInstallSoft.DbOraSid = field.NewString(tableName, "db_ora_sid")
	_serverInstallSoft.DbURLSuffix = field.NewString(tableName, "db_url_suffix")
	_serverInstallSoft.SoftVersion = field.NewString(tableName, "soft_version")
	_serverInstallSoft.WebsqlCacheKey = field.NewString(tableName, "websql_cache_key")
	_serverInstallSoft.OpsConfigGroupID = field.NewString(tableName, "ops_config_group_id")
	_serverInstallSoft.InstallType = field.NewString(tableName, "install_type")
	_serverInstallSoft.InstallSourceType = field.NewString(tableName, "install_source_type")
	_serverInstallSoft.InstallSrcPath = field.NewString(tableName, "install_src_path")
	_serverInstallSoft.InstallOsUser = field.NewInt32(tableName, "install_os_user")
	_serverInstallSoft.InstallScript = field.NewString(tableName, "install_script")
	_serverInstallSoft.UninstallScript = field.NewString(tableName, "uninstall_script")
	_serverInstallSoft.SoftConfPath = field.NewString(tableName, "soft_conf_path")
	_serverInstallSoft.InstallStatus = field.NewString(tableName, "install_status")
	_serverInstallSoft.ServerSoftConfDefID = field.NewString(tableName, "server_soft_conf_def_id")

	_serverInstallSoft.fillFieldMap()

	return _serverInstallSoft
}

type serverInstallSoft struct {
	serverInstallSoftDo

	ALL           field.Asterisk
	ServicePort   field.String // 服务端口 默认空
	InstallSoftID field.String
	ServerID      field.String // 服务器ID
	SoftName      field.String
	/*
		1、Jboss中间件
		2、文件服务器
	*/
	SoftType                 field.String
	SoftHome                 field.String // 备注
	Portoffset               field.String // 端口偏移量
	CreateTime               field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime               field.Time   // 记录修改时间（数据库自动写入）
	Comments                 field.String // 备注说明
	DockerContainerName      field.String // docker 容器名称，用于对容器进行操作。
	StartSoftwareCommand     field.String // 软件启动命令
	StopSoftwareCommand      field.String // 软件关闭命令
	RestartSoftwareCommand   field.String // 软件重启命令
	ServerSoftLogsID         field.String // 软件日志组, server_soft_logs.server_soft_logs_id
	ServerOsUserID           field.Int32  // 软件使用上面系统用启动，server_os_users.server_os_user_id
	SoftID                   field.Int32  // 自增长主键
	SoftwareLiveCheckCommand field.String // 作废；软件状态检查，服务器文件路径和web api
	ServerSoftLogPath        field.String
	IsMainSoft               field.String
	IsDocker                 field.String
	PidShell                 field.String
	DockerServerOsUserID     field.Int32
	DbSchema                 field.String
	DbOraSid                 field.String
	DbURLSuffix              field.String
	SoftVersion              field.String
	WebsqlCacheKey           field.String
	OpsConfigGroupID         field.String
	InstallType              field.String
	InstallSourceType        field.String
	InstallSrcPath           field.String
	InstallOsUser            field.Int32
	InstallScript            field.String
	UninstallScript          field.String
	SoftConfPath             field.String
	InstallStatus            field.String
	ServerSoftConfDefID      field.String

	fieldMap map[string]field.Expr
}

func (s serverInstallSoft) Table(newTableName string) *serverInstallSoft {
	s.serverInstallSoftDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s serverInstallSoft) As(alias string) *serverInstallSoft {
	s.serverInstallSoftDo.DO = *(s.serverInstallSoftDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *serverInstallSoft) updateTableName(table string) *serverInstallSoft {
	s.ALL = field.NewAsterisk(table)
	s.ServicePort = field.NewString(table, "service_port")
	s.InstallSoftID = field.NewString(table, "install_soft_id")
	s.ServerID = field.NewString(table, "server_id")
	s.SoftName = field.NewString(table, "soft_name")
	s.SoftType = field.NewString(table, "soft_type")
	s.SoftHome = field.NewString(table, "soft_home")
	s.Portoffset = field.NewString(table, "portoffset")
	s.CreateTime = field.NewTime(table, "create_time")
	s.ModifyTime = field.NewTime(table, "modify_time")
	s.Comments = field.NewString(table, "comments")
	s.DockerContainerName = field.NewString(table, "docker_container_name")
	s.StartSoftwareCommand = field.NewString(table, "start_software_command")
	s.StopSoftwareCommand = field.NewString(table, "stop_software_command")
	s.RestartSoftwareCommand = field.NewString(table, "restart_software_command")
	s.ServerSoftLogsID = field.NewString(table, "server_soft_logs_id")
	s.ServerOsUserID = field.NewInt32(table, "server_os_user_id")
	s.SoftID = field.NewInt32(table, "soft_id")
	s.SoftwareLiveCheckCommand = field.NewString(table, "software_live_check_command")
	s.ServerSoftLogPath = field.NewString(table, "server_soft_log_path")
	s.IsMainSoft = field.NewString(table, "is_main_soft")
	s.IsDocker = field.NewString(table, "is_docker")
	s.PidShell = field.NewString(table, "pid_shell")
	s.DockerServerOsUserID = field.NewInt32(table, "docker_server_os_user_id")
	s.DbSchema = field.NewString(table, "db_schema")
	s.DbOraSid = field.NewString(table, "db_ora_sid")
	s.DbURLSuffix = field.NewString(table, "db_url_suffix")
	s.SoftVersion = field.NewString(table, "soft_version")
	s.WebsqlCacheKey = field.NewString(table, "websql_cache_key")
	s.OpsConfigGroupID = field.NewString(table, "ops_config_group_id")
	s.InstallType = field.NewString(table, "install_type")
	s.InstallSourceType = field.NewString(table, "install_source_type")
	s.InstallSrcPath = field.NewString(table, "install_src_path")
	s.InstallOsUser = field.NewInt32(table, "install_os_user")
	s.InstallScript = field.NewString(table, "install_script")
	s.UninstallScript = field.NewString(table, "uninstall_script")
	s.SoftConfPath = field.NewString(table, "soft_conf_path")
	s.InstallStatus = field.NewString(table, "install_status")
	s.ServerSoftConfDefID = field.NewString(table, "server_soft_conf_def_id")

	s.fillFieldMap()

	return s
}

func (s *serverInstallSoft) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *serverInstallSoft) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 38)
	s.fieldMap["service_port"] = s.ServicePort
	s.fieldMap["install_soft_id"] = s.InstallSoftID
	s.fieldMap["server_id"] = s.ServerID
	s.fieldMap["soft_name"] = s.SoftName
	s.fieldMap["soft_type"] = s.SoftType
	s.fieldMap["soft_home"] = s.SoftHome
	s.fieldMap["portoffset"] = s.Portoffset
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["modify_time"] = s.ModifyTime
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["docker_container_name"] = s.DockerContainerName
	s.fieldMap["start_software_command"] = s.StartSoftwareCommand
	s.fieldMap["stop_software_command"] = s.StopSoftwareCommand
	s.fieldMap["restart_software_command"] = s.RestartSoftwareCommand
	s.fieldMap["server_soft_logs_id"] = s.ServerSoftLogsID
	s.fieldMap["server_os_user_id"] = s.ServerOsUserID
	s.fieldMap["soft_id"] = s.SoftID
	s.fieldMap["software_live_check_command"] = s.SoftwareLiveCheckCommand
	s.fieldMap["server_soft_log_path"] = s.ServerSoftLogPath
	s.fieldMap["is_main_soft"] = s.IsMainSoft
	s.fieldMap["is_docker"] = s.IsDocker
	s.fieldMap["pid_shell"] = s.PidShell
	s.fieldMap["docker_server_os_user_id"] = s.DockerServerOsUserID
	s.fieldMap["db_schema"] = s.DbSchema
	s.fieldMap["db_ora_sid"] = s.DbOraSid
	s.fieldMap["db_url_suffix"] = s.DbURLSuffix
	s.fieldMap["soft_version"] = s.SoftVersion
	s.fieldMap["websql_cache_key"] = s.WebsqlCacheKey
	s.fieldMap["ops_config_group_id"] = s.OpsConfigGroupID
	s.fieldMap["install_type"] = s.InstallType
	s.fieldMap["install_source_type"] = s.InstallSourceType
	s.fieldMap["install_src_path"] = s.InstallSrcPath
	s.fieldMap["install_os_user"] = s.InstallOsUser
	s.fieldMap["install_script"] = s.InstallScript
	s.fieldMap["uninstall_script"] = s.UninstallScript
	s.fieldMap["soft_conf_path"] = s.SoftConfPath
	s.fieldMap["install_status"] = s.InstallStatus
	s.fieldMap["server_soft_conf_def_id"] = s.ServerSoftConfDefID
}

func (s serverInstallSoft) clone(db *gorm.DB) serverInstallSoft {
	s.serverInstallSoftDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s serverInstallSoft) replaceDB(db *gorm.DB) serverInstallSoft {
	s.serverInstallSoftDo.ReplaceDB(db)
	return s
}

type serverInstallSoftDo struct{ gen.DO }

type IServerInstallSoftDo interface {
	gen.SubQuery
	Debug() IServerInstallSoftDo
	WithContext(ctx context.Context) IServerInstallSoftDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerInstallSoftDo
	WriteDB() IServerInstallSoftDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerInstallSoftDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerInstallSoftDo
	Not(conds ...gen.Condition) IServerInstallSoftDo
	Or(conds ...gen.Condition) IServerInstallSoftDo
	Select(conds ...field.Expr) IServerInstallSoftDo
	Where(conds ...gen.Condition) IServerInstallSoftDo
	Order(conds ...field.Expr) IServerInstallSoftDo
	Distinct(cols ...field.Expr) IServerInstallSoftDo
	Omit(cols ...field.Expr) IServerInstallSoftDo
	Join(table schema.Tabler, on ...field.Expr) IServerInstallSoftDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerInstallSoftDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerInstallSoftDo
	Group(cols ...field.Expr) IServerInstallSoftDo
	Having(conds ...gen.Condition) IServerInstallSoftDo
	Limit(limit int) IServerInstallSoftDo
	Offset(offset int) IServerInstallSoftDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerInstallSoftDo
	Unscoped() IServerInstallSoftDo
	Create(values ...*model.ServerInstallSoft) error
	CreateInBatches(values []*model.ServerInstallSoft, batchSize int) error
	Save(values ...*model.ServerInstallSoft) error
	First() (*model.ServerInstallSoft, error)
	Take() (*model.ServerInstallSoft, error)
	Last() (*model.ServerInstallSoft, error)
	Find() ([]*model.ServerInstallSoft, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerInstallSoft, err error)
	FindInBatches(result *[]*model.ServerInstallSoft, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ServerInstallSoft) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerInstallSoftDo
	Assign(attrs ...field.AssignExpr) IServerInstallSoftDo
	Joins(fields ...field.RelationField) IServerInstallSoftDo
	Preload(fields ...field.RelationField) IServerInstallSoftDo
	FirstOrInit() (*model.ServerInstallSoft, error)
	FirstOrCreate() (*model.ServerInstallSoft, error)
	FindByPage(offset int, limit int) (result []*model.ServerInstallSoft, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerInstallSoftDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverInstallSoftDo) Debug() IServerInstallSoftDo {
	return s.withDO(s.DO.Debug())
}

func (s serverInstallSoftDo) WithContext(ctx context.Context) IServerInstallSoftDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverInstallSoftDo) ReadDB() IServerInstallSoftDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverInstallSoftDo) WriteDB() IServerInstallSoftDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverInstallSoftDo) Session(config *gorm.Session) IServerInstallSoftDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverInstallSoftDo) Clauses(conds ...clause.Expression) IServerInstallSoftDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverInstallSoftDo) Returning(value interface{}, columns ...string) IServerInstallSoftDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverInstallSoftDo) Not(conds ...gen.Condition) IServerInstallSoftDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverInstallSoftDo) Or(conds ...gen.Condition) IServerInstallSoftDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverInstallSoftDo) Select(conds ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverInstallSoftDo) Where(conds ...gen.Condition) IServerInstallSoftDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverInstallSoftDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IServerInstallSoftDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s serverInstallSoftDo) Order(conds ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverInstallSoftDo) Distinct(cols ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverInstallSoftDo) Omit(cols ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverInstallSoftDo) Join(table schema.Tabler, on ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverInstallSoftDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverInstallSoftDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverInstallSoftDo) Group(cols ...field.Expr) IServerInstallSoftDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverInstallSoftDo) Having(conds ...gen.Condition) IServerInstallSoftDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverInstallSoftDo) Limit(limit int) IServerInstallSoftDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverInstallSoftDo) Offset(offset int) IServerInstallSoftDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverInstallSoftDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerInstallSoftDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverInstallSoftDo) Unscoped() IServerInstallSoftDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverInstallSoftDo) Create(values ...*model.ServerInstallSoft) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverInstallSoftDo) CreateInBatches(values []*model.ServerInstallSoft, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverInstallSoftDo) Save(values ...*model.ServerInstallSoft) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverInstallSoftDo) First() (*model.ServerInstallSoft, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerInstallSoft), nil
	}
}

func (s serverInstallSoftDo) Take() (*model.ServerInstallSoft, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerInstallSoft), nil
	}
}

func (s serverInstallSoftDo) Last() (*model.ServerInstallSoft, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerInstallSoft), nil
	}
}

func (s serverInstallSoftDo) Find() ([]*model.ServerInstallSoft, error) {
	result, err := s.DO.Find()
	return result.([]*model.ServerInstallSoft), err
}

func (s serverInstallSoftDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ServerInstallSoft, err error) {
	buf := make([]*model.ServerInstallSoft, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverInstallSoftDo) FindInBatches(result *[]*model.ServerInstallSoft, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverInstallSoftDo) Attrs(attrs ...field.AssignExpr) IServerInstallSoftDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverInstallSoftDo) Assign(attrs ...field.AssignExpr) IServerInstallSoftDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverInstallSoftDo) Joins(fields ...field.RelationField) IServerInstallSoftDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverInstallSoftDo) Preload(fields ...field.RelationField) IServerInstallSoftDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverInstallSoftDo) FirstOrInit() (*model.ServerInstallSoft, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerInstallSoft), nil
	}
}

func (s serverInstallSoftDo) FirstOrCreate() (*model.ServerInstallSoft, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ServerInstallSoft), nil
	}
}

func (s serverInstallSoftDo) FindByPage(offset int, limit int) (result []*model.ServerInstallSoft, count int64, err error) {
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

func (s serverInstallSoftDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverInstallSoftDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverInstallSoftDo) Delete(models ...*model.ServerInstallSoft) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverInstallSoftDo) withDO(do gen.Dao) *serverInstallSoftDo {
	s.DO = *do.(*gen.DO)
	return s
}