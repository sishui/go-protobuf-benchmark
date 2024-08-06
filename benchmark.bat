
@echo off
call build.bat


::==============================googleprotobuf==============================::
cd googleprotobuf
go get -u google.golang.org/protobuf
go mod tidy
go test -bench . -benchmem
cd ..

::==============================fastpb==============================::
cd fastpb
go get -u github.com/cloudwego/fastpb
go mod tidy
go test -bench . -benchmem
cd ..

::==============================csproto==============================::
cd csproto
go get -u github.com/CrowdStrike/csproto
go mod tidy
go test -bench . -benchmem
cd ..

::==============================vtprotobuf==============================::
cd vtprotobuf
go get -u google.golang.org/protobuf
go mod tidy
go test -bench . -benchmem
cd ..


::==============================gogoprotobuf==============================::
cd gogoprotobuf
go get -u github.com/gogo/protobuf
go mod tidy
go test -bench . -benchmem
cd ..

