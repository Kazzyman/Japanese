package main

// TouchTypingExorcise
// The content of "prompt" (i.e., aCardD.KeyR | aCardD.KeyRK | aCardD.KeyK) is set by the calling activity
//
var lastNonRandomCard = len(fileOfCards) - 1
var nonRandomCard = 0

func TouchTypingExorciseSequential(selectedExorcise string) {
	// Accepts the following for ^ ^ ^ ^ :  RomajiNakedPrompt -- RomajiKataPrompt -- KataNakedPrompt
	//

	var usersGuessOrOptionDirective string
	//
	for {
		if nonRandomCard == lastNonRandomCard {
			nonRandomCard = 1
		} else {
			aCardS = fileOfCardsS[nonRandomCard]
			nonRandomCard++
		}
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
		// case of exorcise 8
		case "Sequential-Kata":
			usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK) // Kata prompt, passing K
			// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		case "Sequential-Hira": // 9
			usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH)
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
				case "Sequential-Kata": // 8
					usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK)
				// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "Sequential-Hira": // 9
					usersGuessOrOptionDirective = Kata_Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH)
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
					case "Sequential-Kata": // 3 and 4
						meatOfSequentialKata(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
					case "Sequential-Hira": // 9
						meatOfSequentialHira(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
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
				case "Sequential-Kata":
					meatOfSequentialKata(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				case "Sequential-Hira": // 9
					meatOfSequentialHira(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // end of loop used to handel successive directives without double processing of guesses
	}
}
