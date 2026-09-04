# SurgeQueue

> **High-Throughput Traffic Burst & Atomic Transaction Engine in Go**

SurgeQueue is a production-grade, high-concurrency transaction and traffic buffer engine built with **Go Clean Architecture**. Designed to withstand extreme traffic spikes (flash sales, ticket wars) without overselling inventory, crashing database pools, or dropping transactions.

---

## Architectural Highlights

- **Zero-Overselling Concurrency Protection**: Utilizes PostgreSQL Pessimistic Locking (`SELECT ... FOR UPDATE`) with a mathematical **Lock Ordering Pattern** to guarantee 100% ACID atomicity and prevent deadlocks under reverse concurrent transfers.
- **Asynchronous Traffic Buffer (Worker Pool)**: Protects relational database pools by buffering bursts into buffered Go channels and worker goroutines.
- **Polyglot Persistence**:
  - **PostgreSQL**: Strict relational ACID transactional data (Users, Wallets, Balances).
  - **MongoDB Atlas**: High-throughput append-only document telemetry for audit logs without locking the hot path.
- **Clean Architecture & 12-Factor App**: Decoupled domain layers, zero leaky abstractions with Context-Based `TxManager`, centralized environment configuration, and versioned database migrations.

---

## Tech Stack

- **Language**: Go 1.24
- **Relational Database**: PostgreSQL (GORM + Versioned Migrations)
- **Document Store**: MongoDB Atlas (Official Mongo Go Driver)
- **Mocking & Testing**: Mockery v2, Testify
- **Architecture**: Domain-Driven Modular Clean Architecture

---

## Getting Started

### 1. Prerequisites
- Go 1.24+
- PostgreSQL
- MongoDB Atlas / Local

### 2. Environment Setup
Copy the example environment file:
\`\`\`bash
cp .env.example .env
\`\`\`
Fill in your database credentials in \`.env\`.

### 3. Database Migrations
Run the versioned migrations:
\`\`\`bash
make migrate-up
\`\`\`

### 4. Run the Server
\`\`\`bash
go run main.go
\`\`\`
