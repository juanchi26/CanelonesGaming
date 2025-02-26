@echo off
echo Iniciando proceso de despliegue...

git add . 
git commit -m "Ultimo Commit"
git push

set GOOS=linux
set GOARCH=amd64

echo Compilando Go...
go build -o bootstrap main.go

if not exist bootstrap (
    echo Error: No se generó el archivo bootstrap. Verifica que main.go esté en la carpeta.
    exit /b 1
)

echo Creando archivo ZIP...
powershell Compress-Archive -Path bootstrap -DestinationPath bootstrap.zip -Force

if not exist bootstrap.zip (
    echo Error: No se generó el archivo bootstrap.zip. Verifica la compresión.
    exit /b 1
)

echo Despliegue exitoso. Subir bootstrap.zip a AWS Lambda manualmente o mediante AWS CLI.
exit /b 0