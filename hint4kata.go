package main

import "fmt"

func giveAhintKata(value string) {

	// fix all these hints so that they only refer to the value one needs to type that corresponds to the katakana prompt (to the key)

	if value == "あ" { // a ア, maybe a grotesque A
		fmt.Printf(", hint: middle<- to the 3 char, looks nothing-like the hiragana a, but a lot like a te ア, あ, fuck mae! \n")
	} else if value == "い" { // i イ, shift the two lines of the hiragana
		fmt.Printf(", hint: middle < to the E char, イ looks more like a hiragana te, but at least it is still two mostly-vertical lines イ　\n")
	} else if value == "う" { // u ウ, um-kay
		fmt.Printf(", hint: middle> to the 4 char, u う ウ　is ok, having had to look for angles \n")
	} else if value == "え" { // e エ, eye see it as a ... an, eye
		fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
	} else if value == "お" { // o オ, on-the-go maybe
		fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
	} else if value == "か" { // ka カ
		fmt.Printf(", hint: index --> to the T char \n")
	} else if value == "き" { // ki キ
		fmt.Printf(", hint: L-index > to the G char \n")
	} else if value == "く" { // ku ク, compare to ta タ, and ke ケ
		fmt.Printf(", hint: R-index <- to the H char, ク　no, just no. Starting with one angle, they settled for this? \n")
	} else if value == "け" { // ke ケ, compare to ku ク, and ta タ
		fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits of it are there, just as many curves though \n")
	} else if value == "こ" { // ko コ, compare to ni ニ
		fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles \n")
	} else if value == "さ" { // sa サ, if you 'se' so sa
		fmt.Printf(", hint: ring >v to the X char, at least the sa goes to the left, and it looks a lot like se, I'd say \n")
	} else if value == "し" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ,
		fmt.Printf(", hint: middle on the D char, シ she is sleeping/snoring here, but has no excuse to2 look like this (no angles here, less curve though) \n")
	} else if value == "す" { // su ス, sue is running
		fmt.Printf(", hint: index < to the R char, they were looking for angles (sue is running ス) \n")
	} else if value == "せ" { // se セ
		fmt.Printf(", hint: pinky to the P char \n")
	} else if value == "そ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick
		fmt.Printf(", hint: index <-- to the C char, ソ Looks like しシ or no ノ, or ン starting with all angles - `had to `backtrack?  \n")
	} else if value == "た" { // ta タ, compare to ku ク and ke ケ
		fmt.Printf(", hint: pinky < to the Q char, the top left looks like ta, more at least than ku ク, or he ケ \n")
	} else if value == "ち" { // chi チ, compare to te テ
		fmt.Printf(", hint: pinky on the A char, 'some' resemblance if we see the middle line as the top of the backwards c \n")
	} else if value == "つ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly
		fmt.Printf(", hint: pinky to the Z char, understandable?, if we are generous. Very generous..゛as in adding two ツ  shi シ said \n")
	} else if value == "て" { // te テ, compare to chi チ
		fmt.Printf(", hint: ring < to the W char, te went the wrong way, and gained a flat hat \n")
	} else if value == "と" { // to ト
		fmt.Printf(", hint: ring on the S char, toe is flipped-out. Kicked in the balls, on its head \n")
	} else if value == "な" { // na ナ, compare to me メ, and nu ヌ
		fmt.Printf(", hint: index < to the U char, is simple, very simple \n")
	} else if value == "に" { // ni ニ
		fmt.Printf(", hint: middle < to the I char \n")
	} else if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
		fmt.Printf(", hint: R-ring<--- to the 1 char, at least two lines cross, nu flew  \n")
	} else if value == "ね" { // ne ネ
		fmt.Printf(", hint: L-ring <v to the ,< chars, a hoe that got nailed \n")
	} else if value == "の" { // no ノ, no extra marks needed on this classic no symbol
		fmt.Printf(", hint: middle on the k char, drop the curve, and I'll take it \n")
	} else if value == "ま" { // ma マ, mama is unique, see the breast pump?
		fmt.Printf(", hint: index on the J char, (a breast pump maybe?) \n")
	} else if value == "み" { // mi ミ, and two makes three
		fmt.Printf(", hint: index <v to the N char, me thinks 3 \n")
	} else if value == "む" { // mu ム, an even-fatter-jawed moo-cow
		fmt.Printf(", hint: pinky two over ---> to the ]} chars, now it's an even-fatter-jawed moo-cow  \n")
	} else if value == "め" { // me メ
		fmt.Printf(", hint: pinky -> slide down to the /? chars, cross off the mess. But that's nothing nu ヌ \n")
	} else if value == "も" { // mo モ
		fmt.Printf(", hint: middle <v to the M char, way to hang to the right; you go Joe! \n")
	} else if value == "ら" { // ra ラ
		fmt.Printf(", hint: ring ^ to the o char, Similar ラ, ら \n")
	} else if value == "り" { // ri リ
		fmt.Printf(", hint: ring on the L char, longer on the right ring \n")
	} else if value == "る" { // ru ル, is two
		fmt.Printf(", hint: pinky < to the .> chars, ル, る, ru is now two? \n")
	} else if value == "れ" { // re レ, more, or less, a ru ル - しshe-said
		fmt.Printf(", hint: pinky on the ;+ chars, レ, れ; しshe-it \n")
	} else if value == "ろ" { // ro ロ
		fmt.Printf(", hint: pinky ---> long slide to the _ and backslash chars, maybe it had always been the way to go to do a square O for ro? \n")
	} else if value == "ほ" { // ha ハ
		fmt.Printf(", hint: L-ring---> to the \"-\" char\n")
	} else if value == "ひ" { // hi ヒ
		fmt.Printf(", hint: index > to the V char \n")
	} else if value == "ふ" { // hu フ, squinting it is a ふフ
		fmt.Printf(", hint: ring<, to the 2 char, if we squint? フ, ふ\n")
	} else if value == "へ" { // he ヘ
		fmt.Printf(", hint: ;+ char \n")
	} else if value == "ほ" { // ho ホ, now with a dress and all
		fmt.Printf(", hint: ring ---> to the - char, now looks like a hoe \n")
	} else if value == "や" { // ya ヤ
		fmt.Printf(", hint: index <-- to the 7 char \n")
	} else if value == "ゆ" { // yu ユ
		fmt.Printf(", hint: index->, to the 8 char, yu looks like ユ \n")
	} else if value == "よ" { // yo ヨ
		fmt.Printf(", hint: middle to the 9 char, tripple yo ヨ \n")
	} else if value == "わ" { // wa ワ, compare to wo ヲ
		fmt.Printf(", hint: ring > to the 0 char, so now it's a water fall for wa? \n")
	} else if value == "を" { // wo ヲ, compare to wa ワ
		fmt.Printf(", hint: ring > shifted 0 char, NOW, it looks like wa. Now! For fucks sake, Now? \n")
	} else if value == "ん" { // nh ン, compare to so ソ
		fmt.Printf(", hint: Y char, pointing at the one last remaining bent stroke of the hiragana char \n")
	} else {
		fmt.Printf("\n")
	}
}
