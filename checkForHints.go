package main

import "fmt"

// Note: we only give hints for non-conjunctive-composites below **do-this**
func checkForHints(value string) {
	if value == "あ" { // a ア, maybe a grotesque A
		fmt.Printf(", hint: middle<- to the '3', looks nothing-like a:あ, but a lot like te:て   ア, fuck mae! \n")
	} else if value == "い" { // i イ, shift the two lines of the hiragana
		fmt.Printf(", hint: middle < to the 'E', イ looks like a te:て,  it is still two mostly-vertical lines イ　\n")
	} else if value == "う" { // u ウ, um-kay
		fmt.Printf(", hint: middle> to the 4 char, u ウ　is ok, having had to look for angles \n")
	} else if value == "え" { // e エ, eye see it as a ... an, eye
		fmt.Printf(", hint: index > to the 5 char, it エ　does have a vague, angular resemblance  \n")
	} else if value == "お" { // o オ, on-the-go maybe
		fmt.Printf(", hint: index--> to the 6 char, オ　has only a vague resemblance, albeit with less curves \n")
	} else if value == "か" || value == "が" { // ka カ
		fmt.Printf(", hint: index --> to the 'T' char,  ka-->ga  カ --> ガ   ka-ga \n")
	} else if value == "き" || value == "ぎ" { // ki キ
		fmt.Printf(", hint: L-index > to the 'G' char,  ki-->gi  キ　 ki,gi \n")
	} else if value == "く" || value == "ぐ" { // ku ク, compare to ta タ, and ke ケ
		fmt.Printf(", hint: R-index <- to the 'H', ク　no, just no, One angle & they settled for this?  ku,gu \n")
	} else if value == "け" || value == "げ" { // ke ケ, compare to ku ク, and ta タ
		fmt.Printf(", hint: pinky -> one over to the :* chars, ケ, bits are there, as many curves though  ke,ge \n")
	} else if value == "こ" || value == "ご" { // ko コ, compare to ni ニ
		fmt.Printf(", hint: index <--- to the B char, alright, コ　it makes sense, 'cause angles  ko,go \n")
	} else if value == "さ" || value == "ざ" { // sa サ, if you 'se' so sa
		fmt.Printf(", hint: ring >v to the 'X', at least sa still goes left, but looks 2-much like se  sa-za \n")
	} else if value == "し" || value == "じ" { // shi シ, she sleeps and snores, compare to tsu ツ, and so ソ,
		fmt.Printf(", hint: middle on the 'D' char, シ she is sleeping/snoring  shi,ji \n")
	} else if value == "す" || value == "ず" { // su ス, sue is running
		fmt.Printf(", hint: index < to the 'R' char, they were looking for angles (sue is running ス)  su,zu \n")
	} else if value == "せ" || value == "ぜ" { // se セ
		fmt.Printf(", hint: pinky to the 'P' char, 　sa-->ze  セ \n")
	} else if value == "そ" || value == "ぞ" { // so ソ, compare to ん ン which lays down more and has a more laid-down back-tick
		fmt.Printf(", hint: index <-- to the 'C', ソ  Compare: n:ん:ン starting w angles `backtrack  so,zo \n")
	} else if value == "た" || value == "だ" { // ta タ, compare to ku ク and ke ケ
		fmt.Printf(", hint: pinky < to the 'Q', タ top-left looks like た, more than ku:ク, or ke:ケ  ta-da \n")
	} else if value == "ち" || value == "ぢ" { // chi チ, compare to te テ
		fmt.Printf(", hint: pinky on the 'A' char, チち so 'some' resemblance?, rarely ji \n")
	} else if value == "つ" || value == "づ" { // tsu ツ, compare to shi シ, she who sleeps low and snores loudly
		fmt.Printf(", hint: pinky to the 'Z' char, see water crashing on the she shore, tsu-shi:ツシ   tsu->zu \n")
	} else if value == "て" || value == "で" { // te テ, compare to chi チ
		fmt.Printf(", hint: ring < to the 'W' char, te:テ went the wrong way, and gained a flat hat  te,de \n")
	} else if value == "と" || value == "ど" { // to ト
		fmt.Printf(", hint: ring on the 'S' char, toe:ト is flipped-out. Kicked on its head   to-do \n")
	} else if value == "な" { // na ナ, compare to me メ, and nu ヌ
		fmt.Printf(", hint: index < to the 'U', ナ:な is very simple version  Compare: me:メ,  nu:ヌ  ナ \n")
	} else if value == "に" { // ni ニ
		fmt.Printf(", hint: middle < to the 'I' char,  に:ni ニ \n")
	} else if value == "ぬ" { // nu ヌ, compare to me メ, and na ナ
		fmt.Printf(", hint: R-ring<--- to the '1' char, two lines still cross, ヌ flew, Compare: me:メ,  na:ナ \n")
	} else if value == "ね" { // ne ネ
		fmt.Printf(", hint: L-ring <v to the ',<' chars,   ネ a hoe that got nailed \n")
	} else if value == "の" { // no ノ, no extra marks needed on this classic no symbol
		fmt.Printf(", hint: middle on the 'K' char,  dropped the curve -- I'll take it:  no ノ \n")
	} else if value == "ま" { // ma マ, mama is unique, see the breast pump?
		fmt.Printf(", hint: index on the 'J' char, (a breast pump maybe?) ma マ   Compare: mu ム　\n")
	} else if value == "み" { // mi ミ, and two makes three
		fmt.Printf(", hint: index <v to the 'N' char, mi thinks, mi and 2 makes 3 lines \n")
	} else if value == "む" { // mu ム, an even-fatter-jawed moo-cow
		fmt.Printf(", hint: pinky 2-over ---> to the ']}' chars, even-fatter-jawed mu:ム  Compare: ma マ \n")
	} else if value == "め" { // me メ
		fmt.Printf(", hint: pinky -> slide to the '/?' chars, cross-off the mess. But it's not nu:ヌ me:メ \n")
	} else if value == "も" { // mo モ
		fmt.Printf(", hint: middle <v to the 'M' char, way to hang to the right; you go Joe!  mo モ \n")
	} else if value == "ら" { // ra ラ
		fmt.Printf(", hint: ring ^ to the 'O' char,  ra:ラ Compare: wa:ワ ウ:u  fu:フ wo:ヲ テ:te ラ:ra:ら \n")
	} else if value == "り" { // ri リ
		fmt.Printf(", hint: ring on the 'L' char, longer on the right ring, ri:li:りリ \n")
	} else if value == "る" { // ru ル, is two
		fmt.Printf(", hint: pinky < to the '.>' chars,  ル, る, ru is now two? \n")
	} else if value == "れ" { // re レ
		fmt.Printf(", hint: pinky on the ';+' chars,  れ:re:レ, more-or-less ru:ル sans two \n")
	} else if value == "ろ" { // ro ロ
		fmt.Printf(", hint: pinky ---> long slide to the '_\\' chars,    a square 'O' for ro ロ \n")
	} else if value == "は" || value == "ば" || value == "ぱ" { // ha ハ
		fmt.Printf(", hint: on the 'F' char,  ha ハ は looks a bit like an 'H'  ha,ba,pa \n")
	} else if value == "ひ" || value == "び" || value == "ぴ" { // hi ヒ
		fmt.Printf(", hint: index > to the 'V' char,  smiling hihi ひ or ヒ   hi,bi,pi \n")
	} else if value == "ふ" || value == "ぶ" || value == "ぷ" { // hu フ, squinting it is a ふフ
		fmt.Printf(", hint: ring<, to the '2' char, fu:hu is on 2, if we squint? フ is ふ  fu,bu,pu \n")
	} else if value == "へ" || value == "べ" || value == "ぺ" { // he ヘ
		fmt.Printf(", hint: on the '^' char,  hey-stack is on the hat char  he,be,pe \n")
	} else if value == "ほ" || value == "ぼ" || value == "ぽ" { // ho ホ, now with a dress and all
		fmt.Printf(", hint: ring ---> to the '-=' chars,  ほ now looks even more like a hoe ホ  Ho,Bo,Po \n")
	} else if value == "や" { // ya ヤ
		fmt.Printf(", hint: index <-- to the '7' char,  やヤ is ya\n")
	} else if value == "ゆ" { // yu ユ
		fmt.Printf(", hint: index->, to the '8' char, yu looks like ユ is number 1 \n")
	} else if value == "よ" { // yo ヨ
		fmt.Printf(", hint: middle to the '9' char, tripple yo ヨ \n")
	} else if value == "わ" { // wa ワ, compare to wo ヲ
		fmt.Printf(", hint: ring > to the '0' char,  it's a water fall now for wa ワ,  Compare:  wo ヲ  \n")
	} else if value == "を" { // wo ヲ, compare to wa ワ
		fmt.Printf(", hint: ring > shifted '0' char, wo:ヲ Compare: wa:ワ ウ:u:ウ フ:fu wo:ヲ テ:te ラ:ra \n")
	} else if value == "ん" { // nh ン, compare to so ソ
		fmt.Printf(", hint: 'Y' char,    nh ン is knocked-down,   Compare:  so:ソ \n")
	} else {
		fmt.Printf("\n")
	}
}
