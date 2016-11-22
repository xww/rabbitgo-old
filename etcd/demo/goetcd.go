package demo

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func main() {

	client := etcd.NewClient([]string{"http://127.0.0.1:4001"})
	receiver := make(chan *etcd.Response)
	go client.Watch("/creds", 0, false, receiver, nil)
	for {
		resp, err := client.Get("creds", false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current creds: %s: %s\n", resp.Node.Key, resp.Node.Value)

		r := <-receiver
		log.Printf("Got updated creds: %s: %s\n", r.Node.Key, r.Node.Value)
	}

}
