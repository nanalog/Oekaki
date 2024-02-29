package main

import (
	"fmt"
	"os"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
//型定義
var (
	Pimg *ebiten.Image
	Timg *ebiten.Image
	Rimg *ebiten.Image
	Himg *ebiten.Image
	Bimg *ebiten.Image
)

func init() {
	// 画像を読み込む, エラー処理
	var err error
	Pimg, _, err = ebitenutil.NewImageFromFile("image/sq.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	Timg, _, err = ebitenutil.NewImageFromFile("image/trash.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	Rimg, _, err = ebitenutil.NewImageFromFile("image/return.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	Himg, _, err = ebitenutil.NewImageFromFile("image/hude.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	Bimg, _, err = ebitenutil.NewImageFromFile("image/bar.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
}