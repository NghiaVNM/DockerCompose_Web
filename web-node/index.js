const express = require('express');
const mongoose = require('mongoose');
const User = require('./models/User');

const app = express();
const port = 3000;

mongoose.connect('mongodb://mongo:27017/database', {
    useNewUrlParser: true,
    useUnifiedTopology: true
});

const db = mongoose.connection;
db.on('error', console.error.bind(console, 'connection error:'));
db.once('open', () => {
    console.log('Connected to MongoDB');
})

app.get('/', async (req, res) => {
    try {
        const users = await User.find();
        let userList = users.map(user => `<li>${user.name} - ${user.age}</li>`).join('');
        res.send(`
            <html>
                <head>
                    <title>NodeJS Web App</title>
                </head>
                <body>
                    <h1>Hello World from NodeJS</h1>
                    <h2>User List</h2>
                    <ul>
                        ${userList}
                    </ul>
                </body>
            </html>
        `);
    } catch (err) {
        res.status(500).send(err);
    }
});

app.listen(port, () => {
    console.log(`Server is running at port ${port}`);
});