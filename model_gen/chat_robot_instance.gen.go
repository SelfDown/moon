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

func newChatRobotInstance(db *gorm.DB, opts ...gen.DOOption) chatRobotInstance {
	_chatRobotInstance := chatRobotInstance{}

	_chatRobotInstance.chatRobotInstanceDo.UseDB(db, opts...)
	_chatRobotInstance.chatRobotInstanceDo.UseModel(&model.ChatRobotInstance{})

	tableName := _chatRobotInstance.chatRobotInstanceDo.TableName()
	_chatRobotInstance.ALL = field.NewAsterisk(tableName)
	_chatRobotInstance.RobotID = field.NewString(tableName, "robot_id")
	_chatRobotInstance.SysProjectID = field.NewString(tableName, "sys_project_id")
	_chatRobotInstance.ServerID = field.NewString(tableName, "server_id")
	_chatRobotInstance.RobotLoginPwd = field.NewString(tableName, "robot_login_pwd")
	_chatRobotInstance.CreateTime = field.NewTime(tableName, "create_time")
	_chatRobotInstance.RobotHome = field.NewString(tableName, "robot_home")
	_chatRobotInstance.RobotStartupCmd = field.NewString(tableName, "robot_startup_cmd")
	_chatRobotInstance.RobotSendMessageAPI = field.NewString(tableName, "robot_send_message_api")
	_chatRobotInstance.LoginQcodeBase64 = field.NewString(tableName, "login_qcode_base64")
	_chatRobotInstance.StateChange = field.NewString(tableName, "state_change")
	_chatRobotInstance.Comments = field.NewString(tableName, "comments")
	_chatRobotInstance.ClientLastReportTime = field.NewTime(tableName, "client_last_report_time")
	_chatRobotInstance.RobotShutdownCmd = field.NewString(tableName, "robot_shutdown_cmd")
	_chatRobotInstance.RobotClientVersion = field.NewString(tableName, "robot_client_version")
	_chatRobotInstance.QcodePath = field.NewString(tableName, "qcode_path")
	_chatRobotInstance.RobotType = field.NewString(tableName, "robot_type")
	_chatRobotInstance.ModifyTime = field.NewTime(tableName, "modify_time")
	_chatRobotInstance.DingtalkSendTime = field.NewTime(tableName, "dingtalk_send_time")
	_chatRobotInstance.MessageSentIP = field.NewString(tableName, "message_sent_ip")
	_chatRobotInstance.MessageSentSql = field.NewString(tableName, "message_sent_sql")
	_chatRobotInstance.MessageSentUser = field.NewString(tableName, "message_sent_user")
	_chatRobotInstance.MessageSentPwd = field.NewString(tableName, "message_sent_pwd")
	_chatRobotInstance.SSHSentPort = field.NewString(tableName, "ssh_sent_port")
	_chatRobotInstance.MessageSentPort = field.NewString(tableName, "message_sent_port")

	_chatRobotInstance.fillFieldMap()

	return _chatRobotInstance
}

type chatRobotInstance struct {
	chatRobotInstanceDo

	ALL          field.Asterisk
	RobotID      field.String // 机器人账号
	SysProjectID field.String
	ServerID     field.String
	/*
		机器人登录密码

	*/
	RobotLoginPwd       field.String
	CreateTime          field.Time   // 记录创建时间（数据库自动写入）
	RobotHome           field.String // 机器人安装主目录
	RobotStartupCmd     field.String // 机器人启动命令
	RobotSendMessageAPI field.String // 机器人消息发送接口
	LoginQcodeBase64    field.String // 二维码的base64编码
	/*
		客户端状态变化	旧的状态，新的状态 参见 [客户端状态说明](API.md#客户端运行状态介绍）
		0 离线
		1 在线
		其他的 由上报接口写入
	*/
	StateChange          field.String
	Comments             field.String // 备注说明
	ClientLastReportTime field.Time   // 接受到的客户端POST上报最近一次时间
	RobotShutdownCmd     field.String // 机器人关闭命令
	/*
		机器人客户端版本

	*/
	RobotClientVersion field.String
	/*
		二维码存放路径
		/static/staticuploadsqcode/{{chat_robot_id}}.jpg
	*/
	QcodePath field.String
	/*
		0 是QQ机器人
		1 是微信机器人
	*/
	RobotType        field.String
	ModifyTime       field.Time // 记录修改时间（数据库自动写入）
	DingtalkSendTime field.Time
	MessageSentIP    field.String // 发送消息的服务器ip
	MessageSentSql   field.String // 消息发送脚本
	MessageSentUser  field.String // 消息发送接口登录用户名
	MessageSentPwd   field.String // 消息发送接口登录密码
	SSHSentPort      field.String // 端口
	MessageSentPort  field.String // 端口

	fieldMap map[string]field.Expr
}

func (c chatRobotInstance) Table(newTableName string) *chatRobotInstance {
	c.chatRobotInstanceDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chatRobotInstance) As(alias string) *chatRobotInstance {
	c.chatRobotInstanceDo.DO = *(c.chatRobotInstanceDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chatRobotInstance) updateTableName(table string) *chatRobotInstance {
	c.ALL = field.NewAsterisk(table)
	c.RobotID = field.NewString(table, "robot_id")
	c.SysProjectID = field.NewString(table, "sys_project_id")
	c.ServerID = field.NewString(table, "server_id")
	c.RobotLoginPwd = field.NewString(table, "robot_login_pwd")
	c.CreateTime = field.NewTime(table, "create_time")
	c.RobotHome = field.NewString(table, "robot_home")
	c.RobotStartupCmd = field.NewString(table, "robot_startup_cmd")
	c.RobotSendMessageAPI = field.NewString(table, "robot_send_message_api")
	c.LoginQcodeBase64 = field.NewString(table, "login_qcode_base64")
	c.StateChange = field.NewString(table, "state_change")
	c.Comments = field.NewString(table, "comments")
	c.ClientLastReportTime = field.NewTime(table, "client_last_report_time")
	c.RobotShutdownCmd = field.NewString(table, "robot_shutdown_cmd")
	c.RobotClientVersion = field.NewString(table, "robot_client_version")
	c.QcodePath = field.NewString(table, "qcode_path")
	c.RobotType = field.NewString(table, "robot_type")
	c.ModifyTime = field.NewTime(table, "modify_time")
	c.DingtalkSendTime = field.NewTime(table, "dingtalk_send_time")
	c.MessageSentIP = field.NewString(table, "message_sent_ip")
	c.MessageSentSql = field.NewString(table, "message_sent_sql")
	c.MessageSentUser = field.NewString(table, "message_sent_user")
	c.MessageSentPwd = field.NewString(table, "message_sent_pwd")
	c.SSHSentPort = field.NewString(table, "ssh_sent_port")
	c.MessageSentPort = field.NewString(table, "message_sent_port")

	c.fillFieldMap()

	return c
}

func (c *chatRobotInstance) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chatRobotInstance) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 24)
	c.fieldMap["robot_id"] = c.RobotID
	c.fieldMap["sys_project_id"] = c.SysProjectID
	c.fieldMap["server_id"] = c.ServerID
	c.fieldMap["robot_login_pwd"] = c.RobotLoginPwd
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["robot_home"] = c.RobotHome
	c.fieldMap["robot_startup_cmd"] = c.RobotStartupCmd
	c.fieldMap["robot_send_message_api"] = c.RobotSendMessageAPI
	c.fieldMap["login_qcode_base64"] = c.LoginQcodeBase64
	c.fieldMap["state_change"] = c.StateChange
	c.fieldMap["comments"] = c.Comments
	c.fieldMap["client_last_report_time"] = c.ClientLastReportTime
	c.fieldMap["robot_shutdown_cmd"] = c.RobotShutdownCmd
	c.fieldMap["robot_client_version"] = c.RobotClientVersion
	c.fieldMap["qcode_path"] = c.QcodePath
	c.fieldMap["robot_type"] = c.RobotType
	c.fieldMap["modify_time"] = c.ModifyTime
	c.fieldMap["dingtalk_send_time"] = c.DingtalkSendTime
	c.fieldMap["message_sent_ip"] = c.MessageSentIP
	c.fieldMap["message_sent_sql"] = c.MessageSentSql
	c.fieldMap["message_sent_user"] = c.MessageSentUser
	c.fieldMap["message_sent_pwd"] = c.MessageSentPwd
	c.fieldMap["ssh_sent_port"] = c.SSHSentPort
	c.fieldMap["message_sent_port"] = c.MessageSentPort
}

func (c chatRobotInstance) clone(db *gorm.DB) chatRobotInstance {
	c.chatRobotInstanceDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chatRobotInstance) replaceDB(db *gorm.DB) chatRobotInstance {
	c.chatRobotInstanceDo.ReplaceDB(db)
	return c
}

type chatRobotInstanceDo struct{ gen.DO }

type IChatRobotInstanceDo interface {
	gen.SubQuery
	Debug() IChatRobotInstanceDo
	WithContext(ctx context.Context) IChatRobotInstanceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IChatRobotInstanceDo
	WriteDB() IChatRobotInstanceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IChatRobotInstanceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IChatRobotInstanceDo
	Not(conds ...gen.Condition) IChatRobotInstanceDo
	Or(conds ...gen.Condition) IChatRobotInstanceDo
	Select(conds ...field.Expr) IChatRobotInstanceDo
	Where(conds ...gen.Condition) IChatRobotInstanceDo
	Order(conds ...field.Expr) IChatRobotInstanceDo
	Distinct(cols ...field.Expr) IChatRobotInstanceDo
	Omit(cols ...field.Expr) IChatRobotInstanceDo
	Join(table schema.Tabler, on ...field.Expr) IChatRobotInstanceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IChatRobotInstanceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IChatRobotInstanceDo
	Group(cols ...field.Expr) IChatRobotInstanceDo
	Having(conds ...gen.Condition) IChatRobotInstanceDo
	Limit(limit int) IChatRobotInstanceDo
	Offset(offset int) IChatRobotInstanceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IChatRobotInstanceDo
	Unscoped() IChatRobotInstanceDo
	Create(values ...*model.ChatRobotInstance) error
	CreateInBatches(values []*model.ChatRobotInstance, batchSize int) error
	Save(values ...*model.ChatRobotInstance) error
	First() (*model.ChatRobotInstance, error)
	Take() (*model.ChatRobotInstance, error)
	Last() (*model.ChatRobotInstance, error)
	Find() ([]*model.ChatRobotInstance, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChatRobotInstance, err error)
	FindInBatches(result *[]*model.ChatRobotInstance, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ChatRobotInstance) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IChatRobotInstanceDo
	Assign(attrs ...field.AssignExpr) IChatRobotInstanceDo
	Joins(fields ...field.RelationField) IChatRobotInstanceDo
	Preload(fields ...field.RelationField) IChatRobotInstanceDo
	FirstOrInit() (*model.ChatRobotInstance, error)
	FirstOrCreate() (*model.ChatRobotInstance, error)
	FindByPage(offset int, limit int) (result []*model.ChatRobotInstance, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IChatRobotInstanceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c chatRobotInstanceDo) Debug() IChatRobotInstanceDo {
	return c.withDO(c.DO.Debug())
}

func (c chatRobotInstanceDo) WithContext(ctx context.Context) IChatRobotInstanceDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chatRobotInstanceDo) ReadDB() IChatRobotInstanceDo {
	return c.Clauses(dbresolver.Read)
}

func (c chatRobotInstanceDo) WriteDB() IChatRobotInstanceDo {
	return c.Clauses(dbresolver.Write)
}

func (c chatRobotInstanceDo) Session(config *gorm.Session) IChatRobotInstanceDo {
	return c.withDO(c.DO.Session(config))
}

func (c chatRobotInstanceDo) Clauses(conds ...clause.Expression) IChatRobotInstanceDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chatRobotInstanceDo) Returning(value interface{}, columns ...string) IChatRobotInstanceDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chatRobotInstanceDo) Not(conds ...gen.Condition) IChatRobotInstanceDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chatRobotInstanceDo) Or(conds ...gen.Condition) IChatRobotInstanceDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chatRobotInstanceDo) Select(conds ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chatRobotInstanceDo) Where(conds ...gen.Condition) IChatRobotInstanceDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chatRobotInstanceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IChatRobotInstanceDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chatRobotInstanceDo) Order(conds ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chatRobotInstanceDo) Distinct(cols ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chatRobotInstanceDo) Omit(cols ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chatRobotInstanceDo) Join(table schema.Tabler, on ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chatRobotInstanceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chatRobotInstanceDo) RightJoin(table schema.Tabler, on ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chatRobotInstanceDo) Group(cols ...field.Expr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chatRobotInstanceDo) Having(conds ...gen.Condition) IChatRobotInstanceDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chatRobotInstanceDo) Limit(limit int) IChatRobotInstanceDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chatRobotInstanceDo) Offset(offset int) IChatRobotInstanceDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chatRobotInstanceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IChatRobotInstanceDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chatRobotInstanceDo) Unscoped() IChatRobotInstanceDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chatRobotInstanceDo) Create(values ...*model.ChatRobotInstance) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chatRobotInstanceDo) CreateInBatches(values []*model.ChatRobotInstance, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chatRobotInstanceDo) Save(values ...*model.ChatRobotInstance) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chatRobotInstanceDo) First() (*model.ChatRobotInstance, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatRobotInstance), nil
	}
}

func (c chatRobotInstanceDo) Take() (*model.ChatRobotInstance, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatRobotInstance), nil
	}
}

func (c chatRobotInstanceDo) Last() (*model.ChatRobotInstance, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatRobotInstance), nil
	}
}

func (c chatRobotInstanceDo) Find() ([]*model.ChatRobotInstance, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChatRobotInstance), err
}

func (c chatRobotInstanceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChatRobotInstance, err error) {
	buf := make([]*model.ChatRobotInstance, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chatRobotInstanceDo) FindInBatches(result *[]*model.ChatRobotInstance, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chatRobotInstanceDo) Attrs(attrs ...field.AssignExpr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chatRobotInstanceDo) Assign(attrs ...field.AssignExpr) IChatRobotInstanceDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chatRobotInstanceDo) Joins(fields ...field.RelationField) IChatRobotInstanceDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chatRobotInstanceDo) Preload(fields ...field.RelationField) IChatRobotInstanceDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chatRobotInstanceDo) FirstOrInit() (*model.ChatRobotInstance, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatRobotInstance), nil
	}
}

func (c chatRobotInstanceDo) FirstOrCreate() (*model.ChatRobotInstance, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChatRobotInstance), nil
	}
}

func (c chatRobotInstanceDo) FindByPage(offset int, limit int) (result []*model.ChatRobotInstance, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c chatRobotInstanceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chatRobotInstanceDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chatRobotInstanceDo) Delete(models ...*model.ChatRobotInstance) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chatRobotInstanceDo) withDO(do gen.Dao) *chatRobotInstanceDo {
	c.DO = *do.(*gen.DO)
	return c
}
