package main

import (
	"fmt"
	"math/rand"
	
	tls "saper/tools"
	rl "github.com/gen2brain/raylib-go/raylib"
)
// go build -ldflags="-H=windowsgui -s -w" -tags static -o Saper.exe
func main() {
	// X,Y := 20,20
	// SLOZH := 7
	// xrayfl := false
	
	// if !tls.MainMenu() {return}

	q,xrayfl,SLOZH,X,Y := tls.MainMenu()

	if !q {return}
	
	score := 0
	sp := 21
	bg := rl.Gray
	loosefl := false
	W := int32(10+sp*X )
	H := int32(25+sp*Y)

	btns := make([][]tls.Btn,Y) 
	//rasstanovka min + init polya
	for i := 0;i<Y;i++ {
		btns = append(btns, make([]tls.Btn,X))
		for j := 0;j<X;j++ {
			if rand.Intn(SLOZH) == 0 {
				btns[i] = append(btns[i], tls.NewBtn(5+j*sp,20+i*sp,"*"))
			} else {
				btns[i] = append(btns[i], tls.NewBtn(5+j*sp,20+i*sp,""))
			}
		}
	}
	// rasstanovka ciphor
	for i := 0;i<Y;i++ {
		for j := 0;j<X;j++ {
			if btns[i][j].Txt == "*" {continue}
			txt := tls.ParseBtn(j,i,&btns)
			if (txt == "0") {btns[i][j].Txt = ""} else {btns[i][j].Txt = txt}
		}
	}

	rl.SetConfigFlags(rl.FlagWindowTopmost)
	rl.InitWindow(W,H,"Saper")
	rl.SetTargetFPS(30)
	rl.SetWindowIcon(*rl.LoadImage("icon.png"))

	rl.InitAudioDevice()
	scissors := rl.LoadSound("scissors.mp3")
	mine := rl.LoadSound("mine.mp3")
	defer rl.UnloadSound(scissors)
	defer rl.UnloadSound(mine)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(bg)
		//text
		rl.DrawText("v2.0",5,2,10,rl.Black)
		rl.DrawText("score: "+fmt.Sprint(score),W/2-30,2,20,rl.Black)
		rl.DrawText(fmt.Sprintf("%d:%d",X,Y),W-30,2,10,rl.Black)

		//otrisovka polya
		for i := 0;i<Y;i++ {
			for j := 0;j<X;j++ {
					btns[i][j].Dr(xrayfl)
			}
		}

		//obrabotka nazhatiy
		for i := 0;i<Y;i++ {
			for j := 0;j<X;j++ {
				if btns[i][j].IsCL() {
					if score == 0 && !btns[i][j].IsFlagged {if tls.ParseBtn(j,i,&btns) != "0" {btns[i][j].Txt = tls.ParseBtn(j,i,&btns)} else {btns[i][j].Txt = ""}}
					if btns[i][j].Txt == "*" && !btns[i][j].IsFlagged {loosefl = true}
					if btns[i][j].Txt == "" {tls.OpenZero(j,i,&btns);rl.PlaySound(scissors);continue}
					if !btns[i][j].IsOpen && !btns[i][j].IsFlagged {score++;btns[i][j].IsOpen = true;rl.PlaySound(scissors)}
				}
				
				if btns[i][j].IsFl() {
					if !btns[i][j].IsFlagged {btns[i][j].IsFlagged = true;continue}
					if btns[i][j].IsFlagged {btns[i][j].IsFlagged = false}
				}
			}
		}

		// obrabotka smerty
		if loosefl {
			rl.PlaySound(mine)
			bg = rl.Red
			rl.ClearBackground(bg)
			for i := 0;i<Y;i++ {
				for j := 0;j<X;j++ {
					btns[i][j].Dr(xrayfl)
				}
			}
			rl.EndDrawing()
			rl.WaitTime(2)
			rl.CloseAudioDevice()
			rl.CloseWindow()
		}

		rl.EndDrawing()
	}
	rl.CloseAudioDevice()
}