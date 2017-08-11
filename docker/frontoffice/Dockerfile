FROM node:6.6

RUN mkdir -p /usr/src/front-office
RUN mkdir /var/log/questionnaire/

#COPY docker-entrypoint.sh /docker-entrypoint.sh

WORKDIR /usr/src/front-office

EXPOSE 8081
CMD ["node", "./app/server.js"]
