package _commonUtils

import (
	"github.com/jinzhu/copier"
)

func BeanCopy(toValue interface{}, fromValue interface{}) {
	copier.Copy(&toValue, fromValue)
}
