package detectie

import (
	"regexp"
	"strconv"
	"strings"
)

// DetectIe returns version of IE
func DetectIe(ua string) (bool, float32) {
	var err error
	var ver float64
	re := regexp.MustCompile(`MSIE (\d+.\d+)`)
	a := re.FindStringSubmatch(ua)
	if len(a) == 2 {
		ver, err = strconv.ParseFloat(a[1], 32)
		if err != nil {
			return true, 0
		}
		return true, float32(ver)
	}
	if strings.Index(ua, "Trident/") > 0 {
		re = regexp.MustCompile(`rv:(\d+.\d+)`)
		a = re.FindStringSubmatch(ua)
		if len(a) == 2 {
			ver, err = strconv.ParseFloat(a[1], 32)
			if err != nil {
				return true, 0
			}
			return true, float32(ver)
		}
	}

	re = regexp.MustCompile(`Edge/(\d+.\d+)`)
	a = re.FindStringSubmatch(ua)
	if len(a) == 2 {
		ver, err = strconv.ParseFloat(a[1], 32)
		if err != nil {
			return true, 0
		}
		return true, float32(ver)
	}

	return false, 0
}
