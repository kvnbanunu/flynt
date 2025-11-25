# Flynt

Check it out live at [flyntapp.io](https://flyntapp.io)

Use Flynt to start your Fyre streak!

This app was developed for BCIT COMP 7082

---

## Contributors

Kevin Van Nguyen | DevOps, FullStack, Owner | [github](https://github.com/kvnbanunu)

Evin Gonzales | Frontend QA | [github](https://github.com/evin-gg)

Lucas Laviolette | Database QA | [github](https://github.com/lucaslav05)

Brandon Rada | Backend QA | [github](https://github.com/BrandonRada)

---

## Project Info

- Go 1.25.1
- Next 15.5.3
- SQLite3
- Make

## Requirements

go-sqlite3 requires CGO_ENABLED=1 in your environment

## Usage

1. Clone the project
```sh
git clone https://github.com/kvnbanunu/flynt
```

2. Setup environment variables
```sh
make environment
```

Fill in the .env and client/.env files

3. Install dependencies
```sh
make dependencies
```

4. Seed the database
```sh
make seed
```

5. Start API server
```sh
make server
```

6. Start client
```sh
make next
```

Open [localhost:<port>](http://localhost:3000)
