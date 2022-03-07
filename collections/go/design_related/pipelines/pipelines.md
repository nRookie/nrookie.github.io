
https://go.dev/blog/pipelines

Fan-out, fan-in
Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that’s closed when all the inputs are closed. This is called fan-in.

We can change our pipeline to run two instances of sq, each reading from the same input channel. We introduce a new function, merge, to fan in the results: