const express = require('express');
const app = express();

// Define a GET endpoint
app.get('/', (req, res) => {
    res.send('Hello, World!');
});

module.exports = app;