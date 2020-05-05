package main

import (
	"bufio"
	"fmt"
	"os"
)

type address struct {
	//コロン,月の場所、括弧,日の場所
	cm, ks int
}

type miss struct {
	//何個目、時間
	n int
	t string
}

var (
	//ck...投稿時間確認のためのチャネル
	//dt...データを待機させるためのチャネル
	//ts...テスト用のチャネル
	//end...ゴルーチン終了のお知らせのためのチャネル
	ck  chan string = make(chan string, 594)
	dt1 chan string = make(chan string, 594)
	dt2 chan string = make(chan string, 594)
	ts  chan int    = make(chan int, 594)
	end chan bool   = make(chan bool)

	//箱
	data_100  []byte = make([]byte, 0)
	data_2500 []byte = make([]byte, 0)

	//投稿時間カウント記録
	tim []rune = make([]rune, 24)

	//投稿時間カウント記録集
	times [][]rune = make([][]rune, 0)

	//欠落記録集
	misses []miss = make([]miss, 0)
)

//配列の中身の合計を求める
func add(arr []rune) (sum rune) {
	l := rune(len(arr))
	c := make(chan rune, 23)
	var s, e rune
	for i := rune(0); i < 23; i++ {
		s = l / 23 * i
		e = l / 23 * (i + 1)
		if i == 22 {
			e = l
		}
		go func(arr []rune, s, e rune) {
			var tmp rune
			for j := s; j < e; j++ {
				tmp += arr[j]
			}
			c <- tmp
		}(arr, s, e)
	}
	for i := 0; i < 23; i++ {
		sum += <-c
	}
	return
}

func check() {
	adrs := address{}

	//投稿時間の個数
	l := len(ck)

	//日付用の場所
	var day, da, d rune

	for i := 0; i < l; i++ {

		//受け取った文字列をrune配列に変換
		rn := []rune(<-ck)

		//"月"と"日"の場所取得
		for j, v := range rn {

			//"月"の文字コードは26376
			if v == 26376 {
				adrs.cm = j

				//"日"の文字コードは26085
			} else if v == 26085 {
				adrs.ks = j
				break
			}
		}

		//日付取得
		//2桁の場合
		if len(rn[adrs.cm+1:adrs.ks]) == 2 {
			da = rn[adrs.cm+1]*10 + rn[adrs.cm+2]

			//1桁の場合
		} else {
			da = rn[adrs.cm+1] + 480
		}

		//初日用の処理
		if day == 0 {
			day = da
		}

		//日付が変わったらtimesに追加してtimeリセット
		if day != da {
			times = append(times, tim)
			tim = make([]rune, 24)
		}

		//カウント
		tim[int(rn[adrs.ks+5]*10+rn[adrs.ks+6]-528)]++

		//日付変化確認
		day = da
	}

	//最終日のカウントを追加
	times = append(times, tim)

	//何日分か
	l = len(times) - 1

	//投稿回数カウント
	day = 0

	//欠落チェック
	s := ""
	for i, v := range times {

		//一日の投稿回数の合計
		da = add(v)

		//初日は17、最終日は42、他は48個ならおｋ
		if i == 0 && da < 17 {
			//抜けてるところを探す
			for j, w := range v {

				//15時までは関係ないからスルー
				if j < 15 {
					continue
				}

				//投稿回数を数える
				d += w

				//15時半がなかったら
				if j == 15 && w == 0 {
					s = "//1日目の15時"
					misses = append(misses, miss{-1, s})

					//jがlen-1より小さいとき
				} else if j < len(v)-1 {

					//一個だけで、その次でカバーされてなかったら
					if w == 1 && v[j+1] != 3 {

						//すでに追加されてて、最後に追加した.nと同じなら
						if len(misses) > 0 && int(d-1) == misses[len(misses)-1].n {
							misses[len(misses)-1].t += fmt.Sprintf("//%d日目の%d時－1個", i+1, j)

							//そうじゃなかったら普通に追加
						} else {
							misses = append(misses, miss{int(d - 1), fmt.Sprintf("//%d日目の%d時－1個", i+1, j)})
						}

						//ゼロ個やったら
					} else if w == 0 {
						if len(misses) > 0 && int(d-1) == misses[len(misses)-1].n {
							misses[len(misses)-1].t += fmt.Sprintf("//%d日目の%d時－2個", i+1, j)
						} else {
							misses = append(misses, miss{int(d - 1), fmt.Sprintf("//%d日目の%d時－2個", i+1, j)})
						}
					}
				}
			}

			//二日目以降も同じように処理
		} else if i < l && da < 48 {
			d = day
			for j, w := range v {
				d += w
				if j < len(v)-1 {
					if w == 1 && v[j+1] != 3 {
						if len(misses) > 0 && int(d-1) == misses[len(misses)-1].n {
							misses[len(misses)-1].t += fmt.Sprintf("//%d日目の%d時－1個", i+1, j)
						} else {
							misses = append(misses, miss{int(d - 1), fmt.Sprintf("//%d日目の%d時－1個", i+1, j)})
						}
					} else if w == 0 {
						if len(misses) > 0 && int(d-1) == misses[len(misses)-1].n {
							misses[len(misses)-1].t += fmt.Sprintf("//%d日目の%d時－2個", i+1, j)
						} else {
							misses = append(misses, miss{int(d - 1), fmt.Sprintf("//%d日目の%d時－2個", i+1, j)})
						}
					}
				}
			}
		} else if da < 42 {
			d = day

			//最終日は20時のチェックができたら終わり
			for j, w := range v {
				if j == 21 {
					break
				}
				d += w
				if j < len(v)-1 {
					if w == 1 && v[j+1] != 3 {
						if len(misses) > 0 && int(d-1) == misses[len(misses)-1].n {
							misses[len(misses)-1].t += fmt.Sprintf("//%d日目の%d時－1個", i+1, j)
						} else {
							misses = append(misses, miss{int(d - 1), fmt.Sprintf("//%d日目の%d時－1個", i+1, j)})
						}
					} else if w == 0 {
						if len(misses) > 0 && int(d-1) == misses[len(misses)-1].n {
							misses[len(misses)-1].t += fmt.Sprintf("//%d日目の%d時－2個", i+1, j)
						} else {
							misses = append(misses, miss{int(d - 1), fmt.Sprintf("//%d日目の%d時－2個", i+1, j)})
						}
					}
				}
			}
		}
		day += da
	}
}

func edit(data, noti string) (byt []byte) {
	adrs := address{}

	//ボーダーの文字列をbyte配列に変換
	byt = []byte(data)

	//コロンと括弧の場所を探して記録
	for i, v := range byt {

		//コロンの文字コードは58
		if v == 58 {
			adrs.cm = i

			//括弧の文字コードは40
		} else if v == 40 {
			adrs.ks = i
			break
		}
	}

	//それをもとに数値の部分だけを取り出す
	byt = byt[adrs.cm+2 : adrs.ks-1]

	//コンマを取り除く
	for i, v := range byt {

		//コンマの文字コードは44
		if v == 44 {
			byt = append(byt[:i], byt[i+1:]...)
		}
	}

	//欠落通知を追記
	if noti != "" {
		for _, v := range []byte(noti) {
			byt = append(byt, v)
		}
	}

	//改行コードを追加して返す
	byt = append(byt, 10)
	return
}

func pack(which bool) {
	var (
		l, t int
		not  string
	)

	//2500位の場合
	if which {

		//データの数を取得
		l = len(dt2)

		//記述用データ作成
		for i := 0; i < l; i++ {

			//欠落記録から引用していく
			if t < len(misses) {
				switch misses[t].n {
				case i:
					not = misses[t].t
					t++
				case -1:
					not = misses[0].t
					t++
				default:
					not = ""
				}
			} else {
				not = ""
			}

			//受け取ったデータとかを編集して箱詰め
			for _, v := range edit(<-dt2, not) {
				data_2500 = append(data_2500, v)
			}
		}

		//100位も同じ処理
	} else {
		l = len(dt1)
		for i := 0; i < l; i++ {
			if t < len(misses) {
				switch misses[t].n {
				case i:
					not = misses[t].t
					t++
				case -1:
					not = misses[0].t
					t++
				default:
					not = ""
				}
			} else {
				not = ""
			}
			for _, v := range edit(<-dt1, not) {
				data_100 = append(data_100, v)
			}
		}
	}
}

func mkwr(path, rank string) {

	//ファイル作成
	fl, err := os.Create(path + rank)

	//書いたら閉じる
	defer fl.Close()

	//作れなかったら
	if err != nil {
		fmt.Printf("can't create the file :(\n-----\n%v\n-----\n", err)
	}

	//記述
	if rank == "\\100.txt" {
		fl.Write(data_100)
	} else {
		fl.Write(data_2500)
	}

	//書き終わったら終了をおしらせ
	end <- false
}

func file(eventname string) {

	//先に"0\n"を入れとく
	data_100 = append(data_100, 48, 10)
	data_2500 = append(data_2500, 48, 10)

	//ファイルのパスを途中まで作る
	path := "C:\\Users\\ryogen\\Desktop\\ミリグラ\\" + eventname

	//ファイル開く
	file, err := os.Open(path + ".txt")

	//開けなかったら
	if err != nil {
		fmt.Printf("can't open the file :\\\n-----\n%v\n-----\n", err)
	}

	//スキャナー等準備
	scanner := bufio.NewScanner(file)

	//一行ずつ読み込む
	for scanner.Scan() {

		//改行だけじゃないとき
		if scanner.Text() != "" {
			//投稿時間の行やったら
			if scanner.Text()[4] == 229 {

				//欠落チェック
				ck <- scanner.Text()

				//最初の三文字が"100"やったら
			} else if scanner.Text()[:3] == "100" {

				//その行だけ待機場所に詰める
				dt1 <- scanner.Text()

				//最初の三文字が"2,5"でも100と同じ処理
			} else if scanner.Text()[:3] == "2,5" {
				dt2 <- scanner.Text()
			}
		}
	}

	//ここで閉じる
	file.Close()

	//欠落チェック
	check()

	//読み取ったデータの素を加工して詰める
	go pack(false)
	go pack(true)

	//イベ用のディレクトリ作成
	if err := os.Mkdir(path, 0777); err != nil {
		fmt.Println("\nthis event is probably registered\n")
	}

	//素のデータを上のディレクトリに引っ越し
	if err := os.Rename(path+".txt", path+"\\"+eventname+".txt"); err != nil {
		fmt.Printf("can't move the file :/\n-----\n%v\n-----\n", err)
	}

	//作って書く
	go mkwr(path, "\\100.txt")
	go mkwr(path, "\\2500.txt")

	//リンク張る
	fmt.Println("\n抜けてるところに時間と個数書いたから、確認して自分で埋めて\nhttps://mltd.matsurihi.me/events/\n\n↑ ここから追加するイベントを探して抜けてるところのデータを探す")
}

func main() {
	fmt.Println("漏れを防ぐために怪しいのを見つけるたびに確認してもらうのでお覚悟を\n")
	eventname := ""
	print("記録するイベントの名前\nスペースはめんどくさいから代わりにアンダーバーにして → ")
	fmt.Scan(&eventname)
	go file(eventname)
	<-end
	<-end
}
