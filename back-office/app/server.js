var express = require("express");
var app = express();

var RootDir = '/Users/hy/work/go/src/github.com/hiromaily/go-questionnaires/back-office/app'
var ViewDir = RootDir + '/views';
var StaticDir = RootDir + '/statics'
var port = 8082;

app.use(express.static(StaticDir));

var server = app.listen(port, function(){
    console.log("Node.js is listening to PORT:" + server.address().port);
});


app.get("/admin/", function(req, res, next){
    //res.json(photoList);
    res.sendFile('index.html', { root: ViewDir});
});
