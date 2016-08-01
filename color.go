package golumns

import "fmt"

type colorOption string

const (
	Red     colorOption = "red"
	Green   colorOption = "green"
	Yellow  colorOption = "yellow"
	Blue    colorOption = "blue"
	Magenta colorOption = "magenta"
	Cyan    colorOption = "cyan"
	White   colorOption = "white"
)

type colorConfig struct {
	color colorOption
	items []string
}

type fullConfig struct {
	colorConfigs []colorConfig
}

func parseColor(rawColorString string) (colorOption, error) {
	colors := map[string]colorOption{
		"red":     Red,
		"green":   Green,
		"yellow":  Yellow,
		"blue":    Blue,
		"magenta": Magenta,
		"cyan":    Cyan,
		"white":   White,
	}
	a, ok := colors[rawColorString]
	if !ok {
		return "", fmt.Errorf("Unrecognized color option: %s", rawColorString)
	}
	return a, nil
}
