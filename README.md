# gset-crdt

GSet CRDT Cluster implemented in Go & Docker

## Introduction

CRDTs (Commutative Replicated Data Types) are a certain form of data types that when replicated across several nodes over a network achieve eventual consistency without the need of a consensus round. GSets abbreviated as grow-only sets are CRDT sets modified to only add data into it and becomes consistent across nodes in a cluster having replicated the set.

## Example

After building a cluster of GSet nodes we can now write values to either one or many nodes in the cluster.

```
$ curl -i -X POST localhost:8080/gset/append/user1
$ curl -i -X POST localhost:8081/gset/append/user2
```

When reading the list of values in the set they then sync up with each other and thus return consistent values everytime from any node in the cluster

```
$ curl -i -X GET localhost:8081/gset/list
{
    set: [
        user1,
        user2
    ]
}
```

The values remain consistent for nodes in the cluster that have never added any values into it

```
$ curl -i -X GET localhost:8082/gset/list
{
    set: [
        user1,
        user2
    ]
}
```

We can also lookup if certain values are present in the set

```
$ curl -i -X GET localhost:8082/gset/lookup/user1
> 200 OK
$ curl -i -X GET localhost:8082/gset/lookup/user3
> 404 NOT FOUND
```

## Steps

After cloning the repo. To provision the cluster:

```
$ make provision
```

This creates a 3 node GSet cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

Now we can send requests to append, list and lookup values to any peer node using its port allocated.

```
$ curl -i -X POST localhost:<peer-port>/gset/append/<value>
$ curl -i -X GET localhost:<peer-port>/gset/lookup/<value>
$ curl -i -X GET localhost:<peer-port>/gset/list
```

In the logs for each peer docker container, we can see the logs of the peer nodes getting in sync during read operations.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

This is not certain to clean up all the locally created docker images at times. You can do a docker rmi to delete them.

## GSets

CRDTs aim to be an alternative for consensus based distributed algorithms. Algorithms such as Paxos, Raft, etc tend to be more intensive on the network during their consensus phase and are a hassle when nodes in the cluster are very far apart from each other eg., WAN networks.

GSets are a simple example of CRDTs that are used to implement sets that can only add data & not remove data.

This is useful in certain scenarios like a unique site visitor list that saves a list of unique site visitors as it only increases and values cannot be removed.

To implement this we can build a cluster of GSet nodes of an arbitrary size and a proxy in front of it to direct traffic to the nodes. When a we want to store the ID of a new visitor we can send a write request to any node in the cluster or for the sake of availability, duplicate and send it to multiple nodes in the cluster. A property of CRDTs, immutability ensures that duplicate values are discarded in the set using an union operation. The nodes can either sync up with each other using operation-based replication or state-based replication and thus all of them get in sync and become eventually consistent.

## References

- [A comprehensive study ofConvergent and Commutative Replicated Data Types](https://hal.inria.fr/inria-00555588/document) [Marc Shapiro et al]
- [Strong Eventual Consistency and Conflict-free Replicated Data Types](https://www.youtube.com/watch?v=oyUHd894w18&t=3902s) [Marc Shapiro]
