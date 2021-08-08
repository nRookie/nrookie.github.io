A quick note on troubleshooting connectivity issues before moving on to Service Discovery, if you think you’re experiencing connectivity issues between containers, it’s worth checking the Docker daemon logs as well as container logs.


# Daemon logs

On Windows systems, the daemon logs are stored under ~AppData\Local\Docker, and you can view them in the Windows Event Viewer. On Linux, it depends what init system you’re using. If you’re running a systemd, the logs will go to journald and you can view them with the journalctl -u docker.service command. If you’re not running systemd you should look under the following locations:

- Ubuntu systems running upstart: /var/log/upstart/docker.log
- RHEL-based systems: /var/log/messages
- Debian: /var/log/daemon.log

Managing the log outputs

You can also tell Docker how verbose you want daemon logging to be. To do this, edit the daemon config file (daemon.json) so that debug is set to true and log-level is set to one of the following:

- debug: The most verbose option
- info: The default value and second-most verbose option
- warn: Third most verbose option
- error: Fourth most verbose option
- fatal: Least verbose option

The following snippet from a daemon.json enables debugging and sets the level to debug. It will work on all Docker platforms.

``` shell
{
  <Snip>
  "debug":true,
  "log-level":"debug",
  <Snip>
}
```

Be sure to restart Docker after making changes to the file.

That was the daemon logs. What about container logs?

## Container logs

Logs from standalone containers can be viewed with the docker container logs command, and Swarm service logs can be viewed with the docker service logs command. However, Docker supports lots of logging drivers, and they don’t all work with the docker logs command.


### Logging drivers 

As well as a driver and configuration for daemon logs, every Docker host has a default logging driver and configuration for containers. Some of the drivers include:


- json-file (default)
- journald (only works on Linux hosts running systemd)
- syslog
- splunk 
- gelf


json-file and journald are probably the easiest to configure, and they both work with the docker logs and docker service logs commands. The format of the commands is docker logs <container-name> and docker service logs <service-name>.


The following snippet from a daemon.json shows a Docker host configured to use syslog.

``` shell
{
  "log-driver": "syslog"
}
```


You can configure an individual container, or service, to start with a particular logging driver with the --log-driver and --log-opts flags. These will override anything set in daemon.json.

Container logs work on the premise that your application is running as PID 1 inside the container and sending logs to STDOUT, and errors to STDERR. The logging driver then forwards these “logs” to the locations configured via the logging driver.

If your application logs to a file, it’s possible to use a symlink to redirect log-file writes to STDOUT and STDERR.

The following is an example of running the docker logs command against a container called “vantage-db” configured to use the json-file logging driver.

``` shell

$ docker logs vantage-db
1:C 2 Feb 09:53:22.903 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 2 Feb 09:53:22.904 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1
1:C 2 Feb 09:53:22.904 # Warning: no config file specified, using the default config.
1:M 2 Feb 09:53:22.906 * Running mode=standalone, port=6379.
1:M 2 Feb 09:53:22.906 # WARNING: The TCP backlog setting of 511 cannot be enforced because...
1:M 2 Feb 09:53:22.906 # Server initialized
1:M 2 Feb 09:53:22.906 # WARNING overcommit_memory is set to 0!

```

There’s a good chance you’ll find network connectivity errors reported in the daemon logs or container logs.
