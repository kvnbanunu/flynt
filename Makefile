server:
	go run main.go

next:
	cd client/ && npm run dev

seed: clean
	go run scripts/seed/main.go

docker:
	git pull origin main
	docker compose down
	docker compose build --no-cache
	docker compose up -d
	docker image prune -f
	echo "Deployment complete."

clean:
	rm -f *.db
