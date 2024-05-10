setup-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

dev:
	clear && ./bin/air

restartBuild:
	rm -rf build && mkdir build
	
buildGo:
	GOOS=linux GOARCH=amd64  go build -a -ldflags="-s -w" -o build/main main.go 
	
compress:
	zip -j build/main.zip build/main

glg:
	echo "install the gqlgen dev" && go get -d github.com/99designs/gqlgen && echo "generate the graphql" && go run github.com/99designs/gqlgen generate

mainBuild:
	go build -o main main.go

migrate-account:
	go run github.com/steebchen/prisma-client-go migrate dev --schema=./src/prisma/account/schema.prisma

migrate-channel:
	go run github.com/steebchen/prisma-client-go migrate dev --schema=./src/prisma/channel/schema.prisma

migrate-campaign:
	go run github.com/steebchen/prisma-client-go migrate dev --schema=./src/prisma/campaign/schema.prisma

migrate-transaction:
	go run github.com/steebchen/prisma-client-go migrate dev --schema=./src/prisma/transaction/schema.prisma
migrate:
	make migrate-account && make migrate-channel && make migrate-campaign && make migrate-transaction