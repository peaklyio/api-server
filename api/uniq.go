package api

import (
	"fmt"
	"hash/fnv"
)

type UniqQuery struct {
	Uniq int
}

func U(m string, a ...interface{}) int {
	str := fmt.Sprintf(m, a...)
	h := fnv.New32a()
	h.Write([]byte(str))
	return int(h.Sum32())
}

func UQuery(m string, a ...interface{}) *UniqQuery {
	return &UniqQuery{
		Uniq: U(m, a...),
	}
}
