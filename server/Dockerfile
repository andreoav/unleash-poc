FROM node:slim

WORKDIR /app

ADD package.json yarn.lock ./

RUN yarn

ADD . .

CMD ["node", "index.js"]