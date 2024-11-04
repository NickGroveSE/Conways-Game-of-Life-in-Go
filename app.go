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

var (
	stepCount int = 0
	grid      [][]bool
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

	matrix := make([][]bool, sizeY)
	for i := range matrix {
		matrix[i] = make([]bool, sizeX)
		for j := range matrix[i] {
			matrix[i][j] = false
		}
	}

	grid = matrix

}

func (a *App) CountUp(name string) string {
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
				fmt.Println("Count: ", stepCount)
				runtime.EventsEmit(a.ctx, "CountUp", fmt.Sprintf("Count: %d", stepCount))
			}
		}
	}()
	return fmt.Sprintf("Count: %d", 0)
}

func (a *App) StoreCell(x int, y int, filled bool) string {

	grid[x][y] = filled

	if grid[x][y] {
		return fmt.Sprintf("Cell %d,%d is now filled", x, y)
	} else {
		return fmt.Sprintf("Cell %d,%d is now blank", x, y)
	}

}
