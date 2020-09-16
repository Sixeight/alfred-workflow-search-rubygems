package main

import (
	"strings"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	results, err := search(strings.Join(wf.Args(), " "))
	if err != nil {
		wf.NewWarningItem("Faild to search", "")
	} else {
		for _, result := range results {
			item := wf.NewItem(result.Name)
			item.Subtitle(result.Authors)
			item.Arg(result.ProjectURI)
			item.Valid(true)
		}
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
