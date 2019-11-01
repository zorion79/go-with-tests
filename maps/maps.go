package maps

const (
	ErrWordExist        = DictionaryErr("word exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
	ErrNotFound         = DictionaryErr("could not find the word")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExist
	default:
		return err
	}
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}
