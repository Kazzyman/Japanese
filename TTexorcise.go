package main

// **do-this**
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func TouchTypingExorcise(selectedExorcise string) {
	// Accepts the following for ^ ^ ^ ^ :  RomajiNakedPrompt -- RomajiKataPrompt -- KataNakedPrompt
	rand.Seed(time.Now().UnixNano())
	var usersGuessOrOptionDirective string
	for {
		// Next, pick and assigns a random card to aCard and key globals

		/*
			// this switch is dumb
			switch selectedExorcise {
			case "RomajiNakedPrompt":
				pickARandomCardAndAssign() // ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
			case "RomajiKataPrompt":
				pickARandomCardAndAssign() // ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
			case "KataNakedPrompt":
				pickARandomCardAndAssign() // ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
			case "RespondWithRomajiToNakedKata":
				pickARandomCardAndAssign()
			}
		*/
		pickARandomCardAndAssign()
		// Next, prompt with the new random card and accept guess

		switch selectedExorcise {
		case "RomajiNakedPrompt":
			usersGuessOrOptionDirective = PromptWithOptionsAndScanForNakedRomajiPrompt(aCard.KeyR)
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		case "RomajiKataPrompt":
			usersGuessOrOptionDirective = PromptWithOptionsAndScanForRKprompt(aCard.KeyRK)
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		case "KataNakedPrompt":
			usersGuessOrOptionDirective = PromptWithOptionsAndScanForNakedKataPrompt(aCard.KeyK)
		// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
		case "RespondWithRomajiToNakedKata":
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
				switch selectedExorcise {
				case "RomajiNakedPrompt":
					usersGuessOrOptionDirective = PromptWithOptionsAndScanForRKprompt(aCard.KeyR)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "RomajiKataPrompt":
					usersGuessOrOptionDirective = PromptWithOptionsAndScanForRKprompt(aCard.KeyRK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "KataNakedPrompt":
					usersGuessOrOptionDirective = PromptWithOptionsAndScanForRKprompt(aCard.KeyK)
					// ... , if it is executed after a directive is handled, will be prompting from the same card -- so it is all good
				case "RespondWithRomajiToNakedKata":
					usersGuessOrOptionDirective = PromptWithOptionsAndScanForRespondWithRomajiToNakedKataPrompt(aCard.KeyK)
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

					switch selectedExorcise {
					case "RomajiNakedPrompt":
						meatOfRomajiNakedExorcise(usersGuessOrOptionDirective, aCard.KeyR, aCard.Value) // <-- may or may not be a guess, so check it
					case "RomajiKataPrompt":
						meatOfRomajiKataExorcise(usersGuessOrOptionDirective, aCard.Value) // <-- may or may not be a guess, so check it
					case "KataNakedPrompt":
						meatOfNakedKataExorcise(usersGuessOrOptionDirective, aCard.KeyH, aCard.Value) // <-- may or may not be a guess, so check it
					case "RespondWithRomajiToNakedKata":
						meatOfNakedKataExorcise(usersGuessOrOptionDirective, aCard.KeyR, aCard.Value) // Value is always same as KeyH : the hiragana
						// .Value is used as 'value' to look-up hints in the meat file
					}

					break outOfForLoop
				} else { // It must be a successive directive, so we continue to process it from the top of the loop
					continue outOfForLoop // <-- Immediately reiterate using a labeled loop; CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
				}
			} else {
				switch selectedExorcise {
				case "RomajiNakedPrompt":
					meatOfRomajiNakedExorcise(usersGuessOrOptionDirective, aCard.KeyH, aCard.Value) // <-- may or may not be a guess, so check it
				case "RomajiKataPrompt":
					meatOfRomajiKataExorcise(usersGuessOrOptionDirective, aCard.Value) // <-- may or may not be a guess, so check it
				case "KataNakedPrompt":
					meatOfNakedKataExorcise(usersGuessOrOptionDirective, aCard.Value, aCard.Value) // <-- may or may not be a guess, so check it
				case "RespondWithRomajiToNakedKata":
					meatOfNakedKataExorcise(usersGuessOrOptionDirective, aCard.KeyR, aCard.Value) // Value is always same as KeyH : the hiragana
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
