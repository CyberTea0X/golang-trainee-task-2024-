OPENAPI=docs/openapi.yaml
DOCS=docs/openapi.html
.PHONY: docs
docs:
	npx --yes @redocly/cli build-docs -o $(DOCS) $(OPENAPI)

start-db:
	docker run --rm --name postgres -e POSTGRES_PASSWORD=test -d -p 5432:5432 postgres

build:
	go build -ldflags '-linkmode external -w -extldflags "-static"' .

docker-build:
	docker build . -t cybertea0x/gobanners:latest

docker-run: 
	docker run \
    -v ./.env:/gobanners/.env \
    --network="host" \
    --rm cybertea0x/gobanners

