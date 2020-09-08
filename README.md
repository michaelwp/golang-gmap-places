## absen-service api documentation :
> Version: 1
- ```API-BASE-URL: http://147.139.182.234/api/v1```

-------

### List of available endpoints:

#### map
- `GET /map`

#### Error response format:
 - `status`: `4xx`
 - ```json 
   {
       "code": 0,
       "message": "...",
       "data": null
   }
   ```

#### GET /map
- Query Params :
  - `place: {
         type: String,
         required: false
      }`
     
- Example request:
  - `/map?place=kanggo`
  
- Response:
  - `status`: `200`
  - ```json
       {
           "code": 1,
           "message": "Places",
           "data": [
              "address 1",
              "address 2",
              "address 3"
            ]
       }
    ```