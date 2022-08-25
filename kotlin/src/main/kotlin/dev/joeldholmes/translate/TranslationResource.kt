package dev.joeldholmes.translate

import javax.inject.Inject
import javax.ws.rs.*
import javax.ws.rs.core.MediaType

@Path("/translate")
class TranslationResource {

    @Inject
    private lateinit var service: ITranslationService

    @GET
    @Path("/{word}")
    @Produces(MediaType.APPLICATION_JSON)
    fun translate(
        @PathParam("word") word: String,
        @QueryParam("language") language: String?,
    ) = service.translate(language, word)
}