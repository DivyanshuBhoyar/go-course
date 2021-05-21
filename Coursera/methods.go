package main

import "fmt"

type Animal struct {
	food   string
	motion string
	sound  string
}

func (animal Animal) Eat()   { fmt.Println(animal.food) }
func (animal Animal) Move()  { fmt.Println(animal.motion) }
func (animal Animal) Speak() { fmt.Println(animal.sound) }

func main() {
	var typeQuery, infoQuery string
	 Cow := Animal{"grass", "walk", "moo"}
	 Bird:= Animal{"worms" , "fly" ,	"peep"}
	 Snake:= Animal{"mice", "slither", "hsss"}

	fmt.Println("Available animals : Cow , Bird , Snake")
	fmt.Println("Enter the type and query seperated by single space")
	for {
  fmt.Scanf("%s %s", &typeQuery, &infoQuery)

  switch typeQuery {
	
  case "Cow":
		switch infoQuery {
		case "food":
			Cow.Eat()
		case "motion":
			Cow.Move()
		case "sound":
			Cow.Speak()
		}

	case "Bird":
		switch infoQuery {
		case "food":
			Bird.Eat()
		case "motion":
			Bird.Move()
		case "sound":
			Bird.Speak()
		}

	case "Snake":
		switch infoQuery {
		case "food":
			Snake.Eat()
		case "motion":
			Snake.Move()
		case "sound":
			Snake.Speak()
		}
	}
}
}