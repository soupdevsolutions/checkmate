create-migration name:
    @echo "Creating migration..."
    @migrate create -ext sql -dir ./migrations -seq {{name}} 

start-local:
    @echo "Initializing local database..."
    @chmod +x ./scripts/init_db.sh
    @./scripts/init_db.sh
    @echo "Starting web server..."
    @cd src && go run .