gen_saas_retail:
	protoc --go_out=retail/protocol -I=retail/protocol --go_opt=paths=source_relative byteplus_saas_retail.proto