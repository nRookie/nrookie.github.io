![image-20220320114052613](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320114052613.png)



![image-20220320114329106](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320114329106.png)

![image-20220320114512643](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320114512643.png)



![image-20220320114838542](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320114838542.png)





A very common solution is to use a configuration struct.

This has some advantages.

Using this approach, the configuration struct can grow over time as new options are added, while the public API for creating a server itself remains unchanged.

This method can lead to better documentation.

What was once a massive comment block on the `NewServer` function, becomes a nicely documented struct.

Potentially it also enables the callers to use the zero value to signify they they want the default behaviour for a particular configuration option.



![image-20220320114924594](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320114924594.png)



However, this pattern is not perfect.

It has trouble with defaults, especially if the zero value has a well understood meaning.

For example, in the config structure shown here, when `Port` is not provided, `NewServer` will return a `*Server` for listening on port 8080.

But this has the downside that you can no longer explicitly set `Port` to 0 and have the operating system automatically choose a free port, because that explicit 0 is indistinguishable from the fields’ zero value.



![image-20220320115024790](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320115024790.png)



![image-20220320115149740](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320115149740.png)



A common solution to this empty value problem is to pass a pointer to the value instead, thereby enabling callers to use nil rather than constructing an empty value.

In my opinion this pattern has all the problems of the previous example, and it adds a few more.

We still have to pass *something* for this function’s second argument, but now this value could be `nil`, and most of the time *will* be `nil` for those wanting the default behaviour.

It raises the question, is there a difference between passing `nil`, and passing a pointer to an empty value ?

More concerning to both the package’s author, and its callers, is the Server and the caller can now share a reference to the same configuration value. Which gives rise to questions of what happens if this value is mutated after being passed to the `NewServer` function ?

I believe that well written APIs should not require callers to create dummy values to satisfy those rarer use cases.

I believe that we, as Go programmers, should work hard to ensure that nil is never a parameter that needs to be passed to any public function.

And when we do want to pass configuration information, it should be as self explanatory and as expressive as possible.

So now with these points in mind, I want to talk about what I believe are some better solutions.



![image-20220320115318951](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320115318951.png)



To remove the problem of that mandatory, yet frequently unused, configuration value, we can change the `NewServer` function to accept a variable number of arguments.

Instead of passing `nil`, or some zero value, as a signal that you want the defaults, the variadic nature of the function means you don’t need to pass *anything at all*.

And in my book this solves two big problems.

First, the invocation for the default behaviour becomes as concise as possible.

Secondly, `NewServer` now only accepts `Config` values, not pointers to config values, removing `nil` as a possible argument, and ensuring that the caller cannot retain a reference to the server’s internal configuration.

I think this is a big improvement.

But if we’re being pedantic, it still has a few problems.

Obviously the expectation is for you to provide at most one `Config` value. But as the function signature is variadic, the implementation has to be written to cope with a caller passing multiple, possibly contradictory, configuration structs.

Is there a way to use a variadic function signature *and* improve the expressiveness of configuration parameters when needed ?

I think that there is.



![image-20220320115454222](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320115454222.png)



Inside `NewServer`, applying these options is straightforward.

After opening a `net.Listener`, we declare a `Server` instance using that `listener`.

Then, for each option function provided to `NewServer`, we call that function, passing in a pointer to the `Server` value that was just declared.

Obviously, if no option functions were provided, there is no work to do in this loop and so `srv` is unchanged.

And that’s all there is too it.

Using this pattern we can make an API that has

- sensible defaults
- is highly configurable
- can grow over time
- self documenting
- safe for newcomers
- and never requires nil or an empty value to keep the compiler happy

In the few minutes I have remaining I’d like to show you how I improved one of my own packages by converting it to use functional options.



![image-20220320120432020](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320120432020.png)



I’m an amateur hardware hacker, and many of the devices I work with use a USB serial interface. A so a few months ago I wrote a [terminal handling package](https://github.com/pkg/term).

In the prior version of this package, to open a serial device, change the speed and set the terminal to raw mode, you’d have to do each of these steps individually, checking the error at every stage.

Even though this package is trying to provide a friendlier interface on an even lower level interface, it still left too many procedural warts for the user.

Let’s take a look at the package after applying the functional options pattern.





By converting the `Open` function to use a variadic parameter of function values, we get a much cleaner API.

In fact, it’s not just the `Open` API that improves, the grind of setting an option, checking an error, setting the next option, checking the error, that is gone as well.

The default case, still just takes one argument, the name of the device.

For more complicated use cases, configuration functions, defined in the `term` package, are passed to the `Open` function and are applied in order before returning.

This is the same pattern we saw in the previous example, the only thing that is different is rather than being anonymous, these are public functions. In all other respects their operation is identical.



We’ll take a look at how `Speed`, `RawMode`, and `Open`, are implemented on the next slide.



![image-20220320120749902](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320120749902.png)



![image-20220320121148471](/Users/kestrel/developer/nrookie.github.io/collections/go/xunlianying/工程化/image-20220320121148471.png)





 In summary

- Functional options let you write APIs that can grow over time.
- They enable the default use case to be the simplest.
- They provide meaningful configuration parameters.
- Finally they give you access to the entire power of the language to initialize complex values.

In this talk, I have presented many of the existing configuration patterns, those considered idiomatic and commonly in use today, and at every stage asked questions like:

- Can this be made simpler ?
- Is that parameter necessary ?
- Does the signature of this function make it easy for it to be used safely ?
- Does the API contain traps or confusing misdirection that will frustrate ?

I hope I have inspired you to do the same. To revisit code that you have written in the past and pose yourself these same questions and thereby improve it.



