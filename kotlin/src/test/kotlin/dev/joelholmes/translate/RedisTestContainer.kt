package dev.joelholmes.translate

import io.quarkus.test.common.QuarkusTestResourceLifecycleManager
import org.testcontainers.containers.BindMode
import org.testcontainers.containers.GenericContainer
import org.testcontainers.utility.DockerImageName

class RedisTestContainer : QuarkusTestResourceLifecycleManager {

    private val redisContainer = GenericContainer(DockerImageName.parse("redis:latest"))
        .withExposedPorts(6379)
        .withClasspathResourceMapping("data", "/data", BindMode.READ_ONLY)

    override fun start(): MutableMap<String, String> {
        println("STARTING redis ")
        redisContainer.start()
        println("redis://${redisContainer.getHost()}:${redisContainer.getMappedPort(6379)}")
        return mutableMapOf(Pair("quarkus.redis.hosts", "redis://${redisContainer.getHost()}:${redisContainer.getMappedPort(6379)}"))
    }

    override fun stop() {
        println("STOPPING redis")
        redisContainer.stop()
    }
}
