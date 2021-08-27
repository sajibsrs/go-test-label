package label

type label struct {
	pass string
	fail string
	info string
	warn string
}

func newlabel(block bool) *label {
	if block {
		return &label{
			pass: "\033[102;30m pass ✓ \033[0m ",
			fail: "\033[101;30m fail ✕ \033[0m ",
			info: "\033[106;30m info i \033[0m ",
			warn: "\033[103;30m warn ! \033[0m ",
		}
	}
	return &label{
		pass: "\033[32;1mpass ✓ \033[0m ",
		fail: "\033[31;1mfail ✕ \033[0m ",
		info: "\033[36;1minfo i \033[0m ",
		warn: "\033[33;1mwarn ! \033[0m ",
	}
}