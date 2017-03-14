package engine

var (
	keyStrings = "EscapeUpDownLeftRightLeftMiddleRightMouse Scroll UpMouse Scroll Down"
	keyIndexes = [...]uint16{0, 6, 8, 12, 16, 21, 25, 31, 36, 51, 68}
)

func (k Key) String() string {
	if k >= 10 {
		return "UNKNOWN"
	}
	return keyStrings[keyIndexes[k]:keyIndexes[k+1]]
}
