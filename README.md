# munin-urlhealth-plugin

MUNIN Plugin for monitoring a single URL.

## Usecase

1. Your web server is unstable, it's goes UPs and DOWNs regularly.
2. The performance of your web server is somehow terrible and you want to know when.
3. You use MUNIN to monitor the server but you have no MUNIN plugin that monitor a single URL.

## Get your plugin

First, you have to get to the plugin's binnary.

There're 2 ways to obtain one.

1. Build your self from source code.
2. Get the binary that you can download from our releases.

### Buld your self

This plugin is made using Golang programming language.
Thus you must have installed golang in the computer you use
for building this plugin.

Another thing is, you got to have a machine with `make` installed.

Step 1 : Clone the project into your computer.

```bash
$ git clone git@github.com:hyperjumptech/munin-urlhealth-plugin.git
$ cd munin-urlhealth-plugin
``` 

Step 2 : Depends on your target server, you can build using `make`. For example, build for linux server:

```bash
$ make build-linux
go fmt ./...
mkdir -p build/linux
env GOOS=linux GOARCH=amd64 go build -o build/linux ./...
```
The build result can be obtained from within the `build` folder

You have `build-linux`, `build-macos` and `build-windows`. Or simply `make build` to build them all.

## Installing the binary

Once you've a copy of the plugin's binary, you should copy the binary 
from your build machine to your server (running munin-node).

```bash
scp path/to/munin-urlhealth-plugin myusr@targethost:/home/myusr/munin-urlhealth-plugin
```
Now, ssh to that server.

```bash
ssh myusr@targethost
```

Copy the binary to the location of all munin plugin binary, make sure its owned by root and executable

```bash
$ sudo mv ~/munin-urlhealth-plugin /usr/share/munin/plugins
$ sudo chown root:root /usr/share/munin/plugins
$ sudo chmod -c 755 /usr/share/munin/plugins/munin-urlhealth-plugin
```

make a symlink of your copied plugin to munin's plugin directory

```bash
$ ln -s /usr/share/munin/plugins/munin-urlhealth-plugin /etc/munin/plugins
```

as `root`, edit `/etc/munin/plugin-conf.d/munin-node` and add the following:

```bash
[munin-urlhealth-plugin]
user root
env.MonitorURL http://urltotest:123/somepath
```

Please note the `env.MonitorURL` value is the URL that we want to monitor.
Change it to the URL of your liking.

---

You are done, what's left is to restart `munin-node`

```bash
$ sudo service munin-node restart
```

## How to work with the plugin

This plugin basically monitoring a HTTP URL as 
specified in the `env.MonitorURL` environment variable
you've configured in the `munin-node` file

While monitoring this URL, the plugin reports:

1. **code** The response code of the url when it's called. and
2. **response** The response time taken in millisecond for the server (serving the URL) to response back.

If the call is timed-out, or any http client error
the plugin will yield response code 599.

If the response time is 0, mostlikely that the URL is
never get called due to error, network error or invalid URL.

So, combination of code 599 and response time 0 is bad. 