// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"db_lab_library/mysql/model"
)

func newPublisher(db *gorm.DB, opts ...gen.DOOption) publisher {
	_publisher := publisher{}

	_publisher.publisherDo.UseDB(db, opts...)
	_publisher.publisherDo.UseModel(&model.Publisher{})

	tableName := _publisher.publisherDo.TableName()
	_publisher.ALL = field.NewAsterisk(tableName)
	_publisher.ID = field.NewInt32(tableName, "id")
	_publisher.Name = field.NewString(tableName, "name")
	_publisher.Address = field.NewString(tableName, "address")
	_publisher.Email = field.NewString(tableName, "email")
	_publisher.CreatedAt = field.NewTime(tableName, "created_at")
	_publisher.UpdatedAt = field.NewTime(tableName, "updated_at")

	_publisher.fillFieldMap()

	return _publisher
}

type publisher struct {
	publisherDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String
	Address   field.String
	Email     field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p publisher) Table(newTableName string) *publisher {
	p.publisherDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publisher) As(alias string) *publisher {
	p.publisherDo.DO = *(p.publisherDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publisher) updateTableName(table string) *publisher {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Name = field.NewString(table, "name")
	p.Address = field.NewString(table, "address")
	p.Email = field.NewString(table, "email")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *publisher) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publisher) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["name"] = p.Name
	p.fieldMap["address"] = p.Address
	p.fieldMap["email"] = p.Email
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p publisher) clone(db *gorm.DB) publisher {
	p.publisherDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publisher) replaceDB(db *gorm.DB) publisher {
	p.publisherDo.ReplaceDB(db)
	return p
}

type publisherDo struct{ gen.DO }

type IPublisherDo interface {
	gen.SubQuery
	Debug() IPublisherDo
	WithContext(ctx context.Context) IPublisherDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPublisherDo
	WriteDB() IPublisherDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPublisherDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPublisherDo
	Not(conds ...gen.Condition) IPublisherDo
	Or(conds ...gen.Condition) IPublisherDo
	Select(conds ...field.Expr) IPublisherDo
	Where(conds ...gen.Condition) IPublisherDo
	Order(conds ...field.Expr) IPublisherDo
	Distinct(cols ...field.Expr) IPublisherDo
	Omit(cols ...field.Expr) IPublisherDo
	Join(table schema.Tabler, on ...field.Expr) IPublisherDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPublisherDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPublisherDo
	Group(cols ...field.Expr) IPublisherDo
	Having(conds ...gen.Condition) IPublisherDo
	Limit(limit int) IPublisherDo
	Offset(offset int) IPublisherDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPublisherDo
	Unscoped() IPublisherDo
	Create(values ...*model.Publisher) error
	CreateInBatches(values []*model.Publisher, batchSize int) error
	Save(values ...*model.Publisher) error
	First() (*model.Publisher, error)
	Take() (*model.Publisher, error)
	Last() (*model.Publisher, error)
	Find() ([]*model.Publisher, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Publisher, err error)
	FindInBatches(result *[]*model.Publisher, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Publisher) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPublisherDo
	Assign(attrs ...field.AssignExpr) IPublisherDo
	Joins(fields ...field.RelationField) IPublisherDo
	Preload(fields ...field.RelationField) IPublisherDo
	FirstOrInit() (*model.Publisher, error)
	FirstOrCreate() (*model.Publisher, error)
	FindByPage(offset int, limit int) (result []*model.Publisher, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPublisherDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p publisherDo) Debug() IPublisherDo {
	return p.withDO(p.DO.Debug())
}

func (p publisherDo) WithContext(ctx context.Context) IPublisherDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publisherDo) ReadDB() IPublisherDo {
	return p.Clauses(dbresolver.Read)
}

func (p publisherDo) WriteDB() IPublisherDo {
	return p.Clauses(dbresolver.Write)
}

func (p publisherDo) Session(config *gorm.Session) IPublisherDo {
	return p.withDO(p.DO.Session(config))
}

func (p publisherDo) Clauses(conds ...clause.Expression) IPublisherDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publisherDo) Returning(value interface{}, columns ...string) IPublisherDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publisherDo) Not(conds ...gen.Condition) IPublisherDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publisherDo) Or(conds ...gen.Condition) IPublisherDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publisherDo) Select(conds ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publisherDo) Where(conds ...gen.Condition) IPublisherDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publisherDo) Order(conds ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publisherDo) Distinct(cols ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publisherDo) Omit(cols ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publisherDo) Join(table schema.Tabler, on ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publisherDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publisherDo) RightJoin(table schema.Tabler, on ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publisherDo) Group(cols ...field.Expr) IPublisherDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publisherDo) Having(conds ...gen.Condition) IPublisherDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publisherDo) Limit(limit int) IPublisherDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publisherDo) Offset(offset int) IPublisherDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publisherDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPublisherDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publisherDo) Unscoped() IPublisherDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publisherDo) Create(values ...*model.Publisher) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publisherDo) CreateInBatches(values []*model.Publisher, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publisherDo) Save(values ...*model.Publisher) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publisherDo) First() (*model.Publisher, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publisher), nil
	}
}

func (p publisherDo) Take() (*model.Publisher, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publisher), nil
	}
}

func (p publisherDo) Last() (*model.Publisher, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publisher), nil
	}
}

func (p publisherDo) Find() ([]*model.Publisher, error) {
	result, err := p.DO.Find()
	return result.([]*model.Publisher), err
}

func (p publisherDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Publisher, err error) {
	buf := make([]*model.Publisher, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publisherDo) FindInBatches(result *[]*model.Publisher, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publisherDo) Attrs(attrs ...field.AssignExpr) IPublisherDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publisherDo) Assign(attrs ...field.AssignExpr) IPublisherDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publisherDo) Joins(fields ...field.RelationField) IPublisherDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publisherDo) Preload(fields ...field.RelationField) IPublisherDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publisherDo) FirstOrInit() (*model.Publisher, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publisher), nil
	}
}

func (p publisherDo) FirstOrCreate() (*model.Publisher, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Publisher), nil
	}
}

func (p publisherDo) FindByPage(offset int, limit int) (result []*model.Publisher, count int64, err error) {
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

func (p publisherDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publisherDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publisherDo) Delete(models ...*model.Publisher) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publisherDo) withDO(do gen.Dao) *publisherDo {
	p.DO = *do.(*gen.DO)
	return p
}