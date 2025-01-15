.PHONY: run
dev:
	pymon app.py
clean: 
	rm static/images/*.jpg
ffmpeg_serve:
	ffmpeg -i  videos/alprVideo1.mp4 -listen 1 -f mp4 -movflags frag_keyframe+empty_moov  http://127.0.0.1:5001 -stream_loop -9
clean:
	rm images/*.jpg
