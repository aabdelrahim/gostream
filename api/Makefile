regen-go-pb:
	protoc --go_out=plugins=grpc:./ ./songservice.proto

regen-python-pb:
	python -m grpc_tools.protoc -I./ --python_out=. --grpc_python_out=. ./songservice.proto