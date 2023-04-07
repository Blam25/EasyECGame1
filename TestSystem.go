package main

import E "EasyEC"

type TestSystem struct{}

func NewTestSystem() {
	new := &TestSystem{}
	E.Systems.First = append(E.Systems.First, new)
}

func (s *TestSystem) Execute() {
	//print("yo")
}
