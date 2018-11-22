package articles

import (
	mesg "github.com/mesg-foundation/go-service"
)

const listSuccessOutputKey = "success"

type listSuccessOutputs struct {
	Articles []*Article `json:"articles"`
}

func (s *ArticlesService) listHandler(execution *mesg.Execution) (string, mesg.Data) {
	articles, err := s.st.List()
	if err != nil {
		return newErrorOutput(err)
	}
	return listSuccessOutputKey, listSuccessOutputs{articles}
}
