module.exports = {
  entry: ['./main.js'],
  output: {
    filename: 'octy-shopify.js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'babel-loader'
      }
    ]
  }
}
