package arena

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
	if err != ErrorInvalidArenaNoStart {
		t.Errorf("expected InvalidArenaNoStart error but got %v", err)
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
	if err != ErrorInvalidArenaNoFinish {
		t.Errorf("expected InvalidArenaNoFinish error but got %v", err)
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
	if err != ErrorInvalidArenaMultipleStart {
		t.Errorf("expected InvalidArenaMultipleStart error but got %v", err)
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
	if err != ErrorInvalidArenaMultipleFinish {
		t.Errorf("expected InvalidArenaMultipleFinish error but got %v", err)
	}
}

func compare(t *testing.T, a, b string) {
	if len(a) != len(b) {
		t.Errorf("lengths do not match. len(a) = %d; len(b) = %d\n", len(a), len(b))
	}

	for i := range a {
		if a[i] != b[i] {
			t.Errorf("strings do not match\na[%d] = '%x'; b[%d] = '%x'\n", i, a[i], i, b[i])
		}
	}
}
