package types

type LayoutKey struct {
	Char   string `json:"char"`
	Row    uint8  `json:"row"`
	Col    uint8  `json:"col"`
	Finger uint8  `json:"finger"`
}

type Layout struct {
	Name          string      `json:"name"`
	Authors       []string    `json:"authors"`
	Link          string      `json:"link"`
	CreationTime  Timestamp   `json:"creation_time"`
	PrimaryBoards []string    `json:"primary_boards"`
	Keys          []LayoutKey `json:"keys"`
}

type Layouts map[string]Layout
