# nginx setup
FROM nginx:1.20.0-alpine
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY frontend/ /var/www/html
EXPOSE 80
