package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"strconv"

	"github.com/jpillora/opts"
	log "github.com/sirupsen/logrus"
	shop "github.com/theovassiliou/shopshop/basket"
)

//set this via ldflags (see https://stackoverflow.com/q/11354518)
const pVersion = ".3"

// version is the current version number as tagged via git tag 1.0.0 -m 'A message'
var (
	version = "1.0" + pVersion + "-src"
	commit  string
	branch  string
)

type config struct {
	DropBoxDir string    `help:"Directory with ShopShop lists"`
	ListName   string    `help:"ShopShop listname"`
	LogLevel   log.Level `help:"Log level, one of panic, fatal, error, warn or warning, info, debug, trace"`
}

const shortUsage = "Adding, modifying and deleting items from a ShopShop list"

var conf config

type rm struct {
	Indices []string `type:"arg" name:"Index" help:"item to remove"`
}

const rmUsage = "Removes an item at index position from list"

func (cmd *rm) Run() {
	shoppingList.Remove(cmd.Indices)
	shoppingList.Save()
}

type add struct {
	ItemDescription []string `type:"arg" name:"description" help:"item to add"`
	Quantity        string   `type:"flag"`
}

const addUsage = "Adds an item to the shopping list"

func (cmd *add) Run() {
	shoppingList.AddItem(cmd.Quantity, cmd.ItemDescription)
	shoppingList.Save()
}

func isQuantity(word string) (bool, string) {
	// quantity has the form 400g or 400 or 2cl
	// but not 400g500 or gr400

	for i := len(word); i > 0; i-- {
		if _, err := strconv.Atoi(word[:i]); err == nil {
			return true, word
		}
		if _, err := strconv.Atoi(word[i-1:]); err == nil {
			return false, ""
		}
	}

	return false, ""
}

func execute(sb *shop.Basket, words []string) *shop.Basket {
	cmd := words[0]
	switch cmd {
	case "rm", "remove":
		(&rm{Indices: words[1:]}).Run()
		(&ls{}).Run()
	case "add", "buy":
		quantity := ""
		i := 1
		if isQuantity, q := isQuantity(words[1]); isQuantity {
			i = 2
			quantity = q
		}
		(&add{ItemDescription: words[i:], Quantity: quantity}).Run()
		(&ls{}).Run()
	case "query", "q":
		if len(words) > 1 {
			newListFilename := path.Join(os.ExpandEnv(conf.DropBoxDir), words[1]+".shopshop")
			fi, err := os.Open(newListFilename)
			b, err := ioutil.ReadAll(fi)

			shop.AssertNoError(err)

			newSl := new(shop.Basket)
			err = json.Unmarshal(b, newSl)
			shop.AssertNoError(err)
			newSl.SetFileName(newListFilename)
			return newSl
		}
		(&query{}).Run()

	case "checkout", "co":
		(&co{}).Run()
		(&ls{}).Run()
	case "list", "ls":
		(&ls{}).Run()
	case "help":
		fmt.Println(`Commands:
  add [#] ...   add [quantity] item
  rm # [#]+     remove item(s) at index #
  ls			list items in list
  co            checkout (remove done items)
  query [#]     query for all lists or change to list #`)
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Use 'help' for help")
		(&ls{}).Run()
	}
	return sb
}

type interact struct {
}

const interactUsage = "interactive mode"

func (c *interact) Run() {

	out := os.Stdout
	reader := bufio.NewReader(os.Stdin)

	for {
		out.WriteString("> ")
		switch line, err := reader.ReadString('\n'); err {
		case nil:
			if len(line) < 2 {
				os.Exit(0)
			}
			l := strings.Split(line[:len(line)-1], " ")
			if len(l) > 0 {
				shoppingList = execute(shoppingList, l)
			}
		case io.EOF:
			fmt.Println()
			os.Exit(0)
		default:
			panic(err)
		}

	}
}

type ls struct{}

const lsUsage = "List the shopshop list"

func (cmd *ls) Run() {
	shoppingList.List()
}

type query struct{}

const queryUsage = "Shows a list of all shopping lists"

func (cmd *query) Run() {
	shoppingList.Query(conf.DropBoxDir)
}

type co struct{}

const coUsage = "Checkout (removes done items) from list"

func (cmd *co) Run() {
	shoppingList.Checkout()
	shoppingList.Save()
	shoppingList.List()
}

var fileName string
var shoppingList *shop.Basket

// FormatFullVersion formats for a cmdName the version number based on version, branch and commit
func FormatFullVersion(cmdName, version, branch, commit string) string {
	var parts = []string{cmdName}

	if version != "" {
		parts = append(parts, version)
	} else {
		parts = append(parts, "unknown")
	}

	if branch != "" || commit != "" {
		if branch == "" {
			branch = "unknown"
		}
		if commit == "" {
			commit = "unknown"
		}
		git := fmt.Sprintf("(git: %s %s)", branch, commit)
		parts = append(parts, git)
	}

	return strings.Join(parts, " ")
}

func main() {

	conf = config{
		LogLevel:   log.DebugLevel,
		DropBoxDir: "$HOME/Dropbox/ShopShop/",
		ListName:   "Grocery",
	}

	//parse config
	cmd := opts.
		New(&conf).
		Summary(shortUsage).
		PkgRepo().
		Version(FormatFullVersion("shopshop", version, branch, commit)).
		AddCommand(
			opts.New(&add{}).
				Summary(addUsage)).
		AddCommand(
			opts.New(&ls{}).
				Summary(lsUsage)).
		AddCommand(
			opts.New(&query{}).
				Summary(queryUsage)).
		AddCommand(
			opts.New(&rm{}).
				Summary(rmUsage)).
		AddCommand(
			opts.New(&co{}).
				Summary(coUsage)).
		AddCommand(
			opts.New(&interact{}).
				Summary(interactUsage)).Parse()

	fileName = path.Join(os.ExpandEnv(conf.DropBoxDir), conf.ListName+".shopshop")

	log.SetLevel(conf.LogLevel)
	fi, err := os.Open(fileName)
	b, err := ioutil.ReadAll(fi)

	shop.AssertNoError(err)

	shoppingList = shop.NewBasket()

	err = json.Unmarshal(b, shoppingList)
	shop.AssertNoError(err)
	shoppingList.SetFileName(fileName)

	if cmd.IsRunnable() {
		cmd.Run()
	} else {
		shoppingList.List()
	}
}
