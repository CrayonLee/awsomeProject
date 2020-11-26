package util

import "math/rand"

func RandString(n int) string{
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	ret:=make([]byte,n)

	for i := range ret {
		ret[i]=letters[rand.Intn(len(letters))]
	}
	return string(ret)
}
