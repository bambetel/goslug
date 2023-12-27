go help testflag
go test -fuzz .
go test -fuzz . -fuzztime 5s
go test -fuzz . -fuzztime 1000x
go test -bench .
go test -bench . -benchtime 10s
