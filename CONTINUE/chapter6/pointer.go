package chapter6

import "fmt"

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
}
