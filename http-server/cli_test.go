package poker_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	poker "github.com/vantmet/LearnGoWithTests/http-server"
)

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

type SpyBlindAlerter struct {
	alerts []poker.ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, poker.ScheduledAlert{at, amount, os.Stdout})
}

func TestCLI(t *testing.T) {
	t.Run("It returns an error if invalid text is input after game start", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("3", "Lloyd is a killer")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStarted(t, game, 3)
		assertFinishCalledWith(t, game, poker.BadWinnerInputErrMsg)
	})
	t.Run("start game with 3 players and finish with Chris as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("3", "Chris wins")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStarted(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
	t.Run("start game with 8 players and finish with Chloe as winner", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("8", "Chloe wins")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStarted(t, game, 8)
		assertFinishCalledWith(t, game, "Chloe")

	})
	t.Run("it prompts the user to enter a number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessageSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStarted(t, game, 7)
	})

	t.Run("it prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("Pies")
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessageSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertMessageSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertGameStarted(t testing.TB, game *poker.GameSpy, numberOfPlayers int) {
	if game.StartedWith != numberOfPlayers {
		t.Errorf("wanted a Start called with %d, but got %d", numberOfPlayers, game.StartedWith)
	}
}

func assertGameNotStarted(t testing.TB, game *poker.GameSpy) {
	if game.StartCalled {
		t.Errorf("Game should not have started.")
	}
}

func assertFinishCalledWith(t testing.TB, game *poker.GameSpy, winner string) {
	if game.FinishedWith != winner {
		t.Errorf("wanted a Finish called with %s, but got %s", winner, game.FinishedWith)
	}
}

func userSends(input ...string) io.Reader {
	str := ""
	for i, v := range input {
		str += v
		str += "\n"
		i++
	}
	return strings.NewReader(str)
}
