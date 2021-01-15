package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	active bool
}

func newBullet(renderer *sdl.Renderer, x, y float64) (b bullet) {
	b.tex = loadTextFromBMP(renderer, "sprites/player_bullet.bmp")
	b.x = x
	b.y = y
	return b
}

func (bul bullet) draw(renderer *sdl.Renderer, x, y float64) {
	bul.tex = loadTextFromBMP(renderer, "sprites/player_bullet.bmp")
	bul.x = x
	bul.y = y
}

func (bul bullet) update() {

}

var bullets = []bullet{}

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bullets = append(bullets, newBullet(renderer, 0, 0))
	}
}

func bulletFromPool() (b bullet, ok bool) {
	for _, b := range bullets {
		if !b.active {
			return b, true
		}
	}
	return bullet{}, false
}
