.PHONY: run
clean: 
	sudo  rm -rf goalpr/assets/images
	sudo  mkdir goalpr/assets/images

start:
	docker-compose up -d 

stop:
	docker-compose down 

update:
	docker-compose down 
	docker-compose pull
	docker-compose up -d --build

restart:
	docker-compose restart

remove:
	docker-compose down -v
	docker-compose rm -f

post:
	bash -xv scrippts/post.sh
