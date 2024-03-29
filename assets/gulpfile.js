const gulp = require("gulp");
const dartSass = require("sass");
const gulpSass = require("gulp-sass");
const sass = gulpSass(dartSass);
const autoprefixer = require("gulp-autoprefixer");

function buildStyles(done) {
  gulp
    .src(["./src/style.scss"])
    .pipe(sass().on("error", sass.logError))
    .pipe(autoprefixer())
    .pipe(gulp.dest("../dist"));
  done();
}

function build(done) {
  buildStyles(done);
}
exports.build = build;
