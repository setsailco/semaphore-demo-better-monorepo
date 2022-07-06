const Fastify = require('fastify')

function buildFastify () {
  const fastify = Fastify({
    logger: true
  })

  fastify.get('/', function (request, reply) {
    reply.send({ message: 'Hello from js-service-1, woo again!' })
  })

  return fastify
}

module.exports = buildFastify
