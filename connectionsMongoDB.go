package anton

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Get the Client, will pass this Client into our other functions
func GetClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func DisconnectClient(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func GatherMasters(client *mongo.Client, filter bson.M) []Master {
	var masters []Master
	collection := client.Database("Anton").Collection("Masters")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var master Master
		err = cur.Decode(&master)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		masters = append(masters, master)
	}

	// Close the Cursor
	cur.Close(context.TODO())

	// Return Results
	return masters

}

func GatherSlaves(client *mongo.Client, filter bson.M) []Slave {
	var slaves []Slave
	collection := client.Database("Anton").Collection("Slaves")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var slave Slave
		err = cur.Decode(&slave)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		slaves = append(slaves, slave)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSlaves] Failed to Close Connection", err)
	}

	// Return Results
	return slaves
}

func GatherProfiles(client *mongo.Client, filter bson.M) []Profile {
	var profiles []Profile
	collection := client.Database("Anton").Collection("Profiles")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var profile Profile
		err = cur.Decode(&profile)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		profiles = append(profiles, profile)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherProfiles] Failed to Close Connection", err)
	}

	// Return Results
	return profiles

}

func GatherProxies(client *mongo.Client, filter bson.M) []Proxy {
	var proxies []Proxy
	collection := client.Database("Anton").Collection("Proxies")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var proxy Proxy
		err = cur.Decode(&proxy)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		proxy.ProxyAddress = "http://" + proxy.ProxyAddress
		proxies = append(proxies, proxy)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherProfiles] Failed to Close Connection", err)
	}

	// Return Results
	return proxies
}

func GatherSiteDictionaries(client *mongo.Client, filter bson.M) []SiteDictionary {
	var siteDictionaries []SiteDictionary
	collection := client.Database("Anton").Collection("SitesDictionary")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var siteDictionary SiteDictionary
		err = cur.Decode(&siteDictionary)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		siteDictionaries = append(siteDictionaries, siteDictionary)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSiteDictionaries] Failed to Close Connection", err)
	}

	// Return Results
	return siteDictionaries

}

func GatherFrontEndUsers(client *mongo.Client, filter bson.M) []FrontEndUser {
	var antonUsers []FrontEndUser
	collection := client.Database("Anton").Collection("FrontEndUsers")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var antonUser FrontEndUser
		err = cur.Decode(&antonUser)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		antonUsers = append(antonUsers, antonUser)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSiteDictionaries] Failed to Close Connection", err)
	}

	// Return Results
	return antonUsers

}

func GatherMasterLines(client *mongo.Client, filter bson.M) []Lines {
	var lines []Lines
	collection := client.Database("Anton").Collection("MastersLines")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var line Lines
		err = cur.Decode(&line)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		lines = append(lines, line)
	}

	// Close the Cursor
	cur.Close(context.TODO())

	// Return Results
	return lines

}

func GatherSlaveLines(client *mongo.Client, filter bson.M) []Lines {
	var lines []Lines
	collection := client.Database("Anton").Collection("SlavesLines")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var line Lines
		err = cur.Decode(&line)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		lines = append(lines, line)
	}

	// Close the Cursor
	cur.Close(context.TODO())

	// Return Results
	return lines

}

// Master Method to push all MasterLines to MongoDB
func (master Master) PushMasterLines(URI string) {

	// Start database connections
	client := GetClient(URI)
	results := client.Database("Anton").Collection("MastersLines")

	formattedLinesToInsert := []interface{}{master.MasterLines}
	results.InsertMany(context.Background(), formattedLinesToInsert)
	DisconnectClient(client)

}
