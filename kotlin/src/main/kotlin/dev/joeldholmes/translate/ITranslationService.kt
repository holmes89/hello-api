package dev.joeldholmes.translate

interface ITranslationService {
    fun translate(language: String?, word: String) : Translation?
}