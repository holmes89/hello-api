
let logger = require('morgan');
const redis = require('redis');

class Repository {

    host = process.env.DB_HOST || 'localhost'
    port = process.env.DB_PORT || '6379'
    defaultLanguage = process.env.DEFAULT_LANGUAGE || 'english'
    client = undefined
    constructor(host, port){
        this.host = host ? host : this.host
        this.port = port ? port : this.port
        const connectionURL = `redis://${this.host}:${this.port}`
        console.log(`connecting to ${connectionURL}`)
        this.client = redis.createClient({url : connectionURL});
        this.client.on('connect', ()=> {
            console.log('connected to redis')
        })
        this.client.on('error', err => console.log('client error', err));
        this.client.connect()
    }

     async translate(language, word){
        const lang = language ? language.toLowerCase() : defaultLanguage.toLowerCase()
        const key = `${word.toLowerCase()}:${lang}`
        const val = await this.client.get(key)
        return val
    }
    async close(){
        this.client.quit()
    }
}

module.exports = {Repository};