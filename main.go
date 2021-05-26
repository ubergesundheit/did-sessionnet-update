package main

import (
	"fmt"
	"os"

	colly "github.com/gocolly/colly/v2"
)

// 2021-03-23:
// Somacos GmbH & Co. KG,https://www.somacos.de, SessionNet Version 5.1.8 bi (Layout 5)
// Somacos GmbH & Co. KG,https://www.somacos.de, SessionNet Version 5.2.3 KP3 bi (Layout 5)
// Version changed (from/to)
// V:050108
// V:050203

var expectedAuthor = "Somacos GmbH & Co. KG,https://www.somacos.de, SessionNet Version 5.2.3 KP3 bi (Layout 5)"
var expectedVersion = "V:050203"

func checkVersion() {
	url := "https://www.stadt-muenster.de/sessionnet/sessionnetbi/info.php"
	exitCode := 0

	c := colly.NewCollector()

	c.OnHTML("meta[name=author]", func(e *colly.HTMLElement) {
		newAuthor := e.Attr("content")

		if newAuthor != expectedAuthor {
			fmt.Printf("Author changed (from/to)\n%s\n%s\n", expectedAuthor, newAuthor)
			exitCode = 1
		}
	})

	c.OnHTML("meta[name=sessionnet]", func(e *colly.HTMLElement) {
		newVersion := e.Attr("content")

		if newVersion != expectedVersion {
			fmt.Printf("Version changed (from/to)\n%s\n%s\n", expectedVersion, newVersion)
			exitCode = 1
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnScraped(func(_ *colly.Response) {
		os.Exit(exitCode)
	})

	c.Visit(url)
}

func checkOparl(path string) {
	url := "https://oparl.stadt-muenster.de" + path

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("'%s' is not available. Thats good %v\n", url, err)
	})

	c.OnScraped(func(_ *colly.Response) {
		fmt.Printf("'%s' is available! Thats not good\n", url)
		os.Exit(1)
	})

	c.Visit(url)
}

func main() {
	checkOparl("")
	checkOparl("/system")
	checkVersion()
}
