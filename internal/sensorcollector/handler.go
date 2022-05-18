package sensorcollector

type sensorMonitor interface {
	accept(*reference, string) error
}

type handler interface {
	argLen() int
	consume([]string) (string, sensorMonitor, error)
}
