help:
	@echo "Available commands: run, stop"

run:
	go build -C ./src/master_service -o ../../build/master_service/compiled && \
	go build -C ./src/api_service -o ../../build/api_service/compiled && \
	go build -C ./src/rpc_service -o ../../build/rpc_service/compiled && \
	go build -C ./src/kafka_service -o ../../build/kafka_service/compiled

	docker compose up -d --build
	
stop: 
	docker compose down
