SET BIN_NAME=client_presstest.exe
@REM %cd% !cd! %~dp0
SET BIN_DIR=%~dp0..\deploy\bin\

cd %BIN_DIR%

.\%BIN_NAME% --dir-level=1 --conf-name=..\clientConfig_win\conf.ini