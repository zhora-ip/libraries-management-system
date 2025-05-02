package pkg

import "fmt"

func AddCondition(v *int64, args *[]any, str string, i *int) string {
	if v != nil {
		cond := fmt.Sprintf(str, *i)
		*i++
		*args = append(*args, fmt.Sprint(*v))
		return cond
	}
	return ""
}
