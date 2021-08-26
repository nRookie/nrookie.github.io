# The .netrc file

The .netrc file contains login and initialization information used by the auto-login processes. It generally resides in the user's home directory, but a location outside of the home directory can be set using the environment variable NETRC. Both locations are overridden
by the command line option -N. The selected file must be a regular file, or access will be denied.


The following tokens are recognized; they may be separated by spaces, tabs, or new-lines;


## 'machine name'

Identify a remote machine name. The auto-login process searches the .netrc file for a machine token that matches the remote machine 
specified on the ftp command line or as an open command argument. Once a match is made, the subsequent .netrc tokens are processed, stopping when the end of file is reached or another machine or a default token is encountered.

## 'default'

This is the same as machine name except that default matches any name. There can be only one default token, and it must be after all machine tokens. This is normally used as:

``` shell
default login anonymous password user@site
```

thereby giving the user automatic anonymous ftp login to machines not specified in .netrc. This can be overridden by using the -n flag to disable auto-login.




## ‘login name’

Identify a user on the remote machine. If this token is present, the auto-login process will initiate a login using the specified name.

## 'password string'

Supply a password. If this token is present, the auto-login process will supply the specified string if the remote server requires a password as part of the login process. Note that if this token is present in the .netrc file for any user other than anonymous, ftp will abort the auto-login process if the .netrc is readable by anyone besides the user.



# Ref

https://www.gnu.org/software/inetutils/manual/html_node/The-_002enetrc-file.html