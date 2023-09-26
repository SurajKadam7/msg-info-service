gen-srv:
	microgen -file=msginfo_srv/service.go -out=msginfo_srv/ -package=github.com/SurajKadam7/msg-info-service/msginfo_srv

gen-repo:
	microgen -file=repos/msginfo/repo.go -out=repos/msginfo/ -package=github.com/SurajKadam7/msg-info-service/repos/msginfo

# forward slash bug in -out=/xx