package articles

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateSuccess(t *testing.T) {
	var (
		inputs = createInputs{
			Article: &Article{
				Title:    "Introduction to MESG — Chapter 0/1",
				Content:  "If you’re already familiar with MESG, you can...",
				Location: Location{Country: "Turkey", City: "Ankara"},
			},
		}
		a, sm, server = newTesting(t)
	)

	go server.Start()
	go a.Start()

	sm.On("Save", mock.Anything).Once().Return(nil)

	_, execution, err := server.Execute("create", inputs)
	require.NoError(t, err)
	require.Equal(t, "success", execution.Key())

	var outputs createSuccessOutputs
	require.NoError(t, execution.Data(&outputs))
	require.NotZero(t, outputs.Article.ID)
	require.Equal(t, inputs.Article.Title, outputs.Article.Title)
	require.Equal(t, inputs.Article.Content, outputs.Article.Content)
	require.Equal(t, "introduction-to-mesg-chapter-0-1", outputs.Article.Path)
	require.Equal(t, inputs.Article.Location.Country, outputs.Article.Location.Country)
	require.Equal(t, inputs.Article.Location.City, outputs.Article.Location.City)
	require.NotZero(t, outputs.Article.CreatedAt)

	sm.AssertExpectations(t)
}

func TestCreateError(t *testing.T) {
	var (
		inputs = createInputs{
			Article: &Article{},
		}
		a, sm, server = newTesting(t)
	)
	go server.Start()
	go a.Start()

	_, execution, err := server.Execute("create", inputs)
	require.NoError(t, err)
	require.Equal(t, "error", execution.Key())

	var outputs errorOutput
	require.NoError(t, execution.Data(&outputs))
	require.Contains(t, outputs.Message, "Article.Title")
	require.Contains(t, outputs.Message, "Article.Content")

	sm.AssertExpectations(t)
}
