package service

import (
	"lab/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockMethodWithArfs(t *testing.T) {
	mockDB := mocks.DB{}
	mockDB.On("FetchMessage", "sg").Return("lah", nil)
	g := greeter{&mockDB, "sg"}
	assert.Equal(t, "Message is: lah", g.Greet())
	mockDB.AssertExpectations(t)
}

func TestMockMethodWithArgsIngoreArgs(t *testing.T) {
	mockDB := mocks.DB{}
	mockDB.On("FetchMessage", mock.Anything).Return("lah", nil)
	g := greeter{&mockDB, "in"}
	assert.Equal(t, "Message is: lah", g.Greet())
	mockDB.FetchMessage("in")
	mockDB.AssertCalled(t, "FetchMessage", "in")
	mockDB.AssertNotCalled(t, "FetchMessage", "ch")
	mockDB.AssertExpectations(t)
}

func TestMatchedBy(t *testing.T) {
	mockDB := mocks.DB{}
	mockDB.On("FetchMessage", mock.MatchedBy(
		func(lang string) bool {
			return lang[0] == 'i'
		},
	)).Return("bzzzz", nil)
	g := greeter{&mockDB, "izz"}
	msg := g.Greet()
	assert.Equal(t, "Message is: bzzzz", msg)
	mockDB.AssertExpectations(t)
}
