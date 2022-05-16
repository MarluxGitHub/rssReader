docker-compose down
docker-compose build

echo "Starting backend"

docker-compose up -d backend

echo "Starting frontend"

docker-compose up -d frontend