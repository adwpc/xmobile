
#git reset --hard
go generate -run gen.go
go install ./cmd/gobind
go install ./cmd/gomobile
