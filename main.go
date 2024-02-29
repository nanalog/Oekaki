package main

import (
	"fmt"
	"os"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)
// エラー処理1
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
// ゲーム開始用スクリプト
func run() error {
	g := newGame()
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle("お絵かき Canvas")
	if err := ebiten.RunGame(g); err != nil {
		return err
	}

	return nil
}
