package date

import (
	"time"

	"github.com/charmbracelet/gum/style"
)

// Options are the customization options for the date.
type Options struct {
	Placeholder     string        `help:"Placeholder value" default:"Type something..." env:"GUM_DATE_PLACEHOLDER"`
	Prompt          string        `help:"Prompt to display" default:"> " env:"GUM_DATE_PROMPT"`
	PromptStyle     style.Styles  `embed:"" prefix:"prompt." envprefix:"GUM_DATE_PROMPT_"`
	CursorTextStyle style.Styles  `embed:"" prefix:"cursor." set:"defaultForeground=212" envprefix:"GUM_DATE_CURSOR_"`
	Value           string        `help:"Initial value (can also be passed via stdin)" default:""`
	Header          string        `help:"Header value" default:"" env:"GUM_DATE_HEADER"`
	HeaderStyle     style.Styles  `embed:"" prefix:"header." set:"defaultForeground=240" envprefix:"GUM_DATE_HEADER_"`
	Timeout         time.Duration `help:"Timeout until input aborts" default:"0" env:"GUM_DATE_TIMEOUT"`
}
