package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	i := 1

	/* 絶対値 */
	j := math.Abs(float64(i))

	/* 切り上げ */
	j := math.Ceil(i)

	/* 切り捨て */
	j := math.Trunc(i)

	/* 四捨五入 */
	j := math.Round(i)

	/* 平方根 */
	j := math.Sqrt(i)

	/* N乗 */
	j := math.pow(i, N)

	/* 平方根 */
	j := math.Sqrt(i)

	/* 乱数 */
	rand.Seed(time.Now().Unix())
	j := rand.Intn(N)       // 0 ~ N
	j := rand.Float32()     // 0.0 ~ 1.0
	fmt.Printf("%.Nf\n", j) // 四捨五入して小数第N桁まで出力
}
