package main

// **do-this**
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

// The content of "prompt" (i.e., aCard.KeyR | aCard.KeyRK | aCard.KeyK) is set by the calling activity
//	                                     ^ 1           ^ 2         ^ 3,4

// Ask the user for a Hiragana response (universal for both of the Romaji-containing prompts):
//
//goland:noinspection ALL  **do-this**
func semi_Universal_Prompt_Scan_4_HiraResponse_NOT_a_KataPrompt(prompt string) (usersGuessOrOptionDirective string) {
	// The return signature (above) creates a local var usersGuessOrOptionDirective, used below in the Scan():
	fmt.Printf("%s", prompt)    // Prompt the user, in a pretty blue color: Options: '?' | '??'
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" Type the corresponding Hiragana, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help on a particular Hiragana char ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

//goland:noinspection ALL  **do-this**
// Ask the user for a Romaji|Hira response to the one (unique-and-naked) Katakana prompt:
//goland:noinspection ALL  **do-this**
func Kata_Prompt_Scan_4_Romaji_or_HiraResponse(prompt string) (usersGuessOrOptionDirective string) {
	// The return signature (above) creates a local var usersGuessOrOptionDirective, used below in the Scan():
	fmt.Printf("%s", prompt)    // Prompt the user, in a pretty blue color: Options: '?' | '??'
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" Type the corresponding Romaji|Hira, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help on a particular Hiragana char ... \n") //
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

//goland:noinspection ALL  **do-this**
func pick_RandomCard_Assign_aCard() {
	/*
	  // Use the seed (example by Claude)
	  rand.Seed(seed)
	  fmt.Println(rand.Int())
	*/
	randIndex := rand.Intn(len(fileOfCards))
	aCard = fileOfCards[randIndex] // randomly pick a 'card' from the 'deck' and store it in a global var
}
func pick_Difficult_RandomCard_Assign_aCard() {
	/*
	  // Use the seed (example by Claude)
	  rand.Seed(seed)
	  fmt.Println(rand.Int())
	*/
	randIndex := rand.Intn(len(fileOfCardsMostDifficult))
	aCardD = fileOfCardsMostDifficult[randIndex] // randomly pick a 'card' from the 'deck' and store it in a global var
}

// DIRECTIVES : --------------------------------------------------------------------------------------------
//goland:noinspection ALL  **do-this** Works with Hiragana or Romaji in all activities
func handle_doubleQuestMark_directive() {
	var hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom string // the value to find in locateCard.go
	//
	fmt.Printf("\n  -- Type either the Hiragana char or the Romaji that you need help with:> ")
	_, _ = fmt.Scan(&hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom)
	//
	locateCardAndDisplayHelpFieldsContainedInIt(hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom)
	fmt.Println("")
}

//goland:noinspection ALL  **do-this**
func handle_singleQuestMark_contextSensitive_directive(currentActivity string) {
	switch currentActivity {
	case "Romaji_Prompt":
		giveHintInResponseToSingleQuestionMarkContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
	case "Romaji+Kata_Prompt":
		giveHintInResponseToSingleQuestionMarkContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
	case "Kata_Prompt-Respond-w-Hira|Romaji":
		giveHintInResponseToSingleQuestionMarkContextSensitiveDir_sansRomaji(aCard.KeyH, aCard.HintSansR)
	case "Hira_prompt":
		giveHintInResponseToSingleQuestionMarkContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
	}
}

//goland:noinspection ALL  **do-this**
func reSet_aCard_andThereBy_reSet_thePromptString() {
	var theHiraganaOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Hiragana to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt & \"aCard\\.fields\" :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Hiragana char (instead of, accidentally, an alpha char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theHiraganaOfCardToSilentlyLocate):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric == true {
		fmt.Println("Are you in alphanumeric input mode?")
		fmt.Printf("... if so, change it to Hiragana (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in Hiragana mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)
		// May yet send an Alpha string to the next func, which will itself deal with it elegantly
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		fmt.Println("")
	}
}

func read_map_of_fineOn() {
	for s, f := range frequencyMapOf_IsFineOnChars {
		fmt.Printf(" --- From MapOf_IsFineOn: string is:")
		fmt.Printf(colorCyan)
		fmt.Printf("%s", s)
		fmt.Printf(colorReset)
		fmt.Printf(", freq is:")
		fmt.Printf(colorRed)
		fmt.Printf("%d", f)
		fmt.Printf(colorReset)
		fmt.Printf(" ---\n")
	}
	fmt.Println("")
}

// end of DIRECTIVES -----------------------------------------------------------------------------------

// Activity #5
func drillLines() {
	var firstSlice []string
	var secondSlice []string
	var thirdSlice []string
	var fourthSlice []string
	var fifthSlice []string

	firstSlice = []string{"あ", "い", "う", "え", "お", "か", "き", "く", "け", "こ", "さ", "し", "す", "せ", "そ",
		"た", "ち", "つ", "て", "と", "な", "に", "ぬ", "ね", "の"}

	//                      1     2     3    4     5     6    7    8     9     10   11    12    13   14    15   16    17    18    19   20   21
	secondSlice = []string{"は", "ひ", "ふ", "へ", "ほ", "ま", "み", "む", "め", "も", "や", "ゆ", "よ", "ら", "り", "る", "れ", "ろ", "わ", "を", "ん"}

	//                     1     2     3    4     5    6     7     8    9     10   11    12    13    14   15   16    17
	thirdSlice = []string{"が", "ぎ", "ぐ", "げ", "ご", "ざ", "じ", "ず", "ぜ", "ぞ", "だ", "ぢ", "づ", "で", "ど", "ば", "び",
		//   18   19    20    21    22   23    24   25
		"ぶ", "べ", "ぼ", "ぱ", "ぴ", "ぷ", "ぺ", "ぽ"}

	fourthSlice = []string{"ぎゃ", "ぎゅ", "ぎょ", "じゃ", "じゅ", "じょ", "ぢゃ", "ぢゅ", "ぢょ", "びゃ", "びゅ", "びょ",
		"ぴゃ", "ぴゅ", "ぴょ", "きゃ", "きゅ", "きょ", "しゃ", "しゅ", "しょ", "ちゃ", "ちゅ", "きょ"}

	fifthSlice = []string{"にゃ", "にゅ", "にょ", "ひゃ", "ひゅ", "ひょ", "みゃ", "みゅ", "みょ", "りゃ", "りゅ", "りょ"}
	//
	//
	fmt.Println("  1  2   3  4   5   6   7   8   9   10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25")
	fmt.Println("あ　い　う　え　お　か　き　く　け　こ　さ　し　す　せ　そ　た　ち　つ　て　と　な　に　ぬ　ね　の")
	for_drillLines_obtainAndDealWithUserInput(firstSlice)

	fmt.Println("  1  2   3  4   5  　6  7   8   9   10  11  12  13  14  15  16  17  18  19  20  21")
	fmt.Println("は　ひ　ふ　へ　ほ　ま　み　む　め　も　や　ゆ　よ　ら　り　る　れ　ろ　わ　を　ん")
	for_drillLines_obtainAndDealWithUserInput(secondSlice)

	fmt.Println("  1   2   3  4   5   6    7   8  9   10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25 ")
	fmt.Println("ga  gi  gu  ge  go,  za  ji  zu  ze  zo, da  ji  zu  de  do, ba  bi  bu  be  bo, pa  pi  pu  pe  po")
	for_drillLines_obtainAndDealWithUserInput(thirdSlice)

	fmt.Println(" 1     2      3    4     5   6    7    8    9   　10    11   12   13   14   15   16   17   18   19   20  21   22   23   24")
	fmt.Println("gya  　gyu 　 gyo  ja  　ju  jo  　ja  ju 　 jo  bya 　byu  byo  pya  pyu  pyo  kya  kyu  kyo  sha  shu  sho  cha  chu  cho")
	for_drillLines_obtainAndDealWithUserInput(fourthSlice)

	fmt.Println("1      2    3      4   　  5     6     7    8      9    10    11    12")
	fmt.Println("nya  　nyu  nyo 　 hya 　 hyu  hyo 　 mya  myu　  myo  rya 　 ryu  ryo")
	for_drillLines_obtainAndDealWithUserInput(fifthSlice)

	var pause string
	fmt.Println("Enter 'menu' to return to the main menu")
	_, err := fmt.Scan(&pause)
	if pause == "menu" {
		/*
			case "menu": // From func branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise string)
				// Flush the old stats and hits arrays
				cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
				cyclicArrayHits = CyclicArrayHits{}
				usersGuessOrOptionDirective = "null"
				do_betweenMainMenuSelectionsTTE(selectedExorcise)
				mainMenuPromptScanSelectAndBeginSelectedExorcise()
		*/
		// Flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// finalExorcise = "drillLines"
		//usersGuessOrOptionDirective = "null"
		do_betweenMainMenuSelectionsTTE("drillLines")
		mainMenuPromptScanSelectAndBeginSelectedExorcise()
	}
	if err != nil {
		return
	}
	//
}

// Used 5 times, above in: func drillLines()
func for_drillLines_obtainAndDealWithUserInput(refSlice []string) {
	// Get user input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Convert to strings
	chars := strings.Fields(input)

	match := true
	for i, v := range chars {
		if v != refSlice[i] {
			match = false
			fmt.Printf("\n Failure to match at possition: %d \n\n", i+1)
			if len(refSlice) == i {
				break
			}
		}
	}
	if match {
		fmt.Printf("\nGood job! The slices match: %t \n\n", match)
	}
	var pause string
	fmt.Println("Enter 'menu' to return to the main menu")
	_, err := fmt.Scan(&pause)
	if err != nil {
		// return
	}
	if pause == "menu" {
		/*
			case "menu": // From func branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExorcise string)
				// Flush the old stats and hits arrays
				cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
				cyclicArrayHits = CyclicArrayHits{}
				usersGuessOrOptionDirective = "null"
				do_betweenMainMenuSelectionsTTE(selectedExorcise)
				mainMenuPromptScanSelectAndBeginSelectedExorcise()
		*/
		// Flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		do_betweenMainMenuSelectionsTTE("drillLines")
		mainMenuPromptScanSelectAndBeginSelectedExorcise()
	}
}

/*
.
.
.
*/
// Globals:
// .
// A map v v v v v v v is used to store correct guesses, and is only written to after a correct guess
var frequencyMapOf_IsFineOnChars = make(map[string]int) // create the map, dir 'read_map_of_fineOn' reads it
var frequencyMapOf_need_workOn = make(map[string]int)

// .
// Parse the map to see if the new random card matches any members of the map
// Each time we get a new random card we check the map to see if it has been done correctly 3 or more times ...
// ... if it has, then we pick another card, and check it
// else we break
func check_it_for_fine_on() {
	for s, f := range frequencyMapOf_IsFineOnChars {
		if s == aCard.KeyR { // if it is in the map we need to check the freq
			if f >= 2 { // if the freq is 3+ we need another card
				//read_map_of_fineOn() // we show the map
				//fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
				_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
					s, aCard.KeyH, f)
				check(err2)
				pick_RandomCard_Assign_aCard() // We get that new card ...
				//fmt.Println(" ... so here is a new one ... \n")
				check_it_for_fine_on() // ... and we check THAT new card with a recursive call
			} else { // else the card had a freq less than 3, so ...
				break //  ... we exit the loop and the func -- we will use this card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
}

func check_it_for_fine_onD() {
	for s, f := range frequencyMapOf_IsFineOnChars {
		if s == aCardD.KeyR { // if it is in the map we need to check the freq
			if f >= 2 { // if the freq is 3+ we need another card
				//read_map_of_fineOn() // we show the map
				//fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
				/*
					// Log to a file that this action was taken **do-this**
					fileHandleBig, err := os.OpenFile("JapLogD.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                  // ... gets a file handle to JapLog.txt
					//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
					_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
						s, aCard.KeyH, f)
					check(err2)

				*/
				pick_Difficult_RandomCard_Assign_aCard() // We get that new card ...
				//fmt.Println(" ... so here is a new one ... \n")
				check_it_for_fine_onD() // ... and we check THAT new card with a recursive call
			} else { // else the card had a freq less than 3, so ...
				break //  ... we exit the loop and the func -- we will use this card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
}

func check_it_for_needing_more_practice() {
	var skip_this_step bool
	skip_this_step = false
	for s, f := range frequencyMapOf_need_workOn {
		if s == aCard.KeyR { // Check if the latest random card is in the need_workOn map, and check the freq ...
			if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we keep it
				//read_map_of_fineOn() // we show the map
				fmt.Printf("\n The Random card: %s was missed once or more \n", aCard.KeyH)
				fmt.Println("... so we will keep it and quiz you on it ... ")
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
				_, err2 := fmt.Fprintf(fileHandleBig, "\nPart1of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
					s, f)
				check(err2)
				skip_this_step = true
				break //  ... we exit the loop and the func -- we will keep and use this random card, and skip the next loop
				//check_it_for_fine_on() // ... and we check THAT new card with a recursive call
			} else { // else the card had a freq less than 2, so ...
				continue // keep looking through the map for another instance that may in there, with a significant freq
			}
		}
	}
	if skip_this_step == false {
		// The latest random was not in the map, but it is time to serve-up something difficult ... so:
		for s, f := range frequencyMapOf_need_workOn {
			if s == aCard.KeyR { // Check if the latest random is in the map, and check the freq ...
				if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we set it as aCard
					//read_map_of_fineOn() // we show the map
					fmt.Println("\n This Random card was missed 1 or more times ")
					fmt.Println("... so we will test you on it, since it has been a while")
					// Log to a file that this action was taken **do-this**
					fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                 // ... gets a file handle to JapLog.txt
					//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
					_, err2 := fmt.Fprintf(fileHandleBig, "\nPart2of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
						s, f)
					check(err2)
					practice_this_card(aCard.KeyR) // locate and assign aCard // set it as new aCard
					break                          //  ... we exit the loop and the func -- we will keep and use this random card
					//check_it_for_fine_on() // ... and we check THAT new card with a recursive call
				} else { // else the card had a freq less than 2, so ...
					continue // keep looking through the map for another instance that may in there, with a significant freq
				}
			}
		}
	}
}

func check_it_for_needing_more_practiceD() {
	var skip_this_step bool
	skip_this_step = false
	for s, f := range frequencyMapOf_need_workOn {
		if s == aCardD.KeyR { // Check if the latest random card is in the need_workOn map, and check the freq ...
			if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we keep it
				//read_map_of_fineOn() // we show the map
				fmt.Printf("\n The Random card: %s was missed once or more \n", aCardD.KeyH)
				fmt.Println("... so we will keep it and quiz you on it ... ")
				/*
					// Log to a file that this action was taken **do-this**
					fileHandleBig, err := os.OpenFile("JapLogD.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                                  // ... gets a file handle to JapLog.txt
					//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
					_, err2 := fmt.Fprintf(fileHandleBig, "\nPart1of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
						s, f)
					check(err2)

				*/
				skip_this_step = true
				break //  ... we exit the loop and the func -- we will keep and use this random card, and skip the next loop
				//check_it_for_fine_on() // ... and we check THAT new card with a recursive call
			} else { // else the card had a freq less than 2, so ...
				continue // keep looking through the map for another instance that may in there, with a significant freq
			}
		}
	}
	if skip_this_step == false {
		// The latest random was not in the map, but it is time to serve-up something difficult ... so:
		for s, f := range frequencyMapOf_need_workOn {
			if s == aCardD.KeyR { // Check if the latest random is in the map, and check the freq ...
				if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we set it as aCard
					//read_map_of_fineOn() // we show the map
					fmt.Println("\n This Random card was missed 1 or more times ")
					fmt.Println("... so we will test you on it, since it has been a while")
					/*
						// Log to a file that this action was taken **do-this**
						fileHandleBig, err := os.OpenFile("JapLogD.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
						check(err)                                                                                  // ... gets a file handle to JapLog.txt
						//defer fileHandleBig.Close() // It’s idiomatic to defer a Close immediately after opening a file.
						_, err2 := fmt.Fprintf(fileHandleBig, "\nPart2of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
							s, f)
						check(err2)

					*/
					practice_this_cardD(aCardD.KeyR) // locate and assign aCard // set it as new aCard
					break                            //  ... we exit the loop and the func -- we will keep and use this random card
					//check_it_for_fine_on() // ... and we check THAT new card with a recursive call
				} else { // else the card had a freq less than 2, so ...
					continue // keep looking through the map for another instance that may in there, with a significant freq
				}
			}
		}
	}
}

func stack_the_map() {
	promptToSkip := "shi"
	for i := 0; i < 6; i++ {
		frequencyMapOf_IsFineOnChars[promptToSkip]++
	}
}
