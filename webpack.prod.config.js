var webpack = require('webpack');
var path = require('path');

const StatsWriterPlugin = require("webpack-stats-plugin").StatsWriterPlugin;
const ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports= {
  context: path.resolve(__dirname, 'client'),
  entry: './index.js',
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/assets/',
    filename: 'bundle.js'
  },
  module : {
    rules: [
      {
        test: /\.js$/i,
        exclude: /node_modules/,
        use: ['babel-loader'],
      },
      {
        test: /\.(s[ac]|c)ss$/i,
        use: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: ['css-loader', 'sass-loader', 'import-glob-loader'],
        }),
      }
    ]
  },
  plugins: [
    new StatsWriterPlugin({
        filename: "stats.json"
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      debug: false,
    }),
    new webpack.optimize.UglifyJsPlugin(),
    new ExtractTextPlugin('bundle.css')
  ]
};
