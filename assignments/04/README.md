

## Setup

1. Github [repository](https://github.com/etcd-io/etcd)
2. Download 3.5.1 from etcd releases [page](https://github.com/etcd-io/etcd/releases)
    * [Windows](https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-windows-amd64.zip)
    * [Linux](https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-linux-amd64.tar.gz)
3. Add `etcdctl.exe` to your PATH
4. Start server using `docker-compose.yml`
   ```shell
   docker compose up -d
   ```

## Etcd kv server

* GET/POST/DELETE
* read past versions
* watch client
* compare and swap client


