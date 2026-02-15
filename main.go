package main

import (
	"fmt"
	"math/rand"
	
	tls "saper/tools"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	score := 0
	X,Y := 20,20
	SLOZH := 7
	MAXDIST := 3
	sp := 21
	bg := rl.Gray
	loosefl := false

	btns := make([][]tls.Btn,Y) 
	for i := 0;i<Y;i++ {
		btns = append(btns, make([]tls.Btn,X))
		for j := 0;j<X;j++ {
			if rand.Int()%SLOZH == 0 {
				btns[i] = append(btns[i], tls.Btn{X:int32(5+j*sp),Y:int32(20+i*sp),Txt:"*"})
			} else {
				btns[i] = append(btns[i], tls.Btn{X:int32(5+j*sp),Y:int32(20+i*sp),Txt:""})
			}
		}
	}
	for i := 0;i<Y;i++ {
		for j := 0;j<X;j++ {
			if btns[i][j].Txt == "*" {continue}
			btns[i][j].Txt = fmt.Sprintf("%d",tls.ParseBtn(j,i,&btns))
		}
	}

	rl.SetConfigFlags(rl.FlagWindowTopmost)
	rl.InitWindow(428,440,"Saper")
	rl.SetTargetFPS(30)
	rl.InitAudioDevice()
	rl.SetWindowIcon(*rl.LoadImage("icon.png"))
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if loosefl {
			rl.PlaySound(rl.LoadSound("mine.mp3"))
			bg = rl.Red
			rl.ClearBackground(bg)
			for i := 0;i<Y;i++ {
				for j := 0;j<X;j++ {
					btns[i][j].DrBtn()
				}
			}
			rl.EndDrawing()
			rl.WaitTime(2)
			rl.CloseAudioDevice()
			rl.CloseWindow()
		}

		rl.ClearBackground(bg)
		rl.DrawText(fmt.Sprint(score),200,2,20,rl.Black)

		for i := 0;i<Y;i++ {
			for j := 0;j<X;j++ {
				btns[i][j].DrBtn()
			}
		}
		for i := 0;i<Y;i++ {
			for j := 0;j<X;j++ {
				if btns[i][j].IsCL() {
					if btns[i][j].Txt == "*" {loosefl = true}
					if !btns[i][j].IsOpen {score++}
					if score == 1 {tls.OpenFirst(0,MAXDIST,j,i,&btns);continue}
					btns[i][j].IsOpen = true
				}
				
				if btns[i][j].IsFl() {
					if !btns[i][j].IsFlagged {btns[i][j].IsFlagged = true;continue}
					if btns[i][j].IsFlagged {btns[i][j].IsFlagged = false}
				}
			}
		}

		rl.EndDrawing()
	}
	rl.CloseAudioDevice()
}