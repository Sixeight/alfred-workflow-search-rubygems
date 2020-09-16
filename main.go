package main

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func setupItem(result *searchResult) {
	item := wf.NewItem(result.Name + " " + result.Version)
	item.Subtitle(result.Info)
	item.Arg(result.ProjectURI)
	item.Valid(true)

	cmdMod := item.NewModifier(aw.ModCmd)
	cmdMod.Subtitle("Open source code")

	var uri string
	if result.SourceCodeURI != "" {
		uri = result.SourceCodeURI
	} else if result.HomePageURI != "" {
		uri = result.HomePageURI
	}
	if uri != "" {
		cmdMod.Arg(uri)
		cmdMod.Valid(true)
	}

	if result.DocumentationURI != "" {
		ctrMod := item.NewModifier(aw.ModCtrl)
		ctrMod.Arg(result.DocumentationURI)
		ctrMod.Subtitle("Open documentation")
		ctrMod.Valid(true)
	}
}

func run() {
	results, err := search(strings.Join(wf.Args(), " "))
	if err != nil {
		wf.NewWarningItem("Faild to search", "")
	} else {
		for _, result := range results {
			setupItem(result)
		}
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
