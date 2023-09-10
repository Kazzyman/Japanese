package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete
import "fmt"

/*
	Hint1h    string
	Hint2k    string
	Hint3TT   string
	HintSansR string
*/
// Used only in handle_singleQuestMark_contextSensitive_directive()
// This func calls no other functions and sets no variables
// (key, Hint1, Hint2, Hint3 string) is intended to allow for polymorphism re specific hint lines, yet I now think not!
func giveHintInResponseToSingleQuestionMarkContextSensitiveDir(key, Hint1, Hint2, Hint3 string) {
	// As: giveHintInResponse...^^^...ContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
	// **do-this** decided that only this func is allowed to provide hints, at least for '?' Directive
	//
	// Proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handlers
	fmt.Printf(colorReset)
	//
	//
	// The conditionals which follow all need to be just as they are: as 'if else if' constructs - to handel the last fmt.Println("")
	//
	// First, we do all the forms with dakutens or han-dakutens
	if key == "が" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぎ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぐ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "げ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ご" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	//
	if key == "ざ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "じ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ず" { // duplicate ??? **do-this**
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぜ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぞ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	//

	if key == "だ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぢ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "づ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "で" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ど" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	//

	if key == "ば" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "び" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぶ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "べ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぼ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぱ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぴ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぷ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぺ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぽ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	//
	/*
		.
		.
		.
	*/
	// Next we do the simple cases of no dakutens or han-dakutens
	if key == "あ" { // hints: may want to add additional special hints as special cases above
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "い" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "う" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "え" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "お" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 5

	if key == "か" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "き" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "く" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "け" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "こ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 10

	if key == "さ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "し" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "す" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "せ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "そ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 15

	if key == "た" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ち" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "つ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "て" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "と" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 20

	if key == "な" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "に" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ぬ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ね" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "の" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 25

	if key == "ま" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "み" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "む" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "め" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "も" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 30

	if key == "ら" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "り" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "る" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "れ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ろ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 35

	if key == "は" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ひ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ふ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "へ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ほ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 40

	if key == "や" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "ゆ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "よ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 43

	if key == "わ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "を" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else
	// 45

	if key == "ん" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else {
		fmt.Printf("\n")
	}
	// 46

	fmt.Println()
}

/*
.
.
.
.
*/
// Below, key is aCard. KeyH; and sansRomaji_hint is aCard. HintSansR
func giveHintInResponseToSingleQuestionMarkContextSensitiveDir_sansRomaji(key, sansRomaji_hint string) {
	// **do-this** decided that only this func is allowed to provide hints, at least for '?' Directive
	//
	// Proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handlers
	fmt.Printf(colorReset) // hope this will apply to all that follow ????
	//
	//
	// The conditionals which follow all need to be just as they are: as 'if else if' constructs - to handel the last fmt.Println("")
	//
	// First, we do all of the forms with dakutens or han-dakutens
	if key == "が" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぎ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぐ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "げ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ご" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	//
	if key == "ざ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "じ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ず" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぜ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぞ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	//

	if key == "だ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぢ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "づ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "で" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ど" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	//

	if key == "ば" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "び" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぶ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "べ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぼ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぱ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぴ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぷ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぺ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぽ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	//
	/*
		.
		.
		.
	*/
	// Next we do the simple cases of NO dakutens OR han-dakutens
	if key == "あ" { // hints: may want to add additional special hints as special cases above
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "い" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "う" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "え" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "お" {
		fmt.Printf("%s", sansRomaji_hint)
		fmt.Printf("%s", sansRomaji_hint) // **do-this**
	} else
	// 5

	if key == "か" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "き" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "く" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "け" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "こ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 10

	if key == "さ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "し" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "す" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "せ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "そ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 15

	if key == "た" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ち" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "つ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "て" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "と" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 20

	if key == "な" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "に" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ぬ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ね" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "の" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 25

	if key == "ま" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "み" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "む" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "め" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "も" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 30

	if key == "ら" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "り" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "る" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "れ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ろ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 35

	if key == "は" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ひ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ふ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "へ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ほ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 40

	if key == "や" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "ゆ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "よ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 43

	if key == "わ" {
		fmt.Printf("%s", sansRomaji_hint)
	} else if key == "を" {
		fmt.Printf("%s", sansRomaji_hint)
	} else
	// 45

	if key == "ん" {
		fmt.Printf("%s", sansRomaji_hint)
	} else {
		fmt.Printf("\n")
	}
	// 46

	fmt.Println()
}
