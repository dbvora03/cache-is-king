# Cache is king

## About
I'm a huge fan of how MongoDB atlas works. All you have to do is get a connection string and you are basically connected to a cloud database. Ever since I started using Redis in my side projects, it annoyed me that I had to spin up a Redis server all the time. On top of that, deployments become a pain when dealing with another local server that needs to be run. The purpose of this project is to create a caching server that allows me to simply make requests to the service to cache my responses instead of retrieving them from a database. I recently found out about RedisLabs, which is a cloud solution for Redis, so I thought I'd give it a try. Lastly, I started learning Golang and I kinda like coding in this language, so I thought spinning up a quick REST API helps me learn this language a lot better. 

## Set up
1. Create a .env file with the following variables: `REDISHOST`, `REDISPASSWORD`, `REDISDB`
2. You can use your local Redis connection, or you can get one from RedisLabs or any other cloud provider
3. Run the following commands:
```
go mod init

go mod tidy

go run main.go
```
> I think that's it

And there you go, a caching server that you can use for your projects. 

## Usage
To send something to the caching server, send a POST request to `http://<hostname>/setCache/perm` with the following JSON format:
```json
{
  "project": "<your-project-name>",
  "unitid":"<id-of-your-unit>",
  "value":"<whatever-value-you-go>"
}
```

To send something to the caching server with an expiry time, send a POST request to `http://<hostname>/setCache/temp` with the following JSON format:
```json
{
  "project": "<your-project-name>",
  "unitid":"<id-of-your-unit>",
  "value":"<whatever-value-you-go>",
  "duration": <time-in-integer>
}
```

To retrieve something from the cache, send a GET request to `http://<hostname>/getCache` with the following JSON format:
```json
{
  "project": "<your-project-name>",
  "unitid":"<id-of-your-unit>",
}
```

*IMPORTANT* If passing in an object as a value, stringify the object. DO NOT use an object, or else it will not work