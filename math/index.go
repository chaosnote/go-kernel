package math

import (
	"kernel/conv"
	"regexp"
)

// 預設 Pow10 放大倍率
const magnification int = 2

//------------------------------------------------------------------------------------------------------------[math.Decimal]

// Decimal ...
type Decimal int

// Add + ...p
func (v Decimal) Add(p ...Decimal) Decimal {
	var total Decimal = v
	for _, d := range p {
		total = total + d
	}
	return total
}

// Sub - p
func (v Decimal) Sub(p Decimal) Decimal {
	return v - p
}

// Mul * p
func (v Decimal) Mul(p Decimal) Decimal {
	return v * p
}

// Div / p (1/100)=>0 待思考是否保留
func (v Decimal) Div(p Decimal) Decimal {
	return v / p
}

// ToInt ...
func (v Decimal) ToInt() int {
	return int(v)
}

// ToString ...
func (v Decimal) ToString() string {
	return conv.Int64(v).ToString()
}

//------------------------------------------------------------------------------------------------------------[math.DecimalCollection]

// IntCollection ...
type IntCollection []int

// Accumulate 累加 (例) [0 1 2 3 4] -> [0 1 3 6 10]
func (v IntCollection) Accumulate() []int {
	total := 0
	list := []int{}

	for _, i := range v {
		total = total + i
		list = append(list, total)
	}
	return list
}

// ToStringCollection ...
func (v IntCollection) ToStringCollection() []string {
	list := []string{}

	for _, i := range v {
		list = append(list, conv.Int64(i).ToString())
	}
	return list
}

//------------------------------------------------------------------------------------------------------------[math.String]

// String ...
type String string

// ToDecimal ...
func (v String) ToDecimal() (Decimal, error) {
	i, e := conv.String(v).ToInt64()
	if e != nil {
		return 0, e
	}
	return Decimal(i), nil
}

// IsDecimal ...
func (v String) IsDecimal() (bool, error) {
	// regex group
	// https://golang.org/pkg/regexp/syntax/

	// test 1.
	// rule := `([-]?\d+)(\.\d+)?`
	// val := "-123.123"
	// _, err := regexp.MatchString(rule, val) // 正則測試
	// if err != nil {
	// 	log.Get().Info(err)
	// 	return
	// }

	// test 2.
	// rule := `([-]?\d+)(\.\d{1,2})?`
	// rule := `([-]?\d+)(?P<name>\.\d+)?`
	// m, err := regexp.Compile(rule)
	// log.Get().Debug(m.SubexpNames())
	// log.Get().Debug(m.FindStringSubmatch(price))

	rule := `([-]?\d+)(\.\d{1,2})?`

	m, err := regexp.Compile(rule)
	if err != nil {
		return false, err
	}

	if m.FindStringSubmatch(string(v))[0] != string(v) {
		return false, nil
	}

	return true, nil
}
