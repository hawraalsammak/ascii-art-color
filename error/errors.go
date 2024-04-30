package errormessages

//this is the place you can find all errors you might want to print

const Error1 string = "\033[31mPlease provide text to generate ASCII Art!\033[0m"
const Error2 string = "Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\""
const Error3 string = "\033[31mToo many arguments!\033[0m"
const Error4 string = "\033[31mColor not recognized, using deafult!\033[0m\nList of available colors in README.md"

// go run . --color=red h hello
// go run . --color=red hello
// go run . hello
