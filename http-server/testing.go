package poker

import (
	"fmt"
	"io"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatal("expected a win call but didn't get any")
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q, want %q", store.winCalls[0], winner)
	}
}

func AssertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

// League helpers
func AssertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Assert Helper for body
func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

//Assert helper for Status
func AssertStatus(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Code
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
	To     io.Writer
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

func AssertScheduledAlert(t testing.TB, got ScheduledAlert, want ScheduledAlert) {
	amountGot := got.Amount
	amountWant := want.Amount
	if amountGot != amountWant {
		t.Errorf("got amount %d, want %d", amountGot, amountWant)
	}

	gotScheduledTime := got.At
	wantScheduledTime := want.At
	if gotScheduledTime != wantScheduledTime {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, wantScheduledTime)
	}
}

type GameSpy struct {
	StartCalled bool
	StartedWith int
	BlindAlert  []byte

	FinishedCalled bool
	FinishedWith   string
}

func (g *GameSpy) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
	alertsDestination.Write(g.BlindAlert)

}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}
