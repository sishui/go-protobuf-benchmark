@echo off

del /s /q *.pb.*go

::==============================googleprotobuf==============================::
where protoc-gen-go >nul
if %errorlevel% equ 1 (
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
)
protoc --go_out=googleprotobuf proto/*.proto


::==============================fastpb==============================::
where protoc-gen-fastpb >nul
if %errorlevel% equ 1 (
    go install github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
)
protoc --go_out=fastpb --fastpb_out=fastpb proto/*.proto


::==============================csproto==============================::
where protoc-gen-fastmarshal >nul
if %errorlevel% equ 1 (
    go install github.com/CrowdStrike/csproto/cmd/protoc-gen-fastmarshal@latest
)
protoc --go_out=csproto --fastmarshal_out=csproto --fastmarshal_opt=apiversion=v2 proto/*.proto


::==============================vtprotobuf==============================::
where protoc-gen-go-vtproto >nul
if %errorlevel% equ 1 (
    go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest
)
set VTPROTO_EXE=%GOPATH%\bin\protoc-gen-go-vtproto.exe
protoc -I=. -I=%GOPATH%/src/ ^
    --go_out=vtprotobuf --go-vtproto_out=vtprotobuf ^
    --plugin protoc-gen-go-vtproto="%VTPROTO_EXE%" ^
    --go-vtproto_opt=pool=y.x ^
    --go-vtproto_opt=features=pool+marshal+unmarshal+size+clone ^
    proto/*.proto


::==============================gogoprotobuf==============================::
where protoc-gen-gofast >nul
if %errorlevel% equ 1 (
    go install github.com/gogo/protobuf/protoc-gen-gofast@latest
)
set VTPROTO_EXE=%GOPATH%\bin\protoc-gen-go-vtproto.exe
protoc --gofast_out=gogoprotobuf proto/*.proto