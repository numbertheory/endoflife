# Endoflife Command Line

I'm using this project to learn the Go programming language. 

The [Endoflife.date](https://endoflife.date) API is the source for all the information, and this tool would require a network connection to work.

## Installation

```
$ go build
```

If you want to use `go install`, first run this command to find your go environment's path:

```
$ go env GOPATH
```

Then, add that path with "/bin" attached to it to your local PATH. This will allow you to run `go install` on any local go project and have the binary installed to an accessible PATH. 


## Running

Run the command with just a product to show the End of Life dates for all product cycles:

```
$ endoflife -p ruby
|    EOL     |   Latest Version  |
+============+===================+
| 2026-03-31 | 3.2.1             |
| 2025-03-31 | 3.1.3             |
| 2024-03-31 | 3.0.5             |
| 2023-03-31 | 2.7.7             |
| 2022-03-31 | 2.6.10            |
| 2021-03-31 | 2.5.9             |
| 2020-03-31 | 2.4.10            |
| 2019-03-31 | 2.3.8             |
| 2018-03-31 | 2.2.10            |
| 2017-03-31 | 2.1.10            |
| 2016-02-24 | 2.0.0p648         |
| 2015-02-23 | 1.9.3p551         |
+============+===================+
```

You can also run the command targeting a specific product cycle, and it will return the latest version of that product cycle:

```
$ endoflife -p ruby -c 3.2
|    EOL     |   Latest Version  |
+============+===================+
| 2026-03-31 | 3.2.1             |
+============+===================+
```