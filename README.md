# Financial Solution Backend

![Go Version](https://img.shields.io/badge/Go-1.25.4-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Podman](https://img.shields.io/badge/Podman-892CA0?style=for-the-badge&logo=podman&logoColor=white)

Welcome to the **Financial Solution Backend**. This repository contains the backend services for the Financial Solution platform, built with **Go** for high performance and scalability.

## 🚀 Getting Started

Follow these steps to get your development environment up and running.

### ⚙️ Configuration

Before running the application, create a `.env` file in the root directory to configure the database connection.

**Example `.env` file:**

```env
DB_URL=postgres://[YOUR-USERNAME]:[YOUR-PASSWORD]@localhost:5432/[YOUR-DATABASE]
```

### 🐳 Build & Run with Podman

We use Podman for containerization. You can build and run the environment with the following commands:

**1. Build the image**

VERSION: example 1.0.0

```bash
podman build -t finantial-solution-be:{VERSION} .
```

**2. Run the container**

This command runs the container interactively, maps the current directory to `/app`, and exposes port `8080`.

```bash
podman run -it --rm -p 8080:8080 -v .:/app:Z --userns=keep-id finantial-solution-be:{VERSION} /bin/sh
```

## 📂 Project Structure

- **`main.go`**: Application entry point.
- **`go.mod`**: Go module dependencies.

---

_Generated for Financial Solution Backend_
