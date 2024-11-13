// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/CrazyGang97/go-clean-arch/internal/infra/db/model"
)

func newBookPO(db *gorm.DB, opts ...gen.DOOption) bookPO {
	_bookPO := bookPO{}

	_bookPO.bookPODo.UseDB(db, opts...)
	_bookPO.bookPODo.UseModel(&model.BookPO{})

	tableName := _bookPO.bookPODo.TableName()
	_bookPO.ALL = field.NewAsterisk(tableName)
	_bookPO.ID = field.NewInt64(tableName, "id")
	_bookPO.Name = field.NewString(tableName, "name")
	_bookPO.AuthorID = field.NewInt64(tableName, "author_id")

	_bookPO.fillFieldMap()

	return _bookPO
}

type bookPO struct {
	bookPODo bookPODo

	ALL      field.Asterisk
	ID       field.Int64
	Name     field.String
	AuthorID field.Int64

	fieldMap map[string]field.Expr
}

func (b bookPO) Table(newTableName string) *bookPO {
	b.bookPODo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bookPO) As(alias string) *bookPO {
	b.bookPODo.DO = *(b.bookPODo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bookPO) updateTableName(table string) *bookPO {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.Name = field.NewString(table, "name")
	b.AuthorID = field.NewInt64(table, "author_id")

	b.fillFieldMap()

	return b
}

func (b *bookPO) WithContext(ctx context.Context) *bookPODo { return b.bookPODo.WithContext(ctx) }

func (b bookPO) TableName() string { return b.bookPODo.TableName() }

func (b bookPO) Alias() string { return b.bookPODo.Alias() }

func (b bookPO) Columns(cols ...field.Expr) gen.Columns { return b.bookPODo.Columns(cols...) }

func (b *bookPO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bookPO) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 3)
	b.fieldMap["id"] = b.ID
	b.fieldMap["name"] = b.Name
	b.fieldMap["author_id"] = b.AuthorID
}

func (b bookPO) clone(db *gorm.DB) bookPO {
	b.bookPODo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bookPO) replaceDB(db *gorm.DB) bookPO {
	b.bookPODo.ReplaceDB(db)
	return b
}

type bookPODo struct{ gen.DO }

func (b bookPODo) Debug() *bookPODo {
	return b.withDO(b.DO.Debug())
}

func (b bookPODo) WithContext(ctx context.Context) *bookPODo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookPODo) ReadDB() *bookPODo {
	return b.Clauses(dbresolver.Read)
}

func (b bookPODo) WriteDB() *bookPODo {
	return b.Clauses(dbresolver.Write)
}

func (b bookPODo) Session(config *gorm.Session) *bookPODo {
	return b.withDO(b.DO.Session(config))
}

func (b bookPODo) Clauses(conds ...clause.Expression) *bookPODo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookPODo) Returning(value interface{}, columns ...string) *bookPODo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookPODo) Not(conds ...gen.Condition) *bookPODo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookPODo) Or(conds ...gen.Condition) *bookPODo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookPODo) Select(conds ...field.Expr) *bookPODo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookPODo) Where(conds ...gen.Condition) *bookPODo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookPODo) Order(conds ...field.Expr) *bookPODo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookPODo) Distinct(cols ...field.Expr) *bookPODo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookPODo) Omit(cols ...field.Expr) *bookPODo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookPODo) Join(table schema.Tabler, on ...field.Expr) *bookPODo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookPODo) LeftJoin(table schema.Tabler, on ...field.Expr) *bookPODo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookPODo) RightJoin(table schema.Tabler, on ...field.Expr) *bookPODo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookPODo) Group(cols ...field.Expr) *bookPODo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookPODo) Having(conds ...gen.Condition) *bookPODo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookPODo) Limit(limit int) *bookPODo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookPODo) Offset(offset int) *bookPODo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookPODo) Scopes(funcs ...func(gen.Dao) gen.Dao) *bookPODo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookPODo) Unscoped() *bookPODo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookPODo) Create(values ...*model.BookPO) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookPODo) CreateInBatches(values []*model.BookPO, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookPODo) Save(values ...*model.BookPO) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookPODo) First() (*model.BookPO, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookPO), nil
	}
}

func (b bookPODo) Take() (*model.BookPO, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookPO), nil
	}
}

func (b bookPODo) Last() (*model.BookPO, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookPO), nil
	}
}

func (b bookPODo) Find() ([]*model.BookPO, error) {
	result, err := b.DO.Find()
	return result.([]*model.BookPO), err
}

func (b bookPODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookPO, err error) {
	buf := make([]*model.BookPO, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookPODo) FindInBatches(result *[]*model.BookPO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookPODo) Attrs(attrs ...field.AssignExpr) *bookPODo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookPODo) Assign(attrs ...field.AssignExpr) *bookPODo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookPODo) Joins(fields ...field.RelationField) *bookPODo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookPODo) Preload(fields ...field.RelationField) *bookPODo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookPODo) FirstOrInit() (*model.BookPO, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookPO), nil
	}
}

func (b bookPODo) FirstOrCreate() (*model.BookPO, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookPO), nil
	}
}

func (b bookPODo) FindByPage(offset int, limit int) (result []*model.BookPO, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookPODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookPODo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookPODo) Delete(models ...*model.BookPO) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookPODo) withDO(do gen.Dao) *bookPODo {
	b.DO = *do.(*gen.DO)
	return b
}
