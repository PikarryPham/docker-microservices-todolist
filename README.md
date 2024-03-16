Instruction - MongoDB Cluster (Primary - Secondary - Secondary)
=========================================

---
## Mongo Components

* Config Server (3 member replica set): `configsvr01`,`configsvr02`,`configsvr03`
* 3 Shards (each a 3 member `PSS` replica set):
	* `shard01-a`,`shard01-b`, `shard01-c`
	* `shard02-a`,`shard02-b`, `shard02-c`
	* `shard03-a`,`shard03-b`, `shard03-c`
* 2 Routers (mongos): `router01`, `router02`

## ✨ Steps

### 👉 Step 1: Start all of the containers

```bash
docker-compose up -d
```

### 👉 Step 2: Initialize the replica sets (config servers and shards)

```bash
docker-compose exec configsvr01 sh -c "mongosh < /scripts/init-configserver.js"
docker-compose exec shard01-a sh -c "mongosh < /scripts/init-shard01.js"
docker-compose exec shard02-a sh -c "mongosh < /scripts/init-shard02.js"
docker-compose exec shard03-a sh -c "mongosh < /scripts/init-shard03.js"
```

### 👉 Step 3: Initializing the router

>Note: Wait a bit for the config server and shards to elect their primaries before initializing the router

```bash
docker-compose exec router01 sh -c "mongosh < /scripts/init-router.js"
```

### 👉 Step 4: Enable sharding and setup sharding-key
```bash
docker-compose exec router01 mongosh --port 27017

// Enable sharding for database `MyDatabase`
sh.enableSharding("MyDatabase")

// Setup shardingKey for collection `MyCollection`**
db.adminCommand( { shardCollection: "MyDatabase.MyCollection", key: { oemNumber: "hashed", zipCode: 1, supplierId: 1 } } )

exit
```

## 📋 Verify

### ✅ Verify the status of the sharded cluster

```bash
docker-compose exec router01 mongosh --port 27017
sh.status()
exit
```

### ✅ Verify status of replica set for each shard
> 1 PRIMARY, 2 SECONDARY

```bash
docker exec -it shard-01-node-a bash -c "echo 'rs.status()' | mongosh --port 27017" 
docker exec -it shard-02-node-a bash -c "echo 'rs.status()' | mongosh --port 27017" 
docker exec -it shard-03-node-a bash -c "echo 'rs.status()' | mongosh --port 27017"
```

### ✅ Check database status
```bash
docker-compose exec router01 mongosh --port 27017
use MyDatabase
db.stats()
db.MyCollection.getShardDistribution()
exit
```

## 🔎 More commands

```bash
docker exec -it mongo-config-01 bash -c "echo 'rs.status()' | mongosh --port 27017"
docker exec -it shard-01-node-a bash -c "echo 'rs.help()' | mongosh --port 27017"
docker exec -it shard-01-node-a bash -c "echo 'rs.status()' | mongosh --port 27017" 
docker exec -it shard-01-node-a bash -c "echo 'rs.printReplicationInfo()' | mongosh --port 27017" 
docker exec -it shard-01-node-a bash -c "echo 'rs.printSlaveReplicationInfo()' | mongosh --port 27017"
```
## 🔎 Overall systems

![alt text](https://ibb.co/9nPFKYz)
