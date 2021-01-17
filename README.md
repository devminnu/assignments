# assignments

1.Make sure mysql is running at port 3306 and DB offers is created with below credentails

MYSQL_ADDON_USER: "admin"
MYSQL_ADDON_PASSWORD: "admin"

If not then launch docker container for mysql db

docker container run -d -p 3306:3306 --network net-location-app \
-e MYSQL_ROOT_PASSWORD=admin \
-e MYSQL_DATABASE=offers \
-e MYSQL_USER=admin \
-e MYSQL_PASSWORD=admin \
--name mysql mysql

2.Make sure rabbitmq is running on default port with default user and password
if not then create rabbitmq container

docker run -it -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

3.Go to assignments/offers/offerspub and run below command
go run main.go

4.Go to assignments/offers/offerssub and run below command
go run main.go

5.Now send post request using below json
Method:POST
URL:http://localhost:30003/publishOffers
Body:{
"offers": [
{
"cm_offer_id": "8f6995366e854c9faf1d9f3d233702b8",
"hotel": {
"hotel_id": "BH~46456",
"name": "Hawthorn Suites by Wyndham Eagle CO",
"country": "US",
"address": "0315 Chambers Avenue, 81631",
"latitude": 39.660193,
"longitude": -106.824123,
"telephone": "+1-970-3283000",
"amenities": [
"Business Centre",
"Fitness Room/Gym",
"Pet Friendly",
"Disabled Access",
"Air Conditioned",
"Free WIFI",
"Elevator / Lift",
"Parking"
],
"description": "Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this h",
"room_count": 1,
"currency": "USD"
},
"room": {
"hotel_id": "BH~46456",
"room_id": "S2Q",
"description": "JUNIOR SUITES WITH 2 QUEEN BEDS",
"name": "JUNIOR SUITES WITH 2 QUEEN BEDS",
"capacity": {
"max_adults": 2,
"extra_children": 2
}
},
"rate_plan": {
"hotel_id": "BH~46456",
"rate_plan_id": "BAR",
"cancellation_policy": [
{
"type": "Free cancellation",
"expires_days_before": 2
}
],
"name": "BEST AVAILABLE RATE",
"other_conditions": [
"CXL BY 2 DAYS PRIOR TO ARRIVAL-FEE 1 NIGHT 2 DAYS PRIOR TO ARRIVAL",
"BEST AVAILABLE RATE"
],
"meal_plan": "Room only"
},
"original_data": {
"GuaranteePolicy": {
"Required": true
}
},
"capacity": {
"max_adults": 2,
"extra_children": 2
},
"number": 1,
"price": 1520,
"currency": "USD",
"check_in": "2020-11-18",
"check_out": "2020-11-20",
"fees": [
{
"type": "CountyTax",
"description": "COUNTY TAX PER STAY",
"included": true,
"percent": 17.5
}
]
}
]
}

6.Check data in mysql db
7.connect mysql db using below credentails.

MYSQL_ROOT_PASSWORD=admin
MYSQL_DATABASE=offers
MYSQL_USER=admin
MYSQL_PASSWORD=admin

# Deployment using local kubernetes cluster

if you have docker for desktop installed and local kubernetes cluster such as minikube running execute below command to deploy apps

1. Go to assignments/offers/deployment

2. Run below command to create all deployment resources
   kubectl apply \
   -f mysql-db-secret.yaml \
   -f rabbitmq-secret.yaml \
   -f offerspub-secret.yaml \
   -f offerssub-secret.yaml \
   -f mysql.yaml \
   -f rabbitmq.yaml \
   -f offerspub-deployment.yaml \
   -f offerspub-svc.yaml \
   -f offerssub-deployment.yaml

3. Delete all resources later
   kubectl delete \
   -f mysql-db-secret.yaml \
   -f rabbitmq-secret.yaml \
   -f offerspub-secret.yaml \
   -f offerssub-secret.yaml \
   -f mysql.yaml \
   -f rabbitmq.yaml \
   -f offerspub-deployment.yaml \
   -f offerspub-svc.yaml \
   -f offerssub-deployment.yaml
