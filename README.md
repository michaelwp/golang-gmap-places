## gmap-service api documentation :
> Version: 1
- ```API-BASE-URL: http://0.0.0.0/api/v1```

-------

### List of available endpoints:

#### map
- `GET /map`

#### Error response format:
 - `status`: `4xx`,`5xx`
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
         required: Optional
      }`
     
- Example request:
  - `/map?place=kanggo`
  
- Response:
  - `status`: `200`
  - ```json
       {
           "code": 1,
           "message": "places",
           "data": [
               {
                   "keyword": "kanggo",
                   "place_id": "ChIJAfHFvc8GQi4RrT44xQGKR2E",
                   "name": "Kanggo Sampean",
                   "address": "Pasir Nangka, Tigaraksa Kecamatan, Tangerang, Banten 15720, Indonesia",
                   "lat": -6.237410499999999,
                   "lon": 106.4740201
               },
               {
                   "keyword": "kanggo",
                   "place_id": "ChIJH0MO-RH7aS4RcpWP5dXevfM",
                   "name": "Kanggo",
                   "address": "Ruko Foresta Business Loft 1, Jl. BSD Raya Utama No.32, BSD City, Kec. Pagedangan, Tangerang, Banten 15339, Indonesia",
                   "lat": -6.2975198,
                   "lon": 106.6408388
               },
               {
                   "keyword": "kanggo",
                   "place_id": "ChIJl4RJTZ72aS4RQZP86Tc-zI8",
                   "name": "Kanggo",
                   "address": "Jl. Kebon Kacang Raya No.G7, RW.1, Kb. Melati, Kecamatan Tanah Abang, Kota Jakarta Pusat, Daerah Khusus Ibukota Jakarta 10240, Indonesia",
                   "lat": -6.1945564,
                   "lon": 106.8157077
               }
           ]
       }
    ```
  - ```json
         {
             "code": 3,
             "message": "places not found",
             "data": null
         }
      ```