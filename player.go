package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex  *sdl.Texture
	x, y float64
}

const (
	playerSpeed = 0.2
	playerSize  = 105
)

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = loadTextFromBMP(renderer, "sprites/player.bmp")
	p.x = screenWidth/2.0 - playerSize/2
	p.y = screenHeight - playerSize
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		nil,
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: playerSize, H: playerSize})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		p.y += playerSpeed
	} else if keys[sdl.SCANCODE_UP] == 1 {
		p.y -= playerSpeed
	}
}
