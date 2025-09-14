# ---------- Build Stage ----------
FROM node:22.15.0 AS build

WORKDIR /app

COPY ./frontend/package*.json ./

RUN yarn install --frozen-lockfile
RUN yarn global add @angular/cli

COPY ./frontend ./
RUN ng build client --configuration production


# ---------- Serve Stage ----------
FROM nginx:latest

COPY --from=build /app/dist/client/browser /usr/share/nginx/html
COPY ./frontend/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
