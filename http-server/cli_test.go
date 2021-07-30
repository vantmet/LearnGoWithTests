package poker_test

import (
	"strings"
	"testing"

	poker "github.com/vantmet/LearnGoWithTests/http-server"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		input := strings.NewReader("Chris wins\n")

		cli := poker.NewCLI(playerStore, input)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		input := strings.NewReader("Cleo wins\n")

		cli := poker.NewCLI(playerStore, input)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

}
