FROM node:12.22.2

WORKDIR /front
COPY ./dist ./dist
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install serve

CMD [ "npm", "run", "production-serve" ]