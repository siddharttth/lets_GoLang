package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// generating an english greeting logic
	return "hi there!"
}

func (spanishBot) getGreeting() string {
	// generating an spanisg greeting logic
	return "hola!"
}

// ---------------EXAMPLE--------------------
//
// type user struct{
// 		name string
// }
// type bot interface {
// 		getbreeting(string, int) (string, error)
// 		getBotVersion() float64
// 		respondToUser(user) string
// }
//
// --------------------------------------------
