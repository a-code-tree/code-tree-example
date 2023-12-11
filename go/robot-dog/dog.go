package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math"
)

func main() {
	const (
		width  = 800
		height = 600
	)

	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 设置背景颜色为白色
	dc.Clear()

	// 绘制机器狗身体
	dc.SetRGB(0.5, 0.5, 0.5)
	dc.DrawEllipse(300, 400, 100, 50)
	dc.Fill()

	// 绘制机器狗头部
	dc.DrawCircle(400, 350, 40)
	dc.Fill()

	// 绘制机器狗四条腿
	legSpacing := 80.0
	for i := 0; i < 4; i++ {
		x := 300 + legSpacing*float64(i)
		y := 450
		dc.DrawEllipse(x, float64(y), 20, 40)
		dc.Fill()
	}

	// 绘制机器狗尾巴
	tailX := 250.0
	tailY := 380.0
	tailLength := 40.0
	tailAngle := math.Pi / 6.0
	dc.DrawLine(250, 380, tailX+tailLength*math.Cos(tailAngle), tailY+tailLength*math.Sin(tailAngle))
	dc.SetLineWidth(5)
	dc.Stroke()

	// 保存绘制结果为图片文件
	err := dc.SavePNG("running_robot_dog.png")
	if err != nil {
		fmt.Println("Error saving PNG:", err)
	}
}
