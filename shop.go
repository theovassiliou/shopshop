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
var version = ".1"

// VERSION is the current version number.
var VERSION = "0.0" + version + "-src"

type config struct {
	DropBoxDir string    `help:"Directory with ShopShop lists"`
	FileName   string    `help:"ShopShop filename, without dir"`
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

func execute(shoppingList *shop.Basket, line []string) {
	cmd := line[0]
	switch cmd {
	case "rm", "remove":
		(&rm{Indices: line[1:]}).Run()
	case "add", "buy":
		count := ""
		i := 1
		if _, err := strconv.Atoi(line[1]); err == nil {
			i = 2
			count = line[1]
		}
		(&add{ItemDescription: line[i:], Quantity: count}).Run()
	case "checkout", "co":
		(&co{}).Run()
	case "list", "ls":
		(&ls{}).Run()
	case "help":
		fmt.Println(`Commands:
  add [#] ...   add [quantity] item
  rm # [#]+     remove item(s) at index #
  co            checkout (remove done items)`)
	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Use 'help' for help")
		return
	}
	(&ls{}).Run()
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
				execute(shoppingList, l)
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

type co struct{}

const coUsage = "Checkout (removes done items) from list"

func (cmd *co) Run() {
	shoppingList.Checkout()
	shoppingList.Save()
}

var fileName string
var shoppingList *shop.Basket

func main() {

	conf = config{
		LogLevel:   log.DebugLevel,
		DropBoxDir: "$HOME/Dropbox/ShopShop/",
		FileName:   "Lidl2.shopshop",
	}

	//parse config
	cmd := opts.
		New(&conf).
		Summary(shortUsage).
		PkgRepo().
		Version(VERSION).
		AddCommand(
			opts.New(&add{}).
				Summary(addUsage)).
		AddCommand(
			opts.New(&ls{}).
				Summary(lsUsage)).
		AddCommand(
			opts.New(&rm{}).
				Summary(rmUsage)).
		AddCommand(
			opts.New(&co{}).
				Summary(coUsage)).
		AddCommand(
			opts.New(&interact{}).
				Summary(interactUsage)).Parse()

	fileName = path.Join(os.ExpandEnv(conf.DropBoxDir), conf.FileName)

	log.SetLevel(conf.LogLevel)
	fi, err := os.Open(fileName)
	b, err := ioutil.ReadAll(fi)

	shop.AssertNoError(err)

	shoppingList = new(shop.Basket)
	err = json.Unmarshal(b, shoppingList)
	shop.AssertNoError(err)
	shoppingList.SetFileName(fileName)

	cmd.Run()

	shoppingList.Save()

	(&ls{}).Run()
}
