const gulp = require("gulp");
const browserSync = require("browser-sync");
const dartSass = require("sass");
const gulpSass = require("gulp-sass");
const sass = gulpSass(dartSass);
const autoprefixer = require("gulp-autoprefixer");

const { createGulpEsbuild } = require("gulp-esbuild");
const gulpEsbuild = createGulpEsbuild({
  incremental: false,
});
const isProd = process.env.NODE_ENV === '"prod"';

function buildScripts(done) {
  gulp
    .src("./src/*.+(js)")
    .on("error", console.log)
    .pipe(
      gulpEsbuild({
        bundle: true,
        platform: "node",
        outdir: "../dist/",
        color: true,
        minify: isProd,
        sourcemap: !isProd,
      })
    )
    .pipe(gulp.dest("../dist/"));
  done();
}

function buildStyles(done) {
  gulp.src(["./src/style.scss"]).pipe(sass().on("error", sass.logError)).pipe(autoprefixer()).pipe(gulp.dest("../dist"));
  done();
}

function build(done) {
  buildScripts(done);
  buildStyles(done);
}
exports.build = build;
