package dbclient

import (
	"github.com/devopsdays/gather-flag/rooms/model"
	"github.com/stretchr/testify/mock"
)

// MockMongoClient is a mock implementation of a datastore client for testing purposes.
type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) QueryRoom(roomID string) (model.Room, error) {
	args := m.Mock.Called(roomID)
	return args.Get(0).(model.Room), args.Error(1)
}

func (m *MockMongoClient) OpenMongoDB() {

}

func (m *MockMongoClient) Seed() {

}
