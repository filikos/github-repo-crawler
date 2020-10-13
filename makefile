all : install stop
.PHONY : all

install :
	#runns all migration script. To run specific migration version build more targets.
	@echo "Get latest https://github.com/golang-migrate/migrate image"
	docker pull migrate/migrate
	
	@echo "Starting postgreSQL database"
	docker run -d \
	--name dev \
	--env-file $(shell pwd)/config/dbConfig.env \
	-p 5432:5432 \
	-v $(shell pwd)/postgres:/var/lib/postgresql/data postgres
	
	@echo "Wait for database"
	sleep 5

	@echo "Sun migration scripts"
	docker run \
	-v $(shell pwd)/migrations:/migrations \
	--network host migrate/migrate \
	-path=/migrations/ \
	-database "postgres://postgres:postgresPW@localhost:5432/?sslmode=disable" up
	# for specific migration: ....432/?sslmode=disable" up {{add integer migration number}}

	@echo "Stop and remove PostgreSQL container"
	docker stop `docker ps -aqf "name=dev"`
	docker rm `docker ps -aqf "name=dev"`