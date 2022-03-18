package main

func main() {

	//dbConfig := DB.Configure("./", "postgres")
	//DB.DB = dbConfig.InitConnection()
	server := NewServer(":3000")
	server.Handle("POST", "/create", PostRequest)
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/checkauth", server.AddMidleware(HandleRoot, CheckAuth()))
	server.Listen()

}
