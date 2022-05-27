package chapter6

import (
	"container/list"
	"fmt"
	"time"
)

type Data struct {
	key, value string
}

func invertKeyValue(data Data) {
	var temp = data.key
	data.key = data.value
	data.value = temp
}

// this one will change the real variable, not just copy it
// the param 'data' has type of "*Data" / pointer to Data
func realInvertKeyValue(data *Data) {
	var temp = data.key
	data.key = data.value
	data.value = temp
}

func Pointer() {
	// POINTER
	fmt.Println("[1] POINTER")

	dataName := Data{
		key:   "name",
		value: "MRX",
	}

	fmt.Println("initial dataName: ", dataName)

	var copyOfDataName Data = dataName // it just a copy, dude. Not a pointer to
	copyOfDataName.value = "MRY"

	fmt.Println("copyOfDataName: ", copyOfDataName)
	fmt.Println("after dataName: ", dataName) // dataName remain the same because when copyOfDataName created, it just copy the value inside it, not the address/refer to original dataName

	var pointToDataName *Data = &dataName
	pointToDataName.value = "Real MRY"

	fmt.Println("after REAL dataName: ", dataName)

	println()

	var directPointer *Data = &Data{key: "age", value: "33"}
	fmt.Println("direct pointer: ", directPointer)

	println()

	// FUNCTION POINTER
	fmt.Println("[2] FUNCTION POINTER")

	job := Data{key: "job", value: "programmer"}
	fmt.Println("before inverse: ", job)

	invertKeyValue(job) // it doesnt change!
	fmt.Println("after inverse: ", job)

	realInvertKeyValue(&job) // get the "reference" from variable job. If we only use 'job' it will get the value, not the pointer
	fmt.Println("after REAL inverse: ", job)

	println()

	// LINKED LIST
	fmt.Println("[3] LINKED LIST")

	var itemList *list.List = list.New()
	itemList.PushBack("potion") // insert into the last element
	itemList.PushFront("mana")  // insert into the first element

	for i := itemList.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i)
	}
	println()
	fmt.Printf("first: %v\n", itemList.Front().Value)

	println()

	// PACKAGE TIME
	fmt.Println("[4] PACKAGE TIME")

	now := time.Now()
	fmt.Println("current time: ", now)

	var layout = "2006-01-02"
	parseTime, _ := time.Parse(layout, "1990-03-20") // WTF is this??
	fmt.Println("parsed time: ", parseTime)

	println()
}
