// ebiten4 project ebiten4.go
package main

import (
	E "EasyEC"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//EasyEC.Components.

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&E.G); err != nil {
		log.Fatal(err)
	}
}

// var G Game
var entityOrder int = 1
var test int = 1.0
var myComponents *Components

//var entities []*Entity
//var renderedNonPlayers []*renderedNonPlayer

//var op *ebiten.DrawImageOptions

func init() {
	initMyComponents()
	initMySystems()

	//toDataArg([]string{"hej"})
	player := E.NewEntity()
	E.NewCompPlayer(player)
	E.NewCompHealth(player, 100)
	E.NewCompPosition(player, 200, 125)
	E.NewCompRend(player, "assets/gopher.png")
	//sprite1 := E.New_Sprite()
	//sprite1.Xpos = 500

	entity1 := E.NewEntity()
	E.NewCompRend(entity1, "assets/gopher.png")
	E.NewCompCollider(entity1)
	E.NewEventDmgPlayer(entity1, 20, E.Components.EventCollisionMap)
	E.NewCompPosition(entity1, 500, 500)
	E.NewCompMoveWithCamera(entity1)
	/*entity2 := E.NewEntity().With(
		E.NewCompRend, "assets/gopher.png").With(
		E.NewCompCollider, E.NewDataArgEmpty()).With(
		E.NewCompMoveRand, &E.Data{})
	E.NewCollisionEventTest(entity2, 3)
	entity4 := E.NewEntity().With(
		E.NewCompRend, "assets/gopher.png").With(
		E.NewCompMoveRand, E.NewDataArgEmpty()).With(
		E.NewCompCollider, E.NewDataArgEmpty())
	print(&entity4)
	E.NewCollisionEventTest(entity4, 6)
	/*E.NewEntity().With(
		E.NewCompRend, E.NewDataArgFloat(400, 400)).With(
		E.NewCompCollider, E.NewDataArgEmpty())
	//print(&entity5)

	entity6 := E.NewEntity().
		With(E.NewCompRend, E.NewDataArgFloat(400, 400)).With(
		E.NewCompCollider, E.NewDataArgEmpty())
	E.NewCollisionEventTest(entity6, 9)
	print(entity6)

	//		With(NewCompRend, NewDataArgFloat(1000, 1000)).
	//						With(NewCompCollider, &dataArgs{})
	print(&entity2)
	//entity2.(NewCompRend, NewDataArgFloat(1000, 1000)).
	//With(NewCompCollider, &dataArgs{})

	entity3 := E.NewEntity()
	entity3.With(
		E.NewCompRend, E.NewDataArgFloat(1000, 0)).With(
		E.NewCompCollider, &E.Data{})

	//NewCompRenderedNonPlayer(entity2, 1000, 1000)*/
}

func initMyComponents() {
	myComponents = &Components{}
	myComponents.TestCompMap = make(map[int]*TestComp)
}

func initMySystems() {
	NewTestSystem()
	E.NewSysCollide()
}
