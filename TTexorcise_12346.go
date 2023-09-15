package main

// **do-this**
import (
	"math/rand"
	"time"
)

// The content of "prompt" (i.e., aCard.KeyR | aCard.KeyRK | aCard.KeyK) is set by the calling activity
//
// Used exclusively for exorcise 6
var Mixed_prompts_KeyH string
var Mixed_prompts_KeyK string
var Mixed_prompts_KeyX string
var Mixed_prompt_is string

// Entry point for numbers 1, 2, 3, and 6 Exorcises
func TouchTypingExorcise(selectedExorcise string) { //       - -
	//
	rand.Seed(time.Now().UnixNano())

	var usersGuessOrOptionDirective string // It's all in the name
	//
	isThis_a_cardWeNeedMoreWorkOn := 0 // now and then we will consider working on a char the user has had trouble with
	// /
	// /
	// --- Here begins the main loop, which encapsulates the entirety of this func ---
	for {
		// Does not work well with duplicates in file
		/*
			if fromEndOfFileCounter%2 == 0 {
				pick_RandomCard_Assign_aCard() // Assigns a random card to the global aCard
				fromEndOfFileCounter++
			} else {
				aCard = fileOfCards[nonRandomCardFromEnd]
				nonRandomCardFromEnd--
			}
		*/
		pick_RandomCard_Assign_aCard() // Assigns a random card to the global aCard

		if isThis_a_cardWeNeedMoreWorkOn > 2 { // if we have gone without augmenting the random picks with cards we have prev missed ...
			/*
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				_, err2 := fmt.Fprintf(fileHandleBig, "\nChecked card:%s in the frequencyMapOf_need_workOn after:%d cycles\n",
					aCard.KeyH, isThis_a_cardWeNeedMoreWorkOn)
				check(err2)
				fileHandleBig.Close()
			*/
			isThis_a_cardWeNeedMoreWorkOn = 0    // ... for a while now, then let's go get a card we've missed before, instead of that random one
			check_it_for_needing_more_practice() // **do-this** this func probably need some work to function with exorcise 6
		}
		// In any case:
		if selectedExorcise == "Mixed_prompts" {
			check_it_for_fine_on6() // Attempts to check for both Hira and Romaji
		} else {
			check_it_for_fine_on() // Checks for Romaji only
		}
		isThis_a_cardWeNeedMoreWorkOn++

		// Prompt with the appropriate field from the latest-read card and accept user's guess or Option Directive:
		switch selectedExorcise {
		//
		// case of exorcise 1
		case "Romaji_Prompt":
			// Semi-universal prompt, passing R
			usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
		// ... if executed after a Directive is handled, will be prompting from the same card -- it's all good
		//
		// case of exorcise 2
		case "Romaji_w_Kata_Prompt":
			// Semi-universal prompt, passing RK
			usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
		// ... if executed after a Directive is handled, will be prompting from the same card -- it's all good
		//
		// case of exorcise 3 and 4
		case "Respond_w_Hira_or_Romaji":
			usersGuessOrOptionDirective = Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK) // Kata prompt, passing K
			// ... if executed after a Directive is handled, will be prompting from the same card -- it's all good
		//
		// case of exorcise 6
		case "Mixed_prompts":
			Mixed_prompts_KeyH = aCard.KeyH
			Mixed_prompts_KeyK = aCard.KeyK
			// Generate random 0 or 1
			random := rand.Intn(2)

			if random == 0 {
				Mixed_prompt_is = "hira"
				Mixed_prompts_KeyX = aCard.KeyH
				usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Mixed_prompts_KeyH)
			} else {
				Mixed_prompt_is = "kata"
				Mixed_prompts_KeyX = aCard.KeyK
				usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Mixed_prompts_KeyK)
			}
		}

	outOfForLoop: // this loop is only needed because we want to handel the case of successive directives (tricky)
		// ESPECIALLY, this labeled 'for-loop' is used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
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
				case "Romaji_Prompt": // 1
					usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyR)
					// If executed after a Directive is handled, will be prompting from the same card -- it's all good
				case "Romaji_w_Kata_Prompt": // 2
					usersGuessOrOptionDirective = semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(aCard.KeyRK)
					// ... , as before -- it is all good
				case "Respond_w_Hira_or_Romaji": // 3 or 4
					usersGuessOrOptionDirective = Prompt_Scan_4_Romaji_or_HiraResponse(aCard.KeyK)
					// ... , as before -- it is all good
				case "Mixed_prompts": // 6
					// Generate random 0 or 1
					Mixed_prompts_KeyH = aCard.KeyH
					Mixed_prompts_KeyK = aCard.KeyK
					random := rand.Intn(2)

					if random == 0 {
						Mixed_prompt_is = "hira"
						Mixed_prompts_KeyX = aCard.KeyH
						usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Mixed_prompts_KeyH)
					} else {
						Mixed_prompt_is = "kata"
						Mixed_prompts_KeyX = aCard.KeyK
						usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Mixed_prompts_KeyK)
					}
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
					case "Romaji_Prompt": // 1
						meatOfRomajiExorcise(usersGuessOrOptionDirective, true)
					case "Romaji_w_Kata_Prompt": // 2
						meatOfRomajiKataExorcise(usersGuessOrOptionDirective, true)
					case "Respond_w_Hira_or_Romaji": // 3 and 4
						meatOfKataExorcise(usersGuessOrOptionDirective, true)
					case "Mixed_prompts": // 6
						meatOf_Mixed_HiraKataExorcise(usersGuessOrOptionDirective, true)
					}
					break outOfForLoop
				} else { // It must be a successive Directive, so we continue to process it from the top of the loop,
					continue outOfForLoop // reiterate at labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "Romaji_Prompt": // 1
					meatOfRomajiExorcise(usersGuessOrOptionDirective, true)
				case "Romaji_w_Kata_Prompt": // 2
					meatOfRomajiKataExorcise(usersGuessOrOptionDirective, true)
				case "Respond_w_Hira_or_Romaji": // 3, 4
					meatOfKataExorcise(usersGuessOrOptionDirective, true)
				case "Mixed_prompts": // 6
					meatOf_Mixed_HiraKataExorcise(usersGuessOrOptionDirective, true)
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // End of loop used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
	} // End of outer loop
}
