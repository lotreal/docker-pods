# docker-pods

Run the kubernetes pods.yaml

安装文件路径：/root/.go/bin/docker-pods

配置文件路径：/etc/docker-pods/config.yaml


## 示例

```bash
$ docker-pods run
/data/maokai/pods.yaml
/usr/bin/docker run -d \
-l "pid=maokai" \
-e "X_ENV=staging" \
-e "X_CONSUL_SERVER=\"master.consul.local\"" \
-e "X_CONSUL_DC=nine" \
-v "/data/maokai/.:/srv/www" \
-v "/data/maokai/etc/staging/consul.json:/etc/consul.d/web.json" \
-v "/data/maokai/etc/staging/httpd.conf:/etc/httpd/vhost.d/default.conf" \
docker.web.dm/apache-php
=============================
PODS ID PODS                    CONTAINER ID
maokai  /data/maokai/pods.yaml  055d1b6a69bd63b753f4713b4c5c70f0e03619be010402bfc8248f5263eea448

$ docker-pods ps
RUNNING PODS ID CONTAINER ID    IP ADDR         IMAGE                           COMMAND                 STATUS          PORTS
YES     maokai  055d1b6a69bd    10.0.1.12       docker.web.dm/apache-php        "/usr/local/bin/star    Up 1 seconds    80/tcp
```
