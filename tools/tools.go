package tls

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Btn struct {
	X int32;
	Y int32;
	Txt string;
	IsOpen bool `default: false`;
	IsFlagged bool `default: false`;
}

var BtnSz int32 = 20

func (b Btn) Pr() {
	fmt.Println(b.X,b.Y,b.Txt,b.IsOpen)
}

func (b Btn) DrBtn() {
	if b.IsOpen {
		rl.DrawRectangle(b.X-1,b.Y-1,BtnSz+1,BtnSz+1,rl.Black)
		rl.DrawRectangle(b.X,b.Y,BtnSz,BtnSz,rl.LightGray)
		rl.DrawText(b.Txt,b.X+5,b.Y,22,rl.Black)
	} else if b.IsFlagged {
			rl.DrawRectangle(b.X,b.Y,BtnSz,BtnSz,rl.Red)
	} else {
		rl.DrawRectangle(b.X-1,b.Y-1,BtnSz+1,BtnSz+1,rl.Black)
		rl.DrawRectangle(b.X,b.Y,BtnSz,BtnSz,rl.DarkGray)
	}
}

func ParseBtn(x int,y int,arr *[][]Btn) int {
	cntr := 0
	for i := y-1;i<=y+1;i++{
		for j := x-1;j<=x+1;j++{
			if i < 0 || i > len((*arr))-1 || j < 0 || j > len((*arr)[i])-1 {continue}
			if (*arr)[i][j].Txt == "*" {fmt.Println(j,i);cntr++}
		}
	}
	if cntr > 0 {fmt.Println()}
	return cntr
}

func (b Btn) IsCL() bool {
	ms := rl.GetMousePosition()
	return rl.IsMouseButtonPressed(rl.MouseButtonLeft) && ms.X	> float32(b.X) && ms.X	< float32(b.X+BtnSz) && ms.Y > float32(b.Y) && ms.Y < float32(b.Y+BtnSz)
}

func (b Btn) IsFl() bool {
	ms := rl.GetMousePosition()
	return rl.IsMouseButtonPressed(rl.MouseButtonRight) && ms.X	> float32(b.X) && ms.X	< float32(b.X+BtnSz) && ms.Y > float32(b.Y) && ms.Y < float32(b.Y+BtnSz)
}

func OpenFirst(d int,mxd int,x int,y int,arr *[][]Btn) {
	if y < 0 || y > len((*arr))-1 || x < 0 || x > len((*arr)[0])-1 || (*arr)[y][x].Txt == "*" || d > mxd {return}
	(*arr)[y][x].IsOpen = true
	OpenFirst(d+1,3,x+1,y,arr)
	OpenFirst(d+1,3,x-1,y,arr)
	OpenFirst(d+1,3,x,y-1,arr)
	OpenFirst(d+1,3,x,y+1,arr)
}