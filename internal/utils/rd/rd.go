package rd

import "github.com/hashicorp/terraform-plugin-testing/helper/acctest"

type RandomData struct {
	Int int
	Str string
}

func NewRandomData(strlen int) RandomData {
	return RandomData{
		Int: acctest.RandInt(),
		Str: acctest.RandString(strlen),
	}
}
