package banner_test

import (
	"fmt"
	"testing"

	banner "github.com/waldirborbajr/bpbanner"
)

func TestBanner(t *testing.T) {
	fmt.Println(banner.Banner("0.0.1"))
}
