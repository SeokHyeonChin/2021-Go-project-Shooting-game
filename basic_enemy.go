package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemyWidth  = 100
	basicEnemyHeight = 100
)

type basicEnemy struct {
	tex      *sdl.Texture
	x, y     float64
	radius   float64
	Speed    float64
	left_hit bool
	active   bool
}

func newBasicEnemy(renderer *sdl.Renderer, x, y, speed float64) (be basicEnemy) {
	be.tex = textureFromBMP(renderer, "sprites/basic_enemy.bmp")
	be.x = x
	be.y = y
	be.Speed = speed
	be.radius = 25
	be.active = true
	return be
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	if !be.active {
		return
	}
	x := be.x - basicEnemyWidth/2.0
	y := be.y - basicEnemyHeight/2.0

	renderer.Copy(be.tex,
		&sdl.Rect{X: 120, Y: 55, W: 480, H: 545},
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemyWidth, H: basicEnemyHeight})

}

func (be *basicEnemy) update() {
	if !be.left_hit {
		be.x -= be.Speed
		if be.x-(playerWidth/2.0) <= 0 {
			be.left_hit = true
		}
	} else {
		be.x += be.Speed
		if be.x+(playerWidth/2.0) >= screenWidth {
			be.left_hit = false
		}
	}
}

var enemies []*basicEnemy

func initBasicEnemy(renderer *sdl.Renderer) {
	Speed := 0.5
	for i := 0; i < 5; i++ {
		for j := 1; j < 4; j++ {
			var x float64
			if j%2 == 0 {
				x = (float64(i)/5)*screenWidth + basicEnemyWidth/1.8
			} else {
				x = (float64(i)/5)*screenWidth + basicEnemyWidth/1.1
			}

			y := float64(j)*basicEnemyHeight + basicEnemyHeight*1.4

			enemy := newBasicEnemy(renderer, x, y, Speed-float64(j)*0.1)
			enemies = append(enemies, &enemy)
		}
	}
}
