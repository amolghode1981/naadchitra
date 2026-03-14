# Museum Collection Explorer

A web application for exploring museum collection APIs, built with a Go backend and React frontend.

## Project Structure

```
NAADCHITRA/
├── connectors/               # External service connectors/integrations
├── frontend/                 # Frontend application
├── knowledgerepo/            # Core knowledge repository module
│   ├── import/               # Data import logic
│   │   ├── csv/              # CSV source files for bulk import
│   │   └── import_csv.go     # CSV import handler
│   ├── types/                # Domain model definitions
│   │   ├── instrument.go     # Instrument type
│   │   ├── painting-type.go  # Painting category/type
│   │   ├── painting.go       # Painting type
│   │   └── performer.go      # Performer type
│   ├── schemas/              # JSON schemas for data validation
│   │   ├── instrument.json
│   │   ├── painting-type.json
│   │   ├── painting.json
│   │   └── performer.json
│   ├── utils/                # Shared utilities
│   │   ├── jsonhelper.go     # JSON parsing/serialization helpers
│   │   └── logger.go         # Logging utility
│   ├── go.mod                # Go module definition
│   └── go.sum                # Dependency lock file
├── main.go                   # Application entry point
├── README.md
└── TODO.md
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

## Coding conventions
- Type should be designed to represent a single object.
- Type names will be singular. Same typename will be used for the JSON schema.
- Hyphen should be used in the filenames.

## Technical notes
**AI assisted image analysis**
- Code interacts with AI code using GRPC. Details of what AI is suppose to do are still developing. Though one important thing it should do is use a vision model on the miniature paintings to find out details.
- AI systems so far, still do not objects from Indian miniature paintings with acceptable accuracy. Expecting them to recognize music making and musical instrument is yet to be achieved. Though with the limiated accuracy, the system can still be made effective by running the first iteration through the vision models and later submitted for expert review.
- Project that facilitates this is called Kosambi, https://github.com/amolghode1981/kosambi
- How it does that, what models are used are concealed from NaadChitra.
- It is accessed using gRPC. IDLs and protofiles are a part of project meghdoot https://github.com/amolghode1981/meghdoot
- meghdoot project is added as a submodule to this project.
- Use `git clone --recurse-submodules git@github.com:amolghode1981/naadchitra.git` to also pull the files for the submodule.
- Use `make rpc t=go` to generate go stubs from the proto files. This would create a directory called `meghdoot/output`.
- After this rpc client and data could be imported for use.
