# msw golang exploration

As I did with Python many moons ago, I'm starting my golang initiation with
a wordlist generator to exercise the libraries.

I'm still trying to understand how packages fit together and I'm hoping
that this can contain a couple of projects.

# reading list

**practical-go** https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html

A dense and rich technical blog with many concepts presented concisely with
examples.

**the go blog** https://go.dev/blog

I've barely scratched the surface of this

**Learn Go with Tests** https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces

A test-driven dev elopment guide that describes many features of the go environment that
are essential but not obvious through the main documentation. For example
`go test -bench` benchmarking is presented quite early and `go mod init` is in the
first section.


# thoughts while learning Go

I feel like I'd been seduced by the Syntactic Sugar of Python. I have long thought
that there is a geek self-impression factor that C++ and Java seem to generate:
"Look how clever this piece of code shows me (and others) that I am" provides
an endogenous dopamine hit. I've long thought those languages to be the exemplars
of "cleverness junkie", but in the face of Go, I'm beginning to think I'd been
in thrall of Python for similar reasons.

Yes, Python is more readable on its face than most other languages, but once
you start throwing in comprehensions and abstract base classes (Collections, Iterators,
Iterables, Sequences, et. seq.) type annotations, f-strings, dataclasses to fudge
necessary dunder-methods, etc., etc. Python has gotten pretty damn complex. Somewhere
around my 10-year mark in 2020 as a Pythonista, I found myself wondering why
there were so many new releases and wishing that they'd slow down. This was not
because I was having trouble keeping up, but because fundamental features (for example,
structural pattern matching) are continually added that I've yet to find a use case for.

The simplicity and the durability of the Go language is no accident, and given
its intellectual heritage I'm not surprised.

The advisory type annotations have really been a boon to writing correct code in
Python. Where it really helps is in making use of diverse libraries and optional
parameter lists that can number twenty or more. The thing is, my maxim prior to
annotations was "if method tank can't take an object cap, why did you pass it one?";
