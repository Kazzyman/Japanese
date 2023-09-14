package main

// Finished September 12 of 2023 at 0112 hours
// TouchTypingExorcise
// The content of "prompt" (i.e., aCardS.KeyH | aCardS.KeyK) is set by the calling activity
//
var lastNonRandomCard = len(fileOfCardsS) // Index of the last card in the deck
var nonRandomCard = 0                     // Index of the first card in the deck

func TouchTypingExorciseSequential(selectedExorcise string) {
	// ^ ^ ^ ^ Accepts the following for ^ ^ ^ ^ :  "Sequential_Hira" or "Sequential_Kata" ...
	// ... given in mainMenuPromptScanSelectAndBeginSelectedExorcise() , found in menu.go

	var usersGuessOrOptionDirective string // It's all in the name
	//
	// --- Here begins the main loop, which encapsulates the entirety of this func ---
	for {
		if nonRandomCard == lastNonRandomCard { // If the last-read card was the last card in the deck
			nonRandomCard = 0 // ... Then, begin a second pass through the deck
		} else {
			aCardS = fileOfCardsS[nonRandomCard] // Get the next (and, eventually the last) card
			nonRandomCard++                      // On the last increment, nonRandomCard will be the address of the last card in the deck
		}
		switch selectedExorcise {
		// Prompt with the appropriate field from the latest-read card and accept user's guess or Option Directive:
		// case of exorcise 8
		case "Sequential_Kata": // 8
			usersGuessOrOptionDirective = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK) // Kata prompt
			// ... if executed after a Directive is handled, will be prompting from the same card -- it's all good
		case "Sequential_Hira": // 9
			usersGuessOrOptionDirective = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // Hira prompt
			// ... if executed after a Directive is handled, will be prompting from the same card -- it's all good
		}

	outOfForLoop: // This loop is needed because we want to handel the case of successive Directives (tricky) ...
		// ... the labeled 'for-loop' is used to handel successive Directives WITHOUT DOUBLE PROCESSING of guesses
		for {
			if usersGuessOrOptionDirective == "set" || usersGuessOrOptionDirective == "?" || // <-- If it IS a Directive
				usersGuessOrOptionDirective == "??" || usersGuessOrOptionDirective == "menu" ||
				usersGuessOrOptionDirective == "reset" || usersGuessOrOptionDirective == "stat" ||
				usersGuessOrOptionDirective == "dir" || usersGuessOrOptionDirective == "notes" ||
				usersGuessOrOptionDirective == "quit" || usersGuessOrOptionDirective == "exit" ||
				usersGuessOrOptionDirective == "stats" || usersGuessOrOptionDirective == "rm" ||
				usersGuessOrOptionDirective == "stack" { // <-- If it IS a Directive: v v v
				// v v v v v   HANDEL A DIRECTIVE    v v v v v v v v v v v v v v   v v v v v v v
				branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise) // <-- Do Directive
				// v v v v v   RE-PROMPT   v v v v v v
				// -- Re-prompt following the execution of a Directive ---------------------------------------- v v v v
				// ****************************************************************************************************
				switch selectedExorcise { // Identical to the above switch for prompt selection ***********************
				case "Sequential_Kata": // 8
					usersGuessOrOptionDirective = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyK) // Kata prompt
				// ... if executed after a Directive is handled, will be prompting from the same card -- it's all good
				case "Sequential_Hira": // 9
					usersGuessOrOptionDirective = Prompt_Scan_4_Romaji_or_HiraResponse(aCardS.KeyH) // Hira prompt
				}
				// v v v v v  Do not process directives from the above re-prompting   v v v v v
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
					//  v ^ v ^ At this point we know that the usersGuessOrOptionDirective is probably a valid guess,
					// ... AND, we will therefore, need to leave the loop after processing the user's guess
					// Based on the Selected Exorcise, process the user's guess (determine if it is correct, etc.)
					switch selectedExorcise {
					case "Sequential_Kata": // 8
						meatOfSequentialKata(usersGuessOrOptionDirective, true)
					case "Sequential_Hira": // 9
						meatOfSequentialHira(usersGuessOrOptionDirective, true)
					}
					break outOfForLoop
				} else { // It must be a successive Directive, so we continue to process it from the top of the loop,
					continue outOfForLoop // reiterate at labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "Sequential_Kata": // 8
					meatOfSequentialKata(usersGuessOrOptionDirective, true)
				case "Sequential_Hira": // 9
					meatOfSequentialHira(usersGuessOrOptionDirective, true)
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // End of loop used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
	} // End of outer loop
}
