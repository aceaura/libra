package spark_test

import (
	"fmt"
	"testing"

	"github.com/aceaura/libra/repo/redis"
	"github.com/aceaura/libra/spark"
	"github.com/alicebob/miniredis/v2"
)

func TestMap(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	spark.Map([]interface{}{nil}, func(interface{}) (interface{}, error) {
		c := redis.NewClient(fmt.Sprintf("redis://%s/0", s.Addr()))
		return c.Command("SET", "test", 1)
	})
}
