# Shop

A simple go program to update [ShopShop][1] shopping lists.

[ShopShop][1] is a free iOS shopping list program that we are using in our family for many years now.
In particular the synching via Dropbox is a great help.

ShopShop stores it's data in traditional JSON format, and the files can be located in a DropBox folder.
While sitting on the desktop, being notified by the DropBox message that a shopping list is being updated, with shopshop-cl you can now easily check and update shopping lists.

To use, you should have a `$HOME/Dropbox/ShopShop/Shopping List.shopshop` file, created by ShopShop.

WARNING: THIS SOFTWARE CAN'T BE ERROR FREE, SO USE IT AT YOUR OWN RISK. DON'T USE IT IF YOU HAVEN'T MADE AN ACTUAL BACKUP COPY OF YOUR FIBARO HC2 SYSTEM. IF YOU DO NOT HOW TO DO THIS, PLEASE RECONSIDER TO USE THIS SOFTWARE ANYWAY. I HAVE DONE MY BEST TO MAKE SURE THAT THE TOOLS BEHAVE AS EXPECTED. BUT AGAIN ... USE IT AT YOUR OWN RISK. I AM NOT GIVING ANY KIND OF WARRANTY, NEITHER EXPLICITELY NOR IMPLICITELY.

[1]: https://itunes.apple.com/us/app/shopshop-shopping-list/id288350249?mt=8

## Installation binaries

You can download the binaries directly from the [releases](https://github.com/theovassiliou/shopshop/releases) section.  Unzip/untar the downloaded archive and copy the files to a location of your choice, e.g. `/usr/local/bin/` on *NIX or MacOS. If you install only the binaries, make sure that they are accessible from the command line. Ideally, they are accessible via `$PATH` or `%PATH%`, respectively.

## Example

```shell
$ shopshop
Items in:  Grocery
 0:  Milk  
 1:  Honey @done
 2:  Butter  
 3: 3 Soda
 ```

 The list "Grocery" has 4 items. "Honey" is marked as done and the quantity of "Soda" is set to three

 ```shell
shopshop -h

  Usage: shopshop [options] <command>

  Adding, modifying and deleting items from a ShopShop list

  Options:
  --drop-box-dir, -d  Directory with ShopShop lists (default $HOME/Dropbox/ShopShop/)
  --list-name, -l     ShopShop listname (default Groccery)
  --log-level         Log level, one of panic, fatal, error, warn or warning, info, debug, trace
                      (default debug)
  --version, -v       display version
  --help, -h          display help

  Commands:
  · rm        Removes an item at index position from list
  · co        Checkout (removes done items) from list
  · interact  interactive mode
  · add       Adds an item to the shopping list
  · ls        List the shopshop list

  Version:
    shopshop 1.0.0 (git: master 824843e)
```

## Usage

```shell
    shopshop
```

Will list current items.

```shell
    shopshop add green cheese
```

Will add "green cheese" to your shopping list.

```shell
    shopshop add 5 milk
```

Will add "milk" with a quantitiy of 5 to your shopping list.

```shell
    shopshop rm 3
```

Will remove the item at index 3.

```shell
    shopshop co
```

Will remove items marked as done.

```shell
    shopshop -h
```

Will print the help.

```shell
    shopshop -interact
```

Launch shopshop in interactive mode. Use "help" to find out more

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/theovassiliou/shopshop/tags).

## Authors

* **Theo Vassiliou** - *Complete recoding* - [Theo Vassiliou](https://github.com/theovassiliou)
* **Steve Dunham** - *Inital idea and work* [Steve Dunham](https://github.com/dunhamsteve)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

Thanks to all the people out there that produce amazing open-source software, which supported the creation of this piece of software. In particular I wasn't only able to use libraries etc. But also, to learn and understand golang better. In particular I wanted to thank

* [Nikolaj Schumacher](https://apps.apple.com/de/app/shopshop-einkaufsliste/id288350249) for ShoPShop, the incredible usefull small little iOS app.
* [Jaime Pillora](https://github.com/jpillora) for [jpillora/opts](https://github.com/jpillora/opts). Nice piece of work!
* [Simon Eskildsen](https://github.com/sirupsen) for  [sirupsen/logrus](https://github.com/sirupsen/logrus), which I continously use.
* [PurpleBooth](https://gist.github.com/PurpleBooth) for the well motivated [README-template](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)
* [Steve Dunham](https://github.com/dunhamsteve) for the initial idea and structure of the golang Shop.
