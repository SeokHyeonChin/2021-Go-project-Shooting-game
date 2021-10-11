package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletWidth  = 7
	bulletHeight = 15
	bulletSpeed  = 0.5
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	radius float64
	active bool
}

func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprites/player_bullet.bmp")
	bul.radius = 3
	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}
	x := bul.x - bulletWidth/2.0
	y := bul.y - bulletHeight/2.0

	renderer.Copy(bul.tex,
		&sdl.Rect{X: 140, Y: 0, W: 600, H: 1200},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletWidth, H: bulletHeight})
}

func (bul *bullet) update() {
	bul.x += bulletSpeed * math.Cos(bul.angle)
	bul.y += bulletSpeed * math.Sin(bul.angle)

	if bul.x > screenWidth || bul.x < 0 || bul.y > screenHeight || bul.y < 0 {
		bul.active = false
	}
}

var bulletPool []*bullet

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
