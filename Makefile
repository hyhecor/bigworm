init: \
	npm-install npm-update 

npm-install:
	npm install

npm-update:
	npm update

grpc-install:
	make install -f Makefile.grpc
	
grpc-genertate:
	PACKAGE=pkg/rpc/broker/v1 make genertate -f Makefile.grpc
	PACKAGE=pkg/rpc/broker/v2 make genertate -f Makefile.grpc

