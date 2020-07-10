package image

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetImage(t *testing.T) {
	image, err := GetImage("IF-ee265d5e-5027-40ce-9779-e515ab78fb79")
	assert.Nil(t, err)
	fmt.Println("result: ", image)

}
