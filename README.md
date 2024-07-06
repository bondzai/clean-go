# clean-go

## cURL Commands for Testing Endpoints

### Create a New Event
```bash
curl --location 'http://localhost:9000/api/v1/events' \
--header 'Content-Type: application/json' \
--data '{
  "ID": "1",
  "Message": "This is a new event"
}'
```

### Get All Events
```bash
curl --location 'http://localhost:9000/api/v1/events'
```

### Get Event by ID
```bash
curl --location 'http://localhost:9000/api/v1/events/1'
```

### Delete Event by ID
```bash
curl --location --request DELETE 'http://localhost:9000/api/v1/events/1'
```

### Update Event by ID
```bash
curl --location --request PUT 'http://localhost:9000/api/v1/events/1' \
--header 'Content-Type: application/json' \
--data '{
  "Message": "Updated Event Message"
}'
```
