# instashell

A wise person once said "the slowest part in a query is typing it in".

This is a quick attempt to bring this vision to the shell.

In terms of user experience, this seems promising: seeing results without having to hit "Enter" is satisfying and feels efficient. However, for an optimal experience it would need to be integrated in the shell itself. This would allow integration with existing shell constructs, e.g. aliases, tab completion, globbing, redirection, etc.

## Features

* Input is kept between keystrokes and colorized to indicate a succeeding or failing command
* Control + w deletes back one word (in a manner similar to most shells)
* Control + u deletes the whole line

## Warning

This script may not work properly on Windows.

