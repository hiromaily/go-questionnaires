var express = require("express");
var app = express();
var debugMode = 0;
//debug mode
//debugMode = 1;

//Docker
var RootDir = '/usr/src/front-office/app'

if (debugMode && debugMode == 1){
    //Local
    RootDir = '/Users/hy/work/go/src/github.com/hiromaily/go-questionnaires/front-office/app'
}


var ViewDir = RootDir + '/views';
var StaticDir = RootDir + '/statics'
var port = 8081;

//app.use(express.static(StaticDir));
app.use('/', express.static(StaticDir));

var server = app.listen(port, function(){
    console.log("Node.js is listening to PORT:" + server.address().port);
});


app.get("/", function(req, res, next){
    if (debugMode && debugMode == 1) {
        res.sendFile('debug.html', { root: ViewDir});
    }else{
        //res.json(photoList);
        res.sendFile('index.html', { root: ViewDir});
    }
});
