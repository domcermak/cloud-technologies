

1. Init 
    ```shell
    docker version
    docker ps
    ```
2. Images
   ```shell
   docker images 
   docker search nginx
   // go to https://hub.docker.com/_/nginx
   docker pull nginx
   docker pull nginx:1.20.2
   docker pull mongo:5.0.6
   docker images
   docker image inspect nginx:1.20.2   
   
   docker rmi nginx
   ```
   * [Dive](https://github.com/wagoodman/dive)
     ```shell
     dive nginx:1.20.2
     dive mongo:5.0.6   
     // layers https://hub.docker.com/layers/mongo/library/mongo/5.0.6/images/sha256-416922e55119bd1e49d487bfef2deefe1e7f40b839a863b335fd3c7b10ce0f86?context=explore
     ```   
3. Basic operations
   ```   
   docker run --name nginx -it -p 80:80 nginx
      
   // list, start, stop, remove
   docker ps -a
   docker start nginx
   docker stop nginx
   docker rm nginx   
      
   docker run --name nginx -d -p 80:80 nginx   
   // enter running container
   docker exec -it nginx sh
   
   // container stats 
   docker stats nginx 
   
   // container processes
   docker top nginx
   
   // container logs  
   docker logs nginx
   docker logs -f nginx
   ```   
4. Docker networking
   ```shell
   docker port nginx
   ```
5. Docker volumes
   ```shell
   docker volume ls
   docker volume create my-data
   docker volume inspect my-data
   docker run --rm -v my-data:/data -it ubuntu:20.05 /bin/bash 
   docker run --rm -v my-data:/data -it ubuntu:20.05 /bin/bash
   
   docker run --rm -v "$(PWD)/testdir:/testdir" -it ubuntu:20.05 /bin/bash			 
   ```
6. Commit and tag
   ```shell
   docker tag nginx:1.20.2 my-nginx:1.20.2
   docker tag d933d21f my-nginx:1.20.2
   docker run -it ubuntu:20.05 /bin/bash
   // >> mkdir mydir 
   docker commit $(CONTAINER_ID)
   docker tag $(IMAGE_ID) my-ubuntu:20.05
   docker run -it my-ubuntu:20.05 /bin/bash      
   ```
7. Export/Import
   ```shell
   docker save mongo:5.0.6  > mongo.tar
   ls -sh mongo.tar
   docker load < mongo.ta
   ```   
8. Push/Pull
   ```shell
   docker tag ubuntu:20.05 ondrejsmola/ubuntu:20.05
   docker push ondrejsmola/ubuntu:20.05
   docker pull ondrejsmola/ubuntu:20.05
   ```
