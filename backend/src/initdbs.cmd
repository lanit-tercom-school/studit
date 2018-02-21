set postgres64=C:\Program Files\PostgreSQL\9.6\bin
set postgres32=C:\Program Files (x86)\PostgreSQL\9.6\bin

IF EXIST "%postgres64%" set path=%path%;%postgres64%
IF EXIST "%postgres32%" set path=%path%;%postgres32%

cd auth-service
start initdb.cmd
cd ../file-service
start initdb.cmd
cd ../data-service
start initdb.cmd