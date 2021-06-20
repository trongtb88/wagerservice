## How to run this service at docker
1. Run DOCKER DEAMON at your machine successfully
2. Make sure don't have any image mysql is running at port 3306, otherwise you will have 1 error
3. Make sure don't have any image is running at port 8080, otherwise you will have 1 error
2. Git clone this repo
3. Change file .env which is mapped with your config (DB_USER, DB_PASSWORD, DB_HOST), please note uncommment DB_HOST at docker, comment using 127.0.0.1
5.  Go to terminal at root of project
```sh
   chmod 755 start.sh
   ./start.sh
```

6. If have some logs at console like, server started and worked successfully

```sh
wager_app      | 
wager_app      | 2021/06/20 02:26:07 /go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:194
wager_app      | [0.215ms] [rows:-] SELECT DATABASE()
wager_app      | 
wager_app      | 2021/06/20 02:26:07 /go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:203
wager_app      | [0.961ms] [rows:-] SELECT column_name, is_nullable, data_type, character_maximum_length, numeric_precision, numeric_scale , datetime_precision FROM information_schema.columns WHERE table_schema = 'wager_db' AND table_name = 'buy_wagers'
wager_app      | 
wager_app      | 2021/06/20 02:26:07 /go/pkg/mod/gorm.io/driver/mysql@v1.0.5/migrator.go:83
wager_app      | [1.749ms] [rows:0] ALTER TABLE `buy_wagers` MODIFY COLUMN `wager_id` bigint(20)
wager_app      | 2021/06/20 02:26:07 Starting server at port:  8080
wager_app      | 
wager_app      | 2021/06/20 02:26:19 /app/src/business/domain/wager/wager_sql.go:78
wager_app      | [0.489ms] [rows:1] SELECT * FROM `wagers` LIMIT 20

```

7. Go to Postman, import like that

```
curl --location --request POST 'http://localhost:8080/api/v1/wagers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "total_wager_value": 5000,
    "odds": 4000,
    "selling_percentage": 34,
    "selling_price": 34454.67
}'
```


## How to run this service at localhost
1. Start your mysql at your localhost machine successfully
2. Git clone this repo
3. Change file .env which is mapped with your config (DB_USER, DB_PASSWORD, DB_HOST), please note commment DB_HOST at docker, uncomment using 127.0.0.1
4. Create database wager_db by yourself
5.  Go to terminal at root of project
```sh
   go get .    
   go run src/cmd/main.go
```