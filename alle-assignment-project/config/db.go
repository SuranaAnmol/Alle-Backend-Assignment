package config

import (
    "context"
    "fmt"
    "log"
    "sync"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    DB     *mongo.Database
    client *mongo.Client
    once   sync.Once
)

func ConnectDB() {
    once.Do(func() { // Ensures this runs only once
        fmt.Println("🔄 Connecting to MongoDB...")

        clientOptions := options.Client().ApplyURI("mongodb+srv://suranaanmol999as:vCbMy5RivHgHQUyt@cluster0.1854f.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

        var err error
        client, err = mongo.Connect(context.TODO(), clientOptions)
        if err != nil {
            log.Fatal("❌ Error connecting to MongoDB:", err)
        }

        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        err = client.Ping(ctx, nil)
        if err != nil {
            log.Fatal("❌ MongoDB connection failed:", err)
        }

        DB = client.Database("taskDB")
        fmt.Println("✅ Successfully connected to MongoDB!")
    })
}

func GetDB() *mongo.Database {
    if DB == nil {
        log.Fatal("🚨 Database is not initialized. Call ConnectDB() first.")
    }
    return DB
}
