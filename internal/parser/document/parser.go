package document

type Parser interface {
	Parse(path string) (string, error)
}
