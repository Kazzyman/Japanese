package main

// **do-this**
import (
	"fmt"
	"os"
)

var caOfOldFriends CaOfOldFriends // the final resting place of our oldest friends

type CaOfOldFriends struct {
	jcharFriend [30]string
	index       int
}

func (ca *CaOfOldFriends) InsertOldFriend(JcharFriend string) {
	ca.jcharFriend[ca.index] = JcharFriend
	ca.index = (ca.index + 1) % len(ca.jcharFriend)
}
func loadFrequencyMapOf_IsFineOnChars_from_cyclicArrayUserIsFineOn(m map[string]int) map[string]int {
	// Parse the global cyclic array to extract the strings and put them into the map that was passed as m:
	for i := 0; i < len(cyclicArrayUserIsFineOn.skipThisChar); i++ {
		str := cyclicArrayUserIsFineOn.skipThisChar[i] // As was done above ...
		//fmt.Printf("\n loaded map at pos:%d with str:%s \n", i, str)
		//
		// Apparently this loads a string into; and increments the frequency of, that particular string, in the map
		//m[str]++ // ... this, apparently, increments the int mapped to a particular 'str' in said map
		if str == "" {
			// do nada
		} else {
			m[str]++ // Specifically, the '++' must increment the int value associated with str
		}
	}
	return m
}
func build_caOf3orMoreHitsPerChar(m map[string]int) {
	// build a ca_slice of elements from that map with frequencies of 3 or more
	// read through the map m
	for str, freq := range m {
		if freq >= 3 {
			fmt.Printf("\nfound a str with freq >= 3 , str:%s, freq:%d \n", str, freq)
			// put the str into the ca_slice
			caOfOldFriends.InsertOldFriend(str) // caOfOldFriends is global
		} else if str == "" {
			fmt.Println(" -- str was MT ---")
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the ca from which the map was made
		}
	}
}

/*
.
.
*/

func TouchTypingExorcise(selectedExorcise string) {
	// Accepts the following for ^ ^ ^ ^ :  RomajiNakedPrompt -- RomajiKataPrompt -- KataNakedPrompt
	//
	// rand.Seed(seed) is now done in main.go, at the top of main()
	// was: // rand.Seed(time.Now().UnixNano())
	// Example of seed Use per Claude: fmt.Println(rand.Int())

	var usersGuessOrOptionDirective string
	frequencyMapOf_IsFineOnChars := make(map[string]int) // create a map to associate a unique string with an int
	for {
	outHere:
		for { // pick a card loop
			pick_RandomCard_Assign_aCard() // Assigns a random card to aCard, so there is no need to pass any aCard fields below
			//	/*

			// load the map from cyclicArrayUserIsFineOn [which, is a global] // done v v v  *****************
			m := loadFrequencyMapOf_IsFineOnChars_from_cyclicArrayUserIsFineOn(frequencyMapOf_IsFineOnChars) // func created above ***********
			//for {
			// v v v read and assess the stats from that map m:
			// build a global ca_slice of elements from that map m with frequencies of 3 or more // done *************
			build_caOf3orMoreHitsPerChar(m) // builds (inserts into) caOfOldFriends
			// that  ^ ^ ^ is a global, by ^ ^ ^ ^ ^
			//
			// ...then, range over that CA (as a slice?) to see if the latest-randomly-picked card is in that slice
			for i := 0; i < len(caOfOldFriends.jcharFriend); i++ {
				str := caOfOldFriends.jcharFriend[i] // pull a char (an old friend) from the slice
				if str == aCard.KeyH {               // if the latest-randomly-picked card IS an old friend ...
					fmt.Printf("\n %s was an old friend\n", str)
					pick_RandomCard_Assign_aCard()
					// we then continue the loop back at the 'if', to see if the user is competent on the latest-randomly-picked card
				} else { // the latest-randomly-picked card is NOT in that slice, and we need look no more
					fmt.Printf("\n %s was not in caOfOldFriends as a jcharFriend ---- break ------\n", str)
					break outHere // else Exit the loop with the latest-randomly-picked card, which seems to be a good one to try as the next card
				} //end of if-else
				// since the else clause was not met, no 'break' has occurred, so we re-iterate the loop
			} // end of loop
			//} // end of loop
		} // end of loop
		// As we exit the loop, we now have a card which is not known to be an old friend
		fmt.Println("caOfOldFriends is:")
		fmt.Println(caOfOldFriends)
		fmt.Println("cyclicArrayUserIsFineOn is:")
		fmt.Println(cyclicArrayUserIsFineOn)
		//
		//		*/
		//
		// Prompt with the appropriate field from the new random card and accept user's guess:
		switch selectedExorcise {
		//
		//case of exorcise 1
		case "Romaji_Prompt":
			usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR) // Semi-universal prompt, passing R
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		//
		//case of exorcise 2
		case "Romaji+Kata_Prompt":
			usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK) // Semi-universal prompt, passing RK
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		//
		// case of exorcise 3 and 4
		case "Kata_Prompt-Respond-w-Hira|Romaji":
			usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // Kata prompt, passing K
			// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		}

	outOfForLoop: // this loop is only needed because we want to handel the case of successive directives (tricky)
		// ESPECIALLY, this labeled 'for-loop' is used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
		for {
			if usersGuessOrOptionDirective == "set" || usersGuessOrOptionDirective == "?" || // <-- if it IS a directive
				usersGuessOrOptionDirective == "??" || usersGuessOrOptionDirective == "menu" || usersGuessOrOptionDirective == "reset" ||
				usersGuessOrOptionDirective == "stat" || usersGuessOrOptionDirective == "dir" || usersGuessOrOptionDirective == "notes" ||
				usersGuessOrOptionDirective == "quit" || usersGuessOrOptionDirective == "exit" || usersGuessOrOptionDirective == "stats" {
				// v v v v v   HANDEL A DIRECTIVE  v v v v v v v v
				branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise) // <-- handle the directive
				// v v v v v   re-prompt   v v v v v v
				// -- re-prompt following the execution of a directive -------------------------------------------------------- v v v v
				// ****************************************************************************************************
				switch selectedExorcise { // Identical to the above switch for prompt selection ***********************
				case "Romaji_Prompt": // 1
					usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "Romaji+Kata_Prompt": // 2
					usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "Kata_Prompt-Respond-w-Hira|Romaji": // 3 or 4
					usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				}
				// v v v v v   do not process directives from the re-prompting   v v v v v
				if usersGuessOrOptionDirective != "set" &&
					usersGuessOrOptionDirective != "?" &&
					usersGuessOrOptionDirective != "??" &&
					usersGuessOrOptionDirective != "menu" &&
					usersGuessOrOptionDirective != "reset" &&
					usersGuessOrOptionDirective != "stat" &&
					usersGuessOrOptionDirective != "stats" &&
					usersGuessOrOptionDirective != "quit" &&
					usersGuessOrOptionDirective != "notes" &&
					usersGuessOrOptionDirective != "exit" &&
					usersGuessOrOptionDirective != "dir" {
					//  v ^ v ^ at this point we know that the usersGuessOrOptionDirective is probably a valid guess,
					// ... AND, we need to leave the loop after processing it
					switch selectedExorcise {
					case "Romaji_Prompt": // 1
						meatOfRomajiNakedExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
					case "Romaji+Kata_Prompt": // 2
						meatOfRomajiKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
					case "Kata_Prompt-Respond-w-Hira|Romaji": // 3 and 4
						meatOfNakedKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
					}
					break outOfForLoop
				} else { // It must be a successive directive, so we continue to process it from the top of the loop
					continue outOfForLoop // <-- Immediately reiterate using a labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "Romaji_Prompt":
					meatOfRomajiNakedExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
				case "Romaji+Kata_Prompt":
					meatOfRomajiKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
				case "Kata_Prompt-Respond-w-Hira|Romaji":
					meatOfNakedKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // end of loop used to handel successive directives without double processing of guesses
	} // end of primary loop
}

func branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise string) {
	switch usersGuessOrOptionDirective {
	case "menu":
		// Flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		usersGuessOrOptionDirective = "null"
		betweenMainMenuSelectionsTTE(selectedExorcise)
		mainMenuPromptScanSelectAndBeginSelectedExorcise()
	case "reset":
		// flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
	case "quit":
		os.Exit(1)
	case "exit":
		os.Exit(1)
	case "??": // Directives follow:
		handle_doubleQuestMark_directive()
	case "?":
		handle_singleQuestMark_contextSensitive_directive()
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
		switch selectedExorcise {
		case "Romaji_Prompt": // 1
			reDisplay_Romaji_instructions()
		case "Romaji+Kata_Prompt": // 2
			reDisplay_Romaji_plus_Kata_instructions()
		case "Kata_Prompt-Respond-w-Hira|Romaji": // 3 & 4
			re_display_KataExorciseInstructions()
		}
	}
}
