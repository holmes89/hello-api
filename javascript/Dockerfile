FROM node:17

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY package*.json ./

RUN npm ci --only=production

# Bundle app source
COPY . ./
ENV PORT 8080
ENV NODE_ENV production
EXPOSE 8080
CMD [ "node", "./bin/www" ]