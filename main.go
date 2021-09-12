package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"

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

func ListThemes(opts *docopt.Opts) []string {
	var palette map[string]bool

	useLight, _ := opts.Bool("--light") // am I the only one that will use this flag?
	useDark, _ := opts.Bool("--dark")

	filter := Any
	if useLight {
		filter = Light
	}
	if useDark {
		filter = Dark
	}

	switch filter {
	case Any:
		palette = AllThemes
	case Dark:
		palette = DarkThemes
	case Light:
		palette = LightThemes
	}

	idx := 0
	themes := make([]string, len(palette))

	for theme, _ := range palette {
		themes[idx] = theme
		idx++
	}

	sort.Strings(themes)

	return themes
}

func CheckDconf() {
	_, err := exec.LookPath("dconf")
	if err != nil {
		fmt.Println("guake-cl: guake's configuration is managed through a program called dconf, but we can't find dconf in PATH")
		os.Exit(1)
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
    guake-cl list-themes [(-l|--light)|(-d|--dark)]
    guake-cl (-h|--help|--version)

  Description:
    guake-cl is a command-line tool for update guake's theme. They can be picked by name, or at random
      from light / dark schemes.

  guake-cl updates the dconf setting '/apps/guake/style/font/palette'. You can update the guake-theme from any
  terminal, it doesn't need to be run from guake.

  themes can be previewed interactively using an fzf alias mentioned in examples.

  Options:
    --default   Restore to default theme for guake
    -r, --random    Choose a random theme; can be constrained by --light, --dark
    -l, --light     Select a light colour scheme when used with --random
    -d, --dark      Select a dark colour scheme when used with --random
    -h, --help      Display this documentation
    --version       Display the version-number

  Examples:
    guake-cl --default
    > rollback to guake's default theme

    guake-cl --solarized-dark
    > apply the theme 'solarized-dark' to your guake terminal

    guake-cl list-themes --dark | fzf --preview 'guake-cl --{}'
    > interactively choose a colour-scheme (requires fzf to be installed and on PATH). Please, install fzf and try this! If you like it,
      alias it in your .bashrc or .zshrc as alias guake-cl-fzf="guake-cl list-themes --dark | fzf --preview 'guake-cl --{}'"

    guake-cl -r
    > select and applies a random theme

    guake-cl -rl
    > select and applies a random light scheme

    guake-cl -rd
    > select and applies a random dark scheme

    guake-cl list-themes --light
    > list all light guake-cl themes

    guake-cl list-themes --dark
    > list all dark guake-cl themes

    guake-cl list-themes
    > list all guake-cl themes

    See Also:
      guake, which has a --change-palette option which is similar
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

	listThemes, _ := opts.Bool("list-themes")

	if listThemes {
		themes := ListThemes(&opts)

		for _, theme := range themes {
			fmt.Println(theme)
		}

		return
	}

	CheckDconf()
	SetTheme(&opts)
}
