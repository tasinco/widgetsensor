package sensorcollector

type handler interface {
	argLen() int
	consume([]string) error
}
