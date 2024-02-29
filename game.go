package main

import (
	"image/color"
	"math"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type sl []int
var slit sl

type game struct {
	x      int
	y      int
	n      int
	m      int
	x_     float64
	y_     float64
	number int
	scale  float64
	width  int
	height int
	slices []sl
}

func newGame() *game {
	var g game
	g.reset()
	return &g
}

func (g *game) reset() {
	g.x = 0
	g.y = 0
	g.n = 0
	g.m = 0
	g.x_ = -70
	g.y_ = 0
	g.number = 0
	g.scale = 0.01
	g.width = 740
	g.height = 500
	g.slices = nil
}

func (g *game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.reset()
	}
	// 矢印キーの入力を取得し, 四角形の位置を更新する
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.x, g.y = ebiten.CursorPosition()
		// ゴミ箱選択時のみリセットをかける, 筆の太さ変更時以外であれば描画可能
		if g.x <= 56 && g.y <= 70 {
			g.reset()
		} else if g.y_ == 0 {
			slit = []int{g.x, g.y}
			g.slices = append(g.slices, slit)
			g.n += 1
		}
	}
	// 一つ前の動作に戻る設定, 筆の太さ変更に関する画面の消失設定, 実際に筆の太さを変更する設定
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.x <= 56 && g.y >= 80 && g.y <= 143 && g.number != 0 {
			g.number = len(g.slices) - g.number - 1
			g.slices = g.slices[:g.number]
		} else if g.x <= 56 && g.y >= 180 && g.y <= 235 {
			g.x_, g.y_ = 56.0, 40.0
		} else if g.y_ != 0 && g.x >= 110 && g.x <= 116 && g.y >= 221 && g.y <= 232 {
			g.x_, g.y_ = -70.0, 0
		} else if g.y_ != 0 && g.x > 89 && g.x <= 104 && g.y >= 225 && g.y <= 345 {
			g.scale = g.scale * math.Abs(345-float64(g.y)) / 60
		}
		//96,345
	}
	// 一つ前の動作に戻る設定に関連する箇所
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.m != g.n {
			g.number = g.n - g.m
			g.m = g.n
		}
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	var p, q float64 = 0, 0
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	op := &ebiten.DrawImageOptions{}
	// ゴミ箱画像の配置
	op.GeoM.Translate(0, 0)
	screen.DrawImage(Timg, op)
	// リターン画像の配置
	op.GeoM.Translate(0, 80)
	screen.DrawImage(Rimg, op)
	// 筆マーク画像の配置
	op.GeoM.Translate(0, 100)
	screen.DrawImage(Himg, op)
	// 筆マーク画像の設定に関する画面
	op.GeoM.Translate(g.x_, g.y_)
	screen.DrawImage(Bimg, op)
	// 線の図の大きさ設定
	op.GeoM.Scale(g.scale, g.scale)
	// 線の画像の配置
	for _, n := range g.slices {
		//for で動かすと、値は正常であるが、前の値を加算してしまうため、消去する。
		op.GeoM.Translate(float64(n[0])-p, float64(n[1])-q)
		p = float64(n[0])
		q = float64(n[1])
		screen.DrawImage(Pimg, op)
	}
}
//screenの大きさ
func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}
