package dbclient

import (
	"github.com/devopsdays/gather-flag/accounts/model"
	"github.com/stretchr/testify/mock"
)

// MockMongoClient is a mock implementation of a datastore client for testing purposes.
type MockMongoClient struct {
	mock.Mock
}

// QueryAccount provides a mock for the QueryAccount function
func (m *MockMongoClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

// OpenMongoDB provides a mock for the OpenMongoDB function
func (m *MockMongoClient) OpenMongoDB() {

}

// Seed provides a mock for the Seed function
func (m *MockMongoClient) Seed() {

}
