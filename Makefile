docs:
	npx --yes @redocly/cli build-docs -o docs/openapi.html docs/openapi.yaml

start-db:
	docker run --rm --name postgres -e POSTGRES_PASSWORD=test -d -p 5432:5432 postgres
