# Customization - Prompt

The prompt consists of three parts:
 - Prefix contents
 - User input area
 - Postfix contents

The prefix and postfix contents are set by the current `PromptStyle` implementation. Each implementation must provide two functions: 
 - `LeftPrompt(*Gosh, int)`
 - `RightPrompt(*Gosh, int)`

Both functions can expect to be passed the shell and which line should be drawn. If the prompt consists of multiple lines, they start with 0.

The default implementation is in `simplePromptStyle.go`.