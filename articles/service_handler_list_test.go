package articles

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/stretchr/testify/require"
)

func TestListSuccess(t *testing.T) {
	var (
		articles = []*Article{
			{
				ID:    bson.NewObjectId(),
				Title: "1",
			},
			{
				ID:    bson.NewObjectId(),
				Title: "2",
			},
		}
		a, sm, server = newTesting(t)
	)

	go server.Start()
	go a.Start()

	sm.On("List").Once().Return(articles, nil)

	_, execution, err := server.Execute("list", nil)
	require.NoError(t, err)
	require.Equal(t, "success", execution.Key())

	var outputs listSuccessOutputs
	require.NoError(t, execution.Data(&outputs))
	require.Equal(t, articles, outputs.Articles)

	sm.AssertExpectations(t)
}
