package main

import (
	"fmt"
)

// Constants - declared at package level
const (
	PI           = 3.14159
	MaxUsers     = 1000
	AppName      = "MyApp"
	IsProduction = false
)

// Type constants
const (
	Monday    = iota // 0
	Tuesday          // 1
	Wednesday        // 2
)

func basics() {
	// ========== VARIABLE DECLARATIONS ==========

	// Method 1: var keyword with type
	var name string = "John"
	var age int = 30

	// Method 2: var keyword with type inference
	var city = "New York" // Go infers string type

	// Method 3: Short declaration (most common)
	country := "USA"

	// Method 4: Multiple variable declaration
	/* var (
	    firstName string = "Jane"
	    lastName  string = "Doe"
	    userID    int    = 12345
	)*/

	// Method 5: Multiple assignment
	x, y := 10, 20
	a, b, c := 1, 2.5, "hello"

	// ========== ZERO VALUES ==========
	var defaultInt int       // 0
	var defaultString string // ""
	var defaultBool bool     // false
	var defaultFloat float64 // 0.0

	// ========== DIFFERENT TYPES ==========

	// Integers
	//  var int8Val int8 = 127
	//  var int16Val int16 = 32767
	//  var int32Val int32 = 2147483647
	//  var int64Val int64 = 9223372036854775807
	//  var uintVal uint = 42

	// Floating point
	//  var float32Val float32 = 3.14
	//  var float64Val float64 = 3.141592653589793

	// Complex numbers
	//  var complexVal complex64 = 1 + 2i

	// Strings and runes
	// var str string = "Hello, 世界"
	// var runeVal rune = '世'  // rune is alias for int32
	//var byteVal byte = 'A'   // byte is alias for uint8

	// Arrays (fixed size)
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	//colors := [3]string{"red", "green", "blue"}

	// Slices (dynamic arrays)
	var fruits []string = []string{"apple", "banana", "orange"}
	//scores := []int{95, 87, 92}

	// Maps
	var userAges map[string]int = make(map[string]int)
	userAges["Alice"] = 25
	userAges["Bob"] = 30

	// Or initialize directly
	grades := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}

	// Pointers
	/*var ptr *int = &age
	  value := 100
	  valuePtr := &value*/

	// Structs
	type Person struct {
		Name string
		Age  int
		City string
	}

	var person1 Person = Person{
		Name: "John",
		Age:  30,
		City: "NYC",
	}

	person2 := Person{"Jane", 25, "LA"}

	// Channels
	/*var ch chan int = make(chan int)
	  messageCh := make(chan string, 5) // buffered channel

	  // Interfaces
	  var anything interface{} = "can hold any type"
	  anything = 42
	  anything = []int{1, 2, 3}*/

	// Functions as variables
	var add func(int, int) int = func(a, b int) int {
		return a + b
	}

	multiply := func(x, y int) int {
		return x * y
	}

	// ========== USING CONSTANTS ==========
	radius := 5.0
	area := PI * radius * radius
	fmt.Println("=== Constants ===")

	fmt.Printf("App: %s, Max Users: %d\n", AppName, MaxUsers)
	fmt.Printf("Today is day %d (Monday=0)\n", Tuesday)

	// ========== PRINTING VALUES ==========
	fmt.Println("=== Variables ===")
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("City: %s, Country: %s\n", city, country)
	fmt.Printf("Coordinates: (%d, %d)\n", x, y)
	fmt.Printf("Mixed: %d, %.1f, %s\n", a, b, c)

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("Int: %d, String: '%s', Bool: %t, Float: %.1f\n",
		defaultInt, defaultString, defaultBool, defaultFloat)

	fmt.Println("\n=== Collections ===")
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Printf("Grades: %v\n", grades)

	fmt.Println("\n=== Structs ===")
	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 2: %+v\n", person2)

	fmt.Println("\n=== Functions ===")
	fmt.Printf("5 + 3 = %d\n", add(5, 3))
	fmt.Printf("4 * 7 = %d\n", multiply(4, 7))

	fmt.Printf("\nCircle area: %.2f\n", area)

	// ========== VARIABLE REASSIGNMENT ==========
	name = "Updated Name" // Can reassign variables
	age = 31

	// Constants cannot be reassigned (this would cause compile error):
	// PI = 3.14  // Error: cannot assign to PI

	fmt.Printf("\nUpdated - Name: %s, Age: %d\n", name, age)
}
