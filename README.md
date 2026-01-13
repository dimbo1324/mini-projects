# 📂 Mini Projects Collection

A curated collection of mini-projects and coding challenges implemented in various programming languages. This repository serves as a playground for practicing algorithms, system design patterns, concurrency, and language-specific features.

## 🛠 Structure

The repository is organized by programming language:

* **`go_projects/`** - Projects written in Go (Golang).
* **`python_projects/`** - Projects written in Python.

---

## 🐹 Go Projects

Implementations focusing on concurrency patterns, data structures, and simulation.

| Project Name                                    | Description                                                                                                               | Key Concepts                                                            |
| :---------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------ | :---------------------------------------------------------------------- |
| **[Batcher Queue](./go_projects/BatcherQueue)** | A generic library for grouping items into batches based on time intervals or size limits.                                 | `Generics`, `Concurrency`, `Timeouts`, `Graceful Shutdown`              |
| **[DbSim (TypeBox)](./go_projects/DbSim)**      | An in-memory key-value database simulator that supports scalars, lists, and nested objects with SQL-like command parsing. | `In-memory Storage`, `Interface Implementation`, `Parsers`, `Recursion` |
| **[Async Pipeline](./go_projects/pipeline)**    | A multi-stage asynchronous data processing pipeline (Unix-pipe style) for handling user data and spam checks.             | `Channels`, `Worker Pools`, `Fan-out/Fan-in`, `Sync/Atomic`             |

### Quick Start (Go)

To run any of the Go projects, navigate to the specific directory and use standard Go commands.

**Example (Batcher Queue):**
```bash
cd go_projects/BatcherQueue/app
go test -v ./...