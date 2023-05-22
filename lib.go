package ldtkbridgego

// Return all json data of a LDtk Level with a given name
func (world *LdtkJSON) GetLevelByName(name string) *Level {

	for _, l := range world.Levels {
		if l.Identifier == name {
			return &l
		}
	}

	return nil
}

func (world *LdtkJSON) GetLevelIndex(name string) int {
	for index, l := range world.Levels {
		if l.Identifier == name {
			return index
		}
	}
	return -1
}

func (level *Level) GetLayerByName(name string) *LayerInstance {
	for _, layer := range level.LayerInstances {
		if layer.Identifier == name {
			return &layer
		}
	}

	return nil
}

func (layer *LayerInstance) GetInt(x, y int) int {
	if x < 0 || x > layer.CWid-1 || y < 0 || y > layer.CHei-1 {
		// Check if out of level
		return -1
	} else {

		return layer.IntGridCsv[x+y*layer.CWid]
	}
}
