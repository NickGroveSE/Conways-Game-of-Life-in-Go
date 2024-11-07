package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type Cell struct {
	Value            int
	ImpendingValue   int
	Coordinates      [2]int
	Neighbors        [8]*Cell
	NeighborSnapshot [8]int
}

// type Neighbors struct {
// 	Top      *Cell
// 	TopRight *Cell
// 	Right    *Cell
// 	BotRight *Cell
// 	Bot      *Cell
// 	BotLeft  *Cell
// 	Left     *Cell
// 	TopLeft  *Cell
// }

var (
	stepCount int = 0
	grid      [][]*Cell
)

const (
	TOP   int = -1
	BOT   int = 1
	LEFT  int = -1
	RIGHT int = 1
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.buildMatrix(50, 40)
}

func (a *App) buildMatrix(sizeX int, sizeY int) {

	matrix := make([][]*Cell, sizeY)
	for i := range matrix {
		matrix[i] = make([]*Cell, sizeX)
		for j := range matrix[i] {
			matrix[i][j] = &Cell{
				0,
				0,
				[2]int{i, j},
				[8]*Cell{nil},
				[8]int{0},
			}

		}
	}

	for i := range matrix {
		for j := range matrix[i] {

			// fmt.Printf("X: %d, Y: %d \n", i, j)
			// fmt.Println(len(matrix))

			if i != 0 {
				// fmt.Println("1")
				matrix[i][j].Neighbors[0] = matrix[i+TOP][j]
			}
			if i != 39 {
				// fmt.Println("2")
				matrix[i][j].Neighbors[4] = matrix[i+BOT][j]
			}
			if j != 0 {
				// fmt.Println("3")
				matrix[i][j].Neighbors[6] = matrix[i][j+LEFT]
			}
			if j != 49 {
				// fmt.Println("4")
				matrix[i][j].Neighbors[2] = matrix[i][j+RIGHT]
			}

			if matrix[i][j].Neighbors[0] != nil && matrix[i][j].Neighbors[6] != nil {
				// fmt.Println("5")
				matrix[i][j].Neighbors[7] = matrix[i+TOP][j+LEFT]
			}
			if matrix[i][j].Neighbors[0] != nil && matrix[i][j].Neighbors[2] != nil {
				// fmt.Println("6")
				matrix[i][j].Neighbors[1] = matrix[i+TOP][j+RIGHT]
			}
			if matrix[i][j].Neighbors[4] != nil && matrix[i][j].Neighbors[6] != nil {
				// fmt.Println("7")
				matrix[i][j].Neighbors[5] = matrix[i+BOT][j+LEFT]
			}
			if matrix[i][j].Neighbors[4] != nil && matrix[i][j].Neighbors[2] != nil {
				// fmt.Println("8")
				matrix[i][j].Neighbors[3] = matrix[i+BOT][j+RIGHT]
			}
		}
	}

	grid = matrix

}

func (c *Cell) snap() {

	for idx, cell := range c.Neighbors {
		if cell != nil {
			c.NeighborSnapshot[idx] = cell.Value
		} else {
			c.NeighborSnapshot[idx] = 0
		}
	}

}

func (a *App) Start() string {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		done := make(chan bool)

		for {
			select {
			case <-done:
				fmt.Println("Done!")
				return
			case <-ticker.C:
				stepCount++
				fmt.Println("Iterations: ", stepCount)
				runtime.EventsEmit(a.ctx, "Start", fmt.Sprintf("Iterations: %d", stepCount))
			}
		}
	}()
	return fmt.Sprintf("Iterations: %d", 0)
}

func (a *App) Next() string {

	stepCount++

	go func() {

		filledCells := make([][2]int, 0)

		for i := range grid {
			for j := range grid[i] {
				neighborCount := 0
				grid[i][j].snap()
				for _, value := range grid[i][j].NeighborSnapshot {
					if value == 1 {
						neighborCount++
					}

					if neighborCount == 4 {
						break
					}
				}

				if grid[i][j].Value == 1 {
					if neighborCount > 1 && neighborCount < 4 {
						grid[i][j].ImpendingValue = 1
					} else {
						grid[i][j].ImpendingValue = 0
					}
				} else {
					if neighborCount == 3 {
						grid[i][j].ImpendingValue = 1
					} else {
						grid[i][j].ImpendingValue = 0
					}
				}

				if grid[i][j].ImpendingValue == 1 {
					coordinates := [2]int{i, j}
					filledCells = append(filledCells, coordinates)
				}

			}
		}

		for i := range grid {
			for j := range grid[i] {
				grid[i][j].Value = grid[i][j].ImpendingValue
			}
		}

		runtime.EventsEmit(a.ctx, "Next", filledCells)

	}()

	return fmt.Sprintf("Iterations: %d", stepCount)
}

func (a *App) Reset() string {
	stepCount = 0

	go func() {
		a.buildMatrix(50, 40)

		runtime.EventsEmit(a.ctx, "Next", make([][2]int, 0))
	}()

	return fmt.Sprintf("Iterations: %d", stepCount)
}

func (a *App) StoreCell(x int, y int, filled bool) string {

	if filled {
		grid[y][x].Value = 1
		return fmt.Sprintf("Cell %d,%d is now filled", x, y)
	} else {
		grid[y][x].Value = 0
		return fmt.Sprintf("Cell %d,%d is now blank", x, y)
	}

}
