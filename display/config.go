package display

// Config has switches for setting configurations for output
type Config struct {
	autoColumn  bool
	colorConfig color
}

// Color has switches for setting configuration options for color in the output
type color struct {
	enabled               bool     // on/off
	colorPattern          []string // pattern of colors to alternate between
	listOfElementsToColor []string // option to only color specified elements
}
