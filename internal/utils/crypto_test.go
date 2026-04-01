package utils

import (
	"testing"
)

func TestAESDecrypt(t *testing.T) {
	encrypted := "bH5mI8iK0tK7aQ5x"
	str := "64LMkeF2DarK90YuRF/4DTvzJD09AczKOYYa60v8+XdE5Si/fF6WvyPQK0qrTWZFeYOopKlbH4sAa0c+Tqdm7QaVPd0oz0t3xEPSsLf46H/AMXwooquSEQe2KKWlCfIDEVcAi5CFxvLttBTReSJ9vd2vqSK39g+C7jMxbTI0eDDNcNHot1PLJ3cUOw2HNgteSQv7q7A8AwGY+sGmhvoSL97KNVodVICnFM7i/OLXDcrIVHiEcfYepLlkXawR/ohA0LnP3LM5sdTskUXLBeGu1WBhV8pew2+lsacVLqjgOh0IaWj2cAPdUEV89JdScMqIpzCQXLV+idJttEU2J2PGoBXbooAXtXilpvzyHV6QQzntghhMefjkW0OLjM0zNRXE8/WNqoflFELKKwOFYw1XQvQh83O753sbuvaWzvd0xbTrIxrRl70K82E7oByMj5jAQWXK3ydVET73B3L5sMBjKuIDlI7A5HZqDDDEhF9jFlnddtihnp5XP/91ybjPNP42z5cZGQaHyW3BXrR1Gszm3Z8TCf9leC3pHJLE1Ey/zSs8ODVl5XjSRBllFwAq18mFCUyPJouAmbcEwr8auYCClg=="
	decrypted, err := AESDecrypt(str, encrypted)
	if err != nil {
		t.Errorf("AESDecrypt failed: %v", err)
	}
	t.Log("decrypted:", decrypted)
}

func TestMD5(t *testing.T) {
	str := "123456"
	hash, _ := MD5Hash(str)
	t.Log("hash:", hash)
}

func TestBase64Encode(t *testing.T) {
	str := "Hello World!"
	encoded := Base64Encode([]byte(str))
	t.Log("encoded:", encoded)
}

func TestBase64Decode(t *testing.T) {
	str := "SGVsbG8gV29ybGQh"
	decoded, err := Base64Decode(str)
	if err != nil {
		t.Errorf("Base64Decode failed: %v", err)
	}
	t.Log("decoded:", decoded)
}
