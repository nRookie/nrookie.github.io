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

