# genpasswd

A random password generator written in Go, mostly because I wanted a cross-platform password generator but also to learn how the testing and benchmarking features work in Go.

## Usage

Running the program will generate a password consisting of 12 random alphanumeric characters. For example:

```
genpasswd
```

will generate something like:

```
ppM3m4qLOwDP
```

There is an optional length parameter:

```
genpasswd 4
```

which will generate something like:

```
38qF
```

You can remove the trailing newline in scripts with `tr`:

```
genpasswd | tr -d '\n'
```

## Testing

The test ensures the correct length of password is returned for a given input, utilising a testing table. It compiles a new executable and runs it instead of mucking about with os.Args.

To run the test (while in the genpasswd directory):

```
go test
```

## Benchmarking

The benchmark is for the generation of the default 12-character password. 

To run the benchmark (while in the genpasswd directory) run:

```
go test -bench=.
```

## Profiling

To launch the profiler (while in the genpasswd directory), there is a shell script for convenience:

```
./prof.sh
```