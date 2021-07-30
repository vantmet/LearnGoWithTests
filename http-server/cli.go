package poker

import "io"

type CLI struct {
	playerstore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayPoker() {
	cli.playerstore.RecordWin("Chris")
}
