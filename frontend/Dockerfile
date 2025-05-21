FROM node:20-alpine

WORKDIR /app

# ติดตั้ง dependencies ล่วงหน้า
COPY package*.json ./
RUN npm install

# คัดลอกโค้ด
COPY . .


EXPOSE 3000

CMD ["npm", "run", "dev"]
