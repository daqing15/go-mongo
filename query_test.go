package mongo

import (
	"testing"
)

var countTests = []struct {
	query interface{}
	limit int
	skip  int
	count int64
}{
	{limit: 100, count: 10},
	{limit: 5, count: 5},
	{skip: 5, count: 5},
	{skip: 100, count: 0},
	{query: D{{"x", 1}}, count: 1},
}

func TestCount(t *testing.T) {
	c := dialAndDrop(t, "go-mongo-test", "test")
	defer c.Conn.Close()

	for i := 0; i < 10; i++ {
		err := c.Insert(map[string]int{"x": i})
		if err != nil {
			t.Fatal("insert", err)
		}
	}

	for _, tt := range countTests {
		n, err := c.Find(tt.query).Limit(tt.limit).Skip(tt.skip).Count()
		if err != nil {
			t.Fatal("count", err)
		}
		if n != tt.count {
			t.Errorf("test: %+v, actual: %d", tt, n)
		}
	}
}

func TestQuery(t *testing.T) {
	c := dialAndDrop(t, "go-mongo-test", "test")
	defer c.Conn.Close()

	for i := 0; i < 10; i++ {
		err := c.Insert(map[string]int{"x": i})
		if err != nil {
			t.Fatal("insert", err)
		}
	}

	var m M
	err := c.Find(nil).Sort(D{{"x", -1}}).One(&m)
	if err != nil {
		t.Fatal("findone", err)
	}

	if m["x"] != 9 {
		t.Fatal("expect max value for descending sort")
	}
}