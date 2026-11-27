package wsaa

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
	"testing"
)

func readREADME(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("read README.md: %v", err)
	}
	return string(b)
}

func TestREADMEHasNoPortfolioReference(t *testing.T) {
	readme := readREADME(t)
	for _, s := range []string{"leanucci/portfolio", "APPROACH.md"} {
		if strings.Contains(readme, s) {
			t.Errorf("README.md contains %q", s)
		}
	}
}

func TestREADMEKeepsClientDescription(t *testing.T) {
	readme := readREADME(t)
	for _, s := range []string{"WSAA", "TRA", "signs", "loginCms", "TOKEN", "SIGN"} {
		if !strings.Contains(readme, s) {
			t.Errorf("README.md does not contain %q", s)
		}
	}
}

// readmeSection returns the text of the "## <name>" section of README.md.
func readmeSection(t *testing.T, name string) string {
	t.Helper()
	readme := readREADME(t)
	start := strings.Index(readme, "\n## "+name+"\n")
	if start < 0 {
		t.Fatalf("README.md has no %q section", name)
	}
	rest := readme[start+len("\n## "+name+"\n"):]
	if end := strings.Index(rest, "\n## "); end >= 0 {
		rest = rest[:end]
	}
	return rest
}

// goCodeBlock returns the first Go code block in text.
func goCodeBlock(t *testing.T, text string) string {
	t.Helper()
	start := strings.Index(text, "```go\n")
	if start < 0 {
		t.Fatal("no Go code block")
	}
	rest := text[start+len("```go\n"):]
	end := strings.Index(rest, "```")
	if end < 0 {
		t.Fatal("Go code block has no end")
	}
	return rest[:end]
}

// packageAPI parses the package source files. It returns the exported
// package-level names and the method names.
func packageAPI(t *testing.T) (names, methods map[string]bool) {
	t.Helper()
	names = map[string]bool{}
	methods = map[string]bool{}
	fset := token.NewFileSet()
	for _, file := range []string{"environment.go", "errors.go"} {
		f, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", file, err)
		}
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if d.Recv != nil {
					methods[d.Name.Name] = true
				} else if d.Name.IsExported() {
					names[d.Name.Name] = true
				}
			case *ast.GenDecl:
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						names[s.Name.Name] = true
					case *ast.ValueSpec:
						for _, n := range s.Names {
							names[n.Name] = true
						}
					}
				}
			}
		}
	}
	return names, methods
}

func TestREADMEUsageCallsOnlyExistingAPI(t *testing.T) {
	code := goCodeBlock(t, readmeSection(t, "Usage"))
	f, err := parser.ParseFile(token.NewFileSet(), "usage.go", code, 0)
	if err != nil {
		t.Fatalf("Usage code example is not valid Go: %v", err)
	}
	names, methods := packageAPI(t)
	usedMethods := map[string]bool{}
	ast.Inspect(f, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		x, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		switch x.Name {
		case "fmt":
		case "wsaa":
			if !names[sel.Sel.Name] {
				t.Errorf("Usage uses wsaa.%s, which does not exist", sel.Sel.Name)
			}
		default:
			if !methods[sel.Sel.Name] {
				t.Errorf("Usage calls %s.%s, which does not exist", x.Name, sel.Sel.Name)
			}
			usedMethods[sel.Sel.Name] = true
		}
		return true
	})
	for _, m := range []string{"URL", "String"} {
		if !usedMethods[m] {
			t.Errorf("Usage does not call %s()", m)
		}
	}
}

func TestREADMEErrorsListsEveryErrorType(t *testing.T) {
	section := readmeSection(t, "Errors")
	names, _ := packageAPI(t)
	for name := range names {
		if strings.HasSuffix(name, "Error") && !strings.Contains(section, "wsaa."+name) {
			t.Errorf("Errors section does not list %s", name)
		}
	}
	for _, m := range regexp.MustCompile(`wsaa\.(\w+Error)\b`).FindAllStringSubmatch(section, -1) {
		if !names[m[1]] {
			t.Errorf("Errors section lists %s, which does not exist", m[1])
		}
	}
}

func TestREADMEEndpointsMatchEnvironment(t *testing.T) {
	section := readmeSection(t, "Endpoints")
	for _, url := range []string{TestingURL, ProductionURL} {
		if !strings.Contains(section, "`"+url+"`") {
			t.Errorf("Endpoints section does not contain %q", url)
		}
	}
	for _, url := range regexp.MustCompile("https://[^`\\s|]+").FindAllString(section, -1) {
		if url != TestingURL && url != ProductionURL {
			t.Errorf("Endpoints section contains an unknown URL %q", url)
		}
	}
}

func TestREADMEDevelopmentHasCLAUDECommands(t *testing.T) {
	section := readmeSection(t, "Development")
	b, err := os.ReadFile("CLAUDE.md")
	if err != nil {
		t.Fatalf("read CLAUDE.md: %v", err)
	}
	var commands []string
	re := regexp.MustCompile("^- \\*\\*(Install|Test|Lint):\\*\\* `([^`]+)`")
	for _, line := range strings.Split(string(b), "\n") {
		if m := re.FindStringSubmatch(line); m != nil {
			commands = append(commands, m[2])
		}
	}
	if len(commands) != 3 {
		t.Fatalf("found %d commands in CLAUDE.md, want 3", len(commands))
	}
	for _, c := range commands {
		if !strings.Contains(section, c) {
			t.Errorf("Development section does not contain %q", c)
		}
	}
}

func TestREADMEInstallSection(t *testing.T) {
	section := readmeSection(t, "Install")
	b, err := os.ReadFile("go.mod")
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	m := regexp.MustCompile(`(?m)^go (\S+)$`).FindStringSubmatch(string(b))
	if m == nil {
		t.Fatal("go.mod has no go directive")
	}
	for _, s := range []string{"go get github.com/leanucci/wsaa-go", "Go " + m[1]} {
		if !strings.Contains(section, s) {
			t.Errorf("Install section does not contain %q", s)
		}
	}
}

func TestREADMESaysTheClientDoesNotExistYet(t *testing.T) {
	section := readmeSection(t, "Status")
	want := "The login call, the TRA generation, the CMS signing, and the SOAP client do not exist yet."
	if !strings.Contains(section, want) {
		t.Errorf("Status section does not contain %q", want)
	}
	if strings.Contains(readREADME(t), ".Login(") {
		t.Error("README.md shows a Login call")
	}
}
