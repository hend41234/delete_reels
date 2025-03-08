* create file ".env"
  ```.env
  #GRAPH API
  BASE_URL_GRAPH_API="https://graph.facebook.com/v22.0"
  USER_ID="USER ID of main account"
  SECRET_APP="your SECRET APP"
  APP_ID="your APP ID"
  SYSTEM_USER_ACCESS="your system user TOKEN of developer account"
  ```
* run
    ```bash
    go build -ldflags="-s -w" -o delete_reels .
    ./delete_reels
    ```
  
