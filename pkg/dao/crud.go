package dao

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("not found")
	ErrDupKey   = errors.New("dup key")
)

func Insert(o interface{}) error {
	err := Db.Create(o).Error
	if err != nil {
		nerr, ok := err.(*mysql.MySQLError)
		if ok {
			if nerr.Number == 1062 {
				return ErrDupKey
			}
		}
		return err
	}
	return nil
}

func TxInsert(tx *gorm.DB, o interface{}) error {
	err := tx.Create(o).Error
	if err != nil {
		nerr, ok := err.(*mysql.MySQLError)
		if ok {
			if nerr.Number == 1062 {
				return ErrDupKey
			}
		}
		return err
	}
	return nil
}

func Get(id int, o interface{}) error {
	return GetByQuery(o, "id = ?", id)
}

func GetByQuery(o interface{}, query interface{}, args ...interface{}) error {
	err := Db.Where(query, args...).First(o).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

// update updated_at field
func Update(id int, o interface{}, m map[string]interface{}) error {
	if m == nil {
		return Db.Where("id = ?", id).Updates(o).Error
	} else {
		return Db.Model(o).Where("id = ?", id).Updates(m).Error
	}
}

func Remove(id int, o interface{}) error {
	return RemoveByCond(o, "id = ?", id)
}

func RemoveByCond(o interface{}, query interface{}, args ...interface{}) error {
	return Db.Where(query, args...).Delete(o).Error
}

func List(o interface{}, query interface{}, args ...interface{}) error {
	if query == nil {
		return Db.Find(o).Order("id desc").Error
	}
	return Db.Where(query, args...).Find(o).Order("id desc").Error
}

func ListInById(o interface{}, query interface{}, ids []int) error {
	if query == nil {
		return Db.Find(o).Error
	}
	return Db.Where(query, ids).Find(o).Error
}

// sql template related functions
func RawGet(tpl string, q map[string]interface{}, o interface{}) error {
	tx, err := SqlTemplateClient(tpl, q)
	if err != nil {
		return err
	}
	err = tx.First(o).Error
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func RawFind(tpl string, q map[string]interface{}, o interface{}) error {
	tx, err := SqlTemplateClient(tpl, q)
	if err != nil {
		return err
	}
	return tx.Find(o).Error
}

func RawExec(q map[string]interface{}, tpls ...string) error {
	if len(tpls) == 1 {
		sql, err := htmlTemplate.Execute(tpls[0], &q)
		if err != nil {
			return err
		}
		return Db.Exec(sql).Error
	}
	// 手动开启事务
	rawTx := Db.Exec("START TRANSACTION")
	for _, tpl := range tpls {
		sql, err := htmlTemplate.Execute(tpl, &q)
		if err != nil {
			return err
		}
		rawTx.Exec(sql)
	}
	return rawTx.Exec("commit").Error
}

func SqlTemplateClient(tpl string, q map[string]interface{}) (*gorm.DB, error) {
	sql, err := htmlTemplate.Execute(tpl, &q)
	if err != nil {
		return nil, err
	}
	return Db.Raw(sql), nil
}
