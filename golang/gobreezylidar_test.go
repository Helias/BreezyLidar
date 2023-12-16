package gobreezylidar

import (
	"fmt"
	_ "fmt"
	"testing"
)

func TestNum(t *testing.T) {
	breezyLidar := New()

	fmt.Print("############################ \n START SCAN \n############################\n")

	breezyLidar.getScan()

	fmt.Print("############################ \n END SCAN \n############################\n")
}
