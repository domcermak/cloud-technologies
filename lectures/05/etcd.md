1. Test connection
   ```shell
   etcdctl version
   ```
2. Basic operations
   ```shell   
   etcdctl put foo bar
   etcdctl put foo1 bar1
   etcdctl put foo2 bar2
   
   etcdctl get foo
   etcdctl get foo --hex
   etcdctl get foo --print-value-only
   etcdctl get foo foo3   
   etcdctl get --prefix foo
   etcdctl get --prefix --limit=2 foo
      
   etcdctl del foo
   etcdctl del --prefix foo
   
   etcdctl put goo bar
   # gte byte value
   etcdctl get --from-key foo   
   ```
3. Revisions
   ```shell
   etcdctl put foo bar -w fields
   etcdctl put foo bar2 -w fields
   etcdctl put foo2 bar2 -w fields
      
   etcdctl get --prefix foo
   etcdctl get --prefix foo --rev=43454
   ```
4. Watch
   ```shell
   etcdctl watch --prefix foo
   etcdctl watch --prefix --prev-kv foo
   etcdctl watch --rev=2 --prefix --prev-kv foo
   ```
5. Kompakce
   ```shell
   etcdctl compact 5
   ```
6. Lease
   ```shell
   etcdctl lease grant 60
   etcdctl put --lease=$LEASE test test
   etcdctl lease revoke $LEASE
   
   etcdctl lease keep-alive $LEASE
   ```
7. RBAC
   ```shell

   ```
8. Membership 
   ```
   
   
   ```
9. Metriky - more info in next lectures
   ```
   curl http://localhost:2379/metrics > metrics.txt
   cat metrics.txt
   
   ```
