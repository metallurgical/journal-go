# made with love by norlihazmey.ghazali@gmail.com
# run `make` to compile source proto file into GO and PHP language
# run `make compile` to compile source code into executable binary file
# run `make clean` to delete all generated GO and PHP file

PROTO_PATH=api/proto
PHP_PATH=api/php
GO_PATH=api/golang

all: init build
	
init:
	@echo "Build source proto file.."

build:
	@echo "Build in progress.."
	@echo "Build go proto file.."
	if [ ! -d "api/golang" ]; then \
		echo "Directory api/golang doesnt exist..."; \
		echo "Create directory api/golang.."; \
		mkdir api/golang; \
	else \
		echo "Directory api/golang already exist"; \
		echo "Continue build go proto file.."; \
	fi;	
	protoc -I $(PROTO_PATH) --go_out=plugins=grpc,import_path=golang:$(GO_PATH) $(PROTO_PATH)/journal.proto
	@echo "Build PHP proto file.."
	if [ ! -d "api/php" ]; then \
		echo "Directory api/php doesnt exist..."; \
		echo "Create directory api/php.."; \
		mkdir api/php; \
	else \
		echo "Directory api/php already exist"; \
		echo "Continue build php proto file.."; \
	fi;	
	protoc -I $(PROTO_PATH) --php_out=$(PHP_PATH) --grpc_out=$(PHP_PATH) --plugin=protoc-gen-grpc=/Users/metallurgical/temp/grpc/bins/opt/grpc_php_plugin $(PROTO_PATH)/journal.proto

compile:
	@echo "Compile code into executable file"
	go build

clean:
	@echo "Cleaning all generated proto file(PHP and GO).. Yezzaa.."
	rm api/go/journal.pb.go
	rm -rf api/php/*
	@echo "All generated files removed. DONE!"

