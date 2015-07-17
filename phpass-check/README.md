A command line tool to check a PHPass hashed password

Installation:

```bash
go get github.com/mhemmings/phpass/phpass-check
```

or download a build from http://gobuild.io/download/github.com/mhemmings/phpass/phpass-check

Usage:

```bash
phpass-check -hash="$P$passwordhash" -pw="myunhashedpassword"
```

or

```bash
phpass-check '$P$passwordhash' 'myunhashedpassword'
```
