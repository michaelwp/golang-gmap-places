# Geocoding & Place API service:
![forthebadge made-with-python](https://forthebadge.com/images/badges/made-with-go.svg)

## Prerequisite 
Before using this api, please create your own API key and enable it for Geocoding API & Places API services (visit [here](https://developers.google.com/maps/documentation/geocoding/get-api-key))  

### List of available endpoints:
```API-BASE-URL: http://0.0.0.0:8093/api/v1```


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
         required: Required
      }`
  - `country: {
         type: String (ISO 3166-1 alpha-2),
         required: Optional
      }`
     
- Example request:
  - `/map?place=google&country=id`
  
- Response:
  - `status`: `200`
  - ```json
       {
           "code": 1,
           "message": "places",
           "data": [
               {
                   "keyword": "google",
                   "place_id": "ChIJ8YAu8BuTaS4RtlSUVSaGNnc",
                   "name": "Google ClassRoom",
                   "address": "#79A Umer Block Abbasia Town Main Road KOMP TNI AL BLOK E1 NO 13, RT.04/RW.20, Ciangsana, Kec. Gn. Putri, Bogor, Jawa Barat 16968",
                   "country": "id",
                   "lat": -6.3637314,
                   "lon": 106.9483217
               },
               {
                   "keyword": "google",
                   "place_id": "ChIJyd0fX0nxaS4RCnYrvC2HPf4",
                   "name": "Google Indonesia",
                   "address": "Pacific Century Place Tower Level 45 SCBD Lot 10, Jl. Jend. Sudirman No.53, RT.5/RW.3, Senayan, Kec. Kby. Baru, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12190",
                   "country": "id",
                   "lat": -6.2276334,
                   "lon": 106.8085406
               },
               {
                   "keyword": "google",
                   "place_id": "ChIJW1V7ne6NcC4RKYLMrI1pqTA",
                   "name": "Ratu Ayu Salon Spa Pucanggading",
                   "address": "Jl.Sarwo Edi Wibowo No.200 A-Raya, Pucanggading, Batursari, Kec. Mranggen, Kabupaten Demak, Jawa Tengah 51254",
                   "country": "id",
                   "lat": -7.030869499999999,
                   "lon": 110.4878426
               },
               {
                   "keyword": "google",
                   "place_id": "ChIJoU6jfurxaS4RxeKEpbqynto",
                   "name": "Google maps",
                   "address": "Jl. Palem 2, Petukangan Utara, Kec. Pesanggrahan, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12260",
                   "country": "id",
                   "lat": -6.225632,
                   "lon": 106.7579
               },
               {
                   "keyword": "google",
                   "place_id": "ChIJhYN0EkVeei4RGlXl65xR_AI",
                   "name": "Google Chicken",
                   "address": "Jl. Raya Turi, Bunder, Purwobinangun, Kec. Pakem, Kabupaten Sleman, Daerah Istimewa Yogyakarta 55582",
                   "country": "id",
                   "lat": -7.659457099999998,
                   "lon": 110.3937444
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