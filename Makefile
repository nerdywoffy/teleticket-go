schema_location = db/schema/schema.prisma

generate:
	go run github.com/steebchen/prisma-client-go generate --schema $(schema_location)

migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev --schema $(schema_location)

migrate-deploy:
	go run github.com/steebchen/prisma-client-go migrate deploy --schema $(schema_location)