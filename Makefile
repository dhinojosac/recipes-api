.PHONY: mongo-run

mongo-run:
	docker run -d --name mongodb \
	-v $(pwd)/database/data:/data/db \
	-e MONGO_INITDB_ROOT_USERNAME=admin \
	-e MONGO_INITDB_ROOT_PASSWORD=password \
	-p 27017:27017 \
	mongo:4.4.3

mongo-stop:
	docker stop mongodb

mongo-remove:
	docker rm -f mongodb || true