
FROM node:20 AS builder
RUN mkdir -p /usr/src/app
RUN mkdir -p /usr/src/app/corte.frontend
# Crear el directorio de trabajo
WORKDIR /usr/src/app/corte.frontend

# Copiar los archivos necesarios
COPY . .

# Instalar dependencias
RUN npm install

EXPOSE 5173
CMD [ "npm", "run", "dev"]
