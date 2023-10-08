SET BIN_NAME=client_presstest.exe
@REM %cd% !cd! %~dp0
SET BIN_DIR=%~dp0..\..\..\deploy\bin\
SET BUILD_DIR =%~dp0

if not exist %BIN_DIR% mkdir %BIN_DIR%

cd %BUILD_DIR%

if exist %BIN_NAME% del %BIN_NAME%

@REM windows
SET CGO_ENABLED=0 GOOS=windows GOARCH=amd64

go build -ldflags="-w -s" -o %BIN_NAME%

move /Y %BIN_NAME% %BIN_DIR%