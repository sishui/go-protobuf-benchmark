
@echo off
call build.bat


::==============================googleprotobuf==============================::
cd googleprotobuf
go test -bench . -benchmem
cd ..

::==============================fastpb==============================::
cd fastpb
go test -bench . -benchmem
cd ..

::==============================csproto==============================::
cd csproto
go test -bench . -benchmem
cd ..

::==============================vtprotobuf==============================::
cd vtprotobuf
go test -bench . -benchmem
cd ..


::==============================gogoprotobuf==============================::
cd gogoprotobuf
go test -bench . -benchmem
cd ..

