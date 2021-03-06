package label

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
	label := NewLabel(true)

	t.Logf("Next test should pass")
	t.Logf("%sTest that passed", label.Pass)

	fmt.Println()

	t.Logf("Next test should fail")
	t.Logf("%sTest that failed", label.Fail)

	fmt.Println()

	t.Logf("Next test should have info")
	t.Logf("%sTest that displayed some info", label.Info)

	fmt.Println()

	t.Logf("Next test should have warning")
	t.Logf("%sTest that generated warning", label.Warn)
}
