package main

import (
	"fmt"
)

type apple struct {
	sweetness int
	color     string
}

func (a apple) describeMe() string {
	return fmt.Sprint("I am a ", a.color, " apple. I have ", a.sweetness, " sweetness. [0-5]")
}

// This is a composed apple
type resizedApple struct {
	apple
	size int
}

// This is basically an overridden method
func (ra resizedApple) describeMe() string {
	return fmt.Sprint("I am a ", ra.size, " sized ", ra.color, " apple. I have ", ra.sweetness, " sweetness. [0-5]")
}

// All fruit can be described...
// bam polymorphism
// I like this natural inheritance. Pretty cool
// It's cool - the original type itself does not need to define the interface.
// The interface itself is automatically assumed based on the object it aligns to
type fruit interface {
	describeMe() string
}

func talkAboutAFruit(f fruit) {
	fmt.Println(f.describeMe())
}

func udemyReviewMain() {
	fmt.Println("Loaded Udemy")
	// Composite Literal
	arr := []int{1, 2, 3, 5, 7}
	fmt.Println("This is a slice! -> ", arr)

	m := map[string]int{
		"asdf": 1337,
		"what": 690,
	}
	fmt.Println("This is a map! -> ", m)

	a := apple{3, "red"}
	fmt.Println("Apple is: ", a)
	talkAboutAFruit(a)

	small := resizedApple{a, 2}
	large := resizedApple{a, 7}

	talkAboutAFruit(small)
	talkAboutAFruit(large)

	aNewApple := apple{4, "dark red"}
	talkAboutAFruit(aNewApple)
}
