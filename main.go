package main

import (
	sdl "github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic("Failed to initialize window")
	}
	
}