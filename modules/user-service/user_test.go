package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func last() User {
	return userStorage[idSeq]
}

func resetStorage() {
	idSeq = 0
	userStorage = make(map[int]User)
}

func TestInsertUser(t *testing.T) {
	defer resetStorage()

	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Email:     "singpor@gmail.com",
	}
	err := Insert(&u)
	assert.NoError(t, err)
	u.ID = 1
	assert.Equal(t, u, last())

	u = User{
		FirstName: "Kanokorn",
		LastName:  "Chongnguluam",
		Email:     "kanokorn@gmail.com",
	}
	err = Insert(&u)
	assert.NoError(t, err)
	u.ID = 2
	assert.Equal(t, u, last())
}

func TestUpdate(t *testing.T) {
	defer resetStorage()

	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Email:     "singpor@gmail.com",
	}
	err := Insert(&u)
	assert.NoError(t, err)

	u = last()
	u.Email = "singpor@outlook.com"
	err = Update(&u)
	assert.NoError(t, err)
	assert.Equal(t, u, last())
}

func TestFindByID(t *testing.T) {
	defer resetStorage()

	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Email:     "singpor@gmail.com",
	}
	err := Insert(&u)
	assert.NoError(t, err)

	r, err := FindByID(1)
	assert.NoError(t, err)
	assert.Equal(t, u, *r)

	r, err = FindByID(2)
	assert.Nil(t, r)
	assert.Error(t, err)
}

func TestDelete(t *testing.T) {
	defer resetStorage()
	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Email:     "singpor@gmail.com",
	}
	err := Insert(&u)
	assert.NoError(t, err)

	r, err := FindByID(1)
	assert.NoError(t, err)

	err = Delete(r)
	assert.NoError(t, err)

	err = Delete(r)
	assert.Error(t, err)
}

func TestAll(t *testing.T) {
	// Exercise
}

// func TestRaceCondition(t *testing.T) {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	insertFunc := func() {
// 		defer wg.Done()
// 		defer resetStorage()
// 		u := User{
// 			FirstName: "Weerasak",
// 			LastName:  "Chongnguluam",
// 		}
// 		Insert(&u)
// 		u.ID = 1
// 		assert.Equal(t, u, last())
// 	}
// 	go insertFunc()
// 	go insertFunc()
// 	wg.Wait()
// }
