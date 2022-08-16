package translation_test

import (
	"errors"
	"testing"

	"github.com/holmes89/hello-api/translation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestRemoteServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteServiceTestSuite))
}

type RemoteServiceTestSuite struct {
	suite.Suite
	client    *MockHelloClient
	underTest *translation.RemoteService
}

func (suite *RemoteServiceTestSuite) SetupTest() {
	suite.client = new(MockHelloClient)
	suite.underTest = translation.NewRemoteService(suite.client)
}

type MockHelloClient struct {
	mock.Mock
}

func (m *MockHelloClient) Translate(word, language string) (string, error) {
	args := m.Called(word, language)
	return args.String(0), args.Error(1)
}

func (suite *RemoteServiceTestSuite) TestTranslate() {
	// Arrange
	suite.client.On("Translate", "foo", "bar").Return("baz", nil)

	// Act
	res := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.Equal(res, "baz")
	suite.client.AssertExpectations(suite.T())
}

func (suite *RemoteServiceTestSuite) TestTranslate_CaseSensitive() {
	// Arrange
	suite.client.On("Translate", "foo", "bar").Return("baz", nil)

	// Act
	res := suite.underTest.Translate("Foo", "bar")

	// Assert
	suite.Equal(res, "baz")
	suite.client.AssertExpectations(suite.T())
}

func (suite *RemoteServiceTestSuite) TestTranslate_Error() {
	// Arrange
	suite.client.On("Translate", "foo", "bar").Return("baz", errors.New("failure"))

	// Act
	res := suite.underTest.Translate("foo", "bar")

	// Assert
	suite.Equal(res, "")
	suite.client.AssertExpectations(suite.T())
}

func (suite *RemoteServiceTestSuite) TestTranslate_Cache() {
	// Arrange
	suite.client.On("Translate", "foo", "bar").Return("baz", nil).Times(1)

	// Act
	res1 := suite.underTest.Translate("foo", "bar")
	res2 := suite.underTest.Translate("Foo", "bar")

	// Assert
	suite.Equal(res1, "baz")
	suite.Equal(res2, "baz")
	suite.client.AssertExpectations(suite.T())
}
