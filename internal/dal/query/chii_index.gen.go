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

	"github.com/bangumi/server/internal/dal/dao"
)

func newIndex(db *gorm.DB) index {
	_index := index{}

	_index.indexDo.UseDB(db)
	_index.indexDo.UseModel(&dao.Index{})

	tableName := _index.indexDo.TableName()
	_index.ALL = field.NewField(tableName, "*")
	_index.ID = field.NewUint32(tableName, "idx_id")
	_index.Type = field.NewUint8(tableName, "idx_type")
	_index.Title = field.NewString(tableName, "idx_title")
	_index.Desc = field.NewString(tableName, "idx_desc")
	_index.Replies = field.NewUint32(tableName, "idx_replies")
	_index.SubjectTotal = field.NewUint32(tableName, "idx_subject_total")
	_index.Collects = field.NewUint32(tableName, "idx_collects")
	_index.Stats = field.NewString(tableName, "idx_stats")
	_index.Dateline = field.NewInt32(tableName, "idx_dateline")
	_index.Lasttouch = field.NewUint32(tableName, "idx_lasttouch")
	_index.CreatorID = field.NewUint32(tableName, "idx_uid")
	_index.Ban = field.NewBool(tableName, "idx_ban")

	_index.fillFieldMap()

	return _index
}

type index struct {
	indexDo indexDo

	ALL          field.Field
	ID           field.Uint32
	Type         field.Uint8
	Title        field.String
	Desc         field.String
	Replies      field.Uint32
	SubjectTotal field.Uint32
	Collects     field.Uint32
	Stats        field.String
	Dateline     field.Int32
	Lasttouch    field.Uint32
	CreatorID    field.Uint32
	Ban          field.Bool

	fieldMap map[string]field.Expr
}

func (i index) Table(newTableName string) *index {
	i.indexDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i index) As(alias string) *index {
	i.indexDo.DO = *(i.indexDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *index) updateTableName(table string) *index {
	i.ALL = field.NewField(table, "*")
	i.ID = field.NewUint32(table, "idx_id")
	i.Type = field.NewUint8(table, "idx_type")
	i.Title = field.NewString(table, "idx_title")
	i.Desc = field.NewString(table, "idx_desc")
	i.Replies = field.NewUint32(table, "idx_replies")
	i.SubjectTotal = field.NewUint32(table, "idx_subject_total")
	i.Collects = field.NewUint32(table, "idx_collects")
	i.Stats = field.NewString(table, "idx_stats")
	i.Dateline = field.NewInt32(table, "idx_dateline")
	i.Lasttouch = field.NewUint32(table, "idx_lasttouch")
	i.CreatorID = field.NewUint32(table, "idx_uid")
	i.Ban = field.NewBool(table, "idx_ban")

	i.fillFieldMap()

	return i
}

func (i *index) WithContext(ctx context.Context) *indexDo { return i.indexDo.WithContext(ctx) }

func (i index) TableName() string { return i.indexDo.TableName() }

func (i index) Alias() string { return i.indexDo.Alias() }

func (i *index) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *index) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 12)
	i.fieldMap["idx_id"] = i.ID
	i.fieldMap["idx_type"] = i.Type
	i.fieldMap["idx_title"] = i.Title
	i.fieldMap["idx_desc"] = i.Desc
	i.fieldMap["idx_replies"] = i.Replies
	i.fieldMap["idx_subject_total"] = i.SubjectTotal
	i.fieldMap["idx_collects"] = i.Collects
	i.fieldMap["idx_stats"] = i.Stats
	i.fieldMap["idx_dateline"] = i.Dateline
	i.fieldMap["idx_lasttouch"] = i.Lasttouch
	i.fieldMap["idx_uid"] = i.CreatorID
	i.fieldMap["idx_ban"] = i.Ban
}

func (i index) clone(db *gorm.DB) index {
	i.indexDo.ReplaceDB(db)
	return i
}

type indexDo struct{ gen.DO }

func (i indexDo) Debug() *indexDo {
	return i.withDO(i.DO.Debug())
}

func (i indexDo) WithContext(ctx context.Context) *indexDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i indexDo) ReadDB(ctx context.Context) *indexDo {
	return i.WithContext(ctx).Clauses(dbresolver.Read)
}

func (i indexDo) WriteDB(ctx context.Context) *indexDo {
	return i.WithContext(ctx).Clauses(dbresolver.Write)
}

func (i indexDo) Clauses(conds ...clause.Expression) *indexDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i indexDo) Returning(value interface{}, columns ...string) *indexDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i indexDo) Not(conds ...gen.Condition) *indexDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i indexDo) Or(conds ...gen.Condition) *indexDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i indexDo) Select(conds ...field.Expr) *indexDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i indexDo) Where(conds ...gen.Condition) *indexDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i indexDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *indexDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i indexDo) Order(conds ...field.Expr) *indexDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i indexDo) Distinct(cols ...field.Expr) *indexDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i indexDo) Omit(cols ...field.Expr) *indexDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i indexDo) Join(table schema.Tabler, on ...field.Expr) *indexDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i indexDo) LeftJoin(table schema.Tabler, on ...field.Expr) *indexDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i indexDo) RightJoin(table schema.Tabler, on ...field.Expr) *indexDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i indexDo) Group(cols ...field.Expr) *indexDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i indexDo) Having(conds ...gen.Condition) *indexDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i indexDo) Limit(limit int) *indexDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i indexDo) Offset(offset int) *indexDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i indexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *indexDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i indexDo) Unscoped() *indexDo {
	return i.withDO(i.DO.Unscoped())
}

func (i indexDo) Create(values ...*dao.Index) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i indexDo) CreateInBatches(values []*dao.Index, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i indexDo) Save(values ...*dao.Index) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i indexDo) First() (*dao.Index, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Index), nil
	}
}

func (i indexDo) Take() (*dao.Index, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Index), nil
	}
}

func (i indexDo) Last() (*dao.Index, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Index), nil
	}
}

func (i indexDo) Find() ([]*dao.Index, error) {
	result, err := i.DO.Find()
	return result.([]*dao.Index), err
}

func (i indexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Index, err error) {
	buf := make([]*dao.Index, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i indexDo) FindInBatches(result *[]*dao.Index, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i indexDo) Attrs(attrs ...field.AssignExpr) *indexDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i indexDo) Assign(attrs ...field.AssignExpr) *indexDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i indexDo) Joins(fields ...field.RelationField) *indexDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i indexDo) Preload(fields ...field.RelationField) *indexDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i indexDo) FirstOrInit() (*dao.Index, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Index), nil
	}
}

func (i indexDo) FirstOrCreate() (*dao.Index, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Index), nil
	}
}

func (i indexDo) FindByPage(offset int, limit int) (result []*dao.Index, count int64, err error) {
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

func (i indexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i *indexDo) withDO(do gen.Dao) *indexDo {
	i.DO = *do.(*gen.DO)
	return i
}
