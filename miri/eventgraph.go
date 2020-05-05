package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

var (
	c chan bool = make(chan bool)
)

//差の絶対値
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//あらゆる角度の直線をひく関数
//今はまだ仕組みが分からんただのコピペ
func drawline(img *image.RGBA, x1, y1, x2, y2 int, color color.RGBA) {
	step := 0
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	if dx > dy {
		if x1 > x2 {
			if y1 > y2 {
				step = 1
			} else {
				step = -1
			}
			x1, x2 = x2, x1
			y1 = y2
		} else {
			if y1 < y2 {
				step = 1
			} else {
				step = -1
			}
		}
		img.Set(x1, y1, color)
		s := dx >> 1
		x1++
		for x1 <= x2 {
			s -= dy
			if s < 0 {
				s += dx
				y1 += step
			}
			img.Set(x1, y1, color)
			x1++
		}
	} else {
		if y1 > y2 {
			if x1 > x2 {
				step = 1
			} else {
				step = -1
			}
			y1, y2 = y2, y1
			x1 = x2
		} else {
			if x1 < x2 {
				step = 1
			} else {
				step = -1
			}
		}
		img.Set(x1, y1, color)
		s := dy >> 1
		y1++
		for y1 <= y2 {
			s -= dx
			if s < 0 {
				s += dy
				x1 += step
			}
			img.Set(x1, y1, color)
			y1++
		}
	}
}

func write(filename, eventname, rank string, r, g, b byte) {

	//ボーダーのファイルを開く
	file, err := os.Open(filename)

	//終わったら閉じる
	defer file.Close()

	//エラー
	if err != nil {
		fmt.Println(err)
	}

	//データ入れる箱
	datas := make([]int, 0, 590)

	//スキャナー等準備
	scanner := bufio.NewScanner(file)
	var data int

	//スキャン
	for scanner.Scan() {

		//読んだのをintに変換して箱詰め
		data, _ = strconv.Atoi(scanner.Text())
		datas = append(datas, data)
	}

	//画像生成
	width := 2360
	height := 2500
	upleft := image.Point{0, 0}
	lowright := image.Point{width, height}
	//ボーダー用
	border := image.NewRGBA(image.Rectangle{upleft, lowright})
	//差用
	dif := image.NewRGBA(image.Rectangle{upleft, lowright})

	//指定した色作成
	col := color.RGBA{r, g, b, 255}

	//座標を一時記録する場所
	x := 0
	y := 0

	//グラフ作成
	//軸を書く
	//一時的にx, yを借りる
	axis1 := color.RGBA{238, 238, 238, 130}
	axis2 := color.RGBA{240, 240, 240, 140}
	for ; x < width; x++ {
		for y = 0; y < height; y++ {
			//1時間毎に緑、
			if x%8 == 0 {
				dif.Set(x, y, axis2)
				border.Set(x, y, axis2)

				//30分毎に青い縦線を引く
			} else if x%4 == 0 {
				dif.Set(x, y, axis1)
				border.Set(x, y, axis1)

				//ボーダーの方は10万P毎に緑、
			} else if y%100 == 0 {
				border.Set(x, y, axis2)

				//1万P毎に青い横線を引く
			} else if y%10 == 0 {
				border.Set(x, y, axis1)
			}

			//差の方は一万ポイント毎に緑
			if y%250 == 0 {
				dif.Set(x, y, axis2)

				//千ポイント毎に青い横線を引く
			} else if y%25 == 0 {
				dif.Set(x, y, axis1)
			}
		}
	}

	//折れ線書く
	x = 0
	y = 2500
	yy := 2500
	//データを一時記録する場所
	wait := 0
	for i, v := range datas {
		drawline(dif, x, yy, i*4, 2500-(v-wait)/40, col)
		drawline(border, x, y, i*4, 2500-v/1000, col)
		x = i * 4
		y = 2500 - v/1000
		yy = 2500 - (v-wait)/40
		wait = v
	}

	//画像ファイル作成
	pic1, err := os.Create("C:\\Users\\ryogen\\Desktop\\ミリグラ\\0登録済み\\" + rank + eventname + "_bor.png")
	if err != nil {
		fmt.Println(err)
	}
	pic2, err := os.Create("C:\\Users\\ryogen\\Desktop\\ミリグラ\\0登録済み\\" + rank + eventname + "_dif.png")
	if err != nil {
		fmt.Println(err)
	}

	//ファイルに書き込む
	png.Encode(pic1, border)
	png.Encode(pic2, dif)

	//閉じる
	pic1.Close()
	pic2.Close()

	//終了のお知らせ
	fmt.Println("complete!")
	c <- false
}

func main() {
	//イベント名入力
	eventname := ""
	fmt.Print("グラフを作るイベントの名前\nスペースはめんどくさいから代わりにアンダーバーにして → ")
	fmt.Scan(&eventname)

	//色は指定してもらう
	var r, g, b byte
	fmt.Print("線の色入力\nR, G, Bで指定して\nR → ")
	fmt.Scan(&r)
	fmt.Print("G → ")
	fmt.Scan(&g)
	fmt.Print("B → ")
	fmt.Scan(&b)

	filename := "C:\\Users\\ryogen\\Desktop\\ミリグラ\\" + eventname
	go write(filename+"\\100.txt", eventname, "100\\", r, g, b)
	go write(filename+"\\2500.txt", eventname, "2500\\", r, g, b)
	<-c
	<-c

	//ディレクトリを隅に移動させる
	if err := os.Rename(filename, "C:\\Users\\ryogen\\Desktop\\ミリグラ\\0登録済み\\"+eventname); err != nil {
		fmt.Println(err)
	}
}
