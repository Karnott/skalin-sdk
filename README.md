# Skalin SDK (golang)

## How to use it ?

```golang
  package main

  import (
    "github.com/karnott/skalin-sdk"
  )

  skalinApi, err := skalinsdk.New("GetSkalinAppClientID", "GetSkalinClientApiID", "GetSkalinClientApiSecret")
  if err != nil {
    panic(err)
  }
  customerId := "1"
  contact := Contact{
    RefId:     "2",
    Customer:  &customerId,
    LastName:  "Ceci est un test de l'API (nom de famille)",
    FirstName: "Ceci est un test de l'API (prénom)",
    Email:     "contact+testapi@karnott.fr",
    Phone:     "0123456789",
    Tags:      []string{"tag1", "tag2", "tag3"},
    CustomAttributes: CustomAttributes{
      "customAttributeId": "Ceci est un test de l'API (attribut personnalisé 2)",
    },
  }
  contactSaved, err := skalinApi.SaveContact(contact)
  if err != nil {
    panic(err)
  }
}
```

## About the test

Because an API SDK need to call real URLs, we add mock to simulate API response.
But to test the SDK calling the API, 8 ENV var can be set:
```bash
TEST_SKALIN_APP_CLIENT_ID=""
TEST_SKALIN_CLIENT_API_ID=""
TEST_SKALIN_CLIENT_API_SECRET=""
TEST_SKALIN_CONTACT_CUSTOM_ATTRIBUTE_ID=""
TEST_SKALIN_CUSTOMER_CUSTOM_ATTRIBUTE_ID=""
TEST_SKALIN_EXISTING_CUSTOMER_REF_ID=""
TEST_SKALIN_EXISTING_CUSTOMER_ID=""
TEST_SKALIN_EXISTING_CONTACT_ID=""
```
If these 8 env var are defined, the `GET`, `POST` and `PATCH` APIs will be test with data from Skalin API


## Env var

```bash
LOG_FORMAT=json # define the skalin sdk log format
LOG_LEVEL=info # define the skalin sdk log level
```
~
~