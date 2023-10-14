## Ports
`9901` : Envoy

`8080` : Envoy Proxy (Frontend hits this)

`9090` : Go Backend

## Generate Proto Files
```
bash gen-backend.sh
bash gen-frontend.sh
```

## Backend
```
cd backend-go
```

### Run server
```
go run .
```

## Envoy Proxy
### Install Envoy
```
docker pull envoyproxy envoy:dev
```

### Run Envoy
```
docker compose up
```

## Frontend
```
cd frontend-react-ts
```

### Install dependencies
```
npm i
```

### Run app
```
npm run dev
```