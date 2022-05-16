package chapter5

import "fmt"

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