const { GenericContainer } = require('testcontainers');
const { App } = require('./app');
const request = require('supertest');

let container;
let app;
let api;

beforeAll(async () => {
  container = await new GenericContainer('redis')
    .withExposedPorts(6379)
    .withCopyFileToContainer('./data/dump.rdb', '/data/dump.rdb')
    .start();
  const port = container.getMappedPort(6379);
  const host = container.getHost();
  api = new App(host, port);
  app = api.app;
});

afterAll(async () => {
  await api.close();
  await container.stop();
});

describe('Translate', () => {
  test('hello translation in english to be hello', async () => {
    const response = await request(app).get('/translate/hello');
    expect(response.body).toEqual({
      translation: 'Hello',
      language: 'english',
    });
    expect(response.statusCode).toBe(200);
    return response;
  });
  test('hello translation in german to be hallo', async () => {
    const response = await request(app).get('/translate/hello?language=GERMAN');
    expect(response.body).toEqual({ translation: 'Hallo', language: 'german' });
    expect(response.statusCode).toBe(200);
    return response;
  });
});
