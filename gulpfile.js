var gulp = require('gulp')
    , concat = require('gulp-concat')
    , uglify = require('gulp-uglify')
    , ngAnnotate = require('gulp-ng-annotate');

gulp.task('default', function() {
  gulp.src(['public/js/**/*.js'])
  .pipe(concat('app.min.js'))
  .pipe(ngAnnotate())
  .pipe(uglify())
  .pipe(gulp.dest('public/js/'))
});
