package workflow

// Placeer interface for objects that has place and can change place
type Placeer interface {
	GetPlace() Place
	SetPlace(Place) error
}
