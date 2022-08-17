package solid

import "fmt"

type User struct {
	ID   string
	Name string
}

type Storage struct {
	items map[string]interface{}
}

// NewStorage provides a Storage interface.
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]interface{}),
	}
}

// Add adds a new item.
func (s *Storage) Add(key string, item interface{}) error {
	s.items[key] = item
	return nil
}

// Get retrieves an item.
func (s *Storage) Get(key string) (interface{}, error) {
	user, ok := s.items[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return user, nil
}

// Delete deletes an item.
func (s *Storage) Delete(key string) (interface{}, error) {
	item, ok := s.items[key]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	delete(s.items, key)
	return item, nil
}

type userStorage interface {
	add(key string, item interface{}) error
	get(key string) (interface{}, error)
	delete(key string) error
}

type userStorageHandler struct {
	storage *Storage
}

func (s *userStorageHandler) add(key string, item interface{}) error {
	return s.storage.Add(key, item)
}

func (s *userStorageHandler) get(key string) (interface{}, error) {
	item, err := s.storage.Get(key)
	return item, err
}

func (s *userStorageHandler) delete(key string) error {
	_, err := s.storage.Delete(key)
	return err
}

// Manager defines the user data manager.
type Manager struct {
	storage userStorage
}

// NewManager creates a new Manager.
func NewManager(storage *Storage) *Manager {
	return &Manager{
		storage: &userStorageHandler{
			storage: storage,
		},
	}
}

func (m *Manager) AddUser(user *User) error {
	if user == nil || user.ID == "" {
		return fmt.Errorf("invalid user")
	}
	return m.storage.add(user.ID, user)
}

// GetUser retrieves a user data.
func (m *Manager) GetUser(id string) (*User, error) {
	item, err := m.storage.get(id)
	if err != nil {
		return nil, err
	}
	user, ok := item.(*User)
	if !ok {
		return nil, fmt.Errorf("invalid user data")
	}
	return user, nil
}

// DeleteUser deletes a user data.
func (m *Manager) DeleteUser(id string) error {
	return m.storage.delete(id)
}