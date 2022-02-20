

1. [Hello-world](hello-world)
   ```shell
   docker build -t helloworld:0.0.1 .
   ```
2. [Nginx](nginx)
3. [Flask](flask)
4. [Reference](https://docs.docker.com/engine/reference/builder/#entrypoint)
   ```dockerfile
   FROM scratch
   FROM ubuntu
   FROM ubuntu:204
   
   EXPOSE 80 443
   
   ENV MY_ENV="value"
   WORKDIR /opt
   
   COPY myapp myapp
   COPY test.txt relativeDir/
   
   # accepts URL and extracts archives 
   # prefer COPY
   ADD myapp myapp
      
   CMD ["executable","param1","param2"]
   ENTRYPOINT ["executable", "param1", "param2"]   
   
   LABEL org.opencontainers.image.authors="ondrej.smola@tul.cz"
      
   ARG BUILD_DIR
   ARG BUILD_DIR="/opt/build"
   
   USER student
   USER student:students
   ```


