package workflow

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	// ErrTransitionNotFound error if no transition with this name
	ErrTransitionNotFound = errors.New("transition not found")
)

// Workflow state machine
type Workflow struct {
	transitions  map[Place][]Place
	initialPlace Place
}

// NewWorkflow returns new Workflow instance
func NewWorkflow(initialPlace Place) *Workflow {
	return &Workflow{initialPlace: initialPlace, transitions: map[Place][]Place{}}
}

// Can returns nil if transition applicable to object and error if not
func (w *Workflow) Can(obj Placeer, to Place) error {
	currentPlace := obj.GetPlace()
	if currentPlace == "" {
		currentPlace = w.initialPlace
	}
	tr, ok := w.transitions[currentPlace]
	if !ok {
		return ErrTransitionNotFound
	}
	for _, f := range tr {
		if f == to {
			return nil
		}
	}
	return ErrTransitionNotFound
}

// GetEnabledTransitions return all applicable transitions for object
func (w *Workflow) GetEnabledTransitions(obj Placeer) []Place {
	currentPlace := obj.GetPlace()
	if currentPlace == "" {
		currentPlace = w.initialPlace
	}
	if _, ok := w.transitions[currentPlace]; !ok {
		return nil
	}
	return w.transitions[currentPlace]
}

// Apply next state from transition to object
func (w *Workflow) Apply(obj Placeer, to Place) error {
	currentPlace := obj.GetPlace()
	if currentPlace == "" {
		currentPlace = w.initialPlace
	}
	tr, ok := w.transitions[currentPlace]
	if !ok {
		return ErrTransitionNotFound
	}
	for _, f := range tr {
		if f == to {
			return obj.SetPlace(to)
		}
	}
	return ErrTransitionNotFound
}

// AddTransition to workflow
func (w *Workflow) AddTransition(from Place, to Place) {
	if _, ok := w.transitions[from]; !ok {
		w.transitions[from] = []Place{}
	}
	w.transitions[from] = append(w.transitions[from], to)
}

// DumpToDot dumps transitions to Graphviz Dot format
func (w *Workflow) DumpToDot() []byte {
	buf := bytes.NewBufferString(fmt.Sprintf("digraph {\n%s[color=\"blue\"]\n", w.initialPlace))
	for from, to := range w.transitions {
		for _, place := range to {
			_, _ = buf.WriteString(fmt.Sprintf("%s -> %s[label=\"%s\"];\n", from, place, fmt.Sprintf("%s → %s", from, place)))
		}
	}
	buf.WriteString("}")
	return buf.Bytes()
}

// Place is one of state
type Place string
