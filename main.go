package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	height = 20
	width  = 10
)

const (
	RefreshRate = 100.0 // 画面刷新帧率
	UpdateRate  = 200.0 // 游戏状态更新帧率
)

func main() {
	// 初始化termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	game := &Game{}
	game.Init(height, width)
	game.Render()

	// 创建一个定时器，用于定时刷新屏幕
	refreshRateTicker := time.NewTicker(1000.0 / RefreshRate * time.Millisecond)
	defer refreshRateTicker.Stop()

	// 创建一个定时器，用于定时更新游戏状态
	gameStepTicker := time.NewTicker(1000.0 / UpdateRate * time.Millisecond)
	defer gameStepTicker.Stop()

	// 使用channel来处理事件
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
loop:
	for {
		select {
		case ev := <-eventQueue: // 从channel中取出事件
			if ev.Type == termbox.EventKey {
				if !game.pause {
					switch {
					case ev.Ch == 'W' || ev.Ch == 'w' || ev.Key == termbox.KeyArrowUp:
						game.Up()
					case ev.Ch == 'S' || ev.Ch == 's' || ev.Key == termbox.KeyArrowDown:
						game.Down()
					case ev.Ch == 'A' || ev.Ch == 'a' || ev.Key == termbox.KeyArrowLeft:
						game.Left()
					case ev.Ch == 'D' || ev.Ch == 'd' || ev.Key == termbox.KeyArrowRight:
						game.Right()
					case ev.Ch == 'P' || ev.Ch == 'p':
						game.Pause()
					case ev.Key == termbox.KeyEsc:
						fmt.Println("游戏退出")
						break loop
					}
				} else {
					if ev.Ch == 'P' || ev.Ch == 'p' {
						game.StopPause()
					}
				}
			} else if ev.Type == termbox.EventError {
				panic(ev.Err)
			}
		case <-refreshRateTicker.C: // 画面渲染定时器事件
			// 重新渲染游戏状态
			game.Clear()
			game.Render()
		case <-gameStepTicker.C: // 游戏状态更新定时器事件
			// 更新游戏状态
			if !game.pause {
				gameOver := game.Step()
				if gameOver {
					fmt.Println("游戏结束")
					break loop
				}
			}
		}
	}
}
