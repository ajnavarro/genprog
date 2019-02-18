package utils

import (
	"math/rand"
	"sync"
	"time"

	"github.com/ajnavarro/genprog/node"
)

var _random = rand.New(rand.NewSource(time.Now().UnixNano()))
var mutex = &sync.Mutex{}

func getRandom() *rand.Rand {
	mutex.Lock()
	defer mutex.Unlock()
	return _random
}

func pickOneNode(fn []node.Node) node.Node {
	if len(fn) <= 0 {
		panic("CATCHED 22222222")
	}
	idx := getRandom().Intn(len(fn))
	return fn[idx]
}

func pickOneFactory(fn []*node.FactoryNode) *node.FactoryNode {
	if len(fn) <= 0 {
		panic("CATCHED")
	}
	return fn[getRandom().Intn(len(fn))]
}
