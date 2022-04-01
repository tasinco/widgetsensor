package service

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Service struct {
}

func (s *Service) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		lineBytes, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		}
		line := string(lineBytes[:len(lineBytes)-1])
		log.Println(line)
	}
}
