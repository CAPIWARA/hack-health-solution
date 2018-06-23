package main

import (
	"net/http"
	"context"
	"github.com/gorilla/mux"
	"runtime"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"hack-health-solution/server/graphqlschema"
	"github.com/kr/pretty"
)

func main() {
	router := mux.NewRouter()

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &graphqlschema.Schema,
	})
	//router.HandleFunc("/login", loginAuth)
	router.Handle("/graphql", requireAuth(handler))
	pretty.Log("Server started at http://localhost:8080/graphql")
	pretty.Log(http.ListenAndServe(":8080", router))
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			pretty.Log("options method")
			return
		}
		//token := r.Header.Get("Authorization")
		//log.Println("token: ", token)
		//res, err := users.Decode(token)
		//		//if err != nil {
		//		//	log.Printf("Permission denied: %v", err)
		//		//	w.WriteHeader(http.StatusForbidden)
		//		//	return
		//		//}
		//		//if res.Id == "" {
		//		//	w.WriteHeader(http.StatusInternalServerError)
		//		//	return
		//		//}
		ctx := context.WithValue(r.Context(), "email", "1")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
