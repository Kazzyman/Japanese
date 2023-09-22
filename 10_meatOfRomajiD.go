package main

import (
	"fmt"
)
/*
The first of these two functions is passed the user's first guess, and, potentially, it obtains a second guess
.. the second func is passed that second guess, and, potentially, it obtains the third and final guess
Finally, the first func is re-called by the second: the first func is, then, passed the third and final guess ... 
... with skipFlag false signifying that the offer to try again is to be replaced with a messages containing the correct 
answer, and also a set of hints relating to how one could have known the correct answer. 

Since this Exercise (10) expects only Hira guesses, we test for and handle Directives with a simple Alpha test in 
... semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt()
*/
func meatOfRomajiExerciseD(in, selectedExercise string, skipFlag bool) { //  - -
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
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR, selectedExercise)
				secondTry_meatOfRomajiExerciseD(in, selectedExercise)
			} else
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
			// commented-out line
		} else {
			log_oops(aCardD.KeyR, aCardD.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			if skipFlag == true { // skipFlag is true on first entry, & false when recalled after third failed attempt
				// When current func is re-called by secondTry_, with skipFlag false, skip this, and do next ...
				fmt.Println("Try again") // ... in other words, this only prints on first pass
				fmt.Printf(colorReset)
				// Re-prompt, second guess
				in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR, selectedExercise)
				secondTry_meatOfRomajiExerciseD(in, selectedExercise)
			} else 
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
func secondTry_meatOfRomajiExerciseD(in, selectedExercise string) { //  - -
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
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR, selectedExercise)
			meatOfRomajiExerciseD(in, selectedExercise, false)// Process the third and final guess
		}
	} else
	// Not "zu" [except for two lines, i.e.: if in == "づ" || in == "ず" {  ~~>  if in == aCardD.KeyH {
	// , and the one commented-out line, the remainder of this func is a copy and paste from above] plus one \n
		if in == aCardD.KeyH {
			log_right(aCardD.KeyR)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! \n")
			fmt.Printf("%s", colorReset)
			// commented-out line
		} else {
			log_oops(aCardD.KeyR, aCardD.KeyH, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! Try again\n")
			fmt.Printf(colorReset)
			// Re-prompt, final guess
			in = semi_Universal_Prompt_Scan_4_HiraResponse_to_RomajiPrompt(aCardD.KeyR, selectedExercise)
			meatOfRomajiExerciseD(in, selectedExercise,false)// Process the third and final guess
		}
}