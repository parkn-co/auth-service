#!/bin/bash
set -e -v

rsync -avzh --exclude-from '.gitignore' . parkn-staging:/home/ubuntu/parkn
rsync -avzh ./vendor/ parkn-staging:/home/ubuntu/parkn/vendor/
scp .env parkn-staging:/home/ubuntu/parkn/.env

ssh parkn-staging 'bash -s' <<'ENDSSH'
  docker rm -f `docker ps -a -q`
  docker rmi -f `docker images -q`
  cd parkn
  ./build.sh
  ./run.sh
ENDSSH
