package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	fmt.Println("---------- STARTING APPLICATION 🏃----------")
	fmt.Println(" -> INITIALIZING SDL 🌏 ..............")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Error initializing SDL 🛑: ", err)
		return
	}

	fmt.Println("------------- SDL INITIALIZED 🌱 -----------")

	window, err := sdl.CreateWindow(
		"Gaming in go",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL,
	)

	if err != nil {
		fmt.Println("Error initializing window 🛑: ", err)
		return
	}

	defer window.Destroy()

	renderer, err := 	sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error initializing renderer 🛑: ", err)
		return
	}
	defer renderer.Destroy()
}
