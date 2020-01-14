package app

// Configuration of the app
type Configuration struct {
	// Project directory
	DirPath string
	// Main documentation file
	MainFile string
	// Endpoints root folder
	EndsRoot string
	// Output documentation folder
	Output string
	// Verbose mode, i.e. show warnings
	Verbose bool
}
