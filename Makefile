SHELL := /bin/bash


NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m

BINARY="bluebell"



# all 默认命令
all: gotool build


build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}


run:
	@echo ''
	@printf '$(OK_COLOR)快糙猛，跑一哈 .. 🚀$(NO_COLOR)\n'
	@ go run main.go --filename ./settings/config.yaml
	@echo '🎯'
	@echo ''
	


gotool:
	go fmt ./
	go vet ./


clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi


help:
	@echo ''
	@printf '$(OK_COLOR) 选项 🎯$(NO_COLOR)\n'
	@echo ''
	@echo "make - 格式化 Go 代码，并编译生成二进制文件"
	@echo "make build - 编译 Go 代码，生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件"
	@echo "make gotool - 运行 Go 工具 'fmt' 和 'vet' "
	@echo ''



container_open:
	@echo ''
	@printf '$(OK_COLOR)打开容器服务 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml up -d 
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''


container_close:
	@echo ''
	@printf '$(OK_COLOR)关闭容器服务 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml down 
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''


login_mysql:
	@echo ''
	@printf '$(OK_COLOR)登录 MySQL 容器 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml exec mysql-db sh -c 'mysql -uroot -p${MYSQL_ROOT_PASSWORD}'
	@printf '$(OK_COLOR)退出 .. 🎯$(NO_COLOR)\n'
	@echo ''



login_redis:
	@echo ''
	@printf '$(OK_COLOR)登录 Redis 容器 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml exec redis-cache sh -c 'redis-cli'
	@printf '$(OK_COLOR)退出 .. 🎯$(NO_COLOR)\n'
	@echo ''

 
container_detail:
	@echo ''
	@printf '$(OK_COLOR)查看容器配置 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml config
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''


container_net:
	@echo ''
	@printf '$(OK_COLOR)查看 MySQL 容器 IP 地址 .. 🚀$(NO_COLOR)\n'
	@docker inspect mysql-db | grep IPAddress
	@echo ''
	@printf '$(OK_COLOR)查看 Redis 容器 IP 地址 .. 🚀$(NO_COLOR)\n'
	@docker inspect redis-cache | grep IPAddress
	@echo ''
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''



