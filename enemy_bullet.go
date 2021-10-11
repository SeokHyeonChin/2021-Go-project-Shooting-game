package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	ebulletWidth  = 50
	ebulletHeight = 50
	ebulletSpeed  = 0.3
)

type ebullet struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	radius float64
	active bool
}

func newEBullet(renderer *sdl.Renderer) (ebul ebullet) {
	ebul.tex = textureFromBMP(renderer, "sprites/enemy_bullet.bmp")
	ebul.radius = 20
	return ebul
}

func (ebul *ebullet) draw(renderer *sdl.Renderer) {
	if !ebul.active {
		return
	}
	x := ebul.x - ebulletWidth/2.0
	y := ebul.y - ebulletHeight/2.0

	renderer.Copy(ebul.tex,
		&sdl.Rect{X: 0, Y: 0, W: 128, H: 128},
		&sdl.Rect{X: int32(x), Y: int32(y), W: ebulletWidth, H: ebulletHeight})
}

func (ebul *ebullet) update() {
	ebul.x += ebulletSpeed * math.Cos(ebul.angle)
	ebul.y += ebulletSpeed * math.Sin(ebul.angle)

	if ebul.x > screenWidth || ebul.x < 0 || ebul.y > screenHeight || ebul.y < 0 {
		ebul.active = false
	}
}

var ebulletPool []*ebullet

func initEBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		ebul := newEBullet(renderer)
		ebulletPool = append(ebulletPool, &ebul)
	}
}

func ebulletFromPool() (*ebullet, bool) {
	for _, ebul := range ebulletPool {
		if !ebul.active {
			return ebul, true
		}
	}
	return nil, false
}
