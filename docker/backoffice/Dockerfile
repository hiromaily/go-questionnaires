FROM node:6.6

RUN mkdir -p /usr/src/back-office
RUN mkdir /var/log/questionnaire/

#COPY docker-entrypoint.sh /docker-entrypoint.sh

WORKDIR /usr/src/back-office

EXPOSE 8082
CMD ["node", "./app/server.js"]
