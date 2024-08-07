


init-docker:
	@echo '拉取镜像'
	docker pull bitnami/mysql:latest
	docker pull bitnami/redis:latest
	docker pull bitnami/etcd:latest


deploy-mysql:
	docker run -itd \
		--name db \
		-p 13306:3306 \
		-e ALLOW_EMPTY_PASSWORD=yes \
		-e MYSQL_ROOT_PASSWORD=my-secret-pw \
		bitnami/mysql:latest

deploy-redis:
    docker run -itd \
		--name cache \
		-p 16379:6379 \
		-e ALLOW_EMPTY_PASSWORD=yes \
		bitnami/redis:latest

deploy-etcd:    
	docker run -itd \
        --name etcd-standalone \
        -p 12379:2379 \
        -p 12380:2380 \
        -e ETCDCTL_API=3 \
        -e ALLOW_NONE_AUTHENTICATION=yes \
        -e ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
        bitnami/etcd:latest

exec-mysql:
	@echo '登录 MySQL 控制台'
    docker exec -it db sh

# 连接 mysql 服务器
# mysql --host=127.0.0.1 --port=3306 --user=root --password=my-secret-pw


exec-redis:
	@echo '登录 Redis 控制台'
    docker exec -it cache sh

# 连接 redis 服务器
# redis-cli


exec-etcd:
	@echo '登录 etcd 控制台'
	docker exec -it etcd-standalone sh

# 查看是否已经注册
# etcdctl get --prefix user.rpc

# 查看租约
# etcdctl lease list

# 查看租约剩余时间
# etcdctl lease timetolive xxx
