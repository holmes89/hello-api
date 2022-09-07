const translation = (translationService) => {
  return async (req, res) => {
    let language = req.query.language || "english";
    const resp = await translationService.translate(language, req.params.word);
    resp
      ? res.json({ language: language.toLowerCase(), translation: resp })
      : res.status(404).send("Missing translation");
  };
};

module.exports = translation;
