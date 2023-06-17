package utils

import "fmt"

type Log struct {
}

func (p *Log) Info(format string, a ...any) {
	fmt.Println("dig log test")
	fmt.Printf(format, a...)
}

func (p *Log) Error(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func GetLog() *Log {
	return &Log{}
}
