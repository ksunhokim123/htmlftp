var webpack = require('webpack');
var path = require('path');

var StatsWriterPlugin = require("webpack-stats-plugin").StatsWriterPlugin;

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
        use: [
          'style-loader',
          'css-loader',
          'sass-loader',
          'import-glob-loader',
        ],
      }
    ]
  },
  plugins: [
    new StatsWriterPlugin({
        filename: "stats.json"
    })
  ]
};
