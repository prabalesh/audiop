package util

func PrintError(err error) {
	Colorize(ColorRed, "Error: "+err.Error()+"\n")
}
