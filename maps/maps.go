package maps

import "errors"


type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not found")
var ErrWordExists = errors.New("word already exists")

func (d Dictionary) Lookup(word string) (string, error) {

	meaning, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}
	return meaning, nil
	
}

func (d Dictionary) Add(word, meaning string) error {

	if _, ok := d[word]; ok {
		return ErrWordExists
	}

	d[word] = meaning
	return nil
}


func (d Dictionary) Update(word, meaning string) error {

	if _, ok := d[word]; !ok {
		return ErrWordNotFound
	}

	d[word] = meaning
	return nil
}

func (d Dictionary) Delete(word string) error {

	if _, ok := d[word]; !ok {
		return ErrWordNotFound
	}

	delete(d, word)
	return nil
}