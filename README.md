# Introduction
Microservice implementation to handle journal data using GRPC with PHP as client and GO as server. - Mainly used in internal application -

## Add .env
Copy from example. `cp .env.example .env`

## Setup
```bash
$ go get github.com/metallurgical/journal-go
$ make # to build services files -> use in server
$ make build-local # same as above, use in local # may need to change directory path in Makefile to reflect yours
$ make compile # to compile code into executable file
$ make clean # clean all generated files and executable file 
```

## Global Command

```
sudo cp $GOPATH/bin/journal-go /usr/local/bin/journal-go
```

## Autostart Script
Create a service file in `/lib/systemd/system/journalgo.service` and put following command:

```
[Unit]
Description=Autorun golang grpc server for journal application to consume
Wants=network.target
After=network.target

[Service]
Type=simple
DynamicUser=yes
TimeoutSec=30
Restart=always
RestartSec=3
ExecStart=/path/to/main/binary -env=/path/to/env/file
# Eg: ExecStart=/usr/local/bin/journal-go -env=/path/to/env/file

```
Enable Services:
```
sudo systemctl enable journalgo
```

Start Services:
```
sudo systemctl start journalgo
```

Stop Services:
```
sudo systemctl stop journalgo
```

Restart Services:
```
sudo systemctl restart journalgo
```

Check Services's status:
```
sudo systemctl status journalgo
```