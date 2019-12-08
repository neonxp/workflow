package workflow

import "testing"

func getTestWorkflow() *Workflow {
	w := NewWorkflow("initial")
	w.AddTransition("From initial to A", []Place{"initial"}, "A")
	w.AddTransition("From initial to B", []Place{"initial"}, "B")
	w.AddTransition("From A to C", []Place{"A"}, "C")
	w.AddTransition("From B,C to D", []Place{"B", "C"}, "D")
	w.AddTransition("From C,D to Finish", []Place{"C", "D"}, "Finish")
	return w
}

type testObject struct {
	place Place
}

func (t *testObject) GetPlace() Place {
	return t.place
}

func (t *testObject) SetPlace(p Place) error {
	t.place = p
	return nil
}

func TestWorkflow_Can(t *testing.T) {
	o := new(testObject)
	w := getTestWorkflow()
	if err := w.Can(o, "From initial to A"); err != nil {
		t.Error("Must has transition")
	}
	if err := w.Can(o, "From A to C"); err == nil {
		t.Error("Must has no transition")
	}
}

func TestWorkflow_GetEnabledTransitions(t *testing.T) {
	w:=getTestWorkflow()
	o := new(testObject)
	if len(w.GetEnabledTransitions(o)) != 2 {
		t.Error("Must be exactly 2 transitions from initial")
	}
}

func TestWorkflow_Apply(t *testing.T) {
	o := new(testObject)
	w := getTestWorkflow()
	if err := w.Apply(o, "From initial to A"); err != nil {
		t.Error(err)
	}
	if o.GetPlace() != "A" {
		t.Error("Must be at A place")
	}
	if err := w.Apply(o, "From B,C to D"); err != ErrCantApply {
		t.Error("Must be cant move")
	}
	if err := w.Apply(o, "From A to D"); err != ErrTransitionNotFound {
		t.Error("Must be transition not found")
	}
	if err := w.Apply(o, "From A to C"); err != nil {
		t.Error(err)
	}
	if o.GetPlace() != "C" {
		t.Error("Must be at C place")
	}
}

func TestWorkflow_DumpToDot(t *testing.T) {
	dump := getTestWorkflow().DumpToDot()
	if len(dump) != 288 {
		t.Error("Len must be 288")
	}
}
