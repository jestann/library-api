package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// main view and serve static assets
	// http.HandleFunc("/", homeHandler)
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	err := http.ListenAndServe(":"+os.Getenv(`PORT`), r)
	/* if err != nil {
	    log.Fatal("ListenAndServe: ", err)
	} */

	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	// Here we are implementing the NotImplemented handler. Whenever an API endpoint is hit
	// we will simply return the message "Not Implemented"
	var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Implemented"))
	})
}

/*
type Product struct {
    Id int
    Name string
    Slug string
    Description string
}

/* We will create our catalog of VR experiences and store them in a slice. */ /*
var products = []Product{
  Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description : "Shoot your way to the top on 14 different hoverboards"},
  Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description : "Explore the depths of the sea in this one of a kind underwater experience"},
  Product{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park", Description : "Go back 65 million years in the past and ride a T-Rex"},
  Product{Id: 4, Name: "Cars VR", Slug : "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
  Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description : "Pick up the bow and arrow and master the art of archery"},
  Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description : "Explore the seven wonders of the world in VR"
}

/* The status handler will be invoked when the user calls the /status route
   It will simply return a string with the message "API is up and running" */ /*
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("API is up and running"))
})

/* The products handler will be called when the user makes a GET request to the /products endpoint.
   This handler will return a list of products available for users to review */ /*
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  // Here we are converting the slice of products to json
  payload, _ := json.Marshal(products)

  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(payload))
})

/* The feedback handler will add either positive or negative feedback to the product
   We would normally save this data to the database - but for this demo we'll fake it
   so that as long as the request is successful and we can match a product to our catalog of products
   we'll return an OK status. */ /*
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  var product Product
  vars := mux.Vars(r)
  slug := vars["slug"]

  for _, p := range products {
    if p.Slug == slug {
        product = p
    }
  }

  w.Header().Set("Content-Type", "application/json")
  if product.Slug != "" {
    payload, _ := json.Marshal(product)
    w.Write([]byte(payload))
  } else {
    w.Write([]byte("Product Not Found"))
  }
})

*/

//
// webserver.go
//
// An example of a golang web server.
//
// Usage:
//
//   # run go server in the background
//   $ go run webserver &
//
//   While that's running, use a browser to visit a page.
//   It responds in one of several ways :
//
//  (0) for the URL /home it sends a home HTML page,
//      that runs an AJAX secondary get
//
//   (1) For URLS that start with /generic/
//       it sends some text/plain diagnostics.
//
//       URL: http://localhost:8097/generic/page?color=purple
//       browser (text/plain) :
//           FooWebHandler says ...
//             request.Method      'GET'
//             request.RequestURI  '/generic/page?color=purple'
//             request.URL.Path    '/generic/page'
//             request.Form        'map[color:[purple]]'
//             request.Cookies()   '[testcookiename=testcookievalue]'
//
//   (2) For URLs of the form /item/textstring,
//       it sends back a simplistic JSON response.
//       (In a real application, texstring could for example be
//       the name of an item, and the response could describe it.)
//
//       URL: http://localhost:8097/item/yellow
//       browser (application/json) :
//           {"name":"yellow", "what":"item"}
//
//   (3) Other pages give an error.
//
//       URL: http://localhost:8097/other/path
//       browser :
//           404 page not found
//
// Each visit sets a simple cookie, which
// can be accessed from the request after the 1st visit.
//
// For use in an AJAX setting, you should first decide on a way to
// encode requests for information or submission of data into the URL.
// A REST API would for example use GET and PUT along with URLs that
// put the information requested or sent in the path, like the
// /item/name example here. Or you could use form or data passed in
// the ?keyword=value part of the URL, though I think that's less
// clean. Then to pass the data back to the javascript at the client,
// JSON as shown in the /item/name example is a good choice.
//
// For a discussion of REST see
// en.wikipedia.org/wiki/Representational_state_transfer#Central_principle
//
// Go also has a 3rd party gorilla/mux package that looks interesting,
// setting up fancier ways to extract information from the URL and
// decide which function will respond to a request.  See
// http://www.gorillatoolkit.org/pkg/mux for its details.
//
// Docs and examples for this stuff can be found at
//   http://golang.org/pkg/net/http      particularly #Request
//   http://golang.org/pkg/net/url/#URL  what's in request.URL
//   https://devcharm.com/pages/8-golang-net-http-handlers
//   http://www.alexedwards.net/blog/a-recap-of-request-handling
//   http://blog.golang.org/json-and-go
//
// Jim Mahoney | cs.marlboro.edu | MIT License | March 2014
/*

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func SetMyCookie(response http.ResponseWriter) {
	// Add a simplistic cookie to the response.
	cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue"}
	http.SetCookie(response, &cookie)
}

// Respond to URLs of the form /generic/...
func GenericHandler(response http.ResponseWriter, request *http.Request) {

	// Set cookie and MIME type in the HTTP headers.
	SetMyCookie(response)
	response.Header().Set("Content-type", "text/plain")

	// Parse URL and POST data into the request.Form
	err := request.ParseForm()
	if err != nil {
		http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
	}

	// Send the text diagnostics to the client.
	fmt.Fprint(response, "FooWebHandler says ... \n")
	fmt.Fprintf(response, " request.Method     '%v'\n", request.Method)
	fmt.Fprintf(response, " request.RequestURI '%v'\n", request.RequestURI)
	fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.URL.Path)
	fmt.Fprintf(response, " request.Form       '%v'\n", request.Form)
	fmt.Fprintf(response, " request.Cookies()  '%v'\n", request.Cookies())
}

// Respond to the URL /home with an html home page
func HomeHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("home.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("home.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

// Respond to URLs of the form /item/...
func ItemHandler(response http.ResponseWriter, request *http.Request) {

	// Set cookie and MIME type in the HTTP headers.
	SetMyCookie(response)
	response.Header().Set("Content-type", "application/json")

	// Some sample data to be sent back to the client.
	data := map[string]string{"what": "item", "name": ""}

	// Was the URL of the form /item/name ?
	var itemURL = regexp.MustCompile(`^/item/(\w+)$`)
	var itemMatches = itemURL.FindStringSubmatch(request.URL.Path)
	// itemMatches is captured regex matches i.e. ["/item/which", "which"]
	if len(itemMatches) > 0 {
		// Yes, so send the JSON to the client.
		data["name"] = itemMatches[1]
		json_bytes, _ := json.Marshal(data)
		fmt.Fprintf(response, "%s\n", json_bytes)

	} else {
		// No, so send "page not found."
		http.Error(response, "404 page not found", 404)
	}
}

func main() {
	port := 8080
	portstring := strconv.Itoa(port)

	// Register request handlers for two URL patterns.
	// (The docs are unclear on what a 'pattern' is,
	// but seems be the start of the URL, ending in a /).
	// See gorilla/mux for a more powerful matching system.
	// Note that the "/" pattern matches all request URLs.
	mux := http.NewServeMux()
	mux.Handle("/home", http.HandlerFunc(HomeHandler))
	mux.Handle("/item/", http.HandlerFunc(ItemHandler))
	mux.Handle("/generic/", http.HandlerFunc(GenericHandler))

	// Start listing on a given port with these routes on this server.
	// (I think the server name can be set here too , i.e. "foo.org:8080")
	log.Print("Listening on port " + portstring + " ... ")
	err := http.ListenAndServe(":"+portstring, mux)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}


*/
