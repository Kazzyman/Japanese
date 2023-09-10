package main

// **do-this**
import (
	"math/rand"
	"time"
)

// TouchTypingExorcise
// The content of "prompt" (i.e., aCard.KeyR | aCard.KeyRK | aCard.KeyK) is set by the calling activity
//
// Used exclusively for exorcise 6
var Mixed_prompts_KeyH string
var Mixed_prompts_KeyK string
var Mixed_prompts_KeyX string
var Mixed_prompt_is string

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
	///
	///
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
			//
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
		// in any case:
		/*
			if Mixed_prompt_is == "kata" {
				check_it_for_fine_on() // Checks for Romaji only
			} else if Mixed_prompt_is == "hira" {
				check_it_for_fine_onH() // Checks for Hiragana only
			}
		*/
		if selectedExorcise == "Hira_prompt" {
			check_it_for_fine_on6() // Attempts to check for both Hira and Romaji
		} else {
			check_it_for_fine_on() // Checks for Romaji only
		}
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
		//
		//case of exorcise 6
		case "Hira_prompt":
			Mixed_prompts_KeyH = aCard.KeyH
			Mixed_prompts_KeyK = aCard.KeyK
			// Generate random 0 or 1
			random := rand.Intn(2)

			if random == 0 {
				Mixed_prompt_is = "hira"
				Mixed_prompts_KeyX = aCard.KeyH
				usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Mixed_prompts_KeyH) // RomajiResponse prompt, passing H
			} else {
				Mixed_prompt_is = "kata"
				Mixed_prompts_KeyX = aCard.KeyK
				usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(Mixed_prompts_KeyK) // RomajiResponse prompt, passing K
			}
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
				case "Hira_prompt":
					//for {
					// Generate random 0 or 1
					random := rand.Intn(2)

					if random == 0 {
						Mixed_prompt_is = "hira"
						Mixed_prompts_KeyX = aCard.KeyH
						usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(aCard.KeyH) // RomajiResponse prompt, passing H
					} else {
						Mixed_prompt_is = "kata"
						Mixed_prompts_KeyX = aCard.KeyK
						usersGuessOrOptionDirective = prompt_and_Scan_4_RomajiResponse_to_HiraPrompt(aCard.KeyK) // RomajiResponse prompt, passing K
					}
					//}
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
					case "Hira_prompt":
						meatOf_Mixed_HiraKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
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
				case "Hira_prompt":
					meatOf_Mixed_HiraKataExorcise(usersGuessOrOptionDirective, true) // <-- may or may not be a guess, so check it
				}
				// It is probably a valid guess, AND we need to leave the loop after processing it
				break outOfForLoop
			}
		} // end of loop used to handel successive directives without double processing of guesses
	}
}
