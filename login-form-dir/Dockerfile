FROM node:18-alpine
RUN apk update && apk add git

WORKDIR /usr/src/app

COPY . .
RUN npm install
RUN npm run build
CMD ["npm", "run", "dev"]