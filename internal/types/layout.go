package types

import "time"

type LayoutKey struct {
	Char   string `json:"char"`
	Row    uint8  `json:"row"`
	Col    uint8  `json:"col"`
	Finger uint8  `json:"finger"`
}

type Layout struct {
	Name     string      `json:"name"`
	Owner    int64       `json:"owner"`
	Author   string      `json:"author"`
	Link     string      `json:"link"`
	Created  time.Time   `json:"created"`
	Modified time.Time   `json:"modified"`
	Boards   []string    `json:"boards"`
	Keys     []LayoutKey `json:"keys"`
}

type Layouts map[string]Layout
