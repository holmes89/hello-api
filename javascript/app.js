
let express = require('express');
let logger = require('morgan');

let translateRouter = require('./routes/translation');
const repo = require('./repository/translation')

let app = express();

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.get('/translate/:word', translateRouter(repo));

module.exports = app;