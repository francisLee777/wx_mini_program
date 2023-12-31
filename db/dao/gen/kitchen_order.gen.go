// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"wxcloudrun-golang/db/model"
)

func newOrderDBModel(db *gorm.DB, opts ...gen.DOOption) orderDBModel {
	_orderDBModel := orderDBModel{}

	_orderDBModel.orderDBModelDo.UseDB(db, opts...)
	_orderDBModel.orderDBModelDo.UseModel(&model.OrderDBModel{})

	tableName := _orderDBModel.orderDBModelDo.TableName()
	_orderDBModel.ALL = field.NewAsterisk(tableName)
	_orderDBModel.ID = field.NewInt64(tableName, "id")
	_orderDBModel.UniqueCode = field.NewString(tableName, "unique_code")
	_orderDBModel.OpenID = field.NewString(tableName, "openId")
	_orderDBModel.CreateName = field.NewString(tableName, "create_name")
	_orderDBModel.Info = field.NewString(tableName, "info")
	_orderDBModel.Status = field.NewInt32(tableName, "status")
	_orderDBModel.TargetPeriod = field.NewString(tableName, "target_period")
	_orderDBModel.Extra = field.NewString(tableName, "extra")
	_orderDBModel.CreateTime = field.NewTime(tableName, "create_time")
	_orderDBModel.UpdateTime = field.NewTime(tableName, "update_time")

	_orderDBModel.fillFieldMap()

	return _orderDBModel
}

type orderDBModel struct {
	orderDBModelDo

	ALL          field.Asterisk
	ID           field.Int64  // 自增id
	UniqueCode   field.String // 订单号[全局唯一]
	OpenID       field.String // 创建人id
	CreateName   field.String // 创建人name
	Info         field.String // 订单内容，暂时不做索引直接怼进去
	Status       field.Int32  // 订单状态， 1-新建[提单]  2-已支付
	TargetPeriod field.String // 订餐的目标时间，取值只有9个,已校验.明天中午、后天早上、今天晚上等
	Extra        field.String // 扩展字段
	CreateTime   field.Time   // 创建时间
	UpdateTime   field.Time   // 最后更新时间

	fieldMap map[string]field.Expr
}

func (o orderDBModel) Table(newTableName string) *orderDBModel {
	o.orderDBModelDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o orderDBModel) As(alias string) *orderDBModel {
	o.orderDBModelDo.DO = *(o.orderDBModelDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *orderDBModel) updateTableName(table string) *orderDBModel {
	o.ALL = field.NewAsterisk(table)
	o.ID = field.NewInt64(table, "id")
	o.UniqueCode = field.NewString(table, "unique_code")
	o.OpenID = field.NewString(table, "openId")
	o.CreateName = field.NewString(table, "create_name")
	o.Info = field.NewString(table, "info")
	o.Status = field.NewInt32(table, "status")
	o.TargetPeriod = field.NewString(table, "target_period")
	o.Extra = field.NewString(table, "extra")
	o.CreateTime = field.NewTime(table, "create_time")
	o.UpdateTime = field.NewTime(table, "update_time")

	o.fillFieldMap()

	return o
}

func (o *orderDBModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *orderDBModel) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 10)
	o.fieldMap["id"] = o.ID
	o.fieldMap["unique_code"] = o.UniqueCode
	o.fieldMap["openId"] = o.OpenID
	o.fieldMap["create_name"] = o.CreateName
	o.fieldMap["info"] = o.Info
	o.fieldMap["status"] = o.Status
	o.fieldMap["target_period"] = o.TargetPeriod
	o.fieldMap["extra"] = o.Extra
	o.fieldMap["create_time"] = o.CreateTime
	o.fieldMap["update_time"] = o.UpdateTime
}

func (o orderDBModel) clone(db *gorm.DB) orderDBModel {
	o.orderDBModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o orderDBModel) replaceDB(db *gorm.DB) orderDBModel {
	o.orderDBModelDo.ReplaceDB(db)
	return o
}

type orderDBModelDo struct{ gen.DO }

type IOrderDBModelDo interface {
	gen.SubQuery
	Debug() IOrderDBModelDo
	WithContext(ctx context.Context) IOrderDBModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrderDBModelDo
	WriteDB() IOrderDBModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrderDBModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrderDBModelDo
	Not(conds ...gen.Condition) IOrderDBModelDo
	Or(conds ...gen.Condition) IOrderDBModelDo
	Select(conds ...field.Expr) IOrderDBModelDo
	Where(conds ...gen.Condition) IOrderDBModelDo
	Order(conds ...field.Expr) IOrderDBModelDo
	Distinct(cols ...field.Expr) IOrderDBModelDo
	Omit(cols ...field.Expr) IOrderDBModelDo
	Join(table schema.Tabler, on ...field.Expr) IOrderDBModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrderDBModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrderDBModelDo
	Group(cols ...field.Expr) IOrderDBModelDo
	Having(conds ...gen.Condition) IOrderDBModelDo
	Limit(limit int) IOrderDBModelDo
	Offset(offset int) IOrderDBModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderDBModelDo
	Unscoped() IOrderDBModelDo
	Create(values ...*model.OrderDBModel) error
	CreateInBatches(values []*model.OrderDBModel, batchSize int) error
	Save(values ...*model.OrderDBModel) error
	First() (*model.OrderDBModel, error)
	Take() (*model.OrderDBModel, error)
	Last() (*model.OrderDBModel, error)
	Find() ([]*model.OrderDBModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderDBModel, err error)
	FindInBatches(result *[]*model.OrderDBModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.OrderDBModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrderDBModelDo
	Assign(attrs ...field.AssignExpr) IOrderDBModelDo
	Joins(fields ...field.RelationField) IOrderDBModelDo
	Preload(fields ...field.RelationField) IOrderDBModelDo
	FirstOrInit() (*model.OrderDBModel, error)
	FirstOrCreate() (*model.OrderDBModel, error)
	FindByPage(offset int, limit int) (result []*model.OrderDBModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrderDBModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o orderDBModelDo) Debug() IOrderDBModelDo {
	return o.withDO(o.DO.Debug())
}

func (o orderDBModelDo) WithContext(ctx context.Context) IOrderDBModelDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o orderDBModelDo) ReadDB() IOrderDBModelDo {
	return o.Clauses(dbresolver.Read)
}

func (o orderDBModelDo) WriteDB() IOrderDBModelDo {
	return o.Clauses(dbresolver.Write)
}

func (o orderDBModelDo) Session(config *gorm.Session) IOrderDBModelDo {
	return o.withDO(o.DO.Session(config))
}

func (o orderDBModelDo) Clauses(conds ...clause.Expression) IOrderDBModelDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o orderDBModelDo) Returning(value interface{}, columns ...string) IOrderDBModelDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o orderDBModelDo) Not(conds ...gen.Condition) IOrderDBModelDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o orderDBModelDo) Or(conds ...gen.Condition) IOrderDBModelDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o orderDBModelDo) Select(conds ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o orderDBModelDo) Where(conds ...gen.Condition) IOrderDBModelDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o orderDBModelDo) Order(conds ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o orderDBModelDo) Distinct(cols ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o orderDBModelDo) Omit(cols ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o orderDBModelDo) Join(table schema.Tabler, on ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o orderDBModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o orderDBModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o orderDBModelDo) Group(cols ...field.Expr) IOrderDBModelDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o orderDBModelDo) Having(conds ...gen.Condition) IOrderDBModelDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o orderDBModelDo) Limit(limit int) IOrderDBModelDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o orderDBModelDo) Offset(offset int) IOrderDBModelDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o orderDBModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrderDBModelDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o orderDBModelDo) Unscoped() IOrderDBModelDo {
	return o.withDO(o.DO.Unscoped())
}

func (o orderDBModelDo) Create(values ...*model.OrderDBModel) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o orderDBModelDo) CreateInBatches(values []*model.OrderDBModel, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o orderDBModelDo) Save(values ...*model.OrderDBModel) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o orderDBModelDo) First() (*model.OrderDBModel, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderDBModel), nil
	}
}

func (o orderDBModelDo) Take() (*model.OrderDBModel, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderDBModel), nil
	}
}

func (o orderDBModelDo) Last() (*model.OrderDBModel, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderDBModel), nil
	}
}

func (o orderDBModelDo) Find() ([]*model.OrderDBModel, error) {
	result, err := o.DO.Find()
	return result.([]*model.OrderDBModel), err
}

func (o orderDBModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.OrderDBModel, err error) {
	buf := make([]*model.OrderDBModel, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o orderDBModelDo) FindInBatches(result *[]*model.OrderDBModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o orderDBModelDo) Attrs(attrs ...field.AssignExpr) IOrderDBModelDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o orderDBModelDo) Assign(attrs ...field.AssignExpr) IOrderDBModelDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o orderDBModelDo) Joins(fields ...field.RelationField) IOrderDBModelDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o orderDBModelDo) Preload(fields ...field.RelationField) IOrderDBModelDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o orderDBModelDo) FirstOrInit() (*model.OrderDBModel, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderDBModel), nil
	}
}

func (o orderDBModelDo) FirstOrCreate() (*model.OrderDBModel, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.OrderDBModel), nil
	}
}

func (o orderDBModelDo) FindByPage(offset int, limit int) (result []*model.OrderDBModel, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o orderDBModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o orderDBModelDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o orderDBModelDo) Delete(models ...*model.OrderDBModel) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *orderDBModelDo) withDO(do gen.Dao) *orderDBModelDo {
	o.DO = *do.(*gen.DO)
	return o
}
