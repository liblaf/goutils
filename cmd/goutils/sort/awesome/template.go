package awesome

import (
	"strings"
	"text/template"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var columns = [][2]string{
	{"Name", `[{{ .Name }}]({{ .HTMLURL }})`},
	{"Social", `![GitHub Repo stars](https://img.shields.io/github/stars/{{ .FullName }}) ![GitHub commit activity](https://img.shields.io/github/commit-activity/y/{{ .FullName }})`},
	{"Description", `{{ .Description }}`},
}

var (
	templateRow     *template.Template
	templateSection *template.Template
)

func initTemplate() {
	var err error

	rowText := `|`
	for _, column := range columns {
		rowText += ` ` + column[1] + ` |`
	}
	templateRow = template.Must(template.New("row").Parse(rowText))
	_errors.Check(err)

	sectionText := `## {{ .Section }}` + "\n"
	sectionText += `|`
	for _, column := range columns {
		sectionText += ` ` + column[0] + ` |`
	}
	sectionText += "\n"
	sectionText += `|`
	for range columns {
		sectionText += ` --- |`
	}
	sectionText += "\n"
	sectionText += `{{ .Rows }}` + "\n"
	templateSection = template.Must(template.New("section").Parse(sectionText))
}

func titleCase(s string) string {
	s = strings.ReplaceAll(s, `-`, ` `)
	s = cases.Title(language.English).String(s)
	return s
}
