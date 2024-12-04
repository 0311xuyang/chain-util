// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/0311xuyang/chain-util/dal/model"
)

func newTopDelegation(db *gorm.DB, opts ...gen.DOOption) topDelegation {
	_topDelegation := topDelegation{}

	_topDelegation.topDelegationDo.UseDB(db, opts...)
	_topDelegation.topDelegationDo.UseModel(&model.TopDelegation{})

	tableName := _topDelegation.topDelegationDo.TableName()
	_topDelegation.ALL = field.NewAsterisk(tableName)
	_topDelegation.Delegator = field.NewString(tableName, "delegator")
	_topDelegation.Amount = field.NewInt64(tableName, "amount")
	_topDelegation.ChainName = field.NewString(tableName, "chain_name")
	_topDelegation.Validator = field.NewString(tableName, "validator")

	_topDelegation.fillFieldMap()

	return _topDelegation
}

type topDelegation struct {
	topDelegationDo

	ALL       field.Asterisk
	Delegator field.String
	Amount    field.Int64
	ChainName field.String
	Validator field.String

	fieldMap map[string]field.Expr
}

func (t topDelegation) Table(newTableName string) *topDelegation {
	t.topDelegationDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t topDelegation) As(alias string) *topDelegation {
	t.topDelegationDo.DO = *(t.topDelegationDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *topDelegation) updateTableName(table string) *topDelegation {
	t.ALL = field.NewAsterisk(table)
	t.Delegator = field.NewString(table, "delegator")
	t.Amount = field.NewInt64(table, "amount")
	t.ChainName = field.NewString(table, "chain_name")
	t.Validator = field.NewString(table, "validator")

	t.fillFieldMap()

	return t
}

func (t *topDelegation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *topDelegation) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 4)
	t.fieldMap["delegator"] = t.Delegator
	t.fieldMap["amount"] = t.Amount
	t.fieldMap["chain_name"] = t.ChainName
	t.fieldMap["validator"] = t.Validator
}

func (t topDelegation) clone(db *gorm.DB) topDelegation {
	t.topDelegationDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t topDelegation) replaceDB(db *gorm.DB) topDelegation {
	t.topDelegationDo.ReplaceDB(db)
	return t
}

type topDelegationDo struct{ gen.DO }

func (t topDelegationDo) Debug() *topDelegationDo {
	return t.withDO(t.DO.Debug())
}

func (t topDelegationDo) WithContext(ctx context.Context) *topDelegationDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t topDelegationDo) ReadDB() *topDelegationDo {
	return t.Clauses(dbresolver.Read)
}

func (t topDelegationDo) WriteDB() *topDelegationDo {
	return t.Clauses(dbresolver.Write)
}

func (t topDelegationDo) Session(config *gorm.Session) *topDelegationDo {
	return t.withDO(t.DO.Session(config))
}

func (t topDelegationDo) Clauses(conds ...clause.Expression) *topDelegationDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t topDelegationDo) Returning(value interface{}, columns ...string) *topDelegationDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t topDelegationDo) Not(conds ...gen.Condition) *topDelegationDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t topDelegationDo) Or(conds ...gen.Condition) *topDelegationDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t topDelegationDo) Select(conds ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t topDelegationDo) Where(conds ...gen.Condition) *topDelegationDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t topDelegationDo) Order(conds ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t topDelegationDo) Distinct(cols ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t topDelegationDo) Omit(cols ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t topDelegationDo) Join(table schema.Tabler, on ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t topDelegationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t topDelegationDo) RightJoin(table schema.Tabler, on ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t topDelegationDo) Group(cols ...field.Expr) *topDelegationDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t topDelegationDo) Having(conds ...gen.Condition) *topDelegationDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t topDelegationDo) Limit(limit int) *topDelegationDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t topDelegationDo) Offset(offset int) *topDelegationDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t topDelegationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *topDelegationDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t topDelegationDo) Unscoped() *topDelegationDo {
	return t.withDO(t.DO.Unscoped())
}

func (t topDelegationDo) Create(values ...*model.TopDelegation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t topDelegationDo) CreateInBatches(values []*model.TopDelegation, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t topDelegationDo) Save(values ...*model.TopDelegation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t topDelegationDo) First() (*model.TopDelegation, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopDelegation), nil
	}
}

func (t topDelegationDo) Take() (*model.TopDelegation, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopDelegation), nil
	}
}

func (t topDelegationDo) Last() (*model.TopDelegation, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopDelegation), nil
	}
}

func (t topDelegationDo) Find() ([]*model.TopDelegation, error) {
	result, err := t.DO.Find()
	return result.([]*model.TopDelegation), err
}

func (t topDelegationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TopDelegation, err error) {
	buf := make([]*model.TopDelegation, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t topDelegationDo) FindInBatches(result *[]*model.TopDelegation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t topDelegationDo) Attrs(attrs ...field.AssignExpr) *topDelegationDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t topDelegationDo) Assign(attrs ...field.AssignExpr) *topDelegationDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t topDelegationDo) Joins(fields ...field.RelationField) *topDelegationDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t topDelegationDo) Preload(fields ...field.RelationField) *topDelegationDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t topDelegationDo) FirstOrInit() (*model.TopDelegation, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopDelegation), nil
	}
}

func (t topDelegationDo) FirstOrCreate() (*model.TopDelegation, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopDelegation), nil
	}
}

func (t topDelegationDo) FindByPage(offset int, limit int) (result []*model.TopDelegation, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t topDelegationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t topDelegationDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t topDelegationDo) Delete(models ...*model.TopDelegation) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *topDelegationDo) withDO(do gen.Dao) *topDelegationDo {
	t.DO = *do.(*gen.DO)
	return t
}
