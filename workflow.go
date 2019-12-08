package workflow

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	// ErrCantApply error if transition is not applicable to object
	ErrCantApply = errors.New("cant apply transition")
	// ErrTransitionNotFound error if no transition with this name
	ErrTransitionNotFound = errors.New("transition not found")
)

// Workflow state machine
type Workflow struct {
	transitions  map[string]transition
	initialPlace Place
}

// NewWorkflow returns new Workflow instance
func NewWorkflow(initialPlace Place) *Workflow {
	return &Workflow{initialPlace: initialPlace, transitions: map[string]transition{}}
}

// Can returns nil if transition applicable to object and error if not
func (w *Workflow) Can(obj Placeer, transition string) error {
	currentPlace := obj.GetPlace()
	if currentPlace == "" {
		currentPlace = w.initialPlace
	}
	tr, ok := w.transitions[transition]
	if !ok {
		return ErrTransitionNotFound
	}
	for _, f := range tr.From {
		if f == currentPlace {
			return nil
		}
	}
	return ErrCantApply
}

// GetEnabledTransitions return all applicable transitions for object
func (w *Workflow) GetEnabledTransitions(obj Placeer) []string {
	currentPlace := obj.GetPlace()
	if currentPlace == "" {
		currentPlace = w.initialPlace
	}
	var result = make([]string, 0)
	for name, t := range w.transitions {
		for _, f := range t.From {
			if f == currentPlace {
				result = append(result, name)
				break
			}
		}
	}
	return result
}

// Apply next state from transition to object
func (w *Workflow) Apply(obj Placeer, transition string) error {
	currentPlace := obj.GetPlace()
	if currentPlace == "" {
		currentPlace = w.initialPlace
	}
	tr, ok := w.transitions[transition]
	if !ok {
		return ErrTransitionNotFound
	}
	for _, f := range tr.From {
		if f == currentPlace {
			return obj.SetPlace(tr.To)
		}
	}
	return ErrCantApply
}

// AddTransition to workflow
func (w *Workflow) AddTransition(name string, from []Place, to Place) {
	w.transitions[name] = transition{
		From: from,
		To:   to,
	}
}

// DumpToDot dumps transitions to Graphviz Dot format
func (w *Workflow) DumpToDot() []byte {
	buf := bytes.NewBufferString(fmt.Sprintf("digraph {\n%s[color=\"blue\"]\n", w.initialPlace))
	for name, t := range w.transitions {
		for _, f := range t.From {
			_, _ = buf.WriteString(fmt.Sprintf("%s -> %s[label=\"%s\"];\n", f, t.To, name))
		}
	}
	buf.WriteString("}")
	return buf.Bytes()
}

// Place is one of state
type Place string

type transition struct {
	From []Place
	To   Place
}
