FROM node:lts-alpine

WORKDIR /app

COPY package*.json ./

# Install python and make (required for node-gyp)
RUN apk add --update python3 make g++

RUN npm install

COPY . .

RUN npm run build

RUN npm install -g http-server

EXPOSE 8080

ENV VITE_APP_BASE_URL="http://localhost:3000/api"

CMD [ "http-server", "dist", " --gzip", "--proxy", "http://localhost:8080?" ]