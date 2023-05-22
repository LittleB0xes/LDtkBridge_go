package ldtkbridgego

import (
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)


// Render a layer with a given name
func (level *Level) RenderLayerByName(screen *ebiten.Image, tilesetLib map[int]*ebiten.Image, name string, offsetX, offsetY, scale float64) {
	for _, layer := range level.LayerInstances {
		if layer.Type == "Tiles" && layer.Identifier == name {
			uid := layer.TilesetDefUid
			tileset := tilesetLib[uid]
			gridSize := layer.GridSize
			for _, tile := range layer.GridTiles {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(tile.Px[0])-offsetX, float64(tile.Px[1])-offsetY)

				op.GeoM.Scale(scale, scale)

				screen.DrawImage(tileset.SubImage(image.Rect(tile.Src[0], tile.Src[1], gridSize+tile.Src[0], gridSize+tile.Src[1])).(*ebiten.Image), op)
			}
		}
	}
}
