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
	item := wf.NewItem(result.Name + " " + result.Version).
		Subtitle(result.Info).
		Arg(result.ProjectURI).
		Valid(true)

	var uri string
	if result.SourceCodeURI != "" {
		uri = result.SourceCodeURI
	} else if result.HomePageURI != "" {
		uri = result.HomePageURI
	}
	if uri != "" {
		item.NewModifier(aw.ModCmd).
			Subtitle("Open source code").
			Arg(uri).Valid(true)
	}

	if result.DocumentationURI != "" {
		item.NewModifier(aw.ModCtrl).
			Arg(result.DocumentationURI).
			Subtitle("Open documentation").
			Valid(true)
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
