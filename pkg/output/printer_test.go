package output

import "testing"

func Test_interface(t *testing.T) {
	var _ Interface = &Printer{}
}
