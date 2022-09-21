let express = require('express');
let logger = require('morgan');

let translateRouter = require('./routes/translation');
const { Repository } = require('./repository/translation');

class App {
	app = express();
	repo = undefined;
	constructor(host, port) {
		this.app.use(logger('dev'));
		this.app.use(express.json());
		this.app.use(express.urlencoded({ extended: false }));
		this.repo = new Repository(host, port);
		this.app.get('/translate/:word', translateRouter(this.repo));
	}

	async close() {
		return this.repo.close();
	}
}

module.exports = { App };
