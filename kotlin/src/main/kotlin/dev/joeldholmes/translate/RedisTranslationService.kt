package dev.joeldholmes.translate

import io.quarkus.redis.datasource.RedisDataSource
import org.eclipse.microprofile.config.inject.ConfigProperty
import javax.enterprise.context.ApplicationScoped
import javax.inject.Inject

@ApplicationScoped
class RedisTranslationService: ITranslationService {

    @Inject
    private lateinit var redisAPI: RedisDataSource

    @ConfigProperty(name="default.language")
    var defaultLanguage: String? = "english"

    override fun translate(language: String?, word: String) : Translation? {
        val commands = redisAPI?.string(String::class.java)
        val lang = language?.lowercase() ?: defaultLanguage
        val key = "$word:$lang"
        val translation = commands?.get(key)
        return if (translation == null) {
            null
        } else {
            Translation(language = lang, translation= translation)
        }
    }
}