package main

// **do-this**
import (
	"fmt"
	"strings"
)

// LOGGERS:
// 3 'Right' or 'Oops' loggers
//
// Used in RomajiKata Exorcise
func logHitsRomajiKata(RightOrOops string) {
	cyclicArrayHits.InsertRightOrOops(RightOrOops)
	cyclicArrayHits.InsertChar(aCard.KeyRK) // this works because this func is not passed the Key
}

// Used in NakedKata exorcise
func logHitsKata(RightOrOops, kataChar string) {
	cyclicArrayHits.InsertRightOrOops(RightOrOops)
	cyclicArrayHits.InsertChar(kataChar)
}

// Used in NakedRomaji exorcise
func logHitsRomaji(RightOrOops, romaji string) {
	cyclicArrayHits.InsertRightOrOops(RightOrOops)
	cyclicArrayHits.InsertChar(romaji)
}

// 3 GottenWrong loggers
//
// Used in NakedKata exorcise
func logKataGottenWrong(kataChar string) {
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(kataChar)
}

func logRomajiKataGottenWrong(hiraChar string) {
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(hiraChar)
}

// Used in NakedRomaji exorcise
func logRomajiGottenWrong(romajiChars string) {
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(romajiChars)
}

// directives:
func hits() {
	// Create a map to store the frequency of each string (as was done above in the stats directive)
	frequencyMapRightOrOops := make(map[string]int)
	frequencyMapChar := make(map[string]int) // These, apparently, creates a map to associate a unique string with an int
	frequencyMapWrongs := make(map[string]int)

	//
	// Parse the relevant cyclic array to extract the strings and put them into the relevant map
	//
	// RightOrOops map
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i]
		//
		// Apparently this loads a string into; and increments the frequency of, that particular string, in the map
		//frequencyMapChar[str]++ // ... this, apparently, increments the int mapped to a particular 'str' in said map
		frequencyMapRightOrOops[str]++
	}
	// char map
	for i := 0; i < len(cyclicArrayHits.jchar); i++ {
		str := cyclicArrayHits.jchar[i]
		//
		// Apparently this loads a string into; and increments the frequency of, that particular string, in the map
		frequencyMapChar[str]++
	}
	// load wrongs map
	for i := 0; i < len(cyclicArrayOfTheJcharsGottenWrong.jchar); i++ {
		str := cyclicArrayOfTheJcharsGottenWrong.jchar[i]
		frequencyMapWrongs[str]++
	}

	// -- PRINT -- the 'Right' and 'Oops' and their frequencies (Right or Oops) (top of printout)
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" {
			fmt.Printf(colorGreen)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorGreen)
			fmt.Printf(" %d\n", freq)
			fmt.Printf(colorReset)
		} else if str == "Oops" { // it is Oops
			fmt.Printf(colorRed)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorRed)
			fmt.Printf(" %d\n", freq)
			fmt.Printf(colorReset)
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		}
	}
	// Print the unique Chars and their frequencies (continuing printout from above)
	numberOfUniqueCharsHit := 0
	for str, freq := range frequencyMapChar {
		if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			numberOfUniqueCharsHit++
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorReset)
			fmt.Printf(" %d\n", freq)
		}
	}
	fmt.Printf("Number of unique chars: ")
	fmt.Printf(colorPurple)
	fmt.Printf("%d \n\n", numberOfUniqueCharsHit)
	fmt.Printf(colorReset)

	// Print the ones gotten wrong  (continuing printout from above)
	for str, freq := range frequencyMapWrongs {
		if str == "" {
			// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			fmt.Printf("Gotten Wrong:")
			fmt.Printf(colorRed)

			fieldsOfStr := strings.Split(str, ":")

			fmt.Printf("%s", fieldsOfStr[0]) // Kata
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[1]) // it was
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[2]) // Romaji
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[3]) // but you had guessed
			fmt.Printf(colorRed)
			fmt.Printf("%s ", fieldsOfStr[4]) // the bad guess

			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorReset)
			fmt.Printf(" %d \n", freq)
		}
	}
	fmt.Println("")
}
