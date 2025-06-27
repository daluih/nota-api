@echo off
echo Testando API RESTful - Nota Fiscal

echo.
echo 1. Teste /health
curl http://localhost:8080/health
echo.
echo ----------------------------------------

echo 2. Teste /notas/12345/itens (valido)
curl http://localhost:8080/notas/12345/itens
echo.
echo ----------------------------------------

echo 3. Teste /notas/12349/itens (item com quantidade 0 - erro esperado)
curl http://localhost:8080/notas/12349/itens
echo.
echo ----------------------------------------

echo 4. Teste /notas/99999/itens (nota nao existe - erro esperado)
curl http://localhost:8080/notas/99999/itens
echo.
echo ----------------------------------------

pause
