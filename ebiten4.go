// ebiten4 project ebiten4.go
package main

import (
	"log"

	"github.com/Blam25/Test/Pkg/eccomp"
	"github.com/Blam25/Test/Pkg/ecentity"
	"github.com/Blam25/Test/Pkg/ecevent"
	"github.com/Blam25/Test/Pkg/ecsys"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&ecsys.G); err != nil {
		log.Fatal(err)
	}
}

func init() {
	initMyComponents()
	initMySystems()
	print("hej")

	player := ecentity.NewEntity()
	eccomp.NewPlayer(player)
	eccomp.NewHealth(player, 100)
	eccomp.NewPosition(player, 200, 125)
	eccomp.NewRend(player, "assets/gopher.png")

	entity1 := ecentity.NewEntity()
	eccomp.NewRend(entity1, "assets/gopher.png")
	eccomp.NewCollider(entity1)
	ecevent.NewDmgPlayer(entity1, 20, ecevent.Events.CollisionMap)
	eccomp.NewPosition(entity1, 200, -200)
	eccomp.NewMoveWithCamera(entity1)
	eccomp.NewMovePatrol(entity1)

	entity2 := ecentity.NewEntity()
	eccomp.NewRend(entity2, "assets/gopher.png")
	eccomp.NewCollider(entity2)
	ecevent.NewDmgPlayer(entity2, 40, ecevent.Events.CollisionMap)
	eccomp.NewPosition(entity2, -200, 200)
	eccomp.NewMoveWithCamera(entity2)

	entity3 := ecentity.NewEntity()
	eccomp.NewRend(entity3, "assets/gopher.png")
	eccomp.NewCollider(entity3)
	ecevent.NewDmgPlayer(entity3, 60, ecevent.Events.CollisionMap)
	eccomp.NewPosition(entity3, 100, 100)
	eccomp.NewMoveWithCamera(entity3)

}

func initMyComponents() {
	//myComponents = &Components{}
	//myComponents.TestCompMap = make(map[int]*TestComp)
}

func initMySystems() {
	//NewTestSystem()
	//E.NewSysCollide()
}
