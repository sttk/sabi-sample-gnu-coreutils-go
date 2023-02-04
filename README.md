# sabi-sample-gnu-coreutils

This project is to develops sample implementations using sabi framework.

This is a private project to verify usability of sabi framework for command line program, and this target platform is only macOS.


### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/lib	0.131s	coverage: 87.5% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/tty	0.159s	coverage: 56.5% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/whoami	0.221s	coverage: 69.0% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/yes	0.138s	coverage: 55.6% of statements

Now using version go1.19.5
go version go1.19.5 darwin/amd64
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/lib	0.132s	coverage: 87.5% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/tty	0.168s	coverage: 56.5% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/whoami	0.232s	coverage: 69.0% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/yes	0.141s	coverage: 55.6% of statements

Now using version go1.20
go version go1.20 darwin/amd64
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/lib	0.142s	coverage: 87.5% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/tty	0.196s	coverage: 56.5% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/whoami	0.297s	coverage: 69.0% of statements
ok  	github.com/sttk-go/sabi-sample-gnu-coreutils/yes	0.400s	coverage: 55.6% of statements

Back to go1.20
Now using version go1.20
%
```


## License

Copyright (C) 2022-2023 Takayuki Sato

This program is free software under GPLv3 License.<br>
See the file LICENSE in this distribution for more details.

GNU Coreutils is the one of GNU projects by the Free Software Foundation (FSF).<br>
See https://www.gnu.org/software/coreutils/coreutils.html for more details.
