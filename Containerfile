from docker.io/library/golang:alpine as generate-openapi-spec
	workdir /build

	copy Makefile .swaggo go.* .
	copy internal/web/controller.go ./internal/web/controller.go

	run apk --no-cache add make gcc musl-dev
	run make target/swagger.json

from docker.io/library/node:lts-alpine as build-frontend
	workdir /build

	copy Makefile .
	copy frontend ./frontend
	copy --from=generate-openapi-spec /build/target ./target

	run ls
	run apk --no-cache add make
	run  make \
		--assume-old target/swagger.json \
		frontend/build

from docker.io/library/golang:alpine as build-backend
	workdir /build

	copy Makefile *.yaml go.* .
	copy internal ./internal
	copy cmd ./cmd
	copy --from=generate-openapi-spec /build/target ./target
	copy --from=build-frontend /build/frontend/build ./frontend/build
	copy --from=build-frontend /build/frontend/*.go ./frontend/

	run apk --no-cache add make gcc musl-dev
	run make \
			--assume-old target/swagger.json \
			--assume-old frontend/build \
			build

from docker.io/library/alpine
	workdir /app

	copy --from=build-backend /build/target/plaincooking .

	label org.opencontainers.image.authors="Lukas Dietrich <lukas@lukasdietrich.com>"
	label org.opencontainers.image.source="https://github.com/lukasdietrich/plaincooking"

	volume /data
	expose 8080/tcp

	env PLAINCOOKING_DATABASE_FILENAME "/data/plaincooking.sqlite"
	env PLAINCOOKING_DATABASE_JOURNALMODE "wal"

	cmd [ "/app/plaincooking" ]
