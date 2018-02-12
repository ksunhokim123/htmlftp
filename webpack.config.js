var {resolve} = require('path');
var webpack =require('webpack');

module.exports = {
    entry: [
        './index.tsx'
    ],
    output:{
        filename: 'bundle.js',
        path: resolve(__dirname, 'dist'),
        publicPath: ''
    },
    resolve:{
        extensions: ['.js', '.jsx', '.ts', '.tsx', '.css']
    },
    context: resolve(__dirname, 'client'),
    devtool: 'inline-source-map',
    devServer:{
        hot: true,
        contentBase: resolve(__dirname, 'dist'),
        publicPath: ''
    },
    module: {
        rules:[{
                test: /\.(ts|tsx)$/,
                use: ['awesome-typescript-loader']
            },{
                 test:/\.(s*)css$/,
                 use:['style-loader','css-loader', 'sass-loader']
            },{
              test: /\.(ttf|eot|woff|woff2)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
              loader: "file-loader?name=fonts/[hash].[ext]",
            },{
              test: /\.(jpe?g|png|gif|svg)$/i,
              loader: "file-loader?name=imgs/[hash].[ext]"
            }
        ]
    },
    plugins: []
};
