package e001

import "net/http"

func main()  {

	http.HandleFunc("/", homepage)
	http.ListenAndServe(":3000", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello.........."))
}