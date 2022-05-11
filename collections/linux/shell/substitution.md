### $() command Substitution



*“Command substitution allows the output of a command to replace the command itself. Bash performs the expansion by executing the command and replacing the command substitution with the standard output of the command, with any trailing newlines deleted. Embedded newlines are not deleted, but they may be removed during word splitting.”* Command substitution occurs when a command is enclosed as follows:



``` shell
$(command)
or
`command`

$ echo $(date)
$ echo ‘date’
```





### ${} Parameter Substitution/Expansion

A parameter, in Bash, is an entity that is used to store values. A parameter can be referenced by a number, a name, or by a special symbol. When a parameter is referenced by a number, it is called a **positional parameter**. When a parameter is referenced by a name, it is called a **variable**. When a parameter is referenced by a special symbol, it means they are autoset parameters with special uses.

**Parameter expansion/substitution** is the process of fetching the value from the referenced entity/parameter. It is like you are expanding a variable to fetch its value.



The simplest possible parameter expansion syntax is the following:

Here is how you can use the parameter expansion in Bash:

``` shell
*${parameter}*
```



For example, the simplest usage is to substitute the parameter by its value:

``` shell
$ name="john doe"
$ echo “*${name}*”
```



This command will substitute the value of the variable “name” to be used by the echo command:

You might be thinking that the same can be achieved by avoiding the curly braces as follows:





The answer is that during parameter expansion, these curly braces help in delimiting the variable name. Let us explain what we mean by limiting here. Let me run the following command on my system:



The result did not print the value of the variable name as the system thought that I was referring to the variable “name_”. Thus, my variable name was not “delimited”. The curly braces in the following example will delimit the variable name and return the results as follows:



Here are all the ways in which variables are substituted in Shell:



