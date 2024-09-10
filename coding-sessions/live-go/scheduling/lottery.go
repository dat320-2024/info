package scheduling

import (
	"math/rand"
	"time"
)

type Lottery struct {
	PID     int
	Tickets int
}

func (l *Lottery) String() string {
	return "PID: " + string(l.PID) + ", Tickets: " + string(l.Tickets)
}

func NewLottery(pid, tickets int) *Lottery {
	return &Lottery{PID: pid, Tickets: tickets}
}

func GenerateRandomTicket(lotteries []*Lottery) int {
	totalTickets := 0
	for _, lottery := range lotteries {
		totalTickets += lottery.Tickets
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(totalTickets)
}
func FindWinner(number int, lotteries []*Lottery) int {
	counter := 0
	for _, lottery := range lotteries {
		counter += lottery.Tickets
		if number < counter {
			return lottery.PID
		}

	}
	return -1 // No winner found
}

func Schedule(slots int, lotteries []*Lottery) map[int]int {
	scheduleMap := make(map[int]int)

	for i := 0; i < len(lotteries); i++ {
		scheduleMap[lotteries[i].PID] = 0
	}

	for i := 0; i < slots; i++ {
		random := GenerateRandomTicket(lotteries)
		winner := FindWinner(random, lotteries)
		scheduleMap[winner]++
	}

	return scheduleMap
}
