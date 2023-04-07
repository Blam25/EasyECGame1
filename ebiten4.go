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

	//sprite1 := E.New_Sprite()
	//sprite1.Xpos = 500

	entity1 := E.NewEntity()
	print(&entity1)
	entity1.
		With(E.NewCompRend, E.NewDataArg(nil, nil, E.FArr(-500, -500), nil)).
		With(E.NewCompCollider, &E.Data{Strings: E.SArr("hej", "yo"), Ints: E.IArr(9, 2)})
	E.NewCollisionEventTest(entity1, 1)
	entity2 := E.NewEntity().With(
		E.NewCompRend, E.NewDataArgFloat(100, 100)).With(
		E.NewCompCollider, E.NewDataArgEmpty()).With(
		E.NewCompMoveRand, &E.Data{})
	E.NewCollisionEventTest(entity2, 3)
	entity4 := E.NewEntity().With(
		E.NewCompRend, E.NewDataArgFloat(400, 400)).With(
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
	print(entity6)*/

	//		With(NewCompRend, NewDataArgFloat(1000, 1000)).
	//						With(NewCompCollider, &dataArgs{})
	print(&entity2)
	//entity2.(NewCompRend, NewDataArgFloat(1000, 1000)).
	//With(NewCompCollider, &dataArgs{})

	entity3 := E.NewEntity()
	entity3.With(
		E.NewCompRend, E.NewDataArgFloat(1000, 0)).With(
		E.NewCompCollider, &E.Data{})

	//NewCompRenderedNonPlayer(entity2, 1000, 1000)
}

func initMyComponents() {
	myComponents = &Components{}
	myComponents.TestCompMap = make(map[int]*TestComp)
}

func initMySystems() {
	NewTestSystem()
	E.NewSysCollide()
}
