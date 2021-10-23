https://www.gnu.org/software/autoconf/manual/autoconf-2.66/autoconf.html


## Introduction

https://www.gnu.org/software/autoconf/manual/autoconf-2.66/autoconf.html#Introduction

## The GNU Build System


integrated utilities to finish the job Autoconf started: the GNU build system, whose most important components are Autoconf, Automake, and Libtool. In this chapter, we introduce you to those tools, point you to sources of more information, and try to convince you to use the entire GNU build system for your software.

Automake: Escaping makefile hell
Gnulib: The GNU portability library
Libtool: Building libraries portably
Pointers: More info on the GNU build system


## 3 Making configure Scripts


### Using autoscan to Create configure.ac


### 3.4 Using autoconf to Create configure



## Initializing configure

Every configure script must call AC_INIT before doing anything else that produces output. Calls to silent macros, such as AC_DEFUN, may also occur prior to AC_INIT, although these are generally used via aclocal.m4, since that is implicitly included before the start of configure.ac. The only other required macro is AC_OUTPUT (see Output).


— Macro: AC_INIT (package, version, [bug-report], [tarname], [url])


Process any command-line arguments and perform various initializations and verifications.

Set the name of the package and its version. These are typically used in --version support, including that of configure. The optional argument bug-report should be the email to which users should send bug reports. The package tarname differs from package: the latter designates the full package name (e.g., ‘GNU Autoconf’), while the former is meant for distribution tar ball names (e.g., ‘autoconf’). It defaults to package with ‘GNU ’ stripped, lower-cased, and all characters other than alphanumerics and underscores are changed to ‘-’. If provided, url should be the home page for the package.


## 4.2 Dealing with Autoconf versions

The following optional macros can be used to help choose the minimum version of Autoconf that can successfully compile a given configure.ac.

— Macro: AC_PREREQ (version)
Ensure that a recent enough version of Autoconf is being used. If the version of Autoconf being used to create configure is earlier than version, print an error message to the standard error output and exit with failure (exit status is 63). For example:

          AC_PREREQ([2.66])
This macro may be used before AC_INIT.

— Macro: AC_AUTOCONF_VERSION
This macro was introduced in Autoconf 2.62. It identifies the version of Autoconf that is currently parsing the input file, in a format suitable for m4_version_compare (see m4_version_compare); in other words, for this release of Autoconf, its value is ‘2.66’. One potential use of this macro is for writing conditional fallbacks based on when a feature was added to Autoconf, rather than using AC_PREREQ to require the newer version of Autoconf. However, remember that the Autoconf philosophy favors feature checks over version checks.

You should not expand this macro directly; use ‘m4_defn([AC_AUTOCONF_VERSION])’ instead. This is because some users might have a beta version of Autoconf installed, with arbitrary letters included in its version string. This means it is possible for the version string to contain the name of a defined macro, such that expanding AC_AUTOCONF_VERSION would trigger the expansion of that macro during rescanning, and change the version string to be different than what you intended to check.

## 4.3 Notices in configure
The following macros manage version numbers for configure scripts. Using them is optional.

— Macro: AC_COPYRIGHT (copyright-notice)
State that, in addition to the Free Software Foundation's copyright on the Autoconf macros, parts of your configure are covered by the copyright-notice.

The copyright-notice shows up in both the head of configure and in ‘configure --version’.

— Macro: AC_REVISION (revision-info)
Copy revision stamp revision-info into the configure script, with any dollar signs or double-quotes removed. This macro lets you put a revision stamp from configure.ac into configure without RCS or CVS changing it when you check in configure. That way, you can determine easily which revision of configure.ac a particular configure corresponds to.

For example, this line in configure.ac:
          AC_REVISION([$Revision: 1.30 $])
produces this in configure:

          #!/bin/sh
          # From configure.ac Revision: 1.30



## 4.4 Finding configure Input

— Macro: AC_CONFIG_SRCDIR (unique-file-in-source-dir)
unique-file-in-source-dir is some file that is in the package's source directory; configure checks for this file's existence to make sure that the directory that it is told contains the source code in fact does. Occasionally people accidentally specify the wrong directory with --srcdir; this is a safety check. See configure Invocation, for more information.

Packages that do manual configuration or use the install program might need to tell configure where to find some other shell scripts by calling AC_CONFIG_AUX_DIR, though the default places it looks are correct for most cases.

— Macro: AC_CONFIG_AUX_DIR (dir)
Use the auxiliary build tools (e.g., install-sh, config.sub, config.guess, Cygnus configure, Automake and Libtool scripts, etc.) that are in directory dir. These are auxiliary files used in configuration. dir can be either absolute or relative to srcdir. The default is srcdir or srcdir/.. or srcdir/../.., whichever is the first that contains install-sh. The other files are not checked for, so that using AC_PROG_INSTALL does not automatically require distributing the other auxiliary files. It checks for install.sh also, but that name is obsolete because some make have a rule that creates install from it if there is no makefile.

The auxiliary directory is commonly named build-aux. If you need portability to DOS variants, do not name the auxiliary directory aux. See File System Conventions.

— Macro: AC_REQUIRE_AUX_FILE (file)
Declares that file is expected in the directory defined above. In Autoconf proper, this macro does nothing: its sole purpose is to be traced by third-party tools to produce a list of expected auxiliary files. For instance it is called by macros like AC_PROG_INSTALL (see Particular Programs) or AC_CANONICAL_BUILD (see Canonicalizing) to register the auxiliary files they need.

Similarly, packages that use aclocal should declare where local macros can be found using AC_CONFIG_MACRO_DIR.

— Macro: AC_CONFIG_MACRO_DIR (dir)
Specify dir as the location of additional local Autoconf macros. This macro is intended for use by future versions of commands like autoreconf that trace macro calls. It should be called directly from configure.ac so that tools that install macros for aclocal can find the macros' declarations.

Note that if you use aclocal from Automake to generate aclocal.m4, you must also set ACLOCAL_AMFLAGS = -I dir in your top-level Makefile.am. Due to a limitation in the Autoconf implementation of autoreconf, these include directives currently must be set on a single line in Makefile.am, without any backslash-newlines.



## 4.5 Outputting Files
Every Autoconf script, e.g., configure.ac, should finish by calling AC_OUTPUT. That is the macro that generates and runs config.status, which in turn creates the makefiles and any other files resulting from configuration. This is the only required macro besides AC_INIT (see Input).


— Macro: AC_OUTPUT
Generate config.status and launch it. Call this macro once, at the end of configure.ac.

config.status performs all the configuration actions: all the output files (see Configuration Files, macro AC_CONFIG_FILES), header files (see Configuration Headers, macro AC_CONFIG_HEADERS), commands (see Configuration Commands, macro AC_CONFIG_COMMANDS), links (see Configuration Links, macro AC_CONFIG_LINKS), subdirectories to configure (see Subdirectories, macro AC_CONFIG_SUBDIRS) are honored.

The location of your AC_OUTPUT invocation is the exact point where configuration actions are taken: any code afterwards is executed by configure once config.status was run. If you want to bind actions to config.status itself (independently of whether configure is being run), see Running Arbitrary Configuration Commands.

Historically, the usage of AC_OUTPUT was somewhat different. See Obsolete Macros, for a description of the arguments that AC_OUTPUT used to support.

If you run make in subdirectories, you should run it using the make variable MAKE. Most versions of make set MAKE to the name of the make program plus any options it was given. (But many do not include in it the values of any variables set on the command line, so those are not passed on automatically.) Some old versions of make do not set this variable. The following macro allows you to use it even with those versions.


— Macro: AC_PROG_MAKE_SET
If the Make command, $MAKE if set or else ‘make’, predefines $(MAKE), define output variable SET_MAKE to be empty. Otherwise, define SET_MAKE to a macro definition that sets $(MAKE), such as ‘MAKE=make’. Calls AC_SUBST for SET_MAKE.

If you use this macro, place a line like this in each Makefile.in that runs MAKE on other directories:

     @SET_MAKE@


## 4.6 Performing Configuration Actions
configure is designed so that it appears to do everything itself, but there is actually a hidden slave: config.status. configure is in charge of examining your system, but it is config.status that actually takes the proper actions based on the results of configure. The most typical task of config.status is to instantiate files.

This section describes the common behavior of the four standard instantiating macros: AC_CONFIG_FILES, AC_CONFIG_HEADERS, AC_CONFIG_COMMANDS and AC_CONFIG_LINKS. They all have this prototype:
     AC_CONFIG_ITEMS(tag..., [commands], [init-cmds])
where the arguments are:

tag...
A blank-or-newline-separated list of tags, which are typically the names of the files to instantiate.
You are encouraged to use literals as tags. In particular, you should avoid

          ... && my_foos="$my_foos fooo"
          ... && my_foos="$my_foos foooo"
          AC_CONFIG_ITEMS([$my_foos])
and use this instead:

          ... && AC_CONFIG_ITEMS([fooo])
          ... && AC_CONFIG_ITEMS([foooo])
The macros AC_CONFIG_FILES and AC_CONFIG_HEADERS use special tag values: they may have the form ‘output’ or ‘output:inputs’. The file output is instantiated from its templates, inputs (defaulting to ‘output.in’).

‘AC_CONFIG_FILES([Makefile:boiler/top.mk:boiler/bot.mk)]’, for example, asks for the creation of the file Makefile that contains the expansion of the output variables in the concatenation of boiler/top.mk and boiler/bot.mk.

The special value ‘-’ might be used to denote the standard output when used in output, or the standard input when used in the inputs. You most probably don't need to use this in configure.ac, but it is convenient when using the command line interface of ./config.status, see config.status Invocation, for more details.

The inputs may be absolute or relative file names. In the latter case they are first looked for in the build tree, and then in the source tree. Input files should be text files, and a line length below 2000 bytes should be safe.

commands
Shell commands output literally into config.status, and associated with a tag that the user can use to tell config.status which commands to run. The commands are run each time a tag request is given to config.status, typically each time the file tag is created.
The variables set during the execution of configure are not available here: you first need to set them via the init-cmds. Nonetheless the following variables are precomputed:

srcdir
The name of the top source directory, assuming that the working directory is the top build directory. This is what the configure option --srcdir sets.
ac_top_srcdir
The name of the top source directory, assuming that the working directory is the current build directory.
ac_top_build_prefix
The name of the top build directory, assuming that the working directory is the current build directory. It can be empty, or else ends with a slash, so that you may concatenate it.
ac_srcdir
The name of the corresponding source directory, assuming that the working directory is the current build directory.
tmp
The name of a temporary directory within the build tree, which you can use if you need to create additional temporary files. The directory is cleaned up when config.status is done or interrupted. Please use package-specific file name prefixes to avoid clashing with files that config.status may use internally.
The current directory refers to the directory (or pseudo-directory) containing the input part of tags. For instance, running

          AC_CONFIG_COMMANDS([deep/dir/out:in/in.in], [...], [...])
with --srcdir=../package produces the following values:

          # Argument of --srcdir
          srcdir='../package'
          # Reversing deep/dir
          ac_top_build_prefix='../../'
          # Concatenation of $ac_top_build_prefix and srcdir
          ac_top_srcdir='../../../package'
          # Concatenation of $ac_top_srcdir and deep/dir
          ac_srcdir='../../../package/deep/dir'
independently of ‘in/in.in’.

init-cmds
Shell commands output unquoted near the beginning of config.status, and executed each time config.status runs (regardless of the tag). Because they are unquoted, for example, ‘$var’ is output as the value of var. init-cmds is typically used by configure to give config.status some variables it needs to run the commands.
You should be extremely cautious in your variable names: all the init-cmds share the same name space and may overwrite each other in unpredictable ways. Sorry...

All these macros can be called multiple times, with different tag values, of course!


## 4.7 Creating Configuration Files
Be sure to read the previous section, Configuration Actions.


— Macro: AC_CONFIG_FILES (file..., [cmds], [init-cmds])
Make AC_OUTPUT create each file by copying an input file (by default file.in), substituting the output variable values. This macro is one of the instantiating macros; see Configuration Actions. See Makefile Substitutions, for more information on using output variables. See Setting Output Variables, for more information on creating them. This macro creates the directory that the file is in if it doesn't exist. Usually, makefiles are created this way, but other files, such as .gdbinit, can be specified as well.

Typical calls to AC_CONFIG_FILES look like this:

          AC_CONFIG_FILES([Makefile src/Makefile man/Makefile X/Imakefile])
          AC_CONFIG_FILES([autoconf], [chmod +x autoconf])
You can override an input file name by appending to file a colon-separated list of input files. Examples:

          AC_CONFIG_FILES([Makefile:boiler/top.mk:boiler/bot.mk]
                          [lib/Makefile:boiler/lib.mk])
Doing this allows you to keep your file names acceptable to DOS variants, or to prepend and/or append boilerplate to the file.




###  Standard configure.ac Layout 


The order in which configure.ac calls the Autoconf macros is not important, with a few exceptions. Every configure.ac must contain a call to AC_INIT before the checks, and a call to AC_OUTPUT at the end (see Output). Additionally, some macros rely on other macros having been called first, because they check previously set values of some variables to decide what to do. These macros are noted in the individual descriptions (see Existing Tests), and they also warn you when configure is created if they are called out of order.

To encourage consistency, here is a suggested order for calling the Autoconf macros. Generally speaking, the things near the end of this list are those that could depend on things earlier in it. For example, library functions could be affected by types and libraries.

     Autoconf requirements
     AC_INIT(package, version, bug-report-address)
     information on the package
     checks for programs
     checks for libraries
     checks for header files
     checks for types
     checks for structures
     checks for compiler characteristics
     checks for library functions
     checks for system services
     AC_CONFIG_FILES([file...])
     AC_OUTPUT


## 8 Programming in M4

Autoconf is written on top of two layers: M4sugar, which provides convenient macros for pure M4 programming, and M4sh, which provides macros dedicated to shell script generation.

The most common problem with existing macros is an improper quotation. This section, which users of Autoconf can skip, but which macro writers must read, first justifies the quotation scheme that was chosen for Autoconf and then ends with a rule of thumb. Understanding the former helps one to follow the latter.


## 9.3 Initialization Macros

Initialize the M4sh environment. This macro calls m4_init, then outputs the #! /bin/sh line, a notice about where the output was generated from, and code to sanitize the environment for the rest of the script. Among other initializations, this sets SHELL to the shell chosen to run the script (see CONFIG_SHELL), and LC_ALL to ensure the C locale. Finally, it changes the current diversion to BODY. AS_INIT is called automatically by AC_INIT and AT_INIT, so shell code in configure, config.status, and testsuite all benefit from a sanitized shell environment.



— Macro: m4_esyscmd_s (command)
Like m4_esyscmd, this macro expands to the result of running command in a shell. The difference is that any trailing newlines are removed, so that the output behaves more like shell command substitution.



