package service

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/techschool/pcbook-go/pb"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

type Store interface {
	//// Save saves the laptop to the store
	Save(laptop *pb.Laptop) error
	//// Find finds a laptop by ID
	//Find(id string) (*pb.Laptop, error)
	//// Search searches for laptops with filter, returns one by one via the found function
	//Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

// Save saves the laptop to the store
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[other.Id] = other
	return nil
}
