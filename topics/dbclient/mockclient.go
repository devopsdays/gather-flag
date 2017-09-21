package dbclient

import (
	"github.com/devopsdays/gather-flag/topics/model"
	"github.com/stretchr/testify/mock"
)

// MockMongoClient is a mock implementation of a datastore client for testing purposes.
type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) QueryTopic(topicID string) (model.Topic, error) {
	args := m.Mock.Called(topicID)
	return args.Get(0).(model.Topic), args.Error(1)
}

func (m *MockMongoClient) OpenMongoDB() {

}

func (m *MockMongoClient) Seed() {

}
