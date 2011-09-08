#!/bin/bash

8g src/copyright_notice.go
8g src/paths.go
8g src/browserbridge_config.go
8g src/browserbridge_client.go
8g src/browserbridge_server.go

8l -o bin/client/browser-bridge browserbridge_client.8
8l -o bin/server/browser-bridge_server browserbridge_server.8

rm *.8
