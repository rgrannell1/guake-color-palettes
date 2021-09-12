package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	"github.com/docopt/docopt-go"
)

func UpdateSetting(name *string) {
	palette, ok := Palettes[*name]

	if !ok {
		fmt.Println("guake-cl: no configuration present for theme " + *name)
		os.Exit(1)
	}

	// dconf exists?
	quote := "'"

	cmd := exec.Command("dconf", "write", DCONF_PALETTE_KEY, quote+palette+quote)
	setErr := cmd.Run()

	if setErr != nil {
		fmt.Printf("guake-cl: failed to update colour-theme via dconf: %v\n", setErr)
		os.Exit(1)
	}
}

type ThemeFilter int

const (
	Any ThemeFilter = iota
	Dark
	Light
)

func ChooseRandomTheme(filter ThemeFilter) string {
	var palette map[string]bool

	switch filter {
	case Any:
		palette = AllThemes
	case Dark:
		palette = DarkThemes
	case Light:
		palette = LightThemes
	}

	choice := rand.Intn(len(palette))

	idx := 0
	for theme, _ := range palette {
		if choice == idx {
			return theme
		}
		idx += 1
	}

	return DEFAULT_THEME
}

func SetTheme(opts *docopt.Opts) {
	rand, _ := opts.Bool("--random")

	defaulted, _ := opts.Bool("--default")
	if defaulted {
		defaults := "default"
		UpdateSetting(&defaults)
	}

	if rand {
		useLight, _ := opts.Bool("--light") // am I the only one that will use this flag?
		useDark, _ := opts.Bool("--dark")

		var theme string
		if useLight {
			theme = ChooseRandomTheme(Light)
		} else if useDark {
			theme = ChooseRandomTheme(Dark)
		} else {
			theme = ChooseRandomTheme(Any)
		}

		UpdateSetting(&theme)

		return
	} else {
		// a particular theme was selected; enumerate and find it

		for theme, _ := range Palettes {
			selected, _ := opts.Bool("--" + theme)

			if selected {
				UpdateSetting(&theme)
				break
			}
		}
	}

}

func main() {
	usage := `guake-cl
  Usage:` + "\n"

	for key, _ := range Palettes {
		usage += "    guake-cl (--" + key + ")\n"
	}

	usage += `    guake-cl [-r|--random] [(-l|--light)|(-d|--dark)]
    guake-cl (--default)
    guake-cl (-h|--help|--version)

  Description:
    guake-cl is a command-line tool for update guake's theme. They can be picked by name, or at random
      from light / dark schemes.

  Options:
    -r, --random    Choose a random theme
    -l, --light     Select a light colour scheme when used with --random
    -d, --dark      Select a dark colour scheme when used with --random
    -h, --help      Display this documentation
    --version       Display the version-number
		`

	opts, err := docopt.ParseDoc(usage)

	if err != nil {
		fmt.Printf("guake-cl: failed to read arguments: %v\n", err)
		os.Exit(1)
	}

	showVersion, _ := opts.Bool("--version")

	if showVersion {
		fmt.Println(APP_VERSION)
		return
	}

	SetTheme(&opts)
}
