# chessgo

Practicing what I learned from [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

Using FGDD (see [chesspy](https://github.com/mikepartelow/chesspy/)) to generate meaningful TDD test cases.

Go isn't on my resume yet.

## Famous Game Driven Development

1. Add moves from [The Immortal Game](https://en.wikipedia.org/wiki/Immortal_Game) until the [integration test](https://github.com/mikepartelow/chessgo/blob/main/game_immortal_test.go) fails
2. Comment out the failing integration test case
3. Add [unit tests](https://github.com/mikepartelow/chessgo/blob/main/game_test.go) for the failing integration test case
4. Make the unit tests pass
5. Uncomment the failed integration test case, which should now pass
6. GOTO 1