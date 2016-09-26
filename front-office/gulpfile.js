var gulp = require('gulp');
var browserify = require('browserify');
var babelify = require('babelify');
var uglify = require('gulp-uglify');
var source = require('vinyl-source-stream');
var buffer = require('vinyl-buffer');
//var webserver = require('gulp-webserver');
//var beautify = require('gulp-beautify');
var server = require('gulp-express');

var inDir = './app/src/'
var outDir = './app/statics/dist/'
var inFiles = ['index'];

//browserify for release version
gulp.task('release', function() {
  process.env.NODE_ENV = 'production';
  inFiles.forEach(function(file,i,ar){
    browserify(inDir+file+".js", { debug: false }) //debug: true is for sourcemap
      .transform(babelify)
      .transform('browserify-shim', { global: true })
      .bundle()
      .on("error", function (err) { console.log("Error : " + err.message); })
      .pipe(source(outDir+file+".bundle.js"))
      .pipe(buffer())
      .pipe(uglify())
      .pipe(gulp.dest('./'))
  });
});

//browserify for rdev version
gulp.task('dev', function() {
  inFiles.forEach(function(file,i,ar){
    browserify(inDir+file+".js", { debug: true }) //debug: true is for sourcemap
      .transform(babelify)
      .transform('browserify-shim', { global: true })
      .bundle()
      .on("error", function (err) { console.log("Error : " + err.message); })
      .pipe(source(outDir+file+".bundle.js"))
      .pipe(gulp.dest('./'))
  });
});

//watch
gulp.task('watch', function() {
   gulp.watch('**/*.jsx', ['dev'])
});


//webserver express
gulp.task('web', function() {
  server.run(['app/server.js']);
});

gulp.task('do', ['dev', 'watch', 'web']);
gulp.task('default', ['release', 'watch', 'web']);
