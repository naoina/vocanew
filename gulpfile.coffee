gulp = require 'gulp'
gp = require('gulp-load-plugins')()
browserify = require 'browserify'
source = require 'vinyl-source-stream'
buffer = require 'vinyl-buffer'
mainBowerFiles = require 'main-bower-files'

ASSETSDIR = './app/assets'
CSSDIR = "#{ASSETSDIR}/css"
JSDIR = "#{ASSETSDIR}/js"
IMAGEDIR = "#{ASSETSDIR}/image"
DESTDIR = './public/assets'
TMPDIR = './tmp'

isDist = gp.util.env.dist?

gulp.task 'css', ->
  css = mainBowerFiles(filter: /\.s?css$/).concat [
    "#{CSSDIR}/bootstrap_override.scss"
    "#{CSSDIR}/**/*.scss"
  ]
  gulp
    .src(css)
    .pipe(gp.cached 'css')
    .pipe(gp.if not isDist, do gp.sourcemaps.init)
    .pipe(do gp.plumber)
    .pipe(gp.sass imagePath: '/assets')
    .pipe(gp.autoprefixer 'last 2 version')
    .pipe(gp.remember 'css')
    .pipe(gp.concat 'app.css')
    .pipe(gp.if not isDist, do gp.sourcemaps.write)
    .pipe(gp.if isDist, do gp.minifyCss)
    .pipe(gulp.dest DESTDIR)

gulp.task 'scripts', ->
  browserify
    entries: ["#{JSDIR}/main.coffee"]
    extensions: ['.coffee']
    debug: not isDist
  .bundle()
  .pipe(source 'app.js')
  .pipe(do buffer)
  .pipe(gp.if isDist, do gp.uglify)
  .pipe(gulp.dest DESTDIR)

gulp.task 'image', ->
  gulp
    .src("#{IMAGEDIR}/**/*.{jpg,png,gif}")
    .pipe(do gp.imagemin)
    .pipe(gulp.dest DESTDIR)

gulp.task 'bower', ->
  filter = if isDist then /\.(?:otf|eot|svg|ttf|woff)$/ else /.*/
  gulp
    .src(mainBowerFiles filter: filter, includeDev: not isDist)
    .pipe(gulp.dest DESTDIR)

gulp.task 'watch', ['css', 'scripts', 'image', 'bower'], ->
  return if isDist
  gulp
    .watch("#{CSSDIR}/**/*.{sass,scss}", ['css'])
    .on 'change', (ev) ->
      if ev.type == 'deleted'
        delete gp.cached.caches.sass[ev.path]
        gp.remember.forget 'sass', ev.path
  gulp.watch "#{JSDIR}/**/*.{js,coffee}", ['scripts']
  gulp.watch "#{IMAGEDIR}/**/*.{jpg,png,gif}", ['image']

gulp.task 'default', ['css', 'scripts', 'image', 'bower', 'watch']
