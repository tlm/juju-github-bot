package checklist

import (
	"text/template"
)

var ChecklistTemplate = `G'Day @{{ .PullRequest.User.Login}},

Please complete  the following checklist:
- [ ] Checked if it requires a [pylibjuju](https://github.com/juju/python-libjuju) change?
- [ ] Added [integration tests](https://github.com/juju/juju/tree/develop/tests) for the PR?
- [ ] Added or updated [doc.go](https://discourse.jujucharms.com/t/readme-in-packages/451) related to packages changed?
- [ ] Do comments answer the question of why design decisions were made?
`

var checklist *template.Template

func init() {
	tmpl, err := template.New("checklist").Parse(ChecklistTemplate)
	if err != nil {
		panic(err)
	}
	checklist = tmpl
}
