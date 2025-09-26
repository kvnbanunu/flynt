run-server:
	go run main.go

insert-dummy:
	go run scripts/dummy/main.go

run-client:
	cd client/ && npm run dev

clean:
	rm *.db
