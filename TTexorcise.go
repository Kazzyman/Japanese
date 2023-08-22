package main

// **do-this**
import (
	"fmt"
	"os"
)

func TouchTypingExorcise(selectedExorcise string) {
	// Accepts the following for ^ ^ ^ ^ :  RomajiNakedPrompt -- RomajiKataPrompt -- KataNakedPrompt
	//
	// rand.Seed(seed) is now done in main.go, at the top of main()
	// was: // rand.Seed(time.Now().UnixNano())
	// Example of seed Use per Claude: fmt.Println(rand.Int())

	var usersGuessOrOptionDirective string
	for {
		pickARandomCardAndAssign() // Assigns a random card to aCard, so there is no need to pass any aCard fields below
		// Next, prompt with the appropriate field from the new random card and accept user's guess
		/*
			case:  1 = "RomajiNakedPrompt"
			case:  2 = "RomajiKataPrompt"
			case:  3 = "KataNakedPrompt"
			case:  4 = "RespondWithRomajiToNakedKata"
		*/
		// Prompt with the appropriet field from the new random card and accept user's guess:
		switch selectedExorcise {
		// case 1
		case "RomajiNakedPrompt":
			usersGuessOrOptionDirective = PromptWithOptionsAndScan(aCard.KeyR) // universal prompt, passing R
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		// case 2
		case "RomajiKataPrompt":
			// PromptWithOptionsAndScan(aCard.KeyRK)
			usersGuessOrOptionDirective = PromptWithOptionsAndScan(aCard.KeyRK) // universal prompt, passing RK
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		//
		// These next two should be re-evaluated for justifyability
		// case 3
		case "KataNakedPrompt":
			usersGuessOrOptionDirective = PromptWithOptionsAndScan(aCard.KeyK) // universal prompt, passing K
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		// case 4
		case "RespondWithRomajiToNakedKata": // special prompt (needed??), passing K, as does the other kata prompt, above
			usersGuessOrOptionDirective = PromptWithOptionsAndScanForRespondWithRomajiToNakedKataPrompt(aCard.KeyK)
		}

	outOfForLoop: // this loop is only needed because we want to handel the case of successive directives (tricky)
		// ESPECIALLY, this labeled 'for-loop' is used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
		for {
			if usersGuessOrOptionDirective == "set" || usersGuessOrOptionDirective == "?" || // <-- if it IS a directive
				usersGuessOrOptionDirective == "??" || usersGuessOrOptionDirective == "menu" || usersGuessOrOptionDirective == "reset" ||
				usersGuessOrOptionDirective == "stat" || usersGuessOrOptionDirective == "dir" || usersGuessOrOptionDirective == "notes" ||
				usersGuessOrOptionDirective == "quit" || usersGuessOrOptionDirective == "exit" {
				// v v v v v   HANDEL A DIRECTIVE  v v v v v v v v
				branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise) // <-- handle the directive
				// v v v v v   re-prompt   v v v v v v
				// -- re-prompt following the execution of a directive -------------------------------------------------------- v v v v
				/*
					case:  1 = "RomajiNakedPrompt"
					case:  2 = "RomajiKataPrompt"
					case:  3 = "KataNakedPrompt"
					case:  4 = "RespondWithRomajiToNakedKata"
				*/
				// ****************************************************************************************************
				switch selectedExorcise { // Identical to the above switch for prompt selection ***********************
				case "RomajiNakedPrompt":
					usersGuessOrOptionDirective = PromptWithOptionsAndScan(aCard.KeyR)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "RomajiKataPrompt":
					usersGuessOrOptionDirective = PromptWithOptionsAndScan(aCard.KeyRK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "KataNakedPrompt":
					usersGuessOrOptionDirective = PromptWithOptionsAndScan(aCard.KeyK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "RespondWithRomajiToNakedKata":
					usersGuessOrOptionDirective = PromptWithOptionsAndScanForRespondWithRomajiToNakedKataPrompt(aCard.KeyK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				}
				// v v v v v   do not process directives from the re-prompting   v v v v v
				if usersGuessOrOptionDirective != "set" &&
					usersGuessOrOptionDirective != "?" &&
					usersGuessOrOptionDirective != "??" &&
					usersGuessOrOptionDirective != "menu" &&
					usersGuessOrOptionDirective != "reset" &&
					usersGuessOrOptionDirective != "stat" &&
					usersGuessOrOptionDirective != "quit" &&
					usersGuessOrOptionDirective != "notes" &&
					usersGuessOrOptionDirective != "exit" &&
					usersGuessOrOptionDirective != "dir" {
					//  v ^ v ^ at this point we know that the usersGuessOrOptionDirective is probably a valid guess,
					// ... AND, we need to leave the loop after processing it

					/*
						case:  1 = "RomajiNakedPrompt"
						case:  2 = "RomajiKataPrompt"
						case:  3 = "KataNakedPrompt"
						case:  4 = "RespondWithRomajiToNakedKata"
					*/
					// Only need to be passing the user's guess, as both the prompt and the other aCard fields can now be acessed directly fr within each
					switch selectedExorcise {
					case "RomajiNakedPrompt":
						meatOfRomajiNakedExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
					case "RomajiKataPrompt":
						meatOfRomajiKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
					// The next two call the same func???
					case "KataNakedPrompt":
						meatOfNakedKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
					case "RespondWithRomajiToNakedKata":
						meatOfNakedKataExorcise(usersGuessOrOptionDirective)
					}

					break outOfForLoop
				} else { // It must be a successive directive, so we continue to process it from the top of the loop
					continue outOfForLoop // <-- Immediately reiterate using a labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "RomajiNakedPrompt":
					meatOfRomajiNakedExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
				case "RomajiKataPrompt":
					meatOfRomajiKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
				case "KataNakedPrompt":
					meatOfNakedKataExorcise(usersGuessOrOptionDirective) // <-- may or may not be a guess, so check it
				case "RespondWithRomajiToNakedKata":
					meatOfNakedKataExorcise(usersGuessOrOptionDirective) // Value is always same as KeyH : the hiragana
					// .Value is used as 'value' to look-up hints in the meat file
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
		// flush the old stats and hits arrays
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
	case "??":
		doubleQuestMark()
	case "?":
		singleQuestMark()
	case "set":
		setKey()
	case "stat":
		hits()
	case "notes":
		fmt.Println("\nIn the traditional Hepburn romanization system, the sound じ in hiragana is romanized as \"ji\" \n" +
			"and the katakana ジ is also romanized as \"ji\" \n\n" +
			"However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as\n" +
			" \"zi\" instead of \"ji\"\n\n" +
			"The sound gi:ぎ in hiragana is romanized as \"gi\" and the katakana ギ is also romanized as \"gi\"\n")
		fmt.Println("゜is called \"handakuten\" 半濁点 translates to \"half-voicing mark\" or \"semi-voiced mark\"\n" +
			"゛is called \"dakuten\" 濁点 meaning 'voiced mark' or 'voicing mark'")
	case "dir":
		switch selectedExorcise {
		case "RomajiNakedPrompt":
			reDisplayNakedRomajiOptions()
		case "RomajiKataPrompt":
			reDisplayRomajiKataOptions()
		case "KataNakedPrompt":
			reDisplayNakedKataOptions()
		}
	}
}
