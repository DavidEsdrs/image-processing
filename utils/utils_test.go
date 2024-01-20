package utils_test

import (
	"testing"

	"github.com/DavidEsdrs/image-processing/utils"
)

func TestGaussianKernel(t *testing.T) {
	res := utils.GaussianKernel(5, 1.0)
	t.Logf("%v", res)
}
