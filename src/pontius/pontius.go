package pontius

type Application struct{
    server http
}

func (app Application) New(server http) {
    app.server = server;

    http.HandleFunc("/", )
    http.HandleFunc("/ping", ping)
    http.ListenAndServe(":8080", nil)
}

func (app Application03)run() {
    app.http.HandleFunc("/", )
    app.http.HandleFunc("/ping", ping)
    app.http.ListenAndServe(":8080", nil)
}