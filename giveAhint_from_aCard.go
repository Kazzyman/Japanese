package main

// **do-this** 2 : is the tag I have used to denote sections that I have yet to complete

import "fmt"

// This func provides all four hint fields **do-this** make a version that only reports|provides one or two
// This func calls no other functions and sets no variables **do-this** maybe make use of it
func giveAhintHira(key, value, Hint1h, Hint2k, Hint3TT, HintSansR string) {
	// Decided that this func ALONE is allowed to provide hints from a card (aCard)
	//
	// proper use of the ji sound from the sa group (sometimes spelled zi) : sa->za,ji , special key handlers
	fmt.Printf("%s", colorReset)
	if key == "jiジ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
		// さ　し　 す　せ　そ
		// sa shi su se so
		// za ji  zu ze zo゛
	}
	// proper use of the ji sound from the ta group : ta->da,ji , a special key handler
	if key == "jiヂ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
		// た　ち   つ　て　と
		// ta chi tsu te to
		// da ji  zu  de do゛
	}
	// proper use of the gi sound from the ka group : ka->ga , a special key handler
	if key == "giギ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
		// か　き　く　け　こ
		// ka ki ku ke ko
		// ga gi gu ge go゛
	}
	// properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
	// , special key handlers
	if key == "jaジャ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	if key == "juジュ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	if key == "joジョ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	// obsolete forms of ja, ju and jo , special key handlers
	if key == "ja obsヂャ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	if key == "ju obsヂュ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	if key == "jo obsヂョ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	//fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	//ContextSHskip = true   // will amend all above ??? // will amend all above ??? // AND below

	//
	// then we check one very-special prompt(key)/value events or conditions,  , a very-special key handler
	if key == "zuズ" || key == "zuヅ" || key == "zuズ from つ or す" || key == "zuヅ from つ or す" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
		//ContextSHskip = true
	} else if value == "あ" { // hints: may want to add additional special hints as special cases above
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "い" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "う" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "え" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "お" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 5

	if value == "か" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "き" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "く" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "け" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "こ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 10

	if value == "さ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "し" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "す" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "せ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "そ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 15

	if value == "た" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ち" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "つ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "て" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "と" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 20

	if value == "な" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "に" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ぬ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ね" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "の" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 25

	if value == "ま" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "み" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "む" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "め" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "も" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 30

	if value == "ら" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "り" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "る" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "れ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ろ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 35

	if value == "ほ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	}
	if value == "ひ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ふ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "へ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ほ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 40

	if value == "や" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "ゆ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "よ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 43

	if value == "わ" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else if value == "を" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else
	// 45

	if value == "ん" {
		fmt.Printf("%s\n%s\n%s\n%s", Hint1h, Hint2k, Hint3TT, HintSansR)
	} else {
		fmt.Printf("\n")
	}
	// 46

	fmt.Println()
}
