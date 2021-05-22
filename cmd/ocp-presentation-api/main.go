package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("It is an empty service.\n")

	for i := 1; i <= 10; i++ {
		filename := strconv.Itoa(i)

		doSomething := func(filename string) (err error) {
			f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return
			}

			defer func() {
				cerr := f.Close()
				if err == nil {
					err = cerr
				}
			}()

			if _, err = f.WriteString("Bar"); err != nil {
				return err
			}

			return nil
		}

		if err := doSomething(filename); err != nil {
			log.Fatal(err)
		}

		if err := os.Remove(filename); err != nil {
			log.Fatal(err)
		}
	}
}
