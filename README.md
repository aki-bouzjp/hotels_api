
### docker-compose command
Build containers
```bash
docker build ./ -t app --progress=plain
```

Start containers
```bash
docker-compose up -d
```

Remove volume
```bash
docker-compose down --rmi all --volumes --remove-orphans
```

Connect mysql container
```bash
docker-compose exec app /bin/bash
```
