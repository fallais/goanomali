# goanomali

**goanomali** is a library written in Golang that helps you with the **Anomali REST API**.

## Usage

```go
import "github.com/fallais/goanomali"
```

Construct a new Anomali client, then use the various services on the client to
access different parts of the Anomali API. For example:

```go
client := qradar.NewClient(nil)
```

If you want to provide your own `http.Client`, you can do it :

```go
httpClient := &http.Client{}
client := qradar.NewClient(httpClient)
```
