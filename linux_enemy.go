package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	linuxEnemyWidth      = 140
	linuxEnemyHeight     = 160
	linuxEnemySpeed      = 0.2
	linuxAttackTime      = 3000
	linuxSpecialTime     = 6000
	linuxSuperAttackTime = 10000
)

type linuxEnemy struct {
	tex          *sdl.Texture
	x, y         float64
	radius       float64
	left_hit     bool
	active       bool
	life         float64
	attack_time  float64
	special_time float64
	super_time   float64
}

func newLinuxEnemy(renderer *sdl.Renderer, x, y float64) (le linuxEnemy) {
	le.tex = textureFromBMP(renderer, "sprites/linux_enemy.bmp")
	le.x = x
	le.y = y
	le.radius = 50
	le.active = true
	le.life = 10
	le.attack_time = linuxAttackTime
	le.special_time = linuxSpecialTime
	le.super_time = linuxSuperAttackTime
	return le
}

func (le *linuxEnemy) draw(renderer *sdl.Renderer) {
	if !le.active {
		return
	}
	x := le.x - linuxEnemyWidth/2.0
	y := le.y - linuxEnemyHeight/2.0

	renderer.Copy(le.tex,
		&sdl.Rect{X: 0, Y: 0, W: 1057, H: 1280},
		&sdl.Rect{X: int32(x), Y: int32(y), W: linuxEnemyWidth, H: linuxEnemyHeight})

}

func (le *linuxEnemy) update() {
	if !le.left_hit {
		le.x -= linuxEnemySpeed
		if le.x-(playerWidth/2.0) <= 0 {
			le.left_hit = true
		}
	} else {
		le.x += linuxEnemySpeed
		if le.x+(playerWidth/2.0) >= screenWidth {
			le.left_hit = false
		}
	}
	le.attack_time -= 1
	le.special_time -= 1
	le.super_time -= 1
	if le.attack_time <= 0 {
		le.enemy_attack(85 * (math.Pi / 180))
		le.enemy_attack(90 * (math.Pi / 180))
		le.enemy_attack(95 * (math.Pi / 180))
		le.attack_time = linuxAttackTime
	}
	if le.special_time <= 0 {
		le.enemy_attack(80 * (math.Pi / 180))
		le.enemy_attack(85 * (math.Pi / 180))
		le.enemy_attack(90 * (math.Pi / 180))
		le.enemy_attack(95 * (math.Pi / 180))
		le.enemy_attack(100 * (math.Pi / 180))
		le.special_time = linuxSpecialTime
	}
	if le.super_time <= 0 {
		le.enemy_attack(35 * (math.Pi / 180))
		le.enemy_attack(50 * (math.Pi / 180))
		le.enemy_attack(65 * (math.Pi / 180))
		le.enemy_attack(80 * (math.Pi / 180))
		le.enemy_attack(85 * (math.Pi / 180))
		le.enemy_attack(90 * (math.Pi / 180))
		le.enemy_attack(95 * (math.Pi / 180))
		le.enemy_attack(100 * (math.Pi / 180))
		le.enemy_attack(115 * (math.Pi / 180))
		le.enemy_attack(130 * (math.Pi / 180))
		le.enemy_attack(145 * (math.Pi / 180))
		le.super_time = linuxSuperAttackTime
	}
}

func (le *linuxEnemy) enemy_attack(angle float64) {
	if ebul, ok := ebulletFromPool(); ok {
		ebul.active = true
		ebul.x = le.x
		ebul.y = le.y
		ebul.angle = angle
	}
}
