package scheduling

import (
	"testing"
)

func TestFindWinner(t *testing.T) {
	lotteries := []*Lottery{
		{PID: 1, Tickets: 5},
		{PID: 2, Tickets: 10},
		{PID: 3, Tickets: 3},
	}
	expected := map[int]int{
		7:  2,
		2:  1,
		17: 3,
		20: -1,
		15: 3,
	}
	for number, expectedWinner := range expected {
		actualWinner := FindWinner(number, lotteries)
		if actualWinner != expectedWinner {
			t.Errorf("Expected winner to be %d, but got %d", expectedWinner, actualWinner)
		}
	}

}

func TestFindWinner2(t *testing.T) {
	lotteries := []*Lottery{
		{PID: 1, Tickets: 5},
		{PID: 2, Tickets: 10},
		{PID: 3, Tickets: 3},
	}
	expected := []struct{ in, want int }{
		{7, 2},
		{2, 1},
		{17, 3},
		{20, -1},
		{15, 3},
	}
	for _, expectedWinner := range expected {
		got := FindWinner(expectedWinner.in, lotteries)
		if got != expectedWinner.want {
			t.Errorf("Random number is %d, got winner PID %d, should be PID %d", expectedWinner.in, got, expectedWinner.want)
		}
	}
}
