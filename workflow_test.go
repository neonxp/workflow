package workflow

import (
	"testing"
)

func getTestWorkflow() *Workflow {
	w := NewWorkflow("Start")
	w.AddTransition("Start", "A")
	w.AddTransition("Start", "B")
	w.AddTransition("A", "C")
	w.AddTransition("B", "D")
	w.AddTransition("C", "D")
	w.AddTransition("C", "Finish")
	w.AddTransition("D", "Finish")
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
	if err := w.Can(o, "A"); err != nil {
		t.Error("Must has transition")
	}
	if err := w.Can(o, "C"); err == nil {
		t.Error("Must has no transition")
	}
}

func TestWorkflow_GetEnabledTransitions(t *testing.T) {
	w := getTestWorkflow()
	o := new(testObject)
	if len(w.GetEnabledTransitions(o)) != 2 {
		t.Error("Must be exactly 2 transitions from initial")
	}
}

func TestWorkflow_Apply(t *testing.T) {
	o := new(testObject)
	w := getTestWorkflow()
	if err := w.Apply(o, "A"); err != nil {
		t.Error(err)
	}
	if o.GetPlace() != "A" {
		t.Error("Must be at A place")
	}
	if err := w.Apply(o, "Finish"); err != ErrTransitionNotFound {
		t.Error("Must be transition not found")
	}
	if err := w.Apply(o, "C"); err != nil {
		t.Error(err)
	}
	if o.GetPlace() != "C" {
		t.Error("Must be at C place")
	}
}

func TestWorkflow_DumpToDot(t *testing.T) {
	dump := getTestWorkflow().DumpToDot()
	if len(dump) != 242 {
		t.Errorf("Len must be 242, got %d", len(dump))
	}
}
