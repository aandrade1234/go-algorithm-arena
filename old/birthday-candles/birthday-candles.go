package main

func main() {
}

// time complexity O(n) space complexity O(1)
func solution2(inputs any) any {
	candles := inputs.([]int)

	max := 0
	count := 0

	for _, height := range candles {
		if height > max {
			max = height
			count = 1
		} else if height == max {
			count++
		}
	}

	return count
}

// time complexity O(n) space complexity O(n)
func solution1(inputs any) any {
	candles := inputs.([]int)

	tallestCandle := 0
	cake := make(map[int]int)

	for _, candle := range candles {
		cake[candle] = cake[candle] + 1

		if candle > tallestCandle {
			tallestCandle = candle
		}
	}

	return cake[tallestCandle]
}

func (*BirthdayCandles) GetInputs() []any {
	return []any{
		[]int{3, 2, 1, 3},
		[]int{18, 90, 90, 13, 90, 75, 90, 8, 90, 43},
		[]int{82, 49, 82, 82, 41, 82, 15, 63, 38, 25},
		[]int{44, 53, 31, 27, 77, 60, 66, 77, 26, 36},
	}
}

func (*BirthdayCandles) GetOutputs() []any {
	return []any{2, 5, 4, 2}
}
