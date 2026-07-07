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


# .dockerignore
means "do not send this files when building the image"



Docker Workflow Checklist
Phase 1 — Build
✅ docker build -t ascii-art-web . (Done)
Phase 2 — Inspect Images
docker images
docker image ls
Phase 3 — Create & Run Containers
docker run
docker ps
docker ps -a
Phase 4 — Inspect Objects
docker inspect <container>
docker inspect <image>
Phase 5 — Container Interaction
docker logs
docker exec
docker stop
docker start
docker restart
Phase 6 — Cleanup
docker rm
docker rmi
Phase 7 — Metadata & Information
docker history
docker image inspect
docker container inspect
Phase 8 — Networking (Commands)
docker port
docker network ls
docker network inspect
Phase 9 — Volumes (Introduction)
docker volume ls
docker volume inspect