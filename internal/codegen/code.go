package codegen

import "strings"

// Code builds indented source code line by line. Its methods return the
// receiver for chaining. Create one with [NewCode].
type Code struct {
	line   string
	lines  []string
	indent int
	tab    string
}

// NewCode returns an empty Code that indents with tab per level.
func NewCode(tab string) *Code {
	return &Code{
		line:   "",
		lines:  []string{},
		indent: 0,
		tab:    tab,
	}
}

// Ind changes the indentation level by inc. The level does not go below 0.
func (c *Code) Ind(inc int) *Code {
	c.indent += inc
	if c.indent < 0 {
		c.indent = 0
	}

	return c
}

// NL ends the current line at the current indentation level.
func (c *Code) NL() *Code {
	c.lines = append(c.lines, strings.Repeat(c.tab, c.indent)+c.line)
	c.line = ""

	return c
}

// L appends line to the current line and ends it.
func (c *Code) L(line string) *Code {
	c.App(line).NL()
	return c
}

// App appends str to the current line.
func (c *Code) App(str string) *Code {
	c.line += str
	return c
}

// String ends the current line, if not empty, and returns all lines joined
// by newlines.
func (c *Code) String() string {
	if c.line != "" {
		c.lines = append(c.lines, c.line)
		c.line = ""
	}

	return strings.Join(c.lines, "\n")
}
