package day_6

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

func Function() {
	/** Function */
	val := randomNumber(1, 10)
	fmt.Println("Random Number: ", val)

	fmt.Println("Discount: ", calculateDiscount(25, 30000))

	/** Multiple Return Values */
	total, used := getSysInfo()
	fmt.Println("Total: ", total)
	fmt.Println("Used: ", used)
}

/* func functionName(paramA , paramB type) returnType {} */
func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	random := rand.Int()%(max-min+1) + min
	return random
}

/* func functionName(paramA typeA, paramB typeB) returnType {} */
func calculateDiscount(discount int, price float64) float64 {
	return (float64(discount) / 100) * price
}

func getSysInfo() (uint64, uint64) {
	v, _ := mem.VirtualMemory()

	return v.Total, v.Used
}
