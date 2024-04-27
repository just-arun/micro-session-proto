sessionProto:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=. --go-grpc_out=. session.proto
mailingProto:
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --go_out=./mailing --go-grpc_out=./mailing ./mailing/mailing.proto
sessionProtoJava:
	sh javabuild.sh
	export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --java_out=. --grpc-java_out=. content.proto





