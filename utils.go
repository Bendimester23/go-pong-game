package main

type GameObject interface {
	Update()
	Render()
	Reset()
}
