FROM node:14.18-alpine3.12
EXPOSE 8080

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . ./

CMD ["npm", "run", "serve"]