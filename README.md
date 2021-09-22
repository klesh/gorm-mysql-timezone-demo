# Time zone handling for gorm/mysql

## How to run

```sh
git clone https://github.com/klesh/gorm-mysql-timezone-demo
cd devlake-tz-demo
go get
go run main.go
```

## Sample output

Use local timezone (UTC+8)
```
$ go run main.go

origin values
 local now: 2021-09-22 14:47:58.898758837 +0800 CST m=+0.006028471
 tz0: 2021-01-19 00:00:00 +0000 +0000
 tz1: 2021-01-19 01:00:00 +0100 +0100

save to database
+-----------+----------------------------+
| timezone  | the_datetime               |
+-----------+----------------------------+
| local now | 2021-09-22 14:47:58.899000 |
| tz0       | 2021-01-19 08:00:00        |
| tz1       | 2021-01-19 08:00:00        |
+-----------+----------------------------+

load from database
 local now: 2021-09-22 14:47:58.899 +0800 CST
 tz0: 2021-01-19 08:00:00 +0800 CST
 tz1: 2021-01-19 08:00:00 +0800 CST
```

Use a specific timezone
```
$ TZ=Canada/Atlantic go run main.go                                                                                                          0.417s
origin values
 local now: 2021-09-22 03:49:34.12075714 -0300 ADT m=+0.009208072
 tz0: 2021-01-19 00:00:00 +0000 +0000
 tz1: 2021-01-19 01:00:00 +0100 +0100

save to database
+-----------+----------------------------+
| timezone  | the_datetime               |
+-----------+----------------------------+
| local now | 2021-09-22 03:49:34.121000 |
| tz0       | 2021-01-18 20:00:00        |
| tz1       | 2021-01-18 20:00:00        |
+-----------+----------------------------+


load from database
 local now: 2021-09-22 03:49:34.121 -0300 ADT
 tz0: 2021-01-18 20:00:00 -0400 AST
 tz1: 2021-01-18 20:00:00 -0400 AST
```

## Conclusion

- golang `time.Time` contains time zone information
- `time.Time` was converted to the TimeZone specified in connection string and stored without timezone information
- all datetime values would break if `loc` setting didn't match
