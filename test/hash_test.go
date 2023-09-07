package util_test

import (
	"fmt"
	"testing"

	"github.com/MephistoSolsist/mysql-practice/util"
)

func Test_BcryptHash(t *testing.T) {
	a := util.BcryptHash("hello")
	b := util.BcryptHash("hello")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(util.BcryptCheck("hello", b))
}
