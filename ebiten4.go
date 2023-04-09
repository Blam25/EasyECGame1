// ebiten4 project ebiten4.go
package main

import (
	"log"

	Comp "github.com/Blam25/Test/Pkg/Components"
	Ent "github.com/Blam25/Test/Pkg/Entities"
	Event "github.com/Blam25/Test/Pkg/Events"
	S "github.com/Blam25/Test/Pkg/Systems"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1000, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&S.G); err != nil {
		log.Fatal(err)
	}
}

func init() {
	initMyComponents()
	initMySystems()
	print("hej")

	player := Ent.NewEntity()
	Comp.NewPlayer(player)
	Comp.NewHealth(player, 100)
	Comp.NewPosition(player, 200, 125)
	Comp.NewRend(player, "assets/gopher.png")

	entity1 := Ent.NewEntity()
	Comp.NewRend(entity1, "assets/gopher.png")
	Comp.NewCollider(entity1)
	Event.NewDmgPlayer(entity1, 20, Event.Events.CollisionMap)
	Comp.NewPosition(entity1, 200, -200)
	Comp.NewMoveWithCamera(entity1)

	entity2 := Ent.NewEntity()
	Comp.NewRend(entity2, "assets/gopher.png")
	Comp.NewCollider(entity2)
	Event.NewDmgPlayer(entity2, 40, Event.Events.CollisionMap)
	Comp.NewPosition(entity2, -200, 200)
	Comp.NewMoveWithCamera(entity2)

	entity3 := Ent.NewEntity()
	Comp.NewRend(entity3, "assets/gopher.png")
	Comp.NewCollider(entity3)
	Event.NewDmgPlayer(entity3, 60, Event.Events.CollisionMap)
	Comp.NewPosition(entity3, 100, 100)
	Comp.NewMoveWithCamera(entity3)

}

func initMyComponents() {
	//myComponents = &Components{}
	//myComponents.TestCompMap = make(map[int]*TestComp)
}

func initMySystems() {
	//NewTestSystem()
	//E.NewSysCollide()
}
