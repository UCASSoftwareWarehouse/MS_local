//https://www.w3schools.com/nodejs/nodejs_mongodb.asp

var MongoClient = require('mongodb').MongoClient;
var url = "mongodb://localhost:27017/Files";
//create database
MongoClient.connect(url, function(err, db) {
    if (err) throw err;
    console.log("Database created!");
    db.close();
});

// create collection
var url1 = "mongodb://localhost:27017"
MongoClient.connect(url1, function(err, db) {
    if (err) throw err;
    var dbo = db.db("Files");
    // dbo.collection("code").drop(function(err, delOK) {
    //     if (err) throw err;
    //     if (delOK) console.log("Collection deleted");
    // });
    // dbo.collection("binary").drop(function(err, delOK) {
    //     if (err) throw err;
    //     if (delOK) console.log("Collection deleted");
    // });
    dbo.createCollection("code", function(err, res) {
        if (err) throw err;
        console.log("Collection code created!");
    });
    dbo.createCollection("binary", function(err, res) {
        if (err) throw err;
        console.log("Collection binary created!");
        db.close();
    });
});










