package filters_test

import (
	"testing"

	"github.com/DavidEsdrs/image-processing/filters"
	"github.com/DavidEsdrs/image-processing/logger"
	"github.com/DavidEsdrs/image-processing/utils"
)

func TestBlur(t *testing.T) {
	bf, err := filters.NewBlurFilter(logger.Logger{}, 1, 7)
	if err != nil {
		t.FailNow()
	}
	img, err := utils.LoadImage("../images/almoço.png")
	if err != nil {
		t.FailNow()
	}
	tensor := utils.ConvertIntoTensor(img)
	err = bf.Execute(&tensor)
	if err != nil {
		t.FailNow()
	}
}
