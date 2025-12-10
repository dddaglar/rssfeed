package main

import "errors"

type command struct {
	name string
	arg  []string
}

type commands struct {
	dict map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.dict[cmd.name]; ok {
		return handler(s, cmd)
	}
	return errors.New("command not found")
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.dict[name] = f
}
