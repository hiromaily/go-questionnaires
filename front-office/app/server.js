var express = require("express");
var app = express();

//Local
//var RootDir = '/Users/hy/work/go/src/github.com/hiromaily/go-questionnaires/back-office/app'

//Docker
var RootDir = '/usr/src/front-office/app'

var ViewDir = RootDir + '/views';
var StaticDir = RootDir + '/statics'
var port = 8081;

//app.use(express.static(StaticDir));
app.use('/', express.static(StaticDir));

var server = app.listen(port, function(){
    console.log("Node.js is listening to PORT:" + server.address().port);
});


app.get("/", function(req, res, next){
    //res.json(photoList);
    res.sendFile('index.html', { root: ViewDir});
});
