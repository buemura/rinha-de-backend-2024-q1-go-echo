# Rinha de Backend 2024 Q1

- [Original challenge repository](https://github.com/zanfranceschi/rinha-de-backend-2024-q1)
- [Source code](https://github.com/buemura/rinha-de-backend-2024-q1-go-echo)

## Tech Stack

- Go
- echo
- pgx
- PostgreSQL
- NGINX
- Docker

## How to run

- Run: **(Requires Docker and docker-compose)**

```bash
# Prepare environment
sh scripts/env_up.sh
```

```bash
# Stress Test
cd .gatling
sh executar-teste-local.sh
```

```bash
# Run app locally
go run cmd/http/main.go
```

## Author

<div>
  <a href="https://www.linkedin.com/in/bruno-uemura/"><img src="https://img.shields.io/badge/linkedin-0077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white"></a>
  <a href="https://github.com/buemura/"><img src="https://img.shields.io/badge/github-3b4c52.svg?style=for-the-badge&logo=github&logoColor=white"></a>
</div>

## Result

- Tested machine specs: Virtual Machine - Linux Ubuntu 22.04.3 LTS 64-bit | 16GB RAM | 4-core - AMD® Ryzen 5 5600 6-core processor

![Result](.docs/result.png)
