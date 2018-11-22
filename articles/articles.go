package articles

import (
	mesg "github.com/mesg-foundation/go-service"
)

// ArticlesService is a MESG service to manage articles.
type ArticlesService struct {
	s *mesg.Service

	// st is a storage for storing articles.
	st Storage
}

// New creates a new articles service with given mesg service and storage.
func New(service *mesg.Service, st Storage) (*ArticlesService, error) {
	s := &ArticlesService{
		s:  service,
		st: st,
	}
	return s, nil
}

// Start starts the service.
func (s *ArticlesService) Start() error {
	defer s.Close()
	return s.listenTasks()
}

// Close gracefully closes the articles service.
func (s *ArticlesService) Close() error {
	defer s.st.Close()
	return s.s.Close()
}
