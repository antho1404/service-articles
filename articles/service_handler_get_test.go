package articles

import (
	"errors"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/stretchr/testify/require"
)

func TestGetSuccess(t *testing.T) {
	var (
		article = &Article{
			ID:    bson.NewObjectId(),
			Title: "1",
		}
		inputs = getInputs{
			ID: article.ID.String(),
		}
		a, sm, server = newTesting(t)
	)

	go server.Start()
	go a.Start()

	sm.On("Get", article.ID.String()).Once().Return(article, nil)

	_, execution, err := server.Execute("get", inputs)
	require.NoError(t, err)
	require.Equal(t, "success", execution.Key())

	var outputs createSuccessOutputs
	require.NoError(t, execution.Data(&outputs))
	require.Equal(t, article, outputs.Article)

	sm.AssertExpectations(t)
}

func TestGetError(t *testing.T) {
	var (
		article = &Article{
			ID: bson.NewObjectId(),
		}
		inputs = getInputs{
			ID: article.ID.String(),
		}
		notFoundErr   = errors.New("not found")
		a, sm, server = newTesting(t)
	)
	go server.Start()
	go a.Start()

	sm.On("Get", article.ID.String()).Once().Return(nil, notFoundErr)

	_, execution, err := server.Execute("get", inputs)
	require.NoError(t, err)
	require.Equal(t, "error", execution.Key())

	var outputs errorOutput
	require.NoError(t, execution.Data(&outputs))
	require.Equal(t, notFoundErr.Error(), outputs.Message)

	sm.AssertExpectations(t)
}
