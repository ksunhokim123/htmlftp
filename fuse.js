const { FuseBox, CSSPlugin, WebIndexPlugin } = require("fuse-box");

const fuse = FuseBox.init({
  homeDir: "client",
  target : "browser@es5",
  useTypescriptCompiler : true,
  output: "dist/$name.js",
  plugins: [
               CSSPlugin(),
               WebIndexPlugin({
                   template : "client/index.html"
                 })
                 ]
});

fuse.bundle("app")
  .instructions("> index.ts");

fuse.run();
