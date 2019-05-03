# made with love by norlihazmey.ghazali@gmail.com
# run `make` to compile source proto file into GO and PHP language
# run `make compile` to compile source code into executable binary file
# run `make clean` to delete all generated GO and PHP file

PROTO_PATH=api/proto
PHP_PATH=api/php
GO_PATH=api/go

all: init build
	
init:
	@echo "Build source proto file.."

build:
	@echo "Build in progress.."
	@echo "Build go proto file.."
	if [ ! -d "api/go" ]; then \
		echo "Directory api/go doesnt exist..."; \
		echo "Create directory api/go.."; \
		mkdir api/go; \
	else \
		echo "Directory go already exist"; \
		echo "Continue build go proto file.."; \
	fi;	
	protoc -I $(PROTO_PATH) --go_out=plugins=grpc:$(GO_PATH) $(PROTO_PATH)/journal.proto
	@echo "Build PHP proto file.."
	if [ ! -d "api/php" ]; then \
		echo "Directory api/php doesnt exist..."; \
		echo "Create directory api/php.."; \
		mkdir api/php; \
	else \
		echo "Directory api/php already exist"; \
		echo "Continue build php proto file.."; \
	fi;	
	protoc -I $(PROTO_PATH) --php_out=plugins=grpc:$(PHP_PATH) $(PROTO_PATH)/journal.proto

compile:
	@echo "Compile code into executable file"
	go build

clean:
	@echo "Cleaning all generated proto file(PHP and GO).. Yezzaa.."
	rm api/go/journal.pb.go
	rm -rf api/php/*
	@echo "All generated files removed. DONE!"

