package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

const (
	basicEnemySize = 105
)

func newBasicEnemy(renderer *sdl.Renderer) (be basicEnemy, err error) {
	img, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading basic enemy sprite: %v", err)
	}
	defer img.Free()
	be.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return basicEnemy{}, fmt.Errorf("loading texture from image: %v", err)
	}
	be.x = screenWidth/2.0 - basicEnemySize/2
	be.y = screenHeight/2.0 - basicEnemySize/2
	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	renderer.CopyEx(be.tex, nil,
		&sdl.Rect{X: int32(be.x), Y: int32(be.y), W: basicEnemySize, H: basicEnemySize},
		180, &sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2}, sdl.FLIP_NONE)
}
