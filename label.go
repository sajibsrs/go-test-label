package label

type Label struct {
	Pass string
	Fail string
	Info string
	Warn string
}

func NewLabel(block bool) *Label {
	if block {
		return &Label{
			Pass: "\033[102;30m Pass ✓ \033[0m ",
			Fail: "\033[101;30m Fail ✕ \033[0m ",
			Info: "\033[106;30m Info i \033[0m ",
			Warn: "\033[103;30m Warn ! \033[0m ",
		}
	}
	return &Label{
		Pass: "\033[32;1mPass ✓ \033[0m ",
		Fail: "\033[31;1mfail ✕ \033[0m ",
		Info: "\033[36;1minfo i \033[0m ",
		Warn: "\033[33;1mwarn ! \033[0m ",
	}
}
