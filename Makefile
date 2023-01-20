GOCMD=go
GOBUILD=${GOCMD} build -mod=mod
GOCLEAN=${GOCMD} clean

build: accountbff

.PHONY: \
    accountbff accountbffprod accountbffprod balancesvc balancesvcprod balancetask balancetaskprod paysvc paysvcprod paytask paytaskprod balacksvc balacksvcprod

clean:
	${GOCLEAN}

accountbff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account github.com/zxq97/account/cmd/accountbff

accountbffprod:
	${GOBUILD} -o /home/work/run/account github.com/zxq97/account/cmd/accountbff

balancesvc:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account_balancesvc github.com/zxq97/account/cmd/balancesvc

balancesvcprod:
	${GOBUILD} -o /home/work/run/account_balancesvc github.com/zxq97/account/cmd/balancesvc

balancetask:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account_balancetask github.com/zxq97/account/cmd/balancetask

balancetaskprod:
	${GOBUILD} -o /home/work/run/account_balancetask github.com/zxq97/account/cmd/balancetask

paysvc:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account_paysvc github.com/zxq97/account/cmd/paysvc

paysvcprod:
	${GOBUILD} -o /home/work/run/account_paysvc github.com/zxq97/account/cmd/paysvc

paytask:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account_paytask github.com/zxq97/account/cmd/paytask

paytaskprod:
	${GOBUILD} -o /home/work/run/account_paytask github.com/zxq97/account/cmd/paytask

usersvc:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account_usersvc github.com/zxq97/account/cmd/usersvc

usersvcprod:
	${GOBUILD} -o /home/work/run/account_usersvc github.com/zxq97/account/cmd/usersvc

blacksvc:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account_blacksvc github.com/zxq97/account/cmd/blacksvc

blacksvcprod:
	${GOBUILD} -o /home/work/run/account_blacksvc github.com/zxq97/account/cmd/blacksvc