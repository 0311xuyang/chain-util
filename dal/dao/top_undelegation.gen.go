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

func newTopUndelegation(db *gorm.DB, opts ...gen.DOOption) topUndelegation {
	_topUndelegation := topUndelegation{}

	_topUndelegation.topUndelegationDo.UseDB(db, opts...)
	_topUndelegation.topUndelegationDo.UseModel(&model.TopUndelegation{})

	tableName := _topUndelegation.topUndelegationDo.TableName()
	_topUndelegation.ALL = field.NewAsterisk(tableName)
	_topUndelegation.Delegator = field.NewString(tableName, "delegator")
	_topUndelegation.Validator = field.NewString(tableName, "validator")
	_topUndelegation.Amount = field.NewInt64(tableName, "amount")
	_topUndelegation.ReleaseTime = field.NewTime(tableName, "release_time")
	_topUndelegation.ChainName = field.NewString(tableName, "chain_name")

	_topUndelegation.fillFieldMap()

	return _topUndelegation
}

type topUndelegation struct {
	topUndelegationDo

	ALL         field.Asterisk
	Delegator   field.String
	Validator   field.String
	Amount      field.Int64
	ReleaseTime field.Time
	ChainName   field.String

	fieldMap map[string]field.Expr
}

func (t topUndelegation) Table(newTableName string) *topUndelegation {
	t.topUndelegationDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t topUndelegation) As(alias string) *topUndelegation {
	t.topUndelegationDo.DO = *(t.topUndelegationDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *topUndelegation) updateTableName(table string) *topUndelegation {
	t.ALL = field.NewAsterisk(table)
	t.Delegator = field.NewString(table, "delegator")
	t.Validator = field.NewString(table, "validator")
	t.Amount = field.NewInt64(table, "amount")
	t.ReleaseTime = field.NewTime(table, "release_time")
	t.ChainName = field.NewString(table, "chain_name")

	t.fillFieldMap()

	return t
}

func (t *topUndelegation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *topUndelegation) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["delegator"] = t.Delegator
	t.fieldMap["validator"] = t.Validator
	t.fieldMap["amount"] = t.Amount
	t.fieldMap["release_time"] = t.ReleaseTime
	t.fieldMap["chain_name"] = t.ChainName
}

func (t topUndelegation) clone(db *gorm.DB) topUndelegation {
	t.topUndelegationDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t topUndelegation) replaceDB(db *gorm.DB) topUndelegation {
	t.topUndelegationDo.ReplaceDB(db)
	return t
}

type topUndelegationDo struct{ gen.DO }

func (t topUndelegationDo) Debug() *topUndelegationDo {
	return t.withDO(t.DO.Debug())
}

func (t topUndelegationDo) WithContext(ctx context.Context) *topUndelegationDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t topUndelegationDo) ReadDB() *topUndelegationDo {
	return t.Clauses(dbresolver.Read)
}

func (t topUndelegationDo) WriteDB() *topUndelegationDo {
	return t.Clauses(dbresolver.Write)
}

func (t topUndelegationDo) Session(config *gorm.Session) *topUndelegationDo {
	return t.withDO(t.DO.Session(config))
}

func (t topUndelegationDo) Clauses(conds ...clause.Expression) *topUndelegationDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t topUndelegationDo) Returning(value interface{}, columns ...string) *topUndelegationDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t topUndelegationDo) Not(conds ...gen.Condition) *topUndelegationDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t topUndelegationDo) Or(conds ...gen.Condition) *topUndelegationDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t topUndelegationDo) Select(conds ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t topUndelegationDo) Where(conds ...gen.Condition) *topUndelegationDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t topUndelegationDo) Order(conds ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t topUndelegationDo) Distinct(cols ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t topUndelegationDo) Omit(cols ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t topUndelegationDo) Join(table schema.Tabler, on ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t topUndelegationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t topUndelegationDo) RightJoin(table schema.Tabler, on ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t topUndelegationDo) Group(cols ...field.Expr) *topUndelegationDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t topUndelegationDo) Having(conds ...gen.Condition) *topUndelegationDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t topUndelegationDo) Limit(limit int) *topUndelegationDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t topUndelegationDo) Offset(offset int) *topUndelegationDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t topUndelegationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *topUndelegationDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t topUndelegationDo) Unscoped() *topUndelegationDo {
	return t.withDO(t.DO.Unscoped())
}

func (t topUndelegationDo) Create(values ...*model.TopUndelegation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t topUndelegationDo) CreateInBatches(values []*model.TopUndelegation, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t topUndelegationDo) Save(values ...*model.TopUndelegation) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t topUndelegationDo) First() (*model.TopUndelegation, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopUndelegation), nil
	}
}

func (t topUndelegationDo) Take() (*model.TopUndelegation, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopUndelegation), nil
	}
}

func (t topUndelegationDo) Last() (*model.TopUndelegation, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopUndelegation), nil
	}
}

func (t topUndelegationDo) Find() ([]*model.TopUndelegation, error) {
	result, err := t.DO.Find()
	return result.([]*model.TopUndelegation), err
}

func (t topUndelegationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TopUndelegation, err error) {
	buf := make([]*model.TopUndelegation, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t topUndelegationDo) FindInBatches(result *[]*model.TopUndelegation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t topUndelegationDo) Attrs(attrs ...field.AssignExpr) *topUndelegationDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t topUndelegationDo) Assign(attrs ...field.AssignExpr) *topUndelegationDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t topUndelegationDo) Joins(fields ...field.RelationField) *topUndelegationDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t topUndelegationDo) Preload(fields ...field.RelationField) *topUndelegationDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t topUndelegationDo) FirstOrInit() (*model.TopUndelegation, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopUndelegation), nil
	}
}

func (t topUndelegationDo) FirstOrCreate() (*model.TopUndelegation, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TopUndelegation), nil
	}
}

func (t topUndelegationDo) FindByPage(offset int, limit int) (result []*model.TopUndelegation, count int64, err error) {
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

func (t topUndelegationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t topUndelegationDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t topUndelegationDo) Delete(models ...*model.TopUndelegation) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *topUndelegationDo) withDO(do gen.Dao) *topUndelegationDo {
	t.DO = *do.(*gen.DO)
	return t
}
