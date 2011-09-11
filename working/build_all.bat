:: please make sure to apply the patches before compiling.
:: The source is optimized for gomingw (http://code.google.com/p/gomingw/)

8g src\copyright_notice.go

8g src\paths.go

8g src\browserbridge_config.go

8g src\browserbridge_client.go
8g src\browserbridge_server.go



:: 8l -o .\browser-bridge.exe browserbridge_client.8

8l -o .\browser-bridge_server.exe browserbridge_server.8


del *.8