package asciiArt

import "strings"

func PrintColor(c string) string {
	var colorCode string
	c = strings.ToLower(c)
	switch c {
	case "red":
		colorCode = "\033[31m" // Red color
	case "green":
		colorCode = "\033[32m" // Green color
	case "yellow":
		colorCode = "\033[33m" // Yellow color
	case "blue":
		colorCode = "\033[34m" // Blue color
	case "orange":
		colorCode = "\033[38;5;208m" // Orange color
	case "purple":
		colorCode = "\033[35m" // Purple color
	case "pink":
		colorCode = "\033[95m" // Pink color
	default:
		colorCode = "" // Default color code (empty string)
	}
	// fmt.Println(colorCode)
	return colorCode
}
