package chiken

import (
	"superman/nodes/leaf/jettengame/gamedata/goclib/mahjonglib_go/wall"
)

// GenChikenDraw 生成翻牌鸡
func (mc *MChiken) GenChikenDraw(w *wall.Wall) int {
	if !w.IsAllDrawn() {
		mc.draw = w.GetFrowrdNextTile()
	}
	return mc.draw
}

// GetChikenDraw 获取翻牌鸡翻的那张牌
func (mc *MChiken) GetChikenDraw() int {
	return mc.draw
}
