@echo off
setlocal

:: **1. Menampilkan versi dan informasi**
set SCRIPT_VERSION=1.0.0
set DEFAULT_PACKAGE_FRAMEWORK=github.com/yon-module/yon-framework

echo =======================================
echo  Golang Project Setup Script v%SCRIPT_VERSION%
echo =======================================

:: **2. Input Project Name**
set /p PROJECT_NAME=Enter Project Name: 

:: Ubah nama ke kebab-case (lowercase & ganti spasi dengan "-")
for /f "delims=" %%A in ('powershell -command "[System.Text.RegularExpressions.Regex]::Replace('%PROJECT_NAME%', '\s+', '-').ToLower()"') do set "FOLDER_NAME=%%A"
echo Project folder name: %FOLDER_NAME%

:: **3. Input Package Name (default = folder name)**
set /p PACKAGE_NAME=Enter Package Name (default: %FOLDER_NAME%): 
if "%PACKAGE_NAME%"=="" set PACKAGE_NAME=%FOLDER_NAME%

:: **4. Run install Go packages**
echo Installing required Go packages...
cd %FOLDER_NAME%
go mod init %PACKAGE_NAME%
go get -u %DEFAULT_PACKAGE_FRAMEWORK%
cd ..

:: **5. Generate Folder Structure**
echo Creating project structure...
mkdir %FOLDER_NAME%\src\main\controllers
mkdir %FOLDER_NAME%\src\main\model\dto\request
mkdir %FOLDER_NAME%\src\main\model\dto\response
mkdir %FOLDER_NAME%\src\main\model\entity
mkdir %FOLDER_NAME%\src\main\repository
mkdir %FOLDER_NAME%\src\main\service
mkdir %FOLDER_NAME%\src\main\util
mkdir %FOLDER_NAME%\src\main\helper
mkdir %FOLDER_NAME%\src\tests
type nul > %FOLDER_NAME%\main.go

:: **6. Create .env.sample**
echo Creating .env.sample file...
(
echo yon.server.appName=%PACKAGE_NAME%
echo yon.server.port=8080
echo.
echo yon.database.type=
echo yon.database.host=
echo yon.database.user=
echo yon.database.password=
echo yon.database.port=
echo yon.database.db=
) > %FOLDER_NAME%\.env.sample

:: Selesai
echo Project setup complete! 🎉
echo Navigate to your project folder: cd %FOLDER_NAME%

endlocal
