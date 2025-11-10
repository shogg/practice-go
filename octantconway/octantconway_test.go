package octantconway_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/plutov/practice-go/octantconway"
)

func TestOctantConway(t *testing.T) {
	for i, test := range tests {
		conf := []byte(test.start)

		numLines := bytes.Count(conf, []byte("\n"))
		pause := min(100*time.Millisecond, 1*time.Second/time.Duration(test.N))
		if !testing.Short() {
			fmt.Println(string(conf))
		}

		for range test.N {
			conf = octantconway.OctantConway(conf)

			if !testing.Short() {
				time.Sleep(pause)
				fmt.Printf("\x1b[%dA\x1b[J", numLines+1) // clear last output
				fmt.Println(string(conf))
				numLines = bytes.Count(conf, []byte("\n"))
			}
		}

		a := trimTrailingWhitespace([]byte(test.expected))
		b := trimTrailingWhitespace(conf)
		if !bytes.Equal(a, b) {
			t.Errorf("test #%d failed\nexpected\n%sactual\n%s",
				i, a, b)
		}
	}
}

var regexpTrailingWhitespace = regexp.MustCompile(`[ ]*\n|\n*$`)

func trimTrailingWhitespace(s []byte) []byte {
	return regexpTrailingWhitespace.ReplaceAll(s, []byte("\n"))
}

var tests = []struct {
	start    string
	N        int
	expected string
}{
	{
		"𜴂𜴯",
		1,
		"𜴂𜴯",
	},
	{
		start: "" +
			"𜴩🯦",
		N: 20,
		expected: "" +
			"\n" +
			"  𜺠𜷏",
	},
	{
		start: "" +
			"▀       𜺠𜴧𜺣  ▀\n" +
			"      𜵑𜴜𜺣🮂\n" +
			"     𜺠𜴊𜶑\n" +
			"   𜺠𜴧𜺣𜴄𜺨\n" +
			"𜴳   🮂        𜴳",
		N: 144,
		expected: "" +
			"▀       𜺠𜴧𜺣  ▀\n" +
			"      𜵑𜴜𜺣🮂    \n" +
			"     𜺠𜴊𜶑      \n" +
			"   𜺠𜴧𜺣𜴄𜺨      \n" +
			"𜴳   🮂        𜴳\n",
	},
	{
		start: "" +
			"\n" +
			"   𜴣  𜴩", // diehard
		N:        130,
		expected: "",
	},
	{
		start: "" +
			"\n" +
			" 𜺠𜴧𜴧𜺣\n" +
			"  🮂▘𜴇𜴀",
		N: 43,
		expected: "" +
			"     ▀\n" +
			"\n" +
			"      🮂𜺨",
	},
}
