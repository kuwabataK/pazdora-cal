package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
)

func main() {

	x := flag.Int("x", 5, "盤面の縦方向の大きさ")
	y := flag.Int("y", 6, "盤面の横方向の大きさ")
	l := flag.Int("loopCnt", 100000, "ループする回数")
	flag.Parse()
	x_range := *x
	y_range := *y
	loop_cnt := *l

	num_ok, num_ng := 0, 0

	field := generate_fields(x_range, y_range, loop_cnt)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 5 || d >= 5
	}, field)
	print_prob(num_ok, num_ng, "2色いずれかが5色以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 3 && d >= 3
	}, field)
	print_prob(num_ok, num_ng, "指定2色がある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d, r}
		sort.Ints(array)
		return array[1] >= 3
	}, field)
	print_prob(num_ok, num_ng, "6色のうち5色がある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d}
		sort.Ints(array)
		return array[1] >= 3
	}, field)
	print_prob(num_ok, num_ng, "5色のうち4色がある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 6
	}, field)
	print_prob(num_ok, num_ng, "指定色が6個以上ある確率(ヨグとか)", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 3
	}, field)
	print_prob(num_ok, num_ng, "指定色が3個以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 5
	}, field)
	print_prob(num_ok, num_ng, "指定色が5個以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d, r}
		sort.Ints(array)
		return array[0] >= 3
	}, field)
	print_prob(num_ok, num_ng, "全色ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d}
		sort.Ints(array)
		return array[0] >= 3
	}, field)
	print_prob(num_ok, num_ng, "ガードブレイクが発動できる確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 8
	}, field)
	print_prob(num_ok, num_ng, "指定一色が8個以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 4
	}, field)
	print_prob(num_ok, num_ng, "指定一色が4個以上ある確率 進化前リーチェ", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 7
	}, field)
	print_prob(num_ok, num_ng, "7コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 8
	}, field)
	print_prob(num_ok, num_ng, "8コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 9
	}, field)
	print_prob(num_ok, num_ng, "9コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 10
	}, field)
	print_prob(num_ok, num_ng, "10コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 11
	}, field)
	print_prob(num_ok, num_ng, "11コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 12
	}, field)
	print_prob(num_ok, num_ng, "12コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 13
	}, field)
	print_prob(num_ok, num_ng, "13コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/3)+(b/3)+(g/3)+(l/3)+(d/3)+(r/3) >= 14
	}, field)
	print_prob(num_ok, num_ng, "14コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/4)+(b/4)+(g/4)+(l/4)+(d/4)+(r/4) >= 5
	}, field)
	print_prob(num_ok, num_ng, "(4個消しで)5コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/4)+(b/4)+(g/4)+(l/4)+(d/4)+(r/4) >= 6
	}, field)
	print_prob(num_ok, num_ng, "(4個消しで)6コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/4)+(b/4)+(g/4)+(l/4)+(d/4)+(r/4) >= 7
	}, field)
	print_prob(num_ok, num_ng, "(4個消しで)7コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/4)+(b/4)+(g/4)+(l/4)+(d/4)+(r/4) >= 8
	}, field)
	print_prob(num_ok, num_ng, "(4個消しで)8コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/4)+(b/4)+(g/4)+(l/4)+(d/4)+(r/4) >= 9
	}, field)
	print_prob(num_ok, num_ng, "(4個消しで)9コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return (f/4)+(b/4)+(g/4)+(l/4)+(d/4)+(r/4) >= 10
	}, field)
	print_prob(num_ok, num_ng, "(4個消しで)10コンボ以上ある確率", x_range, y_range)

	num_ok, num_ng = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{b, g, l, d, r}
		sort.Ints(array)
		return f >= 4 && array[4] >= 5
	}, field)
	print_prob(num_ok, num_ng, "指定一色が４つ以上あり、指定色以外が５個以上ある確率(ユウキ（SAO）)", x_range, y_range)
}

// 引数に指定したドロップたちが3つ以上つながっていないことを確認
func check_normal_drops(drops [][]int) bool {

	for i := range drops {
		for j := range drops[i] {
			if j+2 < len(drops[i]) {
				if drops[i][j] == drops[i][j+1] && drops[i][j] == drops[i][j+2] {
					return false
				}
			}

			if i+2 < len(drops) {
				if drops[i][j] == drops[i+1][j] && drops[i][j] == drops[i+2][j] {
					return false
				}
			}

		}
	}

	return true
}

// ドロップを生成する
func generate_drops(x int, y int) [][]int {
	drops := make([][]int, y)
	for i := range drops {
		drops[i] = make([]int, x)
	}
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	for i := range drops {
		for j := range drops[i] {
			drops[i][j] = rand.Intn(6)
		}
	}
	return drops
}

func cnt_drops(target_drop int, drops [][]int) int {

	total := 0

	for i := range drops {
		for j := range drops[i] {
			if drops[i][j] == target_drop {
				total++
			}
		}
	}

	return total
}

// 盤面の配列に対して、指定したbool式に一致する盤面、一致しない盤面の数を返します
//
func monte_carlo_freq(fn func(fire_drops_num int, blue_drops_num int, green_drops_num int,
	light_drops_num int, black_drops_num int, recovery_drops_num int) bool, fields [][]int) (int, int) {

	num_ok := 0
	num_ng := 0

	for i := range fields {

		if fn(fields[i][0], fields[i][1], fields[i][2], fields[i][3], fields[i][4], fields[i][5]) {
			num_ok++
		} else {
			num_ng++
		}

	}

	return num_ok, num_ng

}

// loop_cntの数だけランダムなパズドラ盤面の生成を試みます。
// 3つ以上ドロップがつながっている盤面は破棄され、つながっていない盤面のみ保持されます
func generate_fields(x_range int, y_range int, loop_cnt int) [][]int {

	field := make([][]int, 0, loop_cnt/2)

	for i := 0; i < loop_cnt; i++ {
		drops := generate_drops(x_range, y_range)
		if !check_normal_drops(drops) {
			continue
		}

		f := cnt_drops(0, drops)
		b := cnt_drops(1, drops)
		g := cnt_drops(2, drops)
		l := cnt_drops(3, drops)
		d := cnt_drops(4, drops)
		r := cnt_drops(5, drops)

		do := []int{f, b, g, l, d, r}

		field = append(field, do)
	}

	return field

}

func print_prob(num_ok int, num_ng int, message string, x int, y int) {
	fmt.Println("===============")
	fmt.Println(message)

	fmt.Println(strconv.Itoa(x) + strconv.Itoa(y) + "盤面")
	fmt.Println("試行回数は" + strconv.Itoa(num_ok+num_ng))
	fmt.Println("成功した回数は" + strconv.Itoa(num_ok))
	fmt.Println("失敗した回数は" + strconv.Itoa(num_ng))
	kakuritu := float64(num_ok) / float64(num_ng+num_ok) * 100
	fmt.Println("確率は" + strconv.FormatFloat(kakuritu, 'f', 10, 64))
}
