package dao_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "gitlab.shlab.tech/xurui/pdf-reader-backend/pkg/dao"
)

type TestRecord struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Age       int    `json:"age,omitempty"`
	Salary    int    `json:"salary,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty" gorm:"autoUpdateTime:milli"`
}

func TestCRUD(t *testing.T) {
	err := Db.AutoMigrate(new(TestRecord))
	assert.NoError(t, err)
	defer func() {
		Db.Migrator().DropTable(new(TestRecord))
	}()

	t1 := TestRecord{Name: "tom", Age: 20}
	t2 := TestRecord{Name: "hanks", Age: 30}
	t3 := TestRecord{Name: "jim", Age: 40}
	t4 := TestRecord{Name: "green", Age: 15}
	t5 := TestRecord{Name: "green2", Age: 15}

	t.Run("create", func(t *testing.T) {
		err := Insert(&t1)
		assert.NoError(t, err)
		assert.Equal(t, t1.Id, int64(1))

		err = Insert(&t2)
		assert.NoError(t, err)

		err = Insert(&t3)
		assert.NoError(t, err)

		err = Insert(&t4)
		assert.NoError(t, err)

		err = Insert(&t5)
		assert.NoError(t, err)

		err = Insert(&t5)
		assert.Equal(t, err, ErrDupKey)
	})

	t.Run("list", func(t *testing.T) {
		records := make([]TestRecord, 0)
		err := List(&records, nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, 5, len(records))
	})

	t.Run("list | in", func(t *testing.T) {
		records := make([]TestRecord, 0)
		err := ListInById(&records, "id in ?", []int{1, 3})
		assert.NoError(t, err)
		assert.Equal(t, 2, len(records))
	})

	t.Run("get", func(t *testing.T) {
		var exp1 TestRecord
		err := Get(1, &exp1)
		assert.NoError(t, err)
		assert.Equal(t, exp1.Age, t1.Age)
		assert.Equal(t, exp1.Name, t1.Name)

		var notFound TestRecord
		err = Get(100, &notFound)
		assert.Error(t, err)
	})

	t.Run("update", func(t *testing.T) {
		u1 := TestRecord{Id: 1, Name: "tom-u1"}
		err := Update(1, &u1, nil)
		assert.NoError(t, err)

		var exp1 TestRecord
		err = Get(1, &exp1)
		assert.NoError(t, err)
		assert.Equal(t, exp1.Age, t1.Age)
		assert.Equal(t, exp1.Name, u1.Name)
		assert.True(t, exp1.UpdatedAt > time.Now().UnixMilli()-86400*1000)

		err = Db.Where("age = ?", 15).Updates(&TestRecord{Salary: 1024}).Error
		assert.NoError(t, err)
		var record2 TestRecord
		err = Get(4, &record2)
		assert.NoError(t, err)
		assert.Equal(t, 1024, record2.Salary)

	})

	t.Run("remove", func(t *testing.T) {
		err := Remove(1, new(TestRecord))
		assert.NoError(t, err)
		var exp1 TestRecord
		err = Get(1, &exp1)
		assert.Error(t, err)

		// remove by cond
		err = RemoveByCond(new(TestRecord), "name = ?", "non-existed-name-for-cond-remove")
		assert.NoError(t, err)

		err = RemoveByCond(new(TestRecord), "age = ?", 15)
		assert.NoError(t, err)

		records := make([]TestRecord, 0)
		err = List(&records, "age = ?", 15)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(records))
	})

}
