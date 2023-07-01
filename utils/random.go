package util

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
)
const alpha = "aqwzsxedcrfvtgbyhnujikolpm"
func init() {
	rand.Seed(time.Now().UnixNano())

}
// nombre alea entre min et max 
func RandomInt(min , max int64) int64 {
	return min+ rand.Int63n(max-min+1)
}

func RandomString(n int ) string {
var sb strings.Builder
k := len(alpha)
for i:=0;i< n ;i++{
	c := alpha[rand.Intn(k)]
	sb.WriteByte(c)
} 
return sb.String()


}

func RandOwner() string {
	return RandomString(6)
}

func RandMoney() int64 {
	return RandomInt(0 , 1000)
}

func RandCurrency() string {
	currencies := []string{"euro" , "TND" , "usd "}
	n := len (currencies)
	return currencies[rand.Intn(n)]
}


// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}