package main

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorCyan = "\033[36m"
var colorPurple = "\033[35m"

// Presumably, aCard will always be used to refer to a single card from fileOfCards
// e.g. aCard.KeyR where KeyR is a field in a card from fileOfCards
var aCard = charSetStruct{}

// The structure of a single 'card' (aCard) from fileOfCards
type charSetStruct struct {
	KeyK  string
	KeyH  string
	KeyR  string
	KeyRK string

	Value string

	Hint1k string
	Hint2k string
	Hint1h string
	Hint2h string
}

// Instantiate a series of struct objects as a slice of instances of the charSetStruct type
// fileOfCards being that (slice) of structures of type charSetStruct
/*
We are creating a slice named fileOfCards that holds instances of the charSetStruct type.

Each element in the slice is initialized using the composite literal syntax whereby
we are providing values for each field of the charSetStruct struct: i.e.,
each set of values enclosed in curly braces { ... } represents an instance of the struct.
*/
var fileOfCards = []charSetStruct{
	// **do-this** finish composing hint content (one specific line of each card should contain no Romaji)

	// We will probably never want to be prompted with Hiragana chars, ever! Because, that would just be too easy??? ...
	// ... It is more useful to only see a Katakana or Romaji prompt, so we would only ever want a Hiragana hint.

	// But when being prompted to enter Romaji from a Katakana prompt we would like a hint sans the KeyR ...
	// ... which is to say: just a naked aCard.KeyH or aCard.Value is what we want (ONLY a Hiragana char as hint)
	// ... that, and|or maybe just one dedicated line from the hints

	// vowels: (a-i-u-e-o) ==========================================================================================
	//
	{" ア", "あ", "a", "aア", "あ",

		" a:あ:ア ...maybe? a grotesque 'A':ア ?? , compare: me:め  to  a:あ  to  nu:ぬ    あ Fuck mae!",
		" ... , compare:  nu:ぬ  ne:ね  me:め , to a:あ",
		" The Kata a:ア looks nothing-like the hiragana 'a':あ, more like a hiragana te:て,  あ Fuck mae! ",
		" TT: middle<- to the '3' char for あ,  a:あ:ア  ?is maybe a grotesque A:ア ?? "},
	//
	//
	{" イ", "い", "i", "iイ", "い",

		" i:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: middle< to the E char  い:イ ",
		" i: ... at least it is still two mostly-vertical lines: イ "},
	//
	//
	{" ウ", "う", "u", "uウ", "う",

		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ウ more-angular and with a tick for its top line ",
		" TT: middle> to the '4' char  う : Kata looks like the wa:ワ albeit with a mohawk ' ",
		" u:ウ ,  compare: to wa:ワ the cascading wa-ter fall,　um-'ca'y "},
	//
	//
	{" エ", "え", "e", "eエ", "え",

		" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
		" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
		" TT: index> to the 5 char   e:え:エ ",
		" e e エ, eye see it as a ... an, eye ??  e:え:エ"},
	//
	//
	{" オ", "お", "o", "oオ", "お",

		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" o:お:オ, on-the-go, cloak trailing behind ... maybe? "},
	//
	// ka group: (ka-ga) ============================================================================================
	//
	{" カ", "か", "ka", "kaカ", "か",

		" ka:か:カ ... is an easy one!  It simply looks like a 'K' ",
		" ... Kata is the same カ , albeit more-angular and with one-less line to draw than か ",
		" TT: index--> to the T char   か ",
		""},
	//
	//
	{" キ", "き", "ki", "kiキ", "き",

		" ki:き:キ ... is an easy one!   キ has the same top as ki:き ... and they look like 'keys' ",
		" ki:き ... they both have the same top:  き キ",
		" TT: L-index> to the 'G' char,  Think: ki-of-G  き ,  compare to: sa:さ--ki:き  さき ... saki ",
		""},
	//
	//
	{" ク", "く", "ku", "kuク", "く",

		" ku:く:ク ,  Kata: starting with one angle, they settled for this??  く:ク ",
		" ku:く:ク ,  compare:  ta:タ  ke:ケ  ku:ク  ,  it's kuku, kookoo I tell you! ",
		" TT: R-index<- to the 'H' char  く:ク ",
		""},
	//
	//
	{" ケ", "け", "ke", "keケ", "け",

		" ke:け:ケ  bits of it are there, just as many curves though ,  Compare to:  ku:ク  ta:タ ",
		" ke:け  ... looks enough like a Keg ...  Foam-ment THAT image in your brain ",
		" TT: pinky-> one-over to the :* chars   け:ケ ",
		""},
	//
	//
	{" コ", "こ", "ko", "koコ", "こ",

		" ko:こ:コ ... it makes sense, 'cause angles:  ko:コ ,  compare:  ni:に:ニ ",
		" ... see the two Koy fish?   'ko'oy fish? ",
		" TT: index <--- to the 'B' char,  ko way-down below  こ ",
		""},
	//
	// ka becomes ga group: -----------------------------------------------------------------------------------------
	//
	{" ガ", "が", "ga", "gaガ", "が",

		" ga:が same ガ albeit more angular with one-less line to draw than が",
		" ga:が is an easy one",
		" ga index --> to the T char    が ",
		" ga"},
	//
	//
	{" ギ", "ぎ", "gi", "giギ", "ぎ",

		" gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ,",
		" ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" gi: is on the G char, ぎ:ギ, compare to za:ざ from sa:さ",
		" gi: the sound gi is always from ki:き ,and NEVER from shi:し or from chi:ち"},
	//
	//
	{" グ", "ぐ", "gu", "guグ", "ぐ",

		" gu:ぐ:グ Starting with one angle, they settled for this? Compare to ta:za:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ Compare za ダ, ge ゲ",
		" H char    ぐ:グ ",
		" gu"},
	//
	//
	{" ゲ", "げ", "ge", "geゲ", "げ",

		" ge:げ bits of it are there, just as many curves though. Compare to ku:gu:グ, and ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ",
		" :* chars    げ",
		" ge"},
	//
	//
	{" ゴ", "ご", "go", "goゴ", "ご",

		" go:ご it makes sense, 'cause angles' ",
		" go:",
		" index <--- to the B char    ご",
		" go:"},
	//
	// ya, yu, yo's of ki:き ---------------------------------------------------------------------------
	//
	{" キャ", "きゃ", "kya", "kyaキャ", "きゃ",

		" kya:きゃ is an easy one, キャ has the same top and the ya is similar",
		" kya:きゃ キャ has the same top and the ya is similar",
		" G char,    きゃ, compare to sa:さ ",
		" kya:"},
	//
	//
	{" キュ", "きゅ", "kyu", "kyuキュ", "きゅ",

		" kyu",
		" kyu",
		" kyu",
		" kyu"},
	//
	//
	{" キョ", "きょ", "kyo", "kyoキョ", "きょ",

		" kyo",
		" kyo",
		" kyo",
		" kyo"},
	//
	// ya, yu, yo's of gi:ぎ ---------------------------------------------------------------------------
	//
	{" ギャ", "ぎゃ", "gya", "gyaギャ", "ぎゃ",

		" gya",
		" gya",
		" gya",
		" gya"},
	//
	//
	{" ギュ", "ぎゅ", "gyu", "gyuギュ", "ぎゅ",

		" gyu",
		" gyu",
		" gyu",
		" gyu"},
	//
	//
	{" ギョ", "ぎょ", "gyo", "gyoギョ", "ぎょ",

		" gyo",
		" gyo",
		" gyo",
		" gyo"},
	//
	// sa group: (sa-za,ji,zu,za,zo) ================================================================================
	//
	{" サ", "さ", "sa", "saサ", "さ",

		" sa:さ:サ  sa:サ ??, if you 'se':せ そ sa ,  compare:  Hira se:せ  to  Kata sa:サ",
		"",
		" TT: ring>v to the 'X' char  --  at least Kata 'sa' still goes to the left ... ",
		" ... , but it looks an awful-lot like Hira se:せ, I'd say   ... just 'Hir'-say ?? "},
	//
	//
	{" シ", "し", "shi", "shiシ", "し",

		" shi:し:シ  'she' sleeps & snores゛,  compare: Kata:  tsu:ツ   so:ソ　to the more-laid-back shi:シ ",
		" shi:し:シ  し looks like a sheep hook. Is there even such a thing?? ,  compare:  ツシ  tsu-shi ",
		" TT: middle on the 'D' char, シ 'she' is sleeping/snoring here, but has no excuse to'2' ゛look like this (no angles here, less curve though)",
		" ... what? ... two nasal blasts ... obviously! Yes while reclining, obviously "},
	//
	//
	{" ス", "す", "su", "suス", "す",

		" su",
		" su",
		" su",
		" su"},
	//
	//
	{" セ", "せ", "se", "seセ", "せ",

		" se:せ:セ  ... セ is just a bit more angular, as is the way with most Katakana ",
		" se:せ looks like a face saying something ... 'se'ing something ",
		" TT: pinky to the 'P' char",
		""},
	//
	//
	{" ソ", "そ", "so", "soソ", "そ",

		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like:  し:シ  or  no:ノ, but, MOSTLY, it looks like n:ん:ン ",
		" Seriously?  so:ソ vs nh:ン  ... I guess nh got nh-ocked-down by an upper-cut to the chin "},
	//
	// sa becomes za, ji, zu, ze, zo --------------------------------------------------------------------------------
	//
	{" ザ", "ざ", "za", "zaザ", "ざ",

		" za",
		" za",
		" za",
		" za"},
	//
	//
	{" ジ", "じ", "ji", "jiジ", "じ",

		" ji:ジ is from shi:し and Seldom is it from chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" ji: the sound, can SOMETIMES also be from chi:ち:ヂ ,it would also be the sound ji, Not gi (that being ぎ:gi:ギ)",
		" ... So, remember, the ヂ:ジ sound:ji is always from either chi:ち or shi:し with a dakuten 濁点",
		"ジ is from shi:し and Seldom is it from chi:ち ,it's the sound gee, NEVER gi (that being ぎ:gi:ギ)"},
	//
	// zu has 2 KeyR values **** when prompting with naked Romaji **** 1:"zu", and 2::づ
	{" ズ", "ず", "zu", "zuズ", "ず",

		" zu:ず:ズ:づ:ヅ: , is either ず:ズ or づ:ヅ , as they are both the same sound",
		" zu: for the Kata think 'Sue' ズ is running with her hair flowing behind her (from su:す) ",
		"  TT: index< to the 'R' char,  they were looking for angles (so, now Sue is running ス) ",
		"  ... or for the 'tsu:つ' form think 'tsu' sounds like 'zu':づ "},
	//
	//
	{" ゼ", "ぜ", "ze", "zeゼ", "ぜ",

		" ze",
		" ze",
		" ze",
		" ze"},
	//
	//
	{" ゾ", "ぞ", "zo", "zoゾ", "ぞ",

		" zo",
		" zo",
		" zo",
		" zo"},
	//
	// ya, yu, yo's of shi:し:シ ------------------------------------------------------------------------
	//
	{" シャ", "しゃ", "sha", "shaシャ", "しゃ",

		" sha",
		" sha",
		" sha",
		" sha"},
	//
	//
	{" シュ", "しゅ", "shu", "shuシュ", "しゅ",

		" shu",
		" shu",
		" shu",
		" shu"},
	//
	//
	{" ショ", "しょ", "sho", "shoショ", "しょ",

		" sho",
		" sho",
		" sho",
		" sho"},
	//
	// ya, yu, yo's of ji:じ:ジ -------------------------------------------------------------------------
	//
	{" ジャ", "じゃ", "ja", "jaジャ", "じゃ",

		" ja:じゃ:ジャ is nearly-always from shi:し and really never from chi:ち , ヂャ is obsolete",
		" ja:ジャ:じゃ , obsolete:ヂャ is an obsolete construct for ja",
		" ja:",
		" ja:"},
	//
	//
	{" ジュ", "じゅ", "ju", "juジュ", "じゅ",

		" ju:ジュ:じゅ is nearly-always from shi:し and really never from chi:ち",
		" ju:じゅ:ジュ , obsolete:ヂュ is an obsolete construct for ju",
		" ju: ",
		" ju: "},
	//
	//
	{" ジョ", "じょ", "jo", "joジョ", "じょ",

		" jo:じょ:ジョ is nearly-always from shi:し and really never from chi:ち",
		" jo:ジョ:じょ , obsolete:ヂョ is an obsolete construct for jo",
		" jo: ",
		" jo: "},
	//
	// ta group: (ta-da,ji,zu,de,do) ================================================================================
	//
	{" タ", "た", "ta", "taタ", "た",

		" ta:た:タタタタ  ta-da!... it's a ku:くクタ with a drool ... and that's くool I guess ",
		" ta:た:タ ,  compare:  ku:ク  &  ke:ケ ",
		" TT: pinky< to the 'Q' char ",
		" top of ta:タ looks like た, more at least than ku:ク looks like く , or than he:ケ looks like へ "},
	//
	//
	{" チ", "ち", "chi", "chiチ", "ち",

		" chi:ち:チ , try to see the Kata as having its middle line as the top of the backwards 'c' of ち チ ",
		" ... I know it is hard.  chi:ち:チ ,  &-compare: the Kata chi:チ  to the  Kata te:テ ",
		" TT: pinky on the 'A' char ",
		" Kata chi bares 'some' resemblance if we see its middle line as the top of the backwards 'c' of ち "},
	//
	//
	{" ツ", "つ", "tsu", "tsuツ", "つ",

		" tsu",
		" tsu",
		" tsu",
		" tsu"},
	//
	//
	{" テ", "て", "te", "teテ", "て",

		" te:て:テ ",
		" te:て:テ ,  compare:  chi:チ テ ",
		" TT: ring< to the 'W' char, te:テ:て  went the wrong way, and gained a flat hat ",
		""},
	//
	//
	{" ト", "と", "to", "toト", "と",

		" to:と:ト   Kata 'toe' is flipped-out. Kicked in the balls. on its head: と -> ト ",
		" Katakana to:と is flipped-out. Kicked in the balls, on its head ",
		" TT: ring on the 'S' char, Kata toe is flipped-out. Kicked in the balls. on its head: と -> ト ",
		" Kata toe is flipped-out. Kicked in the balls. on its head: と -> ト "},
	//
	// ta becomes da, ji, zu, de, do --------------------------------------------------------------------------------
	//
	{" ダ", "だ", "da", "daダ", "だ",

		" da",
		" da",
		" da",
		" da"},
	//
	//
	/*
		{" ヂ", "ぢ", "ji", "jiヂ", "ぢ",

			" ji",
			" ji",
			" ji",
			" ji"},
		//
	*/
	//
	// zu has 2 KeyR values **** when prompting with naked Romaji **** 1:"zu", and 2::づ
	{" ヅ", "づ", "zu", "zuヅ", "づ",

		" zu:づ:ヅ:ず:ズ , is either ず:ズ or づ:ヅ , as they are both the same sound ",
		" zu: for the Kataす think 'Sue' ズ is running with her hair flowing behind her (from su:す) ",
		" TT: L-pinky>v to the 'Z' char ",
		"  ... for the Kata of the tsunami 'tsu:つ' form think 'tsu' sounds like 'zu':づ "},
	//
	//
	{" デ", "で", "de", "deデ", "で",

		" de",
		" de",
		" de",
		" de"},
	//
	//
	{" ド", "ど", "do", "doド", "ど",

		" do",
		" do",
		" do",
		" do"},
	//
	// ya, yu, yo's of chi:ち:チ ------------------------------------------------------------------------
	//
	{" チャ", "ちゃ", "cha", "chaチャ", "ちゃ",

		" cha",
		" cha",
		" cha",
		" cha"},
	//
	//
	{" チュ", "ちゅ", "chu", "chuチュ", "ちゅ",

		" chu",
		" chu",
		" chu",
		" chu"},
	//
	//
	{" チョ", "ちょ", "cho", "choチョ", "ちょ",

		" cho",
		" cho",
		" cho",
		" cho"},
	//
	// ya, yu, yo's of ji:ぢ:ヂ -------------------------------------------------------------------------
	//
	/*
		{" ヂャ", "ぢゃ", "ja", "jaヂャ", "ぢゃ",

			" ja",
			" ja",
			" ja",
			" ja"},
		//
	*/
	//
	/*
			{" ヂュ", "ぢゅ", "ju", "juヂュ", "ぢゅ",

				" ju",
				" ju",
				" ju",
				" ju"},
			//

		//
		{" ヂョ", "ぢょ", "jo", "joヂョ", "ぢょ",

			" jo",
			" jo",
			" jo",
			" jo"},
		//
	*/
	// na group: (always a naked group) =============================================================================
	//
	{" ナ", "な", "na", "naナ", "な",

		" na:な:ナ  the Kata ナ is just like the left side of the Hira な  --  more-so even: な:ナ ",
		" na:な:ナ  , compare Kata:na:ナ  to  me:メ  &  nu:ヌ  ( those last two are still similar: め ぬ )",
		" TT: index< to the 'U' char,   ナ is simple, very simple ... na-t complex at-all as Kata ",
		" Hiragana t-ies a knot, sorta like there are two chars thar "},
	//
	//
	{" ニ", "に", "ni", "niニ", "に",

		" ni:に:ニ  I can see a knee-cap in the Hiragana ",
		" ni:に:ニ  ",
		" TT: middle< to the 'I' char ",
		""},
	//
	//
	{" ヌ", "ぬ", "nu", "nuヌ", "ぬ",

		" nu:ぬ:ヌ  compare:  me:メ  &  na:ナ ",
		"",
		" TT: R-ring<--- to the '1' char, at least two lines cross, nu flew ヌ -- till it hit a ceiling ",
		""},
	//
	//
	{" ネ", "ね", "ne", "neネ", "ね",

		" ne:ね:ネ  Kata looks like a hoe that got nailed ネ ",
		"",
		" TT: R-ring<v to the ,< chars  ...  Yea, to me it is a hoe:ほ that got nailed ネ ",
		""},
	//
	//
	{" ノ", "の", "no", "noノ", "の",

		" no",
		" no",
		" no",
		" no"},
	//
	// ya, yu, yo's of ni:に:ニ -------------------------------------------------------------------------
	//
	{" ニャ", "にゃ", "nya", "nyaニャ", "にゃ",

		" nya",
		" nya",
		" nya",
		" nya"},
	//
	//
	{" ニュ", "にゅ", "nyu", "nyuニュ", "にゅ",

		" nyu",
		" nyu",
		" nyu",
		" nyu"},
	//
	//
	{" ニョ", "にょ", "nyo", "nyoニョ", "にょ",

		" nyo",
		" nyo",
		" nyo",
		" nyo"},
	//
	// ha group: (ha-fu ba:pa) ha:hi:fu:he:ho =======================================================================
	//
	{" ハ", "は", "ha", "haハ", "は",

		" ha",
		" ha",
		" ha",
		" ha"},
	//
	//
	{" ヒ", "ひ", "hi", "hiヒ", "ひ",

		" hi",
		" hi",
		" hi",
		" hi"},
	//
	//
	{" フ", "ふ", "fu", "fuフ", "ふ",

		" fu",
		" fu: think Mt Fuji 'fu'ji ",
		" fu",
		" fu"},
	//
	//
	{" ヘ", "へ", "he", "heヘ", "へ",

		" he",
		" he",
		" he",
		" he"},
	//
	//
	{" ホ", "ほ", "ho", "hoホ", "ほ",

		" ho",
		" ho",
		" ho",
		" ho"},
	//
	// ba, bi, bu, be, bo -------------------------------------------------------------------------------------------
	//
	{" バ", "ば", "ba", "baバ", "ば",

		" ba",
		" ba",
		" ba",
		" ba"},
	//
	//
	{" ビ", "び", "bi", "biビ", "び",

		" bi",
		" bi",
		" bi",
		" bi"},
	//
	//
	{" ブ", "ぶ", "bu", "buブ", "ぶ",

		" bu",
		" bu",
		" bu",
		" bu"},
	//
	//
	{" ベ", "べ", "be", "beベ", "べ",

		" be",
		" be",
		" be",
		" be"},
	//
	//
	{" ボ", "ぼ", "bo", "boボ", "ぼ",

		" bo",
		" bo",
		" bo",
		" bo"},
	//
	// pa, pi, pu, pe, po -------------------------------------------------------------------------------------------
	//
	{" パ", "ぱ", "pa", "paパ", "ぱ",

		" pa",
		" pa",
		" pa",
		" pa"},
	//
	//
	{" ピ", "ぴ", "pi", "piピ", "ぴ",

		" pi",
		" pi",
		" pi",
		" pi"},
	//
	//
	{" プ", "ぷ", "pu", "puプ", "ぷ",

		" pu",
		" pu",
		" pu",
		" pu"},
	//
	//
	{" ペ", "ぺ", "pe", "peペ", "ぺ",

		" pe",
		" pe",
		" pe",
		" pe"},
	//
	//
	{" ポ", "ぽ", "po", "poポ", "ぽ",

		" po",
		" po",
		" po",
		" po"},
	//
	// ya, yu, yo's of hi:ひ:ヒ -------------------------------------------------------------------------
	//
	{" ヒャ", "ひゃ", "hya", "hyaヒャ", "ひゃ",

		" hya",
		" hya",
		" hya",
		" hya"},
	//
	//
	{" ヒュ", "ひゅ", "hyu", "hyuヒュ", "ひゅ",

		" hyu",
		" hyu",
		" hyu",
		" hyu"},
	//
	//
	{" ヒョ", "ひょ", "hyo", "hyoヒョ", "ひょ",

		" hyo",
		" hyo",
		" hyo",
		" hyo"},
	//
	// ya, yu, yo's of bi:び:ビ -------------------------------------------------------------------------
	//
	{" ビャ", "びゃ", "bya", "byaビャ", "びゃ",

		" bya",
		" bya",
		" bya",
		" bya"},
	//
	//
	{" ビュ", "びゅ", "byu", "byuビュ", "びゅ",

		" byu",
		" byu",
		" byu",
		" byu"},
	//
	//
	{" ビョ", "びょ", "byo", "byoビョ", "びょ",

		" byo",
		" byo",
		" byo",
		" byo"},
	//
	//
	{" ピャ", "ぴゃ", "pya", "pyaピャ", "ぴゃ",

		" pya",
		" pya",
		" pya",
		" pya"},
	//
	// ya, yu, yo's of pi:ぴ:ピ -------------------------------------------------------------------------
	//
	{" ピュ", "ぴゅ", "pyu", "pyuピュ", "ぴゅ",

		" pyu",
		" pyu",
		" pyu",
		" pyu"},
	//
	//
	{" ピョ", "ぴょ", "pyo", "pyoピョ", "ぴょ",

		" pyo",
		" pyo",
		" pyo",
		" pyo"},
	//
	// ma group: ====================================================================================================
	//
	{" マ", "ま", "ma", "maマ", "ま",

		" ma",
		" ma",
		" ma",
		" ma"},
	//
	//
	{" ミ", "み", "mi", "miミ", "み",

		" mi",
		" mi",
		" mi",
		" mi"},
	//
	//
	{" ム", "む", "mu", "muム", "む",

		" mu",
		" mu",
		" mu",
		" mu"},
	//
	//
	{" メ", "め", "me", "meメ", "め",

		" me",
		" me",
		" me",
		" me"},
	//
	//
	{" モ", "も", "mo", "moモ", "も",

		" mo",
		" mo",
		" mo",
		" mo"},
	//
	// ya, yu, yo's of mi:み:ミ -------------------------------------------------------------------------
	//
	{" ミャ", "みゃ", "mya", "myaミャ", "みゃ",

		" mya",
		" mya",
		" mya",
		" mya"},
	//
	//
	{" ミュ", "みゅ", "myu", "myuミュ", "みゅ",

		" myu",
		" myu",
		" myu",
		" myu"},
	//
	//
	{" ミョ", "みょ", "myo", "myoミョ", "みょ",

		" myo",
		" myo",
		" myo",
		" myo"},
	//
	// group ya yu yo ===============================================================================================
	//
	{" ヤ", "や", "ya", "yaヤ", "や",

		" ya",
		" ya",
		" ya",
		" ya"},
	//
	//
	{" ユ", "ゆ", "yu", "yuユ", "ゆ",

		" yu",
		" yu",
		" yu",
		" yu"},
	//
	//
	{" ヨ", "よ", "yo", "yoヨ", "よ",

		" yo",
		" yo",
		" yo",
		" yo"},
	//
	// ra group: (pronounced more like la) ==========================================================================
	//
	{" ラ", "ら", "ra", "raラ", "ら",

		" ra",
		" ra",
		" ra",
		" ra"},
	//
	//
	{" リ", "り", "ri", "riリ", "り",

		" ri",
		" ri",
		" ri",
		" ri"},
	//
	//
	{" ル", "る", "ru", "ruル", "る",

		" ru",
		" ru",
		" ru",
		" ru"},
	//
	//
	{" レ", "れ", "re", "reレ", "れ",

		" re",
		" re",
		" re",
		" re"},
	//
	//
	{" ロ", "ろ", "ro", "roロ", "ろ",

		" ro",
		" ro",
		" ro",
		" ro"},
	//
	// ya, yu, yo's of ri:り:リ ---------------------------------------------------------------------------
	//
	{" リャ", "りゃ", "rya", "ryaリャ", "りゃ",

		" rya",
		" rya",
		" rya",
		" rya"},
	//
	//
	{" リュ", "りゅ", "ryu", "ryuリュ", "りゅ",

		" ryu",
		" ryu",
		" ryu",
		" ryu"},
	//
	//
	{" リョ", "りょ", "ryo", "ryoリョ", "りょ",
		" ryo",
		" ryo",
		" ryo",
		" ryo"},
	//
	// wa and wo ====================================================================================================
	//
	{" ワ", "わ", "wa", "waワ", "わ",

		" wa",
		" wa",
		" wa",
		" wa"},
	//
	//
	{" ヲ", "を", "wo", "woヲ", "を",

		" wo",
		" wo",
		" wo",
		" wo"},
	//
	// n ============================================================================================================
	//
	{" ン", "ん", "n", "nン", "ん",

		" n",
		" n",
		" n",
		" n"},
}
