package main

import (
        "fmt"
	"strings"
        "os"

        "github.com/jedib0t/go-pretty/v6/table"
        "github.com/gocolly/colly/v2"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [acronyms]\n", os.Args[0])
    os.Exit(2)
}

func main() {
    if (len(os.Args) != 2){
        usage()
    }

    c := colly.NewCollector()
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"Class", "Name", "Unit"})

    c.OnHTML(".courseblock strong", func(e *colly.HTMLElement) {
                
		info := e.Text
                s := strings.Replace(info, "\n", "", -1)
                s1 := strings.Split(s, ".")
                t.AppendRow([]interface{}{s1[0], s1[1], s1[2]})
                t.AppendSeparator()

	})

    c.Visit("https://catalog.calpoly.edu/coursesaz/" + string(os.Args[1]) + "/")
    t.Render()
}

