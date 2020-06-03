FROM nginx 
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY wait-for /bin/wait-for
RUN apt-get update && apt-get install -y netcat
ENTRYPOINT wait-for --timeout=300 backend:80 -- nginx -g 'daemon off;'