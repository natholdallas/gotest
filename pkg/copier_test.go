package pkg

import (
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

type CopierModel struct {
	ID uint
}

type CopierTestModel struct {
	Model
	Username string
	Password string
}

type CopierTestModelIn struct {
	ID       uint `copier:"Model"`
	Username string
	Password string
}

func TestCopier(t *testing.T) {
	test := CopierTestModel{}
	testIn := CopierTestModelIn{1, "Username", "Password"}
	copier.Copy(&test, testIn)
	fmt.Println(Marshal(test))
}
