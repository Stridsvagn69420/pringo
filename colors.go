package pringo

// Font Colors
type Color string

const (
	reset   Color = "\u001b[0m"
	Black   Color = "\u001b[30m"
	Red     Color = "\u001b[31m"
	Green   Color = "\u001b[32m"
	Yellow  Color = "\u001b[33m"
	Blue    Color = "\u001b[34m"
	Magenta Color = "\u001b[35m"
	Cyan    Color = "\u001b[36m"
	White   Color = "\u001b[37m"
	None    Color = ""
)

// Font Colors Bright
const (
	BlackBright   Color = "\u001b[30;1m"
	RedBright     Color = "\u001b[31;1m"
	GreenBright   Color = "\u001b[32;1m"
	YellowBright  Color = "\u001b[33;1m"
	BlueBright    Color = "\u001b[34;1m"
	MagentaBright Color = "\u001b[35;1m"
	CyanBright    Color = "\u001b[36;1m"
	WhiteBright   Color = "\u001b[37;1m"
)

// Background Colors

const (
	BackgroundBlack   Color = "\u001b[40m"
	BackgroundRed     Color = "\u001b[41m"
	BackgroundGreen   Color = "\u001b[42m"
	BackgroundYellow  Color = "\u001b[43m"
	BackgroundBlue    Color = "\u001b[44m"
	BackgroundMagenta Color = "\u001b[45m"
	BackgroundCyan    Color = "\u001b[46m"
	BackgroundWhite   Color = "\u001b[47m"
)

// Background Colors Bright
const (
	BackgroundBlackBright   Color = "\u001b[40;1m"
	BackgroundRedBright     Color = "\u001b[41;1m"
	BackgroundGreenBright   Color = "\u001b[42;1m"
	BackgroundYellowBright  Color = "\u001b[43;1m"
	BackgroundBlueBright    Color = "\u001b[44;1m"
	BackgroundMagentaBright Color = "\u001b[45;1m"
	BackgroundCyanBright    Color = "\u001b[46;1m"
	BackgroundWhiteBright   Color = "\u001b[47;1m"
)
