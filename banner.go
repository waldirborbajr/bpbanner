package banner

func init() {}

func Banner(sysVersion string) string {
	var banner string

	banner += "\n"
	banner += " ____              _______                 _    _ \n"
	banner += "|  _ \\     _      |__   __|               | |  | |\n"
	banner += "| |_) |  _| |_       | |      ___    ___  | |__| |\n"
	banner += "|  _ <  |_   _|      | |     / _ \\  / __| |  __  |\n"
	banner += "| |_) |   |_|        | |    |  __/ | (__  | |  | |\n"
	banner += "|____/               |_|     \\___|  \\___| |_|  |_|\n"
	banner += "\n"
	banner += sysVersion
	banner += "\n"

	return banner
}
