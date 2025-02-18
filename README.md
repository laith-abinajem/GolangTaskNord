# **Multi-Tenant POS Transaction Processor**

## ** About the Project**

This project is a **Multi-Tenant POS (Point of Sale) Transaction Processor** built using **Golang (Fiber framework)** and **React.js (Material-UI)** for the frontend.  
It is designed to:
Process **sales transactions** across multiple tenants.  
 Maintain **product, branch, and tenant records**.  
 Provide **Prometheus monitoring** and **observability metrics**.  
 Offer **graceful shutdown & resilience** with retry mechanisms.  
 Use **Docker & Makefile commands** for easy setup and deployment.

---

## ** Features**

### **Backend (Golang + Fiber)**

- **Multi-Tenant Support**: Different tenants, products, and branches.
- **Transaction Processing**: API endpoints for creating and fetching sales transactions.
- **Caching & Observability**: Uses **Redis** for caching & **Prometheus** for monitoring.
- **Resilience & Graceful Shutdown**: Handles **SIGINT (Ctrl+C)** and **database reconnections**.
- **Role-Based Access Control (RBAC)**: Implements roles for **admin** and **staff**.
- **Database Seeding**: Automatically seeds **tenants, branches, products**.

### **Frontend (React + Material-UI)**

- **Create Transactions**: Allows users to create new sales transactions.
- **View Top Products**: Displays the best-selling products in a table.
- **Monitor Total Sales**: Shows **total sales per product** in a dashboard.

---

## ** Setup & Installation**

### **Prerequisites**

Before running the project, ensure you have:

- **Golang** (v1.18+)
- **Docker & Docker Compose**
- **Make** (for running commands)

---

## ** Quick Start - Running the Project**

### **Prerequisites**

Before running the project, ensure you have:

- **Golang** (v1.18+)
- **Docker & Docker Compose**
- **Make** (for running commands)

---

### **Set Up Environment Variables**

Load the environment variables from `exports.sh`:

```sh
source exports.sh
```

---

### ** Start Database & Services (MySQL, Redis)**

Run the following command to start **MySQL, Redis, and other services**:

```sh
make up
```

- This will:
  Start **MySQL** and **Redis** inside Docker  
   Recreate containers if necessary  
   Remove orphaned Docker containers

🔹 **Alternative: Run Manually**

```sh
docker compose up -d --build
```

---

### ** Seed Database**

To seed **tenants, branches, products**, and roles:

```sh
make seed
```

- This will:
  Run `cmd/seed/main.go` to insert **default data**.

🔹 **Alternative: Run Manually**

```sh
go run cmd/seed/main.go
```

---

### **Start the Backend**

To build and start the backend server:

```sh
make run
```

- This will:
  Compile the backend  
   Start the server (`bin/server`)  
   Load environment variables (`exports.sh`)

🔹 **Alternative: Run Manually**

```sh
go run main.go
```

### **Start the Frontend**

To run the **React frontend**:

```sh
cd webTask
yarn install   # Install dependencies
yarn start     # Start frontend on localhost:3000
```

- **Open in browser**: [http://localhost:3000](http://localhost:3000)

---

## ** API Endpoints**

| Endpoint                                   | Method | Description                       |
| ------------------------------------------ | ------ | --------------------------------- |
| `/api/transactions`                        | `POST` | **Create a new transaction**      |
| `/api/transactions/{tenantID}/{productID}` | `GET`  | **Get total sales for a product** |
| `/api/transactions/top-products`           | `GET`  | **Get top-selling products**      |

🔹 **Example Transaction Payload (`POST /api/transactions`)**

```json
{
  "tenant_id": 1,
  "branch_id": 1,
  "product_id": 1,
  "quantity": 3,
  "price_per_unit": 11
}
```

---

## ** Monitoring & Observability**

**Prometheus Metrics Endpoint:**

```sh
http://localhost:8888/metrics
```

It provides:

- **API Request Duration**
- **Transactions Processed**
- **Cache Hit/Miss Rates**
- **Go Runtime Metrics**

---
