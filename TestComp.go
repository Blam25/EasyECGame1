package main

import E "EasyEC"

type TestComp struct {
	entity *E.Entity
}

func NewTestComp(entity *E.Entity) {
	new := TestComp{}
	new.entity = entity
	myComponents.TestComps = append(myComponents.TestComps, &new)
	myComponents.TestCompMap[entity.GetId()] = &new
}
