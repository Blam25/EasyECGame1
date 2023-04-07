package main

import E "EasyEC"

type CompPosition struct {
	entity *E.Entity
	Xpos   int
	Ypos   int
}

func NewCompPosition(entity *E.Entity, xpos int, ypos int) {
	new := CompPosition{}
	new.entity = entity
	new.Xpos = xpos
	new.Ypos = ypos
	myComponents.CompPosition = append(myComponents.CompPosition, &new)
	myComponents.CompPositionMap[entity.GetId()] = &new
}
