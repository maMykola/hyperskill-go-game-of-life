package main

import (
	"fmt"
	"math/rand"
)

type State byte
type UniverseState [][]State
type Universe struct {
	Generation int
	Size       int
	State      UniverseState
}

const (
	Alive State = 'O'
	Dead  State = ' '
)

func main() {
	var size int

	fmt.Scan(&size)

	universe := buildUniverse(size)
	for i := 0; i < 10; i++ {
		universe.NextGeneration()
		universe.Print()
	}
}

func buildUniverse(size int) *Universe {
	cells := make(UniverseState, size)
	for i := 0; i < size; i++ {
		cells[i] = make([]State, size)
		for j := 0; j < size; j++ {
			if rand.Intn(2) == 1 {
				cells[i][j] = Alive
			} else {
				cells[i][j] = Dead
			}
		}
	}

	universe := Universe{
		Size:  size,
		State: cells,
	}

	return &universe
}

func (u *Universe) Print() {
	fmt.Printf("Generation #%d\n", u.Generation)
	fmt.Printf("Alive: %d\n", u.Count(Alive))
	for i := 0; i < u.Size; i++ {
		fmt.Println(string(u.State[i]))
	}
}

func (u *Universe) NextGeneration() {
	u.Generation++
	currentState := u.GetState()
	for x := 0; x < u.Size; x++ {
		for y := 0; y < u.Size; y++ {
			u.State[x][y] = currentState.NextGen(x, y)
		}
	}
}

func (u *Universe) Count(s State) int {
	var total int
	for i := 0; i < u.Size; i++ {
		for j := 0; j < u.Size; j++ {
			if u.State[i][j] == s {
				total++
			}
		}
	}
	return total
}

func (u *Universe) GetState() UniverseState {
	state := make(UniverseState, u.Size)
	for i := 0; i < u.Size; i++ {
		state[i] = make([]State, u.Size)
		copy(state[i], u.State[i])
	}
	return state
}

func (us *UniverseState) NextGen(x, y int) State {
	var total int

	total += us.CountNeighbor(x-1, y-1)
	total += us.CountNeighbor(x-1, y)
	total += us.CountNeighbor(x-1, y+1)
	total += us.CountNeighbor(x, y-1)
	total += us.CountNeighbor(x, y+1)
	total += us.CountNeighbor(x+1, y-1)
	total += us.CountNeighbor(x+1, y)
	total += us.CountNeighbor(x+1, y+1)

	switch (*us)[x][y] {
	case Alive:
		if total == 2 || total == 3 {
			return Alive
		}
	case Dead:
		if total == 3 {
			return Alive
		}
	}

	return Dead
}

// CountNeighbor return 1 if the cell is alive at (x, y) otherwise return 0
func (us *UniverseState) CountNeighbor(x, y int) int {
	size := len(*us)

	if x < 0 {
		x = size - 1
	} else if x >= size {
		x = 0
	}

	if y < 0 {
		y = size - 1
	} else if y >= size {
		y = 0
	}

	if (*us)[x][y] == Alive {
		return 1
	}

	return 0
}
