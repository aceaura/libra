package route_test

import (
	"fmt"
	"testing"

	"github.com/aceaura/libra/boost/magic"
	"github.com/aceaura/libra/core/route"
)

func TestRoute(t *testing.T) {
	style := magic.ChainStyle{
		ChainSeperator: magic.SeparatorPeriod,
		WordSeparator:  magic.SeparatorUnderscore,
	}
	src := style.Chain("bus.test_route.src")
	dst := style.Chain("bus.test_route.dst")
	route := route.NewChainRoute(src, dst)
	t.Logf("%v", route)
	if fmt.Sprint(route) != "[Bus:TestRoute:Src] -> [<Bus>:TestRoute:Dst]" {
		t.Fatal("route string is not as expected")
	}
}
