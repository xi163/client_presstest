SET BIN_NAME=client_presstest.exe
SET BIN_DIR=E:\winshare\presstest\src\client_presstest\main

@REM cd E:\winshare\presstest\src\client_presstest\main
cd %BIN_DIR%

del %BIN_NAME%

@REM windows
SET CGO_ENABLED=0 GOOS=windows GOARCH=amd64

go build -ldflags="-w -s" -o %BIN_NAME%

@REM move /Y %BIN_NAME% %BIN_DIR%