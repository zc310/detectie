package detectie

import (
	"regexp"
	"strconv"
	"strings"
)

// DetectIe returns version of IE
func DetectIe(ua string) (bool, float32) {
	var (
		err error
		ver float64
		re  *regexp.Regexp
		a   []string
		n   int
	)
	if n = strings.Index(ua, "MSIE "); n > 0 {
		re = regexp.MustCompile(`MSIE (\d+.\d+)`)
		if a = re.FindStringSubmatch(ua); len(a) == 2 {
			if ver, err = strconv.ParseFloat(a[1], 32); err != nil {
				return true, 0
			}
			return true, float32(ver)
		}
	}
	if n = strings.Index(ua, "Trident/"); n > 0 {
		re = regexp.MustCompile(`rv:(\d+.\d+)`)
		if a = re.FindStringSubmatch(ua); len(a) == 2 {
			if ver, err = strconv.ParseFloat(a[1], 32); err != nil {
				return true, 0
			}
			return true, float32(ver)
		}
	}
	if n = strings.Index(ua, "Edge/"); n > 0 {
		re = regexp.MustCompile(`Edge/(\d+.\d+)`)
		if a = re.FindStringSubmatch(ua); len(a) == 2 {
			if ver, err = strconv.ParseFloat(a[1], 32); err != nil {
				return true, 0
			}
			return true, float32(ver)
		}
	}

	return false, 0
}
