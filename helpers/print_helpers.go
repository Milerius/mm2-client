package helpers

import "github.com/kyokomi/emoji/v2"

func PrintCheck(str string, isCompleted bool) {
	if isCompleted {
		_, _ = emoji.Println(str + " :white_check_mark:")
	} else {
		_, _ = emoji.Println(str + " :x:")
	}
}

func TransformBool(val bool) string {
	if val {
		return emoji.Sprintf(":white_check_mark:")
	} else {
		return emoji.Sprintf(":x:")
	}
}
