package main

func main() {
	s := NewServer()
	s.Listen(":8000")
}