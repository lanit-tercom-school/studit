FROM nginx

RUN apt-get update && apt-get install --no-install-recommends --no-install-suggests -y git

RUN git clone https://github.com/lanit-tercom-school/studit

COPY nginx.conf /etc/nginx/nginx.conf

RUN cd webroot/src && ng build --prod

