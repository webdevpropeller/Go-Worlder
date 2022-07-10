package database

import "strings"

// Table ...
type Table interface {
	NAME() string
	SetAlias(string)
}

// table ...
type table struct {
	name string
}

// NAME ...
func (t table) NAME() string {
	return t.name
}

func (t table) SetAlias(alias string) {
	t.name = strings.Join([]string{
		t.name, "AS", alias,
	}, " ")
}
