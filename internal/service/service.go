package service

import (
	"bufio"
	"io"
	"log"
	"os"
	"widgetsensor/internal/sensorcollector"
)

type Service struct {
}

func (s *Service) Run() {
	sensorCollector := sensorcollector.NewSensor()

	reader := bufio.NewReader(os.Stdin)
	for linecnt := uint64(0); ; linecnt++ {
		lineBytes, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		line := string(lineBytes[:len(lineBytes)-1])
		log.Println(line)
		sensorCollector.Consume(line)
		if linecnt > 0 && !sensorCollector.ReferenceValid() {
			log.Fatal("reference line missing")
		}
	}
}
