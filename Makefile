.PHONY: proto cert startServer startClient clean

cert:
	cd cert && \
	./genkey.sh && \
	cd ..

clean:
	cd cert && \
	rm *.pem ; rm *.srl && \
	cd ..

proto:
	cd proto && ./update.sh && cd ..

startServer:
	cd server && ./startserver.sh && cd ..

startClient:
	cd client && ./startclients.sh && cd ..