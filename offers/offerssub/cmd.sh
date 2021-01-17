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

docker run -d --name rabbitmq-new -p 5672:5672 -p 15672:15672 --network net-location-app rabbitmq:3-management


