

## Prometheus 使用



## 1.部署配置

1.  Prometheus 搭建

```shell

#不挂载外部配置
docker run -p 9090:9090  prom/prometheus

#挂载外部配置
docker run -p 9090:9090 -v /etc/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus


```

2.  Grafana 搭建

   ```shell
   #grafana
   docker run -d --name=grafana -p 3000:3000 grafana/grafana
   ```



## 2.配置  prometheus.yaml

```yaml

# my global config
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).
 
# Alertmanager configuration
alerting:
  alertmanagers:
  - static_configs:
    - targets:
       - alertmanager:9093
 
# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"
 
# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'prometheus'
 
    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.
    scrape_interval: 5s
    static_configs:
    - targets: ['localhost:9090']
 
  - job_name: 'node'
    scrape_interval: 10s
    static_configs:
      - targets: ['localhost:9100']

```



## 3.导入datasource

https://blog.csdn.net/qq_28846087/article/details/100024118



## 4.服务端埋点



#### 默认采集

```go
import (
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "log"
    "net/http"
)
func main() {
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```



#### 选择性采集

```go
func init() {
    //Metrics have to be registered to be exposed:
    prometheus.MustRegister(prometheus.NewGoCollector())
}
```



#### 自定义采集

实现下面的接口

```go
type Collector interface {
	// Describe sends the super-set of all possible descriptors of metrics
	// collected by this Collector to the provided channel and returns once
	// the last descriptor has been sent. The sent descriptors fulfill the
	// consistency and uniqueness requirements described in the Desc
	// documentation.
	//
	// It is valid if one and the same Collector sends duplicate
	// descriptors. Those duplicates are simply ignored. However, two
	// different Collectors must not send duplicate descriptors.
	//
	// Sending no descriptor at all marks the Collector as “unchecked”,
	// i.e. no checks will be performed at registration time, and the
	// Collector may yield any Metric it sees fit in its Collect method.
	//
	// This method idempotently sends the same descriptors throughout the
	// lifetime of the Collector. It may be called concurrently and
	// therefore must be implemented in a concurrency safe way.
	//
	// If a Collector encounters an error while executing this method, it
	// must send an invalid descriptor (created with NewInvalidDesc) to
	// signal the error to the registry.
	Describe(chan<- *Desc)
	// Collect is called by the Prometheus registry when collecting
	// metrics. The implementation sends each collected metric via the
	// provided channel and returns once the last metric has been sent. The
	// descriptor of each sent metric is one of those returned by Describe
	// (unless the Collector is unchecked, see above). Returned metrics that
	// share the same descriptor must differ in their variable label
	// values.
	//
	// This method may be called concurrently and must therefore be
	// implemented in a concurrency safe way. Blocking occurs at the expense
	// of total performance of rendering all registered metrics. Ideally,
	// Collector implementations support concurrent readers.
	Collect(chan<- Metric)
}
```

聚集器

```go
reg := prometheus.NewPedanticRegistry()
reg.MustRegister(自定义的收集器)//

gatherers := prometheus.Gatherers{
        prometheus.DefaultGatherer,
    	reg,
    }
    h := promhttp.HandlerFor(gatherers,
        promhttp.HandlerOpts{
        ErrorHandling: promhttp.ContinueOnError,
    })

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
    h.ServeHTTP(w, r)
    })
    if err := http.ListenAndServe(":8080", nil); err != nil {
    }
```



ref：

https://studygolang.com/articles/17959

https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#NewBuildInfoCollector

https://blog.csdn.net/u014029783/article/details/80001251