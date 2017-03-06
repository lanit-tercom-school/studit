###Backend Requirements:
* [golang](https://golang.org/)
* [postgresql](https://www.postgresql.org/)

###Start Service:
1. Clone repository
2. Go to `Path/To/Repository/src/service/`
3. Create DB named `studit` and then execute `/schema/data_base_init.sql` to create tables (username: `postgres`, password: `postgres`)
4. Run `InstallPackages.cmd`
5. Go to `/schema`
6. Run `go run db_init_data.go`
7. `cd ..`
8. `go run.cmd`
9. ...
10. Profit
