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
	fmt.Println(" -> Creating window 🌏 ..............")
 
	window, err := sdl.CreateWindow(
		"Space Shooter 🌠",
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
	fmt.Println("------------- WINDOW CREATED 🌱 -----------")

	defer window.Destroy()

	fmt.Println(" -> Creating renderer 🌏 ..............")
	renderer, err := 	sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error initializing renderer 🛑: ", err)
		return
	}
	fmt.Println("------------- RENDERER CREATED 🌱 -----------")

	defer renderer.Destroy()
}
