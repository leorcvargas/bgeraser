package images

type Storage interface {
	Write(filename string, content []byte) error
}
