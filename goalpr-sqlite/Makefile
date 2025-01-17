.PHONY: run
dev:
	nodemon --exec go run main.go --signal SIGTERM
clean:
	rm public/images/*.jpg
clean_db:
	rm data/sqlite.db	
ffmpeg_serve:
	ffmpeg -i public/videos/1.mp4 -listen 1 -f mp4 -movflags frag_keyframe+empty_moov  http://127.0.0.1:5000
