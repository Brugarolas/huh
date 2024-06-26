package huh

import (
	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

// Theme is a collection of styles for components of the form.
// Themes can be applied to a form using the WithTheme option.
type Theme struct {
	Form           lipgloss.Style
	Group          lipgloss.Style
	FieldSeparator lipgloss.Style
	Blurred        FieldStyles
	Focused        FieldStyles
	Help           help.Styles
}

// FieldStyles are the styles for input fields.
type FieldStyles struct {
	Base           lipgloss.Style
	Title          lipgloss.Style
	Description    lipgloss.Style
	ErrorIndicator lipgloss.Style
	ErrorMessage   lipgloss.Style

	// Select styles.
	SelectSelector lipgloss.Style // Selection indicator
	Option         lipgloss.Style // Select options
	NextIndicator  lipgloss.Style
	PrevIndicator  lipgloss.Style
	DisabledOption lipgloss.Style

	// FilePicker styles.
	Directory lipgloss.Style
	File      lipgloss.Style

	// Multi-select styles.
	MultiSelectSelector lipgloss.Style
	SelectedOption      lipgloss.Style
	SelectedPrefix      lipgloss.Style
	UnselectedOption    lipgloss.Style
	UnselectedPrefix    lipgloss.Style

	// Textinput and teatarea styles.
	TextInput TextInputStyles

	// Confirm styles.
	FocusedButton lipgloss.Style
	BlurredButton lipgloss.Style

	// Card styles.
	Card      lipgloss.Style
	NoteTitle lipgloss.Style
	Next      lipgloss.Style
}

// TextInputStyles are the styles for text inputs.
type TextInputStyles struct {
	Cursor      lipgloss.Style
	CursorText  lipgloss.Style
	Placeholder lipgloss.Style
	Prompt      lipgloss.Style
	Text        lipgloss.Style
}

// copy returns a copy of a TextInputStyles with all children styles copied.
func (t TextInputStyles) copy() TextInputStyles {
	return TextInputStyles{
		Cursor:      t.Cursor.Copy(),
		Placeholder: t.Placeholder.Copy(),
		Prompt:      t.Prompt.Copy(),
		Text:        t.Text.Copy(),
	}
}

// copy returns a copy of a FieldStyles with all children styles copied.
func (f FieldStyles) copy() FieldStyles {
	return FieldStyles{
		Base:                f.Base.Copy(),
		Title:               f.Title.Copy(),
		Description:         f.Description.Copy(),
		ErrorIndicator:      f.ErrorIndicator.Copy(),
		ErrorMessage:        f.ErrorMessage.Copy(),
		SelectSelector:      f.SelectSelector.Copy(),
		NextIndicator:       f.NextIndicator.Copy(),
		PrevIndicator:       f.PrevIndicator.Copy(),
		Option:              f.Option.Copy(),
		Directory:           f.Directory.Copy(),
		File:                f.File.Copy(),
		MultiSelectSelector: f.MultiSelectSelector.Copy(),
		SelectedOption:      f.SelectedOption.Copy(),
		DisabledOption:      f.DisabledOption.Copy(),
		SelectedPrefix:      f.SelectedPrefix.Copy(),
		UnselectedOption:    f.UnselectedOption.Copy(),
		UnselectedPrefix:    f.UnselectedPrefix.Copy(),
		FocusedButton:       f.FocusedButton.Copy(),
		BlurredButton:       f.BlurredButton.Copy(),
		TextInput:           f.TextInput.copy(),
		Card:                f.Card.Copy(),
		NoteTitle:           f.NoteTitle.Copy(),
		Next:                f.Next.Copy(),
	}
}

const (
	buttonPaddingHorizontal = 2
	buttonPaddingVertical   = 0
)

// ThemeBase returns a new base theme with general styles to be inherited by
// other themes.
func ThemeBase() *Theme {
	var t Theme

	t.FieldSeparator = lipgloss.NewStyle().SetString("\n\n")

	button := lipgloss.NewStyle().
		Padding(buttonPaddingVertical, buttonPaddingHorizontal).
		MarginRight(1)

	// Focused styles.
	t.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true)
	t.Focused.Card = lipgloss.NewStyle().PaddingLeft(1)
	t.Focused.ErrorIndicator = lipgloss.NewStyle().SetString(" *")
	t.Focused.ErrorMessage = lipgloss.NewStyle().SetString(" *")
	t.Focused.SelectSelector = lipgloss.NewStyle().SetString("> ")
	t.Focused.NextIndicator = lipgloss.NewStyle().MarginLeft(1).SetString("→")
	t.Focused.PrevIndicator = lipgloss.NewStyle().MarginRight(1).SetString("←")
	t.Focused.MultiSelectSelector = lipgloss.NewStyle().SetString("> ")
	t.Focused.SelectedPrefix = lipgloss.NewStyle().SetString("[•] ")
	t.Focused.UnselectedPrefix = lipgloss.NewStyle().SetString("[ ] ")
	t.Focused.FocusedButton = button.Foreground(lipgloss.Color("0")).Background(lipgloss.Color("7"))
	t.Focused.BlurredButton = button.Foreground(lipgloss.Color("7")).Background(lipgloss.Color("0"))
	t.Focused.TextInput.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	t.Help = help.New().Styles

	// Blurred styles.
	t.Blurred = t.Focused
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.MultiSelectSelector = lipgloss.NewStyle().SetString("  ")
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return &t
}

// ThemeCharm returns a new theme based on the Charm color scheme.
func ThemeCharm() *Theme {
	t := ThemeBase()

	var (
		normalFg = lipgloss.AdaptiveColor{Light: "235", Dark: "252"}
		indigo   = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
		cream    = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#FFFDF5"}
		fuchsia  = lipgloss.Color("#F780E2")
		green    = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
		red      = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
		comment  = lipgloss.AdaptiveColor{Light: "", Dark: "243"}
	)

	f := &t.Focused
	f.Base = f.Base.BorderForeground(lipgloss.Color("238"))
	f.Title.Foreground(indigo).Bold(true)
	f.NoteTitle.Foreground(indigo).Bold(true).MarginBottom(1)
	f.Directory.Foreground(indigo)
	f.Description.Foreground(comment)
	f.ErrorIndicator.Foreground(red)
	f.ErrorMessage.Foreground(red)
	f.SelectSelector.Foreground(fuchsia)
	f.NextIndicator.Foreground(fuchsia)
	f.PrevIndicator.Foreground(fuchsia)
	f.Option.Foreground(normalFg)
	f.MultiSelectSelector.Foreground(fuchsia)
	f.SelectedOption.Foreground(green)
	f.DisabledOption.Foreground(comment)
	f.SelectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#02CF92", Dark: "#02A877"}).SetString("✓ ")
	f.UnselectedPrefix = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "", Dark: "243"}).SetString("• ")
	f.UnselectedOption.Foreground(normalFg)
	f.FocusedButton.Foreground(cream).Background(fuchsia)
	f.Next = f.FocusedButton.Copy()
	f.BlurredButton.Foreground(normalFg).Background(lipgloss.AdaptiveColor{Light: "252", Dark: "237"})

	f.TextInput.Cursor.Foreground(green)
	f.TextInput.Placeholder.Foreground(lipgloss.AdaptiveColor{Light: "248", Dark: "238"})
	f.TextInput.Prompt.Foreground(fuchsia)

  t.Blurred = f.copy()
	t.Blurred.Base = f.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}

// ThemeDracula returns a new theme based on the Dracula color scheme.
func ThemeDracula() *Theme {
	t := ThemeBase()

	var (
		background = lipgloss.AdaptiveColor{Dark: "#282a36"}
		selection  = lipgloss.AdaptiveColor{Dark: "#44475a"}
		foreground = lipgloss.AdaptiveColor{Dark: "#f8f8f2"}
		comment    = lipgloss.AdaptiveColor{Dark: "#6272a4"}
		green      = lipgloss.AdaptiveColor{Dark: "#50fa7b"}
		purple     = lipgloss.AdaptiveColor{Dark: "#bd93f9"}
		red        = lipgloss.AdaptiveColor{Dark: "#ff5555"}
		yellow     = lipgloss.AdaptiveColor{Dark: "#f1fa8c"}
	)

	f := &t.Focused
	f.Base.BorderForeground(selection)
	f.Title.Foreground(purple)
	f.NoteTitle.Foreground(purple)
	f.Description.Foreground(comment)
	f.ErrorIndicator.Foreground(red)
	f.Directory.Foreground(purple)
	f.File.Foreground(foreground)
	f.ErrorMessage.Foreground(red)
	f.SelectSelector.Foreground(yellow)
	f.NextIndicator.Foreground(yellow)
	f.PrevIndicator.Foreground(yellow)
	f.Option.Foreground(foreground)
	f.MultiSelectSelector.Foreground(yellow)
	f.SelectedOption.Foreground(green)
	f.DisabledOption.Foreground(comment)
	f.SelectedPrefix.Foreground(green)
	f.UnselectedOption.Foreground(foreground)
	f.UnselectedPrefix.Foreground(comment)
	f.FocusedButton.Foreground(yellow).Background(purple).Bold(true)
	f.BlurredButton.Foreground(foreground).Background(background)

	f.TextInput.Cursor.Foreground(yellow)
	f.TextInput.Placeholder.Foreground(comment)
	f.TextInput.Prompt.Foreground(yellow)

  t.Blurred = f.copy()
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}

// ThemeBase16 returns a new theme based on the base16 color scheme.
func ThemeBase16() *Theme {
	t := ThemeBase().copy()

	f := &t.Focused
	f.Base.BorderForeground(lipgloss.Color("8"))
	f.Title.Foreground(lipgloss.Color("6"))
	f.NoteTitle.Foreground(lipgloss.Color("6"))
	f.Directory.Foreground(lipgloss.Color("6"))
	f.Description.Foreground(lipgloss.Color("8"))
	f.ErrorIndicator.Foreground(lipgloss.Color("9"))
	f.ErrorMessage.Foreground(lipgloss.Color("9"))
	f.SelectSelector.Foreground(lipgloss.Color("3"))
	f.NextIndicator.Foreground(lipgloss.Color("3"))
	f.PrevIndicator.Foreground(lipgloss.Color("3"))
	f.Option.Foreground(lipgloss.Color("7"))
	f.MultiSelectSelector.Foreground(lipgloss.Color("3"))
	f.SelectedOption.Foreground(lipgloss.Color("2"))
	f.DisabledOption.Foreground(lipgloss.Color("8"))
	f.SelectedPrefix.Foreground(lipgloss.Color("2"))
	f.UnselectedOption.Foreground(lipgloss.Color("7"))
	f.FocusedButton.Foreground(lipgloss.Color("7")).Background(lipgloss.Color("5"))
	f.BlurredButton.Foreground(lipgloss.Color("7")).Background(lipgloss.Color("0"))

	f.TextInput.Cursor.Foreground(lipgloss.Color("5"))
	f.TextInput.Placeholder.Foreground(lipgloss.Color("8"))
	f.TextInput.Prompt.Foreground(lipgloss.Color("3"))

	t.Blurred = f.copy()
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NoteTitle = t.Blurred.NoteTitle.Foreground(lipgloss.Color("8"))
	t.Blurred.Title = t.Blurred.NoteTitle.Foreground(lipgloss.Color("8"))

	t.Blurred.TextInput.Prompt = t.Blurred.TextInput.Prompt.Foreground(lipgloss.Color("8"))
	t.Blurred.TextInput.Text = t.Blurred.TextInput.Text.Foreground(lipgloss.Color("7"))

	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}

// ThemeCatppuccin returns a new theme based on the Catppuccin color scheme.
func ThemeCatppuccin() *Theme {
	t := ThemeBase()

	light := catppuccin.Latte
	dark := catppuccin.Mocha
	var (
		base     = lipgloss.AdaptiveColor{Light: light.Base().Hex, Dark: dark.Base().Hex}
		text     = lipgloss.AdaptiveColor{Light: light.Text().Hex, Dark: dark.Text().Hex}
		subtext1 = lipgloss.AdaptiveColor{Light: light.Subtext1().Hex, Dark: dark.Subtext1().Hex}
		subtext0 = lipgloss.AdaptiveColor{Light: light.Subtext0().Hex, Dark: dark.Subtext0().Hex}
		overlay1 = lipgloss.AdaptiveColor{Light: light.Overlay1().Hex, Dark: dark.Overlay1().Hex}
		overlay0 = lipgloss.AdaptiveColor{Light: light.Overlay0().Hex, Dark: dark.Overlay0().Hex}
		green    = lipgloss.AdaptiveColor{Light: light.Green().Hex, Dark: dark.Green().Hex}
		red      = lipgloss.AdaptiveColor{Light: light.Red().Hex, Dark: dark.Red().Hex}
		pink     = lipgloss.AdaptiveColor{Light: light.Pink().Hex, Dark: dark.Pink().Hex}
		mauve    = lipgloss.AdaptiveColor{Light: light.Mauve().Hex, Dark: dark.Mauve().Hex}
		cursor   = lipgloss.AdaptiveColor{Light: light.Rosewater().Hex, Dark: dark.Rosewater().Hex}
	)

	f := &t.Focused
	f.Base.BorderForeground(subtext1)
	f.Title.Foreground(mauve)
	f.NoteTitle.Foreground(mauve)
	f.Directory.Foreground(mauve)
	f.Description.Foreground(subtext0)
	f.ErrorIndicator.Foreground(red)
	f.ErrorMessage.Foreground(red)
	f.SelectSelector.Foreground(pink)
	f.NextIndicator.Foreground(pink)
	f.PrevIndicator.Foreground(pink)
	f.Option.Foreground(text)
	f.MultiSelectSelector.Foreground(pink)
	f.SelectedOption.Foreground(green)
	f.DisabledOption.Foreground(subtext0)
	f.SelectedPrefix.Foreground(green)
	f.UnselectedPrefix.Foreground(text)
	f.UnselectedOption.Foreground(text)
	f.FocusedButton.Foreground(base).Background(pink)
	f.BlurredButton.Foreground(text).Background(base)

	f.TextInput.Cursor.Foreground(cursor)
	f.TextInput.Placeholder.Foreground(overlay0)
	f.TextInput.Prompt.Foreground(pink)

	t.Blurred = f.copy()
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())

	t.Help.Ellipsis = t.Help.Ellipsis.Foreground(subtext0)
	t.Help.ShortKey = t.Help.ShortKey.Foreground(subtext0)
	t.Help.ShortDesc = t.Help.ShortDesc.Foreground(overlay1)
	t.Help.ShortSeparator = t.Help.ShortSeparator.Foreground(subtext0)
	t.Help.FullKey = t.Help.FullKey.Foreground(subtext0)
	t.Help.FullDesc = t.Help.FullDesc.Foreground(overlay1)
	t.Help.FullSeparator = t.Help.FullSeparator.Foreground(subtext0)

	return t
}
