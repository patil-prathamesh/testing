package maps

type Dictionary map[string]string

const ErrNotFound = DictionaryErr("could not find the word you were looking for")
const ErrWordExists = DictionaryErr("cannot add word because it already exists")
const ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	_, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return d[word], nil
}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
