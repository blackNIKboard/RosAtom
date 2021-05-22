# node.js build
FROM  node:15.10.0-alpine as build
WORKDIR /app

COPY package*.json ./
RUN npm ci --only-production

COPY . .
RUN npm run build

# nginx setup
FROM nginx:1.20.0-alpine
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /var/www/html
EXPOSE 80
