``` shell
Name:             go-demo-2
│ Namespace:        default
│ Address:          10.23.184.141
│ Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
│ Rules:
│   Host        Path  Backends
│   ----        ----  --------
│   *
│               /demo   go-demo-2-api:8080 (10.16.1.30:8080,10.16.1.31:8080,10.16.2.14:8080)
│ Annotations:  ingress.kubernetes.io/ssl-redirect: false
│               kubernetes.io/ingress.class: nginx
│               nginx.ingress.kubernetes.io/ssl-redirect: false
│ Events:
│   Type    Reason  Age                From                      Message
│   ----    ------  ----               ----                      -------
│   Normal  Sync    12m (x2 over 13m)  nginx-ingress-controller  Scheduled for sync
```

(<error: endpoints "default-http-backend" not found>)

https://github.com/nginxinc/kubernetes-ingress/issues/966

https://kubernetes.github.io/ingress-nginx/troubleshooting/



delete some svc and ingress

``` shell
 nginx-ingress-controller-59f5c88994-wjlxx I0314 13:24:44.613486       1 controller.go:159] "Configuration changes detected, backend reload required"
│ nginx-ingress-controller-59f5c88994-wjlxx I0314 13:24:44.666707       1 controller.go:176] "Backend successfully reloaded"
│ nginx-ingress-controller-59f5c88994-wjlxx I0314 13:24:44.667436       1 event.go:282] Event(v1.ObjectReference{Kind:"Pod", Namespace:"network", Name:"nginx-ingress-controller-59f5c88994-wjlxx"
│
```







``` yaml
 Context: kubernetes-admin@kubernetes              <c>       Copy                                                                                                          ____  __.________
 Cluster: kubernetes                               <n>       Next Match                                                                                                   |    |/ _/   __   \______
 User:    kubernetes-admin                         <shift-n> Prev Match                                                                                                   |      < \____    /  ___/
 K9s Rev: v0.25.18                                 <r>       Toggle Auto-Refresh                                                                                          |    |  \   /    /\___ \
 K8s Rev: v1.23.4                                  <f>       Toggle FullScreen                                                                                            |____|__ \ /____//____  >
 CPU:     3%                                                                                                                                                                      \/            \/
 MEM:     48%
┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─o─s─n─t─o─k─n─i─L─g─(─e Describe(network/nginx-ingress-controller) ───────────────────────────────────────────────────────────────────────────
│ Name:                     nginx-ingress-controller
│ Namespace:                network
│ Labels:                   app.kubernetes.io/component=controller
│                           app.kubernetes.io/instance=nginx-ingress-controller
│                           app.kubernetes.io/managed-by=Helm
│                           app.kubernetes.io/name=nginx-ingress-controller
│                           helm.sh/chart=nginx-ingress-controller-9.1.11
│ Annotations:              meta.helm.sh/release-name: nginx-ingress-controller
│                           meta.helm.sh/release-namespace: network
│ Selector:                 app.kubernetes.io/component=controller,app.kubernetes.io/instance=nginx-ingress-controller,app.kubernetes.io/name=nginx-ingress-controller
│ Type:                     LoadBalancer
│ IP Family Policy:         SingleStack
│ IP Families:              IPv4
│ IP:                       10.104.194.180
│ IPs:                      10.104.194.180
│ Port:                     http  80/TCP
│ TargetPort:               http/TCP
│ NodePort:                 http  32065/TCP
│ Endpoints:                10.16.1.23:80
│ Port:                     https  443/TCP
│ TargetPort:               https/TCP
│ NodePort:                 https  32115/TCP
│ Endpoints:                10.16.1.23:443
│ Session Affinity:         None
│ External Traffic Policy:  Cluster
│ Events:                   <none>
```



``` shell
[root@10-23-75-240 k8s-specs]# curl   10.16.1.23:80/demo
<html>
<head><title>404 Not Found</title></head>
<body>
<center><h1>404 Not Found</h1></center>
<hr><center>nginx</center>
</body>
</html>
[root@10-23-75-240 k8s-specs]# curl   10.16.1.23:80/demo/hello
<html>
<head><title>503 Service Temporarily Unavailable</title></head>
<body>
<center><h1>503 Service Temporarily Unavailable</h1></center>
<hr><center>nginx</center>
</body>
</html>
[root@10-23-75-240 k8s-specs]# curl   10.16.1.23:80/demo/hello
<html>
<head><title>503 Service Temporarily Unavailable</title></head>
<body>
<center><h1>503 Service Temporarily Unavailable</h1></center>
<hr><center>nginx</center>
</body>
</html>
```





Redeploy the go-demo2

``` shell
[root@10-23-75-240 k8s-specs]# kubectl create     -f deploy/go-demo-2.yml
deployment.apps/go-demo-2-db created
service/go-demo-2-db created
deployment.apps/go-demo-2-api created
service/go-demo-2-api created


```





## ingress related service

![](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314213622822.png)



![image-20220314213646900](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314213646900.png)



![image-20220314213731821](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314213731821.png)





## ingress related pods



![image-20220314213910507](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314213910507.png)

![image-20220314214032710](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314214032710.png)



## ingress related deploy



![image-20220314214137152](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314214137152.png)



![image-20220314214153915](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314214153915.png)



![image-20220314214224414](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314214224414.png)



The `go-demo-2` Service we’re currently using is no longer properly configured for our Ingress setup. Using `type: NodePort`, it is configured to export the port `8080` on all of the nodes. Since we’re expecting users to access the application through the Ingress Controller on port `80`, there’s probably no need to allow external access through the port `8080` as well.

We should switch to the `ClusterIP` type. That will allow direct access to the Service only within the cluster, thus limiting all external communication through Ingress.



``` yaml
---

apiVersion: v1
kind: Service
metadata:
  name: go-demo-2-api
spec:
  type: ClusterIP
  ports:
  - port: 8080
  selector:
    type: api
    service: go-demo-2
```

``` shell
kubectl create \

​    -f ingress/go-demo-2.yml \

​    --record --save-config



curl -i "http://$IP/demo/hello"
```







IP is the address

![image-20220314221413972](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314221413972.png)





## create devops-tools kit



``` shell
kubectl create \
    -f ingress/devops-toolkit.yml \
    --record --save-config
```





``` shell
[root@10-23-75-240 k8s-specs]# kubectl get ing
NAME             CLASS    HOSTS   ADDRESS         PORTS   AGE
devops-toolkit   <none>   *       10.23.184.141   80      37s
go-demo-2        <none>   *       10.23.184.141   80      13m
```





``` shell
[root@10-23-75-240 k8s-specs]# curl 10.104.194.180/demo/hello
hello, world!
[root@10-23-75-240 k8s-specs]# curl 10.104.194.180
<!DOCTYPE HTML>
<html lang="en-us">

	<head>
	<meta charset="utf-8" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no" />
	<meta name="description" content="Books and courses dedicated to DevOps practices and tools">
	<meta name="author" content="Viktor Farcic">
	<meta name="generator" content="Hugo 0.93.2" />
	<title>The DevOps Toolkit Series</title>
	<!-- Stylesheets -->

	<link rel="stylesheet" href="/css/main.css"/>





	<!-- Custom Fonts -->
	<link href="/css/font-awesome.min.css" rel="stylesheet" type="text/css">


	<link rel="shortcut icon" type="image/x-icon" href="/favicon.ico">
	<link rel="icon" type="image/x-icon" href="/favicon.ico">


	<!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
	<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
	<!--[if lt IE 9]>
	<script src="js/ie/html5shiv.js"></script>
	<script src="js/ie/html5shiv.jsrespond.min.js"></script>
	<![endif]-->

<script type="application/javascript">
var doNotTrack = false;
if (!doNotTrack) {
	(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');
	ga('create', 'UA-174219852-1', 'auto');

	ga('send', 'pageview');
}
</script>
</head>

	<body>

	<!-- Wrapper -->
		<div id="wrapper">

			<!-- Header -->
    <header id="header" class="alt">
        <a href="/" class="logo"><strong>The DevOps Toolkit</strong> <span>By Viktor Farcic</span></a>
        <nav>
            <a href="#menu">Menu</a>
        </nav>
    </header>

<!-- Menu -->
    <nav id="menu">
        <ul class="links">

                <li><a href="/">Home</a></li>


        </ul>
        <ul class="actions vertical">


        </ul>
    </nav>

			<!-- Banner -->
    <section id="banner" class="major">
        <div class="inner">
            <header class="major">
                <h1>The DevOps Toolkit Series</h1>
            </header>
            <div class="content">
                <p>Where DevOps becomes practice</p>
                <ul class="actions">
                    <li><a href="#one" class="button next scrolly">Get Started</a></li>
                </ul>
            </div>
        </div>
    </section>

		<!-- Main -->
			<div id="main">


					<!-- Header -->
    <section id="one" class="tiles">

        <article>
            <span class="image">
                <img src="img/youtube.png" alt="" />
            </span>
            <header class="major">
                <h3><a href="/posts/youtube" class="link">The YouTube Channel</a></h3>
                <p>DevOps Toolkit</p>
            </header>
        </article>

        <article>
            <span class="image">
                <img src="img/catalog-small.jpg" alt="" />
            </span>
            <header class="major">
                <h3><a href="/posts/catalog" class="link">The DevOps Toolkit</a></h3>
                <p>Catalog, Patterns, And Blueprints</p>
            </header>
        </article>

        <article>
            <span class="image">
                <img src="img/devops23-small.jpg" alt="" />
            </span>
            <header class="major">
                <h3><a href="/posts/devops-23" class="link">The DevOps 2.3 Toolkit</a></h3>
                <p>Kubernetes</p>
            </header>
        </article>

    </section>




			</div>

		<!-- Contact -->


		<!-- Footer -->

				<!-- Footer -->
    <footer id="footer">
        <div class="inner">
            <ul class="icons">

                    <li><a href="https://www.twitter.com/vfarcic" class="icon alt fa-twitter" target="_blank"><span class="label">Twitter</span></a></li>

                    <li><a href="https://www.github.com/vfarcic" class="icon alt fa-github" target="_blank"><span class="label">GitHub</span></a></li>

                    <li><a href="https://www.linkedin.com/in/viktorfarcic/" class="icon alt fa-linkedin" target="_blank"><span class="label">LinkedIn</span></a></li>

            </ul>
            <ul class="copyright">
                <li>&copy; Viktor Farcic</li>

            </ul>
        </div>
    </footer>



		</div>

	<!-- Scripts -->
		<!-- Scripts -->
    <!-- jQuery -->
    <script src="/js/jquery.min.js"></script>
    <script src="/js/jquery.scrolly.min.js"></script>
    <script src="/js/jquery.scrollex.min.js"></script>
    <script src="/js/skel.min.js"></script>
    <script src="/js/util.js"></script>



    <!-- Main JS -->
    <script src="/js/main.js"></script>


<script type="application/javascript">
var doNotTrack = false;
if (!doNotTrack) {
	window.ga=window.ga||function(){(ga.q=ga.q||[]).push(arguments)};ga.l=+new Date;
	ga('create', 'UA-174219852-1', 'auto');

	ga('send', 'pageview');
}
</script>
<script async src='https://www.google-analytics.com/analytics.js'></script>
</body>
```



Ingress is a (kind of) Service that runs on all nodes of a cluster. A user can send requests to any and, as long as they match one of the rules, they will be forwarded to the appropriate Service.

![image-20220314222446680](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/network/image-20220314222446680.png)



