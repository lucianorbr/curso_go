package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em aquitetura amd64")
	}
	t.Fail()
}
