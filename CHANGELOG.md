# Alpha v0.0.1

This is an early preview version that showcases basic functionality of `gosh`.
Current features are:
 - Prompt
    * Prompt drawing
    * Key capturing, exit at any time by pressing C-d
    * Basic command evaluation for testing purposes: `cd` changes directory, `gst` runs git status, `exit` closes gosh
 - Runner
    * Creates a map from all executables in the path and their location
 - Debugger
    * Ability to send arbitrary debugging messages to a second window
 - GoshScript
    * Basic Lexer implementation for reading goshscript files
    * Lexer parses numbers of any radix from 2 to 36
    * Lexer parses escape sequences (see [Escape codes](goshscript/STRINGS?id=escape-codes))
 - Plugin
    * Plugins are event-driven, every plugin will be able to implement different functions that will be called by the EventHandler
 - Documentation
    * docsify-based documentation available via rythenglyth.github.io/gosh/
 - Continuous Integration
    * gosh is built automatically on all commits to main and for pull requests