package main

import E "EasyEC"

type CompDamage struct {
	entity *E.Entity
	Damage int
}

func NewCompDamage(entity *E.Entity, damage int) {
	new := CompDamage{}
	new.entity = entity
	new.Damage = damage
	myComponents.CompDamage = append(myComponents.CompDamage, &new)
	myComponents.CompDamageMap[entity.GetId()] = &new
}
