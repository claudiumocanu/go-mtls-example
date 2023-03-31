# go-mtls-example

## Goal

The goal is to cover as many scenarios as possible:

- alpha listens on plain HTTP
- bravo listens on HTTPS without mTLS enabled
- charlie listens on HTTPS with mTLS enabled, trusting only alpha

## step-by-step

I developed this in small increments, starting from the basics. Each milestone has a reference to the commit.


### 1. Three simple HTTP servers
Start with 3 regular HTTP servers, each listening on a different port.

- `make servers` runs main.go which laches the 3 routines, each listening on its own port on plain HTTP
- all three can be accessed via the browser at `http://localhost:20001/hello`, `http://localhost:20002/hello`, `http://localhost:20003/hello`.


### 2. Enable HTTPs

#### 2.1 Generate the certificates
First, we need to generate some certificates.
Check `Makefile` and run each command individually for a more detailed view of what is happening under the hood, or just run `make certs`. This command will:
- delete all files from ./cert
- create a CA
- create 3 keys, one for each server
- create CSRs, one for each server
- sign the 3 CSRs with the CA

IMPORTANT: 
- when prompted for the CSRs details, make sure you fill a correct value when asked for the CN.
- the directory `./cert` is added in `.gitignore`: you don't want to push sensitive info in git. 

If everything worked fine, the following files should be generated under the `./cert` directory:

```
├── cert
│   ├── alpha.crt
│   ├── alpha.csr
│   ├── alpha.key
│   ├── bravo.crt
│   ├── bravo.csr
│   ├── bravo.key
│   ├── ca.crt
│   ├── ca.key
│   ├── charlie.crt
│   ├── charlie.csr
│   └── charlie.key
```

#### 2.2 Listen TLS

- update the code in bravo and charlie to start those two service on HTTPs. TODO: reference commit here
- start the servers: `make servers`
- access them again from the browser: to access bravo or charlie,`https://` must be explicitly typed in the browser hence they're not listening on the default `443`.
- of course, the browser will complain, because the certificates presented by the services are not signed by a trusted authority; these were self-signed with a CA that was self-generated in the previous step. The self-generated CA can be trusted in the operating system or the browser if needed.


