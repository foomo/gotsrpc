package codegen

import "strings"

type Code struct {
	line   string
	lines  []string
	indent int
	tab    string
}

func NewCode(tab string) *Code {
	return &Code{
		line:   "",
		lines:  []string{},
		indent: 0,
		tab:    tab,
	}
}

func (c *Code) Ind(inc int) *Code {
	c.indent += inc
	if c.indent < 0 {
		c.indent = 0
	}

	return c
}

func (c *Code) NL() *Code {
	c.lines = append(c.lines, strings.Repeat(c.tab, c.indent)+c.line)
	c.line = ""

	return c
}

func (c *Code) L(line string) *Code {
	c.App(line).NL()
	return c
}

func (c *Code) App(str string) *Code {
	c.line += str
	return c
}

func (c *Code) String() string {
	if c.line != "" {
		c.lines = append(c.lines, c.line)
		c.line = ""
	}

	return strings.Join(c.lines, "\n")
}
