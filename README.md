# Statistics for Machine Learning in Go

This repository aims to contain the code snippets of the book [Statistics for Machine Learning of Packt](https://www.packtpub.com/big-data-and-business-intelligence/statistics-machine-learning) but written in Go, rather than Python and R, in order to experiment, how Go can be used for this purpose.

## Implementation

The snippets are implemented as follow:

* All the code snippets are written in [Go example functions](https://blog.golang.org/examples)
* Each chapter is package, whose names is prefixed by `c` (of chapter) and follow with 2 digits which represent the number of the chapter.
* It uses external dependencies, as the goal is to experiment with Go package which can provide functionalities for statistics and machine learning as the ones used in the book for python an R and when the Go standard library doesn't provide them. Those dependencies are currently informed using [Go dep](https://golang.github.io/dep/).

## License

All the source code keeps the same license as the original source written in Python and R, although I haven't been able to find under which terms the sources are released, the book is under the Packt Publishing Copyright, so the contest of this repository doesn't aim nor change them and don't try to reproduce any part of the book, just only used the examples to experiment how the Go language looks like for Statistics Machine Learning calculation.

In case that there is any problem legally with this work, don't hesitate to open an issue.
