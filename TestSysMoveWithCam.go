package main

import (
	E "EasyEC"

	"github.com/hajimehoshi/ebiten/v2"
)

type SysMoveWithCam struct {
}

func NewSysMoveWithCam() {
	new := &SysMoveWithCam{}
	E.Systems.First = append(E.Systems.First, new)
}

func (s *SysMoveWithCam) Execute() {
	deltax, deltay := s.move()
	for _, z := range E.Components.RendNPO {
		z.Ypos = z.Ypos + deltay
		z.Xpos = z.Xpos + deltax
	}
}

func (s *SysMoveWithCam) move() (float64, float64) {
	var x float64
	var y float64
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		//print("hej")
		//s.ypos = s.ypos - s.moveSpeed
		y = E.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		//s.ypos = s.ypos + s.moveSpeed
		y = -E.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		//s.xpos = s.xpos + s.moveSpeed
		x = -E.Components.Player.MoveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		//s.xpos = s.xpos - s.moveSpeed
		x = E.Components.Player.MoveSpeed
	}
	return x, y
}
