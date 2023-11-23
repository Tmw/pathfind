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
		"#..........F.......#\n" +
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
		"#............F.....#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"####################"

	_, err := Parse(mapInput)
	if err != InvalidMapNoStart {
		t.Fatalf("expected InvalidMapNoStart error but got %v", err)
	}
}

func TestParseValdiation_NoFinish(t *testing.T) {
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
	if err != InvalidMapNoFinish {
		t.Fatalf("expected InvalidMapNoFinish error but got %v", err)
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

func TestParseValdiation_MultipleFinish(t *testing.T) {
	const mapInput = "" +
		"####################\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#.....S............#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"#.....F......F.....#\n" +
		"#..................#\n" +
		"#..................#\n" +
		"####################"

	_, err := Parse(mapInput)
	if err != InvalidMapMultipleFinish {
		t.Fatalf("expected InvalidMapMultipleFinish error but got %v", err)
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
