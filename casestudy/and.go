package main

import (
	"fmt"
	"sync"
)

type Process struct {
	id                  int
	isInCriticalSection bool
	turn                int
	wantToEnter         []bool
	channel             chan bool
	replyCount          int
}

func NewProcess(id int, totalProcesses int) *Process {
	wantToEnter := make([]bool, totalProcesses)
	channel := make(chan bool, totalProcesses)
	return &Process{
		id:                  id,
		isInCriticalSection: false,
		turn:                0,
		wantToEnter:         wantToEnter,
		channel:             channel,
		replyCount:          0,
	}
}

func (p *Process) EnterCriticalSection(processes []*Process) {
	p.wantToEnter[p.id] = true
	p.turn = 1 - p.turn
	p.replyCount = 0

	for i := 0; i < len(processes); i++ {
		if i == p.id {
			continue
		}
		go func(proc *Process) {
			proc.channel <- true
		}(processes[i])
	}

	for i := 0; i < len(processes)-1; i++ {
		<-p.channel
	}

	p.isInCriticalSection = true
	fmt.Printf("Process %d entered the critical section.\n", p.id)
	// Perform critical section operations
	p.ExitCriticalSection(processes)
}

func (p *Process) ExitCriticalSection(processes []*Process) {
	p.isInCriticalSection = false
	p.wantToEnter[p.id] = false
	p.replyCount = 0

	for i := 0; i < len(processes); i++ {
		if i == p.id {
			continue
		}
		go func(proc *Process) {
			proc.channel <- false
		}(processes[i])
	}
}

func (p *Process) Listen(processes []*Process) {
	for {
		<-p.channel
		if p.wantToEnter[p.id] && (p.turn == p.id) {
			p.replyCount++
			if p.replyCount == len(processes)-1 {
				p.isInCriticalSection = true
				fmt.Printf("Process %d entered the critical section.\n", p.id)
				// Perform critical section operations
				p.ExitCriticalSection(processes)
			}
		} else {
			p.channel <- false
		}
	}
}

func main() {
	totalProcesses := 3

	processes := make([]*Process, totalProcesses)
	for i := 0; i < totalProcesses; i++ {
		processes[i] = NewProcess(i, totalProcesses)
		go processes[i].Listen(processes)
	}

	var wg sync.WaitGroup
	wg.Add(totalProcesses)

	for i := 0; i < totalProcesses; i++ {
		go func(p *Process) {
			defer wg.Done()
			p.EnterCriticalSection(processes)
		}(processes[i])
	}

	wg.Wait()
}
