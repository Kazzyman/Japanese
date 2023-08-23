package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete

import "fmt"

// used only in handle_singleQuestMark_contextSensitive_directive() Question
// this func calls no other functions and sets no variables
func giveHintInResponseToSingleQuestionMarkDir(key, value, Hint1k, Hint2k, Hint1h, Hint2h string) {
	// **do-this** decided that only this func is allowed to provide hints, except MAYBE in special cases such as in ?? H-M-B
	// 'in' is always '?' when this func is called
	//
	// proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handlers
	fmt.Printf("%s", colorReset) // hope this will apply to all that follow ????
	if key == "jiジ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
		// さ　し　 す　せ　そ
		// sa shi su se so
		// za ji  zu ze zo゛
	}
	// proper use of the ji sound from the ta group : ta->da,ji , a special key handler
	if key == "jiヂ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
		// た　ち   つ　て　と
		// ta chi tsu te to
		// da ji  zu  de do゛
	}
	// proper use of the gi sound from the ka group : ka->ga , a special key handler
	if key == "giギ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
		// か　き　く　け　こ
		// ka ki ku ke ko
		// ga gi gu ge go゛
	}
	// properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
	// , special key handlers
	if key == "jaジャ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	if key == "juジュ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	if key == "joジョ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	// obsolete forms of ja, ju and jo , special key handlers
	if key == "ja obsヂャ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	if key == "ju obsヂュ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	if key == "jo obsヂョ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	//fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	//ContextSHskip = true   // will amend all above ??? // will amend all above ??? // AND below

	//
	// then we check one very-special prompt(key)/value events or conditions,  , a very-special key handler
	if key == "zu" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if key == "zuズ" || key == "zuヅ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "あ" { // hints: may want to add additional special hints as special cases above
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "い" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "う" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "え" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "お" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 5

	if value == "ず" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "づ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "か" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "き" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "く" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "け" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "こ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 10

	if value == "さ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "し" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "す" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "せ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "そ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 15

	if value == "た" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ち" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "つ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "て" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "と" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 20

	if value == "な" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "に" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ぬ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ね" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "の" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 25

	if value == "ま" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "み" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "む" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "め" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "も" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 30

	if value == "ら" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "り" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "る" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "れ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ろ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 35

	if value == "ほ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	}
	if value == "ひ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ふ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "へ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ほ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 40

	if value == "や" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "ゆ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "よ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 43

	if value == "わ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else if value == "を" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else
	// 45

	if value == "ん" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1k, Hint2k, Hint1h, Hint2h)
	} else {
		fmt.Printf("\n")
	}
	// 46

	fmt.Println()
}
