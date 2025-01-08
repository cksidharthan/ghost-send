# Share Secret - Secure Secret Sharing Platform

A secure platform for sharing sensitive information with end-to-end encryption, automatic expiration, and view limits. Built with Nuxt.js 3 for the frontend and Go for the backend.

## 🌟 Features

- **End-to-End Encryption**: All secrets are encrypted before leaving your browser
- **Auto-Expiration**: Set expiration times (5 minutes to 7 days)
- **View Limits**: Control how many times your secret can be viewed
- **Password Protection**: Additional security layer with password protection
- **Modern UI**: Responsive design with dark mode and animations
- **Secure Backend**: Built with Go and PostgreSQL for robust data handling

## 🏗️ Architecture

### Frontend (Nuxt.js 3)
- Modern, responsive UI built with Nuxt.js 3
- TailwindCSS for styling
- Dark mode support
- Animated components
- Form validation
- Clipboard integration
- Error handling and notifications

### Backend (Go)
- RESTful API built with Gin framework
- PostgreSQL database with migrations
- Secure password hashing
- Automatic cleanup of expired secrets
- Health check endpoints
- Structured logging with Zap
- Dependency injection with Uber FX

## 🚀 Getting Started

### Prerequisites
- Node.js (v22.9.0 or later)
- Go (v1.23.4 or later)
- Docker and Docker Compose
- PostgreSQL (v17.0 or later)

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/share-secret.git
cd share-secret
```

2. Set up environment variables:
```env
# Frontend (.env)
SHARE_SECRET_API_URL=http://localhost:7780

# Backend (.env)
POSTGRES_HOST=localhost
POSTGRES_PORT=5433
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=secret
POSTGRES_SSL_MODE=disable
PORT=7780
```

3. Start the development environment:
```bash
# Using Docker Compose
docker-compose -f deploy/docker-compose.yaml up -d

# Or start services individually:

# Frontend
cd frontend
npm install
npm run dev

# Backend
cd backend
go mod download
go run main.go
```

The application will be available at:
- Frontend: http://localhost:9090
- Backend: http://localhost:7780
- PostgreSQL: localhost:5433

## 🔧 Configuration

### Frontend Configuration
- `nuxt.config.ts`: Nuxt.js configuration with the following modules:
  - `@nuxt/ui`
  - `@nuxtjs/tailwindcss`
  - `@nuxtjs/color-mode`
  - `@nuxt/icon`
  - `@nuxt/image`
- Environment variables:
  - `SHARE_SECRET_API_URL`: Backend API URL

### Backend Configuration
- Database migrations in `backend/db/migrations`
- Environment variables:
  - `POSTGRES_*`: Database configuration
  - `PORT`: API port
  - `LOG_LEVEL`: Logging level
  - `MIGRATION_PATH`: Path to database migrations

## 📦 Deployment

The project includes Docker configurations for easy deployment:

```bash
docker-compose -f deploy/docker-compose.yaml up -d
```

This will start:
1. PostgreSQL database (port 5433)
2. Backend API service (port 7780)
3. Frontend application (port 9090)

### Production Build

```bash
# Frontend
cd frontend
npm run build

# Backend
cd backend
go build -o secret main.go
```

## 🔒 Security Features

1. **Database Security**
   - Encrypted secret storage using PGP encryption
   - Automatic deletion of expired secrets
   - Password hashing with bcrypt
   - Secure salt generation

2. **API Security**
   - CORS protection
   - Input validation
   - Error handling
   - Secure password verification

3. **Frontend Security**
   - XSS protection
   - Secure password handling
   - Client-side validation
   - Secure clipboard operations

## 📝 API Endpoints

- `GET /healthz`: Health check endpoint
- `POST /secrets`: Create a new secret
  - Parameters:
    - `secret_text`: The text to encrypt
    - `password`: Access password
    - `expires_at`: Expiration time
    - `views`: Number of allowed views
- `GET /secrets/:id`: Retrieve a secret
  - Parameters:
    - `id`: Secret UUID
    - `password`: Access password
- `GET /secrets/:id/status`: Check secret status
  - Parameters:
    - `id`: Secret UUID

## 🛠️ Development

### Frontend Pages
- `/`: Home page with secret creation form
- `/access/:id`: Secret access page
- `/about`: About page with feature information

### Backend Structure
- `cmd/`: Application entry point
- `pkg/`: Main application packages
- `db/`: Database migrations and queries
- `deploy/`: Deployment configurations

## 📚 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Nuxt.js](https://nuxt.com/)
- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [TailwindCSS](https://tailwindcss.com/)
- [Gin](https://gin-gonic.com/)