package loader

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	testTextFile = "/index.html"
	testImgFile  = "/image.png"
)

func TestCompression(t *testing.T) {
	t.Run("compressed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, testTextFile, nil)
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		res := httptest.NewRecorder()
		testLoader.ServeHTTP(res, req)
		if res.Code != http.StatusOK {
			t.Errorf("Not OK: %d", res.Code)
		}
		if res.Header().Get("Content-Encoding") != "gzip" {
			t.Errorf("Wrong encoding: %s", res.Header().Get("Content-Encoding"))
		}
		if res.Body.Len() != len(testLoader.content[testTextFile].CompressedBytes) {
			t.Errorf("Wrong content length, expected %d, got %d", res.Body.Len(), len(testLoader.content[testTextFile].CompressedBytes))
		}
	})
	t.Run("uncompressed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, testTextFile, nil)
		req.Header.Set("Accept", "*/*")
		res := httptest.NewRecorder()
		testLoader.ServeHTTP(res, req)
		if res.Code != http.StatusOK {
			t.Errorf("Not OK: %d", res.Code)
		}
		if res.Header().Get("Content-Encoding") != "" {
			t.Errorf("Wrong encoding: %s", res.Header().Get("Content-Encoding"))
		}
		if res.Body.Len() != len(testLoader.content[testTextFile].RawBytes) {
			t.Errorf("Wrong content length, expected %d, got %d", res.Body.Len(), len(testLoader.content[testTextFile].CompressedBytes))
		}
	})
	t.Run("uncompressable", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, testImgFile, nil)
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		res := httptest.NewRecorder()
		testLoader.ServeHTTP(res, req)
		if res.Code != http.StatusOK {
			t.Errorf("Not OK: %d", res.Code)
		}
		if res.Header().Get("Content-Encoding") != "" {
			t.Errorf("Wrong encoding: %s", res.Header().Get("Content-Encoding"))
		}
		if res.Body.Len() != len(testLoader.content[testImgFile].RawBytes) {
			t.Errorf("Wrong content length, expected %d, got %d", res.Body.Len(), len(testLoader.content[testTextFile].CompressedBytes))
		}
	})
}

func TestConditional(t *testing.T) {
	// Test If-Modified-Since 200 & 304
	// Test Etag/If-None-Match 200 & 304
}

func TestNotFound(t *testing.T) {
	// Test file with nonexistent path
	// Test folder with nonexistent path
}

func TestIndex(t *testing.T) {
	// Test / vs /index.html
	// Test when index does/doesn't exist
}

func TestRange(t *testing.T) {
	// Test range with and without compression
}

func TestContentType(t *testing.T) {
	// Test content types
}

var testLoader = New()

func init() {
	testLoader.Add(&Content{
		Path:     testImgFile,
		Hash:     "5typ8Zrehr5H07eDP0GNuQ",
		Modified: time.Unix(1506288628, 0),
		Raw: `
iVBORw0KGgoAAAANSUhEUgAAAMgAAAC3CAMAAABg8uG4AAABGlBMVEVHcEyoqKqytLVzc3ONj5HCwsTW
1trDysqPk5OWmJqZmcyZmZucnqGztLWQkZLb3N2Hh4eChYbd3d3d3d66u72EhYfCw8Xc3N3Ly81XV1dY
WFjh4eJPT0/h4eLo6Onr6+zt7e3k5OVra2vc3N3h4eLn5+ju7u/09PW8vb/Q0NLq6uvS09Tw8PCvsLKy
s7W/v8GkpafZ2dvy8vPIyMqsra/Fxse1t7m4ubvKy8zV1demqKrNzs/8/Pypqq3BwsSQkZPd3d+IiIvX
2NmYmpyen6DCxMX19faMjY+foaObnJ85NjeWl5n29ve6u70/PT5RT1D39/j5+fozMDFgX2CTlJdzcnSD
hIZnZmfDw8RZWFl6ent/f4BIRkiho6Xc+NdTAAAAHnRSTlMA6Y/kcFctESFABeS90Jqm3sR4v7Kq8t/Y
ynnj7s8DVns2AAAvBElEQVR4AezTwYqbQBzHcR1iotbZwCLJXnoIelBEo8yMowbdWTfSWYIsxAa2aX3/
1+gYpeyWHhJWb/0+wMCH//ykKfufOjPAQpqs+T2E4nnFhJMi7gw5xxhDVZomfVkSSsz7dbLlYEKIcch9
EcayJk3R4mv0ViMvssoK83b0u6uaomjdDZTw5RBeJD4GU0CWFnLciDGW1Jxn4yKUO+jHtK6DmSTJYcgP
LxdJjo3xh6JYzLOYqCQZ52DUXft0u3WcWoRVyQjDsH3uJFgka+M7LtksiW86yBzAh9Vq9WAq/1QAGNRF
sd0OEmroOhCS/Xd/kOSzUfexdL9lzO5yi/aWgwCSlmWZpinxdeljugJlWhBSVYOkFlFTlTS5fd77fyTG
F2Ossyhrtv9p2yzduS7BnMvXO1AUlV1pFSw+zsKkAUkQIn9LYlmWwzwIBokoazOgj8CYLxmjTfepKqdw
xEC4cvUl49SKegp5B9EArNEuSXBKeknxXkLjOA5EAyQXZdnnVz9fCwL64QkHbfkl8/ohx46QXCherA2n
ALRK3J0oqd92qJNUg8QZJB3k8fTaPPYnGUUytxhlUYPsza7mtzok9SmuU8vqJStomiYskOeKegn2SC/p
IO8lwen4mjfHJvfxPh8k+ucc9jm0T/vNpiQh5y1QwE3Da04UWVZPEZhU5IkGSUJ8D4lI5SXIQ4NEQOj5
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
AElFTkSuQmCC
`,
	})

	testLoader.Add(&Content{
		Path:     testTextFile,
		Hash:     "BqZLX-Gc5JCfp_iX562hjA",
		Modified: time.Unix(1506288652, 0),
		Compressed: `
H4sIAAAAAAAA/3yTTY/TMBCG7/yKYbmA1DQtsFClbgQSSFy4ceE4jcfNaG1PZE/SVlX/O0qyZUF8XCJ5
Yr/vk0eOeW6l0XNH0GrwtXl8EtraKKun+vMJQ+cJPklAjqacpyaQIjQtpky669UVm//MIgbaDUzHTpJC
I1Ep6u7uyFbbnaWBGyqmxYIjK6MvcoOeduu72mQ9e6r3Ys+XPTYPhyR9tEUjXlL1wq3cyr3eBkwHjtVq
26G1HA/VauskauEwsD9X0lGEjDEvWvIDKTcIkXpafLktFx8To1+Me4pMid3V8nCZkKp3q1V3ulXcUwDs
VX423Y8v/wLm3HYvyVIqElruc7WmcMXKc3xYYDVwZiV7edz9ZvN2s3FbpZMWlhpJqCyxihLp+iGQZXwZ
8DQrqt6PPK8u/zLifmWfUG92ZvDfqZ6UjXxXU866TTlfgbGkNpaH2rTrP65Cu65NV39rOYOdRsAZKCvu
PeeWLKjAnqDPZMFJAva+zzp+3EBAc1gGjmCl6QNFzUv4Lj0EPI+HQFvOz27J8enEkbWVXqFLLAkakWQ5
TspAEmB+4HiYCjtKgXNmiUtTdiOsQWgTuV2r2lVleTwel4wRl5IO5dyUy8ee+qskAo5OUpjCl8ulKbGe
ksrJSTn7Kae/5kcAAAD///DU6xJLAwAA
`,
	})
}
