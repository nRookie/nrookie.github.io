Q：想复制含有某个关键字的所有行到另外一个文件中，该如何操作呢？
例如：
<1>this is a book;
<2>this is a dog;
<3>this is a english book;

要将所有含book的行copy出来。

方法：
" Clear register A
:let @a=""
" Append all lines which matchs book to register A
:g/book/y A
" Open a new buffer
:new
" Paste content of register A into the new buffer
:put a

解释：
:let @a="" 使用let命令寄存器a里的内容清空
:g/book/y A 把所有包含book的行都添加到寄存器a中。注：此处是A而不是a，A意味着符合要求的行都被追加到寄存器a中，而a则意味着符合要求的行都会替代寄存器里的内容，如果用a就会导致最后寄存器里只有符合要求的最后一行。
至此，所有包含book的行都在寄存器a里面了。
:put a 把寄存器a里的内容粘贴出来
也可以直接写到文件中去：
:g/book/. w >> filename


附：vim剪切板与命令行
Q: 如何将命令模式下复制的东西黏贴到:命令行里？
:<C-r>"
:help quotequote
或者：
：shift+insert
