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
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
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
