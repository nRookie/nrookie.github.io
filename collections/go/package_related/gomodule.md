```
GO111MODULE=on go get golang.org/x/tools/gopls@latest

```

https://maelvls.dev/go111module-everywhere/


GO111MODULE is an environment variable that can be set when using go for changing how Go imports packages. One of the first pain-points is that depending on the Go version, its semantics change.



GO111MODULE=on will force using Go modules even if the project is in your GOPATH. Requires go.mod to work.



GO111MODULE=off forces Go to behave the GOPATH way, even outside of GOPATH.

GO111MODULE=auto is the default mode. In this mode, Go will behave
similarly to GO111MODULE=on when you are outside of GOPATH,
similarly to GO111MODULE=off when you are inside the GOPATH even if a go.mod is present.

Whenever you are in your GOPATH and you want to do an operation that requires Go modules (e.g., go get a specific version of a binary), you need to do

``` shell
GO111MODULE=on go get github.com/golang/mock/tree/master/mockgen@v1.3.1

```

Now that we know that GO111MODULE can be very useful for enabling the Go Modules behavior, here is the answer: that’s because GO111MODULE=on allows you to select a version. Without Go Modules, go get fetches the latest commit from master. With Go Modules, you can select a specific version based on git tags.

I use GO111MODULE=on very often when I want to switch between the latest version and the HEAD version of gopls (the Go Language Server):



## Private Go Modules and Dockerfile
Many companies choose to use private repositories as import paths. As explained above, we can use GOPRIVATE in order to tell Go (as of Go 1.13) to skip the package proxy and fetch our private packages directly from Github.

But what about building Docker images? How can go get fetch our private repositories from a docker build?

Solution 1: vendoring
With go mod vendor, no need to pass Github credentials to the docker build context. We can just put everything in vendor/ and the problem is solved. In the Dockerfile, -mod=vendor will be required, but developers don’t even have to bother with -mod=vendor since they have access to the private Github repositories anyway using their local Git config

Pros: faster build on CI (~10 to 30 seconds less)
Cons: PRs are bloated with vendor/ changes and the repo’s size might be big
Solution 2: no vendoring
If vendor/ is just too big (e.g., for Kubernetes controllers, vendor/ is about 30MB), we can very well do it without vendoring. That would require to pass some form of GITHUB_TOKEN as argument of docker build, and in the Dockerfile, set something like:

