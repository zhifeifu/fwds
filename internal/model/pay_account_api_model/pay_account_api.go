///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pay_account_api_model

import (
	"fmt"
	"fwds/internal/model"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PayAccountApi {
	return new(PayAccountApi)
}

func NewQueryBuilder() *payAccountApiModelQueryBuilder {
	return new(payAccountApiModelQueryBuilder)
}

func (t *PayAccountApi) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *PayAccountApi) Update(db *gorm.DB) error {
	return db.Where("id", t.Id).Updates(t).Error
}

type payAccountApiModelQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *payAccountApiModelQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *payAccountApiModelQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PayAccountApi{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *payAccountApiModelQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PayAccountApi{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *payAccountApiModelQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PayAccountApi{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
		res.Error = nil
	}
	return c, res.Error
}

func (qb *payAccountApiModelQueryBuilder) First(db *gorm.DB) (*PayAccountApi, error) {
	ret := &PayAccountApi{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
		res.Error = nil
	}
	return ret, res.Error
}

func (qb *payAccountApiModelQueryBuilder) QueryOne(db *gorm.DB) (*PayAccountApi, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *payAccountApiModelQueryBuilder) QueryAll(db *gorm.DB) ([]*PayAccountApi, error) {
	var ret []*PayAccountApi
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *payAccountApiModelQueryBuilder) Limit(limit int) *payAccountApiModelQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *payAccountApiModelQueryBuilder) Offset(offset int) *payAccountApiModelQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereId(p model.Predicate, value int64) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereIdIn(value []int64) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereIdNotIn(value []int64) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderById(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WherePayAccountId(p model.Predicate, value int64) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pay_account_id", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WherePayAccountIdIn(value []int64) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pay_account_id", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WherePayAccountIdNotIn(value []int64) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "pay_account_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderByPayAccountId(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "pay_account_id "+order)
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereMethod(p model.Predicate, value string) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereMethodIn(value []string) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereMethodNotIn(value []string) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderByMethod(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "method "+order)
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereApi(p model.Predicate, value string) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereApiIn(value []string) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereApiNotIn(value []string) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderByApi(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "api "+order)
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereStatus(p model.Predicate, value int32) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereStatusIn(value []int32) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereStatusNotIn(value []int32) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderByStatus(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereCreateTime(p model.Predicate, value time.Time) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereCreateTimeIn(value []time.Time) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderByCreateTime(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereUpdateTime(p model.Predicate, value time.Time) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", p),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereUpdateTimeIn(value []time.Time) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) WhereUpdateTimeNotIn(value []time.Time) *payAccountApiModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountApiModelQueryBuilder) OrderByUpdateTime(asc bool) *payAccountApiModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_time "+order)
	return qb
}
