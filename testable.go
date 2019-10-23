package testutils

type Field struct {
	Name  string
	Value interface{}
}

//
type Testable interface {
	Fields() []Field
}

func AssertEqualTestable(t Tester, exp, got Testable) bool {
	exps, gots := exp.Fields(), got.Fields()

	if ex, gt := len(exps), len(gots); ex != gt {
		t.Errorf("wrong count of fields: expected %d, got %s", ex, gt)
		return false
	}

	results := make(map[string]bool)

	for _, e := range exps {
		for _, g := range gots {
			if e.Name == g.Name && e.Value != g.Value {
				t.Errorf("wrong %s: expected %v, got %v", e.Name, e.Value, g.Value)
			}
		}
	}

	return len(results) == 0
}

