package label

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
	label := NewLabel(false)

	t.Logf("Next test should pass")
	t.Logf("%sTest that passed", label.pass)

	fmt.Println()

	t.Logf("Next test should fail")
	t.Logf("%sTest that failed", label.fail)

	fmt.Println()

	t.Logf("Next test should have info")
	t.Logf("%sTest that displayed some info", label.info)

	fmt.Println()

	t.Logf("Next test should have warning")
	t.Logf("%sTest that generated warning", label.warn)
}
