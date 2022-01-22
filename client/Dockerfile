FROM node:17-alpine as build

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build

FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY --from=build /app/build .
