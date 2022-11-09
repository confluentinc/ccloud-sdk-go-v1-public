# ccloud-sdk-go-v1-public

> Public Golang SDK for Confluent Cloud.

## Usage

```go
import (
  "fmt"
  "context"

  "github.com/confluentinc/ccloud-sdk-go-v1-public"
  schedv1 "github.com/confluentinc/cc-structs/kafka/scheduler/v1"
)

const (
  baseURL = "https://confluent.cloud"
)

func main() {
  // create an unauthenticated client
  client := ccloud.NewClient(baseURL, nil, nil)
  token, err := client.Auth.Login("john@doe.io", "password")
  if err != nil {
    panic(err)
  }

  // create an authenticated client
  client := sdk.NewClientWithJWT(context.Background(), token, baseURL, nil)
  authConfig, err := client.Auth.User()
  if err != nil {
    panic(err)
  }

  apikey, err := client.APIKey.Create(&schedv1.ApiKey{
  	Description: "test key",
  	UserId:      authConfig.User.Id,
  	Clusters:    []string{"lkc-abc123"},
  })

  fmt.Println("key: %s\nsecret: %s",
    apikey.Key,
    apikey.Secret,
  )
}
```
