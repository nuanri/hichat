var path = require('path');
var node_modules_dir = path.resolve(__dirname, 'node_modules');

module.exports = {
    entry: './src/index.js',
    output: {
        path: path.join(__dirname, '/dist/static/js/'),
        filename: 'bundle.js'
    },
    resolve: {
        extensions: ['', '.js', '.jsx']
    },
    module: {
/*      loaders: [{
        test: /\.js[x]?$/,
        exclude: [node_modules_dir],
        loader: 'babel-loader',
        query: {
          //plugins: ['transform-runtime'],
          // stage-0 -> ES7
          presets: ['react', 'es2015', 'stage-0'],
        }
      }]*/
      loaders:[
        {
          test: /\.js[x]?$/,
          exclude: /node_modules/,
          loader: 'babel-loader?presets[]=es2015&presets[]=react',
        },
      ]
    }
}
