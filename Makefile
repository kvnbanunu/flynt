run-server:
	go run main.go

run-client:
	cd client/ && npm run dev

insert-dummy:
	go run scripts/dummy/main.go

docker-deploy:
	git pull origin main
	docker compose down
	docker compose build --no-cache
	docker compose up -d
	docker image prune -f
	echo "Deployment complete."

clean:
	rm *.db
