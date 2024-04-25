package asciiart

import (
	"os"
	"testing"
)

func readFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(data)
	//       
}


func TestDisplayArt2(t *testing.T) {
	test1 := []struct {
		input    string
		expected string
	}{ {
		"1Hello 2There",
		readFile("test1.txt"),
	},
		
	}

	for _, test := range test1 {
		filePath := "../standard.txt"
		output, err := DisplayArt(filePath, test.input)
		if err != nil {
			t.Errorf("Error %s", err)
		}
		if output != test.expected {
			os.WriteFile("test2.txt", []byte(output), 0644)
			t.Errorf("DisplayArt(%s), expexted(\n%s\n), got(\n%s\n)", test.input, test.expected, output)
		}
	}

}

// func TestDisplayArt(t *testing.T) {
// 	test1 := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{
// 			"Hello\n",
// 			`
// 		 _    _          _   _          
// 		| |  | |        | | | |         
// 		| |__| |   ___  | | | |   ___   
// 		|  __  |  / _ \ | | | |  / _ \  
// 		| |  | | |  __/ | | | | | (_) | 
// 		|_|  |_|  \___| |_| |_|  \___/  
										
										
// 	`,
// 		},
// 		{
// 			"hello",
// 			`
// 	 _              _   _          
// 	| |            | | | |         
// 	| |__     ___  | | | |   ___   
// 	|  _ \   / _ \ | | | |  / _ \  
// 	| | | | |  __/ | | | | | (_) | 
// 	|_| |_|  \___| |_| |_|  \___/  
								   
								   
// 	`,
// 		},

// 		{
// 			"HeLlO",
// 			`
// 	 _    _          _        _    ____   
// 	| |  | |        | |      | |  / __ \  
// 	| |__| |   ___  | |      | | | |  | | 
// 	|  __  |  / _ \ | |      | | | |  | | 
// 	| |  | | |  __/ | |____  | | | |__| | 
// 	|_|  |_|  \___| |______| |_|  \____/  
										  
										  
// 	`,
// 		},

// 		{
// 			"Hello There",

// 			`
// 	 _    _          _   _               _______   _                           
// 	| |  | |        | | | |             |__   __| | |                          
// 	| |__| |   ___  | | | |   ___          | |    | |__     ___   _ __    ___  
// 	|  __  |  / _ \ | | | |  / _ \         | |    |  _ \   / _ \ | '__|  / _ \ 
// 	| |  | | |  __/ | | | | | (_) |        | |    | | | | |  __/ | |    |  __/ 
// 	|_|  |_|  \___| |_| |_|  \___/         |_|    |_| |_|  \___| |_|     \___| 
																			   
											   
// 	`,
// 		},

// 		{"1Hello 2There",
// `
// 		_    _          _   _                      _______   _                           $
// 	_  | |  | |        | | | |              ____  |__   __| | |                          $
//    / | | |__| |   ___  | | | |   ___       |___ \    | |    | |__     ___   _ __    ___  $
//    | | |  __  |  / _ \ | | | |  / _ \        __) |   | |    |  _ \   / _ \ | '__|  / _ \ $
//    | | | |  | | |  __/ | | | | | (_) |      / __/    | |    | | | | |  __/ | |    |  __/ $
//    |_| |_|  |_|  \___| |_| |_|  \___/      |_____|   |_|    |_| |_|  \___| |_|     \___| $
// 																						 $
// 																						 $
// `,
// 		},
// 	}

// 	for _, test := range test1 {
// 		filePath := "../standard.txt"
// 		output, err := DisplayArt(filePath, test.input)
// 		if err != nil {
// 			t.Errorf("Error %s", err)
// 		}
// 		if output != test.expected {
// 			t.Errorf("DisplayArt(%s), expexted(%s), got(%s)", test.input, test.expected, output)
// 		}
// 	}
// }
