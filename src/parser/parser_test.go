package parser

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("an empty program returns no instructions", func(t *testing.T) {
		input := []string{""}
		want := []Instruction{}
		got := Parse(input)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("Parse(%+v) should return %+v, but returned %+v", input, want, got)
		}
	})
	t.Run("A single @instruction returns an @instruction", func(t *testing.T) {
		input := []string{"@1234"}
		want := []Instruction{
			{"@", 1234},
		}
		got := Parse(input)

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("Parse(%+v) should return %+v, but returned %+v", input, want, got)
		}
	})
}
