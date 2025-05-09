# develop stage
FROM node:24-alpine as develop-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

# build stage
FROM develop-stage as build-stage
RUN npm run build

# production stage
FROM nginx:alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY entrypoint.sh /usr/share/nginx/
ENTRYPOINT ["/usr/share/nginx/entrypoint.sh"]
EXPOSE 90
CMD ["nginx", "-g", "daemon off;"]