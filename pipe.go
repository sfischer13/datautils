package datautils

import "bufio"
import "fmt"
import "os"

func StdinSource() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			out <- line
		}
	}()
	return out
}

func StdoutSink(src <-chan string) {
	for s := range src {
		fmt.Println(s)
	}
}

func TransformPipe(src <-chan string, transform func(string) string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range src {
			out <- transform(s)
		}
	}()
	return out
}
