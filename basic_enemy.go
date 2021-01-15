package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

const (
	basicEnemySize = 105
)

func newBasicEnemy(renderer *sdl.Renderer) (be basicEnemy) {
	be.tex = loadTextFromBMP(renderer, "sprites/basic_enemy.bmp")
	be.x = screenWidth/2.0 - basicEnemySize/2
	be.y = screenHeight/2.0 - basicEnemySize/2
	return be
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	renderer.CopyEx(be.tex, nil,
		&sdl.Rect{X: int32(be.x), Y: int32(be.y), W: basicEnemySize, H: basicEnemySize},
		180, &sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2}, sdl.FLIP_NONE)
}
