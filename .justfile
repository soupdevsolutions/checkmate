start-local:
    @echo "Initializing local database..."
    @chmod +x ./scripts/init_db.sh
    @./scripts/init_db.sh
    @echo "Starting web server..."
    @cd src && go run .