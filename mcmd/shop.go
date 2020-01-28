package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"strconv"

	"github.com/jpillora/opts"
	log "github.com/sirupsen/logrus"
	shop "github.com/theovassiliou/shopshop"
)

//set this via ldflags (see https://stackoverflow.com/q/11354518)
var version = ".1"

// VERSION is the current version number.
var VERSION = "0.0" + version + "-src"

type config struct {
	DropBoxDir  string    `help:"Directory with ShopShop lists"`
	FileName    string    `help:"ShopShop filename, without dir"`
	Interactive bool      `help:"Start in interactive mode"`
	LogLevel    log.Level `help:"Log level, one of panic, fatal, error, warn or warning, info, debug, trace"`
}

const shortUsage = "Adding, modifying and deleting items from a ShopShop list"

var conf config

func assertNoError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func save(sl *shop.Basket) {
	shoppingJSON, _ := json.MarshalIndent(sl, "", "  ")
	err := ioutil.WriteFile(fileName, shoppingJSON, 0644)
	assertNoError(err)
}

type add struct {
	ItemDescription []string `type:"arg" name:"description" help:"item to add"`
	Quantity        string   `type:"flag"`
}

const addUsage = "Adds an item to the shopping list"

func (cmd *add) Run() {
	count := cmd.Quantity
	buf := bytes.NewBuffer(nil)
	for _, word := range cmd.ItemDescription[0:] {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(word)
	}
	name := buf.String()
	shoppingList.ShoppingList = append(shoppingList.ShoppingList, shop.Item{Done: false, Count: count, Name: name})
	fmt.Println("Adding:", count, name)
	save(shoppingList)
}

type ls struct{}

const lsUsage = "List the shopshop list"

func (cmd *ls) Run() {
	fmt.Println("Items:")
	for i, item := range shoppingList.ShoppingList {
		check := " "
		if item.Done {
			check = "@done"
		}
		fmt.Printf("%2d: %s %s %s\n", i, item.Count, item.Name, check)
	}
	fmt.Println()
}

type co struct{}

const coUsage = "Checkout (removes done items) from list"

func (cmd *co) Run() {
	var newList []shop.Item
	for _, item := range shoppingList.ShoppingList {
		if !item.Done {
			newList = append(newList, item)
		}
	}
	shoppingList.ShoppingList = newList
	save(shoppingList)
}

type rm struct {
	Index string `type:"arg" name:"Index" help:"item to remove"`
}

const rmUsage = "Removes an item at index position from list"

func (cmd *rm) Run() {
	idx, err := strconv.Atoi(cmd.Index)
	if err != nil {
		fmt.Println(err)
		return
	}
	item := shoppingList.ShoppingList[idx]
	fmt.Println("Removing:", item.Count, item.Name)
	shoppingList.ShoppingList = append(shoppingList.ShoppingList[:idx], shoppingList.ShoppingList[idx+1:]...)
	save(shoppingList)
}

var fileName string
var shoppingList *shop.Basket

func main() {

	conf = config{
		LogLevel:   log.DebugLevel,
		DropBoxDir: "$HOME/Dropbox/ShopShop/",
		FileName:   "Lidl.shopshop",
	}

	//parse config
	cmd := opts.New(&conf).
		Summary(shortUsage).
		Repo("github.com/theovassiliou/shopshop").
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
		Parse()

	fileName = path.Join(os.ExpandEnv(conf.DropBoxDir), conf.FileName)

	log.SetLevel(conf.LogLevel)
	fi, err := os.Open(fileName)
	b, err := ioutil.ReadAll(fi)

	assertNoError(err)

	shoppingList = new(shop.Basket)
	err = json.Unmarshal(b, shoppingList)
	assertNoError(err)
	cmd.Run()

	save(shoppingList)

	(&ls{}).Run()
}
