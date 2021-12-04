package main

import (
	"fmt"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var (
	sum int
)

type Coord struct {
	x int
	y int
}

func tourMain() {
	fmt.Println("Loaded Tour")
	defer fmt.Println("\nExited Tour...")

	// These infer the right value from the constant
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum = 1
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("sum is ", sum)

	sum = 0
	for sum < 1000 {
		sum += 100
	}

	fmt.Println("sum is ", sum)

	if v := 7; v == 7 {
		fmt.Println("v is ", v)
	} else {
		// can still access v here
		fmt.Println("v isn't 7, it's ", v)
	}

	// Pointer stuff
	var p *int // p is a pointer to an int
	x := 5

	fmt.Println("p and x are: ", p, x)
	fmt.Println("updating...")
	p = &x // p is now pointing to the address of x
	*p = 4 // deference pointer p, update value of x
	fmt.Println("p and x are: ", p, x)
	fmt.Println("value of dereferenced p (and x) are now:", *p)

	coord := Coord{3, 5}
	coordP := &coord
	//both of these are are valid
	fmt.Println("Coords are x: ", coord.x, " and y: ", coord.y)
	fmt.Println("Coords are x: ", (*coordP).x, " and y: ", coordP.y)

	tourArraysAndSlices()
	tourInterfaces()

}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func tourArraysAndSlices() {

	//Array syntax
	var nums [4]int
	nums = [4]int{1, 2, 3, 4}
	fmt.Println("Array is: ", nums)
	sliced := nums[1:3]
	fmt.Println("Sliced is: ", sliced)

	// Slice is not a new array!!!
	nums[1] = 7
	fmt.Println("Array changed: ", nums)
	fmt.Println("So did slice: ", sliced)

	// Easy way to iterate over array (or map)
	for i, v := range nums {
		fmt.Println("Value ", v, " at index ", i)
	}

	// Must use make to create maps
	var coordMap map[string]Coord
	coordMap = make(map[string]Coord)
	coordMap["point-one"] = Coord{1, 2}
	coordMap["point-two"] = Coord{4, 2}
	fmt.Println("Map is: ", coordMap)

	// Work with arrays
	aCoord, ok := coordMap["point-one"]
	if ok {
		fmt.Println("Key is in the map! Value: ", aCoord)
	}
	delete(coordMap, "point-one")
	aCoord, ok = coordMap["point-one"]
	if !ok {
		// returns {0,0}, not nil, because coord values are primitives
		fmt.Println("Key is not in the map... Value: ", aCoord)
	}
}

func tourInterfaces() {

	// generics... kind of
	var anything interface{} = "someString"

	s := anything.(string) // okay
	fmt.Println("Cast to string: ", s)
	f, ok := anything.(float64) // compiles but fails
	if ok {
		fmt.Println("Cast to float: ", f)
	}

	fmt.Println(trickyCast(3))
	fmt.Println(trickyCast(3.214234))
	fmt.Println(trickyCast("DEFINITELY_not_a_number"))

}

func trickyCast(any interface{}) float64 {
	switch v := any.(type) {
	case int:
		return float64(v)

	case float32:
		return float64(v)

	case float64:
		return float64(v)

	default:
		fmt.Println("Couldn't cast!! ", any)
		return 0.0
	}

}
