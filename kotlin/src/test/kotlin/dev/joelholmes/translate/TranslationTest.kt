package dev.joelholmes.translate

import io.quarkus.test.common.QuarkusTestResource
import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured.given
import org.hamcrest.CoreMatchers.equalTo
import org.junit.jupiter.api.Test

@QuarkusTestResource(RedisTestContainer::class)
@QuarkusTest
class TranslationTest {

    @Test
    fun testHelloEndpoint() {
        given()
            .`when`().get("/translate/hello")
            .then()
            .statusCode(200)
            .body("translation", equalTo("Hello"))
            .body("language", equalTo("english"))
    }

    @Test
    fun testHelloEndpointGerman() {
        given()
            .`when`().get("/translate/hello?language=GERMAN")
            .then()
            .statusCode(200)
            .body("translation", equalTo("Hallo"))
            .body("language", equalTo("german"))
    }
}
