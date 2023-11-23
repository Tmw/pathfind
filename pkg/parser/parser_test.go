package parser

import (
	"strings"
	"testing"
)

func TestParseThenRender(t *testing.T) {
	const mapInput = "" +
		"####################\n" +
		"#..................#\n" +
		"#...S..............#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..........E.......#\n" +
		"#..................#\n" +
		"####################"

	m, err := Parse(mapInput)
	if err != nil {
		t.Fatal(err)
		return
	}

	var out strings.Builder
	m.Render(&out)

	actual, expected := out.String(), strings.TrimSpace(mapInput)
	compare(t, actual, expected)
}

func TestParseValdiation_NoStart(t *testing.T) {
	const mapInput = "" +
		"####################\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#............E.....#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"####################"

	_, err := Parse(mapInput)
	if err != InvalidMapNoStart {
		t.Fatalf("expected InvalidMapNoStart error but got %v", err)
	}
}

func TestParseValdiation_NoEnd(t *testing.T) {
	const mapInput = "" +
		"####################\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#............S.....#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"####################"

	_, err := Parse(mapInput)
	if err != InvalidMapNoEnd {
		t.Fatalf("expected InvalidMapNoEnd error but got %v", err)
	}
}

func TestParseValdiation_MultipleStart(t *testing.T) {
	const mapInput = "" +
		"####################\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#.....S............#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#............S.....#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"####################"

	_, err := Parse(mapInput)
	if err != InvalidMapMultipleStart {
		t.Fatalf("expected InvalidMapMultipleStart error but got %v", err)
	}
}

func TestParseValdiation_MultipleEnd(t *testing.T) {
	const mapInput = "" +
		"####################\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#.....S............#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#.....E......E.....#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"####################"

	_, err := Parse(mapInput)
	if err != InvalidMapMultipleEnd {
		t.Fatalf("expected InvalidMapMultipleEnd error but got %v", err)
	}
}

func compare(t *testing.T, a, b string) {
	if len(a) != len(b) {
		t.Fatalf("lengths do not match. len(a) = %d; len(b) = %d\n", len(a), len(b))
	}

	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("strings do not match\na[%d] = '%x'; b[%d] = '%x'\n", i, a[i], i, b[i])
		}
	}
}
