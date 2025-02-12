package main

// create a closure function which returns a function that generates positive integers
func generator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {
	numGen := generator()
	for i := 0; i < 5; i++ {
		println(numGen(), "\t")
	}
}
