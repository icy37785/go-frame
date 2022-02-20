package util

import "github.com/jinzhu/copier"

func Copy(dst, src interface{}) error {
	return copier.Copy(src, dst)
}
