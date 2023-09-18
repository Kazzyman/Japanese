package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			case "menu": // From func branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExercise string)
				// Flush the old stats and hits arrays
				cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
				cyclicArrayHits = CyclicArrayHits{}
				usersGuessOrOptionDirective = "null"
				do_betweenMainMenuSelectionsTTE(selectedExercise)
				mainMenuPromptScanSelectAndBeginSelectedExercise()
		*/
		// Flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// finalExercise = "drillLines"
		// usersGuessOrOptionDirective = "null"
		do_betweenMainMenuSelectionsTTE("drillLines")
		mainMenuPromptScanSelectAndBeginSelectedExercise()
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
			case "menu": // From func branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExercise string)
				// Flush the old stats and hits arrays
				cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
				cyclicArrayHits = CyclicArrayHits{}
				usersGuessOrOptionDirective = "null"
				do_betweenMainMenuSelectionsTTE(selectedExercise)
				mainMenuPromptScanSelectAndBeginSelectedExercise()
		*/
		// Flush the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		do_betweenMainMenuSelectionsTTE("drillLines")
		mainMenuPromptScanSelectAndBeginSelectedExercise()
	}
}
