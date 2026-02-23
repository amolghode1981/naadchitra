# Museum Collection Explorer

A web application for exploring museum collection APIs, built with a Go backend and React frontend.

## Project Structure

```
museum-project/
├── backend/                  # Go API server
│   ├── main.go               # HTTP server, routes, registry
│   ├── go.mod
│   └── connectors/
│       ├── connector.go      # Connector interface + ArtObject struct
│       └── met.go            # Metropolitan Museum of Art connector
└── frontend/                 # React + Vite + Tailwind
    ├── index.html
    ├── package.json
    ├── tailwind.config.js
    ├── postcss.config.js
    ├── vite.config.js
    └── src/
        ├── main.jsx
        ├── App.jsx
        ├── App.css
        └── components/
            ├── MuseumSearch.jsx
            ├── ArtworkCard.jsx
            └── ErrorMessage.jsx
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org) 18+ and npm

## Running Locally

You need two terminals running at the same time.

**Terminal 1 — Backend:**
```bash
cd backend
go run main.go
# Runs on http://localhost:8080
```

**Terminal 2 — Frontend:**
```bash
cd frontend
npm install
npm run dev
# Runs on http://localhost:5173
```

Then open `http://localhost:5173` in your browser.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/museums` | List available museum connectors |
| GET | `/api/object?museum=met&id=50939` | Fetch a single object by ID |

## Adding a New Museum Connector

1. Create a new file in `backend/connectors/` e.g. `rijks.go`
2. Implement the `Connector` interface:
```go
type Connector interface {
    Name() string
    FetchObject(id string) (*ArtObject, error)
}
```
3. Register it in `backend/main.go`:
```go
var registry = map[string]connectors.Connector{
    "met":   connectors.NewMetConnector(),
    "rijks": connectors.NewRijksConnector(), // add here
}
```

## Museums Currently Supported

- [Metropolitan Museum of Art](https://metmuseum.github.io/) — `met`
  - API documentation and full endpoint reference available at https://metmuseum.github.io

## Tech Stack

- **Backend:** Go, standard library only (no frameworks)
- **Frontend:** React 19, Vite, Tailwind CSS
- **Design:** Indian miniature painting palette — lapis, saffron, gold, ivory
- **Wireframes & Initial Design:** Basic wireframes and component structure designed using [Claude](https://claude.ai) by Anthropic
