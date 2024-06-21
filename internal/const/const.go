package cnst

var HomeBase = "base"

var StaticFileDir = "./ui/static/"

const (
	GreenBackground = "\033[42m"
	RedBackground   = "\033[41m"
	Reset           = "\033[0m"
)

var HomeHTMLLists = []string{
	"./ui/html/base.html",
	"./ui/html/pages/home.html",
	"./ui/html/partials/nav.html",
}

var ViewSnippetHTMLLists = []string{
	"./ui/html/base.html",
	"./ui/html/partials/nav.html",
	"./ui/html/pages/view.html",
}
