**logrotate** is designed to ease administration of systems that generate large numbers of log files. It allows automatic rotation, compression, removal, and mailing of log files. Each log file may be handled daily, weekly, monthly, or when it grows too large.

Normally, **logrotate** is run as a daily cron job. It will not modify a log multiple times in one day unless the criterion for that log is based on the log's size and **logrotate** is being run multiple times each day, or unless the **-f** or **--force** option is used.



Any number of config files may be given on the command line. Later config files may override the options given in earlier files, so the order in which the **logrotate** config files are listed is important. Normally, a single config file which includes any other config files which are needed should be used. See below for more information on how to use the **include** directive to accomplish this. If a directory is given on the command line, every file in that directory is used as a config file.



If no command line arguments are given, **logrotate** will print version and copyright information, along with a short usage summary. If any errors occur while rotating logs, **logrotate** will exit with non-zero status.

