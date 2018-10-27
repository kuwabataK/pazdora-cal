package main

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
)

func main() {

	num_ok, num_ng, x_range, y_range := 0, 0, 0, 0

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 5 || d >= 5
	})
	print_prob(num_ok, num_ng, "2色いずれかが5色以上ある確率", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 3 && d >= 3
	})
	print_prob(num_ok, num_ng, "指定2色がある確率", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d, r}
		sort.Ints(array)
		return array[1] >= 3
	})
	print_prob(num_ok, num_ng, "6色のうち5色がある確率", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d}
		sort.Ints(array)
		return array[1] >= 3
	})
	print_prob(num_ok, num_ng, "5色のうち4色がある確率", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 6
	})
	print_prob(num_ok, num_ng, "指定色が6個以上ある確率(ヨグとか)", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 3
	})
	print_prob(num_ok, num_ng, "指定色が3個以上ある確率", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		return f >= 5
	})
	print_prob(num_ok, num_ng, "指定色が5個以上ある確率", x_range, y_range)

	num_ok, num_ng, x_range, y_range = monte_carlo_freq(func(f int, b int, g int, l int, d int, r int) bool {
		array := []int{f, b, g, l, d, r}
		sort.Ints(array)
		return array[0] >= 3
	})
	print_prob(num_ok, num_ng, "全色ある確率", x_range, y_range)

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

func cnt_drops(target_num int, drops [][]int) int {

	total := 0

	for i := range drops {
		for j := range drops[i] {
			if drops[i][j] == target_num {
				total++
			}
		}
	}

	return total
}

func monte_carlo_freq(fn func(fire_drops_num int, blue_drops_num int, green_drops_num int,
	light_drops_num int, black_drops_num int, recovery_drops_num int) bool) (int, int, int, int) {

	loop_cnt := 100000
	x_range := 5
	y_range := 6
	num_ok := 0
	num_ng := 0

	for i := 0; i < loop_cnt; i++ {
		drops := generate_drops(x_range, y_range)
		if !check_normal_drops(drops) {
			continue
		}

		f_num := cnt_drops(0, drops)
		blu_num := cnt_drops(1, drops)
		g_num := cnt_drops(2, drops)
		l_num := cnt_drops(3, drops)
		bla_num := cnt_drops(4, drops)
		r_num := cnt_drops(5, drops)

		if fn(f_num, blu_num, g_num, l_num, bla_num, r_num) {
			num_ok++
		} else {
			num_ng++
		}

	}

	return num_ok, num_ng, x_range, y_range

}

func print_prob(num_ok int, num_ng int, message string, x int, y int) {
	fmt.Println("===============")
	fmt.Println(message)

	fmt.Println(strconv.Itoa(x) + strconv.Itoa(y) + "盤面")
	fmt.Println("試行回数は" + strconv.Itoa(num_ok+num_ng))
	fmt.Println("成功した回数は" + strconv.Itoa(num_ok))
	fmt.Println("失敗した回数は" + strconv.Itoa(num_ng))
	fmt.Println("確率は")
	kakuritu := float64(num_ok) / float64(num_ng+num_ok) * 100
	fmt.Println(kakuritu)
}
