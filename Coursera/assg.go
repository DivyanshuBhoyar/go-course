package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name string
	props []string
}
type Snake struct {
	name string
	props []string
}
type Bird struct {
	name string
	props []string
}
var propsMap = map[string] []string{
	"Cow": {"grass", "walk" ,"moo"},
	"Snake": {"mice", "slither" ,"hss"},
	"Bird": {"worms", "fly" ,"peep"},
}
func (c Cow) Eat() {fmt.Println(c.props[0])}
func (c Cow) Move() {fmt.Println(c.props[1])}
func (c Cow) Speak() {fmt.Println(c.props[2])}
func (s Snake) Eat() {fmt.Println(s.props[0])}
func (s Snake) Move() {fmt.Println(s.props[1])}
func (s Snake) Speak() {fmt.Println(s.props[2])}
func (b Bird) Eat() {fmt.Println(b.props[0])}
func (b Bird) Move() {fmt.Println(b.props[1])}
func (b Bird) Speak() {fmt.Println(b.props[2])}


var  (
	all_cows   = map[string]Cow{}
	all_snakes = map[string]Snake{}
	all_birds  = map[string]Bird{}
)

func  newAnimal( name , animal string ) {
	switch animal {
	case "cow" :
		all_cows[name] = Cow{ name ,  propsMap["Cow"] }
	case "bird":
		all_birds[name] = Bird{ name ,  propsMap["Bird"] }
	case "snake":
		all_snakes[name] = Snake{ name ,  propsMap["Snake"] }
	}
	fmt.Println("Created it!")
}
func main() {
	var mode, name, query string
	var a Animal
	fmt.Println("Enter <mode> <name> <query>")
	fmt.Println("MODE : newanimal OR query		NAME : name		QUERY : Property OR newanimal's type")
	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s %s\n", &mode, &name, &query)
		switch mode {
		case "newanimal":
			newAnimal(name, query)
		case "query":
			switch query {
			case "eat":
				if object, ok := all_cows[name]; ok {
					a = object
					a.Eat()
				} else {
					if object, ok := all_snakes[name]; ok {
						a = object
						a.Eat()
					} else {
						if object, ok := all_birds[name]; ok {
							a = object
							a.Eat()
						}
					}
				}
			case "move":
				if object, ok := all_cows[name]; ok {
					a = object
					a.Move()
				} else {
					if object, ok := all_snakes[name]; ok {
						a = object
						a.Move()
					} else {
						if object, ok := all_birds[name]; ok {
							a = object
							a.Move()
						}
					}
				}
			case "speak":
				if object, ok := all_cows[name]; ok {
					a = object
					a.Speak()

				} else {
					if object, ok := all_snakes[name]; ok {
						a = object
						a.Speak()
					} else {
						if object, ok := all_birds[name]; ok {
							a = object
							a.Speak()
						}
					}
				}

			}
		}
	}
}