package maps

type Dictionary map[string]string

type DictionaryError string

func (d DictionaryError) Error() string {
	return string(d)
}

const (
	ErrorWordNotFound     = DictionaryError("could not find the word you were looking for")
	ErrorWordExists       = DictionaryError("cannot add word because it already exists")
	ErrorWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, searchErr := d.Search(word)
	switch searchErr {
	case ErrorWordNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrorWordExists
	default:
		return searchErr
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, searchErr := d.Search(word)
	switch searchErr {
	case ErrorWordNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return searchErr
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
