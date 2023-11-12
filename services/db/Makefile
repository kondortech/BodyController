.PHONY: spin-up-db clean

spin-up-db:
	sudo docker run -d -p 27017:27017 --name body-controller-mongo-db mongo:latest 

clean:
	sudo docker kill body-controller-mongo-db
	sudo docker container prune