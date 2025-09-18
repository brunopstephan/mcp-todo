# MCP to-do list

This project presents a simple to-do list operation in a MCP ecosystem (Server and Client).

<img width="1757" height="450" alt="image" src="https://github.com/user-attachments/assets/a25a9ad4-2c85-4655-ad2b-2eeb977eb9f2" />

## Technologies

### MCP Server:
- Golang;
- GORM (Postgres);
- [MCP go-sdk](https://github.com/modelcontextprotocol/go-sdk);

### MCP Client:
- NodeJS;
- Typescript;
- [mcp-use-ts](https://github.com/mcp-use/mcp-use-ts);
- Anthropic's Claude API as LLM (Model: claude-3-5-haiku-20241022);

## Setup

First of all, you need create a `.env` file with the variables in `.env.example`.

You can use the Postgres credentials in the example running `docker compose up -d` to start a Postgres container.

### MCP Server

- Enter the  `mcp-server` directory:

```bash
cd mcp-server
```

- Install dependencies:

```bash
go mod download
```

- Build the binary:

```bash
go build -o build/mcp-server
```

### MCP Client

- Enter the  `mcp-client` directory:

```bash
cd mcp-client
```

- Install dependencies:

```bash
npm install
```

- Build the executable .js file:

```bash
npm run build
```

- Run it:

```bash
npm start
```
