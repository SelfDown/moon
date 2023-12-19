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

func newZabbixMaintenanceList(db *gorm.DB, opts ...gen.DOOption) zabbixMaintenanceList {
	_zabbixMaintenanceList := zabbixMaintenanceList{}

	_zabbixMaintenanceList.zabbixMaintenanceListDo.UseDB(db, opts...)
	_zabbixMaintenanceList.zabbixMaintenanceListDo.UseModel(&model.ZabbixMaintenanceList{})

	tableName := _zabbixMaintenanceList.zabbixMaintenanceListDo.TableName()
	_zabbixMaintenanceList.ALL = field.NewAsterisk(tableName)
	_zabbixMaintenanceList.MaintenanceID = field.NewInt64(tableName, "maintenance_id")
	_zabbixMaintenanceList.MaintenanceName = field.NewString(tableName, "maintenance_name")
	_zabbixMaintenanceList.MaintenanceStatus = field.NewInt32(tableName, "maintenance_status")
	_zabbixMaintenanceList.MaintenanceSysProjectID = field.NewString(tableName, "maintenance_sys_project_id")
	_zabbixMaintenanceList.MaintenanceActiveSince = field.NewTime(tableName, "maintenance_active_since")
	_zabbixMaintenanceList.MaintenanceActiveTill = field.NewTime(tableName, "maintenance_active_till")
	_zabbixMaintenanceList.CreateTime = field.NewTime(tableName, "create_time")
	_zabbixMaintenanceList.ModifyTime = field.NewTime(tableName, "modify_time")
	_zabbixMaintenanceList.Comments = field.NewString(tableName, "comments")

	_zabbixMaintenanceList.fillFieldMap()

	return _zabbixMaintenanceList
}

type zabbixMaintenanceList struct {
	zabbixMaintenanceListDo

	ALL                     field.Asterisk
	MaintenanceID           field.Int64  // 维护周期id
	MaintenanceName         field.String // 维护周期名字
	MaintenanceStatus       field.Int32  // 维护周期状态 1 启用，0 停用
	MaintenanceSysProjectID field.String // 要创建维护周期的项目点
	MaintenanceActiveSince  field.Time   // 维护周期开始时间
	MaintenanceActiveTill   field.Time
	CreateTime              field.Time
	ModifyTime              field.Time   // 记录修改时间（数据库自动写入）
	Comments                field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (z zabbixMaintenanceList) Table(newTableName string) *zabbixMaintenanceList {
	z.zabbixMaintenanceListDo.UseTable(newTableName)
	return z.updateTableName(newTableName)
}

func (z zabbixMaintenanceList) As(alias string) *zabbixMaintenanceList {
	z.zabbixMaintenanceListDo.DO = *(z.zabbixMaintenanceListDo.As(alias).(*gen.DO))
	return z.updateTableName(alias)
}

func (z *zabbixMaintenanceList) updateTableName(table string) *zabbixMaintenanceList {
	z.ALL = field.NewAsterisk(table)
	z.MaintenanceID = field.NewInt64(table, "maintenance_id")
	z.MaintenanceName = field.NewString(table, "maintenance_name")
	z.MaintenanceStatus = field.NewInt32(table, "maintenance_status")
	z.MaintenanceSysProjectID = field.NewString(table, "maintenance_sys_project_id")
	z.MaintenanceActiveSince = field.NewTime(table, "maintenance_active_since")
	z.MaintenanceActiveTill = field.NewTime(table, "maintenance_active_till")
	z.CreateTime = field.NewTime(table, "create_time")
	z.ModifyTime = field.NewTime(table, "modify_time")
	z.Comments = field.NewString(table, "comments")

	z.fillFieldMap()

	return z
}

func (z *zabbixMaintenanceList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := z.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (z *zabbixMaintenanceList) fillFieldMap() {
	z.fieldMap = make(map[string]field.Expr, 9)
	z.fieldMap["maintenance_id"] = z.MaintenanceID
	z.fieldMap["maintenance_name"] = z.MaintenanceName
	z.fieldMap["maintenance_status"] = z.MaintenanceStatus
	z.fieldMap["maintenance_sys_project_id"] = z.MaintenanceSysProjectID
	z.fieldMap["maintenance_active_since"] = z.MaintenanceActiveSince
	z.fieldMap["maintenance_active_till"] = z.MaintenanceActiveTill
	z.fieldMap["create_time"] = z.CreateTime
	z.fieldMap["modify_time"] = z.ModifyTime
	z.fieldMap["comments"] = z.Comments
}

func (z zabbixMaintenanceList) clone(db *gorm.DB) zabbixMaintenanceList {
	z.zabbixMaintenanceListDo.ReplaceConnPool(db.Statement.ConnPool)
	return z
}

func (z zabbixMaintenanceList) replaceDB(db *gorm.DB) zabbixMaintenanceList {
	z.zabbixMaintenanceListDo.ReplaceDB(db)
	return z
}

type zabbixMaintenanceListDo struct{ gen.DO }

type IZabbixMaintenanceListDo interface {
	gen.SubQuery
	Debug() IZabbixMaintenanceListDo
	WithContext(ctx context.Context) IZabbixMaintenanceListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IZabbixMaintenanceListDo
	WriteDB() IZabbixMaintenanceListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IZabbixMaintenanceListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IZabbixMaintenanceListDo
	Not(conds ...gen.Condition) IZabbixMaintenanceListDo
	Or(conds ...gen.Condition) IZabbixMaintenanceListDo
	Select(conds ...field.Expr) IZabbixMaintenanceListDo
	Where(conds ...gen.Condition) IZabbixMaintenanceListDo
	Order(conds ...field.Expr) IZabbixMaintenanceListDo
	Distinct(cols ...field.Expr) IZabbixMaintenanceListDo
	Omit(cols ...field.Expr) IZabbixMaintenanceListDo
	Join(table schema.Tabler, on ...field.Expr) IZabbixMaintenanceListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IZabbixMaintenanceListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IZabbixMaintenanceListDo
	Group(cols ...field.Expr) IZabbixMaintenanceListDo
	Having(conds ...gen.Condition) IZabbixMaintenanceListDo
	Limit(limit int) IZabbixMaintenanceListDo
	Offset(offset int) IZabbixMaintenanceListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IZabbixMaintenanceListDo
	Unscoped() IZabbixMaintenanceListDo
	Create(values ...*model.ZabbixMaintenanceList) error
	CreateInBatches(values []*model.ZabbixMaintenanceList, batchSize int) error
	Save(values ...*model.ZabbixMaintenanceList) error
	First() (*model.ZabbixMaintenanceList, error)
	Take() (*model.ZabbixMaintenanceList, error)
	Last() (*model.ZabbixMaintenanceList, error)
	Find() ([]*model.ZabbixMaintenanceList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ZabbixMaintenanceList, err error)
	FindInBatches(result *[]*model.ZabbixMaintenanceList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ZabbixMaintenanceList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IZabbixMaintenanceListDo
	Assign(attrs ...field.AssignExpr) IZabbixMaintenanceListDo
	Joins(fields ...field.RelationField) IZabbixMaintenanceListDo
	Preload(fields ...field.RelationField) IZabbixMaintenanceListDo
	FirstOrInit() (*model.ZabbixMaintenanceList, error)
	FirstOrCreate() (*model.ZabbixMaintenanceList, error)
	FindByPage(offset int, limit int) (result []*model.ZabbixMaintenanceList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IZabbixMaintenanceListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (z zabbixMaintenanceListDo) Debug() IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Debug())
}

func (z zabbixMaintenanceListDo) WithContext(ctx context.Context) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.WithContext(ctx))
}

func (z zabbixMaintenanceListDo) ReadDB() IZabbixMaintenanceListDo {
	return z.Clauses(dbresolver.Read)
}

func (z zabbixMaintenanceListDo) WriteDB() IZabbixMaintenanceListDo {
	return z.Clauses(dbresolver.Write)
}

func (z zabbixMaintenanceListDo) Session(config *gorm.Session) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Session(config))
}

func (z zabbixMaintenanceListDo) Clauses(conds ...clause.Expression) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Clauses(conds...))
}

func (z zabbixMaintenanceListDo) Returning(value interface{}, columns ...string) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Returning(value, columns...))
}

func (z zabbixMaintenanceListDo) Not(conds ...gen.Condition) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Not(conds...))
}

func (z zabbixMaintenanceListDo) Or(conds ...gen.Condition) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Or(conds...))
}

func (z zabbixMaintenanceListDo) Select(conds ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Select(conds...))
}

func (z zabbixMaintenanceListDo) Where(conds ...gen.Condition) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Where(conds...))
}

func (z zabbixMaintenanceListDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IZabbixMaintenanceListDo {
	return z.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (z zabbixMaintenanceListDo) Order(conds ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Order(conds...))
}

func (z zabbixMaintenanceListDo) Distinct(cols ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Distinct(cols...))
}

func (z zabbixMaintenanceListDo) Omit(cols ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Omit(cols...))
}

func (z zabbixMaintenanceListDo) Join(table schema.Tabler, on ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Join(table, on...))
}

func (z zabbixMaintenanceListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.LeftJoin(table, on...))
}

func (z zabbixMaintenanceListDo) RightJoin(table schema.Tabler, on ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.RightJoin(table, on...))
}

func (z zabbixMaintenanceListDo) Group(cols ...field.Expr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Group(cols...))
}

func (z zabbixMaintenanceListDo) Having(conds ...gen.Condition) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Having(conds...))
}

func (z zabbixMaintenanceListDo) Limit(limit int) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Limit(limit))
}

func (z zabbixMaintenanceListDo) Offset(offset int) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Offset(offset))
}

func (z zabbixMaintenanceListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Scopes(funcs...))
}

func (z zabbixMaintenanceListDo) Unscoped() IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Unscoped())
}

func (z zabbixMaintenanceListDo) Create(values ...*model.ZabbixMaintenanceList) error {
	if len(values) == 0 {
		return nil
	}
	return z.DO.Create(values)
}

func (z zabbixMaintenanceListDo) CreateInBatches(values []*model.ZabbixMaintenanceList, batchSize int) error {
	return z.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (z zabbixMaintenanceListDo) Save(values ...*model.ZabbixMaintenanceList) error {
	if len(values) == 0 {
		return nil
	}
	return z.DO.Save(values)
}

func (z zabbixMaintenanceListDo) First() (*model.ZabbixMaintenanceList, error) {
	if result, err := z.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ZabbixMaintenanceList), nil
	}
}

func (z zabbixMaintenanceListDo) Take() (*model.ZabbixMaintenanceList, error) {
	if result, err := z.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ZabbixMaintenanceList), nil
	}
}

func (z zabbixMaintenanceListDo) Last() (*model.ZabbixMaintenanceList, error) {
	if result, err := z.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ZabbixMaintenanceList), nil
	}
}

func (z zabbixMaintenanceListDo) Find() ([]*model.ZabbixMaintenanceList, error) {
	result, err := z.DO.Find()
	return result.([]*model.ZabbixMaintenanceList), err
}

func (z zabbixMaintenanceListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ZabbixMaintenanceList, err error) {
	buf := make([]*model.ZabbixMaintenanceList, 0, batchSize)
	err = z.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (z zabbixMaintenanceListDo) FindInBatches(result *[]*model.ZabbixMaintenanceList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return z.DO.FindInBatches(result, batchSize, fc)
}

func (z zabbixMaintenanceListDo) Attrs(attrs ...field.AssignExpr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Attrs(attrs...))
}

func (z zabbixMaintenanceListDo) Assign(attrs ...field.AssignExpr) IZabbixMaintenanceListDo {
	return z.withDO(z.DO.Assign(attrs...))
}

func (z zabbixMaintenanceListDo) Joins(fields ...field.RelationField) IZabbixMaintenanceListDo {
	for _, _f := range fields {
		z = *z.withDO(z.DO.Joins(_f))
	}
	return &z
}

func (z zabbixMaintenanceListDo) Preload(fields ...field.RelationField) IZabbixMaintenanceListDo {
	for _, _f := range fields {
		z = *z.withDO(z.DO.Preload(_f))
	}
	return &z
}

func (z zabbixMaintenanceListDo) FirstOrInit() (*model.ZabbixMaintenanceList, error) {
	if result, err := z.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ZabbixMaintenanceList), nil
	}
}

func (z zabbixMaintenanceListDo) FirstOrCreate() (*model.ZabbixMaintenanceList, error) {
	if result, err := z.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ZabbixMaintenanceList), nil
	}
}

func (z zabbixMaintenanceListDo) FindByPage(offset int, limit int) (result []*model.ZabbixMaintenanceList, count int64, err error) {
	result, err = z.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = z.Offset(-1).Limit(-1).Count()
	return
}

func (z zabbixMaintenanceListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = z.Count()
	if err != nil {
		return
	}

	err = z.Offset(offset).Limit(limit).Scan(result)
	return
}

func (z zabbixMaintenanceListDo) Scan(result interface{}) (err error) {
	return z.DO.Scan(result)
}

func (z zabbixMaintenanceListDo) Delete(models ...*model.ZabbixMaintenanceList) (result gen.ResultInfo, err error) {
	return z.DO.Delete(models)
}

func (z *zabbixMaintenanceListDo) withDO(do gen.Dao) *zabbixMaintenanceListDo {
	z.DO = *do.(*gen.DO)
	return z
}
