package scli

import "strings"

type actionPattern struct {
	commands []string
	handler  Handler
}

func (p *actionPattern) parse(pattern string, h Handler) *actionPattern {
	words := strings.Split(pattern, " ")
	p.commands = make([]string, len(words))
	for i, word := range words {
		p.commands[i] = word
	}
	p.handler = h
	return p
}

func (p *actionPattern) resolve(args []string) (m map[string]string, ok bool) {
	if len(args) != len(p.commands) {
		return
	}
	m = make(map[string]string)
	for i, cmd := range p.commands {
		arg := args[i]
		if strings.HasPrefix(cmd, ":") {
			m[cmd[1:]] = arg
		} else if arg != cmd {
			return
		}
	}
	ok = true
	return
}
