browser-bridge v0.1   Open URLs on another computer.


More info:
http://code.google.com/p/browser-bridge/



Building on Linux

First, install Google's Go compilers.
Then, use the build_all.sh to build this.
You will find a new directory named "bin",
containing 2 executable binary files:
	browser-bridge
	browser-bridge_server

The first one is the client, the second is
the server. Put them whereever you want.





Building on Windows

Sure, there are some experimental windows patchs.
But they are not really working right now. Sorry.

If you need a windows server, You may just build it.
Get Go for windows (http://code.google.com/p/gomingw/)
And then manually execute the commands you find in 
build_all.sh






Config

There are two places to put the config file:

	/etc/browserbridge.conf
	~/.browserbridge.conf

If there are config files in both directories,
only ~/.browserbridge.conf will be used.

The config file itself is well documented.






Usage

Client: Just start it, giving it the url the client shall open.

Server: Just let it run; close through kill or Ctrl-C
	(I know this is very untidy, I will improve this.)
