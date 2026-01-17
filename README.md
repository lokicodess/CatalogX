## CatalogX

**CatalogX** is a production-grade **Product Catalog REST API** built using **Go**, designed to deeply explore real-world backend engineering concepts.

The system supports **JWT-based authentication**, **product and category management**, **full-text search**, and **advanced filtering** with pagination and sorting.

To achieve high performance and scalability, **Redis** is used for **caching and rate limiting**, while **PostgreSQL** serves as the primary datastore. The application is fully containerized using **Docker** and instrumented with **Prometheus** and **Grafana** for observability, metrics, and performance monitoring.

CatalogX is engineered to:
- Handle **1000+ requests per second**
- Achieve **70%+ cache hit ratio**
- Maintain **sub-100ms p95 latency**
- Ensure reliability through **extensive unit and integration tests** with **80%+ coverage**
