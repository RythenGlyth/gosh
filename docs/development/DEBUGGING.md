# Development

Whether you're working on the prompt or a completion function, it can often be useful to be able to print debugging messages at any time.

For this purpose, the `debug` module has been created.
During development, write debugging messages using
`gosh.DebugMessage(int, string)`.

The first parameter specifies a module identifier which can be set for easier readability in the debugging log. It must be greater than 0, these values already exist:

|   |                |
| - | -------------- |
| 1 | Main Loop      |
| 2 | Prompt drawing |
| 3 | Plugin Handler |

Then start a debugging server in a seperate window using
```bash
gosh/src > go run . start-debug
```
And start `gosh` with debugging enabled:
```bash
gosh/src > go run . debug
```
You will see the debugging messages being sent:
![Screenshot of `gosh` with debugging enabled](../assets/DEBUGGING-1.png)


Stop `gosh` as usual and the debugger by pressing `C-c`.
