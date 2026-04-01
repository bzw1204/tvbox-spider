package utils

import (
	"testing"
)

func TestGzip(t *testing.T) {
	str := "Hello World! This is a test string for gzip compression."
	compressed, err := Gzip(str)
	if err != nil {
		t.Fatalf("Gzip failed: %v", err)
	}
	t.Log("compressed:", compressed)

	if compressed == "" {
		t.Error("Gzip returned empty string")
	}
}

func TestUnGzip(t *testing.T) {
	// 先压缩一个字符串
	original := "Hello World! This is a test string for gzip compression."
	compressed, err := Gzip(original)
	if err != nil {
		t.Fatalf("Gzip failed: %v", err)
	}

	// 然后解压缩
	decompressed, err := UnGzip(compressed)
	if err != nil {
		t.Fatalf("UnGzip failed: %v", err)
	}

	if decompressed != original {
		t.Errorf("UnGzip failed: expected %q, got %q", original, decompressed)
	}

	t.Log("decompressed:", decompressed)
}

func TestGzipAndUnGzip_RoundTrip(t *testing.T) {
	testCases := []string{
		"",
		"Hello",
		"Hello World!",
		"This is a longer string with more content to test the compression and decompression functionality.",
		"特殊字符测试！@#$%^&*()_+{}|:<>?",
		"中文测试字符串，包括一些特殊的 Unicode 字符。",
		"1234567890",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyz",
	}

	for i, tc := range testCases {
		t.Run(string(rune(i)), func(t *testing.T) {
			compressed, err := Gzip(tc)
			if err != nil {
				t.Fatalf("Gzip failed for test case %d: %v", i, err)
			}
			decompressed, err := UnGzip(compressed)
			if err != nil {
				t.Fatalf("UnGzip failed for test case %d: %v", i, err)
			}

			if decompressed != tc {
				t.Errorf("Round trip failed for test case %d: expected %q, got %q", i, tc, decompressed)
			}
		})
	}
}

func TestGzip_EmptyString(t *testing.T) {
	empty := ""
	compressed, err := Gzip(empty)
	if err != nil {
		t.Fatalf("Gzip of empty string failed: %v", err)
	}
	if compressed == "" {
		t.Error("Gzip of empty string should not return empty string")
	}

	decompressed, err := UnGzip(compressed)
	if err != nil {
		t.Fatalf("UnGzip of empty string failed: %v", err)
	}
	if decompressed != empty {
		t.Errorf("UnGzip of empty string failed: expected %q, got %q", empty, decompressed)
	}
}

func TestGzip_LargeString(t *testing.T) {
	// 创建一个较大的字符串
	largeStr := ""
	for i := 0; i < 1000; i++ {
		largeStr += "This is line " + string(rune(i)) + " of a large string. "
	}

	compressed, err := Gzip(largeStr)
	if err != nil {
		t.Fatalf("Gzip of large string failed: %v", err)
	}
	decompressed, err := UnGzip(compressed)
	if err != nil {
		t.Fatalf("UnGzip of large string failed: %v", err)
	}

	if decompressed != largeStr {
		t.Error("Large string round trip failed")
	}

	t.Logf("Original size: %d bytes, Compressed size: %d bytes", len(largeStr), len(compressed))
}
