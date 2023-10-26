# [sabi-sample-gnu-coreutils-go][repo-url] [![CI Status][ci-img]][ci-url] [![GPL3 License][gpl3-img]][gpl3-url]

This project is to develop sample implementations using [sabi][sabi-url] framework, of which behaviors are same with programs in [GNU coreutils][gnu-url].

This is a private project to verify usability of [sabi][sabi-url] framework for command line programs, and this target platform is only macOS.


## Supporting Go versions

Programs in this project support Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/lib	0.169s	coverage: 91.7% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/tty	0.235s	coverage: 88.2% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/whoami	0.374s	coverage: 96.9% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/yes	0.201s	coverage: 95.1% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/lib	0.112s	coverage: 91.7% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/tty	0.208s	coverage: 88.2% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/whoami	0.276s	coverage: 96.9% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/yes	0.157s	coverage: 95.1% of statements

Now using version go1.20.8
go version go1.20.8 darwin/amd64
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/lib	0.142s	coverage: 91.7% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/tty	0.177s	coverage: 88.2% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/whoami	0.284s	coverage: 96.9% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/yes	0.384s	coverage: 95.1% of statements

Now using version go1.21.1
go version go1.21.1 darwin/amd64
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/lib	0.128s	coverage: 91.7% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/tty	0.161s	coverage: 88.2% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/whoami	0.368s	coverage: 96.9% of statements
ok  	github.com/sttk/sabi-sample-gnu-coreutils-go/yes	0.264s	coverage: 95.1% of statements

Back to go1.21.1
Now using version go1.21.1
```


## License

Copyright (C) 2022-2023 Takayuki Sato

This program is free software under GPLv3 License.<br>
See the file LICENSE in this distribution for more details.

GNU Coreutils is the one of GNU projects by the Free Software Foundation (FSF).<br>
See https://www.gnu.org/software/coreutils/coreutils.html for more details.


[repo-url]: https://github.com/sttk/sabi-sample-gnu-coreutils-go
[ci-img]: https://github.com/sttk/sabi-sample-gnu-coreutils-go/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/sabi-sample-gnu-coreutils-go/actions
[gpl3-img]: https://img.shields.io/badge/license-GPL3-green.svg
[gpl3-url]: https://opensource.org/license/gpl-3-0/
[sabi-url]: https://github.com/sttk/sabi
[gnu-url]: https://www.gnu.org/software/coreutils/

