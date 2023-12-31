.PHONY: build-proto

build-proto:
	protoc --proto_path=pkg/proto --go_out=pkg/build/ --go-grpc_out=pkg/build/ pkg/proto/*.proto

build-proto-python:
	python -m grpc_tools.protoc --proto_path=pkg/proto --python_out=src --grpc_python_out=src pkg/proto/*.proto

build-proto-python:
	python -m grpc_tools.protoc --proto_path=pkg/proto/*.proto --python_out=pkg/build --grpc_python_out=pkg/build

run-app:
	go run app.go

run-postgres:
	docker run --rm -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pswd postgres:latest -c ssl=on -c ssl_cert_file=/etc/ssl/certs/ssl-cert-snakeoil.pem -c ssl_key_file=/etc/ssl/private/ssl-cert-snakeoil.key

# Migrations:

migrate-users-up:
	migrate -path ${CURDIR}/Users/pkg/migrations/ -database postgres://postgres:pswd@localhost:5432/postgres up 1

migrate-users-down:
	migrate -path ${CURDIR}/Users/pkg/migrations/ -database postgres://postgres:pswd@localhost:5432/postgres down 1

migrate-users-fix:
	migrate -path ${CURDIR}/Users/pkg/migrations/ -database postgres://postgres:pswd@localhost:5432/postgres force 1

migrate-users-clean: migrate-users-fix migrate-users-down

migrate-products-up:
	migrate -path ${CURDIR}/Products/pkg/migrations/ -database postgres://postgres:pswd@localhost:5432/postgres up 1

migrate-products-down:
	migrate -path ${CURDIR}/Products/pkg/migrations/ -database postgres://postgres:pswd@localhost:5432/postgres down 1

migrate-products-fix:
	migrate -path ${CURDIR}/Products/pkg/migrations/ -database postgres://postgres:pswd@localhost:5432/postgres force 1

migrate-products-clean: migrate-products-fix migrate-products-down

#	docker run -v {{ migration dir }}:/migrations --network host migrate/migrate
#   	-path=/migrations/ -database postgres://localhost:5432/database up 2
