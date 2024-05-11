.PHONY: run
clean: 
	rm -rf goalpr/assets/images
	mkdir goalpr/assets/images

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

ffmpeg_serve:
	ffmpeg -i flaskalpr/static/uploaded-video/alprVideo.mp4 -listen 1 -f mp4 -movflags frag_keyframe+empty_moov  http://127.0.0.1:5001

ffmpeg_comp:
	 ffmpeg -i alpr-video.mp4 -vcodec libx265 -crf 40 alpr-video-small.mp4