.PHONY: proto cert server client clean report

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

server:
	cd server && ./startserver.sh && cd ..

client:
	cd client && ./startclients.sh && cd ..

report:
	cd report && rm -f report.pdf && \
	pandoc -f markdown -t pdf report.md -o report.pdf && \
	cd ..