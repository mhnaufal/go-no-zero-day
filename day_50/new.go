package day_50

const greeting = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World!"
	}
	return greeting + name
}
