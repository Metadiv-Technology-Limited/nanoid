package nanoid

import (
	"strings"

	"github.com/Metadiv-Technology-Limited/nanoid/internal/constant"
	gonanoid "github.com/matoous/go-nanoid"
)

type Opts struct {
	Numbers   bool
	Lowercase bool
	Uppercase bool
	/*
		Symbols are "_-"
	*/
	Symbols bool
	/*
		LookAlike are "1lI0Oouv5Ss"
	*/
	ExcludeAlike bool
	Length       int
}

/*
New returns a random string with the given options.
*/
func New(opts Opts) string {
	chars := ""
	if opts.Numbers {
		chars += constant.NUMBERS
	}
	if opts.Lowercase {
		chars += constant.LOWERCASE
	}
	if opts.Uppercase {
		chars += constant.UPPERCASE
	}
	if opts.Symbols {
		chars += constant.SYMBOLS
	}
	if opts.ExcludeAlike {
		for _, c := range constant.LOOK_ALIKE {
			chars = strings.ReplaceAll(chars, string(c), "")
		}
	}
	if chars == "" {
		panic("no chars")
	}
	if opts.Length == 0 {
		panic("length is zero")
	}
	s, err := gonanoid.Generate(chars, opts.Length)
	if err != nil {
		panic(err)
	}
	return s
}

/*
NewSafe quickly returns a random string with the safe options.
*/
func NewSafe() string {
	return New(Opts{
		Numbers:      true,
		Lowercase:    true,
		Uppercase:    true,
		Symbols:      false,
		ExcludeAlike: true,
		Length:       constant.SAFE_LENGTH,
	})
}
