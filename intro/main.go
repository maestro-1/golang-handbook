package main


import "fmt"

/*
This is an introduction to programming in go
*/


func stringLessons(name string)  {
	/**
	There are 2 ways to create variables

	Using the var keyword or using the ":=" operator

	var age = 35
	age := 35 

	are the same
	*/
	greeting := "Hello"
	fmt.Println(greeting, name)
}

func numberLesson(age int8)  {

	/*
	There are different kind of numbers
	int8, int16, int32, int64

	int8, int16, int32, int64

	float32, float64

	Each one using different size of bytes, with 64 being the most memory hungry
	uint start from 0
	int split starting halfway from negative values 

	eg
	uint8 0 to 255
	int8 -127 to 128

	int is perceived as the native style of machines, and is more widely used so just stick to it
	*/
	fmt.Println("Are you", age, "years old")
}

func stringFormattingLesson(lesson string, score float32, total uint8)  {
	fmt.Printf("this is the %s lesson and I scored %0.1f out of %d ", lesson, score, total)
}

func arraySlicesLesson()  {
	/*
	Arrays in go are immutable, so they must be copied over

	If an array has a number list in the bracket, you cannot append to it
	The size is permanently fixed
	*/ 

	firstArray := []int8{8, 9, 10}
	fmt.Println(firstArray)

	// multiple elements can be added to an array at once by placing them in the append function at once
	firstArray = append(firstArray, 24)
	fmt.Println(firstArray[2])

	// size of this array is immuatable
	secondArray := [2]string{"oghenemarho", "onothoja"}
	fmt.Println(secondArray)
}

func loopsLesson()  {
	x := 5
	
	for x < 6 {
		fmt.Println("ran code and broke out of loop")
		// this can be akin to a while loop if this break statement isn't present
		break
	}

	for x < 6 {
		fmt.Println("ran 2nd for loop")
		// this can be akin to a while loop if this doesn't increment to equal the expression above
		x++
	}

	for i:=1; i<5; i++ {
		fmt.Println("ran for loop")
	}

	secondArray := [2]string{"oghenemarho", "onothoja"}

	for i:=0; i<len(secondArray); i++ {
		fmt.Println("the value of i is:", secondArray[i])
	}

	for index, value := range secondArray {
		fmt.Printf("the postion is %v and the value is %v \n", index, value)
	}
}

func mapsLesson() {
	/*
	Maps in go are basically hashmaps
	All the keys in a hashmap must have the same type
	All the values in a single map must have the same type
	*/

	menu := map[string]float32{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"puddin": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	for key, value:= range menu {
		fmt.Println(key, value)
	}
}

func main()  {
	// list of lessons I have done in go
	
	stringLessons("Onothoja")
	numberLesson(18)
	stringFormattingLesson("string formatting", 97.8, 100)
	arraySlicesLesson()
	loopsLesson()

	// conditionals()
	// functions()

	mapsLesson()
}