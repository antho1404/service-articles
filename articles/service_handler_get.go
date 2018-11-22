package articles

import (
	mesg "github.com/mesg-foundation/go-service"
)

const getSuccessOutputKey = "success"

type getInputs struct {
	// ID or human readable id of requested article.
	ID string `json:"id"`
}

type getSuccessOutputs struct {
	Articles []*Article `json:"articles"`
}

func (s *ArticlesService) getHandler(execution *mesg.Execution) (string, mesg.Data) {
	var inputs getInputs
	if err := execution.Data(&inputs); err != nil {
		return newErrorOutput(err)
	}

	article, err := s.st.Get(inputs.ID)
	if err != nil {
		return newErrorOutput(err)
	}

	return createSuccessOutputKey, createSuccessOutputs{article}
}
