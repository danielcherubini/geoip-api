package ip

import (
	"fmt"
	"testing"
)

func TestGetIP(t *testing.T) {
	singleIP := "193.215.2.26"
	singleIPString := GetIP(singleIP)
	fmt.Println(singleIPString)

	xFFor := "193.215.2.26, 193.215.2.25"
	xFForString := GetIP(xFFor)
	fmt.Println(xFForString)

}
