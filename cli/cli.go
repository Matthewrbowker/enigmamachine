package cli

type CLI struct {
}

func New() CLI {
	return CLI{}
}

func (c CLI) Exec() int {
	return 0
}
