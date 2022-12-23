package oracle

import "testing"

func Test(t *testing.T) {
	oraclePlugin := OraclePlugin{}
	if !oraclePlugin.Enabled() {
		t.Fatal("oracle default enabled")
	}
	if oraclePlugin.Order() != defaultConfig["order"] {
		t.Fatalf("oracle order should be %v", defaultConfig["order"])
	}

}
