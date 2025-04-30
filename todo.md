# 🧩 URL Shortener – System Design + Golang Feature Checklist

A practical roadmap to build and learn full-scale system design concepts using Go.

---

## 🔰 Phase 1: Core Functionality (MVP)
- [ ] Generate short URLs from long URLs
- [ ] Redirect from short URL to original URL
- [ ] Save mappings in in-memory store (e.g., `map[string]string`)
- [ ] Base62 or UUID-based slug generator

---

## 🛢️ Phase 2: Persistent Storage
- [ ] Use PostgreSQL / Redis / BadgerDB for persistent storage
- [ ] Design `urls` table: `id`, `short_url`, `original_url`, `created_at`, `expires_at`
- [ ] Add TTL support (URL expiry)

---

## 🧠 Phase 3: System Design Concepts
- [ ] Use layers: Handler ➝ Service ➝ Repository
- [ ] Separate internal packages: `/api`, `/models`, `/db`, `/utils`
- [ ] Clean Architecture / Hexagonal (optional)

---

## 🧰 Phase 4: Developer Tools
- [ ] Live reload with `Air`
- [ ] Environment config with `godotenv` or `viper`
- [ ] Logging (`logrus` or `zap`)
- [ ] Middleware (rate limiter, logging, recovery)

---

## 🚀 Phase 5: Advanced Features
- [ ] Rate limiting per IP/user (redis sliding window)
- [ ] Analytics: click count, referrer, location
- [ ] User accounts and auth (JWT)
- [ ] Custom aliases for short URLs
- [ ] Expiring links (set TTL at creation)
- [ ] Admin dashboard (optional)

---

## ⚙️ Phase 6: Scalability & Reliability
- [ ] Use Redis/Memcached as a cache layer
- [ ] Horizontal scaling discussion (load balancer)
- [ ] High availability DB config
- [ ] Sharding short URL space (hash-based)

---

## 🧪 Phase 7: Testing & CI
- [ ] Unit tests for handlers/services
- [ ] Integration tests
- [ ] GitHub Actions / CI pipeline

---

## 🐳 Phase 8: Docker & Deployment
- [ ] Dockerfile for Go app
- [ ] Docker Compose (Go + Postgres/Redis)
- [ ] Deploy to Render / Railway / GCP / AWS / DigitalOcean

---

## 🧹 Phase 9: Optional Polishing
- [ ] Custom 404 page
- [ ] QR code generation
- [ ] Preview page before redirect

---

## 🏁 Final Touch
- [ ] README with setup instructions
- [ ] Swagger/OpenAPI docs
- [ ] Postman collection
