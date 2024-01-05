package config

const (
	MaxGoRoutines = 10

	Reset    = "\033[0m"
	Red      = "\033[31m"
	Yellow   = "\033[33m"
	BgRed    = "\033[41m"
	BgYellow = "\033[43m"
	BgGreen  = "\033[42m"
	GreenTxt = "\033[32m"
)

type SpellCheckerConfig struct {
	MaxRoutines int
}

func NewSpellCheckerConfig() *SpellCheckerConfig {
	return &SpellCheckerConfig{
		MaxRoutines: MaxGoRoutines,
	}
}
