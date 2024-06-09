package os_exit_analyzer

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestOsExitChecker(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, OsExitChecker, "./...")
}
