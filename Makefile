gen-srv:
	microgen -file=msginfo_srv/service.go -out=msginfo_srv/ -package=github.com/suraj.kadam7/msg-info-srv/msginfo_srv

gen-repo:
	microgen -file=repos/msginfo/repo.go -out=repos/msginfo/ -package=github.com/suraj.kadam7/msg-info-srv/repos/msginfo

# forward slash bug in -out=/xx