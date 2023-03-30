make servers:
	go run main.go

# Remove all files from ./cert
cert-delete-certificates:
	rm ./cert/*

# Create a new root CA key and crt. Fill your own values when prompted, or just hit enter and accept the defaults if you're just playing 
cert-gen-root-ca:
	openssl req -newkey rsa:2048 -nodes -x509 -days 365 -out cert/ca.crt -keyout cert/ca.key

# Create 3 keys, one for each server
cert-gen-server-keys:
	openssl genrsa -out cert/alpha.key 2048
	openssl genrsa -out cert/bravo.key 2048
	openssl genrsa -out cert/charlie.key 2048

# Create 3 CSRs. For this demo, the CN matter, therefore enter localhost when prompted.
cert-gen-server-CSRs:
	openssl req -new -key cert/alpha.key -days 365 -out cert/alpha.csr 
	openssl req -new -key cert/bravo.key -days 365 -out cert/bravo.csr 
	openssl req -new -key cert/charlie.key -days 365 -out cert/charlie.csr 

# Sign the CSRs with the CA:
cert-sign-CSRs:
	openssl x509 -req -in cert/alpha.csr -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -out cert/alpha.crt -days 365 -sha256
	openssl x509 -req -in cert/bravo.csr -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -out cert/bravo.crt -days 365 -sha256
	openssl x509 -req -in cert/charlie.csr -CA cert/ca.crt -CAkey cert/ca.key -CAcreateserial -out cert/charlie.crt -days 365 -sha256

certs: cert-delete-certificates cert-gen-root-ca cert-gen-server-keys cert-gen-server-CSRs cert-sign-CSRs


.PHONY: servers cert-delete-certificates cert-gen-root-ca cert-gen-server-keys cert-gen-server-CSRs cert-sign-CSRs certs