package dev.joeldholmes.translate

import javax.inject.Inject
import javax.ws.rs.GET
import javax.ws.rs.Path
import javax.ws.rs.PathParam
import javax.ws.rs.Produces
import javax.ws.rs.QueryParam
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
        @QueryParam("language") language: String?
    ) = service.translate(language, word)
}
