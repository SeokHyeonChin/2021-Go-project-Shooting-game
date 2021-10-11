package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 0.5
	playerWidth        = 90
	playerHeight       = 80
	playerShotCooldown = time.Millisecond * 250
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	radius   float64
	active   bool
	lastShot time.Time
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "sprites/player.bmp")
	p.x = screenWidth / 2.0
	p.y = screenHeight - playerHeight/1.5
	p.radius = 40
	p.active = true
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	if !p.active {
		return
	}
	x := p.x - playerWidth/2.0
	y := p.y - playerHeight/2.0

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 425, H: 325},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerWidth, H: playerHeight})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && p.x-(playerWidth/2.0) > 0 {
		p.x -= playerSpeed

	} else if keys[sdl.SCANCODE_RIGHT] == 1 && p.x+(playerWidth/2.0) < screenWidth {
		p.x += playerSpeed
	} else if keys[sdl.SCANCODE_UP] == 1 && p.y > 600 {
		p.y -= playerSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 && p.y+(playerHeight/2.0) < 800 {
		p.y += playerSpeed
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown {
			p.shoot()
			p.lastShot = time.Now()
		}
	}
}

func (p *player) shoot() {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true
		bul.x = p.x
		bul.y = p.y - playerHeight/2
		bul.angle = 270 * (math.Pi / 180)
	}
}
