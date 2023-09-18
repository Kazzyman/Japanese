package main

// **do-this**
import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
)

// The content of "prompt" (i.e., aCard.KeyR | aCard.KeyRK | aCard.KeyK) is set by the calling activity
//	                                     ^ 1           ^ 2         ^ 3,4

// Ask the user for a Hiragana response (universal for both of the Romaji-containing prompts):
// Used in exercises 1 and 2 (see GoLand's usages at right)
func semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(prompt string) (in string) { //         - -
	usersGuessOrOptionDirective := in
		// The return signature (above) creates a local var 'in'', used below in the Scan():
	fmt.Printf("%s", prompt)    // Prompt the user, in a pretty blue color: Options: '?' | '??'
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" Type the corresponding Hiragana, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
// Used in exercises 6 and 9 (see GoLand's usages at right)
func prompt_and_Scan_4_RomajiResponse_to_Any_Prompt(prompt string) (in string) { //                     - -
	usersGuessOrOptionDirective := in
		// The return signature (above) creates a local var 'in', used below in the Scan():
	fmt.Printf("%s", prompt)    // Prompt the user, in a pretty blue color: Options: '?' | '??'
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" Type the corresponding Romaji, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// Ask the user for a Romaji|Hira response to the (naked, sans romaji) Katakana or Hiragana prompt:
func Prompt_Scan_4_Romaji_or_HiraResponse(prompt string) (usersGuessOrOptionDirective string) { //      - -
	fmt.Printf("%s", prompt)    // Prompt the user, in a pretty blue color: Options: '?' | '??'
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" Type the corresponding Romaji|Hira, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n") //
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// Exercises 8 and 9 don't use one of these two, and, instead, they pick their own aCardS 
func pick_Difficult_RandomCard_Assign_aCard() { //           - -
	randIndex := rand.Intn(len(fileOfCardsMostDifficult))
	aCardD = fileOfCardsMostDifficult[randIndex] // randomly pick a 'card' from the 'deck' and store it in a global var
}
func pick_RandomCard_Assign_aCard() { //                     - -
	randIndex := rand.Intn(len(fileOfCards))
	aCard = fileOfCards[randIndex] // randomly pick a 'card' from the 'deck' and store it in a global var
}


// DIRECTIVES : --------------------------------------------------------------------------------------------
// 
func handle_doubleQuestMark_directive() { //                 - -
	// Works with Hiragana or Romaji user input in all Exercises 
	var hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom string // the value to find in locateCard.go
	//
	fmt.Printf("\n  -- Type either the Hiragana char or the Romaji that you need help with:> ")
	_, _ = fmt.Scan(&hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom)
	//
	locateCardAndDisplayHelpFieldsContainedInIt(hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom)
	fmt.Println("")
}
// Handles the set Directive 'set'
func reSet_aCard_andThereBy_reSet_thePromptString() { //     - -
	var theHiraganaOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Hiragana to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt & \"aCard\\.fields\" :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Hiragana char (instead of, accidentally, an alpha char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theHiraganaOfCardToSilentlyLocate):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric == true {
		fmt.Println("Are you in alphanumeric input mode?")
		fmt.Printf("... if so, change it to Hiragana (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in Hiragana mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)
		// May yet send an Alpha string to the next func, which will itself deal with it elegantly
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		fmt.Println("")
	}
}

func read_map_of_fineOn() { //     - -
	if len(frequencyMapOf_IsFineOnChars) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe map is empty\n")
		fmt.Printf(colorReset)
	}
	//
	for s, f := range frequencyMapOf_IsFineOnChars {
		fmt.Printf(" --- From MapOf_IsFineOn: string is:")
		fmt.Printf(colorCyan)
		fmt.Printf("%s", s)
		fmt.Printf(colorReset)
		fmt.Printf(", freq is:")
		fmt.Printf(colorRed)
		fmt.Printf("%d", f)
		fmt.Printf(colorReset)
		fmt.Printf(" ---\n")
	}
	fmt.Println("")
}

// end of DIRECTIVES -----------------------------------------------------------------------------------

/*
.
.
.
*/
// Globals:
// .
// A map v v v v v v v is used to store correct guesses, and is only written to after a correct guess
var frequencyMapOf_IsFineOnChars = make(map[string]int) // create the map, dir 'read_map_of_fineOn' reads it
var frequencyMapOf_need_workOn = make(map[string]int)

// .
// Parse the map to see if the new random card matches any members of the map
// Each time we get a new random card we check the map to see if it has been done correctly 3 or more times ...
// ... if it has, then we pick another card, and check it
// else we break
func check_it_for_fine_on() { //          - -
	for s, f := range frequencyMapOf_IsFineOnChars {
		if s == aCard.KeyR { // if it is in the map we need to check the freq
			if f >= 2 { // if the freq is 3+ we need another card
				// read_map_of_fineOn() // we show the map
				// fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
					s, aCard.KeyH, f)
				check(err2)
				pick_RandomCard_Assign_aCard() // We get that new card ...
				// fmt.Println(" ... so here is a new one ... \n")
				check_it_for_fine_on() // ... and we check THAT new card with a recursive call
				_ = fileHandleBig.Close()
			} else { // else the card had a freq less than 3, so ...
				break //  ... we exit the loop and the func -- we will use this card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
}

func check_it_for_fine_on6() { //       - -
	for s, f := range frequencyMapOf_IsFineOnChars {
		if s == aCard.KeyH || s == aCard.KeyK { // if it is in the map we need to check the freq
			if f >= 2 { // if the freq is 3+ we need another card
					// read_map_of_fineOn() // we show the map
					// fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
				// Log to a file that this action was taken **do-this**
				if s == aCard.KeyH {
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
						s, aCard.KeyH, f)
					check(err2)
					_ = fileHandleBig.Close()
				} else if s == aCard.KeyK {
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyK:%s freq:%d \n",
						s, aCard.KeyK, f)
					check(err2)
					_ = fileHandleBig.Close()
				}
				pick_RandomCard_Assign_aCard() // We get that new card ...
					// fmt.Println(" ... so here is a new one ... \n")
				check_it_for_fine_on6() // ... and we check THAT new card with a recursive call
			} else { // else the card had a freq less than 3, so ...
				break //  ... we exit the loop and the func -- we will use this card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
}
func check_it_for_fine_on7() { //       - -
	for s, f := range frequencyMapOf_IsFineOnChars {
		if s == aCard.KeyH || s == aCard.KeyK { // if it is in the map we need to check the freq
			if f >= 2 { // if the freq is 3+ we need another card
					// read_map_of_fineOn() // we show the map
					// fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
				// Log to a file that this action was taken **do-this**
				if s == aCard.KeyH {
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
						s, aCard.KeyH, f)
					check(err2)
					_ = fileHandleBig.Close()
				} else if s == aCard.KeyK {
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyK:%s freq:%d \n",
						s, aCard.KeyK, f)
					check(err2)
					_ = fileHandleBig.Close()
				}
				pick_Difficult_RandomCard_Assign_aCard() // We get that new card ...
					// fmt.Println(" ... so here is a new one ... \n")
				check_it_for_fine_on7() // ... and we check THAT new card with a recursive call
			} else { // else the card had a freq less than 3, so ...
				break //  ... we exit the loop and the func -- we will use this card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
}

func check_it_for_needing_more_practice() { //       - -
	var skip_this_step bool
	skip_this_step = false
	for s, f := range frequencyMapOf_need_workOn {
		if s == aCard.KeyR { // Check if the latest random card is in the need_workOn map, and check the freq ...
			if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we keep it
					// read_map_of_fineOn() // we show the map
				fmt.Printf("\n The Random card: %s was missed once or more \n", aCard.KeyH)
				fmt.Println("... so we will keep it and quiz you on it ... ")
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				_, err2 := fmt.Fprintf(fileHandleBig, "\nPart1of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
					s, f)
				check(err2)
				skip_this_step = true
				_ = fileHandleBig.Close()
				break //  ... we exit the loop and the func -- we will keep and use this random card, and skip the next loop
			} else { // else the card had a freq less than 2, so ...
				continue // keep looking through the map for another instance that may in there, with a significant freq
			}
		}
	}
	if skip_this_step == false {
		// The latest random was not in the map, but it is time to serve-up something difficult ... so:
		for s, f := range frequencyMapOf_need_workOn {
			if s == aCard.KeyR { // Check if the latest random is in the map, and check the freq ...
				if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we set it as aCard
					// read_map_of_fineOn() // we show the map
					fmt.Println("\n This Random card was missed 1 or more times ")
					fmt.Println("... so we will test you on it, since it has been a while")
					// Log to a file that this action was taken **do-this**
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandleBig, "\nPart2of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
						s, f)
					check(err2)
					_ = fileHandleBig.Close()
					practice_this_card(aCard.KeyR) // locate and assign aCard // set it as new aCard
					break                          //  ... we exit the loop and the func -- we will keep and use this random card
				} else { // else the card had a freq less than 2, so ...
					continue // keep looking through the map for another instance that may in there, with a significant freq
				}
			}
		}
	}
}
func check_it_for_needing_more_practiceD_7() { //      - -
	var skip_this_step bool
	skip_this_step = false
	for s, f := range frequencyMapOf_need_workOn {
		if s == aCard.KeyR { // Check if the latest random card is in the need_workOn map, and check the freq ...
			if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we keep it
					// read_map_of_fineOn() // we show the map
				fmt.Printf("\n The Random card: %s was missed once or more \n", aCard.KeyH)
				fmt.Println("... so we will keep it and quiz you on it ... ")
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				_, err2 := fmt.Fprintf(fileHandleBig, "\nPart1of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
					s, f)
				check(err2)
				skip_this_step = true
				_ = fileHandleBig.Close()
				break //  ... we exit the loop and the func -- we will keep and use this random card, and skip the next loop
			} else { // else the card had a freq less than 2, so ...
				continue // keep looking through the map for another instance that may in there, with a significant freq
			}
		}
	}
	if skip_this_step == false {
		// The latest random was not in the map, but it is time to serve-up something difficult ... so:
		for s, f := range frequencyMapOf_need_workOn {
			if s == aCard.KeyR { // Check if the latest random is in the map, and check the freq ...
				if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we set it as aCard
						// read_map_of_fineOn() // we show the map
					fmt.Println("\n This Random card was missed 1 or more times ")
					fmt.Println("... so we will test you on it, since it has been a while")
					// Log to a file that this action was taken **do-this**
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandleBig, "\nPart2of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
						s, f)
					check(err2)
					_ = fileHandleBig.Close()
					practice_this_cardD(aCard.KeyR) // locate and assign aCard // set it as new aCard
					break                           //  ... we exit the loop and the func -- we will keep and use this random card
				} else { // else the card had a freq less than 2, so ...
					continue // keep looking through the map for another instance that may in there, with a significant freq
				}
			}
		}
	}
}

/*
const seedFile = "randomSeed.dat"

var seed int64

//goland:noinspection ALL
func seedFile_Maker() {
	// Try to read existing seed
	if data, err := os.ReadFile(seedFile); err == nil {
		seed = int64(binary.LittleEndian.Uint64(data))
	} else {
		// No existing seed, create a new one
		err := binary.Read(rand.Reader, binary.LittleEndian, &seed)
		if err != nil {
			return
		}
		f, _ := os.Create(seedFile)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
			}
		}(f)
		err2 := binary.Write(f, binary.LittleEndian, seed)
		if err2 != nil {
			return
		}
	}
}

func createAndWrite_seedFile() {
	f, _ := os.Create(seedFile)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	err := binary.Write(f, binary.LittleEndian, seed)
	if err != nil {
		return
	}
}
*/
func stack_the_map() { //             - -
	promptToSkip := "shi"
	for i := 0; i < 6; i++ {
		frequencyMapOf_IsFineOnChars[promptToSkip]++
	}
	fmt.Printf("\nSix occurrences of 'shi' have been added to frequencyMapOf_IsFineOnChars\n\n")
}

func branchOnUserSelectedDirectiveIfGiven(in, selectedExercise string) { //             - -
	if in == "set" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "menu" ||
		in == "reset" ||
		in == "stat" ||
		in == "dir" ||
		in == "notes" ||
		in == "quit" ||
		in == "exit" ||
		in == "stats" ||
		in == "rm" ||
		in == "stack" {
		switch in {
		case "menu":
			// Flush the old stats and hits arrays
			cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
			cyclicArrayHits = CyclicArrayHits{}
			// usersGuessOrOptionDirective = "null"
			nonRandomCard = 0
			// Also, flush the map
			frequencyMapOf_IsFineOnChars = make(map[string]int)
			do_betweenMainMenuSelectionsTTE(selectedExercise) // This only writes transition entries to the log file
			mainMenuPromptScanSelectAndBeginSelectedExercise()
		case "reset":
			// Flush (clear) the old stats and hits arrays
			cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
			cyclicArrayHits = CyclicArrayHits{}
			// Also, flush (clear) the map
			frequencyMapOf_IsFineOnChars = make(map[string]int)
		case "quit":
			os.Exit(1)
		case "exit":
			os.Exit(1)
		case "??": // Directives follow:
			handle_doubleQuestMark_directive()
		case "?":
			switch selectedExercise {
			/*
			After the third failed guess, there is never any point in showing less than the first 3 hint lines !!!!
			So, I think that all this has been done already (though 3 has the sans Romaji for the first Kata section)

			However, in the case of '?' Directive handling, the following guidelines would be optimal
			and, the same would be true of the '??' Directive handling: the following guidelines would be optimal

			(in the first case we would need to switch on the selected exercise, but, for the second case,
			we will know what the user knows since he has already 'given' either a hira or a romaji as the seed)
			*/
			case "Romaji_Prompt": // 1 respond with Hira
					/*
						hint should be sans hira, but may include all else,
						though romaji-EX-clusion would be entirely-pointless
					*/
				fmt.Printf("\n%s\n%s\n%s\n\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
			//
			case "Romaji_w_Kata_Prompt": // 2 respond with Hira
					/*
						hint should be sans hira, but may also include all else,
						though romaji-and/or-kata-EX-clusion would be entirely-pointless
					*/
				fmt.Printf("\n%s\n%s\n%s\n\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
			//
			case "Respond_w_Hira_or_Romaji_to_kataPrompt_3": // 3 & '4'
				// fmt.Printf("\n\n%s\n", aCard.HintSansR)
				fmt.Printf("\n%s\n%s\n%s\n\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
					/*
						ideally, only a vague hint sans any mention of hira, kata, or romaji is preferred,
						though kata-EX-clusion would be entirely-pointless
					*/
			//
			case "Mixed_prompts": // 6
				// fmt.Printf("\n\n%s\n", aCard.HintSansR)
				fmt.Printf("\n%s\n%s\n%s\n\n", aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT)
					/*
					   6 prompts w a mixture of kata and hira : accepts MAINLY romaji :
						so, hint SHOULD BE SANS ROMAJI; and the pointless cases are not really worth worrying about !!
					*/
			//
			//	
			case "Most_Difficult": // 7
				fmt.Printf("\n%s\n%s\n%s\n\n", aCardD.Hint1h, aCardD.Hint2k, aCardD.Hint3TT)
					// 7 prompts w kata only : accepts either hira OR romaji,
					// so, ideally, only a vague hint sans any mention of hira, kata, or romaji is preferred,
					// though kata-EX-clusion would be entirely-pointless
			//
			case "Sequential_Kata": // 8
				fmt.Printf("\n%s\n%s\n%s\n\n", aCardS.Hint1h, aCardS.Hint2k, aCardS.Hint3TT)
				   // 8 prompts w kata only : accepts either hira OR romaji,
				   // so, ideally, only a vague hint sans any mention of hira, kata, or romaji is preferred,
				   // though kata-EX-clusion would be entirely-pointless
			//
			case "Sequential_Hira": // 9
				fmt.Printf("\n%s\n%s\n%s\n\n", aCardS.Hint1h, aCardS.Hint2k, aCardS.Hint3TT)
					// 9 prompts w hira only : and, ACTUALLY, only ACCEPTS romaji :
					// hint should be sans romaji -- hira EX-clusion being entirely-pointless
			//
			} // End of switch within a switch 

		case "set":
			reSet_aCard_andThereBy_reSet_thePromptString()
		case "stat":
			hits()
		case "stats":
			hits()
		case "notes":
					//goland:noinspection ALL  **do-this**
			fmt.Println("\nIn the traditional Hepburn romanization system, the sound じ in hiragana is romanized as \"ji\" \n" +
				"and the katakana ジ is also romanized as \"ji\" \n\n" +
				"However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as\n" +
				" \"zi\" instead of \"ji\"\n\n" +
				"The sound gi:ぎ in hiragana is romanized as \"gi\" and the katakana ギ is also romanized as \"gi\"\n")
					//goland:noinspection ALL  **do-this**
			fmt.Println("゜is called \"handakuten\" 半濁点 translates to \"half-voicing mark\" or \"semi-voiced mark\"\n" +
				"゛is called \"dakuten\" 濁点 meaning 'voiced mark' or 'voicing mark'")
		case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
			switch selectedExercise {
				case "Romaji_Prompt": // 1
					body_of_Romaji_instructions()
					fmt.Println("You are doing Exercise 1")
				case "Romaji_w_Kata_Prompt": // 2
					body_of_Romaji_plus_Kata_instructions()
					fmt.Println("You are doing Exercise 2")
				case "Respond_w_Hira_or_Romaji_to_kataPrompt_3": // 3 & '4'
					body_of_KataExerciseInstructions()
					fmt.Println("You are doing Exercise 3")
				case "Most_Difficult": // 7
					body_of_Difficult_instructions()
					fmt.Println("You are doing Exercise 7")
				case "Mixed_prompts": // 6
					body_of_instructions_for_Romaji_responces_only()
					fmt.Println("You are doing Exercise 6")
				case "Sequential_Kata": // 8
					body_of_KataExerciseInstructions()
					fmt.Println("You are doing Exercise 8")
				case "Sequential_Hira": // 9
					body_of_instructions_for_Romaji_responces_only()
					fmt.Println("You are doing Exercise 9")
			}
		case "rm":
			read_map_of_fineOn()
		case "stack":
			// Load six occurrences of 'shi' to the map_of_fineOn 
			stack_the_map()
		}
	}
}

// Creates a func named check which takes one parameter "e" of type error
func check(e error) { //      - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
