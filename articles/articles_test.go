package articles

import (
	"testing"

	mesg "github.com/mesg-foundation/go-service"
	"github.com/mesg-foundation/go-service/mesgtest"
	"github.com/stretchr/testify/require"
)

const (
	token    = "token"
	endpoint = "endpoint"
)

func newTesting(t *testing.T) (*ArticlesService, *MockStorage, *mesgtest.Server) {
	testServer := mesgtest.NewServer()
	service, err := mesg.New(
		mesg.DialOption(testServer.Socket()),
		mesg.TokenOption(token),
		mesg.EndpointOption(endpoint),
	)
	require.NoError(t, err)
	require.NotNil(t, service)

	ms := &MockStorage{}

	a, err := New(service, ms)
	require.NoError(t, err)

	return a, ms, testServer
}
