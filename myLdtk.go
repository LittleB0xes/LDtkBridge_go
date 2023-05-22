package ldtkbridgego

import "encoding/json"

func UnmarshalLdtkJSON(data []byte) (LdtkJSON, error) {
	var r LdtkJSON
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LdtkJSON) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LdtkJSON struct {
	Levels []Level     `json:"levels"`
	Defs   Definitions `json:"defs"`
}

type Level struct {
	Identifier     string          `json:"identifier"`
	PxWid          int             `json:"pxWid"`
	PxHei          int             `json:"PxHei"`
	LayerInstances []LayerInstance `json:"layerInstances"`
	FieldInstances []FieldInstance `json:"fieldInstances"`
}

type LayerInstance struct {
	Type            string           `json:"__type"`
	Identifier      string           `json:"__identifier"`
	CHei            int              `json:"__cHei"`
	CWid            int              `json:"__cWid"`
	EntityInstances []EntityInstance `json:"entityInstances"`
	GridTiles       []TileInstance   `json:"gridTiles"`
	TilesetDefUid   int              `json:"__tilesetDefUid"`
	GridSize        int              `json:"__gridSize"`
	IntGridCsv      []int            `json:"intGridCsv"`
}

type TileInstance struct {
	F   int   `json:"f"`
	T   int   `json:"t"`
	Src []int `json:"src"`
	Px  []int `json:"px"`
}

type EntityInstance struct {
	Identifier     string          `json:"__identifier"`
	Grid           []int           `json:"__grid"`
	Tile           TileRectangle   `json:"__tile"`
	FieldInstances []FieldInstance `json:"fieldInstances"`
}

type TileRectangle struct {
	H          int `json:"h"`
	W          int `json:"w"`
	X          int `json:"x"`
	Y          int `json:"y"`
	TilesetUid int `json:"tilesetUid"`
}

type FieldInstance struct {
	Identifier string        `json:"__identifier"`
	Tile       TileRectangle `json:"__tile"`
	Type       string        `json:"__type"`
	Value      interface{}   `json:"__value"`
}

type Definitions struct {
	Tilesets []TilesetDefinition `json:"tilesets"`
}

type TilesetDefinition struct {
	Identifier string `json:"__identifier"`
	Uid        int    `json:"uid"`
	RelPath    string `json:"relPath"`
}

