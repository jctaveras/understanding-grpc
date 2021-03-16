SERVER_CN = localhost

generate-calc:
		protoc calculator/calcpb/calc.proto --go_out=plugins=grpc:.

generate-greet:
		protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

# Output files
# ca.key: Certificate Authority Private Key file
# ca.crt: Certificate Authority trust certificate
# server.key: Server Private Key password protected
# server.csr: Server Certificate Signed Request
# server.crt: Server Certificate Signed by CA
# server.pem: Conversion of server.key into a format for gRPC

# Sharable Files: ca.crt, server.csr
# Private Files: ca.key, server.key, server.crt, server.pem
gen-ssl:
		openssl genrsa -passout pass:welc0me3 -des3 -out ssl/ca.key 4096
		openssl req -passin pass:welc0me3 -new -x509 -days 365 -key ssl/ca.key -out ssl/ca.crt -subj "/CN=$(SERVER_CN)"

		openssl genrsa -passout pass:welc0me3 -des3 -out ssl/server.key 4096
		openssl req -passin pass:welc0me3 -new -key ssl/server.key -out ssl/server.csr -subj "/CN=$(SERVER_CN)" -config ssl/ssl.cnf

		openssl x509 -req -passin pass:welc0me3 -days 365 -in ssl/server.csr -CA ssl/ca.crt -CAkey ssl/ca.key -set_serial 01 -out ssl/server.crt -extensions req_ext -extfile ssl/ssl.cnf

		openssl pkcs8 -topk8 -nocrypt -passin pass:welc0me3 -in ssl/server.key -out ssl/server.pem
