package main

import E "EasyEC"

type CompHealth struct {
	entity *E.Entity
	Health int
}

func NewCompHealth(entity *E.Entity, health int) {
	new := CompHealth{}
	new.entity = entity
	new.Health = health
	myComponents.CompHealth = append(myComponents.CompHealth, &new)
	myComponents.CompHealthMap[entity.GetId()] = &new
}
