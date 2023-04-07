package main

type Components struct {
	TestComps       []*TestComp
	TestCompMap     map[int]*TestComp
	CompPosition    []*CompPosition
	CompPositionMap map[int]*CompPosition
	CompDamage      []*CompDamage
	CompDamageMap   map[int]*CompDamage
	CompHealth      []*CompHealth
	CompHealthMap   map[int]*CompHealth
}
