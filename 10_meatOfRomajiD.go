package main

import (
	"fmt"
)
/*
The first of these two functions is passed the user's first guess, and, potentially, it obtains a second guess
.. the second func is passed that second guess, and, potentially, it obtains the third and final guess
Finally, the first func is re-called by the second,: the first func is, then, passed the third and final guess

Since this Exercise (10) expects only Hira guesses, we test for embedded Directives with a simple Alpha test
*/
func meatOfRomajiExerciseD(in string, skipFlag bool) { //  - -
	if aCardD.KeyR == "zu" {
		if in == "づ" || in == "ず" {
			log_right(aCardD.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
		} else {
			log_oops(aCardD.KeyR, aCardD.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
				fmt.Println("Try again") // ... in other words, this only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, second guess
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
				// If "in" was a Directive, i.e., a string containing an Alpha char
				isDir := if_it_is_a_Directive(in)
				if isDir {
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_PromptD")
					// Re-prompt, second guess 
					in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
				} // else:
				secondTry_meatOfRomajiExerciseD(in)
			}
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCardD.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCardD.Hint1h, aCardD.Hint2k, aCardD.Hint3TT)
				fmt.Println("")
			}
		} // If the user was ^^Right!, then we return to TouchTypingExercises12346_10(SE) (essentially from this line)
		
	} else 
	// Not "zu" [except for two lines, i.e.: if in == "づ" || in == "ず" {  ~~>  if in == aCardD.KeyH {
	// , and the one commented-out line, the remainder of this func is a copy and paste from above] plus one \n
		if in == aCardD.KeyH {
			log_right(aCardD.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! \n")
			fmt.Printf("%s", colorReset)
			// fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
		} else {
			log_oops(aCardD.KeyR, aCardD.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
				fmt.Println("Try again") // ... in other words, this only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, second guess
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
				// If "in" was a Directive, i.e., a string containing an Alpha char
				isDir := if_it_is_a_Directive(in)
				if isDir {
					branchOnUserSelectedDirectiveIfGiven(in, "Romaji_PromptD")
					// Re-prompt, second guess 
					in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
				} // else:
				secondTry_meatOfRomajiExerciseD(in)
			}
			if skipFlag == false { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// If the user guesses incorrectly on third-and-final try, then, and only then, execute the following
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCardD.KeyH)
				fmt.Printf("%s", colorReset)
				// Display hints in white
				fmt.Printf("\n\n%s\n%s\n%s\n", aCardD.Hint1h, aCardD.Hint2k, aCardD.Hint3TT)
				fmt.Println("")
			}
		} 
}
/*
.
*/
// Second-Try version of the above func (tests the second guess, and then obtains the third and final one)
func secondTry_meatOfRomajiExerciseD(in string) { // NOTE: we have already been prompted with KeyR  - -
	if aCardD.KeyR == "zu" {
		if in == "づ" || in == "ず" {
			log_right(aCardD.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
		} else {
			log_oops(aCardD.KeyR, aCardD.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! Try again\n")
			fmt.Printf(colorReset)
			// Re-prompt, final guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
			// If "in" was a Directive, i.e., a string containing an Alpha char
			isDir := if_it_is_a_Directive(in)
			if isDir {
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_PromptD")
				// Re-prompt, final guess 
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
			} // else:
			meatOfRomajiExerciseD(in, false)// Process the third and final guess
		}
	} else
	// Not "zu" [except for two lines, i.e.: if in == "づ" || in == "ず" {  ~~>  if in == aCardD.KeyH {
	// , and the one commented-out line, the remainder of this func is a copy and paste from above] plus one \n
		if in == aCardD.KeyH {
			log_right(aCardD.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! \n")
			fmt.Printf("%s", colorReset)
			// fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
		} else {
			log_oops(aCardD.KeyR, aCardD.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! Try again\n")
			fmt.Printf(colorReset)
			// Re-prompt, final guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
			// If "in" was a Directive, i.e., a string containing an Alpha char
			isDir := if_it_is_a_Directive(in)
			if isDir {
				branchOnUserSelectedDirectiveIfGiven(in, "Romaji_PromptD")
				// Re-prompt, final guess 
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR)
			} // else:
			meatOfRomajiExerciseD(in, false)// Process the third and final guess
		}
}
