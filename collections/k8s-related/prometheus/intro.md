### What is `Prometheus`? [#](https://www.educative.io/courses/advanced-kubernetes-techniques/mE69W4qvE93#What-is-Prometheus?-)

> 📌 `Prometheus` is a database (of sorts) designed to fetch (pull) and store highly dimensional time series data.



### How `Prometheus`'s works? [#](https://www.educative.io/courses/advanced-kubernetes-techniques/mE69W4qvE93#How-Prometheus's-works?-)

`Prometheus'` query language allows us to easily find data that can be used both for graphs and, more importantly, for alerting. It does not attempt to provide a “great” visualization experience. For that, it integrates with [Grafana](https://grafana.com/).

Unlike most other similar tools, we do not push data to `Prometheus`. Or, to be more precise, that is not the common way of getting metrics. Instead, `Prometheus` is a pull-based system that periodically fetches metrics from exporters. There are many third-party exporters we can use. But, in our case, the most crucial exporter is baked into **Kubernetes**. `Prometheus` can pull data from an exporter that transforms information from Kube API. Through it, we can fetch (almost) everything we might need. Or, at least, that’s where the bulk of the information will be coming from.



## Send alerts using `AlertManager`



Finally, storing metrics in `Prometheus` would not be of much use if we are not notified when there’s something wrong. Even when we do integrate `Prometheus` with [Grafana](https://grafana.com/), that will only provide us with dashboards. I assume that you have better things to do than to stare at colorful graphs. So, we’ll need a way to send alerts from `Prometheus` to, let’s say, Slack. Luckily, [Alertmanager](https://prometheus.io/docs/alerting/alertmanager/) allows us just that. It is a separate application maintained by the same community.



We’ll see how all those pieces fit together through hands-on exercises. So, let’s get going and install `Prometheus`, `Alertmanager`, and a few other applications.



![image-20220403185732954](/Users/kestrel/developer/nrookie.github.io/collections/k8s-related/prometheus/image-20220403185732954.png)

