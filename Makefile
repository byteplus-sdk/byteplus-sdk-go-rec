gen_saas_retail:
	protoc --go_out=retail/protocol -I=retail/protocol --go_opt=paths=source_relative byteplus_saas_retail.proto

gen_saas_content:
	protoc --go_out=content/protocol -I=content/protocol --go_opt=paths=source_relative byteplus_saas_content.proto