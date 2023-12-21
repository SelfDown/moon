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

func newPubfileAggIssueList(db *gorm.DB, opts ...gen.DOOption) pubfileAggIssueList {
	_pubfileAggIssueList := pubfileAggIssueList{}

	_pubfileAggIssueList.pubfileAggIssueListDo.UseDB(db, opts...)
	_pubfileAggIssueList.pubfileAggIssueListDo.UseModel(&model.PubfileAggIssueList{})

	tableName := _pubfileAggIssueList.pubfileAggIssueListDo.TableName()
	_pubfileAggIssueList.ALL = field.NewAsterisk(tableName)
	_pubfileAggIssueList.ID = field.NewInt32(tableName, "id")
	_pubfileAggIssueList.AggID = field.NewString(tableName, "agg_id")
	_pubfileAggIssueList.ReqIssueID = field.NewString(tableName, "req_issue_id")
	_pubfileAggIssueList.IssueKey = field.NewString(tableName, "issue_key")
	_pubfileAggIssueList.IssueTitle = field.NewString(tableName, "issue_title")
	_pubfileAggIssueList.IssueType = field.NewInt32(tableName, "issue_type")
	_pubfileAggIssueList.IssueStatu = field.NewString(tableName, "issue_statu")
	_pubfileAggIssueList.Issuetype = field.NewString(tableName, "issuetype")
	_pubfileAggIssueList.Issupply = field.NewString(tableName, "issupply")
	_pubfileAggIssueList.IssueAssignee = field.NewString(tableName, "issue_assignee")
	_pubfileAggIssueList.CreateTime = field.NewTime(tableName, "create_time")
	_pubfileAggIssueList.ModifyTime = field.NewTime(tableName, "modify_time")
	_pubfileAggIssueList.Comments = field.NewString(tableName, "comments")

	_pubfileAggIssueList.fillFieldMap()

	return _pubfileAggIssueList
}

type pubfileAggIssueList struct {
	pubfileAggIssueListDo

	ALL           field.Asterisk
	ID            field.Int32
	AggID         field.String // ID
	ReqIssueID    field.String
	IssueKey      field.String
	IssueTitle    field.String
	IssueType     field.Int32 // 0 开发类，1，项目类，2-环境类
	IssueStatu    field.String
	Issuetype     field.String // Jira的ISSUE问题类型
	Issupply      field.String // 是否为开始集成测试后补录的ISSUE 0 或空 表示 非补录 1 表示补录
	IssueAssignee field.String // ISSUE经办人
	CreateTime    field.Time   // 记录创建时间（数据库自动写入）
	ModifyTime    field.Time   // 记录修改时间（数据库自动写入）
	Comments      field.String // 备注说明

	fieldMap map[string]field.Expr
}

func (p pubfileAggIssueList) Table(newTableName string) *pubfileAggIssueList {
	p.pubfileAggIssueListDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pubfileAggIssueList) As(alias string) *pubfileAggIssueList {
	p.pubfileAggIssueListDo.DO = *(p.pubfileAggIssueListDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pubfileAggIssueList) updateTableName(table string) *pubfileAggIssueList {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.AggID = field.NewString(table, "agg_id")
	p.ReqIssueID = field.NewString(table, "req_issue_id")
	p.IssueKey = field.NewString(table, "issue_key")
	p.IssueTitle = field.NewString(table, "issue_title")
	p.IssueType = field.NewInt32(table, "issue_type")
	p.IssueStatu = field.NewString(table, "issue_statu")
	p.Issuetype = field.NewString(table, "issuetype")
	p.Issupply = field.NewString(table, "issupply")
	p.IssueAssignee = field.NewString(table, "issue_assignee")
	p.CreateTime = field.NewTime(table, "create_time")
	p.ModifyTime = field.NewTime(table, "modify_time")
	p.Comments = field.NewString(table, "comments")

	p.fillFieldMap()

	return p
}

func (p *pubfileAggIssueList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pubfileAggIssueList) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 13)
	p.fieldMap["id"] = p.ID
	p.fieldMap["agg_id"] = p.AggID
	p.fieldMap["req_issue_id"] = p.ReqIssueID
	p.fieldMap["issue_key"] = p.IssueKey
	p.fieldMap["issue_title"] = p.IssueTitle
	p.fieldMap["issue_type"] = p.IssueType
	p.fieldMap["issue_statu"] = p.IssueStatu
	p.fieldMap["issuetype"] = p.Issuetype
	p.fieldMap["issupply"] = p.Issupply
	p.fieldMap["issue_assignee"] = p.IssueAssignee
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["modify_time"] = p.ModifyTime
	p.fieldMap["comments"] = p.Comments
}

func (p pubfileAggIssueList) clone(db *gorm.DB) pubfileAggIssueList {
	p.pubfileAggIssueListDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p pubfileAggIssueList) replaceDB(db *gorm.DB) pubfileAggIssueList {
	p.pubfileAggIssueListDo.ReplaceDB(db)
	return p
}

type pubfileAggIssueListDo struct{ gen.DO }

type IPubfileAggIssueListDo interface {
	gen.SubQuery
	Debug() IPubfileAggIssueListDo
	WithContext(ctx context.Context) IPubfileAggIssueListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPubfileAggIssueListDo
	WriteDB() IPubfileAggIssueListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPubfileAggIssueListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPubfileAggIssueListDo
	Not(conds ...gen.Condition) IPubfileAggIssueListDo
	Or(conds ...gen.Condition) IPubfileAggIssueListDo
	Select(conds ...field.Expr) IPubfileAggIssueListDo
	Where(conds ...gen.Condition) IPubfileAggIssueListDo
	Order(conds ...field.Expr) IPubfileAggIssueListDo
	Distinct(cols ...field.Expr) IPubfileAggIssueListDo
	Omit(cols ...field.Expr) IPubfileAggIssueListDo
	Join(table schema.Tabler, on ...field.Expr) IPubfileAggIssueListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPubfileAggIssueListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPubfileAggIssueListDo
	Group(cols ...field.Expr) IPubfileAggIssueListDo
	Having(conds ...gen.Condition) IPubfileAggIssueListDo
	Limit(limit int) IPubfileAggIssueListDo
	Offset(offset int) IPubfileAggIssueListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPubfileAggIssueListDo
	Unscoped() IPubfileAggIssueListDo
	Create(values ...*model.PubfileAggIssueList) error
	CreateInBatches(values []*model.PubfileAggIssueList, batchSize int) error
	Save(values ...*model.PubfileAggIssueList) error
	First() (*model.PubfileAggIssueList, error)
	Take() (*model.PubfileAggIssueList, error)
	Last() (*model.PubfileAggIssueList, error)
	Find() ([]*model.PubfileAggIssueList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PubfileAggIssueList, err error)
	FindInBatches(result *[]*model.PubfileAggIssueList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PubfileAggIssueList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPubfileAggIssueListDo
	Assign(attrs ...field.AssignExpr) IPubfileAggIssueListDo
	Joins(fields ...field.RelationField) IPubfileAggIssueListDo
	Preload(fields ...field.RelationField) IPubfileAggIssueListDo
	FirstOrInit() (*model.PubfileAggIssueList, error)
	FirstOrCreate() (*model.PubfileAggIssueList, error)
	FindByPage(offset int, limit int) (result []*model.PubfileAggIssueList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPubfileAggIssueListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pubfileAggIssueListDo) Debug() IPubfileAggIssueListDo {
	return p.withDO(p.DO.Debug())
}

func (p pubfileAggIssueListDo) WithContext(ctx context.Context) IPubfileAggIssueListDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pubfileAggIssueListDo) ReadDB() IPubfileAggIssueListDo {
	return p.Clauses(dbresolver.Read)
}

func (p pubfileAggIssueListDo) WriteDB() IPubfileAggIssueListDo {
	return p.Clauses(dbresolver.Write)
}

func (p pubfileAggIssueListDo) Session(config *gorm.Session) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Session(config))
}

func (p pubfileAggIssueListDo) Clauses(conds ...clause.Expression) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pubfileAggIssueListDo) Returning(value interface{}, columns ...string) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pubfileAggIssueListDo) Not(conds ...gen.Condition) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pubfileAggIssueListDo) Or(conds ...gen.Condition) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pubfileAggIssueListDo) Select(conds ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pubfileAggIssueListDo) Where(conds ...gen.Condition) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pubfileAggIssueListDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPubfileAggIssueListDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pubfileAggIssueListDo) Order(conds ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pubfileAggIssueListDo) Distinct(cols ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pubfileAggIssueListDo) Omit(cols ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pubfileAggIssueListDo) Join(table schema.Tabler, on ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pubfileAggIssueListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pubfileAggIssueListDo) RightJoin(table schema.Tabler, on ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pubfileAggIssueListDo) Group(cols ...field.Expr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pubfileAggIssueListDo) Having(conds ...gen.Condition) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pubfileAggIssueListDo) Limit(limit int) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pubfileAggIssueListDo) Offset(offset int) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pubfileAggIssueListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pubfileAggIssueListDo) Unscoped() IPubfileAggIssueListDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pubfileAggIssueListDo) Create(values ...*model.PubfileAggIssueList) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pubfileAggIssueListDo) CreateInBatches(values []*model.PubfileAggIssueList, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pubfileAggIssueListDo) Save(values ...*model.PubfileAggIssueList) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pubfileAggIssueListDo) First() (*model.PubfileAggIssueList, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggIssueList), nil
	}
}

func (p pubfileAggIssueListDo) Take() (*model.PubfileAggIssueList, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggIssueList), nil
	}
}

func (p pubfileAggIssueListDo) Last() (*model.PubfileAggIssueList, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggIssueList), nil
	}
}

func (p pubfileAggIssueListDo) Find() ([]*model.PubfileAggIssueList, error) {
	result, err := p.DO.Find()
	return result.([]*model.PubfileAggIssueList), err
}

func (p pubfileAggIssueListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PubfileAggIssueList, err error) {
	buf := make([]*model.PubfileAggIssueList, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pubfileAggIssueListDo) FindInBatches(result *[]*model.PubfileAggIssueList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pubfileAggIssueListDo) Attrs(attrs ...field.AssignExpr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pubfileAggIssueListDo) Assign(attrs ...field.AssignExpr) IPubfileAggIssueListDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pubfileAggIssueListDo) Joins(fields ...field.RelationField) IPubfileAggIssueListDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pubfileAggIssueListDo) Preload(fields ...field.RelationField) IPubfileAggIssueListDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pubfileAggIssueListDo) FirstOrInit() (*model.PubfileAggIssueList, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggIssueList), nil
	}
}

func (p pubfileAggIssueListDo) FirstOrCreate() (*model.PubfileAggIssueList, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PubfileAggIssueList), nil
	}
}

func (p pubfileAggIssueListDo) FindByPage(offset int, limit int) (result []*model.PubfileAggIssueList, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pubfileAggIssueListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pubfileAggIssueListDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pubfileAggIssueListDo) Delete(models ...*model.PubfileAggIssueList) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pubfileAggIssueListDo) withDO(do gen.Dao) *pubfileAggIssueListDo {
	p.DO = *do.(*gen.DO)
	return p
}