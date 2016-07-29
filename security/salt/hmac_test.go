package salt

import "testing"

func TestAddHMACMD5Salt(t *testing.T) {
	const SALT string = "LDbrzZtU"
	const VALUE string = "Test测试123ABC"
	if s := AddHMACMD5Salt(SALT, VALUE); s != "13fc1e65074fa2745796ca0cab04ee05" {
		t.Fatal("Error AddHMACMD5Salt:", s)
	}
}

func TestVerifyHMACMD5Salt(t *testing.T) {
	const SALT string = "LDbrzZtU"
	const VALUE string = "Test测试123ABC"
	if s := AddHMACMD5Salt(SALT, VALUE); VerifyHMACMD5Salt(SALT, s, "13fc1e65074fa2745796ca0cab04ee05") {
		t.Fatal("Error VerifyHMACMD5Salt:", s)
	}
}
