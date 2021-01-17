docker image build -t devminnu/offerspub .

docker container run -d -p 30003:30003 --name offerspub devminnu/offerspub

docker container rm -f offerspub

docker push devminnu/offerspub