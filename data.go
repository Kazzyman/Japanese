package main

var fileOfCards = []charSetStruct{

	// vowels: (a-i-u-e-o) ==========================================================================================
	//
	{"ア", "あ", "a", "aア", "あ",

		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: middle<- to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" compare:  me:め  あ  nu:ぬ  ne:ね      あ-Fuck-mae-te!"},
	// 2 times for a:あ:ア
	{"ア", "あ", "a", "aア", "あ",

		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: middle<- to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" compare:  me:め  あ  nu:ぬ  ne:ね      あ-Fuck-mae-te!"},
	//
	//
	{"イ", "い", "i", "iイ", "い",

		" i:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: middle< to the 'E' char  い:イ ",
		" ... at least it is still two mostly-vertical lines: イ "},
	// 2 times for i:い:イ
	{"イ", "い", "i", "iイ", "い",

		" i:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: middle< to the 'E' char  い:イ ",
		" ... at least it is still two mostly-vertical lines: イ "},
	//
	//
	{"ウ", "う", "u", "uウ", "う",

		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ウ more-angular and with a tick for its top line   Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" TT: middle> to the '4' char  う : Kata looks like the wa:ワ albeit with a mohawk ' ",
		" う:ウ ,  compare:  wa:ワ cascading wa-ter fall,　um-'ca'y "},
	// 2 times for u:う:ウ
	{"ウ", "う", "u", "uウ", "う",

		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ウ more-angular and with a tick for its top line ",
		" TT: middle> to the '4' char  う : Kata looks like the wa:ワ albeit with a mohawk ' ",
		" う:ウ ,  compare:  wa:ワ cascading wa-ter fall,　um-'ca'y "},
	//
	//
	{"エ", "え", "e", "eエ", "え",

		" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
		" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
		" TT: index> to the '5' char   e:え:エ ",
		" エ, eye see it as a ... an, eye ??  eh?:え:エ"},
	// 2 times for e:え
	{"エ", "え", "e", "eエ", "え",

		" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
		" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
		" TT: index> to the '5' char   e:え:エ ",
		" エ, eye see it as a ... an, eye ??  eh?:え:エ"},
	//
	//
	{"オ", "お", "o", "oオ", "お",

		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go, cloak trailing behind ... maybe? "},
	// 2 times for o:お:オ
	{"オ", "お", "o", "oオ", "お",

		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go, cloak trailing behind ... maybe? "},
	//
	// ka group: (ka-ga) ============================================================================================
	//
	{"カ", "か", "ka", "kaカ", "か",

		" ka:か:カ ... is an easy one!  It simply looks like a 'K' ",
		" ... Kata is the same カ , albeit more-angular and with one-less line to draw than か ",
		" TT: index--> to the 'T' char   か ",
		" か:カ ... is an easy one!  It simply looks like a 'K'"},
	//
	//
	{"キ", "き", "ki", "kiキ", "き",

		" ki:き:キ ... is an easy one!   キ has the same top as ki:き ... and they look like 'keys' ",
		" ki:き ... they both have the same top:  き キ",
		" TT: L-index> to the 'G' char,  Think: ki-of-G  き ,  compare:  sa:さ-ki:き  さき ... saki ",
		"き:キ ... is an easy one!   キ has the same top as き ... and they look like 'keys' "},
	//
	//
	{"ク", "く", "ku", "kuク", "く",

		" ku:く:ク ,  Kata: starting with one angle, they settled for this??  く:ク ",
		" ku:く:ク ,  compare:  ta:タ  ke:ケ  ku:ク  クケタ,  it's kuku, kookoo I tell you! ",
		" TT: R-index<- to the 'H' char  く:ク ",
		"く:ク ,  compare:  クケタ ta:タ  ke:ケ  ?:ク  ,  it's kookoo I tell you!"},
	//
	//
	{"ケ", "け", "ke", "keケ", "け",

		" ke:け:ケ  bits of it are there, just as many curves though ,  Compare to:  ku:ク  ta:タ ",
		" ke:け  ... looks enough like a Keg ...  Foam-ment THAT image in your brain ",
		" TT: pinky-> one-over to the ':*' chars   け:ケ ",
		"け:ケ  bits still there, as many curves though,  Compare:  ku:ク  ta:タ  クケタ"},
	//
	//
	{"コ", "こ", "ko", "koコ", "こ",

		" ko:こ:コ ... it makes sense, 'cause angles:  ko:コ ,  compare:  ni:に:ニ ",
		" ... see the two Koy fish?   'ko'oy fish? ",
		" TT: index <--- to the 'B' char,  ko way-down below  こ ",
		" こ:コ ... it makes sense, 'cause angles,  compare:  ni:に:ニ"},
	// 2 times for ko:こ:コ
	{"コ", "こ", "ko", "koコ", "こ",

		" ko:こ:コ ... it makes sense, 'cause angles:  ko:コ ,  compare:  ni:に:ニ ",
		" ... see the two Koy fish?   'ko'oy fish? ",
		" TT: index <--- to the 'B' char,  ko way-down below  こ ",
		" こ:コ ... it makes sense, 'cause angles,  compare:  ni:に:ニ"},
	//
	// ka becomes ga group: -----------------------------------------------------------------------------------------

	//
	{"ガ", "が", "ga", "gaガ", "が",

		" ga:が same ガ albeit more angular with one-less line to draw than が",
		" ga:が is an easy one. ka--> ga,gi,gu,ge,go",
		" TT: index --> to the 'T' char    が ",
		" が same ガ albeit more angular with one-less line to draw than. Gaut it? が"},
	//
	//
	{"ギ", "ぎ", "gi", "giギ", "ぎ",

		" gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ",
		" ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" TT: 'G' char, ぎ:ギ, compare to za:ざ from sa:さ",
		" It's sound is always from き ,and NEVER from し or ち  ka--> ga,?,gu,ge,go"},
	//
	//
	{"グ", "ぐ", "gu", "guグ", "ぐ",

		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:za:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  za:ダ, ge:ゲ    ダグゲ　",
		" TT: 'H' char    ぐ:グ ",
		" ぐ:グ  Compare:  za:ダ, ge:ゲ  ダグゲ  google:  ka,ki,ku,ke,ko-->ga,gi,?,ge,go"},
	// 2 times for gu:ぐ:グ
	{"グ", "ぐ", "gu", "guグ", "ぐ",

		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:za:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  za:ダ, ge:ゲ    ダグゲ　",
		" TT: 'H' char    ぐ:グ ",
		" ぐ:グ  Compare:  za:ダ, ge:ゲ  ダグゲ  google:  ka,ki,ku,ke,ko-->ga,gi,?,ge,go"},
	//
	//
	{"ゲ", "げ", "ge", "geゲ", "げ",

		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ",
		" TT:  ':*' chars    げ",
		" げ Compare to ku:gu:グ, and ta:za:ダ   ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go"},
	// 2 times for ge:げ:ゲ
	{"ゲ", "げ", "ge", "geゲ", "げ",

		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ",
		" TT:  ':*' chars    げ",
		" げ Compare to ku:gu:グ, and ta:za:ダ   ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go"},
	//
	//
	{"ゴ", "ご", "go", "goゴ", "ご",

		" go:ご:ゴ  it makes sense, 'cause angles' ",
		" go:ご:ゴ ",
		" TT: index <--- to the 'B' char    ご",
		" ご:ゴ  it makes sense, 'cause angles'   ka,ki,ku,ke,ko-->ga,gi,gu,ge,?"},
	//
	{"ゴ", "ご", "go", "goゴ", "ご",

		" go:ご:ゴ  it makes sense, 'cause angles' ",
		" go:ご:ゴ ",
		" TT: index <--- to the 'B' char    ご",
		" ご:ゴ  it makes sense, 'cause angles'   ka,ki,ku,ke,ko-->ga,gi,gu,ge,?"},
	//
	// ya, yu, yo's of ki:き ---------------------------------------------------------------------------
	//
	{"キャ", "きゃ", "kya", "kyaキャ", "きゃ",

		" kya:きゃ is an easy one, キャ has the same top, & the 'ya' is similar",
		" kya:きゃ キャ has the same top and the 'ya' is similar",
		" TT: 'G' char, then index <-- to the 7 char    きゃ,   Compare:  sa:さ ",
		" きゃ is an easy one, キャ has the same top, & the 'ya' is similar"},
	//
	//
	{"キュ", "きゅ", "kyu", "kyuキュ", "きゅ",

		" kyu:きゅ キュ has the same top",
		" kyu: ゅ:ュ",
		" TT: yu:ゅ index->, to the '8' char, yu looks like:  ユ are #1",
		" Key-ゅ   and:  ユ are #1     キュ has the same top as き "},
	//
	//
	{"キョ", "きょ", "kyo", "kyoキョ", "きょ",

		" kyo:きょ:キョ   キョ  has the same top",
		" kyo:キョ ",
		" TT: ょ:ョ middle to the '9' char, triple yo:ょ ヨ",
		" Key-ょ  キョ has the same top as き   hira:ょ is a yoyo , triple yo:ヨ"},
	//
	// ya, yu, yo's of gi:ぎ ---------------------------------------------------------------------------
	//
	{"ギャ", "ぎゃ", "gya", "gyaギャ", "ぎゃ",

		" gya  gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し",
		" gya:ギャ ",
		" TT: ya:ゃ index <-- to the 7 char",
		" ぎ:ギ is an easy Kata, ギ has the same top as ki:き   ka,ki,ku,ke,ko-->ga,?,gu,ge,go"},
	//
	//
	{"ギュ", "ぎゅ", "gyu", "gyuギュ", "ぎゅ",

		" gyu:ぎゅ:ギュ   ギ has same top as ki:き   ka,(ki),ku,ke,ko-->ga,(gi),gu,ge,go",
		" gyu:ぎゅ:ギュ ",
		" TT: yu:ゅ index->, to the '8' char, yu looks like:  ユ are #1",
		" ぎ:ギ is an easy Kata, ギ has the same top as ki:き    ka,ki,ku,ke,ko-->ga,?,gu,ge,go"},
	//
	//
	{"ギョ", "ぎょ", "gyo", "gyoギョ", "ぎょ",

		" gyo: gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ",
		" gyo: ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" TT: 'G' char, ぎ:ギ, then ょ:ョ:middle to the '9' char, triple yo:ょ ヨ      Compare: ざ",
		" It's sound is always from き ,and NEVER from し or ち   ka,ki,ku,ke,ko--> ga,?,gu,ge,go"},
	//
	// sa group: (sa-za,ji,zu,za,zo) ================================================================================
	//
	{"サ", "さ", "sa", "saサ", "さ",

		" sa:さ:サ , if you 'se':せ そサ ,  compare:  Hira se:せ  to  Kata sa:サ  せサ ",
		" ... , but it looks an awful-lot like Hira se:せ, I'd say   ... just 'Hir'-say ?? ",
		" TT: ring>v to the 'X' char  --  at least Kata 'sa' still goes to the left ... ",
		" さ:サ ,  Compare:  Hira:se:せ  to  Kata:?:サ  せサ se:? "},
	//
	//  ... what? ... two nasal blasts ... obviously! Yes while reclining, obviously
	{"シ", "し", "shi", "shiシ", "し",

		" shi:し:シ  'she' sleeps & snores゛,  compare: Kata:  tsu:ツ   so:ソ　to the more-laid-back shi:シ ",
		" shi:し:シ  し looks like a sheep hook. Is there even such a thing?? ,  compare:  ツシ  tsu-shi ",
		" TT: middle on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し  (no angles here, less curve though)"},
	//
	//
	{"ス", "す", "su", "suス", "す",

		" su:す:ス ",
		" ス sue:す looks like she is running (they were looking for angles)",
		" TT: index < to the 'R' char,  (sue is running ス)",
		" ス:す"},
	//
	//
	{"セ", "せ", "se", "seセ", "せ",

		" se:せ:セ  ... セ is just a bit more angular, as is the way with most Katakana ",
		" se:せ looks like a face saying something ... 'se'ing something ",
		" TT: pinky to the 'P' char",
		" せ:セ  ... セ is just a bit more angular, as is the way with most Katakana"},
	//
	// Seriously?  so:ソ vs nh:ン  ... I guess nh got nh-ocked-down by an upper-cut to the chin
	{"ソ", "そ", "so", "soソ", "そ",

		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick "},
	// 3 times so:そ:ソ
	{"ソ", "そ", "so", "soソ", "そ",

		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick "},
	//
	{"ソ", "そ", "so", "soソ", "そ",

		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" そ:ソ ,  compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick "},
	//
	// sa becomes za, ji, zu, ze, zo --------------------------------------------------------------------------------
	//
	{"ザ", "ざ", "za", "zaザ", "ざ",

		" sa-->za:ざ ",
		" sa-->za   za:za:ザ ",
		" TT: down on the 'X' char for sa:za",
		" ?:ざ:ザ "},
	//
	//
	{"ジ", "じ", "ji", "jiジ", "じ",

		" ji:ジ is from shi:し & Seldom is it fr chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" ji, the sound, VERY RARELY is from chi:ち:ヂ ,but is NEVER the sound gi (that being ぎ:gi:ギ)",
		" TT: 'D' char, remember, the sound:ji is nearly-always from shi:し with a dakuten 濁点 じ:ジ",
		" ジ:じ the sound, VERY Seldom from ち, it's sound:gee:jeep, NEVER gi:geek(that being ぎ:gi:ギ)"},
	//
	// zu has 2 KeyR values **** when prompting with naked Romaji **** 1:"zu", and 2::づ
	{"ズ", "ず", "zu", "zuズ", "ず",

		" zu:ず:ズ:づ:ヅ: , is either ず:ズ or づ:ヅ , as they are both the same sound",
		" zu: for the Kata think 'Sue' ズ is running with her hair flowing behind her (from su:す) ",
		"  TT: index< to the 'R' char,  they were looking for angles (so, now Sue is running ス) ",
		" ?:ず:ズ:づ:ヅ   both are the same sound: it's a wave つ  or a  snake ず "},
	//
	//
	{"ゼ", "ぜ", "ze", "zeゼ", "ぜ",

		" sa-->za   se-->ze ",
		" ze:ゼ:ぜ",
		" TT: pinky to the 'P' char   ze:ゼ:ぜ",
		" ?:ゼ:ぜ   sa-->za  se:? "},
	//
	//
	{"ゾ", "ぞ", "zo", "zoゾ", "ぞ",

		" zo:ぞ:ゾ  Compare: n:ん:ン which lays-down more",
		" zo:ぞ:ゾ    Compare:  し:シ  or  no:ノ,  or  n:ん:ン  ",
		" TT: index <-- to the C char     zo:ぞ:ゾ",
		" ?:ぞ:ゾ   in any case it is not no:ノ  or  n:ん:ン  which lays-down more  "},
	//
	{"ゾ", "ぞ", "zo", "zoゾ", "ぞ",

		" zo:ぞ:ゾ  Compare: n:ん:ン which lays-down more",
		" zo:ぞ:ゾ    Compare:  し:シ  or  no:ノ,  or  n:ん:ン  ",
		" TT: index <-- to the C char     zo:ぞ:ゾ",
		" ?:ぞ:ゾ   in any case it is not no:ノ  or  n:ん:ン  which lays-down more  "},
	//
	// ya, yu, yo's of shi:し:シ ------------------------------------------------------------------------
	//
	{"シャ", "しゃ", "sha", "shaシャ", "しゃ",

		" sha",
		" sha",
		" sha",
		" sha"},
	//
	//
	{"シュ", "しゅ", "shu", "shuシュ", "しゅ",

		" shu",
		" shu",
		" shu",
		" shu"},
	//
	//
	{"ショ", "しょ", "sho", "shoショ", "しょ",

		" sho",
		" sho",
		" sho",
		" sho"},
	//
	// ya, yu, yo's of ji:じ:ジ -------------------------------------------------------------------------
	//
	{"ジャ", "じゃ", "ja", "jaジャ", "じゃ",

		" ja:じゃ:ジャ is nearly-always from shi:し and really never from chi:ち , ヂャ is obsolete",
		" ja:ジャ:じゃ , obsolete:ヂャ is an obsolete construct for ja",
		" ja:",
		" ja:"},
	//
	//
	{"ジュ", "じゅ", "ju", "juジュ", "じゅ",

		" ju:ジュ:じゅ is nearly-always from shi:し and really never from chi:ち",
		" ju:じゅ:ジュ , obsolete:ヂュ is an obsolete construct for ju",
		" ju: ",
		" ju: "},
	//
	//
	{"ジョ", "じょ", "jo", "joジョ", "じょ",

		" jo:じょ:ジョ is nearly-always from shi:し and really never from chi:ち",
		" jo:ジョ:じょ , obsolete:ヂョ is an obsolete construct for jo",
		" jo: ",
		" jo: "},
	//
	// ta group: (ta-da,ji,zu,de,do) ================================================================================
	//
	{"タ", "た", "ta", "taタ", "た",

		" ta:た:タタタタ  ta-da!... it's a ku:くクタ with a drool ... and that's くool I guess ",
		" ta:た:タ ,  compare:  ku:ク  &  ke:ケ ",
		" TT: pinky< to the 'Q' char ",
		" top of ta:タ looks like た,   more than ku:ク looks like く  "},
	//
	//
	{"チ", "ち", "chi", "chiチ", "ち",

		" chi:ち:チ ,   ち チ ",
		" ...  chi:ち:チ ,  Compare:  chi:チ  te:テ ",
		" TT: pinky on the 'A' char ",
		" Kata bares 'some' resemblance to ち "},
	//
	//
	{"ツ", "つ", "tsu", "tsuツ", "つ",

		" tsu:つ:ツ  see water crashing on the she shore",
		" tsu",
		" TT: tsu",
		" ?:つ:ツ "},
	//
	//
	{"テ", "て", "te", "teテ", "て",

		" te:て:テ ",
		" te:て:テ ,  Compare:  chi:チ テ:て ",
		" TT: ring< to the 'W' char, te:テ:て  went the wrong way, and gained a flat hat ",
		" ?:て:テ ,  Compare:  chi:チ テ:て "},
	//
	//
	{"ト", "と", "to", "toト", "と",

		" to:と:ト   Kata 'toe' is flipped-out. Kicked in the balls. on its head: と -> ト ",
		" Katakana to:と is flipped-out. Kicked in the balls, On its head ",
		" TT: ring on the 'S' char, Kata toe is flipped-out. Kicked on its head: と -> ト ",
		" Kata ? is flipped-out. Kicked in the balls. On its head: と --> ト "},
	//
	// ta becomes da, ji, zu, de, do --------------------------------------------------------------------------------
	//
	{"ダ", "だ", "da", "daダ", "だ",

		" ta-->da  だ:ダ ",
		" da",
		" TT: ",
		" ta-->? "},
	//
	//
	// zu has 2 KeyR values **** when prompting with naked Romaji **** 1:"zu", and 2::づ
	{"ヅ", "づ", "zu", "zuヅ", "づ",

		" zu:づ:ヅ:ず:ズ , is either ず:ズ or づ:ヅ , as they are both the same sound ",
		" zu: for the Kataす think 'Sue' ズ is running with her hair flowing behind her (from su:す) ",
		" TT: L-pinky>v to the 'Z' char ",
		"  ... ?:づ:ヅ:ず:ズ    that big wave つ "},
	//
	//
	{"デ", "で", "de", "deデ", "で",

		" de:で:デ ",
		" de",
		" TT: de",
		" ?:で:デ "},
	//
	//
	{"ド", "ど", "do", "doド", "ど",

		" ta-->da  to --> do",
		" do:ど:ド ",
		" TT: do",
		" ?:ど:ド   ta-->da  ?-->do "},
	//
	{"ド", "ど", "do", "doド", "ど",

		" ta-->da  to --> do",
		" do:ど:ド ",
		" TT: do",
		" ?:ど:ド   ta-->da  ?-->do "},
	//
	// ya, yu, yo's of chi:ち:チ ------------------------------------------------------------------------
	//
	{"チャ", "ちゃ", "cha", "chaチャ", "ちゃ",

		" cha",
		" cha",
		" cha",
		" cha"},
	//
	//
	{"チュ", "ちゅ", "chu", "chuチュ", "ちゅ",

		" chu",
		" chu",
		" chu",
		" chu"},
	//
	//
	{"チョ", "ちょ", "cho", "choチョ", "ちょ",

		" cho",
		" cho",
		" cho",
		" cho"},
	//

	// na group: (always a naked group) =============================================================================
	//
	{"ナ", "な", "na", "naナ", "な",

		" na:な:ナ  Kata ナ is just like the left side of the Hira な  --  more-so even: な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  (those last two are still similar: め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" Hiragana t-ies a knot, sorta like there are two chars thar "},
	//
	{"ナ", "な", "na", "naナ", "な",

		" na:な:ナ  Kata ナ is just like the left side of the Hira な  --  more-so even: な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  (those last two are still similar: め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" Hiragana t-ies a knot, sorta like there are two chars thar "},
	//
	{"ナ", "な", "na", "naナ", "な",

		" na:な:ナ  Kata ナ is just like the left side of the Hira な  --  more-so even: な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  (those last two are still similar: め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" Hiragana t-ies a knot, sorta like there are two chars thar "},
	//
	//
	{"ニ", "に", "ni", "niニ", "に",

		" ni:に:ニ  I can see a knee-cap in the Hiragana ",
		" ni:に:ニ  ",
		" TT: middle< to the 'I' char ",
		" ?:に:ニ  I can see a knee-cap in the Hiragana "},
	//
	//
	{"ヌ", "ぬ", "nu", "nuヌ", "ぬ",

		" nu:ぬ:ヌ   Compare:  me:メ  &  na:ナ ",
		" nu flew ヌ -- till it hit a ceiling ",
		" TT: R-ring<--- to the '1' char, at least two lines cross, nu flew ヌ -- hit a ceiling ",
		" ?-flew ヌ -- till it hit a ceiling "},
	//
	//
	{"ネ", "ね", "ne", "neネ", "ね",

		" ne:ね:ネ  Kata looks like a hoe that got nailed ネ ",
		" something got nailed-down by that spike at the left ",
		" TT: R-ring<v to the ,< chars  ...  Yea, to me it is a hoe:ほ that got nailed ネ ",
		" ?:ね:ネ  Kata looks like a hoe that got nailed ネ"},
	//
	//
	{"ノ", "の", "no", "noノ", "の",

		" no:の:ノ",
		" no:の:ノ",
		" TT: no",
		" ?:の:ノ    ?-thank-you "},
	//
	// ya, yu, yo's of ni:に:ニ -------------------------------------------------------------------------
	//
	{"ニャ", "にゃ", "nya", "nyaニャ", "にゃ",

		" nya",
		" nya",
		" nya",
		" nya"},
	//
	//
	{"ニュ", "にゅ", "nyu", "nyuニュ", "にゅ",

		" nyu",
		" nyu",
		" nyu",
		" nyu"},
	//
	//
	{"ニョ", "にょ", "nyo", "nyoニョ", "にょ",

		" nyo",
		" nyo",
		" nyo",
		" nyo"},
	//
	// ha group: (ha-fu ba:pa) ha:hi:fu:he:ho =======================================================================
	//
	{"ハ", "は", "ha", "haハ", "は",

		" ハ:は:ha ",
		" ハ ha-haystack ハ ",
		" TT: ha",
		" ハ:は:? "},
	//
	//
	{"ヒ", "ひ", "hi", "hiヒ", "ひ",

		" hi:ひ:ヒ",
		" hi, pronounced hee ",
		" TT: index > to the 'V' char   hi",
		" ?:ひ:ヒ   ?-man "},
	//
	//
	{"フ", "ふ", "fu", "fuフ", "ふ",

		" fu:ふ:フ",
		" fu: think Mt Fuji 'fu'ji ",
		" fu or hu ",
		" it is the big mountain in Japan "},
	//
	// done to here
	{"ヘ", "へ", "he", "heヘ", "へ",

		" he",
		" he",
		" he",
		" he"},
	//
	//
	{"ホ", "ほ", "ho", "hoホ", "ほ",

		" ho",
		" ho",
		" ho",
		" ho"},
	//
	// ba, bi, bu, be, bo -------------------------------------------------------------------------------------------
	//
	{"バ", "ば", "ba", "baバ", "ば",

		" ba",
		" ba",
		" ba",
		" ba"},
	//
	//
	{"ビ", "び", "bi", "biビ", "び",

		" bi",
		" bi",
		" bi",
		" bi"},
	//
	//
	{"ブ", "ぶ", "bu", "buブ", "ぶ",

		" bu",
		" bu",
		" bu",
		" bu"},
	//
	//
	{"ベ", "べ", "be", "beベ", "べ",

		" be",
		" be",
		" be",
		" be"},
	//
	//
	{"ボ", "ぼ", "bo", "boボ", "ぼ",

		" bo",
		" bo",
		" bo",
		" bo"},
	//
	// pa, pi, pu, pe, po -------------------------------------------------------------------------------------------
	//
	{"パ", "ぱ", "pa", "paパ", "ぱ",

		" pa",
		" pa",
		" pa",
		" pa"},
	//
	//
	{"ピ", "ぴ", "pi", "piピ", "ぴ",

		" pi",
		" pi",
		" pi",
		" pi"},
	//
	//
	{"プ", "ぷ", "pu", "puプ", "ぷ",

		" pu",
		" pu",
		" pu",
		" pu"},
	//
	//
	{"ペ", "ぺ", "pe", "peペ", "ぺ",

		" pe",
		" pe",
		" pe",
		" pe"},
	//
	//
	{"ポ", "ぽ", "po", "poポ", "ぽ",

		" po",
		" po",
		" po",
		" po"},
	//
	// ya, yu, yo's of hi:ひ:ヒ -------------------------------------------------------------------------
	//
	{"ヒャ", "ひゃ", "hya", "hyaヒャ", "ひゃ",

		" hya",
		" hya",
		" hya",
		" hya"},
	//
	//
	{"ヒュ", "ひゅ", "hyu", "hyuヒュ", "ひゅ",

		" hyu",
		" hyu",
		" hyu",
		" hyu"},
	//
	//
	{"ヒョ", "ひょ", "hyo", "hyoヒョ", "ひょ",

		" hyo",
		" hyo",
		" hyo",
		" hyo"},
	//
	// ya, yu, yo's of bi:び:ビ -------------------------------------------------------------------------
	//
	{"ビャ", "びゃ", "bya", "byaビャ", "びゃ",

		" bya",
		" bya",
		" bya",
		" bya"},
	//
	//
	{"ビュ", "びゅ", "byu", "byuビュ", "びゅ",

		" byu",
		" byu",
		" byu",
		" byu"},
	//
	//
	{"ビョ", "びょ", "byo", "byoビョ", "びょ",

		" byo",
		" byo",
		" byo",
		" byo"},
	//
	//
	{"ピャ", "ぴゃ", "pya", "pyaピャ", "ぴゃ",

		" pya",
		" pya",
		" pya",
		" pya"},
	//
	// ya, yu, yo's of pi:ぴ:ピ -------------------------------------------------------------------------
	//
	{"ピュ", "ぴゅ", "pyu", "pyuピュ", "ぴゅ",

		" pyu",
		" pyu",
		" pyu",
		" pyu"},
	//
	//
	{"ピョ", "ぴょ", "pyo", "pyoピョ", "ぴょ",

		" pyo",
		" pyo",
		" pyo",
		" pyo"},
	//
	// ma group: ====================================================================================================
	//
	{"マ", "ま", "ma", "maマ", "ま",

		" ma",
		" ma",
		" ma",
		" ma"},
	//
	//
	{"ミ", "み", "mi", "miミ", "み",

		" mi",
		" mi",
		" mi",
		" mi"},
	//
	//
	{"ム", "む", "mu", "muム", "む",

		" mu",
		" mu",
		" mu",
		" mu"},
	//
	//
	{"メ", "め", "me", "meメ", "め",

		" me",
		" me",
		" me",
		" me"},
	//
	//
	{"モ", "も", "mo", "moモ", "も",

		" mo",
		" mo",
		" mo",
		" mo"},
	//
	// ya, yu, yo's of mi:み:ミ -------------------------------------------------------------------------
	//
	{"ミャ", "みゃ", "mya", "myaミャ", "みゃ",

		" mya",
		" mya",
		" mya",
		" mya"},
	//
	//
	{"ミュ", "みゅ", "myu", "myuミュ", "みゅ",

		" myu",
		" myu",
		" myu",
		" myu"},
	//
	//
	{"ミョ", "みょ", "myo", "myoミョ", "みょ",

		" myo",
		" myo",
		" myo",
		" myo"},
	//
	// group ya yu yo ===============================================================================================
	//
	{"ヤ", "や", "ya", "yaヤ", "や",

		" ya",
		" ya",
		" ya",
		" ya"},
	//
	//
	{"ユ", "ゆ", "yu", "yuユ", "ゆ",

		" yu",
		" yu",
		" yu",
		" yu"},
	//
	//
	{"ヨ", "よ", "yo", "yoヨ", "よ",

		" yo",
		" yo",
		" yo",
		" yo"},
	//
	// ra group: (pronounced more like la) ==========================================================================
	//
	{"ラ", "ら", "ra", "raラ", "ら",

		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra",
		" ra",
		" Compare:  hu:フ u:ウ  ラ:ら:? "},
	//
	//
	{"リ", "り", "ri", "riリ", "り",

		" ri",
		" ri",
		" ri",
		" ri"},
	//
	//
	{"ル", "る", "ru", "ruル", "る",

		" ru",
		" ru",
		" ru",
		" ru"},
	//
	//
	{"レ", "れ", "re", "reレ", "れ",

		" re",
		" re",
		" re",
		" re"},
	//
	//
	{"ロ", "ろ", "ro", "roロ", "ろ",

		" ro",
		" ro",
		" ro",
		" ro"},
	//
	// ya, yu, yo's of ri:り:リ ---------------------------------------------------------------------------
	//
	{"リャ", "りゃ", "rya", "ryaリャ", "りゃ",

		" rya",
		" rya",
		" rya",
		" rya"},
	//
	//
	{"リュ", "りゅ", "ryu", "ryuリュ", "りゅ",

		" ryu",
		" ryu",
		" ryu",
		" ryu"},
	//
	//
	{"リョ", "りょ", "ryo", "ryoリョ", "りょ",
		" ryo",
		" ryo",
		" ryo",
		" ryo"},
	//
	// wa and wo ====================================================================================================
	//
	{"ワ", "わ", "wa", "waワ", "わ",

		" wa",
		" wa",
		" wa",
		" wa"},
	//
	//
	{"ヲ", "を", "wo", "woヲ", "を",

		" wo",
		" wo",
		" wo",
		" wo"},
	//
	// n ============================================================================================================
	//
	{"ン", "ん", "n", "nン", "ん",

		" n",
		" n",
		" n",
		" n"},
}

var aCard = charSetStruct{}

// The structure of a single 'card' (aCard.) from fileOfCards
type charSetStruct struct {
	KeyK  string
	KeyH  string
	KeyR  string
	KeyRK string

	Keyh string // New name for the old name (below): aCard.'Value'
	//Value string

	Hint1h    string
	Hint2k    string
	Hint3TT   string
	HintSansR string
}

/*
 Above, we instantiate a series of struct objects as a slice of instances of the charSetStruct type
 fileOfCards being that "slice" of structures of type charSetStruct

 i.e., We are creating a slice named fileOfCards that holds instances of the charSetStruct type

Each element in the slice is initialized using the composite literal syntax whereby
... we are providing values for each field of the charSetStruct struct: i.e.,
... each set of values enclosed in curly braces { ... } represents an instance of the struct.
*/

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
