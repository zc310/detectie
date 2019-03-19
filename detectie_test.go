package detectie_test

import (
	"github.com/zc310/detectie"
	"testing"
)

var tests = []struct {
	Agent   string
	IE      bool
	Version float32
}{
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36", false, 0},
	{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Safari/537.36 Edge/13.10586", true, float32(13.10586)},
	{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0", true, float32(12.0)},
	{"Mozilla/4.0 (compatible; MSIE 6.0b; Windows NT 5.0; YComp 5.0.0.0) (Compatible;  ;  ; Trident/4.0)", true, float32(6.0)},
	{"Mozilla/5.0 (Windows; U; MSIE 9.0; WIndows NT 9.0; en-US))", true, float32(9.0)},
	{"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; AS; rv:11.0) like Gecko", true, float32(11.0)},
	{"Mozilla/5.0 (compatible; MSIE 10.6; Windows NT 6.1; Trident/5.0; InfoPath.2; SLCC1; .NET CLR 3.0.4506.2152; .NET CLR 3.5.30729; .NET CLR 2.0.50727) 3gpp-gba UNTRUSTED/1.0", true, float32(10.6)},
	{"Mozilla/4.0 (compatible; U; MSIE 6.0; Windows NT 5.1) (Compatible;  ;  ; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)", true, float32(6.0)},
}

func TestDetectIe(t *testing.T) {
	var ie bool
	var ver float32

	for _, test := range tests {
		ie, ver = detectie.DetectIe(test.Agent)
		if ie != test.IE || ver != test.Version {
			t.Errorf("DetectIe(%q) = %t %.2f, want %t %.2f", test.Agent, ie, ver, test.IE, test.Version)
		}

	}
}
