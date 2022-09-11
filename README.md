# lets-go
Example and exercises from the the Book Let's Go

## Generate self-signed TLS

Create `tls` folder in the project's root:

```bash
mkdir tls && cd tls
```

Generate private key and TLS certificate:

```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost 
```