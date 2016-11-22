package main

import (
	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {
	cfg := client.Config{
		Endpoints:               []string{"http://10.0.2.15:2280"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	etcdClient, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}
	api := client.NewKeysAPI(etcdClient)

	//server
	watcher := api.Watcher("workers/", &client.WatcherOptions{
		Recursive: true,
	})
	go func() {
		for {
			res, err := watcher.Next(context.Background())
			if err != nil {
				log.Println("Error watch workers:", err)
				break
			}
			if res.Action == "expire" {
				fmt.Println(res.Action, res.Node.Key, res.Node.Value)

			} else if res.Action == "set" {
				fmt.Println(res.Action, res.Node.Key, res.Node.Value)

			} else if res.Action == "delete" {
				fmt.Println(res.Action, res.Node.Key, res.Node.Value)

			}
		}
	}()

	//client
	time.Sleep(time.Second * 1)
	key := "workers/" + "aa"
	value := "aa"
	_, err = api.Set(context.Background(), key, string(value), &client.SetOptions{
		TTL: time.Second * 30,
	})
	time.Sleep(time.Second * 3)
	key2 := "workers/" + "bb"
	value2 := "bb"
	_, err = api.Set(context.Background(), key2, string(value2), &client.SetOptions{
		TTL: time.Second * 30,
	})
	if err != nil {
		log.Println("Error update workerInfo:", err)
	}
	time.Sleep(time.Second * 1)

	res, _ := api.Get(context.Background(), "workers/", &client.GetOptions{Recursive: true})
	for i := 0; i < len(res.Node.Nodes); i++ {
		//log.Printf("%s,%s,%s",res.Action,res.Node.Nodes[i].Key,res.Action,res.Node.Nodes[i].Value)
		fmt.Println(res.Node.Nodes[i].Value)
	}
	time.Sleep(time.Second * 32)

}
