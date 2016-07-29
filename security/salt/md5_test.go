package salt

import "testing"

func TestAddMD5Salt(t *testing.T) {
	const SALT string = "LDbrzZtU"
	const VALUE string = "Test测试123ABC"
	if s := AddMD5Salt(SALT, VALUE); s != "4f4e07078de261ac55cef16a9972c979" {
		t.Fatal("Error AddMD5Salt:", s)
	}
}

func TestVerifyMD5Salt(t *testing.T) {
	const SALT string = "LDbrzZtU"
	const VALUE string = "Test测试123ABC"
	if s := AddMD5Salt(SALT, VALUE); VerifyMD5Salt(SALT, s, "4f4e07078de261ac55cef16a9972c979") {
		t.Fatal("Error VerifyMD5Salt:", s)
	}
}
