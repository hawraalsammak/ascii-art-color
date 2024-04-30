package main

import (
	"nono"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Args := os.Args[1:]
	if len(Args) == 0 {
		fmt.Println("Error: no arguments.")
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
		os.Exit(0)
	} else if len(Args) > 5 {
		fmt.Println("Error: so many Arguments.")
		fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
		os.Exit(0)
	} else {
		//// --color, supStr , --str , output ,fs
		output := "" // name of file from --output=
		input := ""  // input string to be converted to ascii art
		banner := "standard.txt"
		var color string  // the color from flag --color
		toBeColored := "" // sup text to be colored
		var hasColor bool // if flag --color is found

		Art := ""

		Args, output, hasColor = CheckOutput(Args)
		if len(Args) < 1 {
			fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
			os.Exit(0)
		}
		lastIndex := len(Args) - 1
		banner = CheckBanner(Args[lastIndex])
		if len(Args) == 1 {
			input = Args[0]
			banner = CheckBanner("standard")
		} else if len(Args) == 2 {
			if !hasColor {
				input = Args[0]
				banner = CheckBanner(Args[1])
			} else {
				color = Args[0]
				input = Args[1]
				banner = CheckBanner("standard")
			}
		} else if len(Args) == 3 {
			if !hasColor {
				fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
				os.Exit(0)

			} else if banner == "NotFound" {
				color = Args[0]
				toBeColored = Args[1]
				input = Args[2]
				banner = CheckBanner("standard")
			} else {

				color = Args[0]
				input = Args[1]
			}
		} else if len(Args) == 4 {
			color = Args[0]
			toBeColored = Args[1]
			input = Args[2]
		}
		fmt.Println(output, color, toBeColored)
		////////////////////////////////////////////////////////////////
		var fileLines []string
		// input := os.Args[1]
		if len(fileLines) == 1 {
			file, err := os.Open("standard.txt")
			if err != nil {
				log.Fatal(err)
			}
			fileScanner := bufio.NewScanner(file)
			fileScanner.Split(bufio.ScanLines)
			for fileScanner.Scan() {
				fileLines = append(fileLines, fileScanner.Text())
				defer file.Close()
			}
		}
		if hasColor {
			color = asciiArt.PrintColor(color)
			if color == "" {
				log.Fatal("color not found")
			}
		}

		file, err := os.Open(banner)
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			fileLines = append(fileLines, fileScanner.Text())
		}
		file.Close()
		var fileload []int
		result := strings.ReplaceAll(input, "\\t", "   ")
		result = strings.ReplaceAll(result, "\\n", " \\n ")
		words := strings.Split(result, " \\n ")
		if hasColor {
			PrintBannersWithColors(toBeColored, color, words, fileLines)
		}else {
			for j := 0; j < len(words); j++ {
				if words[j] != "" {
					fileload = ConvertToASCIIArt(words[j])
					fileload1 := PrintLines(fileload, fileLines)
					for i := 0; i < len(fileload1); i++ {
						if strings.Join(fileload1[i], "") != "" {
							if hasColor {
								fmt.Println(color, strings.Join(fileload1[i], ""), "\033[0m")
							} else {
								fmt.Println(strings.Join(fileload1[i], ""))
							}
	
							Art = Art + strings.Join(fileload1[i], "") + "\n"
						}
					}
				} else if j != 1 && words[j] == "" {
					fmt.Println()
					Art = Art + "\n"
				}
		}
		
		}
		CreateFile(output, Art)
	}

}
func CheckOutput(Args []string) ([]string, string, bool) {
	//// --color, supStr , --str , output ,fs
	output := ""
	var hasColor bool
	for i, arg := range Args {
		if strings.HasPrefix(arg, "--output=") {
			output = strings.TrimPrefix(arg, "--output=")
			if output == "" {
				fmt.Println("usage: go run . --output=<fileName.txt> something standard")
				os.Exit(0)
			}
			Args[i] = ""
		} else if strings.HasPrefix(arg, "--") && !strings.HasPrefix(arg, "--color=") {
			fmt.Println("usage: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		} else if strings.HasPrefix(arg, "--color=") {
			hasColor = true
			Args[i] = strings.TrimPrefix(arg, "--color=")
		}
	}
	Args = CleanInput(Args)
	return Args, output, hasColor
}

func CleanInput(args []string) []string {
	var cleanArgs []string
	for _, arg := range args {
		if arg != "" {
			cleanArgs = append(cleanArgs, arg)
		}
	}
	return cleanArgs
}

func CheckBanner(banner string) string {

	switch strings.ToLower(banner) {
	case "standard":
		banner = "standard.txt"
	case "shadow":
		banner = "shadow.txt"
	case "thinkertoy":
		banner = "thinkertoy.txt"
	default:
		banner = "NotFound"
	}
	return banner
}

func ConvertToASCIIArt(input string) []int {
	var word []int
	for i := 0; i < len(input); i++ {
		char := int(rune(input[i]))
		if input[i] == '\\' && i < len(input) {
			if input[i+1] == 'n' {
				char = 127
				word = append(word, char)
				i = i + 1
			} else {
				word = append(word, char)
			}
		} else {
			word = append(word, char)
		}
	}
	return word
}

func CreateFile(file, art string) {
	os.WriteFile(file, []byte(art), 0644)
}

func PrintBannersWithColors(Str, colors string, banners, arr []string) {
	num := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			if Str == "" {
				for _, j := range ch {
					n := (j-32)*9 + 1
					fmt.Print(colors, arr[int(n)+i])
				}
			} else {
				h := 0
				count := 0
				match := false
				for _, j := range ch {
					
					if !match || count >= len(Str) {
						h = h + 1
					}
					
					check := true
					n := (j-32)*9 + 1
					for q := 0; q < len(Str); q++ {
						
							if rune(Str[q]) == j {
								
								word := ch
								if count < len(Str) {
									if  h <= ( 1+ len(word) - len(Str))  {
										if (Str == word[h-1:h+len(Str)-1] || (match && count < len(Str)) ){
											match = true
											count = count + 1
											
											fmt.Print(colors, arr[int(n)+i])
										
											check = false
										}

									}
								
									if count == len(Str) {
										count = 0
										match = false
									}
									break
								}
						
							}
					}
					if check == true{
						fmt.Print("\033[0m", arr[int(n)+i])
					}
					
				}
				count = 0
			}
			fmt.Println("\033[0m")
		}
	}
}
func PrintLines(input []int, str []string) [][]string {
	word := make([][]string, 8)
	for i := 0; i < len(input); i++ {
		input[i] = input[i] - 32
	}
	for i := 0; i < len(input); i++ {
		index := input[i]
		if index == 95 {
			word[7] = append(word[7], "")
		} else if index < 0 || index >= 96 {
			fmt.Println("-Error, unrecognized character")
			break
		}
		for j := 0; j < 8; j++ {
			if index == 95 {
				word[7] = append(word[7], "\n")
				break
			} else {
				word[j] = append(word[j], str[index*9+j+1])
			}
		}
	}
	return word
}


