package users

import "errors"

type user struct {
	Name string
	Age  uint8
}

func (u user) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Age == 0 {
		return errors.New("age is required")
	}
	return nil
}

func New(name string, age uint8) (*user, error) {
	u := &user{
		Name: name,
		Age:  age,
	}

	err := u.Validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func addYear(u user) {
	u.Age++
}
