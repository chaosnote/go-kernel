package conv

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//------------------------------------------------------------------------------------------------------------[conv.String]

// String ...
type String string

func (v String) toFloat(p int) (float64, error) {
	f, e := strconv.ParseFloat(string(v), p)
	if e == nil {
		return f, nil
	}
	return 0, e
}

// ToFloat32 ...
func (v String) ToFloat32() (float32, error) {
	f, e := v.toFloat(32)
	return float32(f), e
}

// ToFloat64 ...
func (v String) ToFloat64() (float64, error) {
	f, e := v.toFloat(32)
	return float64(f), e
}

func (v String) toInt(p int) (int64, error) {
	i, e := strconv.ParseInt(string(v), 10, p)
	if e == nil {
		return i, nil
	}
	return 0, e
}

// ToInt ...
func (v String) ToInt() (int, error) {
	i, e := v.toInt(64)
	return int(i), e
}

// ToInt32 ...
func (v String) ToInt32() (int32, error) {
	i, e := v.toInt(32)
	return int32(i), e
}

// ToInt64 ...
func (v String) ToInt64() (int64, error) {
	i, e := v.toInt(64)
	return i, e
}

// ToIntBytes ...
func (v String) ToIntBytes() ([]byte, error) {
	list := strings.Split(string(v), "")
	dist := []byte{}

	for _, s := range list {
		i, e := String(s).ToInt64()
		if e != nil {
			return []byte{}, e
		}
		dist = append(dist, byte(i))
	}
	return dist, nil
}

//------------------------------------------------------------------------------------------------------------[conv.Uint64]

// Uint64 ...
type Uint64 uint64

// ToString ...
func (v Uint64) ToString() string {
	return strconv.FormatUint(uint64(v), 10)
}

//------------------------------------------------------------------------------------------------------------[conv.Int]

// Int ...
type Int int

// Left 往左補 0 至 N 位數
func (v Int) Left(padding int) string {
	f := "%0" + fmt.Sprintf("%d", padding) + "d"
	n := fmt.Sprintf(f, int(v))
	if len(n) != padding {
		log.Printf("padding.Left: %d %d \n", v, padding)
		return ""
	}
	return n
}

//------------------------------------------------------------------------------------------------------------[conv.Int64]

// Int64 ...
type Int64 int64

// ToString ...
func (v Int64) ToString() string {
	return strconv.FormatInt(int64(v), 10)
}

//------------------------------------------------------------------------------------------------------------[conv.Float64]

// Float64 ...
type Float64 float64

// ToString ...
// prec 小數點後位數( 有效長度 0 - 4 )
func (v Float64) ToString(prec uint8) string {
	if prec > 4 {
		prec = 4
	}
	return strconv.FormatFloat(float64(v), 'f', int(prec), 64)
}

// RoundTo 四捨五入 {p:小數點後幾位}
func (v Float64) RoundTo(p uint32) float64 {
	n := float64(v)
	return math.Round(n*math.Pow(10, float64(p))) / math.Pow(10, float64(p))
}

//------------------------------------------------------------------------------------------------------------[conv.Float64Array]

// Float64Array ...
type Float64Array []float64

// First 取得第一筆值(未處理空值)
func (v Float64Array) First() float64 {
	return v[0]
}

// Last 取得最後一筆值(未處理空值)
func (v Float64Array) Last() float64 {
	return v.LastN(1)
}

// LastN 取得最後N筆位置的值(未處理空值)
func (v Float64Array) LastN(n int) float64 {
	return v[len(v)-n]
}
