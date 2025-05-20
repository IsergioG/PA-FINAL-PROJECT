# 💀 Death Note

Aplicación basada en el anime DeathNote 

## 🧱 Tecnologías principales

- ⚙️ Backend: Go (`gorilla/mux`, GORM)
- 🌐 Frontend: React + Vite + TypeScript
- 🗃️ Base de datos: PostgreSQL
- 🐳 Contenedores: Docker & Docker Compose

## 📁 Estructura del Proyecto

## ⚙️ Requisitos previos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Opcional para desarrollo individual) Node.js y Go instalados

---

## 🚀 Cómo iniciar el proyecto

### 1. Clona el repositorio

```bash
git clone https://github.com/IsergioG/PA-FINAL-PROJECT.git
cd death-note-fullstack

```

## Configura las variables de entorno 


PORT=8000
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=root
POSTGRES_PASSWORD=root
POSTGRES_DB=death_note_db

## Levanta los servicios de docker
```bash

docker-compose up --build

```
