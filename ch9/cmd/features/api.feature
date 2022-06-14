Feature: Translate API
  Users should be able to submit a word to translate

  Scenario: Translation
    Given the word "hello"
    When I translate it to "german"
    Then the response should be "hallo"
  Scenario: Translation unkown
    Given the word "goodbye"
    When I translate it to "german"
    Then the response should be ""