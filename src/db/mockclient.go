package db

import (
	"model"

	"github.com/stretchr/testify/mock"
)

// MockBoltClient ...
type MockBoltClient struct {
	mock.Mock
}

// QueryAccount ...
func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

// Init ...
func (m *MockBoltClient) Init() {

}

// Seed ...
func (m *MockBoltClient) Seed() {

}

func (m *MockBoltClient) initBucket() {

}

func (m *MockBoltClient) seedAccounts() {

}
