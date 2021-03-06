package chapter3

import (
	"errors"
	"fmt"
)

func Composite() {
	// COMPOSITE TYPE
	fmt.Println("COMPOSITE TYPE")

	// ARRAY
	fmt.Println("[1] ARRAY")
	var arr_1 [3]int // zero value array
	fmt.Println("array: ", arr_1)

	var arr_2 = [3]int{10, 10, 10}
	fmt.Println("array: ", arr_2)

	var arr_3 = [...]string{"Joker", "Deathshoot", "The Riddle"} // although we use "..." to define the length, the array size is still fixed based on how many value given
	fmt.Println("array: ", arr_3)

	arr_3[2] = "Captain Cold"
	fmt.Println("array at index 2: ", arr_3[2])

	println()

	// COMPOSITE
	var months = [...]string{
		"january",
		"february",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"december",
	}

	var slice1 = months[4:7]
	fmt.Printf("slice: %v | Length: %v | capacity: %v\n", slice1, len(slice1), cap(slice1)) // the capacity will result 8 get from the rest of the data from the index 4 of the array months

	println()

	// Append if capacity still enough
	var slice2 = months[1:3] // [february, march]
	fmt.Println("enough: ", slice2)

	slice2[0] = "FEBRUARY"
	fmt.Println("enough: ", slice2) // february at index 0 change into FEBRUARY
	fmt.Println("enough: ", months) // february at index 1 also change into FEBRUARY

	slice3 := append(slice2, "Safar") // Safar will be placed at the very end
	fmt.Println("enough: ", slice3)   // will append at index 2 after march)
	fmt.Println("enough: ", months)   // will also be changed because it still has enough capacity

	println()

	// Append if capacity exced will create a new array
	slice4 := months[10:]
	fmt.Println("exced: ", slice4)

	var slice5 = append(slice4, "Rajab")
	fmt.Println("exced: ", slice5)
	fmt.Println("exced: ", months) // doesnt change because slice5 has already exced the capacity so if we try to append on it, it will then create a new array with "Rajab" as a new value inside it

	println()

	slice6 := make([]string, 2, 3)
	slice6[0] = "Batman"
	slice6[1] = "Superman"

	fmt.Printf("slice: %v | Length: %v | capacity: %v\n", slice6, len(slice6), cap(slice6))

	println()

	// MAP
	fmt.Println("[2] MAP")

	villain := map[string]string{
		"Joker":        "Live Laugh Laugh",
		"Captain Cold": "Freeze",
	}

	fmt.Println("map: ", villain) // when printing like this, it will get sorted by key, so that Cpatain Cold comes first
	fmt.Println("map: ", villain["Joker"])

	villain["Captain Cold"] = "Feel cold feeling you"
	fmt.Println("map: ", villain["Captain Cold"])

	delete(villain, "Captain Cold")
	fmt.Println("map: ", villain["Captain Cold"])

	villain["The Riddle"] = "Tik tok, time taking..."

	// comma ok idiom
	v, ok := villain["The Riddle"]
	if ok != false {
		fmt.Println("map: ", v)
	}

	println()

	// STRUCT
	fmt.Println("[3] STRUCT")

	var pikachu Pokemon = Pokemon{"Pikachu", 13}

	fmt.Println("Pokemon struct: ", pikachu)
	fmt.Printf("Pokemon struct: Name: [%v]\n", pikachu.name)

	println()

	pikachu.getInfo()

	println()

	// INTERFACE
	fmt.Println("[4] INTERFACE")

	fmt.Printf("My \"%v\" is on level \"%v\"\n", pikachu.GetName(), pikachu.GetLevel())
	fmt.Printf("empty interface: %v\n", Any())

	println()

	// ERROR INTERFACE
	fmt.Println("[4.5] ERROR INTERFACE")

	var pokemonError error = errors.New("This Pokemon still error!")
	fmt.Println("error interface: ", pokemonError.Error()) // get the string text of from using .Error()

	println()

	// TYPE ASSERTION
	fmt.Println("[5] TYPE ASSERTION")

	var pokePower interface{} = 303 // must be type of interface{}
	fmt.Println("new poke power in int: ", pokePower.(int))

	println()

	fmt.Println("+------------------+")

}

// Struct
type Pokemon struct {
	name  string
	level int
}

// Struct Method
func (pokemon Pokemon) getInfo() {
	fmt.Printf("INFO --> [%v]-[%v]\n", pokemon.name, pokemon.level)
}

// Interface
type BaseMonster interface {
	GetLevel() int
	GetName() string
}

// Implement Interface Method
// see there is no explicit declarations of the interface.
// Go implicitly implement an interface, if the method has 'exactly' the same signature with the method in the interface
func (pokemon Pokemon) GetLevel() int {
	return pokemon.level
}

func (pokemon Pokemon) GetName() string {
	return pokemon.name
}

// Empty Interface
func Any() interface{} {
	return nil
}
