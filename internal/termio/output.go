package termio

import (
	"fmt"
	"golang.design/x/clipboard"
	"golang.org/x/term"
	"os"
	"text/tabwriter"
)

const padding = 3

func init() {
	err := clipboard.Init()
	if err != nil {
		fmt.Println("Error initializing clipboard", err)
		os.Exit(1)
	}
}

func ClearScreen() {
  fmt.Print("\033[2J")
}

func Clip(text string) {
	clipboard.Write(clipboard.FmtText, []byte(text))
}

func PrintList(files []string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	numcolumns := getNumOfColumns()
	fmt.Println("")

	var column = 1
	for i, file := range files {
		fragment := fmt.Sprintf("%d: %s\t", i+1, file)

		writer.Write([]byte(fragment))
		if column == numcolumns {
			writer.Write([]byte("\n"))
			column = 1
		} else {
			column += 1
		}

		if i == len(files)-1 && column < 4 {
			writer.Write([]byte("\n"))
		}
	}
	writer.Flush()
}

func PrintLoopHelp(helps []string) {
	fmt.Println("====================HELP====================")
	for _, help := range helps {
		fmt.Println(help)
	}
	fmt.Println("============================================")
}

func getNumOfColumns() int {
	fd := int(os.Stdin.Fd())
	width, _, _ := term.GetSize(fd)
	if width < 90 {
		return 1
	} else if width < 120 {
		return 2
	} else {
		return 3
	}
}
