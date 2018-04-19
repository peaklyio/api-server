package object

import (
	"hash/fnv"

	"fmt"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"gopkg.in/mgo.v2/bson"
)

func StringToBSONID(n string) bson.ObjectId {
	if len(n) < 2 {
		logger.Critical("Invalid ID!")
		return bson.NewObjectId()
	}
	a := fnv.New32a()
	a.Write([]byte(n))
	ten := a.Sum32()
	f := string(n[0])
	ff := string(n[1])
	twelve := fmt.Sprintf("%d%s%s", ten, f, ff)
	logger.Debug("%d", len([]byte(twelve)))
	return bson.ObjectIdHex(twelve)
}
