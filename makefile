api:
	docker compose up weather-api --build

restart:
	docker compose down --volumes && api

no-cache-api:
	docker compose up weather-api --build --no-cache

down:
	docker compose down

v-down:
	docker compose down --volumes

test:
	export WEATHER_API_KEY=$(API) ENV=local && go test ./...
