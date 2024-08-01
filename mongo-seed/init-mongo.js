db = db.getSiblingDB('database');

db.createCollection('users');
db.users.insertMany([
    { name: 'Nguyen Van A', age: 20 },
    { name: 'Nguyen Van B', age: 21 },
    { name: 'Nguyen Van C', age: 22 }
]);