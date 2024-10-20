.PHONY: proto cert startServer startClient clean

cert:
	cd cert && ./generate_ca.sh && cd .. && \
	cd client && ./generate_keys.sh && cd .. && \
	cd server && ./generate_key.sh && cd ..

clean:
	cd cert && rm -f *.pem ; rm -f *.srl && cd .. && echo "Removed files in cert" && \
	cd server && rm -f *.pem ; rm -f *.srl && cd .. && echo "Removed files in server" && \
	cd client && rm -f *.pem ; rm -f *.srl && cd .. && echo "Removed files in client"

proto:
	cd proto && ./update.sh && cd ..

startServer:
	cd server && ./startserver.sh && cd ..

startClient:
	cd client && ./startclients.sh && cd ..