# My Go Server

A Go web server with MongoDB Atlas integration and internationalization support.

## 🚀 Quick Start

### Prerequisites
- Go 1.25+
- MongoDB Atlas account (or local MongoDB)

### Environment Setup
Create a `.env` file:
```env
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/database?retryWrites=true&w=majority
```

### Development with Live Reloading

Use the `dev` command for automatic server restarts on file changes:

```bash
make dev
```

This uses [Air](https://github.com/air-verse/air) for live reloading. The server will automatically restart when you modify `.go`, `.html`, `.tpl`, or `.tmpl` files.

### Manual Commands

```bash
# Run server once
make run

# Build binary
make build

# Run tests
make test
```

## 📁 Project Structure

```
├── main.go              # Application entry point
├── repositories/        # Data access layer
├── services/           # Business logic
├── routers/            # HTTP routes and handlers
├── clinics/            # Clinic models
├── users/              # User models
├── locales/            # Internationalization files
├── docs/               # API documentation
├── .air.toml           # Air configuration
└── Makefile            # Build commands
```

## 🌐 Internationalization (i18n)

The API supports multiple languages via the `Accept-Language` header:

```bash
# English
curl -H "Accept-Language: en" http://localhost:8080/users

# Spanish
curl -H "Accept-Language: es" http://localhost:8080/users
```

See `I18N_README.md` for more details.

## 🛠 Development

### Adding New Features
1. Update models in respective directories
2. Add repository methods
3. Implement service logic
4. Create/update routes
5. Add translations to `locales/` files

### Live Reloading Setup
The `.air.toml` file configures which files to watch and how to rebuild. By default, it watches:
- `*.go` files
- Template files (`.html`, `.tpl`, `.tmpl`)

## 📚 API Documentation

Visit `/swagger/index.html` when the server is running for interactive API documentation.