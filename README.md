# My Onboaring Project Three

## Description
This is my third project on SRE onboarding to understand the Go library.

## SLO and SLI
* Availability: To be Defined
* Mean Response Time: To be Defined

## Architecture Diagram
```
+--------+        +------------------+         +---------+
|  App   | <--->  |  Ministore Lib   |  <--->  |  MySql  |
+--------+        +------------------+         +---------+
```

## Owner
[Setya Kurniawan](https://github.com/SetyaK)

## Contact and On-Call Information
[Setya Kurniawan](setya.kurniawan@bukalapak.com)

## Links

## Onboarding and Development Guide
### Prerequisite
1. Git,
  [Guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
2. Go,
  [Guide](https://golang.org/doc/install)
3. Govendor,
  `$ go get github.com/kardianos/govendor`
4. Mysql,
  [Guide](https://dev.mysql.com/doc/refman/5.7/en/installing.html)
5. Sqllite3
### Setup
1. Install required go libraries
  `$ govendor sync`
### Development Guide
#### Using supplied cmd/ministore/
1. Create .env file from .env.sample
  `$ cp .env.sample .env`
  Then modify it to match your configuration
  `ENV=development` mean existing table will be recreated when
  `database.Migration.migrate` was called
2. Build app
  Run `$ go build`
  Run `$ ministore` to test the library.
  You can also modify it to test other functionality
  Schema can be created by add call `database.Migration.migrate`

## On-Call Runbooks

## F.A.Q.
