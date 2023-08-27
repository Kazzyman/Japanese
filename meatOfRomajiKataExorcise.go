package main

// **do-this** : is the tag I have used to denote sections that I have yet to complete

import "fmt"

func meatOfRomajiKataExorcise(in string) {
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	/*  **do-this**
	??? Retain the first sections with the "^^Right!" & "^^Oops!" messages and actions, and maybe copy the rest to G-A-H-H ???
		This func provides some special "^^Right!" or "^^Oops!" messages and actions in a "'fall'-through" manner ...
		... and, THAT is ALL that it does! (SOME hints ARE also appended in the case of either a "^^Right!" or an "^^Oops!")
		... It only calls one func: checkForHints(value)
		-----------------------------------------------------------------
	*/

	// Proper use of the "ji" sound from the "sa" group (sometimes spelled zi) : sa->za,ji , special 'key' handlers
	// さ　し　 す　せ　そ
	// sa shi su se so
	// za ji  zu ze zo゛
	if aCard.KeyRK == "jiジ" {
		if in == "じ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("ji:ジ is always from shi し and NEVER from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(": ジ is the sound ji always from shi し ,and NEVER from the other ji: chi\n")
		}
	}
	// Proper use of the ji sound from the "ta" group : ta->da,ji , another special 'key' handler
	// た　ち   つ　て　と
	// ta chi tsu te to
	// da ji  zu  de do゛
	if aCard.KeyRK == "jiヂ" {
		if in == "ぢ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf(" ji:ヂ is always from chi ち ,it is the sound ji, NEVER gi ( that being ぎ:gi ギ )\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(": the ヂ sound ji is always from either chi ち or shi し\n")
		}
	}

	// Proper use of the "gi" sound from the ka group : ka->ga , another special 'key' handler
	// か　き　く　け　こ
	// ka ki ku ke ko
	// ga gi gu ge go゛
	if aCard.KeyRK == "giギ" {
		if in == "ぎ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is always from ki き and NEVER from shi し ,and it is the sound gi, NEVER ji ( that being じ:ji ジ )\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(":the sound gi is always from ki き ,and NEVER from shi し or chi ち\n")
		}
	}

	// Properly-constructed forms of ji used as ja, ju, jo which are nearly always from shi し and really never from chi ち
	// , three-more special 'key' handlers
	if aCard.KeyRK == "jaジャ" { // 1 of 3
		if in == "じゃ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and really never from chi ち\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(": it is nearly-never from chi ち but is nearly-always from shi し\n")
		} // 1 of 3
	}
	if aCard.KeyRK == "juジュ" { // 2 of 3
		if in == "じゅ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and really never from chi ち\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(": it is nearly-never from chi ち but is nearly-always from shi し\n")
		} // 2 of 3
	}
	if aCard.KeyRK == "joジョ" { // 3 of 3
		if in == "じょ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and really never from chi ち\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(": it is nearly-never from chi:ち but is nearly-always from shi:し\n")
		}
	} // end 3 of 3

	// /* obsolete */ forms of ja, ju and jo , special 'key' handlers for rarely-used forms
	if aCard.KeyRK == "ja obsヂャ" { // 1 of 3
		if in == "ぢゃ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("The ji sound is nearly-always from shi:し and really never from chi:ち\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(": it is nearly-never from chi:ち but is nearly-always from shi:し\n")
		} // 1 of 3
	}
	if aCard.KeyRK == "ju obsヂュ" { // 2 of 3
		if in == "ぢゅ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("The ji sound is nearly-always from shi し and really never from chi ち\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf(" And remember that it is nearly-never from chi ち but is nearly-always from shi し\n")
		} // 2 of 3
	}
	if aCard.KeyRK == "jo obsヂョ" { // 3 of 3
		if in == "ぢょ" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It is nearly always from shi し and really never from chi ち\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH)
			fmt.Printf("%s", colorReset)
			fmt.Printf("　And remember that it is nearly-never from chi ち but is nearly-always from shi し\n")
		}
	} // end 3 of 3  ---  'obsolete' :: ya, yu, yo forms of ja, ju, jo "obsolete"
	// */  -- 'fall' through --- "fall-through" --- >>>

	// Next, we check for some additional 'very-special' prompt(key)/value events or conditions ...
	// ... one 'very'-special key handler to emphasize the sameness of the two variants of zu sound:
	if aCard.KeyRK == "zuズ" || aCard.KeyRK == "zuヅ" || aCard.KeyRK == "zuズ from つ or す" || aCard.KeyRK == "zuヅ from つ or す" {
		//      ^^     ^^                ^^     ^^  are in main, but ^^     ^^^^^^^^^^^^^      and      ^^       ^^^^^^^^^^^^^ are NOT!!!!
		// ^^^  **do-this** ^^^ add these to main.go var fileOfCards = []charSetStruct{
		// "zu: from つ or す" IS in thar one time as a KeyR
		if in == "づ" || in == "ず" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			//
			logHitsRomajiKata("Right")
			logSkipThisPrompt(aCard.KeyR)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound\n")
		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("        ^^Oops! ")
			//
			logHitsRomajiKata("Oops")
			logReinforceThisPrompt(aCard.KeyR)
			logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
			//
			fmt.Printf("%s", colorReset)
			fmt.Printf("It was: ")
			fmt.Printf("%s", colorCyan)
			fmt.Printf("%s ", aCard.KeyH) // v v v v v right up here v v v v v v
			fmt.Printf("%s", colorReset)  // ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
			fmt.Printf("It could have been either ず or づ as they are the same sound\n")
		}
	} else {
		// This 'else' (immediately above) covers a lot of if, else ifs: Dozens of them!
		// ...
		// the next two conditions are for all remaining normal (not special) prompt(key)/value events or conditions
		// ...

		if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
			// ********************************************************
			// 'else', no special cases were found, so we process normal cases of "if in == aCard.KeyH"
			// ********************************************************
			if in == aCard.KeyH {
				fmt.Printf("%s", colorGreen)
				fmt.Printf("        ^^Right! \n")
				fmt.Printf("%s", colorReset)
				//
				logHitsRomajiKata("Right")
				logSkipThisPrompt(aCard.KeyR)
				//
			} else {
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　 ^^Oops! ")
				//
				logHitsRomajiKata("Oops")
				logReinforceThisPrompt(aCard.KeyR)
				logRomajiKataGottenWrong(aCard.KeyRK + ":it was:" + aCard.KeyH + ":but you had guessed:" + in)
				//
				fmt.Printf("%s", colorReset)
				fmt.Printf(" It was: ")
				fmt.Printf("%s", colorCyan)
				fmt.Printf("%s", aCard.KeyH)
				fmt.Printf("%s", colorReset)
				// added new func here v v v  **do-this**
				checkForHints(aCard.KeyH) // this new func is found in meatOfNakedKata.go : we only give hints for non-composites **do-this**
			}
		}
	}
}

/*
	if value == "あ" { // a ア, maybe a grotesque A
		fmt.Printf(", hint: L-middle<- to the '3' char, looks nothing-like the hiragana 'a', but a lot like a te ア, あ Fuck mae! ")
	} else if value == "い" { // i イ, shift the two lines of the hiragana
		fmt.Printf(", hint: L-middle< to the 'E' char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines ")
	} else if value == "う" { // u ウ, um-kay
		fmt.Printf(", hint: L-middle>^ to the '4' char, u ウ　is ok, they had to look for angles ")
	} else if value == "え" { // e エ, eye see it as a ... an, eye
		fmt.Printf(", hint: L-index>^^ to the '5' char, it エ　does have a vague, angular resemblance  ")
	} else if value == "お" { // o オ, on-the-go maybe
		fmt.Printf(", hint: L-index--> to the '6' char, オ　has only a vague resemblance, albeit with less curves ")
	} else if value == "か" { // ka カ
		fmt.Printf(", hint: L-index-->^ to the 'T' char ")
	} else if value == "き" { // ki キ
		fmt.Printf(", hint: L-index> to the 'G' char ")
	} else if value == "く" { // ku ク, compare to ta タ, and ke ケ
		fmt.Printf(", hint: R-index<- to the 'H' char, ク　no, just no. Starting with one angle, they settled for this? ")
	} else if value == "け" { // ke ケ, compare to ku ク, and ta タ
		fmt.Printf(", hint: R-pinky-> one over to the ':*' chars, ケ, bits of it are there, just as many curves though ")
	} else if value == "こ" { // ko コ, compare to ni ニ
		fmt.Printf(", hint: R-index<--- to the 'B' char, alright, コ　it makes sense, 'cause angles ")
	} else if value == "さ" { // sa サ, if you 'se' so sa
		fmt.Printf(", hint: L-ring>v to the 'X' char, at least the sa goes to the left, and it looks a lot like se, I'd say ")
	} else if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ,
		fmt.Printf(", hint: L-middle on the 'D' char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) ")
	} else if value == "す" { // su ス, sue is running
		fmt.Printf(", hint: L-index< to the 'R' char, they were looking for angles (sue is running ス) ")
	} else if value == "せ" { // se セ
		fmt.Printf(", hint: R-pinky to the 'P' char ")
	} else if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick
		fmt.Printf(", hint: L-index<-- to the 'C' char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  ")
	} else if value == "た" { // ta タ, compare to ku ク and ke ケ
		fmt.Printf(", hint: L-pinky < to the 'Q' char, the top left looks like ta, more at least than ku ク, or he ケ ")
	} else if value == "ち" { // chi チ, compare to te テ
		fmt.Printf(", hint: L-pinky on the 'A' char, 'some' resemblance if we see the middle line as the top of the backwards c ")
	} else if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly
		fmt.Printf(", hint: L-pinky to the 'Z' char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said ")
	} else if value == "て" { // te テ, compare to chi チ
		fmt.Printf(", hint: L-ring< to the 'W' char, te went the wrong way, and gained a flat hat ")
	} else if value == "と" { // to ト
		fmt.Printf(", hint: L-ring on the 'S' char, toe is flipped-out. Kicked in the balls, on its head ")
	} else if value == "な" { // na ナ, compare to me メ, and nu ヌ
		fmt.Printf(", hint: R-index< to the 'U' char, is simple, very simple ")
	} else if value == "に" { // ni ニ
		fmt.Printf(", hint: R-middle < to the 'I' char ")
	} else if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
		fmt.Printf(", hint: L-ring<--- to the '1' char, at least two lines cross, nu flew  ")
	} else if value == "ね" { // ne ネ
		fmt.Printf(", hint: L-ring<v to the ',<' chars, a hoe that got nailed ")

		// v v v v these all execute with an Oops ^^^^
		// just need to add all the following to the struct **do this**
	} else if value == "の" { // no ノ, no extra marks needed on this classic no symbol
		fmt.Printf(", hint: R-middle on the 'K' char, drop the curve, and I'll take it ")
	} else if value == "ま" { // ma マ, mama is unique, see the breast pump?
		fmt.Printf(", hint: R-index on the 'J' char, (a breast pump maybe?) ")
	} else if value == "み" { // mi ミ, and two makes three
		fmt.Printf(", hint: R-index<v to the 'N' char, me thinks 3 ")
	} else if value == "む" { // mu ム, an even-fatter-jawed moo-cow
		fmt.Printf(", hint: R-pinky two over---> to the ']}' chars, now it's an even-fatter-jawed moo-cow  ")
	} else if value == "め" { // me メ
		fmt.Printf(", hint: L-pinky->v slide down to the '/?' chars, cross off the mess. But that's nothing nu ヌ ")
	} else if value == "も" { // mo モ
		fmt.Printf(", hint: R-middle<v to the 'M' char, way to hang to the right; you go Joe! ")
	} else if value == "ら" { // ra ラ
		fmt.Printf(", hint: R-ring^ to the 'o' char, Similar ラ, ら ")
	} else if value == "り" { // ri リ
		fmt.Printf(", hint: R-ring on the 'L' char, longer on the right ring ")
	} else if value == "る" { // ru ル, is two
		fmt.Printf(", hint: R-pinky<v to the '.>' chars, ル, る, ru is now two? ")
	} else if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
		fmt.Printf(", hint: R-pinky on the ';+' chars, レ, れ; しshe-it ")
	} else if value == "ろ" { // ro ロ
		fmt.Printf(", hint: R-pinky ---> long slide to the '_' and/or '\\' char")
		fmt.Printf("... maybe it had always been the way-to-go to do a square 'O' for ro? ")
	} else if value == "は" { // ha ハ
		fmt.Printf(", hint: L-index on the 'F' char ")
	} else if value == "ひ" { // hi ヒ
		fmt.Printf(", hint: L-index>v to the 'V' char ")
	} else if value == "ふ" { // hu フ, squinting it is a ふフ
		fmt.Printf(", hint: L-ring<-^ to the '2' char, if we squint? フ, ふ ")
	} else if value == "へ" { // he ヘ
		fmt.Printf(", hint: R-ring---->^^ to the '~' char ")
	} else if value == "ほ" { // ho ホ, now with a dress and all
		fmt.Printf(", hint: R-ring--->^ to the '-' char, now looks like a hoe ")
	} else if value == "や" { // ya ヤ
		fmt.Printf(", hint: R-index<--^^ to the '7' char ")
	} else if value == "ゆ" { // yu ユ
		fmt.Printf(", hint: R-index->^^, to the '8' char, yu looks like ユ ")
	} else if value == "よ" { // yo ヨ
		fmt.Printf(", hint: R-middle^^ to the '9' char, triple yo ヨ ")
	} else if value == "わ" { // wa ワ, compare to wo ヲ
		fmt.Printf(", hint: R-ring>^^ to the '0' char, so now it's a water fall for wa? ")
	} else if value == "を" { // wo ヲ, compare to wa ワ
		fmt.Printf(", hint: R-ring>^^ shifted '0' char, NOW, it looks like wa. Now! For fuck's sake, Now? ")
	} else if value == "ん" { // nh ン, compare to so ソ
		fmt.Printf(", hint: 'Y' char, pointing at the one last remaining bent stroke of the hiragana char ")
	} else {
	}
}
}
fmt.Println() // <-- adds formatting space after a ^^Right!
} */
//
