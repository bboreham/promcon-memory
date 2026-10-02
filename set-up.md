Set-up:

(probably not needed) docker run --name=procexp -d --pid=host --net=host ncabatoff/process-exporter
docker run --name=cadvisor -d --privileged --cgroupns=host --pid=host --net=host ghcr.io/google/cadvisor --global_housekeeping_interval=5s --max_housekeeping_interval=5s 
docker run --name=prometheus --pid=host --net=host -d -p 9090:9090 -v /Users/bryan/src/github.com/bboreham/memory/:/etc/prometheus prom/prometheus --config.auto-reload --config.file=/etc/prometheus/prometheus.yml --storage.tsdb.path=/prometheus
docker run -d --net=host --name=grafana --volume grafana-storage:/var/lib/grafana grafana/grafana

docker run --privileged --memory=200M -v /Users/bryan/src/github.com/bboreham/memory:/go/src/github.com/bboreham/memory -ti golang

Query:
{__name__=~".*mem.*",id=~"/docker/9df.*",failure_type=""} >0

Drop Caches:
echo 1 > /proc/sys/vm/drop_caches
