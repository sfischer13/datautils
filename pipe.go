package datautils

import "bufio"
import "fmt"
import "os"

// StdinSource TODO
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

// StdoutSink TODO
func StdoutSink(src <-chan string) {
	for s := range src {
		fmt.Println(s)
	}
}

// TransformPipe TODO
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
