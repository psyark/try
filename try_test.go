package try_test

import (
	"fmt"
	"testing"

	"github.com/psyark/try"
)

func TestXxx(t *testing.T) {
	// tryなし
	{
		val, err := GetInt(123)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(val)
	}
	{
		val, err := GetString("Foo")
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(val)
	}

	// tryあり
	err := try.Try(func() error {
		fmt.Println(try.Must(GetInt(456)))
		fmt.Println(try.Must(GetString("Bar")))
		fmt.Println(try.Must(GetInt()))
		fmt.Println(try.Must(GetString()))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func GetInt(value ...int) (int, error) {
	if len(value) != 0 {
		return value[0], nil
	}
	return 0, fmt.Errorf("no value")
}

func GetString(value ...string) (string, error) {
	if len(value) != 0 {
		return value[0], nil
	}
	return "", fmt.Errorf("no value")
}
