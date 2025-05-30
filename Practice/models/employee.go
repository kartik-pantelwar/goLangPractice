// Package models contains data models for the application
package models

// Employee represents a company employee with both exported and unexported fields
type Employee struct {
	name   string // Unexported field - only accessible within this package
	EMPID  int    // Exported field - accessible from other packages
	Salary int    // Exported field - accessible from other packages
}

// GetName returns the employee's name (accessor for unexported field)
func (e Employee) GetName() string {
	return e.name
}

// SetName sets the employee's name (mutator for unexported field)
func (e *Employee) SetName(newName string) {
	e.name = newName
}

// NewEmployee creates a new Employee with provided values
func NewEmployee(name string, id int, salary int) Employee {
	return Employee{
		name:   name,
		EMPID:  id,
		Salary: salary,
	}
}
