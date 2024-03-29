package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

type Point struct {
	x, y int
}

func ReadMouseLeftClickPoint() Point {
	// 录制单次鼠标点击事件
	defer time.Sleep(time.Millisecond * 300)
	p := Point{}
	if robotgo.AddEvent("mleft") {
		p.x, p.y = robotgo.GetMousePos()
		fmt.Printf("---你按下左键, 坐标为(%d, %d)---\n", p.x, p.y)
	}
	return p
}
func WriteMouseLeftClickList(q *bool, mousePointList *[]Point) {
	// 重放鼠标点击事件队列
	for i := 1; !*q; i++ {
		for _, p := range *mousePointList {
			robotgo.MovesClick(p.x, p.y, "left", false)
			time.Sleep(time.Millisecond * 500)
		}
		fmt.Println("执行", i, "次")
	}
}
func main() {
	if startMsg := robotgo.ShowAlert("提示", "点击确定开始录制3个按键", "确定", "取消"); startMsg == 0 { //确定0，取消1
		mousePointList := []Point{ReadMouseLeftClickPoint(), ReadMouseLeftClickPoint(), ReadMouseLeftClickPoint()}
		if endMsg := robotgo.ShowAlert("提示", "录制完毕, 点击确定开始播放，开始后单击右键退出", "确定", "取消"); endMsg == 0 { //确定0，取消1
			quit := false
			go WriteMouseLeftClickList(&quit, &mousePointList)
			quit = robotgo.AddEvent("mright")
		}
	}
}
