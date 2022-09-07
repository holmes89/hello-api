const redis = require("redis");

class Repository {
  constructor(host, port) {
    this.host = host ? host : process.env.DB_HOST || "localhost";
    this.port = port ? port : process.env.DB_PORT || "6379";
    this.defaultLanguage = process.env.DEFAULT_LANGUAGE || "english";
    const connectionURL = `redis://${this.host}:${this.port}`;
    console.log(`connecting to ${connectionURL}`);
    this.client = redis.createClient({ url: connectionURL });
    this.client.on("connect", () => {
      console.log("connected to redis");
    });
    this.client.on("error", (err) => console.log("client error", err));
    this.client.connect();
  }

  async translate(language, word) {
    const lang = language
      ? language.toLowerCase()
      : this.defaultLanguage.toLowerCase();
    const key = `${word.toLowerCase()}:${lang}`;
    const val = await this.client.get(key);
    return val;
  }
  async close() {
    this.client.quit();
  }
}

module.exports = { Repository };
