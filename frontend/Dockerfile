FROM node:22.15.0 as build

WORKDIR /app

COPY package*.json ./

RUN yarn install

RUN yarn global add @angular/cli

COPY . .

RUN ng build client

FROM nginx:latest

COPY --from=build /app/dist/client/browser /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
