package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	playerStore := &StubPlayerStore{}
	input := strings.NewReader("Chris Wins\n")

	cli := &CLI{playerStore, input}
	cli.PlayPoker()

	assertPlayerWin(t, playerStore, "Chris")
}

func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q, want %q", store.winCalls[0], winner)
	}
}
