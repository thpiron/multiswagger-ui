# Swagger UI

Display swaggers from multiple openAPI specs.

## Config:
```yaml
    {
      "name": "swagger",
      "favicon": "https://static1.smartbear.co/swagger/media/assets/swagger_fav.png",
      "openAPI": [
        {
          "name": "Pet",
          "URL": "https://petstore.swagger.io/v2/swagger.json"
        },
        ...
      ]
    }
```

## run as docker on port 8080
```bash
make docker-run
```