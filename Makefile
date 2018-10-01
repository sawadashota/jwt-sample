bootstrap: copy-envrc ensure generate-certs

copy-envrc:
	cp .envrc.sample .envrc
	echo "edit .envrc for your env"

ensure:
	docker-compose run --rm go dep ensure -v

generate-certs:
	ssh-keygen -t rsa -f certs/id_rsa -P ""
	ssh-keygen -f certs/id_rsa.pub -e -m pkcs8 > certs/id_rsa.pub.pkcs8

test:
	docker-compose run --rm go go test -v ./...

coverage:
	docker-compose run --rm go go test -coverprofile=profile ./... && go tool cover -html=profile -o profile.html
