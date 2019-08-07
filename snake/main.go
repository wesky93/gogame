package main

import tl "github.com/JoelOtter/termloop"

func initLevel() *tl.BaseLevel {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: '∴',
	})

	// 물웅덩이 추가
	level.AddEntity(tl.NewRectangle(4, 5, 10, 8, tl.ColorBlue))

	// 캐릭터 추가
	player := NewPlayer(level)
	level.AddEntity(player)
	return level
}

func main() {
	game := tl.NewGame()

	game.Screen().SetLevel(initLevel())
	game.Start()
}
