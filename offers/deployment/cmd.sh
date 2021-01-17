protoc -I pb/location \
pb/location/location.proto \
--go_out=plugins=grpc:.


//from weather app dir

docker create network 

docker image build -t devminnu/locationapp -f location/Dockerfile .
docker push devminnu/locationapp
docker container run -d --network net-location-app -p 33001:33001 --name location location  
docker container ls -a 
docker container rm -f location
docker container logs location

docker image rm location

//createmongo

docker run -d -p 2020:27017 --network net-location-app  \
    -e MONGO_INITDB_DATABASE=locationdb \
    -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
    -e MONGO_INITDB_ROOT_PASSWORD=mongoadminpass321 \
    --name locationdb mongo

docker container rm -f locationdb


db.createUser({user: 'mongoadmin', pwd: 'mongoadminpass321',roles: [{role: 'readWrite', db: 'locationdb'}]}); 


k run mongodb --image mongo --env=MONGO_INITDB_ROOT_USERNAME=mongoadmin --env=MONGO_INITDB_ROOT_PASSWORD=mongoadminpass321 --port 27017 --expose --dry-run=client -oyaml -- bash "mongod;db.createUser({user: 'mongoadmin', pwd: 'mongoadminpass321',roles: [{role: 'readWrite', db: 'locationdb'}]}); " > mongo.yaml 


# // db\.createUser\(\{user\: $DB_USER\,pwd\: $DB_PASS\,roles\:\[\{role\: "readWrite"\,db\: "locationdb"\}\]\}\);


docker container run -d -p 3306:3306 --network net-location-app \
-e MYSQL_ROOT_PASSWORD=admin \
-e MYSQL_DATABASE=offers \
-e MYSQL_USER=admin \
-e MYSQL_PASSWORD=admin \
--name mysql mysql

docker container rm -f mysql




# New offerspub
docker image build -t devminnu/offerspub .

docker push devminnu/offerspub

docker container run -d -p 30003:30003 --name offerspub devminnu/offerspub

docker container rm -f offerspub


# 
docker image build -t devminnu/offerssub .

docker push devminnu/offerssub

docker container run -d -p 30004:30004 --name offerssub devminnu/offerssub

docker container rm -f offerssub

# 
k create secret generic mysql-db-secret --from-env-file=mysql-db-secret.env --dry-run=client -oyaml > mysql-db-secret.yaml

k create secret generic rabbitmq-secret --from-env-file=rabbitmq-secret.env --dry-run=client -oyaml > rabbitmq-secret.yaml

k run mysql --image=mysql --port 3306 --expose --dry-run=client -oyaml > mysql.yaml

k run rabbitmq --image=rabbitmq:3-management --port 5672 --expose --dry-run=client -oyaml > rabbitmq.yaml

k create secret generic offerspub-secret --from-env-file=offerspub-secret.env --dry-run=client -oyaml > offerspub-secret.yaml
k create secret generic offerssub-secret --from-env-file=offerssub-secret.env --dry-run=client -oyaml > offerssub-secret.yaml

k create deployment offerspub --image=devminnu/offerspub --port 30003  --replicas 1 --dry-run=client -oyaml>offerspub-deployment.yaml
k expose deployment offerspub --port 30003 --target-port 30003 --type NodePort --dry-run=client -oyaml>offerspub-svc.yaml

k delete -f offerspub-secret.yaml -f offerspub-deployment.yaml -f offerspub-svc.yaml
k apply -f offerspub-secret.yaml -f offerspub-deployment.yaml -f offerspub-svc.yaml

k create deployment offerssub --image=devminnu/offerssub --port 30004  --replicas 1 --dry-run=client -oyaml>offerssub-deployment.yaml
# k expose deployment offerssub --port 30004 --target-port 30004 --type NodePort --dry-run=client -oyaml>offerssub-svc.yaml

k delete  -f offerssub-deployment.yaml #-f offerssub-svc.yaml
k apply  -f offerssub-deployment.yaml -f offerssub-svc.yaml

# docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management


