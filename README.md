# golang-learn
Learning golang by following a tutorial and practice in local

Need:
1. some code editor like VS CODE
2. Git
3. Go itself https://go.dev/doc/install

Useful Git commands
git clone https://github.com/kapilepatel/golang-learn.git
git config --global user.name "John Doe"
$ git config --global user.email johndoe@example.com

go mod init golang-learn

go build main.go

./main

go tool dist list
go env
go env GOOS GOARCH

cd 31-interface
go test


cd 42-try-test   
go test -v
go test -bench ValidDate


go test -coverprofile c.out
go tool cover -html=c.out