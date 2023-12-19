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

func newItems(db *gorm.DB, opts ...gen.DOOption) items {
	_items := items{}

	_items.itemsDo.UseDB(db, opts...)
	_items.itemsDo.UseModel(&model.Items{})

	tableName := _items.itemsDo.TableName()
	_items.ALL = field.NewAsterisk(tableName)
	_items.Itemid = field.NewInt64(tableName, "itemid")
	_items.Type = field.NewInt32(tableName, "type")
	_items.SnmpCommunity = field.NewString(tableName, "snmp_community")
	_items.SnmpOid = field.NewString(tableName, "snmp_oid")
	_items.Hostid = field.NewInt64(tableName, "hostid")
	_items.Name = field.NewString(tableName, "name")
	_items.Key = field.NewString(tableName, "key_")
	_items.Delay = field.NewString(tableName, "delay")
	_items.History = field.NewString(tableName, "history")
	_items.Trends = field.NewString(tableName, "trends")
	_items.Status = field.NewInt32(tableName, "status")
	_items.ValueType = field.NewInt32(tableName, "value_type")
	_items.TrapperHosts = field.NewString(tableName, "trapper_hosts")
	_items.Units = field.NewString(tableName, "units")
	_items.Snmpv3Securityname = field.NewString(tableName, "snmpv3_securityname")
	_items.Snmpv3Securitylevel = field.NewInt32(tableName, "snmpv3_securitylevel")
	_items.Snmpv3Authpassphrase = field.NewString(tableName, "snmpv3_authpassphrase")
	_items.Snmpv3Privpassphrase = field.NewString(tableName, "snmpv3_privpassphrase")
	_items.Formula = field.NewString(tableName, "formula")
	_items.Error = field.NewString(tableName, "error")
	_items.Lastlogsize = field.NewInt64(tableName, "lastlogsize")
	_items.Logtimefmt = field.NewString(tableName, "logtimefmt")
	_items.Templateid = field.NewInt64(tableName, "templateid")
	_items.Valuemapid = field.NewInt64(tableName, "valuemapid")
	_items.Params = field.NewString(tableName, "params")
	_items.IpmiSensor = field.NewString(tableName, "ipmi_sensor")
	_items.Authtype = field.NewInt32(tableName, "authtype")
	_items.Username = field.NewString(tableName, "username")
	_items.Password = field.NewString(tableName, "password")
	_items.Publickey = field.NewString(tableName, "publickey")
	_items.Privatekey = field.NewString(tableName, "privatekey")
	_items.Mtime = field.NewInt32(tableName, "mtime")
	_items.Flags = field.NewInt32(tableName, "flags")
	_items.Interfaceid = field.NewInt64(tableName, "interfaceid")
	_items.Port = field.NewString(tableName, "port")
	_items.Description = field.NewString(tableName, "description")
	_items.InventoryLink = field.NewInt32(tableName, "inventory_link")
	_items.Lifetime = field.NewString(tableName, "lifetime")
	_items.Snmpv3Authprotocol = field.NewInt32(tableName, "snmpv3_authprotocol")
	_items.Snmpv3Privprotocol = field.NewInt32(tableName, "snmpv3_privprotocol")
	_items.State = field.NewInt32(tableName, "state")
	_items.Snmpv3Contextname = field.NewString(tableName, "snmpv3_contextname")
	_items.Evaltype = field.NewInt32(tableName, "evaltype")
	_items.JmxEndpoint = field.NewString(tableName, "jmx_endpoint")
	_items.MasterItemid = field.NewInt64(tableName, "master_itemid")
	_items.Timeout = field.NewString(tableName, "timeout")
	_items.URL = field.NewString(tableName, "url")
	_items.QueryFields = field.NewString(tableName, "query_fields")
	_items.Posts = field.NewString(tableName, "posts")
	_items.StatusCodes = field.NewString(tableName, "status_codes")
	_items.FollowRedirects = field.NewInt32(tableName, "follow_redirects")
	_items.PostType = field.NewInt32(tableName, "post_type")
	_items.HTTPProxy = field.NewString(tableName, "http_proxy")
	_items.Headers = field.NewString(tableName, "headers")
	_items.RetrieveMode = field.NewInt32(tableName, "retrieve_mode")
	_items.RequestMethod = field.NewInt32(tableName, "request_method")
	_items.OutputFormat = field.NewInt32(tableName, "output_format")
	_items.SslCertFile = field.NewString(tableName, "ssl_cert_file")
	_items.SslKeyFile = field.NewString(tableName, "ssl_key_file")
	_items.SslKeyPassword = field.NewString(tableName, "ssl_key_password")
	_items.VerifyPeer = field.NewInt32(tableName, "verify_peer")
	_items.VerifyHost = field.NewInt32(tableName, "verify_host")
	_items.AllowTraps = field.NewInt32(tableName, "allow_traps")

	_items.fillFieldMap()

	return _items
}

type items struct {
	itemsDo

	ALL                  field.Asterisk
	Itemid               field.Int64
	Type                 field.Int32
	SnmpCommunity        field.String
	SnmpOid              field.String
	Hostid               field.Int64
	Name                 field.String
	Key                  field.String
	Delay                field.String
	History              field.String
	Trends               field.String
	Status               field.Int32
	ValueType            field.Int32
	TrapperHosts         field.String
	Units                field.String
	Snmpv3Securityname   field.String
	Snmpv3Securitylevel  field.Int32
	Snmpv3Authpassphrase field.String
	Snmpv3Privpassphrase field.String
	Formula              field.String
	Error                field.String
	Lastlogsize          field.Int64
	Logtimefmt           field.String
	Templateid           field.Int64
	Valuemapid           field.Int64
	Params               field.String
	IpmiSensor           field.String
	Authtype             field.Int32
	Username             field.String
	Password             field.String
	Publickey            field.String
	Privatekey           field.String
	Mtime                field.Int32
	Flags                field.Int32
	Interfaceid          field.Int64
	Port                 field.String
	Description          field.String
	InventoryLink        field.Int32
	Lifetime             field.String
	Snmpv3Authprotocol   field.Int32
	Snmpv3Privprotocol   field.Int32
	State                field.Int32
	Snmpv3Contextname    field.String
	Evaltype             field.Int32
	JmxEndpoint          field.String
	MasterItemid         field.Int64
	Timeout              field.String
	URL                  field.String
	QueryFields          field.String
	Posts                field.String
	StatusCodes          field.String
	FollowRedirects      field.Int32
	PostType             field.Int32
	HTTPProxy            field.String
	Headers              field.String
	RetrieveMode         field.Int32
	RequestMethod        field.Int32
	OutputFormat         field.Int32
	SslCertFile          field.String
	SslKeyFile           field.String
	SslKeyPassword       field.String
	VerifyPeer           field.Int32
	VerifyHost           field.Int32
	AllowTraps           field.Int32

	fieldMap map[string]field.Expr
}

func (i items) Table(newTableName string) *items {
	i.itemsDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i items) As(alias string) *items {
	i.itemsDo.DO = *(i.itemsDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *items) updateTableName(table string) *items {
	i.ALL = field.NewAsterisk(table)
	i.Itemid = field.NewInt64(table, "itemid")
	i.Type = field.NewInt32(table, "type")
	i.SnmpCommunity = field.NewString(table, "snmp_community")
	i.SnmpOid = field.NewString(table, "snmp_oid")
	i.Hostid = field.NewInt64(table, "hostid")
	i.Name = field.NewString(table, "name")
	i.Key = field.NewString(table, "key_")
	i.Delay = field.NewString(table, "delay")
	i.History = field.NewString(table, "history")
	i.Trends = field.NewString(table, "trends")
	i.Status = field.NewInt32(table, "status")
	i.ValueType = field.NewInt32(table, "value_type")
	i.TrapperHosts = field.NewString(table, "trapper_hosts")
	i.Units = field.NewString(table, "units")
	i.Snmpv3Securityname = field.NewString(table, "snmpv3_securityname")
	i.Snmpv3Securitylevel = field.NewInt32(table, "snmpv3_securitylevel")
	i.Snmpv3Authpassphrase = field.NewString(table, "snmpv3_authpassphrase")
	i.Snmpv3Privpassphrase = field.NewString(table, "snmpv3_privpassphrase")
	i.Formula = field.NewString(table, "formula")
	i.Error = field.NewString(table, "error")
	i.Lastlogsize = field.NewInt64(table, "lastlogsize")
	i.Logtimefmt = field.NewString(table, "logtimefmt")
	i.Templateid = field.NewInt64(table, "templateid")
	i.Valuemapid = field.NewInt64(table, "valuemapid")
	i.Params = field.NewString(table, "params")
	i.IpmiSensor = field.NewString(table, "ipmi_sensor")
	i.Authtype = field.NewInt32(table, "authtype")
	i.Username = field.NewString(table, "username")
	i.Password = field.NewString(table, "password")
	i.Publickey = field.NewString(table, "publickey")
	i.Privatekey = field.NewString(table, "privatekey")
	i.Mtime = field.NewInt32(table, "mtime")
	i.Flags = field.NewInt32(table, "flags")
	i.Interfaceid = field.NewInt64(table, "interfaceid")
	i.Port = field.NewString(table, "port")
	i.Description = field.NewString(table, "description")
	i.InventoryLink = field.NewInt32(table, "inventory_link")
	i.Lifetime = field.NewString(table, "lifetime")
	i.Snmpv3Authprotocol = field.NewInt32(table, "snmpv3_authprotocol")
	i.Snmpv3Privprotocol = field.NewInt32(table, "snmpv3_privprotocol")
	i.State = field.NewInt32(table, "state")
	i.Snmpv3Contextname = field.NewString(table, "snmpv3_contextname")
	i.Evaltype = field.NewInt32(table, "evaltype")
	i.JmxEndpoint = field.NewString(table, "jmx_endpoint")
	i.MasterItemid = field.NewInt64(table, "master_itemid")
	i.Timeout = field.NewString(table, "timeout")
	i.URL = field.NewString(table, "url")
	i.QueryFields = field.NewString(table, "query_fields")
	i.Posts = field.NewString(table, "posts")
	i.StatusCodes = field.NewString(table, "status_codes")
	i.FollowRedirects = field.NewInt32(table, "follow_redirects")
	i.PostType = field.NewInt32(table, "post_type")
	i.HTTPProxy = field.NewString(table, "http_proxy")
	i.Headers = field.NewString(table, "headers")
	i.RetrieveMode = field.NewInt32(table, "retrieve_mode")
	i.RequestMethod = field.NewInt32(table, "request_method")
	i.OutputFormat = field.NewInt32(table, "output_format")
	i.SslCertFile = field.NewString(table, "ssl_cert_file")
	i.SslKeyFile = field.NewString(table, "ssl_key_file")
	i.SslKeyPassword = field.NewString(table, "ssl_key_password")
	i.VerifyPeer = field.NewInt32(table, "verify_peer")
	i.VerifyHost = field.NewInt32(table, "verify_host")
	i.AllowTraps = field.NewInt32(table, "allow_traps")

	i.fillFieldMap()

	return i
}

func (i *items) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *items) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 63)
	i.fieldMap["itemid"] = i.Itemid
	i.fieldMap["type"] = i.Type
	i.fieldMap["snmp_community"] = i.SnmpCommunity
	i.fieldMap["snmp_oid"] = i.SnmpOid
	i.fieldMap["hostid"] = i.Hostid
	i.fieldMap["name"] = i.Name
	i.fieldMap["key_"] = i.Key
	i.fieldMap["delay"] = i.Delay
	i.fieldMap["history"] = i.History
	i.fieldMap["trends"] = i.Trends
	i.fieldMap["status"] = i.Status
	i.fieldMap["value_type"] = i.ValueType
	i.fieldMap["trapper_hosts"] = i.TrapperHosts
	i.fieldMap["units"] = i.Units
	i.fieldMap["snmpv3_securityname"] = i.Snmpv3Securityname
	i.fieldMap["snmpv3_securitylevel"] = i.Snmpv3Securitylevel
	i.fieldMap["snmpv3_authpassphrase"] = i.Snmpv3Authpassphrase
	i.fieldMap["snmpv3_privpassphrase"] = i.Snmpv3Privpassphrase
	i.fieldMap["formula"] = i.Formula
	i.fieldMap["error"] = i.Error
	i.fieldMap["lastlogsize"] = i.Lastlogsize
	i.fieldMap["logtimefmt"] = i.Logtimefmt
	i.fieldMap["templateid"] = i.Templateid
	i.fieldMap["valuemapid"] = i.Valuemapid
	i.fieldMap["params"] = i.Params
	i.fieldMap["ipmi_sensor"] = i.IpmiSensor
	i.fieldMap["authtype"] = i.Authtype
	i.fieldMap["username"] = i.Username
	i.fieldMap["password"] = i.Password
	i.fieldMap["publickey"] = i.Publickey
	i.fieldMap["privatekey"] = i.Privatekey
	i.fieldMap["mtime"] = i.Mtime
	i.fieldMap["flags"] = i.Flags
	i.fieldMap["interfaceid"] = i.Interfaceid
	i.fieldMap["port"] = i.Port
	i.fieldMap["description"] = i.Description
	i.fieldMap["inventory_link"] = i.InventoryLink
	i.fieldMap["lifetime"] = i.Lifetime
	i.fieldMap["snmpv3_authprotocol"] = i.Snmpv3Authprotocol
	i.fieldMap["snmpv3_privprotocol"] = i.Snmpv3Privprotocol
	i.fieldMap["state"] = i.State
	i.fieldMap["snmpv3_contextname"] = i.Snmpv3Contextname
	i.fieldMap["evaltype"] = i.Evaltype
	i.fieldMap["jmx_endpoint"] = i.JmxEndpoint
	i.fieldMap["master_itemid"] = i.MasterItemid
	i.fieldMap["timeout"] = i.Timeout
	i.fieldMap["url"] = i.URL
	i.fieldMap["query_fields"] = i.QueryFields
	i.fieldMap["posts"] = i.Posts
	i.fieldMap["status_codes"] = i.StatusCodes
	i.fieldMap["follow_redirects"] = i.FollowRedirects
	i.fieldMap["post_type"] = i.PostType
	i.fieldMap["http_proxy"] = i.HTTPProxy
	i.fieldMap["headers"] = i.Headers
	i.fieldMap["retrieve_mode"] = i.RetrieveMode
	i.fieldMap["request_method"] = i.RequestMethod
	i.fieldMap["output_format"] = i.OutputFormat
	i.fieldMap["ssl_cert_file"] = i.SslCertFile
	i.fieldMap["ssl_key_file"] = i.SslKeyFile
	i.fieldMap["ssl_key_password"] = i.SslKeyPassword
	i.fieldMap["verify_peer"] = i.VerifyPeer
	i.fieldMap["verify_host"] = i.VerifyHost
	i.fieldMap["allow_traps"] = i.AllowTraps
}

func (i items) clone(db *gorm.DB) items {
	i.itemsDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i items) replaceDB(db *gorm.DB) items {
	i.itemsDo.ReplaceDB(db)
	return i
}

type itemsDo struct{ gen.DO }

type IItemsDo interface {
	gen.SubQuery
	Debug() IItemsDo
	WithContext(ctx context.Context) IItemsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IItemsDo
	WriteDB() IItemsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IItemsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IItemsDo
	Not(conds ...gen.Condition) IItemsDo
	Or(conds ...gen.Condition) IItemsDo
	Select(conds ...field.Expr) IItemsDo
	Where(conds ...gen.Condition) IItemsDo
	Order(conds ...field.Expr) IItemsDo
	Distinct(cols ...field.Expr) IItemsDo
	Omit(cols ...field.Expr) IItemsDo
	Join(table schema.Tabler, on ...field.Expr) IItemsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IItemsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IItemsDo
	Group(cols ...field.Expr) IItemsDo
	Having(conds ...gen.Condition) IItemsDo
	Limit(limit int) IItemsDo
	Offset(offset int) IItemsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IItemsDo
	Unscoped() IItemsDo
	Create(values ...*model.Items) error
	CreateInBatches(values []*model.Items, batchSize int) error
	Save(values ...*model.Items) error
	First() (*model.Items, error)
	Take() (*model.Items, error)
	Last() (*model.Items, error)
	Find() ([]*model.Items, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Items, err error)
	FindInBatches(result *[]*model.Items, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Items) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IItemsDo
	Assign(attrs ...field.AssignExpr) IItemsDo
	Joins(fields ...field.RelationField) IItemsDo
	Preload(fields ...field.RelationField) IItemsDo
	FirstOrInit() (*model.Items, error)
	FirstOrCreate() (*model.Items, error)
	FindByPage(offset int, limit int) (result []*model.Items, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IItemsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i itemsDo) Debug() IItemsDo {
	return i.withDO(i.DO.Debug())
}

func (i itemsDo) WithContext(ctx context.Context) IItemsDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i itemsDo) ReadDB() IItemsDo {
	return i.Clauses(dbresolver.Read)
}

func (i itemsDo) WriteDB() IItemsDo {
	return i.Clauses(dbresolver.Write)
}

func (i itemsDo) Session(config *gorm.Session) IItemsDo {
	return i.withDO(i.DO.Session(config))
}

func (i itemsDo) Clauses(conds ...clause.Expression) IItemsDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i itemsDo) Returning(value interface{}, columns ...string) IItemsDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i itemsDo) Not(conds ...gen.Condition) IItemsDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i itemsDo) Or(conds ...gen.Condition) IItemsDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i itemsDo) Select(conds ...field.Expr) IItemsDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i itemsDo) Where(conds ...gen.Condition) IItemsDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i itemsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IItemsDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i itemsDo) Order(conds ...field.Expr) IItemsDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i itemsDo) Distinct(cols ...field.Expr) IItemsDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i itemsDo) Omit(cols ...field.Expr) IItemsDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i itemsDo) Join(table schema.Tabler, on ...field.Expr) IItemsDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i itemsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IItemsDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i itemsDo) RightJoin(table schema.Tabler, on ...field.Expr) IItemsDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i itemsDo) Group(cols ...field.Expr) IItemsDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i itemsDo) Having(conds ...gen.Condition) IItemsDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i itemsDo) Limit(limit int) IItemsDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i itemsDo) Offset(offset int) IItemsDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i itemsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IItemsDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i itemsDo) Unscoped() IItemsDo {
	return i.withDO(i.DO.Unscoped())
}

func (i itemsDo) Create(values ...*model.Items) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i itemsDo) CreateInBatches(values []*model.Items, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i itemsDo) Save(values ...*model.Items) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i itemsDo) First() (*model.Items, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Items), nil
	}
}

func (i itemsDo) Take() (*model.Items, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Items), nil
	}
}

func (i itemsDo) Last() (*model.Items, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Items), nil
	}
}

func (i itemsDo) Find() ([]*model.Items, error) {
	result, err := i.DO.Find()
	return result.([]*model.Items), err
}

func (i itemsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Items, err error) {
	buf := make([]*model.Items, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i itemsDo) FindInBatches(result *[]*model.Items, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i itemsDo) Attrs(attrs ...field.AssignExpr) IItemsDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i itemsDo) Assign(attrs ...field.AssignExpr) IItemsDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i itemsDo) Joins(fields ...field.RelationField) IItemsDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i itemsDo) Preload(fields ...field.RelationField) IItemsDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i itemsDo) FirstOrInit() (*model.Items, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Items), nil
	}
}

func (i itemsDo) FirstOrCreate() (*model.Items, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Items), nil
	}
}

func (i itemsDo) FindByPage(offset int, limit int) (result []*model.Items, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i itemsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i itemsDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i itemsDo) Delete(models ...*model.Items) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *itemsDo) withDO(do gen.Dao) *itemsDo {
	i.DO = *do.(*gen.DO)
	return i
}
