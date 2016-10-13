# go-201
Some Go programming exercises.

* q1: A recursive function that outputs a slice of integers as an indented string of XML.
* q2: A function that converts a one-lined XML string into an indented XML string.
* q3: Refer this [sqlfiddle](http://sqlfiddle.com/#!9/c14d2/4) for executable example.
* q4: An application that reads a CSV file.
* common_string: A function that finds the common strings in two given strings. Refer [common_strings.md](common_strings/README.md) for examples.
* common_subsequence: A function that finds the [longest common subsequence](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem) in two given strings. Refer [common_subsequence.md](common_subsequence/README.md) for examples.
* essay_monkey: A function that builds an essay from a collection of adjectives, nouns and verbs. Refer [essay_monkey.md](essay_monkey/README.md) for more information.
* tic_tae_toe: A program that plays tic-tae-toe.

To run the test:
```sh
$ go test -v -cover ./...
```

To run the q4 `main` application:
```sh
$ go build -v
$ ./go-201 q4.data
```

To run the tic-tae-toe application:
```sh
$ cd tic-tae-toe
$ go build
$ ./tic-tae-toe
```
