package user

import (
	"sync"
)

type User struct {
	Id    int64
	Name  string
	Level int32
	Xp    int32
}

type UserStorage struct {
	mu     sync.Mutex
	users  map[int64]*User
	nextID int64
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		users:  make(map[int64]*User),
		nextID: 1,
	}
}

func (s *UserStorage) CreateUser(name string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &User{
		Id:    s.nextID,
		Name:  name,
		Level: 1,
		Xp:    0,
	}

	s.users[s.nextID] = user
	s.nextID++

	return user
}
func (s *UserStorage) GetUser(id int64) *User {
	user, ok := s.users[id]
	if !ok {
		return &User{}
	}

	return user
}
func (s *UserStorage) ListUsers() []User {
	userslist := make([]User, 0, len(s.users))

	for _, user := range s.users {
		userslist = append(userslist, *user)
	}

	return userslist
}
