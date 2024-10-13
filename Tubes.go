package main

func main() {
	//// Serve src files from the "src" directory
	//fs := http.FileServer(http.Dir("./client/dist"))
	////http.Handle("/static/", http.StripPrefix("/src/", fs))
	//
	//http.Handle("/", fs)
	//
	//fmt.Println("Starting server at port 8080")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatalf("Error starting server: %v", err)
	//}
	Route()
}
