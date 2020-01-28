package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/jpillora/opts"
	log "github.com/sirupsen/logrus"
	shop "github.com/theovassiliou/shopshop"
)

//set this via ldflags (see https://stackoverflow.com/q/11354518)
var version = ".1"

// VERSION is the current version number.
var VERSION = "0.0" + version + "-src"

type config struct {
	Cmd         []string  `type:"arg" name:"command" help:"command, one of add, remove/em, co (checkout, remove done items from list)"`
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

func process(shoppingList *shop.Basket, line []string) {
	cmd := line[0]
	switch cmd {
	case "rm", "remove":
		idx, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		item := shoppingList.ShoppingList[idx]
		fmt.Println("Removing:", item.Count, item.Name)
		shoppingList.ShoppingList = append(shoppingList.ShoppingList[:idx], shoppingList.ShoppingList[idx+1:]...)
		save(shoppingList)
	case "add", "buy":
		count := ""
		buf := bytes.NewBuffer(nil)
		i := 1
		if _, err := strconv.Atoi(line[1]); err == nil {
			i = 2
			count = line[1]
		}
		for _, word := range line[i:] {
			if buf.Len() > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteString(word)
		}
		name := buf.String()
		shoppingList.ShoppingList = append(shoppingList.ShoppingList, shop.Item{false, count, name})
		fmt.Println("Adding:", count, name)
		save(shoppingList)
	case "checkout", "co":
		var newList []shop.Item
		for _, item := range shoppingList.ShoppingList {
			if !item.Done {
				newList = append(newList, item)
			}
		}
		shoppingList.ShoppingList = newList
		save(shoppingList)
	case "help":
		fmt.Println(`Commands:
  add ...  add item
  rm #     remove item at index
  co       checkout (remove done items)`)
	case "list", "ls":
		fmt.Println("Items:")
		for i, item := range shoppingList.ShoppingList {
			check := " "
			if item.Done {
				check = "@done"
			}
			fmt.Printf("%2d: %s %s %s\n", i, item.Count, item.Name, check)
		}
		fmt.Println()
	default:
		fmt.Println("Unknown command:", cmd)
	}
	fmt.Println()

}

func interact(shoppingList *shop.Basket) {
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
				process(shoppingList, l)
			}
		case io.EOF:
			fmt.Println()
			os.Exit(0)
		default:
			panic(err)
		}
	}
}

var fileName string

func main() {

	conf = config{
		LogLevel:   log.DebugLevel,
		DropBoxDir: "$HOME/Dropbox/ShopShop/",
		FileName:   "Lidl.shopshop",
	}

	//parse config
	opts.New(&conf).
		Summary(shortUsage).
		Repo("github.com/theovassiliou/shopshop").
		Version(VERSION).
		Parse()

	fileName = path.Join(os.ExpandEnv(conf.DropBoxDir), conf.FileName)

	log.SetLevel(conf.LogLevel)
	fi, err := os.Open(fileName)
	b, err := ioutil.ReadAll(fi)

	assertNoError(err)

	shoppingList := new(shop.Basket)
	err = json.Unmarshal(b, shoppingList)
	assertNoError(err)

	if len(conf.Cmd) > 0 {
		process(shoppingList, conf.Cmd)
		process(shoppingList, []string{"ls"})
	} else if conf.Interactive {
		interact(shoppingList)
	} else {
		process(shoppingList, []string{"ls"})
	}

	save(shoppingList)

}
