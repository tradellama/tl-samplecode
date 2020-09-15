This repository contains example code posted on our website at [TradeLlama](https://www.tradellama.com/blog/)

## Running all tests.

```
$git clone git@github.com:tradellama/tl-samplecode.git
$cd tl-samplecode/
$ go test -cover  ./...
```

And the result will be

```
ok      github.com/tradellama/tl-samplecode/movingaverage       0.001s  coverage: 100.0% of statements
```

## Running tests from a single package

```
$git clone git@github.com:tradellama/tl-samplecode.git
$cd tl-samplecode/
go test -cover github.com/tradellama/tl-samplecode/movingaverage
```

and the result will be

```
ok      github.com/tradellama/tl-samplecode/movingaverage               coverage: 100.0% of statements
```