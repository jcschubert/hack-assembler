package parser

import (
	"reflect"
	"testing"
)

type ParserTestCase struct {
	desc  string
	input []string
	want  []Instruction
}

func TestParser(t *testing.T) {
	cases := []ParserTestCase{
		{
			"An empty program returns no instructions.",
			[]string{""},
			[]Instruction{},
		},
		{
			"A single @instruction with a value returns an @instruction with an address value.",
			[]string{"@1234"},
			[]Instruction{
				{"@", 1234},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := Parse(c.input)
			if !reflect.DeepEqual(c.want, got) {
				t.Fatalf("Parse(%+v) should return %+v, but returned %+v", c.input, c.want, got)
			}
		})
	}
}
