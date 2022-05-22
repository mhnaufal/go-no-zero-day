package chapter5

import (
	"fmt"
	"strings"
)

func Funtion() {
	// FUNCTION
	fmt.Println("FUNCTION")

	isAttackSuccess := doAttack("Pikachu", 33)
	if !isAttackSuccess {
		fmt.Println("is attack succced? ", isAttackSuccess)
	} else {
		fmt.Println("is attack succced? ", isAttackSuccess)
	}

	println()

	effect, enhanced := castSpell("Fire No Jutsu", 23, 300)

	if effect == "" || enhanced == 0 {
		fmt.Println("is spell casted? No")
	} else {
		fmt.Println("is spell casted? Yes")
	}

	println()

	fmt.Println("VARIADIC FUNCTION")
	calculateAfterDamage("Moamoa", 13, 14, 31)

	println()

	evolve("Bruddog", isCanEvolve)

	println()

	fmt.Println("ANONYMOUS FUNCTION")

	potion := "Dark Magic"
	canIBuyThisPotion := func(name string, price float32) bool {
		if strings.ToLower(name) != "dark magic" {
			fmt.Printf("You buy \"%v\" for [%v]\n", name, price)
			return true
		}
		fmt.Println("FORBIDDEN potion can't be bought!")
		return false
	}

	println("-> ", canIBuyThisPotion(potion, 32))

	println()

	fmt.Println("DEFER PANIC RECOVER")

	isCanEvolve("Pee")

	fmt.Println("+------------------+")
}

func doAttack(targetMonster string, damageDealt float32) bool {
	fmt.Printf("Attack the monster [%v] and dealt [%v] damage points\n", targetMonster, damageDealt)

	if damageDealt <= 0 {
		fmt.Println("No damage dealt")
		return false
	}

	return true
}

// if named return value not updated inside the function, it will get zero value automatically
func castSpell(spell string, manaCost int, currentMana int) (effect string, enhancedEffect int) {
	fmt.Printf("Cast spell [%v] with mana cost of [%v] and left out with [%v] mana points\n", spell, manaCost, currentMana-manaCost)

	if manaCost > currentMana {
		fmt.Printf("Can't cast spell [%v]. Lack of mana!\n", spell)
		return
	}

	return "Gaining effect...", (manaCost + currentMana) / 10
}

func calculateAfterDamage(targetMonster string, damages ...int) {
	var totalDamageTaken = 0

	for _, damage := range damages {
		totalDamageTaken += damage
	}

	fmt.Printf("The monster [%v] takes a total of [%v] damage points\n", targetMonster, totalDamageTaken)
}

type Evolve func(monster string) bool

func isCanEvolve(monster string) bool {
	defer logging() // pay attention to what function get defer here

	if monster == "" || len(monster) <= 3 {
		panic("can't evolve an empty monster") // panic will get back up until it meet with defer, then it will go into the defer and see if there is a recover()
		return false
	} else {
		return true
	}
}

func evolve(monster string, isCanEvolve Evolve) {
	if isTheMonsterCanEvolve := isCanEvolve(monster); isTheMonsterCanEvolve {
		fmt.Printf("The monster [%v] evolves to \"Super %v\"\n", monster, monster)
	} else {
		fmt.Printf("Can't evolve the monster!\n")
	}
}

func logging() {
	message := recover() // get called by the panic by popping out the stack

	if message != nil {
		fmt.Println("ERROR: ", message)
	}

	fmt.Println("[#] executing program...")
}
