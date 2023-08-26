package main

// **do-this**
import (
	"fmt"
	"strings"
)

// LOGGERS:
//
// 2 'Reinforce-or-Skip' loggers|Inserters:
func logSkipThisPrompt(promptToSkip string) {
	cyclicArrayUserIsFineOn.InsertMatchedPrompt(promptToSkip)
}
func logReinforceThisPrompt(promptToReinforce string) {
	cyclicArrayUserNeedsWorkOn.InsertMissedPrompt(promptToReinforce)
}

// 3 'Right' or 'Oops' loggers|Inserters:
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

// 3 GottenWrong loggers|Inserters:
//
// Used in NakedKata exorcise
func logKataGottenWrong(kataChar string) {
	// ***am-doing*** save into a file, or some combination of file and cyclic array, so as to drill more on missed
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(kataChar)
}

// Used in RomajiKata exorcise
func logRomajiKataGottenWrong(hiraChar string) {
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(hiraChar)
}

// Used in NakedRomaji exorcise
func logRomajiGottenWrong(romajiChars string) {
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(romajiChars)
}

// directives:
func hits() {
	// Create a map to store the frequency of each string
	frequencyMapRightOrOops := make(map[string]int)
	frequencyMapChar := make(map[string]int) // These, apparently, create a map to associate a unique string with an int
	frequencyMapWrongs := make(map[string]int)

	//
	// Parse the relevant cyclic array to extract the strings and put them into the relevant map:
	//
	// load RightOrOops map
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i]
		//
		// Apparently this loads a string into; and increments the frequency of, that particular string, in the map
		//frequencyMapChar[str]++ // ... this, apparently, increments the int mapped to a particular 'str' in said map
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// load char map
	for i := 0; i < len(cyclicArrayHits.jchar); i++ {
		str := cyclicArrayHits.jchar[i]
		//
		// Apparently this loads a string into; and increments the frequency of, that particular string, in the map
		frequencyMapChar[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// load wrongs map
	for i := 0; i < len(cyclicArrayOfTheJcharsGottenWrong.jchar); i++ {
		str := cyclicArrayOfTheJcharsGottenWrong.jchar[i] // As was done above ...
		frequencyMapWrongs[str]++                         // Specifically, the '++' must increment the int value associated with str
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
			// Print "Gotten Wrong:" + 'str' as multicolored (below)
			fmt.Printf("Gotten Wrong:")            // (in color White)
			fieldsOfStr := strings.Split(str, ":") // Print 'str' as multicolored (below)
			//                              // Gotten Wrong: (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[0]) // KataOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[1]) // it was (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[2]) // RomajiOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[3]) // but you had guessed:_ (in color White) _ is a space char
			fmt.Printf(colorRed)
			fmt.Printf("%s ", fieldsOfStr[4]) // the bad guess_ (in color Red) _ is a space char
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:") // Frequency: (in color Cyan)
			fmt.Printf(colorReset)
			fmt.Printf(" %d \n", freq) // 'a number' (in color White)
		}
	}
	fmt.Println("")
}
