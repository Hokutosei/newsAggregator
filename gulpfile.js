var gulp = require('gulp')
    , concat = require('gulp-concat')
    , uglify = require('gulp-uglify')
    , ngAnnotate = require('gulp-ng-annotate');

var jsDirPath = 'public/js/**/*.js';

gulp.task('default', function() {
  gulp.src([jsDirPath])
  .pipe(concat('app.min.js'))
  .pipe(ngAnnotate())
  .pipe(uglify())
  .pipe(gulp.dest('public/js/'))
});

// var watcher = gulp.watch(jsDirPath, ['default']);
// watcher.on('change', function(event) {
//   console.log('File ' + event.path + ' was ' + event.type + ', running tasks...');
// });
