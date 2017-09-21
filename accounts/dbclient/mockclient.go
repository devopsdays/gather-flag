package dbclient

import (
	"github.com/devopsdays/gather-flag/accounts/model"
	"github.com/stretchr/testify/mock"
)

// MockMongoClient is a mock implementation of a datastore client for testing purposes.
type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockMongoClient) OpenMongoDB() {

}

func (m *MockMongoClient) Seed() {

}
