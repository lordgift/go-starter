package user

import (
	"errors"
	"sort"
)

var ErrNotFound = errors.New("user: not found")

var (
	userStorage = make(map[int]User)
	idSeq       int
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func Insert(user *User) error {
	idSeq++
	user.ID = idSeq
	userStorage[idSeq] = *user
	return nil
}

func Update(user *User) error {
	if _, ok := userStorage[user.ID]; !ok {
		return ErrNotFound
	}
	userStorage[user.ID] = *user
	return nil
}

func Delete(user *User) error {
	if _, ok := userStorage[user.ID]; !ok {
		return ErrNotFound
	}
	delete(userStorage, user.ID)
	return nil
}

func FindByID(id int) (*User, error) {
	if _, ok := userStorage[id]; !ok {
		return nil, ErrNotFound
	}
	user := userStorage[id]
	return &user, nil
}

func All() ([]User, error) {
	users := make([]User, 0, len(userStorage))
	for _, v := range userStorage {
		users = append(users, v)
	}
	sort.Slice(users, func(i, j int) bool { return users[i].ID < users[j].ID })
	return users, nil
}
