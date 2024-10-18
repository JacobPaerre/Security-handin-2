.PHONY: update cert

cert:
	cd cert && ./genkeys.sh && cd ..

clean:
	cd cert && rm *.pem && rm *.srl && cd ..

update:
	./update.sh