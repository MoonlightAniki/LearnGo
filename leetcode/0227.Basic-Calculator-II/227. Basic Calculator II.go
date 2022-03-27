package solution0227

// 执行用时： 84 ms, 在所有 Go 提交中击败了 5.37% 的用户
/*func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	operators := []string{"+", "-", "*", "/"}
	arr := []string{s}
	for _, op := range operators {
		arr = split(arr, op)
	}
	nums := make([]int, 0)
	ops := make([]string, 0)
	i := 0
	for i < len(arr) {
		switch arr[i] {
		case "+":
			ops = append(ops, "+")
			i++
		case "-":
			ops = append(ops, "-")
			i++
		case "*":
			num, _ := strconv.Atoi(arr[i+1])
			nums[len(nums)-1] *= num
			i += 2
		case "/":
			num, _ := strconv.Atoi(arr[i+1])
			nums[len(nums)-1] /= num
			i += 2
		default:
			num, _ := strconv.Atoi(arr[i])
			nums = append(nums, num)
		}
	}
	res := nums[0]
	for j, op := range ops {
		if op == "+" {
			res += nums[j+1]
		} else {
			res -= nums[j+1]
		}
	}
	return res
}

func split(ss []string, splitter string) []string {
	res := make([]string, 0)
	for _, s := range ss {
		splits := strings.Split(s, splitter)
		for i, item := range splits {
			if i != 0 {
				res = append(res, splitter)
			}
			res = append(res, item)
		}
	}
	return res
}*/

func calculate(s string) int {
	stack := make([]int, 0)
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = 10*num + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	res := 0
	for _, num := range stack {
		res += num
	}
	return res
}
