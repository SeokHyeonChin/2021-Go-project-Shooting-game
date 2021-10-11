package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("loading : %v", err))
	}
	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}

func stage_clear(renderer *sdl.Renderer) {
	tex := textureFromBMP(renderer, "sprites/clear.bmp")
	renderer.Copy(tex,
		&sdl.Rect{X: 0, Y: 0, W: 1000, H: 570},
		&sdl.Rect{X: 50, Y: 150, W: 500, H: 270})
}

func game_over(renderer *sdl.Renderer) {
	tex := textureFromBMP(renderer, "sprites/game_over.bmp")
	renderer.Copy(tex,
		&sdl.Rect{X: 0, Y: 0, W: 1080, H: 1080},
		&sdl.Rect{X: 50, Y: 150, W: 500, H: 500})
}

func main() {
	//sdl Initiation
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL: ", err)
		return
	}

	//Window Creation
	window, err := sdl.CreateWindow(
		"Shooting Go Game - Kubernetes Addition",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window: ", err)
		return
	}
	defer window.Destroy()

	//renderer Creation
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer: ", err)
		return
	}
	defer renderer.Destroy()

	//Rendering
	plr := newPlayer(renderer)

	boss := newLinuxEnemy(renderer, 300, 90)

	initBasicEnemy(renderer)

	initBulletPool(renderer)

	initEBulletPool(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.Clear()
		renderer.SetDrawColor(222, 255, 255, 255)

		plr.draw(renderer)
		plr.update()

		boss.draw(renderer)
		boss.update()

		for _, enemy := range enemies {
			enemy.draw(renderer)
			enemy.update()
		}

		for _, bul := range bulletPool {
			bul.draw(renderer)
			bul.update()

			// Linux Enemy Collision detector
			if bul.active && math.Sqrt(math.Pow(boss.x-bul.x, 2)+
				math.Pow(boss.y-bul.y, 2)) <= (boss.radius+bul.radius) {
				boss.life -= 1
				bul.active = false
				if boss.life <= 0 {
					boss.active = false
				}
			}
		}

		for _, ebul := range ebulletPool {
			ebul.draw(renderer)
			ebul.update()
			// Player Collision detector
			if ebul.active && math.Sqrt(math.Pow(plr.x-ebul.x, 2)+
				math.Pow(plr.y-ebul.y, 2)) <= (plr.radius+ebul.radius) {
				plr.active = false
				ebul.active = false
			}
		}
		// Basic Enemy Collision detector
		for _, enemy := range enemies {
			if enemy.active {
				for _, bul := range bulletPool {
					if bul.active && math.Sqrt(math.Pow(enemy.x-bul.x, 2)+
						math.Pow(enemy.y-bul.y, 2)) <= (enemy.radius+bul.radius) {
						enemy.active = false
						bul.active = false
					}
				}
			}
		}

		// Game Over
		if !plr.active {
			game_over(renderer)
			renderer.Present()
			time.Sleep(time.Second * 5)
			break
		}
		// Game Clear -loop
		if !boss.active {
			stage_clear(renderer)
			renderer.Present()
			time.Sleep(time.Second * 5)
			break
		}
		renderer.Present()
	}
}
