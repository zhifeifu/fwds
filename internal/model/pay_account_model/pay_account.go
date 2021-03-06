///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package pay_account_model

import (
	"fmt"
	"fwds/internal/model"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *PayAccount {
	return new(PayAccount)
}

func NewQueryBuilder() *payAccountModelQueryBuilder {
	return new(payAccountModelQueryBuilder)
}

func (t *PayAccount) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *PayAccount) Update(db *gorm.DB) error {
	return db.Where("id", t.Id).Updates(t).Error
}

type payAccountModelQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *payAccountModelQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *payAccountModelQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&PayAccount{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *payAccountModelQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&PayAccount{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *payAccountModelQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&PayAccount{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
		res.Error = nil
	}
	return c, res.Error
}

func (qb *payAccountModelQueryBuilder) First(db *gorm.DB) (*PayAccount, error) {
	ret := &PayAccount{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
		res.Error = nil
	}
	return ret, res.Error
}

func (qb *payAccountModelQueryBuilder) QueryOne(db *gorm.DB) (*PayAccount, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *payAccountModelQueryBuilder) QueryAll(db *gorm.DB) ([]*PayAccount, error) {
	var ret []*PayAccount
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *payAccountModelQueryBuilder) Limit(limit int) *payAccountModelQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *payAccountModelQueryBuilder) Offset(offset int) *payAccountModelQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereId(p model.Predicate, value int64) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereIdIn(value []int64) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereIdNotIn(value []int64) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderById(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereName(p model.Predicate, value string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereNameIn(value []string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereNameNotIn(value []string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderByName(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereAppKey(p model.Predicate, value string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_key", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereAppKeyIn(value []string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_key", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereAppKeyNotIn(value []string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_key", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderByAppKey(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "app_key "+order)
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereAppSecret(p model.Predicate, value string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_secret", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereAppSecretIn(value []string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_secret", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereAppSecretNotIn(value []string) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "app_secret", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderByAppSecret(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "app_secret "+order)
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereStatus(p model.Predicate, value int32) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereStatusIn(value []int32) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereStatusNotIn(value []int32) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderByStatus(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereCreateTime(p model.Predicate, value time.Time) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereCreateTimeIn(value []time.Time) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereCreateTimeNotIn(value []time.Time) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderByCreateTime(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereUpdateTime(p model.Predicate, value time.Time) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", p),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereUpdateTimeIn(value []time.Time) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) WhereUpdateTimeNotIn(value []time.Time) *payAccountModelQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *payAccountModelQueryBuilder) OrderByUpdateTime(asc bool) *payAccountModelQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_time "+order)
	return qb
}
