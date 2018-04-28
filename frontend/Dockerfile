# Stage 0, based on Node.js, to build and compile Angular
FROM node:8.6 as node
WORKDIR /app
COPY package.json /app/
RUN npm install
COPY ./ /app/
RUN npm run build -- --environment=prod --prod

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx
COPY --from=node /app/dist/ /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/nginx.conf