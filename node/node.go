package node

type Node struct {
}

func New() (*Node, error) {
	n := Node{}
	return &n, nil
}
