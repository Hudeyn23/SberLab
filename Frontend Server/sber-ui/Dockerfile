FROM node:12.22.2

WORKDIR /app
COPY ["package.json", "package-lock.json*", "./"]
RUN npm install
COPY ./ .
RUN npm run build

CMD [ "npm", "run", "production-serve" ]