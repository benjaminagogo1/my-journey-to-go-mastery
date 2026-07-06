docker --name ascii-web ascii-art-web

docker rm
docker run -p 8080:8080 ascii-art-web
docker run -d -p 8080:8080 ascii-art-web (Run the container in the background)
docker logs ascii-web (views the logs)
docker ps -a (view both running and stopped)
docker exec
docker exec <container name> command
docker -it ascii-web sh 
docker images (To see the images total numbers of images inside the  container)


LABEL
LABEL author ="Benjamin Agogo"

docker inspect 

docker inspect ascii-art-web
