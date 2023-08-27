package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete

import "fmt"

// Used only in handle_singleQuestMark_contextSensitive_directive()
// This func calls no other functions and sets no variables
func giveHintInResponseToSingleQuestionMarkContextSensitiveDir(key, Hint1, Hint2, Hint3 string) {
	// As: giveHintInResponse...^^^...ContextSensitiveDir(aCard.KeyH, aCard.Hint1h, aCard.Hint2k, aCard.Hint3TT) // dole-out a hint
	// **do-this** decided that only this func is allowed to provide hints
	// 'in' is always '?' when this func is called
	//
	// proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handlers
	fmt.Printf("%s", colorReset) // hope this will apply to all that follow ????
	if key == "jiジ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
		// さ　し　 す　せ　そ
		// sa shi su se so
		// za ji  zu ze zo゛
	}
	// proper use of the ji sound from the ta group : ta->da,ji , a special key handler
	if key == "jiヂ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
		// た　ち   つ　て　と
		// ta chi tsu te to
		// da ji  zu  de do゛
	}
	// proper use of the gi sound from the ka group : ka->ga , a special key handler
	if key == "giギ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
		// か　き　く　け　こ
		// ka ki ku ke ko
		// ga gi gu ge go゛
	}
	// properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
	// , special key handlers
	if key == "jaジャ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	if key == "juジュ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	if key == "joジョ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	// obsolete forms of ja, ju and jo , special key handlers
	if key == "ja obsヂャ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	if key == "ju obsヂュ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	if key == "jo obsヂョ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	//fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	//ContextSHskip = true   // will amend all above ??? // will amend all above ??? // AND below

	//
	// then we check one very-special prompt(key)/value events or conditions,  , a very-special key handler
	if key == "zu" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "zuズ" || key == "zuヅ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "あ" { // hints: may want to add additional special hints as special cases above
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

	if key == "ず" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "づ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	} else if key == "か" {
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

	if key == "ほ" {
		fmt.Printf("%s\n%s\n%s", Hint1, Hint2, Hint3)
	}
	if key == "ひ" {
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
