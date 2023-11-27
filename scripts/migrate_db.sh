DB_USER=${POSTGRES_USER:=postgres}
DB_PASSWORD=${POSTGRES_PASSWORD:=password}
DB_NAME=${POSTGRES_DB:=healthchecker}
DB_PORT=${POSTGRES_PORT:=5432}

migrate -source file://../migrations -database postgres://$DB_USER:$DB_PASSOWRD@$DB_HOST:$DB_PORT/$DB_NAME up