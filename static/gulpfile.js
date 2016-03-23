var gulp = require('gulp'),
    gwatch = require('gulp-watch'),
    gfilter = require('gulp-filter'),
    uglify = require('gulp-uglify'),
    plumber = require('gulp-plumber'),
     htmlmin = require('gulp-htmlmin'),
    ngAnnotate = require('gulp-ng-annotate'),
    concat = require('gulp-concat');



    // var gulp = require('gulp'),
    // gwatch = require('gulp-watch'),
    // sass = require('gulp-sass'),
    // concat = require('gulp-concat'),
    // uglify = require('gulp-uglify'),
    // htmlmin = require('gulp-htmlmin'),
    // phpmin = require('rgulp-phpmin'),
    // prefix = require("rgulp-prefix"),
    // cache = require('gulp-cached'),
    // gfilter = require('gulp-filter'),
    // //merge = require('merge-stream'),
    // haml = require('gulp-haml'),
    // ngAnnotate = require('gulp-ng-annotate'),
    // autoprefixer = require('gulp-autoprefixer'),
    // autoprefixer_conf = {
    //     //browsers: ['> 1%, last 2 versions, Firefox ESR, Opera 12.1'],
    //     cascade: true
    // }
    // //,extreplace = require('gulp-ext-replace')
    // , plumber = require('gulp-plumber'),
    // templateCache = require('gulp-angular-templatecache');

    //--------------------
    function xwatch(p, fn) {
        (function(p, fn) {
            var cb = function() {
                try {
                    console.log('cb:' + p);
                    var files = gulp.src(p) //
                    .pipe(plumber()).pipe(gfilter(function(file) {
                        return !file.isNull();
                    }));
                    fn(files);
                } catch (e) {
                    console.log('err');
                    console.log(e);
                }
            };
            var loopfn = function() {
                console.log('loopfn:' + p);
                var st = gwatch(p, function(events, done) {
                    cb();
                });
                st.on('error', function(err) {
                    console.log(err);
                    st.close();
                    setTimeout(loopfn, 100);
                    console.log('setTimeout:' + p);
                });
                cb();
            }
            loopfn();
        })(p, fn);
        //cb();1
        // fn(gulp.src(p));
        // watch({glob:p,verbose:true,emit:'all'},function(){
        //    fn(gulp.src(p)); 
        // });
    }
    //---------------
gulp.task('default', function() {
    xwatch([
        'theme/assets/global/plugins/jquery.min.js',
        'theme/assets/global/plugins/jquery-migrate.min.js',
        'theme/assets/global/plugins/bootstrap/js/bootstrap.min.js'
        ], function(files) {
        return files //
        //
        .pipe(concat("all.js")) //   
        //.pipe(gulp.dest(dev + '/js')) //  
      //  .pipe(ngAnnotate())
      //.pipe(uglify()) //      
       // .pipe(prefix(copyright)) //
        .pipe(gulp.dest('./dist')); //
    });
    //---------------------------
    xwatch([
        'js/controllers/**/*.js'

        ], function(files) {
        return files //
        //
        .pipe(concat("ctls.js")) //   
        //.pipe(gulp.dest(dev + '/js')) //  
        .pipe(ngAnnotate())
      .pipe(uglify()) //      
       // .pipe(prefix(copyright)) //
        .pipe(gulp.dest('./dist')); //
    });
    //----------------------------------
    xwatch(['views/**/*.html'], function(files) {
        //console.log('aaaa');

        files //
        .pipe(htmlmin({
            collapseWhitespace: true,
            removeComments: true,
            removeAttributeQuotes: false,
        })) // 
        // .pipe(templateCache({
        //     root: '/partials/' //
        //     ,
        //     standalone: true
        // }))
        .pipe(gulp.dest( 'dist/views/')) //
        ; //
    });
    //----

});