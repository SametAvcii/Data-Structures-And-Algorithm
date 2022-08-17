package solid

import "fmt"


type User struct {
	ID   string
	Name string
}

type Storage interface {
	Add(key string, item interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

type storage struct {
	items map[string]interface{}
}

func NewStorage() Storage {
	return &storage{
		items: make(map[string]interface{}),
	}
}

func (s *storage) Add(key string, item interface{}) error {
	s.items[key] = item
	return nil
}

func (s *storage) Get(key string) (interface{}, error) {
	user, ok := s.items[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return user, nil
}

func (s *storage) Delete(key string) error {
	_, ok := s.items[key]
	if !ok {
		return fmt.Errorf("not found")
	}
	delete(s.items, key)
	return nil
}

type Manager struct {
	storage Storage
}

func NewManager(storage Storage) *Manager {
	return &Manager{
		storage: storage,
	}
}

func (m *Manager) AddUser(user *User) error {
	if user == nil || user.ID == "" {
		return fmt.Errorf("invalid user")
	}
	return m.storage.Add(user.ID, user)
}

func (m *Manager) GetUser(id string) (*User, error) {
	item, err := m.storage.Get(id)
	if err != nil {
		return nil, err
	}
	user, ok := item.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid user data")
	}
	return user, nil
}

func (m *Manager) DeleteUser(id string) error {
	return m.storage.Delete(id)
}
