package main
import (
	"fmt"
	"strconv"
)

func prima(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true

}

func valid(n int) bool {
	s := strconv.Itoa(n)
	jumlah := 0

	for i := 0; i < len(s); i++ {
		jumlah += int(s[i] - '0')

		if i > 0 && s[i] < s[i-1] {
			return false
		}
	}

	return prima(jumlah)
}

func main(){
	var L, R int
	var jumlah int
	fmt.Scan(&L, &R)

	for i := L; i <= R; i++ {
		if valid(i) {
			jumlah++
		}
	}

	fmt.Println(jumlah)
}