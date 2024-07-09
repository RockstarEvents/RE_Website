init:
	go mod init eventPlanner
	go get github.com/labstack/echo/v4
	go get github.com/labstack/echo/v4/middleware
	go get github.com/joho/godotenv
clean:
	rm -f eventPlanner
