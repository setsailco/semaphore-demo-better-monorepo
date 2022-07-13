// Initial release
// Test version bump
// Test version bump 2
// Test version bump 3s
const server = require('./app')()

server.listen({ port: 3000 }, (err, address) => {
  if (err) {
    server.log.error(err)
    process.exit(1)
  }
})
