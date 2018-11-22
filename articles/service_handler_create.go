package articles

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gosimple/slug"
	mesg "github.com/mesg-foundation/go-service"
	validator "gopkg.in/go-playground/validator.v9"
)

const createSuccessOutputKey = "success"

type createInputs struct {
	Article *Article `json:"article" validator:"required"`
}

type createSuccessOutputs struct {
	Article *Article `json:"article"`
}

func (s *ArticlesService) createHandler(execution *mesg.Execution) (string, mesg.Data) {
	var inputs createInputs
	if err := execution.Data(&inputs); err != nil {
		return newErrorOutput(err)
	}

	if err := validator.New().Struct(inputs.Article); err != nil {
		return newErrorOutput(err)
	}

	article := inputs.Article
	article.ID = bson.NewObjectId()
	article.Path = slug.Make(article.Title)
	article.CreatedAt = time.Now()

	if err := s.st.Save(article); err != nil {
		return newErrorOutput(err)
	}

	return createSuccessOutputKey, createSuccessOutputs{article}
}
