package main

// **do-this** menu
import (
	"fmt"
	"os"
	"time"
	//"math/rand"
)

//goland:noinspection ALL
func display_ListingOf_OptionsThese_AllHave_inCommon() {
	fmt.Println("    (using US or Alpha-Numeric input mode):")
	fmt.Println("        Enter 'menu' to return to the the main menu ")
	fmt.Println("        Enter 'dir' to redisplay this menu of available directives")
	fmt.Println("        Enter 'notes' for some background on Romaji conventions")
	fmt.Println("        Enter '?' for context-sensitive help ")
	fmt.Println("        Enter '??' for help on a particular Hiragana char")
	fmt.Println("        Enter 'set' to reset the prompt & \"key\" ")
	fmt.Println("        Enter 'stat' to view what you have done so far in the current session")
	fmt.Println("        Enter 'reset' to reset the hits logs")
	fmt.Println("        Enter 'exit' or 'quit' to terminate the app")
	//goland:noinspection ALL
	fmt.Println("\n")
}

func mainMenuPromptScanSelectAndBeginSelectedExorcise() {
	var mainMenuSelection string
	//seedFile_Maker(); createAndWrite_seedFile()
	for {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
		///

		fmt.Printf("" +
			" Menu main: me-i-n me-ni-yu\n" +
			"メインメニュー:\n\n\n" +
			// 1 -
			"-- 1 -- v v v \n" +
			"        nin-shi-ki  t-I-pu-su-ru       todo           Enter                  simple       very     useful\n" +
			"ひらがな を 認識 し タイプする 練習 を するには「1」を 入力  してください。 (シンプル で、非常 に 役立ちます)\n" +
			"^Hiragana   ^^Recognizing      ^^practicing        Enter^      ^please do     ^simple          useful^ \n\n\n" +
			// ^ 1 - Recognizing Hiragana & practicing typing, Enter 1 (simple and very useful) ^ ^
			//
			// 2 -
			"-- 2 -- v v v \n" +
			"                     pair             Hiragana    t-I-pu  practicing   todo         enter                sort-of \n" +
			"ローマ字-カタカナ の ペア を 認識 し、ひらがな を タイプする 練習 を するには「2」を 入力 してください。(ある程度簡単 で、役立ちます)\n" +
			"^Romaji-Katakana     pair    ^Recognizing        t-I-pu-su   ^practice  ^todo         ^enter   ^please do \n\n\n" +
			// 2 - Recognizing Romaji+Katakana, practicing typing Hiragana (sort-of simple and useful)
			//
			// 3 -
			"-- 3 -- v v v \n" +
			"Katakana   recogni   Hiragana  or     Romaji      t-I-pu    practice   todo\n" +
			"カタカナ を 認識 し、ひらがな また は ローマ字 を タイプする 練習 を するには「3」を 入力 してください。(非常に汎用性が高い)\n" +
			"^Katakana   ^Recog   Hiragana  or     Romaji       typing     ^practice              enter please do  \n\n\n" +
			// 3 - Katakana recognizing, practice typing Hiragana or Romaji, 3 Enter ( )
			//
			// 5 -
			"-- 5 -- v v v \n" +
			"   t-I-pu    practice   todo\n" +
			"ドリルラインをタイプする練習をするには「5」を入力してください。\n\n" +
			"" +
			// 5 -
			//
			// 6 -
			"-- 6 -- v v v \n" +
			"   t-I-pu    practice   todo\n" +
			"ひらがなを認識し、ローマ字をタイプする練習をするには「6」を入力してください。 \n\n" +
			"" +
			// 6-
			//
			// 7 -
			"-- 7 -- v v v \n" +
			"   t-I-pu    practice   todo\n" +
			"最も難しいカタカナを練習するには「7」を入力してください。\n\n" +
			"" +
			// 7 -
			//
			// exit -
			"-- exit -- v v v \n" +
			"   t-I-pu    practice   todo\n" +
			"終了するには「exit」と入力してください。" +
			"" +
			// exit -
			"\n\n")
		/// Display the Main Menu
		fmt.Printf("  Main Menu: \n\n" +
			"  Enter '1' to practice recognizing Romaji: and typing Hiragana (simple, quite useful)\n" +
			"  Enter '2' for recognizing Romaji-Katakana pairs: typing Hiragana (somewhat easy, useful)\n" +
			"  Enter '3' to practice recognizing Katakana, type either Hiragana or Romaji (very versatile)\n" +
			"  Enter '5' to practice typing drill lines\n" +
			"  Enter '6' mixed Hira and Kata prompts: answer with Romaji\n" +
			"  Enter '7' to practice the Most-Difficult kata, type either Hiragana or Romaji\n" +
			"  Enter '8' to practice Sequential Kata, type either Hiragana or Romaji\n" +
			"  Enter '9' to practice Sequential Hira, type either Hiragana or Romaji\n" +
			"  Enter 'exit' to quit\n\n\n")
		// Pause to accept (Scan) a string var, e.g., "7" or "exit", from the command line, from the user
		_, _ = fmt.Scan(&mainMenuSelection) // Create address for, and auto-dereference, var mainMenuSelection
		///
		/*
			///
			fmt.Printf("// To prove that mainMenuSelection is a variable of type string :\n")
			fmt.Printf("// First, prove that it is a string type with:\n\n" +
				"fmt.Printf(\"The type of 'mainMenuSelection' is: %%T \\n\\n\", mainMenuSelection)" +
				"\n\nWhich yields: \n\n")
			fmt.Printf("\tThe type of 'mainMenuSelection' is: %T \n\n", mainMenuSelection)

			fmt.Println("Note: in Go, constants don't have a direct reflection type.")

			fmt.Printf("Next, we execute the following code:\n\n" +
				"xType := reflect.TypeOf(mainMenuSelection).Kind()\n" +
				"\tif xType == reflect.String {\n" +
				"\t\t fmt.Println(\"\\t\\\"== reflect.String\\\" has shown that mainMenuSelection is a variable (var)\\n\")\n" +
				"\t}\n" +
				"\tif xType != reflect.String {\n" +
				"\t\t fmt.Println(\"\\t\\\"!= reflect.String\\\" has shown that mainMenuSelection is a constant (const)\\n\")\n" +
				"\t}\n\n" +
				"It should print: \n\n\t\"== reflect.String\" has shown that mainMenuSelection is a variable\n\n" +
				"\nAnd, it actually prints: \n\n")
			//
			xType := reflect.TypeOf(mainMenuSelection).Kind()
			if xType == reflect.String {
				fmt.Println("\t\"== reflect.String\" has shown that mainMenuSelection is a variable\n\n")
			}
			if xType != reflect.String {
				fmt.Println("\t\"!= reflect.String\" has shown that mainMenuSelection is probably a constant\n\n")
			}
			///
		*/
		///
		if mainMenuSelection == "1" {
			selectedExorcise := "Romaji_Prompt" // 1
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_Romaji_instructions_BeginExorcise(selectedExorcise) // 1
		} else if mainMenuSelection == "2" {
			selectedExorcise := "Romaji+Kata_Prompt" // 2
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_Romaji_plus_Kata_instructions_BeginExorcise(selectedExorcise) // 2
		} else if mainMenuSelection == "3" {
			selectedExorcise := "Kata_Prompt-Respond-w-Hira|Romaji" // 3 or 4
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_KataExorciseInstructions_BeginExorcise(selectedExorcise) // 3 or 4
		} else if mainMenuSelection == "5" {
			selectedExorcise := "drillLines"
			log_to_JapLog_file_inception_time(selectedExorcise)
			drillLines() // 5
		} else if mainMenuSelection == "6" {
			selectedExorcise := "Hira_prompt" // 6
			log_to_JapLog_file_inception_time(selectedExorcise)
			display_Hira_instructions_BeginExorcise(selectedExorcise) // 6
		} else if mainMenuSelection == "7" {
			selectedExorcise := "Most_Difficult" // 7
			//log_to_JapLog_file_inception_time(selectedExorcise)
			display_Difficult_instructions_BeginExorcise(selectedExorcise) // 7
		} else if mainMenuSelection == "8" {
			selectedExorcise := "Sequential-Kata" // 8
			//log_to_JapLog_file_inception_time(selectedExorcise)
			display_Seq_instructions_BeginExorcise(selectedExorcise) // 8
		} else if mainMenuSelection == "9" {
			selectedExorcise := "Sequential-Hira" // 9
			//log_to_JapLog_file_inception_time(selectedExorcise)
			display_Seq_instructions_BeginExorcise(selectedExorcise) // 8
		} else if mainMenuSelection == "exit" || mainMenuSelection == "quit" {
			os.Exit(1)
		}
	}
}

// 1
//
//goland:noinspection ALL
func display_Romaji_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Romaji_instructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Romaji_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji prompt\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Romaji_instructions() {
	body_of_Romaji_instructions()
}

// 2
//
//goland:noinspection ALL
func display_Romaji_plus_Kata_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Romaji_plus_Kata_instructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Romaji_plus_Kata_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji-Katakana prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji-Katakana prompt\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji-Katakana prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Romaji_plus_Kata_instructions() {
	body_of_Romaji_plus_Kata_instructions()
}

// 3 and 4
//
//goland:noinspection ALL
func display_KataExorciseInstructions_BeginExorcise(selectedExorcise string) {
	body_of_KataExorciseInstructions()
	TouchTypingExorcise(selectedExorcise)
}

//goland:noinspection ALL
func body_of_KataExorciseInstructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing Katakana chars: Using Hiragana-input-mode on your system ...")
	fmt.Println("... enter the Hiragana chars that correspond to the Katakana.\n")
	fmt.Println("Or, alternatively, type the Romaji that corresponds to the Katakana\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type either the Hiragana or Romaji that corresponds to the Katakana prompt: \n")
}

//goland:noinspection ALL
func re_display_KataExorciseInstructions() {
	body_of_KataExorciseInstructions()
}

// 6
//
//goland:noinspection ALL
func display_Hira_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Hira_instructions()
	TouchTypingExorcise(selectedExorcise)
}
func display_Seq_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Hira_instructions()
	TouchTypingExorciseSequential(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Hira_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Printf("力いまのレーまじ\n\n" +
		"ひらがなのプロンプトに答える形でローマ字のタッチタイピングの練習をしましょう。\n\n" +
		"ひらがなのプロンプトに対応するローマ字をタイプしてください。\n\n" +
		"(米語入力モードまたはアルファベット/数字入力モードを使用):" +
		"\n\nメインメニューに戻るには「menu」と入力してください。\n\n" +
		"利用可能なディレクティブのメニューを再表示するには「dir」と入力してください。\n\n" +
		"ローマ字の規則に関する背景情報を見るには「notes」と入力してください。 \n\n" +
		"コンテキストセンシティブヘルプの場合は「?」と入力してください。\n\n" +
		"特定のひらがな文字のヘルプの場合は「??」と入力してください。\n\n" +
		"プロンプトと「キー」をリセットするには、「set」と入力してください。\n\n" +
		"現在のセッションでこれまでに行ったことを表示するには、「stat」と入力してください。\n\n" +
		"ヒットログをリセットするには、「reset」と入力してください。\n\n" +
		"アプリを終了するには、「exit」または「quit」と入力してください。\n\n" +
		"\n" +
		"ここでは、以下の点に注意して翻訳しました。\n\n" +
		"「プロンプト」は「prompt」と訳しました。\n" +
		"「キー」は「keys」と訳しました。\n" +
		"「アプリ」は「app」と訳しました。\n" +
		"「?」と「??」は、それぞれ「context-sensitive help」と「help with a specific hiragana character」と訳しました。\n" +
		"「ヒットログ」は「hit log」と訳しました。\n" +
		"また、以下の点についても、注意事項として追加しました。\n\n" +
		"ひらがなを入力する際は、米語入力モードまたはアルファベット/数字入力モードを使用してください。\n" +
		"プロンプトと「キー」をリセットするには、「set」と入力してください。\n" +
		"現在のセッションでこれまでに行ったことを表示するには、「stat」と入力してください。\n" +
		"ヒットログをリセットするには、「reset」と入力してください。\n" +
		"\n\n")
	fmt.Printf("Let's practice typing Romanization by responding to hiragana prompts.\n\n" +
		"ひらがなを「hiragana」と、ローマ字を「Romanization」と訳しました。\n" +
		"Type the Romanization corresponding to the hiragana prompt.\n\n" +
		"(Use US keyboard input mode or alphabet/number input mode):\n\n" +
		"To return to the main menu, type \"menu\".\n\n" +
		"To display the menu of available directives again, type \"dir\".\n\n" +
		"To see background information on Romanization rules, type \"notes\".\n\n" +
		"For context-sensitive help, type \"?\".\n\n" +
		"For help with a specific hiragana character, type \"??\", e.g. \"??あ\".\n\n" +
		"To reset the prompt and \"keys\", type \"set\".\n\n" +
		"To display what you have done so far in the current session, type \"stat\".\n\n" +
		"To reset the hit log, type \"reset\".\n\n" +
		"To exit the app, type \"exit\" or \"quit\".")
	//
	fmt.Println("Practicing touch-typing (TT) Romaji in response to Hiragana prompts:\n")
	fmt.Println("Type the Romaji corresponding to the Hiragana prompt\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type the Romaji corresponding to the Hiragana prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Hira_instructions() {
	body_of_Hira_instructions()
}

// 7
//
//goland:noinspection ALL
func display_Difficult_instructions_BeginExorcise(selectedExorcise string) {
	body_of_Difficult_instructions()
	TouchTypingExorciseDifficult(selectedExorcise)
}

//goland:noinspection ALL
func body_of_Difficult_instructions() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing Katakana chars: Using Hiragana-input-mode on your system ...")
	fmt.Println("... enter the Hiragana chars that correspond to the Katakana.\n")
	fmt.Println("Or, alternatively, type the Romaji that corresponds to the Katakana\n")
	display_ListingOf_OptionsThese_AllHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type either the Hiragana or Romaji that corresponds to the Katakana prompt: \n")
}

//goland:noinspection ALL
func reDisplay_Difficult_instructions() {
	body_of_Difficult_instructions()
}

// This is executed each time menu is given as a dir by the user during any activity
// Things to do after an activity, and before beginning another activity
func do_betweenMainMenuSelectionsTTE(selectedExorcise string) {
	currentTime := time.Now()
	if selectedExorcise == "Romaji_Prompt" { // 1
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 1 'Romaji_Prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExorcise == "Romaji+Kata_Prompt" { // 2
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 2 'Romaji+Kata_Prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExorcise == "Kata_Prompt-Respond-w-Hira|Romaji" { // 3, 4
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 3 or 4 'Kata_Prompt-Respond-w-Hira|Romaji' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExorcise == "drillLines" {
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 5 'Drill Lines' occured at: %s \n",
			currentTime.Format("01-02-2006 15:04:05 Monday"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExorcise == "Hira_prompt" { // 3, 4
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nTransition from exorcise 6 'Hira_prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	}
}

/*
.
.
*/
