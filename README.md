# httphelpers

Helpers for working with the go http standard library that I always end up needing

# Struct for an API Answer

Enforces a standard for the API answers from all services

```
type ApiAnswer struct
```

A successful answer will look like this:

```json
{
    "status": 200,
    "message": "Hello, world!",
    "data": {
        "hello": "world"
    }
}
```

An unsuccessful answer will look like this:

```json
{
    "status": 404,
    "message": "The resource was not found",
    "error": "RESOURCE_NOT_FOUND"
}
```

# Writing Successful / Unsuccessful answers

WriteSuccess and WriteError will write a successful or unsuccessful answer to the response writer. 

Note they can be used as http.HandlerFuncs.

```go
func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /write-success", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type AnyData struct {
			Hello string `json:"hello"`
		}

		anyData := AnyData{
			Hello: "world",
		}

		WriteSuccess(http.StatusOK, "Hello, world!", anyData)(w, r)
	}))

	mux.Handle("GET /write-error", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteError(http.StatusInternalServerError, "API_ERROR_CODE", "An error occurred")(w, r)
	}))

	_ = http.ListenAndServe("localhost:8000", mux)
}
```