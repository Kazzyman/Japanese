package main

// **do-this**
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// TouchTypingExorcise
// The content of "prompt" (i.e., aCard.KeyR | aCard.KeyRK | aCard.KeyK) is set by the calling activity
//
func TouchTypingExorcise(selectedExorcise string) {
	// Accepts the following for ^ ^ ^ ^ :  RomajiNakedPrompt -- RomajiKataPrompt -- KataNakedPrompt
	//
	// rand.Seed(seed) is|was|use2be done in main.go, at the top of main()
	// was: // see: seedFile_Maker() , bottom of menu.go file
	//
	rand.Seed(time.Now().UnixNano()) // old way, works with rand.Seed(seed) top of main() ???
	// Example of seed Use per Claude: fmt.Println(rand.Int())

	var usersGuessOrOptionDirective string
	//
	isThis_a_cardWeNeedMoreWorkOn := 0 // now and then we will consider working on a char the user has had trouble with
	for {
		pick_RandomCard_Assign_aCard()         // Assigns a random card to the global aCard
		if isThis_a_cardWeNeedMoreWorkOn > 2 { // if we have gone without augmenting the random picks with cards we have prev missed ...
			//
			// Log to a file that this action was taken **do-this**
			fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
			check(err)                                                                                 // ... gets a file handle to JapLog.txt
			//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
			_, err2 := fmt.Fprintf(fileHandleBig, "\nChecked card:%s in the frequencyMapOf_need_workOn after:%d cycles\n",
				aCard.KeyH, isThis_a_cardWeNeedMoreWorkOn)
			check(err2)
			isThis_a_cardWeNeedMoreWorkOn = 0 // ... for a while now, then let's go get a card we've missed before, instead of that random one
			check_it_for_needing_more_practice()
		}
		// in any case:
		check_it_for_fine_on()
		isThis_a_cardWeNeedMoreWorkOn++

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
				usersGuessOrOptionDirective == "quit" || usersGuessOrOptionDirective == "exit" || usersGuessOrOptionDirective == "stats" ||
				usersGuessOrOptionDirective == "rm" || usersGuessOrOptionDirective == "stack" {
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
					usersGuessOrOptionDirective != "rm" &&
					usersGuessOrOptionDirective != "stack" &&
					usersGuessOrOptionDirective != "exit" &&
					usersGuessOrOptionDirective != "dir" {
					//  v ^ v ^ at this point we know that the usersGuessOrOptionDirective is probably a valid guess,
					// ... AND, we need to leave the loop after processing it
					switch selectedExorcise {
					case "Romaji_Prompt": // 1
						meatOfRomajiNakedExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
						//check_for_match()
					case "Romaji+Kata_Prompt": // 2
						meatOfRomajiKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					case "Kata_Prompt-Respond-w-Hira|Romaji": // 3 and 4
						meatOfKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					}
					break outOfForLoop
				} else { // It must be a successive directive, so we continue to process it from the top of the loop
					continue outOfForLoop // <-- Immediately reiterate using a labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "Romaji_Prompt":
					meatOfRomajiNakedExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					//check_for_match()
				case "Romaji+Kata_Prompt":
					meatOfRomajiKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				case "Kata_Prompt-Respond-w-Hira|Romaji":
					meatOfKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // end of loop used to handel successive directives without double processing of guesses
	}
}

func branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise string) {
	switch usersGuessOrOptionDirective {
	case "menu":
		// Flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		usersGuessOrOptionDirective = "null"
		do_betweenMainMenuSelectionsTTE(selectedExorcise)
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
		handle_singleQuestMark_contextSensitive_directive(selectedExorcise)
		// The above ^^^ func must handle the case of being called from a naked kata (Romaji Prompt) activity w v v v
		//giveHintInResponseToSingleQuestionMarkContextSensitiveDir_sansRomaji(key, sansRomaji_hint string)
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
	case "rm":
		read_map_of_fineOn()
	case "stack":
		stack_the_map()
	}
}

func check(e error) { // create a func named check which takes one parameter "e" of type error
	if e != nil {
		panic(e) // use panic() to display error code
	}
}
