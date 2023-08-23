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

// Ask the user for a Romaji|Hira response to the one (unique-and-naked) Katakana prompt:
//
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

// DIRECTIVES : --------------------------------------------------------------------------------------------
//goland:noinspection ALL  **do-this**
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
func handle_singleQuestMark_contextSensitive_directive() {
	giveHintInResponseToSingleQuestionMarkDir(aCard.KeyH, aCard.Value, aCard.Hint1k, aCard.Hint2k, aCard.Hint1h, aCard.Hint2h) // dole-out a hint
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
	obtainAndDealWithUserInput(firstSlice)

	fmt.Println("  1  2   3  4   5  　6  7   8   9   10  11  12  13  14  15  16  17  18  19  20  21")
	fmt.Println("は　ひ　ふ　へ　ほ　ま　み　む　め　も　や　ゆ　よ　ら　り　る　れ　ろ　わ　を　ん")
	obtainAndDealWithUserInput(secondSlice)

	fmt.Println("  1   2   3  4   5   6    7   8  9   10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25 ")
	fmt.Println("ga  gi  gu  ge  go,  za  ji  zu  ze  zo, da  ji  zu  de  do, ba  bi  bu  be  bo, pa  pi  pu  pe  po")
	obtainAndDealWithUserInput(thirdSlice)

	fmt.Println(" 1     2      3    4     5   6    7    8    9   　10    11   12   13   14   15   16   17   18   19   20  21   22   23   24")
	fmt.Println("gya  　gyu 　 gyo  ja  　ju  jo  　ja  ju 　 jo  bya 　byu  byo  pya  pyu  pyo  kya  kyu  kyo  sha  shu  sho  cha  chu  cho")
	obtainAndDealWithUserInput(fourthSlice)

	fmt.Println("1      2    3      4   　  5     6     7    8      9    10    11    12")
	fmt.Println("nya  　nyu  nyo 　 hya 　 hyu  hyo 　 mya  myu　  myo  rya 　 ryu  ryo")
	obtainAndDealWithUserInput(fifthSlice)

	var pause string
	_, err := fmt.Scan(&pause)
	if err != nil {
		return
	}
	fmt.Println(pause)
}

// Used 5 times, above in: func drillLines()
func obtainAndDealWithUserInput(refSlice []string) {
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
}
