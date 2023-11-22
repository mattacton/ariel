package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/mattacton/ariel/internal/collection"
	"github.com/mattacton/ariel/internal/curl"
	"github.com/mattacton/ariel/internal/request"
	"github.com/mattacton/ariel/internal/stack"
	"github.com/mattacton/ariel/internal/termio"
	"github.com/rbretecher/go-postman-collection"
)

type shortCircuitCode int
type selectionlevel int
type Selection struct {
	level selectionlevel
	input string
}
type ChooseOne struct {
	prompt             string
	headerToast, toast *string
	items              []string
}

const (
	HELP shortCircuitCode = iota
	EXIT
	REFRESH
	BACK
	RESTART
	NONE
)
const (
	COLLECTION selectionlevel = iota
	ITEM
)

var validInputPattern = regexp.MustCompile(`^[0-9]+|[\\?qbrc]{1}$`)

var helps = []string{
	"<num> - Choose an item",
	"<b> - Go back",
	"<c> - Choose a new collection",
	"<r> - Refresh",
	"<q> - Quit",
	"<?> - Help",
}
var toast string
var headerToast string

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 || len(args) > 1 {
		printUsage()
		os.Exit(1)
	}

	collections := collection.FromPaths(args)
	if len(collections) == 0 {
		fmt.Println("No collections found")
		os.Exit(0)
	}

	var selectedCollection postman.Collection
	itemsStack := stack.New[[]*postman.Items]()
	headerToast = termio.HelpText(helps)
	for {
		var items []*postman.Items
		var item *postman.Items
		var selection Selection
		if collection.IsZeroed(selectedCollection) {
			input := chooseAnItem(ChooseOne{
				"Choose a collection: ",
				&headerToast,
				&toast,
				collection.CollectionNames(collections)})
			selection = Selection{COLLECTION, input}
		} else {
			if itemsStack.Empty() {
				itemsStack.Push(request.GetRequests(selectedCollection))
			}

			items = itemsStack.Peek()
			input := chooseAnItem(ChooseOne{
				"Choose an item: ",
				&headerToast,
				&toast,
				request.RequestNames(items)})
			selection = Selection{ITEM, input}
		}

		index, short := shortCircuit(selection.input)
		if short == EXIT {
			os.Exit(0)
		} else if short == HELP {
			headerToast = termio.HelpText(helps)
			continue
		} else if short == REFRESH {
			continue
		} else if short == RESTART {
			itemsStack = stack.New[[]*postman.Items]()
			selectedCollection = postman.Collection{}
			continue
		} else if short == BACK {
			selectedCollection = goBack(selectedCollection, itemsStack)
			continue
		}

		if selection.level == COLLECTION {
			if index > len(collections) || index < 1 {
				toast = "\n 0> Invalid collection number"
				continue
			}
			selectedCollection = collections[index-1]
			continue
		} else {
			if index > len(items) || index < 1 {
				toast = "\n 0> Invalid item number"
				continue
			}
			item = items[index-1]
			if request.IsDir(item) {
				itemsStack.Push(item.Items)
				continue
			}
		}

		toast = fmt.Sprintf("\n 0> Copied %s to clipboard", item.Name)

		termio.Clip(curl.FromRequest(item.Request))
	}
}

func shortCircuit(input string) (int, shortCircuitCode) {
	switch input {
	case "?", "help":
		return 0, HELP
	case "q":
		return 0, EXIT
	case "b":
		return 0, BACK
	case "c":
		return 0, RESTART
	case "r":
		return 0, REFRESH
	default:
		index, err := strconv.Atoi(input)
		if err != nil {
			printInvalidInput()
			return 0, REFRESH
		}
		return index, NONE
	}
}

func goBack(selectedCollection postman.Collection, itemsStack *stack.Stack[[]*postman.Items]) postman.Collection {
	if collection.IsZeroed(selectedCollection) {
		return selectedCollection
	} else {
		itemsStack.Pop()
		if itemsStack.Empty() {
			return postman.Collection{}
		}
		return selectedCollection
	}
}

func printUsage() {
	fmt.Println("usage: ariel <collection directory> | <collection.json>")
	fmt.Println("\nUse ariel to quickly copy curl commands from your Postman v2.1.0 collections to your clipboard.")
}

func printInvalidInput() {
	fmt.Println("Invalid input, choose an item number or '?' for help")
}

func chooseAnItem(chooseOne ChooseOne) string {
	termio.ClearScreen(chooseOne.headerToast)
	termio.PrintList(chooseOne.items)
	termio.PrintToast(chooseOne.toast)
	return termio.PromptRead(chooseOne.prompt, isValidInput)
}

func isValidInput(input string) bool {
	match := validInputPattern.MatchString(input)
	if !match {
		printInvalidInput()
	}
	return match
}
