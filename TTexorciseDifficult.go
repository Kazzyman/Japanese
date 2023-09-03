package main

import (
	"math/rand"
	"time"
)

// TouchTypingExorcise
// The content of "prompt" (i.e., aCardD.KeyR | aCardD.KeyRK | aCardD.KeyK) is set by the calling activity
//
func TouchTypingExorciseDifficult(selectedExorcise string) {
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
		pick_Difficult_RandomCard_Assign_aCard() // Assigns a random card to the global aCard
		/*
			if isThis_a_cardWeNeedMoreWorkOn > 2 {   // if we have gone without augmenting the random picks with cards we have prev missed ...
				//

					// Log to a file that this action was taken **do-this**
					fileHandleBig, err := os.OpenFile("JapLogD.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                  // ... gets a file handle to JapLog.txt
					//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
					_, err2 := fmt.Fprintf(fileHandleBig, "\nChecked card:%s in the frequencyMapOf_need_workOn after:%d cycles\n",
						aCardD.KeyH, isThis_a_cardWeNeedMoreWorkOn)
					check(err2)


				isThis_a_cardWeNeedMoreWorkOn = 0 // ... for a while now, then let's go get a card we've missed before, instead of that random one
				check_it_for_needing_more_practiceD()

			}

		*/
		// in any case:
		//check_it_for_fine_onD()
		isThis_a_cardWeNeedMoreWorkOn++

		// Prompt with the appropriate field from the new random card and accept user's guess:
		switch selectedExorcise {
		//
		//case of exorcise 1
		case "Romaji_Prompt":
			usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCardD.KeyR) // Semi-universal prompt, passing R
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		//
		//case of exorcise 2
		case "Romaji+Kata_Prompt":
			usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCardD.KeyRK) // Semi-universal prompt, passing RK
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		//
		// case of exorcise 3 and 4
		case "Kata_Prompt-Respond-w-Hira|Romaji":
			usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK) // Kata prompt, passing K
			// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		case "Most_Difficult":
			usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK) // Kata prompt, passing K
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
					usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCardD.KeyR)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "Romaji+Kata_Prompt": // 2
					usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCardD.KeyRK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "Kata_Prompt-Respond-w-Hira|Romaji": // 3 or 4
					usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK)
				// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "Most_Difficult":
					usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardD.KeyK)
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
						meatOfRomajiExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
						//check_for_match()
					case "Romaji+Kata_Prompt": // 2
						meatOfRomajiKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					case "Kata_Prompt-Respond-w-Hira|Romaji": // 3 and 4
						meatOfKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					case "Most_Difficult":
						meatOfKataExorciseD(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					}
					break outOfForLoop
				} else { // It must be a successive directive, so we continue to process it from the top of the loop
					continue outOfForLoop // <-- Immediately reiterate using a labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "Romaji_Prompt":
					meatOfRomajiExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					//check_for_match()
				case "Romaji+Kata_Prompt":
					meatOfRomajiKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				case "Kata_Prompt-Respond-w-Hira|Romaji":
					meatOfKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				case "Most_Difficult":
					meatOfKataExorciseD(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // end of loop used to handel successive directives without double processing of guesses
	}
}
