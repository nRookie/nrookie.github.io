## extern

when we write

``` 
int foo(int arg1, char arg2);

```

The compiler treats it as

```
extern int foo(int arg1, char arg2)
```

Since the extern keyword extends the function’s visibility to the whole program, the function can be used (called) anywhere in any of the files of the whole program, provided those files contain a declaration of the function. (With the declaration of the function in place, the compiler knows the definition of the function exists somewhere else and it goes ahead and compiles the file). So that’s all about extern and functions.



To begin with, how would you declare a variable without defining it? You would do something like this:



``` 
extern int var;
```

how would you define var? You would do this

``` 
int var;
```

A declaration can be done any number of times but definition only once.
the extern keyword is used to extend the visibility of variables/functions.
Since functions are visible throughout the program by default, the use of extern is not needed in function declarations or definitions. Its use is implicit.
When extern is used with a variable, it’s only declared, not defined.
As an exception, when an extern variable is declared with initialization, it is taken as the definition of the variable as well.



## char * char[] difference

https://www.geeksforgeeks.org/whats-difference-between-char-s-and-char-s-in-c/