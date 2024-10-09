.PHONY: build run

build:
	docker build -t ecommerce-user-service .

run:
	docker run -p 80:80 ecommerce-user-service
