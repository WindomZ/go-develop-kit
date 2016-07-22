package googleauth

import "testing"

func TestGenerateRandomSecret(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if !ValidSecret(GenerateRandomSecret(20, true), 20, true) {
			t.Fatal("Fail to valid secret")
		}
	}
}
