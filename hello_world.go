// classic hello world

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	pl("Hello World!")

	// types of data
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.1415))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello, Go!"))
	pl(reflect.TypeOf('🐵'))

	// user input
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n') // handle error
	if err == nil {
		pl("Hello,", name, "Nice to meet you!")
	} else {
		log.Fatal(err)
	}

	// convert string to int
	cv3 := "3"
	cv4, err := strconv.Atoi(cv3)
	pl(cv4, err, reflect.TypeOf(cv4))

	// convert int to string
	cv5 := 5
	cv6 := strconv.Itoa(cv5)
	pl(cv6, err, reflect.TypeOf(cv6))

	// convert string to float
	cv7 := "3.1415"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
		pl(cv8, err, reflect.TypeOf(cv8))
	} else {
		log.Fatal(err)
	}

	// convert float to string
	cv9 := fmt.Sprintf("%f", 3.1415)
	pl(cv9, reflect.TypeOf(cv9))

	// if statement

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		pl("important birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("important birthday")
	} else if iAge >= 65 {
		pl("important birthday")
	} else {
		pl("not an important birthday")
	}
	pl("!true = ", !true)

	// strings

	sv1 := "This is a string"
	pl(sv1)
	replacer := strings.NewReplacer("is", "was")
	sv2 := replacer.Replace(sv1)
	pl(sv2)
	// strings operations
	pl("length of string is :", len(sv2))
	pl("contains string :", strings.Contains(sv2, "was"))
	pl("index of string :", strings.Index(sv2, "was"))
	pl("replace string :", strings.Replace(sv2, "i", "1", -1))

	sv3 := "\n String with new line\n"
	sv3 = strings.TrimSpace(sv3)                                // white spaces
	pl("split string :", strings.Split("a-b-c-d", "-"))         // split string
	pl("upper string :", strings.ToUpper(sv3))                  // upper string
	pl("lower string :", strings.ToLower(sv3))                  // lower string
	pl("prefix string :", strings.HasPrefix("tacocat", "taco")) // prefix string
	pl("suffix string :", strings.HasSuffix("tacocat", "cat"))  // suffix string
	pl(sv3)

	// runes
	rstr := "abcdfg"
	pl("Rune Count:", utf8.RuneCountInString(rstr))
	for i, runeVal := range rstr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	// working with time information
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// math operations

	pl("1 + 1 =", 1+1) // addition
	pl("1 - 1 =", 1-1) // subtraction
	pl("1 * 1 =", 1*1) // multiplication
	pl("1 / 1 =", 1/1) // integer division
	pl("1 % 1 =", 1%1) // remainder

	mInt := 1
	mInt++
	mInt += 1
	// precision
	pl("float precision :", 3.14159265359)
	// random numbers
	seedSecs := time.Now().Unix() // get current time in seconds
	rand.Seed(seedSecs)           // seed the random number generator
	randNum := rand.Intn(50) + 1  // generate random number between 1 and 50
	pl("random number :", randNum)

	pl("Abs(-10)) :", math.Abs(-10))       // absolute value
	pl("Ceil(9.25)) :", math.Ceil(9.25))   // round up
	pl("Floor(9.25)) :", math.Floor(9.25)) // round down
	pl("Mod(10,3)) :", math.Mod(10, 3))    // modulus
	pl("Pow(5,3)) :", math.Pow(5, 3))      // power
	pl("Sqrt(25)) :", math.Sqrt(25))       // square root
	pl("Cbrt(27)) :", math.Cbrt(27))       // cube root
	pl("Max(10,3)) :", math.Max(10, 3))    // maximum
	pl("Min(10,3)) :", math.Min(10, 3))    // minimum
	pl("Round(9.25)) :", math.Round(9.25)) // round to nearest integer
	pl("Round(9.75)) :", math.Round(9.75)) // round to nearest integer
	pl("Trunc(9.75)) :", math.Trunc(9.75)) // truncate decimal
	pl("Trunc(9.25)) :", math.Trunc(9.25)) // truncate decimal
	pl("Log(10)) :", math.Log(10))         // natural log
	pl("Log10(10)) :", math.Log10(10))     // log base 10
	pl("Log2(10)) :", math.Log2(10))       // log base 2

	//convert 90 degrees to radians
	r90 := (90.0 * math.Pi) / 180.0
	d90 := (r90 * 180.0) / math.Pi
	fmt.Printf("90 degrees in radians is %.2f\n", r90) // 1.570796
	fmt.Printf("90 radians in degrees is %.2f\n", d90) // 90.000000

	pl("Sin(r90)) :", math.Sin(r90)) // sine

	//format printing
	// %d = decimal integer
	// %f = float
	// %t = boolean
	// %s = string
	// %v = any value
	// %T = type of value
	// %b = binary
	// %c = character
	// %o = octal base 8
	// %x = hexadecimal base 16

	fmt.Printf("%s %d %c %f %t %o %x\n", "string", 10, 65, 3.1415, true, 8, 15)

	fmt.Printf("%9f\n", 3.14159265359)   // width of 9
	fmt.Printf("%.2f\n", 3.14159265359)  // 2 decimal places
	fmt.Printf("%9.2f\n", 3.14159265359) // width of 9 and 2 decimal places
	fmt.Printf("%9.f\n", 3.14159265359)  // width of 9 and 0 decimal places

	sp1 := fmt.Sprintf("%9.f\n", 3.14159265359)
	pl(sp1)

	// for loops
	for i := 0; i < 5; i++ {
		pl(i)
	}

	// while loops
	j := 0
	for j < 5 {
		pl(j)
		j++
	}

	// infinite loop
	k := 0
	for {
		pl(k)
		k++
		if k > 4 {
			break
		}
	}

	// guessing game with infinite loop
	seedSecs = time.Now().Unix() // get current time in seconds
	rand.Seed(seedSecs)          // seed the random number generator
	randNum = rand.Intn(50) + 1  // generate random number between 1 and 10
	pl("random number :", randNum)
	for true {
		fmt.Print("Guess the number between 1 and 50 : ")
		pl("Random number is :", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n') // handle error
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if guessInt > randNum {
			pl("Guess lower")
		} else if guessInt < randNum {
			pl("Guess higher")
		} else {
			pl("You guessed it!")
			break
		}
	}

	// arrays static
	aNums := []int{1, 2, 3}
	for _, i := range aNums {
		pl(i)
	}
	var arr1 [5]int
	arr1[0] = 10
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("index 0 of arr1 is :", arr2[0])
	pl("length of arr1 is :", len(arr2))

	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}
	arr3 := [2][2]int{{1, 2}, {3, 4}} // 2d array

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	// strings

	aStr1 := "abcdefgh"
	rArr := []rune(aStr1)
	for _, i := range rArr {
		fmt.Printf("Rune Array : %d\n", i)
	}
	byteArr := []byte("abcd")
	bStr := string(byteArr[:])
	pl("i'm a string :", bStr)

	// slices are dynamic arrays
	slice1 := make([]string, 6)
	slice1[0] = "Society"
	slice1[1] = "is"
	slice1[2] = "a"
	slice1[3] = "joint-stock"
	slice1[4] = "company"
	pl("Slice size :", len(slice1))    // size of slice
	for i := 0; i < len(slice1); i++ { // iterate over slice
		pl(slice1[i])
	}

	for _, i := range slice1 { // iterate over slice
		pl(i)
	}

	sArr := [5]int{1, 2, 3, 4, 5} // array
	slice3 := sArr[0:2]           // slice
	pl("1st 3 elements of array :", sArr[:3])
	pl("last 3 elements of array :", sArr[2:])

	sArr[0] = 10 // change array
	pl("slice3 :", slice3)
	slice3[0] = 1
	pl("sArr :", sArr)

	slice3 = append(slice3, 6, 7, 8) // append to slice
	pl("slice3 :", slice3)
	pl("sArr :", sArr)

	slice4 := make([]string, 6) // make slice
	pl("slice4 :", slice4)
	pl("slice4[0] :", slice4[0])

}
