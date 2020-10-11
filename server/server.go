package server

type Server struct {
}

func New() (*Server, error) {
	s := Server{}
	return &s, nil
}
